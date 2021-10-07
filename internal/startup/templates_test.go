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
	assert.Equal(t, bFiles[0].Name, common.IstioSecurityChecks)
	fm := utils.NewKFolder()
	err = utils.CreateSecurityFolderIfNotExist("mesh", "v1.0.0", fm)
	assert.NoError(t, err)
	// save benchmark files to folder
	err = SaveSecurityFilesIfNotExist("mesh", "v1.0.0", bFiles)
	assert.NoError(t, err)
	// fetch files from benchmark folder
	bFiles, err = utils.GetMeshSecurityChecksFiles("mesh", "v1.0.0", fm)
	assert.Equal(t, bFiles[0].Name, common.IstioSecurityChecks)
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
