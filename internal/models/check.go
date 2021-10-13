package models

import (
	"github.com/chen-keinan/go-command-eval/utils"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/mitchellh/mapstructure"
)

//Check data model
type Check struct {
	BenchmarkType string     `yaml:"benchmark_type"`
	Categories    []Category `yaml:"categories"`
}

//CheckTotals model
type CheckTotals struct {
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
	Name   string           `yaml:"name"`
	Checks []*SecurityCheck `yaml:"security_checks"`
}

//SecurityCheck data model
type SecurityCheck struct {
	Name                 string   `mapstructure:"name" yaml:"name"`
	ProfileApplicability string   `mapstructure:"profile_applicability" yaml:"profile_applicability"`
	Description          string   `mapstructure:"description" yaml:"description"`
	CheckCommand         []string `mapstructure:"check_command" json:"check_command"`
	CheckType            string   `mapstructure:"check_type" yaml:"check_type"`
	Remediation          string   `mapstructure:"remediation" yaml:"remediation"`
	Impact               string   `mapstructure:"impact" yaml:"impact"`
	AdditionalInfo       string   `mapstructure:"additional_info" yaml:"additional_info"`
	References           []string `mapstructure:"references" yaml:"references"`
	DefaultValue         string   `mapstructure:"default_value" yaml:"default_value"`
	EvalExpr             string   `mapstructure:"eval_expr" yaml:"eval_expr"`
	PolicyName           string   `mapstructure:"policy_name" yaml:"policy_name"`
	EvalMessage          string   `mapstructure:"eval_message" yaml:"eval_message"`
	PolicyResult         []utils.PolicyResult
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
func (at *SecurityCheck) UnmarshalYAML(unmarshal func(interface{}) error) error {
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
