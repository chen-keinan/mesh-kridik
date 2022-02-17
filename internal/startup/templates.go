package startup

import (
	"fmt"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/internal/security/mesh"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/gobuffalo/packr"
	"os"
	"path/filepath"
)

//GenerateMeshSecurityFiles use packr to load benchmark audit test yaml
//nolint:gocyclo
func GenerateMeshSecurityFiles() ([]utils.FilesInfo, error) {
	return mesh.LoadIstioSpecs()
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
