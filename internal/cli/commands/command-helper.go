package commands

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"github.com/chen-keinan/mesh-kridik/internal/logger"
	"github.com/chen-keinan/mesh-kridik/internal/models"
	"github.com/chen-keinan/mesh-kridik/pkg/filters"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/chen-keinan/mesh-kridik/ui"
	"github.com/mitchellh/colorstring"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
	"strings"
	"time"
)

func printTestResults(at []*models.SecurityCheck, table *tablewriter.Table, category string) models.CheckTotals {
	var (
		warnCounter int
		passCounter int
		failCounter int
	)
	for _, a := range at {
		var testType string
		if a.NonApplicable {
			testType = "Manual"
		} else {
			testType = "Automated"
		}
		if a.NonApplicable {
			warnTest := colorstring.Color("[yellow][Warn]")
			warnCounter++
			table.Append([]string{category, warnTest, testType, a.Name})
			continue
		}
		if a.TestSucceed {
			passTest := colorstring.Color("[green][Pass]")
			table.Append([]string{category, passTest, testType, a.Name})

			passCounter++
		} else {
			failTest := colorstring.Color("[red][Fail]")
			table.Append([]string{category, failTest, testType, a.Name})
			failCounter++
		}
	}
	return models.CheckTotals{Fail: failCounter, Pass: passCounter, Warn: warnCounter}
}

func printClassicTestResults(at []*models.SecurityCheck, log *logger.MeshKridikLogger) models.CheckTotals {
	var (
		warnCounter int
		passCounter int
		failCounter int
	)
	// just empty line
	fmt.Println()
	for _, a := range at {
		if a.NonApplicable {
			warnTest := colorstring.Color("[yellow][Warn]")
			log.Console(fmt.Sprintf("%s %s\n", warnTest, a.Name))
			warnCounter++
			continue
		}
		if a.TestSucceed {
			passTest := colorstring.Color("[green][Pass]")
			log.Console(fmt.Sprintf("%s %s\n", passTest, a.Name))
			passCounter++
		} else {
			failTest := colorstring.Color("[red][Fail]")
			log.Console(fmt.Sprintf("%s %s\n", failTest, a.Name))
			failCounter++
		}
		for index, pr := range a.PolicyResult {
			log.Console(fmt.Sprintf("       %d. %s\n", index+1, pr))
		}
	}
	return models.CheckTotals{Fail: failCounter, Pass: passCounter, Warn: warnCounter}
}

func calculateTotals(at []*models.SecurityCheck, log *logger.MeshKridikLogger) models.CheckTotals {
	var (
		warnCounter int
		passCounter int
		failCounter int
	)
	for _, a := range at {
		if a.NonApplicable {
			warnCounter++
			continue
		}
		if a.TestSucceed {
			passCounter++
		} else {
			failCounter++
		}
	}
	return models.CheckTotals{Fail: failCounter, Pass: passCounter, Warn: warnCounter}
}

//AddFailedMessages add failed audit test to report data
func AddFailedMessages(at *models.SecurityCheck, isSucceeded bool) []*models.SecurityCheck {
	av := make([]*models.SecurityCheck, 0)
	at.TestSucceed = isSucceeded
	if !isSucceeded || at.NonApplicable {
		av = append(av, at)
	}
	return av
}

//AddAllMessages add all audit test to report data
func AddAllMessages(at *models.SecurityCheck, isSucceeded bool) []*models.SecurityCheck {
	av := make([]*models.SecurityCheck, 0)
	at.TestSucceed = isSucceeded
	av = append(av, at)
	return av
}

//TestLoader load tests from filesystem
//command-helper.go
//go:generate mockgen -destination=../../mocks/mock_TestLoader.go -package=mocks . TestLoader
type TestLoader interface {
	LoadSecurityChecks(fi []utils.FilesInfo) []*models.SubCategory
}

//AuditTestLoader object
type AuditTestLoader struct {
}

//NewFileLoader create new file loader
func NewFileLoader() TestLoader {
	return &AuditTestLoader{}
}

//LoadSecurityChecks load audit test from benchmark folder
func (tl AuditTestLoader) LoadSecurityChecks(auditFiles []utils.FilesInfo) []*models.SubCategory {
	auditTests := make([]*models.SubCategory, 0)
	audit := models.Check{}
	for _, auditFile := range auditFiles {
		if !strings.HasSuffix(auditFile.Name, common.PolicySuffix) {
			err := yaml.Unmarshal([]byte(auditFile.Data), &audit)
			if err != nil {
				panic("Failed to unmarshal audit test yaml file")
			}
			auditTests = append(auditTests, audit.Categories[0].SubCategory)
		}
	}
	return auditTests
}

//FilterAuditTests filter audit tests by predicate chain
func FilterAuditTests(predicates []filters.Predicate, predicateParams []string, at *models.SubCategory) *models.SubCategory {
	return RunPredicateChain(predicates, predicateParams, len(predicates), at)
}

//RunPredicateChain call every predicate in chain and filter tests
func RunPredicateChain(predicates []filters.Predicate, predicateParams []string, size int, at *models.SubCategory) *models.SubCategory {
	if size == 0 {
		return at
	}
	return RunPredicateChain(predicates[1:size], predicateParams[1:size], size-1, predicates[size-1](at, predicateParams[size-1]))
}

// check weather are exist in array of specificTests
func isArgsExist(args []string, name string) bool {
	for _, n := range args {
		if n == name {
			return true
		}
	}
	return false
}

//GetResultProcessingFunction return processing function by specificTests
func GetResultProcessingFunction(args []string) ResultProcessor {
	if isArgsExist(args, common.Report) || isArgsExist(args, common.ReportFull) {
		return reportResultProcessor
	}
	return simpleResultProcessor
}

//getOutPutGeneratorFunction return output generator function
func getOutputGeneratorFunction(args []string) ui.OutputGenerator {
	switch {
	case isArgsExist(args, common.Report) || isArgsExist(args, common.ReportFull):
		return ReportOutputGenerator
	case isArgsExist(args, common.Classic) || isArgsExist(args, common.ClassicFull):
		return ConsoleOutputGenerator
	default:
		return ClassicOutputGenerator
	}
}

//buildPredicateChain build chain of filters based on command criteria
func buildPredicateChain(args []string) []filters.Predicate {
	pc := make([]filters.Predicate, 0)
	for _, n := range args {
		switch {
		case strings.HasPrefix(n, common.IncludeParam):
			pc = append(pc, filters.IncludeAudit)
		case strings.HasPrefix(n, common.ExcludeParam):
			pc = append(pc, filters.ExcludeAudit)
		case strings.HasPrefix(n, common.NodeParam):
			pc = append(pc, filters.NodeAudit)
		case n == "a":
			pc = append(pc, filters.Basic)
		}
	}
	return pc
}

//buildPredicateParams build chain of filters params based on command criteria
func buildPredicateChainParams(args []string) []string {
	pp := make([]string, 0)
	pp = append(pp, args...)
	return pp
}

func filteredAuditBenchTests(auditTests []*models.SubCategory, pc []filters.Predicate, pp []string) []*models.SubCategory {
	ft := make([]*models.SubCategory, 0)
	for _, adt := range auditTests {
		filteredAudit := FilterAuditTests(pc, pp, adt)
		if len(filteredAudit.Checks) == 0 {
			continue
		}
		ft = append(ft, filteredAudit)
	}
	return ft
}

func loadPolicies(fi []utils.FilesInfo) map[string]string {
	policyMap := make(map[string]string)
	for _, policy := range fi {
		if strings.HasSuffix(policy.Name, ".policy") {
			policyMap[policy.Name] = policy.Data
		}
	}
	return policyMap
}

func executeTests(ft []*models.SubCategory, mc *MeshCheck, policies map[string]string) []*models.SubCategory {
	completedTest := make([]*models.SubCategory, 0)
	mc.log.Console(ui.MeshCheck)
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                   // Start the spinner
	for _, f := range ft {
		s.Prefix = fmt.Sprintf("[Category] %s   ", f.Name)
		s.Start()
		tr := ui.ExecuteSpecs(f, mc.runAuditTest, policies)
		printClassicTestResults(f.Checks, mc.log)
		completedTest = append(completedTest, tr)
		s.Stop()
	}
	return completedTest
}
