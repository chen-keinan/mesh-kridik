package utils

import (
	"fmt"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/pkg/utils/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

//Test_GetHomeFolder test
func Test_GetHomeFolder(t *testing.T) {
	a := GetHomeFolder()
	assert.True(t, strings.HasSuffix(a, ".mesh-kridik"))
}

//Test_CreateHomeFolderIfNotExist test
func Test_CreateHomeFolderIfNotExist(t *testing.T) {
	fm := NewKFolder()
	err := CreateHomeFolderIfNotExist(fm)
	assert.NoError(t, err)
	_, err = os.Stat(GetHomeFolder())
	if os.IsNotExist(err) {
		t.Fatal()
	}
	err = os.RemoveAll(GetHomeFolder())
	assert.NoError(t, err)
}

//Test_GetSecurityFolder test
func Test_GetSecurityFolder(t *testing.T) {
	fm := NewKFolder()
	err := CreateHomeFolderIfNotExist(fm)
	assert.NoError(t, err)
	a, err := GetSecurityFolder("mesh", "istio", fm)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(a, ".mesh-kridik/security/mesh/istio"))
}

//Test_SecurityFolderIfNotExist test
func Test_SecurityFolderIfNotExist(t *testing.T) {
	fm := NewKFolder()
	err := CreateSecurityFolderIfNotExist("mesh", "istio", fm)
	assert.NoError(t, err)
	folder, err := GetSecurityFolder("mesh", "istio", fm)
	assert.NoError(t, err)
	_, err = os.Stat(folder)
	if os.IsNotExist(err) {
		t.Fatal()
	}
	err = os.RemoveAll(folder)
	assert.NoError(t, err)
}

//Test_GetMeshSecureCheckFiles test
func Test_GetMeshSecureCheckFiles(t *testing.T) {
	fm := NewKFolder()
	err := CreateHomeFolderIfNotExist(fm)
	if err != nil {
		t.Fatal(err)
	}
	err = CreateSecurityFolderIfNotExist("mesh", "istio", fm)
	if err != nil {
		t.Fatal(err)
	}
	err = saveFilesIfNotExist([]FilesInfo{{Name: "aaa", Data: "bbb"}, {Name: "ddd", Data: "ccc"}})
	if err != nil {
		t.Fatal(err)
	}
	f, err := GetMeshSecurityChecksFiles("mesh", "istio", fm)
	if err != nil {
		t.Fatal(err)
	}
	folder, err := GetSecurityFolder("mesh", "istio", fm)
	assert.NoError(t, err)
	err = os.RemoveAll(folder)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, f[0].Name, "aaa")
	assert.Equal(t, f[1].Name, "ddd")

}

//Test_GetMeshSecurityChecksNoFolder test
func Test_GetMeshSecurityChecksFolder(t *testing.T) {
	fm := NewKFolder()
	_, err := GetMeshSecurityChecksFiles("mesh", "istio", fm)
	assert.True(t, err != nil)
}

func saveFilesIfNotExist(filesData []FilesInfo) error {
	fm := NewKFolder()
	folder, err := GetSecurityFolder("mesh", "istio", fm)
	if err != nil {
		return err
	}
	for _, fileData := range filesData {
		filePath := filepath.Join(folder, fileData.Name)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			f, err := os.Create(filePath)
			if err != nil {
				panic(err)
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

//Test_GetEnv test getting home mesh-kridik folder
func Test_GetEnv(t *testing.T) {
	os.Setenv(common.MeshKridikHomeEnvVar, "/home/mesh-kridik")
	homeFolder := GetEnv(common.MeshKridikHomeEnvVar, "/home/user")
	assert.Equal(t, homeFolder, "/home/mesh-kridik")
	os.Unsetenv(common.MeshKridikHomeEnvVar)
	homeFolder = GetEnv(common.MeshKridikHomeEnvVar, "/home/user")
	assert.Equal(t, homeFolder, "/home/user")
}

//Test_PluginsSourceFolder test
func Test_PluginsSourceFolder(t *testing.T) {
	fm := NewKFolder()
	err := CreatePluginsSourceFolderIfNotExist(fm)
	assert.NoError(t, err)
	a, err := GetPluginSourceSubFolder(fm)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(a, PluginSourceSubFolder))
}

//Test_PluginsCompiledFolder test
func Test_PluginsCompiledFolder(t *testing.T) {
	fm := NewKFolder()
	err := CreatePluginsCompiledFolderIfNotExist(fm)
	assert.NoError(t, err)
	a, err := GetCompilePluginSubFolder(fm)
	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(a, CompilePluginSubFolder))
}

func TestCreateBenchmarkFoldersErrorHomeFolder(t *testing.T) {
	ctl := gomock.NewController(t)
	fm := mocks.NewMockFolderMgr(ctl)
	fm.EXPECT().GetHomeFolder().Return("homePath", fmt.Errorf("error")).Times(1)
	err := CreateSecurityFolderIfNotExist("mesh", "v1.0.0", fm)
	assert.Error(t, err)
	fmr := NewKFolder()
	path, err := GetSecurityFolder("mesh", "istio", fmr)
	assert.NoError(t, err)
	rhfp := GetHomeFolder()
	fm2 := mocks.NewMockFolderMgr(ctl)
	fm2.EXPECT().GetHomeFolder().Return(rhfp, nil).Times(1)
	fm2.EXPECT().CreateFolder(path).Return(fmt.Errorf("error")).Times(1)
	err = CreateSecurityFolderIfNotExist("mesh", "istio", fm2)
	assert.Error(t, err)
}

func TestCreatePluginsCompiledFolderIfNotExist(t *testing.T) {
	ctl := gomock.NewController(t)
	fm := mocks.NewMockFolderMgr(ctl)
	fm.EXPECT().GetHomeFolder().Return("homePath", fmt.Errorf("error")).Times(1)
	err := CreatePluginsCompiledFolderIfNotExist(fm)
	assert.Error(t, err)
	fmr := NewKFolder()
	path, err := GetCompilePluginSubFolder(fmr)
	assert.NoError(t, err)
	rhfp := GetHomeFolder()
	fm2 := mocks.NewMockFolderMgr(ctl)
	fm2.EXPECT().GetHomeFolder().Return(rhfp, nil).Times(1)
	fm2.EXPECT().CreateFolder(path).Return(fmt.Errorf("error")).Times(1)
	err = CreatePluginsCompiledFolderIfNotExist(fm2)
	assert.Error(t, err)
}

func TestCreatePluginsSourcesFolderIfNotExist(t *testing.T) {
	ctl := gomock.NewController(t)
	fm := mocks.NewMockFolderMgr(ctl)
	fm.EXPECT().GetHomeFolder().Return("homePath", fmt.Errorf("error")).Times(1)
	err := CreatePluginsSourceFolderIfNotExist(fm)
	assert.Error(t, err)
	fmr := NewKFolder()
	path, err := GetPluginSourceSubFolder(fmr)
	assert.NoError(t, err)
	rhfp := GetHomeFolder()
	fm2 := mocks.NewMockFolderMgr(ctl)
	fm2.EXPECT().GetHomeFolder().Return(rhfp, nil).Times(1)
	fm2.EXPECT().CreateFolder(path).Return(fmt.Errorf("error")).Times(1)
	err = CreatePluginsSourceFolderIfNotExist(fm2)
	assert.Error(t, err)
}

func TestGetBenchmarkFoldersErrorHomeFolder(t *testing.T) {
	ctl := gomock.NewController(t)
	fm := mocks.NewMockFolderMgr(ctl)
	fm.EXPECT().GetHomeFolder().Return("homePath", fmt.Errorf("error")).Times(1)
	_, err := GetSecurityFolder("mesh", "istio", fm)
	assert.Error(t, err)
}
func TestGetSourcePluginFoldersErrorHomeFolder(t *testing.T) {
	ctl := gomock.NewController(t)
	fm := mocks.NewMockFolderMgr(ctl)
	fm.EXPECT().GetHomeFolder().Return("homePath", fmt.Errorf("error")).Times(1)
	_, err := GetPluginSourceSubFolder(fm)
	assert.Error(t, err)
}
func TestGetCompiledPluginFoldersErrorHomeFolder(t *testing.T) {
	ctl := gomock.NewController(t)
	fm := mocks.NewMockFolderMgr(ctl)
	fm.EXPECT().GetHomeFolder().Return("homePath", fmt.Errorf("error")).Times(1)
	_, err := GetCompilePluginSubFolder(fm)
	assert.Error(t, err)
}
