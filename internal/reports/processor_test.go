package reports

import (
	"github.com/chen-keinan/mesh-kridik/internal/models"
	"github.com/magiconair/properties/assert"
	"testing"
)

//Test_GenerateAuditReport test
func Test_GenerateAuditReport(t *testing.T) {
	ab := make([]*models.AuditBench, 0)
	ab = append(ab, &models.AuditBench{Name: "aaa", Description: "bbb", Impact: "ccc", Remediation: "ddd"})
	tb := GenerateAuditReport(ab)
	s := tb.String()
	assert.Equal(t, s, "--------------\t-------------------------------------------------------------------------------------------\nStatus:       \tFailed                                                                                     \nName:         \taaa                                                                                        \nDescription:  \tbbb                                                                                        \nAudit:        \t[]                                                                                         \nRemediation:  \tddd                                                                                        \nReferences:   \t[]                                                                                         \n              ")
}
