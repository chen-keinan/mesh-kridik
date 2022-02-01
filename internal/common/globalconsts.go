package common

const (
	//IstioMutualmTLS file name
	IstioMutualmTLS = "1_istio_mutual_mtls.yml"
	//AllowMtlsPermissiveMode policy name
	AllowMtlsPermissiveMode = "allow_mtls_permissive_mode.rego"
	//SaferAuthorizationPolicyPatterns file name
	SaferAuthorizationPolicyPatterns = "2_safer_authorization_policy_patterns.yml"
	//SaferAuthorizationPolicyPatternsPolicy policy name
	SaferAuthorizationPolicyPatternsPolicy = "safer_authorization_policy_pattern.rego"
	//TLSOriginationForEgressTraffic policy name
	TLSOriginationForEgressTraffic = "3_tls_origination_for_egress_traffic.yml"
	//DestinationRulePerformTLSOrigination policy name
	DestinationRulePerformTLSOrigination = "destination_rule_perform_tls_origination.rego"
	//ProtocolDetection spec
	ProtocolDetection = "4_protocol_detection.yml"
	//DetectByProtocol policy name
	DetectByProtocol = "detect_by_protocol.rego"
	//Cni spec
	Cni = "5_cni.yml"
	//PodCapabilitiesExist policy name
	PodCapabilitiesExist = "pod_capabilities_exist.rego"
	//Gateway spec
	Gateway = "6_gateway.yml"
	//ConfigureLimitDownstreamConnections spec
	ConfigureLimitDownstreamConnections = "7_configure_limit_downstream_connections .yml"
	//ConfigureThirdPartyServiceAccountTokens spec
	//nolint:gosec
	ConfigureThirdPartyServiceAccountTokens = "8_configure_third_party_service_account_tokens.yml"
	//UnderstandTrafficCaptureLimitations spec
	UnderstandTrafficCaptureLimitations = "10_understand_traffic_capture_limitations.yml"
	//ControlPlane spec
	ControlPlane = "9_control_plane.yml"
	//AvoidOverlyBroadHostsConfigurations policy name
	AvoidOverlyBroadHostsConfigurations = "avoid_overly_broad_hosts_configurations.rego"
	//RestrictGatewayCreationPrivileges policy name
	RestrictGatewayCreationPrivileges = "restrict_gateway_creation_privileges.rego"
	//PathNormalizationInAuthorization policy name
	PathNormalizationInAuthorization = "path_normalization_in_authorization.rego"
	//DownstreamConnectionLimitConfigMap policy name
	DownstreamConnectionLimitConfigMap = "downstream_connection_limit_config_map.rego"
	//IngressGatewayPatchedDownstreamConnectionLimit policy name
	IngressGatewayPatchedDownstreamConnectionLimit = "ingress_gateway_patched_downstream_connection_limit.rego"
	//ClosePort8008UnauthenticatePlaintext policy
	ClosePort8008UnauthenticatePlaintext = "close_port_8008_as_unauthenticate_plaintext.rego"
	//ClosePort15010UnauthenticatePlaintext policy
	ClosePort15010UnauthenticatePlaintext = "close_port_15010_as_unauthenticate_plaintext.rego"
	//IstioUsing3rdPartyTokens policy
	//nolint:gosec
	IstioUsing3rdPartyTokens = "istio_using_3rd_party_tokens.rego"
	//ProxyBlocksExternalHostWithinMesh policy
	ProxyBlocksExternalHostWithinMesh = "proxy_blocks_external_host_within_mesh.rego"
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
	PolicySuffix = ".rego"
)
