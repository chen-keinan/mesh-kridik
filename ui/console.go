package ui

import (
	"github.com/chen-keinan/mesh-kridik/internal/logger"
	"github.com/chen-keinan/mesh-kridik/internal/models"
)

// OutputGenerator for  audit results
type OutputGenerator func(at []*models.SubCategory, log *logger.MeshKridikLogger)

//PrintOutput print audit test result to console
func PrintOutput(auditTests []*models.SubCategory, outputGenerator OutputGenerator, log *logger.MeshKridikLogger) {
	log.Console(auditResult)
	outputGenerator(auditTests, log)
}

//ExecuteSpecs execute audit test and show progress bar
func ExecuteSpecs(a *models.SubCategory, execTestFunc func(ad *models.SecurityCheck, policies map[string]string) []*models.SecurityCheck, policies map[string]string) *models.SubCategory {
	if len(a.Checks) == 0 {
		return a
	}
	completedTest := make([]*models.SecurityCheck, 0)
	for _, test := range a.Checks {
		ar := execTestFunc(test, policies)
		completedTest = append(completedTest, ar...)
	}
	return &models.SubCategory{Name: a.Name, Checks: completedTest}
}
