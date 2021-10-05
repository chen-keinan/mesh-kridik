package utils

import (
	"fmt"
	"github.com/chen-keinan/kube-mesh-kridik/internal/common"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

//PluginSourceSubFolder plugin source folder
const PluginSourceSubFolder = "plugins/source"

//CompilePluginSubFolder plugins complied folder
const CompilePluginSubFolder = "plugins/compile"

//FolderMgr defines the interface for kube-knark folder
//fileutil.go
//go:generate mockgen -destination=./mocks/mock_FolderMgr.go -package=mocks . FolderMgr
type FolderMgr interface {
	CreateFolder(folderName string) error
	GetHomeFolder() (string, error)
}

//bFolder kube-mesh-kridik folder object
type bFolder struct {
}

//NewKFolder return bFolder instance
func NewKFolder() FolderMgr {
	return &bFolder{}
}

//CreateFolder create new kube-mesh-kridik folder
func (lxdf bFolder) CreateFolder(folderName string) error {
	_, err := os.Stat(folderName)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(folderName, 0750)
		if errDir != nil {
			return err
		}
	}
	return nil
}

//GetHomeFolder return kube-mesh-kridik home folder
func (lxdf bFolder) GetHomeFolder() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	// User can set a custom KUBE_KNARK_HOME from environment variable
	usrHome := GetEnv(common.LxdProbeHomeEnvVar, usr.HomeDir)
	return path.Join(usrHome, ".kube-mesh-kridik"), nil
}

//GetPluginSourceSubFolder return plugins source folder path
func GetPluginSourceSubFolder(fm FolderMgr) (string, error) {
	folder, err := fm.GetHomeFolder()
	if err != nil {
		return "", err
	}
	return path.Join(folder, PluginSourceSubFolder), nil
}

//GetCompilePluginSubFolder return plugin compiled folder path
func GetCompilePluginSubFolder(fm FolderMgr) (string, error) {
	folder, err := fm.GetHomeFolder()
	if err != nil {
		return "", err
	}
	return path.Join(folder, CompilePluginSubFolder), nil
}

//CreatePluginsCompiledFolderIfNotExist create plugins compiled folder if not exist
func CreatePluginsCompiledFolderIfNotExist(fm FolderMgr) error {
	ebpfFolder, err := GetCompilePluginSubFolder(fm)
	if err != nil {
		return err
	}
	return fm.CreateFolder(ebpfFolder)
}

//CreatePluginsSourceFolderIfNotExist plugins source folder if not exist
func CreatePluginsSourceFolderIfNotExist(fm FolderMgr) error {
	pluginfFolder, err := GetPluginSourceSubFolder(fm)
	if err != nil {
		return err
	}
	return fm.CreateFolder(pluginfFolder)
}

//GetHomeFolder return kube-mesh-kridik home folder
func GetHomeFolder() string {
	usr, err := user.Current()
	if err != nil {
		panic("Failed to fetch user home folder")
	}
	// User can set a custom LXD_PROBE_HOME from environment variable
	usrHome := GetEnv(common.LxdProbeHomeEnvVar, usr.HomeDir)
	return path.Join(usrHome, ".kube-mesh-kridik")
}

//CreateHomeFolderIfNotExist create kube-mesh-kridik home folder if not exist
func CreateHomeFolderIfNotExist(fm FolderMgr) error {
	lxdProbeFolder, err := fm.GetHomeFolder()
	if err != nil {
		return err
	}
	_, err = os.Stat(lxdProbeFolder)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(lxdProbeFolder, 0750)
		if errDir != nil {
			return fmt.Errorf("failed to create kube-mesh-kridik home folder at %s", lxdProbeFolder)
		}
	}
	return nil
}

//GetBenchmarkFolder return benchmark folder
func GetBenchmarkFolder(spec, version string, fm FolderMgr) (string, error) {
	folder, err := fm.GetHomeFolder()
	if err != nil {
		return "", err
	}
	return filepath.Join(folder, fmt.Sprintf("benchmarks/%s/%s/", spec, version)), nil
}

//CreateBenchmarkFolderIfNotExist create kube-mesh-kridik benchmark folder if not exist
func CreateBenchmarkFolderIfNotExist(spec, version string, fm FolderMgr) error {
	benchmarkFolder, err := GetBenchmarkFolder(spec, version, fm)
	if err != nil {
		return err
	}
	return fm.CreateFolder(benchmarkFolder)
}

//GetLxdBenchAuditFiles return lxd benchmark file
func GetLxdBenchAuditFiles(spec, version string, fm FolderMgr) ([]FilesInfo, error) {
	filesData := make([]FilesInfo, 0)
	folder, err := GetBenchmarkFolder(spec, version, fm)
	if err != nil {
		return filesData, err
	}
	filesInfo, err := ioutil.ReadDir(filepath.Join(folder))
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range filesInfo {
		filePath := filepath.Join(folder, filepath.Clean(fileInfo.Name()))
		fData, err := ioutil.ReadFile(filepath.Clean(filePath))
		if err != nil {
			return nil, err
		}
		filesData = append(filesData, FilesInfo{fileInfo.Name(), string(fData)})
	}
	return filesData, nil
}

//FilesInfo file data
type FilesInfo struct {
	Name string
	Data string
}

//GetEnv Get Environment Variable value or return default
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
