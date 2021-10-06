package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chen-keinan/mesh-kridik/pkg/models"
	"net/http"
	"strings"
)

//MeshSecurityCheckResultHook this plugin method accept mesh security check results
//event include test data , description , audit, remediation and result
func MeshSecurityCheckResultHook(lxdAuditResults models.MeshCheckResults) error {
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
