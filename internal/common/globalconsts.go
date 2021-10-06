package common

const (
	//FilesystemConfiguration file name
	FilesystemConfiguration = "1.1_filesystem_configuration.yml"
	//ConfigureSoftwareUpdates file name
	ConfigureSoftwareUpdates = "1.2_configure_software_updates.yml"
	//ConfigureSudo file name
	ConfigureSudo = "1.3_configure_sudo.yml"
	//FilesystemIntegrityChecking file name
	FilesystemIntegrityChecking = "1.4_filesystem_integrity_checking.yml"
	//AdditionalProcessHardening file name
	AdditionalProcessHardening = "1.5_additional_process_hardening.yml"
	//MandatoryAccessControl file name
	MandatoryAccessControl = "1.6_mandatory_access_control.yml"
	//WarningBanners file name
	WarningBanners = "1.7_warning_banners.yml"
	//EnsureUpdates file name
	EnsureUpdates = "1.8_ensure_updates.yml"
	//InetdServices file name
	InetdServices = "2.1_inetd_services.yml"
	//SpecialPurposeServices file name
	SpecialPurposeServices = "2.2_special_purpose_services.yml"
	//ServiceClients file name
	ServiceClients = "2.3_service_clients.yml"
	//NonessentialServices file name
	NonessentialServices = "2.4_nonessential_services.yml"
	//NetworkParameters file name
	NetworkParameters = "3.1_network_parameters.yml"
	//NetworkParametersHost file name
	NetworkParametersHost = "3.2_network_parameters_host.yml"
	//TCPWrappers file name
	TCPWrappers = "3.3_tcp_wrappers.yml"
	//FirewallConfiguration file name
	FirewallConfiguration = "3.4_firewall_configuration.yml"
	//ConfigureLogging file name
	ConfigureLogging = "4.1_configure_logging.yml"
	//EnsureLogrotateConfigured file name
	EnsureLogrotateConfigured = "4.2_ensure_logrotate_configured.yml"
	//EnsureLogrotateAssignsAppropriatePermissions file name
	EnsureLogrotateAssignsAppropriatePermissions = "4.3_ensure_logrotate_assigns_appropriate_permissions.yml"
	//ConfigureCron file name
	ConfigureCron = "5.1_configure_cron.yml"
	//SSHServerConfiguration file name
	SSHServerConfiguration = "5.2_ssh_server_configuration.yml"
	//ConfigurePam file name
	ConfigurePam = "5.3_configure_pam.yml"
	//UserAccountsAndEnvironment file name
	UserAccountsAndEnvironment = "5.4_user_accounts_and_environment.yml"
	//RootLoginRestrictedSystemConsole file name
	RootLoginRestrictedSystemConsole = "5.5_root_login_restricted_system_console.yml"
	//EnsureAccessSuCommandRestricted file name
	EnsureAccessSuCommandRestricted = "5.6_ensure_access_su_command_restricted.yml"
	//SystemFilePermissions file name
	SystemFilePermissions = "6.1_system_file_permissions.yml"
	//UserAndGroupSettings file name
	UserAndGroupSettings = "6.2_user_and_group_settings.yml"
	//GrepRegex for tests
	GrepRegex = "[^\"]\\S*'"
	//MultiValue for tests
	MultiValue = "MultiValue"
	//SingleValue for tests
	SingleValue = "SingleValue"
	//EmptyValue for test
	EmptyValue = "EmptyValue"
	//NotValidNumber value
	NotValidNumber = "10000"
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
	//LdxProbeCli Name
	LdxProbeCli = "ldx-probe"
	//LdxProbeVersion version
	LdxProbeVersion = "0.1"
	//IncludeParam param
	IncludeParam = "i="
	//ExcludeParam param
	ExcludeParam = "e="
	//NodeParam param
	NodeParam = "n="
	//LxdProbeHomeEnvVar ldx probe Home env var
	LxdProbeHomeEnvVar = "LXD_PROBE_HOME"
	//LxdProbe binary name
	LxdProbe = "mesh-kridik"
	//RootUser process user owner
	RootUser = "root"
	//NonApplicableTest test is not applicable
	NonApplicableTest = "non_applicable"
	//ManualTest test can only be manual executed
	ManualTest = "manual"
	//LxdBenchAuditResultHook hook name
	LxdBenchAuditResultHook = "LxdBenchAuditResultHook"
)
