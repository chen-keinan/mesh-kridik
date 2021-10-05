package ui

import (
	"github.com/chen-keinan/kube-mesh-kridik/internal/logger"
	"github.com/chen-keinan/kube-mesh-kridik/internal/models"
)

// OutputGenerator for  audit results
type OutputGenerator func(at []*models.SubCategory, log *logger.MeshKridikLogger)

//PrintOutput print audit test result to console
func PrintOutput(auditTests []*models.SubCategory, outputGenerator OutputGenerator, log *logger.MeshKridikLogger) {
	log.Console(auditResult)
	outputGenerator(auditTests, log)
}

//ExecuteSpecs execute audit test and show progress bar
func ExecuteSpecs(a *models.SubCategory, execTestFunc func(ad *models.CheckSpec) []*models.CheckSpec) *models.SubCategory {
	if len(a.AuditTests) == 0 {
		return a
	}
	completedTest := make([]*models.CheckSpec, 0)
	for _, test := range a.AuditTests {
		ar := execTestFunc(test)
		completedTest = append(completedTest, ar...)
	}
	return &models.SubCategory{Name: a.Name, AuditTests: completedTest}
}
