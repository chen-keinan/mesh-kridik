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
	apm, err := box.FindString(common.SaferAuthorizationPolicyPatternsPolicy)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.SaferAuthorizationPolicyPatternsPolicy, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.SaferAuthorizationPolicyPatternsPolicy, Data: apm})
	//4
	toet, err := box.FindString(common.TLSOriginationForEgressTraffic)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.TLSOriginationForEgressTraffic, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.TLSOriginationForEgressTraffic, Data: toet})
	//5
	drto, err := box.FindString(common.DestinationRulePerformTLSOrigination)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.DestinationRulePerformTLSOrigination, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.DestinationRulePerformTLSOrigination, Data: drto})
	//6
	dp, err := box.FindString(common.ProtocolDetection)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.ProtocolDetection, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ProtocolDetection, Data: dp})
	//7
	dbp, err := box.FindString(common.DetectByProtocol)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.DetectByProtocol, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.DetectByProtocol, Data: dbp})
	//8
	cni, err := box.FindString(common.Cni)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.Cni, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.Cni, Data: cni})
	//9
	pce, err := box.FindString(common.PodCapabilitiesExist)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.PodCapabilitiesExist, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.PodCapabilitiesExist, Data: pce})
	//10
	gate, err := box.FindString(common.Gateway)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.Gateway, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.Gateway, Data: gate})
	//11
	aohc, err := box.FindString(common.AvoidOverlyBroadHostsConfigurations)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.AvoidOverlyBroadHostsConfigurations, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.AvoidOverlyBroadHostsConfigurations, Data: aohc})
	//12
	rgcp, err := box.FindString(common.RestrictGatewayCreationPrivileges)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.RestrictGatewayCreationPrivileges, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.RestrictGatewayCreationPrivileges, Data: rgcp})
	//13
	pnla, err := box.FindString(common.PathNormalizationInAuthorization)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.PathNormalizationInAuthorization, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.PathNormalizationInAuthorization, Data: pnla})
	//14
	cldc, err := box.FindString(common.ConfigureLimitDownstreamConnections)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.ConfigureLimitDownstreamConnections, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.ConfigureLimitDownstreamConnections, Data: cldc})
	//15
	dclc, err := box.FindString(common.DownstreamConnectionLimitConfigMap)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.DownstreamConnectionLimitConfigMap, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.DownstreamConnectionLimitConfigMap, Data: dclc})
	//16
	ipcl, err := box.FindString(common.IngressGatewayPatchedDownstreamConnectionLimit)
	if err != nil {
		return []utils.FilesInfo{}, fmt.Errorf("faild to load security checks %s  %s", common.IngressGatewayPatchedDownstreamConnectionLimit, err.Error())
	}
	fileInfo = append(fileInfo, utils.FilesInfo{Name: common.IngressGatewayPatchedDownstreamConnectionLimit, Data: ipcl})
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
