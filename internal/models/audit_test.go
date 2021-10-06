package models

import (
	"fmt"
	"github.com/chen-keinan/mesh-kridik/internal/common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func TestAuditBench_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     string
	}{
		{name: "non applicable test", fileName: "no_applicable.yml", want: common.NonApplicableTest},
		{name: "manual test", fileName: "manual.yml", want: common.ManualTest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ab := Audit{}
			err := yaml.Unmarshal(readTestData(tt.fileName, t), &ab)
			if err != nil {
				t.Errorf("TestAuditBench_UnmarshalYAML failed to unmarshal json %v", err)
			}
			got := ab.Categories[0].SubCategory.AuditTests[0].TestType
			if tt.want != got {
				t.Errorf("TestAuditBench_UnmarshalYAML want %v got %v", tt.want, got)
			}
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
