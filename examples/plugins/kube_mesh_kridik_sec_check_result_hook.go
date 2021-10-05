package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chen-keinan/kube-mesh-kridik/pkg/models"
	"net/http"
	"strings"
)

//MeshKridikBenchAuditResultHook this plugin method accept mesh security check results
//event include test data , description , audit, remediation and result
func MeshKridikBenchAuditResultHook(lxdAuditResults models.MeshKridikSecurityResults) error {
	var sb = new(bytes.Buffer)
	err := json.NewEncoder(sb).Encode(lxdAuditResults)
	fmt.Print(lxdAuditResults)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "http://localhost:8090/audit-results", strings.NewReader(sb.String()))
	if err != nil {
		return err
	}
	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
