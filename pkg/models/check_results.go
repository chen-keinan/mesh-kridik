package models

//MeshCheckResults encapsulate audit test results to be consumed by user plugin
type MeshCheckResults struct {
	BenchmarkType string             `yaml:"benchmark_type"`
	Categories    []AuditBenchResult `yaml:"audit_bench_result"`
}

//AuditBenchResult data model
type AuditBenchResult struct {
	Name                 string   `yaml:"name"`
	ProfileApplicability string   `yaml:"profile_applicability"`
	Category             string   `yaml:"category"`
	Description          string   `yaml:"description"`
	AuditCommand         []string `json:"audit_command"`
	Remediation          string   `yaml:"remediation"`
	Impact               string   `yaml:"impact"`
	AdditionalInfo       string   `yaml:"additional_info"`
	References           []string `yaml:"references"`
	TestResult           string   `yaml:"test_result"`
}
