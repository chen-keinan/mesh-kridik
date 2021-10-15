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
	assert.Equal(t, bFiles[0].Name, common.IstioMutualmTLS)
	fm := utils.NewKFolder()
	err = utils.CreateSecurityFolderIfNotExist("mesh", "istio", fm)
	assert.NoError(t, err)
	// save benchmark files to folder
	err = SaveSecurityFilesIfNotExist("mesh", "istio", bFiles)
	assert.NoError(t, err)
	// fetch files from benchmark folder
	bFiles, err = utils.GetMeshSecurityChecksFiles("mesh", "istio", fm)
	assert.Equal(t, bFiles[0].Name, common.AllowMtlsPermissiveMode)
	assert.Equal(t, bFiles[1].Name, common.AllowWithPositiveMatchingRulesFrom)
	assert.Equal(t, bFiles[2].Name, common.AllowWithPositiveMatchingRulesTo)
	assert.Equal(t, bFiles[3].Name, common.IstioMutualmTLS)
	assert.Equal(t, bFiles[4].Name, common.SaferAuthorizationPolicyPatterns)
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
	folder, err2 := utils.GetSecurityFolder("lxd", "v1.6.0", fm)
	assert.NoError(t, err2)
	err := os.RemoveAll(folder)
	assert.NoError(t, err)
	err = os.RemoveAll(utils.GetHomeFolder())
	assert.NoError(t, err)
}
