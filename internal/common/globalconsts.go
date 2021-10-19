package common

const (
	//IstioMutualmTLS file name
	IstioMutualmTLS = "1_istio_mutual_mtls.yml"
	//AllowMtlsPermissiveMode policy name
	AllowMtlsPermissiveMode = "allow_mtls_permissive_mode.policy"
	//SaferAuthorizationPolicyPatterns file name
	SaferAuthorizationPolicyPatterns = "2_safer_authorization_policy_patterns.yml"
	//SaferAuthorizationPolicyPatternsPolicy policy name
	SaferAuthorizationPolicyPatternsPolicy = "safer_authorization_policy_pattern.policy"
	//TLSOriginationForEgressTraffic policy name
	TLSOriginationForEgressTraffic = "3_tls_origination_for_egress_traffic.yml"
	//DestinationRulePerformTLSOrigination policy name
	DestinationRulePerformTLSOrigination = "destination_rule_perform_tls_origination.policy"
	//ProtocolDetection spec
	ProtocolDetection = "4_protocol_detection.yml"
	//DetectByProtocol policy name
	DetectByProtocol = "detect_by_protocol.policy"
	//Cni spec
	Cni = "5_cni.yml"
	//PodCapabilitiesExist policy name
	PodCapabilitiesExist = "pod_capabilities_exist.policy"
	//Report arg
	Report = "r"
	//ReportFull arg
	ReportFull = "report"
	//Classic arg
	Classic = "c"
	//ClassicFull arg
	ClassicFull = "classic"
	//Synopsis help
	Synopsis = "synopsis"
	//MeshKridikCli Name
	MeshKridikCli = "mesh-kridik"
	//MeshKridikVersion version
	MeshKridikVersion = "0.1"
	//IncludeParam param
	IncludeParam = "i="
	//ExcludeParam param
	ExcludeParam = "e="
	//NodeParam param
	NodeParam = "n="
	//MeshKridikHomeEnvVar ldx probe Home env var
	MeshKridikHomeEnvVar = "MESH_KRIDIK_HOME"
	//MeshKridik binary name
	MeshKridik = "mesh-kridik"
	//NonApplicableTest test is not applicable
	NonApplicableTest = "non_applicable"
	//ManualTest test can only be manual executed
	ManualTest = "manual"
	//MeshSecurityCheckResultHook hook name
	MeshSecurityCheckResultHook = "MeshSecurityCheckResultHook"
	//PolicySuffix name
	PolicySuffix = ".policy"
)
