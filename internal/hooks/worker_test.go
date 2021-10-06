package hooks

import (
	m2 "github.com/chen-keinan/mesh-kridik/pkg/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func Test_NewPluginWorker(t *testing.T) {
	production, err := zap.NewProduction()
	assert.NoError(t, err)
	completedChan := make(chan bool)
	plChan := make(chan m2.LxdAuditResults)
	pw := NewPluginWorker(NewPluginWorkerData(plChan, LxdBenchAuditResultHook{}, completedChan), production)
	assert.True(t, len(pw.cmd.plugins.Plugins) == 0)

}
