package filters

import (
	"github.com/chen-keinan/mesh-kridik/internal/models"
	"github.com/chen-keinan/mesh-kridik/pkg/utils"
	"strings"
)

// Predicate filter audit tests cmd criteria
type Predicate func(tests *models.SubCategory, params string) *models.SubCategory

// IncludeCheck include audit tests , only included tests will be executed
var IncludeCheck Predicate = func(tests *models.SubCategory, params string) *models.SubCategory {
	sat := make([]*models.SecurityCheck, 0)
	spt := utils.GetAuditTestsList("i", params)
	// check if param include category
	for _, sp := range spt {
		if strings.HasPrefix(tests.Name, sp) {
			return tests
		}
	}
	// check tests
	for _, at := range tests.Checks {
		for _, sp := range spt {
			if strings.HasPrefix(at.Name, sp) {
				sat = append(sat, at)
			}
		}
	}
	if len(sat) == 0 {
		return &models.SubCategory{Name: tests.Name, Checks: make([]*models.SecurityCheck, 0)}
	}
	return &models.SubCategory{Name: tests.Name, Checks: sat}
}

// ExcludeCheck audit test from been executed
var ExcludeCheck Predicate = func(tests *models.SubCategory, params string) *models.SubCategory {
	sat := make([]*models.SecurityCheck, 0)
	spt := utils.GetAuditTestsList("e", params)
	// if exclude category
	for _, sp := range spt {
		if strings.HasPrefix(tests.Name, sp) {
			return &models.SubCategory{Name: tests.Name, Checks: []*models.SecurityCheck{}}
		}
	}
	for _, at := range tests.Checks {
		var skipTest bool
		for _, sp := range spt {
			if strings.HasPrefix(at.Name, sp) {
				skipTest = true
			}
		}
		if skipTest {
			continue
		}
		sat = append(sat, at)
	}
	return &models.SubCategory{Name: tests.Name, Checks: sat}
}

// Basic filter by specific audit tests as set in command
var Basic Predicate = func(tests *models.SubCategory, params string) *models.SubCategory {
	return tests
}
