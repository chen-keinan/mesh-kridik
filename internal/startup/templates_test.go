package startup

import (
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//Test_CreateMeshSecurityFilesIfNotExist test
func Test_CreateMeshSecurityFilesIfNotExist(t *testing.T) {
	bFiles, err := GenerateMeshSecurityFiles()
	if err != nil {
		t.Fatal(err)
	}
	// generate test with packr
	fm := utils.NewKFolder()
	err = utils.CreateSecurityFolderIfNotExist("mesh", "istio", fm)
	assert.NoError(t, err)
	// save benchmark files to folder
	err = SaveSecurityFilesIfNotExist("mesh", "istio", bFiles)
	assert.NoError(t, err)
	// fetch files from benchmark folder
	bFiles, err = utils.GetMeshSecurityChecksFiles("mesh", "istio", fm)
	assert.Equal(t, bFiles[0].Name, common.UnderstandTrafficCaptureLimitations)
	assert.Equal(t, bFiles[1].Name, common.IstioMutualmTLS)
	assert.Equal(t, bFiles[2].Name, common.SaferAuthorizationPolicyPatterns)
	assert.Equal(t, bFiles[3].Name, common.TLSOriginationForEgressTraffic)
	assert.Equal(t, bFiles[4].Name, common.ProtocolDetection)
	assert.Equal(t, bFiles[5].Name, common.Cni)
	assert.Equal(t, bFiles[6].Name, common.Gateway)
	assert.Equal(t, bFiles[7].Name, common.ConfigureLimitDownstreamConnections)
	assert.Equal(t, bFiles[8].Name, common.ConfigureThirdPartyServiceAccountTokens)
	assert.Equal(t, bFiles[9].Name, common.ControlPlane)
	assert.Equal(t, bFiles[10].Name, common.AllowMtlsPermissiveMode)
	assert.Equal(t, bFiles[11].Name, common.AvoidOverlyBroadHostsConfigurations)
	assert.Equal(t, bFiles[12].Name, common.ClosePort15010UnauthenticatePlaintext)
	assert.Equal(t, bFiles[13].Name, common.ClosePort8008UnauthenticatePlaintext)
	assert.Equal(t, bFiles[14].Name, common.DestinationRulePerformTLSOrigination)
	assert.Equal(t, bFiles[15].Name, common.DetectByProtocol)
	assert.Equal(t, bFiles[16].Name, common.DownstreamConnectionLimitConfigMap)
	assert.Equal(t, bFiles[17].Name, common.IngressGatewayPatchedDownstreamConnectionLimit)
	assert.Equal(t, bFiles[18].Name, common.IstioUsing3rdPartyTokens)
	assert.Equal(t, bFiles[19].Name, common.PathNormalizationInAuthorization)
	assert.Equal(t, bFiles[20].Name, common.PodCapabilitiesExist)
	assert.Equal(t, bFiles[21].Name, common.ProxyBlocksExternalHostWithinMesh)
	assert.Equal(t, bFiles[22].Name, common.RestrictGatewayCreationPrivileges)
	assert.Equal(t, bFiles[23].Name, common.SaferAuthorizationPolicyPatternsPolicy)

	assert.NoError(t, err)
	err = os.RemoveAll(utils.GetHomeFolder())
	assert.NoError(t, err)
}

//Test_GetHelpSynopsis test
func Test_GetHelpSynopsis(t *testing.T) {
	hs := GetHelpSynopsis()
	assert.True(t, len(hs) != 0)
}

//Test_SaveBenchmarkFilesIfNotExist test
func Test_SaveBenchmarkFilesIfNotExist(t *testing.T) {
	fm := utils.NewKFolder()
	folder, err2 := utils.GetSecurityFolder("mesh", "v1.6.0", fm)
	assert.NoError(t, err2)
	err := os.RemoveAll(folder)
	assert.NoError(t, err)
	err = os.RemoveAll(utils.GetHomeFolder())
	assert.NoError(t, err)
}
