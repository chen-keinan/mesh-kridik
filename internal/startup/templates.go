package startup

import (
	"fmt"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/gobuffalo/packr"
	"os"
	"path/filepath"
)

//GenerateMeshSecurityFiles use packr to load benchmark audit test yaml
//nolint:gocyclo
func GenerateMeshSecurityFiles() ([]utils.FilesInfo, error) {
	fileInfo := make([]utils.FilesInfo, 0)
	box := packr.NewBox("./../security/mesh/istio/")
	// Add Master Node Configuration tests
	//1
	mnc, err := box.FindString(common.IstioMutualmTLS)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.IstioMutualmTLS, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.IstioMutualmTLS, Data: mnc})
	dmpm, err := box.FindString(common.AllowMtlsPermissiveMode)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.AllowMtlsPermissiveMode, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.AllowMtlsPermissiveMode, Data: dmpm})
	//2
	sap, err := box.FindString(common.SaferAuthorizationPolicyPatterns)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.SaferAuthorizationPolicyPatterns, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.SaferAuthorizationPolicyPatterns, Data: sap})
	//3
	apm, err := box.FindString(common.AllowWithPositiveMatchingRulesTo)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.AllowWithPositiveMatchingRulesTo, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.AllowWithPositiveMatchingRulesTo, Data: apm})
	//4
	apmf, err := box.FindString(common.AllowWithPositiveMatchingRulesFrom)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.AllowWithPositiveMatchingRulesFrom, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.AllowWithPositiveMatchingRulesFrom, Data: apmf})
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

//SaveSecurityFilesIfNotExist create benchmark audit file if not exist
func SaveSecurityFilesIfNotExist(spec, version string, filesData []utils.FilesInfo) error {
	fm := utils.NewKFolder()
	folder, err := utils.GetSecurityFolder(spec, version, fm)
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
				return fmt.Errorf("failed to write security file")
			}
			err = f.Close()
			if err != nil {
				return fmt.Errorf("faild to close file %s", filePath)
			}
		}
	}
	return nil
}
