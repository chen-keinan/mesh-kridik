package startup

import (
	"fmt"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/gobuffalo/packr"
	"os"
	"path/filepath"
)

//GenerateLxdBenchmarkFiles use packr to load benchmark audit test yaml
//nolint:gocyclo
func GenerateLxdBenchmarkFiles() ([]utils.FilesInfo, error) {
	fileInfo := make([]utils.FilesInfo, 0)
	box := packr.NewBox("./../benchmark/lxd/v1.0.0/")
	// Add Master Node Configuration tests
	//1
	mnc, err := box.FindString(common.FilesystemConfiguration)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s  %s", common.FilesystemConfiguration, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.FilesystemConfiguration, Data: mnc})
	//2
	su, err := box.FindString(common.ConfigureSoftwareUpdates)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.ConfigureSoftwareUpdates, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ConfigureSoftwareUpdates, Data: su})
	//3
	cs, err := box.FindString(common.ConfigureSudo)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.ConfigureSudo, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ConfigureSudo, Data: cs})
	//4
	ic, err := box.FindString(common.FilesystemIntegrityChecking)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.FilesystemIntegrityChecking, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.FilesystemIntegrityChecking, Data: ic})
	//5
	ah, err := box.FindString(common.AdditionalProcessHardening)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.AdditionalProcessHardening, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.AdditionalProcessHardening, Data: ah})
	//6
	mac, err := box.FindString(common.MandatoryAccessControl)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.MandatoryAccessControl, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.MandatoryAccessControl, Data: mac})
	//7
	wb, err := box.FindString(common.WarningBanners)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.WarningBanners, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.WarningBanners, Data: wb})
	//8
	eu, err := box.FindString(common.EnsureUpdates)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.EnsureUpdates, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.EnsureUpdates, Data: eu})
	//9
	is, err := box.FindString(common.InetdServices)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.InetdServices, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.InetdServices, Data: is})
	//10
	sps, err := box.FindString(common.SpecialPurposeServices)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.SpecialPurposeServices, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.SpecialPurposeServices, Data: sps})
	//11
	sci, err := box.FindString(common.ServiceClients)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.ServiceClients, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ServiceClients, Data: sci})
	//12
	nes, err := box.FindString(common.NonessentialServices)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.NonessentialServices, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.NonessentialServices, Data: nes})
	//13
	np, err := box.FindString(common.NetworkParameters)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.NetworkParameters, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.NetworkParameters, Data: np})
	//14
	nps, err := box.FindString(common.NetworkParametersHost)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.NetworkParametersHost, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.NetworkParametersHost, Data: nps})
	//15
	tw, err := box.FindString(common.TCPWrappers)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.TCPWrappers, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.TCPWrappers, Data: tw})
	//15
	fc, err := box.FindString(common.FirewallConfiguration)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.FirewallConfiguration, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.FirewallConfiguration, Data: fc})
	//16
	cf, err := box.FindString(common.ConfigureLogging)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.ConfigureLogging, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ConfigureLogging, Data: cf})
	//16
	elc, err := box.FindString(common.EnsureLogrotateConfigured)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.EnsureLogrotateConfigured, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.EnsureLogrotateConfigured, Data: elc})
	//17
	elap, err := box.FindString(common.EnsureLogrotateAssignsAppropriatePermissions)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.EnsureLogrotateAssignsAppropriatePermissions, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.EnsureLogrotateAssignsAppropriatePermissions, Data: elap})
	//18
	cc, err := box.FindString(common.ConfigureCron)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.ConfigureCron, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ConfigureCron, Data: cc})
	//19
	ssc, err := box.FindString(common.SSHServerConfiguration)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.SSHServerConfiguration, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.SSHServerConfiguration, Data: ssc})
	//20
	cp, err := box.FindString(common.ConfigurePam)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.ConfigurePam, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ConfigurePam, Data: cp})
	//21
	uaae, err := box.FindString(common.UserAccountsAndEnvironment)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.UserAccountsAndEnvironment, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.UserAccountsAndEnvironment, Data: uaae})
	//22
	rrs, err := box.FindString(common.RootLoginRestrictedSystemConsole)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.RootLoginRestrictedSystemConsole, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.RootLoginRestrictedSystemConsole, Data: rrs})
	//23
	easc, err := box.FindString(common.EnsureAccessSuCommandRestricted)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.EnsureAccessSuCommandRestricted, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.EnsureAccessSuCommandRestricted, Data: easc})
	//24
	sfp, err := box.FindString(common.SystemFilePermissions)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.SystemFilePermissions, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.SystemFilePermissions, Data: sfp})
	//25
	ugs, err := box.FindString(common.UserAndGroupSettings)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load lxd benchmarks audit tests %s %s", common.UserAndGroupSettings, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.UserAndGroupSettings, Data: ugs})
	return fileInfo, nil
}

//GetHelpSynopsis get help synopsis file
func GetHelpSynopsis() string {
	box := packr.NewBox("./../cli/commands/help/")
	// Add Master Node Configuration tests
	hs, err := box.FindString(common.Synopsis)
	if err != nil {
		panic(fmt.Sprintf("faild to load cli help synopsis %s", err.Error()))
	}
	return hs
}

//SaveBenchmarkFilesIfNotExist create benchmark audit file if not exist
func SaveBenchmarkFilesIfNotExist(spec, version string, filesData []utils.FilesInfo) error {
	fm := utils.NewKFolder()
	folder, err := utils.GetBenchmarkFolder(spec, version, fm)
	if err != nil {
		return err
	}
	for _, fileData := range filesData {
		filePath := filepath.Join(folder, fileData.Name)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			f, err := os.Create(filePath)
			if err != nil {
				return fmt.Errorf(err.Error())
			}
			_, err = f.WriteString(fileData.Data)
			if err != nil {
				return fmt.Errorf("failed to write benchmark file")
			}
			err = f.Close()
			if err != nil {
				return fmt.Errorf("faild to close file %s", filePath)
			}
		}
	}
	return nil
}
