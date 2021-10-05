package startup

import (
	"github.com/chen-keinan/kube-mesh-kridik/internal/common"
	"github.com/chen-keinan/kube-mesh-kridik/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//Test_CreateLxdBenchmarkFilesIfNotExist test
func Test_CreateLxdBenchmarkFilesIfNotExist(t *testing.T) {
	bFiles, err := GenerateLxdBenchmarkFiles()
	if err != nil {
		t.Fatal(err)
	}
	// generate test with packr
	assert.Equal(t, bFiles[0].Name, common.FilesystemConfiguration)
	assert.Equal(t, bFiles[1].Name, common.ConfigureSoftwareUpdates)
	assert.Equal(t, bFiles[2].Name, common.ConfigureSudo)
	assert.Equal(t, bFiles[3].Name, common.FilesystemIntegrityChecking)
	assert.Equal(t, bFiles[4].Name, common.AdditionalProcessHardening)
	assert.Equal(t, bFiles[5].Name, common.MandatoryAccessControl)
	assert.Equal(t, bFiles[6].Name, common.WarningBanners)
	assert.Equal(t, bFiles[7].Name, common.EnsureUpdates)
	assert.Equal(t, bFiles[8].Name, common.InetdServices)
	assert.Equal(t, bFiles[9].Name, common.SpecialPurposeServices)
	assert.Equal(t, bFiles[10].Name, common.ServiceClients)
	assert.Equal(t, bFiles[11].Name, common.NonessentialServices)
	assert.Equal(t, bFiles[12].Name, common.NetworkParameters)
	assert.Equal(t, bFiles[13].Name, common.NetworkParametersHost)
	assert.Equal(t, bFiles[14].Name, common.TCPWrappers)
	assert.Equal(t, bFiles[15].Name, common.FirewallConfiguration)
	assert.Equal(t, bFiles[16].Name, common.ConfigureLogging)
	assert.Equal(t, bFiles[17].Name, common.EnsureLogrotateConfigured)
	assert.Equal(t, bFiles[18].Name, common.EnsureLogrotateAssignsAppropriatePermissions)
	assert.Equal(t, bFiles[19].Name, common.ConfigureCron)
	assert.Equal(t, bFiles[20].Name, common.SSHServerConfiguration)
	assert.Equal(t, bFiles[21].Name, common.ConfigurePam)
	assert.Equal(t, bFiles[22].Name, common.UserAccountsAndEnvironment)
	assert.Equal(t, bFiles[23].Name, common.RootLoginRestrictedSystemConsole)
	assert.Equal(t, bFiles[24].Name, common.EnsureAccessSuCommandRestricted)
	assert.Equal(t, bFiles[25].Name, common.SystemFilePermissions)
	assert.Equal(t, bFiles[26].Name, common.UserAndGroupSettings)
	fm := utils.NewKFolder()
	err = utils.CreateBenchmarkFolderIfNotExist("lxd", "v1.0.0", fm)
	assert.NoError(t, err)
	// save benchmark files to folder
	err = SaveBenchmarkFilesIfNotExist("lxd", "v1.0.0", bFiles)
	assert.NoError(t, err)
	// fetch files from benchmark folder
	bFiles, err = utils.GetLxdBenchAuditFiles("lxd", "v1.0.0", fm)
	assert.Equal(t, bFiles[0].Name, common.FilesystemConfiguration)
	assert.Equal(t, bFiles[1].Name, common.ConfigureSoftwareUpdates)
	assert.Equal(t, bFiles[2].Name, common.ConfigureSudo)
	assert.Equal(t, bFiles[3].Name, common.FilesystemIntegrityChecking)
	assert.Equal(t, bFiles[4].Name, common.AdditionalProcessHardening)
	assert.Equal(t, bFiles[5].Name, common.MandatoryAccessControl)
	assert.Equal(t, bFiles[6].Name, common.WarningBanners)
	assert.Equal(t, bFiles[7].Name, common.EnsureUpdates)
	assert.Equal(t, bFiles[8].Name, common.InetdServices)
	assert.Equal(t, bFiles[9].Name, common.SpecialPurposeServices)
	assert.Equal(t, bFiles[10].Name, common.ServiceClients)
	assert.Equal(t, bFiles[11].Name, common.NonessentialServices)
	assert.Equal(t, bFiles[12].Name, common.NetworkParameters)
	assert.Equal(t, bFiles[13].Name, common.NetworkParametersHost)
	assert.Equal(t, bFiles[14].Name, common.TCPWrappers)
	assert.Equal(t, bFiles[15].Name, common.FirewallConfiguration)
	assert.Equal(t, bFiles[16].Name, common.ConfigureLogging)
	assert.Equal(t, bFiles[17].Name, common.EnsureLogrotateConfigured)
	assert.Equal(t, bFiles[18].Name, common.EnsureLogrotateAssignsAppropriatePermissions)
	assert.Equal(t, bFiles[19].Name, common.ConfigureCron)
	assert.Equal(t, bFiles[20].Name, common.SSHServerConfiguration)
	assert.Equal(t, bFiles[21].Name, common.ConfigurePam)
	assert.Equal(t, bFiles[22].Name, common.UserAccountsAndEnvironment)
	assert.Equal(t, bFiles[23].Name, common.RootLoginRestrictedSystemConsole)
	assert.Equal(t, bFiles[24].Name, common.EnsureAccessSuCommandRestricted)
	assert.Equal(t, bFiles[25].Name, common.SystemFilePermissions)
	assert.Equal(t, bFiles[26].Name, common.UserAndGroupSettings)
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
	folder, err2 := utils.GetBenchmarkFolder("lxd", "v1.6.0", fm)
	assert.NoError(t, err2)
	err := os.RemoveAll(folder)
	assert.NoError(t, err)
	err = os.RemoveAll(utils.GetHomeFolder())
	assert.NoError(t, err)
}
