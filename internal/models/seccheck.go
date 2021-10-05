package models

import (
	"github.com/chen-keinan/kube-mesh-kridik/internal/common"
	"github.com/mitchellh/mapstructure"
)

//SecCheck data model
type SecCheck struct {
	BenchmarkType string     `yaml:"benchmark_type"`
	Categories    []Category `yaml:"categories"`
}

//SecCheckTotals model
type SecCheckTotals struct {
	Warn int
	Pass int
	Fail int
}

//Category data model
type Category struct {
	Name        string       `yaml:"name"`
	SubCategory *SubCategory `yaml:"sub_category"`
}

//SubCategory data model
type SubCategory struct {
	Name       string       `yaml:"name"`
	AuditTests []*CheckSpec `yaml:"audit_tests"`
}

//CheckSpec data model
type CheckSpec struct {
	Name                 string   `mapstructure:"name" yaml:"name"`
	ProfileApplicability string   `mapstructure:"profile_applicability" yaml:"profile_applicability"`
	Description          string   `mapstructure:"description" yaml:"description"`
	AuditCommand         []string `mapstructure:"audit" json:"audit"`
	CheckType            string   `mapstructure:"check_type" yaml:"check_type"`
	Remediation          string   `mapstructure:"remediation" yaml:"remediation"`
	Impact               string   `mapstructure:"impact" yaml:"impact"`
	AdditionalInfo       string   `mapstructure:"additional_info" yaml:"additional_info"`
	References           []string `mapstructure:"references" yaml:"references"`
	EvalExpr             string   `mapstructure:"eval_expr" yaml:"eval_expr"`
	TestSucceed          bool
	CommandParams        map[int][]string
	Category             string
	NonApplicable        bool
	TestType             string `mapstructure:"type" yaml:"type"`
}

//AuditResult data
type AuditResult struct {
	NumOfExec    int
	NumOfSuccess int
}

//UnmarshalYAML over unmarshall to add logic
func (at *CheckSpec) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var res map[string]interface{}
	if err := unmarshal(&res); err != nil {
		return err
	}
	err := mapstructure.Decode(res, &at)
	if err != nil {
		return err
	}
	if at.TestType == common.NonApplicableTest || at.TestType == common.ManualTest {
		at.NonApplicable = true
	}
	return nil
}
