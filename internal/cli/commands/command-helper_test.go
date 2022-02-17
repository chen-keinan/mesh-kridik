package commands

import (
	"fmt"
	"github.com/chen-keinan/go-command-eval/eval"
	"github.com/chen-keinan/mesh-kridik/internal/cli/mocks"
	"github.com/chen-keinan/mesh-kridik/internal/logger"
	"github.com/chen-keinan/mesh-kridik/internal/models"
	"github.com/chen-keinan/mesh-kridik/internal/startup"
	"github.com/chen-keinan/mesh-kridik/pkg/filters"
	m2 "github.com/chen-keinan/mesh-kridik/pkg/models"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/olekukonko/tablewriter"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

//Test_AddFailedMessages text
func Test_AddFailedMessages(t *testing.T) {
	atb1 := &models.SecurityCheck{TestSucceed: false}
	afm := AddFailedMessages(atb1, false)
	assert.True(t, len(afm) == 1)
	atb2 := &models.SecurityCheck{TestSucceed: true}
	afm = AddFailedMessages(atb2, true)
	assert.True(t, len(afm) == 0)
}

//Test_isArgsExist
func Test_isArgsExist(t *testing.T) {
	args := []string{"aaa", "bbb"}
	exist := isArgsExist(args, "aaa")
	assert.True(t, exist)
	exist = isArgsExist(args, "ccc")
	assert.False(t, exist)
}

//Test_isArgsExist
func Test_GetProcessingFunction(t *testing.T) {
	args := []string{"r"}
	a := GetResultProcessingFunction(args)
	name := GetFunctionName(a)
	assert.True(t, strings.Contains(name, "commands.glob..func5"))
	args = []string{"report"}
	a = GetResultProcessingFunction(args)
	name = GetFunctionName(a)
	assert.True(t, strings.Contains(name, "commands.glob..func5"))
	args = []string{"c"}
	a = GetResultProcessingFunction(args)
	name = GetFunctionName(a)
	assert.True(t, strings.Contains(name, "commands.glob..func4"))
	args = []string{"classic"}
	a = GetResultProcessingFunction(args)
	name = GetFunctionName(a)
	assert.True(t, strings.Contains(name, "commands.glob..func4"))
	args = []string{}
	a = GetResultProcessingFunction(args)
	name = GetFunctionName(a)
	assert.True(t, strings.Contains(name, "commands.glob..func4"))
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

//Test_getSpecificTestsToExecute test
func Test_getSpecificTestsToExecute(t *testing.T) {
	test := utils.GetAuditTestsList("i", "i=1.2.4,1.2.5")
	assert.Equal(t, test[0], "1.2.4")
	assert.Equal(t, test[1], "1.2.5")
}

//Test_LoadSecurityCheck test
func Test_LoadSecurityCheck(t *testing.T) {
	fm := utils.NewKFolder()
	folder, err2 := utils.GetSecurityFolder("mesh", "istio", fm)
	assert.NoError(t, err2)
	err := os.RemoveAll(folder)
	if err != nil {
		t.Fatal(err)
	}
	err = utils.CreateHomeFolderIfNotExist(fm)
	if err != nil {
		t.Fatal(err)
	}
	err = utils.CreateSecurityFolderIfNotExist("mesh", "istio", fm)
	if err != nil {
		t.Fatal(err)
	}
	bFiles, err := startup.GenerateMeshSecurityFiles()
	if err != nil {
		t.Fatal(err)
	}
	err = startup.SaveSecurityFilesIfNotExist("mesh", "istio", bFiles)
	if err != nil {
		t.Fatal(err)
	}
	at, err := NewFileLoader().LoadSecurityChecks(bFiles)
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, len(at) != 0)
	assert.True(t, strings.Contains(at[0].Checks[0].Name, "10.0 Securing egress traffic"))
}

func Test_LoadSecurityWithFailure(t *testing.T) {
	filesd, err := NewFileLoader().LoadSecurityChecks([]utils.FilesInfo{{Name: "aa.yaml", Data: ",ers"}})
	assert.Nilf(t, filesd, "file expexted to be nil")
	if err == nil {
		t.Errorf("Test_LoadSecurityWithFailure expect err: %s", "yaml: did not find expected node content")
	}
}

//Test_FilterAuditTests test
func Test_FilterAuditTests(t *testing.T) {
	at := &models.SubCategory{Checks: []*models.SecurityCheck{{Name: "1.2.1 aaa"}, {Name: "2.2.2"}}}
	fab := FilterAuditTests([]filters.Predicate{filters.IncludeCheck}, []string{"1.2.1"}, at)
	assert.Equal(t, fab.Checks[0].Name, "1.2.1 aaa")
	assert.True(t, len(fab.Checks) == 1)
}

//Test_buildPredicateChain test
func Test_buildPredicateChain(t *testing.T) {
	fab := buildPredicateChain([]string{"a", "i=1.2.1"})
	assert.True(t, len(fab) == 2)
	fab = buildPredicateChain([]string{"a"})
	assert.True(t, len(fab) == 1)
	fab = buildPredicateChain([]string{"i=1.2.1"})
	assert.True(t, len(fab) == 1)
}

//Test_buildPredicateChainParams test
func Test_buildPredicateChainParams(t *testing.T) {
	p := buildPredicateChainParams([]string{"a", "i=1.2.1"})
	assert.True(t, len(p) == 2)
	assert.Equal(t, p[0], "a")
	assert.Equal(t, p[1], "i=1.2.1")
}

func Test_filteredAuditBenchTests(t *testing.T) {
	asc := []*models.SubCategory{{Checks: []*models.SecurityCheck{{Name: "1.1.0 bbb"}}}}
	fp := []filters.Predicate{filters.IncludeCheck, filters.ExcludeCheck}
	st := []string{"i=1.1.0", "e=1.1.0"}
	fr := filteredAuditBenchTests(asc, fp, st)
	assert.True(t, len(fr) == 0)
}

//Test_executeTests test
func Test_executeTests(t *testing.T) {
	const policy = `package example
	default deny = false
	allow {
		some i
		input.kind == "Pod"
		image := input.spec.containers[i].image
		not startswith(image, "kalpine")
		}`
	ab := &models.SecurityCheck{}
	ab.CheckCommand = []string{"aaa", "bbb"}
	ab.EvalExpr = "'${0}' != ''; && [${1} MATCH no_permission.rego QUERY example.allow RETURN allow]"
	ab.CommandParams = map[int][]string{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	evalcmd := mocks.NewMockCmdEvaluator(ctrl)
	evalcmd.EXPECT().EvalCommandPolicy([]string{"aaa", "bbb"}, ab.EvalExpr, policy).Return(eval.CmdEvalResult{Match: true, CmdEvalExpr: ab.EvalExpr, Error: nil})
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	infos := []utils.FilesInfo{{Name: "no_permission.rego", Data: policy}}
	kb := &MeshCheck{FilesInfo: infos, ResultProcessor: GetResultProcessingFunction([]string{}), PlChan: plChan, CompletedChan: completedChan, Evaluator: evalcmd, log: logger.GetLog()}
	sc := []*models.SubCategory{{Checks: []*models.SecurityCheck{ab}}}
	policyMap := make(map[string]string)
	policyMap["no_permission.rego"] = policy
	executeTests(sc, kb, policyMap)
	assert.True(t, ab.TestSucceed)
	go func() {
		<-plChan
		completedChan <- true
	}()
}

func TestPrintTestResults(t *testing.T) {
	tests := []struct {
		name         string
		tests        []*models.SecurityCheck
		testCategory string
		testType     string
		warn         int
		pass         int
		fail         int
	}{
		{name: "regular result", testCategory: "aaa", tests: []*models.SecurityCheck{{Name: "bbb", TestSucceed: true}, {Name: "ccc", TestSucceed: false}, {Name: "ddd", NonApplicable: true}}, warn: 1, pass: 1, fail: 1, testType: "regular"},
		{name: "classic result", testCategory: "aaa", tests: []*models.SecurityCheck{{Name: "bbb", TestSucceed: true}, {Name: "ccc", TestSucceed: false}, {Name: "ddd", NonApplicable: true}}, warn: 1, pass: 1, fail: 1, testType: "classic"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr models.CheckTotals
			if tt.testType == "regular" {
				tr = printTestResults(tt.tests, tablewriter.NewWriter(os.Stdout), tt.testCategory)
			} else {
				tr = printClassicTestResults(tt.tests, logger.GetLog())
			}
			if tr.Pass != tt.pass {
				t.Errorf("TestPrintTestResults() = %v, want %v", tr.Pass, tt.pass)
			}
			if tr.Fail != tt.fail {
				t.Errorf("TestPrintTestResults() = %v, want %v", tr.Fail, tt.fail)
			}
			if tr.Warn != tt.warn {
				t.Errorf("TestPrintTestResults() = %v, want %v", tr.Warn, tt.warn)
			}
		})
	}
}

func TestLoaPolicies(t *testing.T) {
	tests := []struct {
		name  string
		files []utils.FilesInfo
		got   map[string]string
	}{
		{name: "load policies with policies files", files: []utils.FilesInfo{{Name: "a.rego", Data: "policy data"}}, got: map[string]string{"a.rego": "policy data"}},
		{name: "load policies without policies files", files: []utils.FilesInfo{{Name: "a.yml", Data: "spec data"}}, got: map[string]string{}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := loadPolicies(tt.files)
			if len(want) != len(tt.got) {
				t.Fatal(fmt.Sprintf("TestLoaPolicies: want:%d got:%d", len(want), len(tt.got)))
			}
		})
	}
}

func TestLoaTotals(t *testing.T) {
	tests := []struct {
		name   string
		checks []*models.SecurityCheck
		want   models.CheckTotals
	}{
		{name: "check test total with succeed and failure", checks: []*models.SecurityCheck{{TestSucceed: true}, {TestSucceed: false}}, want: models.CheckTotals{Pass: 1, Fail: 1, Warn: 0}},
		{name: "check fail and warn", checks: []*models.SecurityCheck{{TestSucceed: false}, {NonApplicable: true}}, want: models.CheckTotals{Pass: 0, Fail: 1, Warn: 1}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateTotals(tt.checks)
			if got.Pass != tt.want.Pass {
				t.Fatal(fmt.Sprintf("TestLoaTotals: want:%d got:%d", tt.want.Pass, got.Pass))
			}
			if got.Fail != tt.want.Fail {
				t.Fatal(fmt.Sprintf("TestLoaTotals: want:%d got:%d", tt.want.Fail, got.Fail))
			}
			if got.Warn != tt.want.Warn {
				t.Fatal(fmt.Sprintf("TestLoaTotals: want:%d got:%d", tt.want.Warn, got.Warn))
			}
		})
	}
}
