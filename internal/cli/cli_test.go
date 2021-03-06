package cli

import (
	"github.com/chen-keinan/go-command-eval/eval"
	"github.com/chen-keinan/mesh-kridik/internal/cli/commands"
	"github.com/chen-keinan/mesh-kridik/internal/cli/mocks"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	m3 "github.com/chen-keinan/mesh-kridik/internal/mocks"
	"github.com/chen-keinan/mesh-kridik/internal/models"
	m2 "github.com/chen-keinan/mesh-kridik/pkg/models"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

//Test_StartCli tests
func Test_StartCli(t *testing.T) {
	fm := utils.NewKFolder()
	initSecurityChecksData(fm, ArgsData{SpecType: "mesh", SpecVersion: "istio"})
	files, err := utils.GetMeshSecurityChecksFiles("mesh", "istio", fm)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, len(files), 24)
	assert.Equal(t, files[0].Name, common.UnderstandTrafficCaptureLimitations)
	assert.Equal(t, files[1].Name, common.IstioMutualmTLS)
	assert.Equal(t, files[2].Name, common.SaferAuthorizationPolicyPatterns)
	assert.Equal(t, files[3].Name, common.TLSOriginationForEgressTraffic)
	assert.Equal(t, files[4].Name, common.ProtocolDetection)
	assert.Equal(t, files[5].Name, common.Cni)
	assert.Equal(t, files[6].Name, common.Gateway)
	assert.Equal(t, files[7].Name, common.ConfigureLimitDownstreamConnections)
	assert.Equal(t, files[8].Name, common.ConfigureThirdPartyServiceAccountTokens)
	assert.Equal(t, files[9].Name, common.ControlPlane)
	assert.Equal(t, files[10].Name, common.AllowMtlsPermissiveMode)
	assert.Equal(t, files[11].Name, common.AvoidOverlyBroadHostsConfigurations)
	assert.Equal(t, files[12].Name, common.ClosePort15010UnauthenticatePlaintext)
	assert.Equal(t, files[13].Name, common.ClosePort8008UnauthenticatePlaintext)
	assert.Equal(t, files[14].Name, common.DestinationRulePerformTLSOrigination)
	assert.Equal(t, files[15].Name, common.DetectByProtocol)
	assert.Equal(t, files[16].Name, common.DownstreamConnectionLimitConfigMap)
	assert.Equal(t, files[17].Name, common.IngressGatewayPatchedDownstreamConnectionLimit)
	assert.Equal(t, files[18].Name, common.IstioUsing3rdPartyTokens)
	assert.Equal(t, files[19].Name, common.PathNormalizationInAuthorization)
	assert.Equal(t, files[20].Name, common.PodCapabilitiesExist)
	assert.Equal(t, files[21].Name, common.ProxyBlocksExternalHostWithinMesh)
	assert.Equal(t, files[22].Name, common.RestrictGatewayCreationPrivileges)
	assert.Equal(t, files[23].Name, common.SaferAuthorizationPolicyPatternsPolicy)
}

func Test_ArgsSanitizer(t *testing.T) {
	args := []string{"--a", "-b"}
	ad := ArgsSanitizer(args)
	assert.Equal(t, ad.Filters[0], "a")
	assert.Equal(t, ad.Filters[1], "b")
	assert.False(t, ad.Help)
	args = []string{}
	ad = ArgsSanitizer(args)
	assert.True(t, ad.Filters[0] == "")
	args = []string{"--help"}
	ad = ArgsSanitizer(args)
	assert.True(t, ad.Help)
}

//Test_MeshKridikHelpFunc test
func Test_MeshKridikHelpFunc(t *testing.T) {
	cm := make(map[string]cli.CommandFactory)
	bhf := MeshKridikHelpFunc(common.MeshKridik)
	helpFile := bhf(cm)
	assert.True(t, strings.Contains(helpFile, "Available commands are:"))
	assert.True(t, strings.Contains(helpFile, "Usage: mesh-kridik [--version] [--help] <command> [<args>]"))
}

//Test_createCliBuilderData test
func Test_createCliBuilderData(t *testing.T) {
	cmdArgs := []string{"a"}
	ad := ArgsSanitizer(os.Args[1:])
	cmdArgs = append(cmdArgs, ad.Filters...)
	cmds := make([]cli.Command, 0)
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	// invoke cli
	cmds = append(cmds, commands.NewMeshCheck(ad.Filters, plChan, completedChan, []utils.FilesInfo{}, eval.NewEvalCmd()))
	c := createCliBuilderData(cmdArgs, cmds)
	_, ok := c["a"]
	assert.True(t, ok)

}

//Test_InvokeCli test
func Test_InvokeCli(t *testing.T) {
	const policy = `package example
	default deny = false
	allow {
		some i
		input.kind == "Pod"
		image := input.spec.containers[i].image
		not startswith(image, "kalpine")
		}`
	ab := &models.SecurityCheck{}
	ab.CheckCommand = []string{"aaa"}
	ab.EvalExpr = "'${0}' != '';&& [${1} MATCH no_permission.rego QUERY example.allow RETURN allow]"
	ab.CommandParams = map[int][]string{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	evalCmd := mocks.NewMockCmdEvaluator(ctrl)
	evalCmd.EXPECT().EvalCommandPolicy([]string{"aaa"}, ab.EvalExpr, policy).Return(eval.CmdEvalResult{Match: true}).Times(1)
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	tl := m3.NewMockTestLoader(ctrl)
	infos := []utils.FilesInfo{{Name: "no_permission.rego", Data: policy}}
	tl.EXPECT().LoadSecurityChecks(infos).Return([]*models.SubCategory{{Name: "te", Checks: []*models.SecurityCheck{ab}}}, nil)
	go func() {
		<-plChan
		completedChan <- true
	}()
	kb := &commands.MeshCheck{FilesInfo: infos, Evaluator: evalCmd, ResultProcessor: commands.GetResultProcessingFunction([]string{}), FileLoader: tl, OutputGenerator: commands.ConsoleOutputGenerator, PlChan: plChan, CompletedChan: completedChan}
	cmdArgs := []string{"a"}
	cmds := make([]cli.Command, 0)
	// invoke cli
	cmds = append(cmds, kb)
	c := createCliBuilderData(cmdArgs, cmds)
	a, err := invokeCommandCli(cmdArgs, c)
	assert.NoError(t, err)
	assert.True(t, a == 0)
}

func Test_InitPluginFolder(t *testing.T) {
	fm := utils.NewKFolder()
	initPluginFolders(fm)
}

func Test_InitPluginWorker(t *testing.T) {
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	go func() {
		plChan <- m2.MeshCheckResults{}
		completedChan <- true
	}()
	initPluginWorker(plChan, completedChan)

}
