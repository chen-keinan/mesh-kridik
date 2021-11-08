package commands

import (
	"fmt"
	"github.com/chen-keinan/go-command-eval/eval"
	"github.com/chen-keinan/mesh-kridik/internal/cli/mocks"
	"github.com/chen-keinan/mesh-kridik/internal/models"
	m2 "github.com/chen-keinan/mesh-kridik/pkg/models"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestRunAuditTests(t *testing.T) {
	const policy = `package example
	default deny = false
	deny {
		some i
		input.kind == "Pod"
		image := input.spec.containers[i].image
		not startswith(image, "kalpine")
		}`
	tests := []struct {
		name              string
		testFile          string
		completedChan     chan bool
		plChan            chan m2.MeshCheckResults
		wantTestSucceeded bool
	}{

		{name: "Test_MultiCommandParams_OK", testFile: "CheckMultiParamOK.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: true},
		{name: "Test_MultiCommandParams_OK_With_IN", testFile: "CheckMultiParamOKWithIN.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: true},
		{name: "Test_MultiCommandParams_NOKWith_IN", testFile: "CheckMultiParamNOKWithIN.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: false},
		{name: "Test_MultiCommandParamsPass1stResultToNext", testFile: "CheckMultiParamPass1stResultToNext.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: false},
		{name: "Test_MultiCommandParamsComplex", testFile: "CheckMultiParamComplex.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: true},
		{name: "Test_MultiCommandParamsComplexOppositeEmptyReturn", testFile: "CheckInClauseOppositeEmptyReturn.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: false},
		{name: "Test_MultiCommandParamsComplexOppositeWithNumber", testFile: "CheckInClauseOppositeWithNum.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: false},
		{name: "Test_MultiCommand4_2_13", testFile: "CheckInClause4.2.13.yml", completedChan: make(chan bool), plChan: make(chan m2.MeshCheckResults), wantTestSucceeded: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ab := models.Check{}
			err := yaml.Unmarshal(readTestData(tt.testFile, t), &ab)
			if err != nil {
				t.Errorf("failed to Unmarshal test file %s error : %s", tt.testFile, err.Error())
			}
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			evalCmd := mocks.NewMockCmdEvaluator(ctrl)
			testBench := ab.Categories[0].SubCategory.Checks[0]
			testBench.EvalExpr = "'${0}' != '';&& [${1} MATCH no_permission.policy QUERY example.allow RETURN allow]"
			evalCmd.EXPECT().EvalCommandPolicy(testBench.CheckCommand, testBench.EvalExpr, policy).Return(eval.CmdEvalResult{Match: tt.wantTestSucceeded, Error: nil}).Times(1)
			kb := MeshCheck{Evaluator: evalCmd, ResultProcessor: GetResultProcessingFunction([]string{}), PlChan: tt.plChan, CompletedChan: tt.completedChan}
			policyMap := make(map[string]string)
			policyMap["no_permission.policy"] = policy
			kb.runAuditTest(ab.Categories[0].SubCategory.Checks[0], policyMap)
			assert.Equal(t, ab.Categories[0].SubCategory.Checks[0].TestSucceed, tt.wantTestSucceeded)
			go func() {
				<-tt.plChan
				tt.completedChan <- true
			}()
		})
	}
}

func readTestData(fileName string, t *testing.T) []byte {
	f, err := os.Open(fmt.Sprintf("./fixtures/%s", fileName))
	if err != nil {
		t.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	return b
}

//Test_NewMeshCheck test
func Test_NewMeshCheck(t *testing.T) {
	args := []string{"a", "i=1.2.3"}
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	evaluator := eval.NewEvalCmd()
	ka := NewMeshCheck(args, plChan, completedChan, []utils.FilesInfo{}, evaluator)
	assert.True(t, len(ka.PredicateParams) == 2)
	assert.True(t, len(ka.PredicateChain) == 2)
	assert.True(t, ka.ResultProcessor != nil)
	go func() {
		<-plChan
		completedChan <- true
	}()
}

//Test_Help testR
func Test_Help(t *testing.T) {
	args := []string{"a", "i=1.2.3"}
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	evaluator := eval.NewEvalCmd()
	ka := NewMeshCheck(args, plChan, completedChan, []utils.FilesInfo{}, evaluator)
	help := ka.Help()
	assert.True(t, len(help) > 0)
	go func() {
		<-plChan
		completedChan <- true
	}()
}

//Test_reportResultProcessor test
func Test_reportResultProcessor(t *testing.T) {
	ad := &models.SecurityCheck{Name: "1.2.1 aaa"}
	fm := reportResultProcessor(ad, true)
	assert.True(t, len(fm) == 0)
	fm = reportResultProcessor(ad, false)
	assert.True(t, len(fm) == 1)
	assert.Equal(t, fm[0].Name, "1.2.1 aaa")
}

//Test_MeshSynopsis test
func Test_MeshSynopsis(t *testing.T) {
	args := []string{"a", "i=1.2.3"}
	completedChan := make(chan bool)
	plChan := make(chan m2.MeshCheckResults)
	evaluator := eval.NewEvalCmd()
	ka := NewMeshCheck(args, plChan, completedChan, []utils.FilesInfo{}, evaluator)
	s := ka.Synopsis()
	assert.True(t, len(s) > 0)
	go func() {
		<-plChan
		completedChan <- true
	}()
}

func Test_sendResultToPlugin(t *testing.T) {
	pChan := make(chan m2.MeshCheckResults)
	cChan := make(chan bool)
	auditTests := make([]*models.SubCategory, 0)
	ab := make([]*models.SecurityCheck, 0)
	ats := &models.SecurityCheck{Name: "bbb", TestSucceed: true}
	atf := &models.SecurityCheck{Name: "ccc", TestSucceed: false}
	ab = append(ab, ats)
	ab = append(ab, atf)
	mst := &models.SubCategory{Name: "aaa", Checks: ab}
	auditTests = append(auditTests, mst)
	go func() {
		<-pChan
		cChan <- true
	}()
	sendResultToPlugin(pChan, cChan, auditTests)

}
func Test_calculateFinalTotal(t *testing.T) {
	att := make([]models.CheckTotals, 0)
	atOne := models.CheckTotals{Fail: 2, Pass: 3, Warn: 1}
	atTwo := models.CheckTotals{Fail: 1, Pass: 5, Warn: 7}
	att = append(att, atOne)
	att = append(att, atTwo)
	res := calculateFinalTotal(att)
	assert.Equal(t, res.Warn, 8)
	assert.Equal(t, res.Pass, 8)
	assert.Equal(t, res.Fail, 3)
	str := printFinalResults([]models.CheckTotals{res})
	assert.Equal(t, str, "Test Result Total:    \x1b[32mPass:\x1b[0m 8 , \x1b[33mWarn:\x1b[0m 8 , \x1b[31mFail:\x1b[0m 3 ")
}
