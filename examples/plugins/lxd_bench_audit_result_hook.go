package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/chen-keinan/mesh-kridik/pkg/models"
	"net/http"
	"strings"
)

//LxdBenchAuditResultHook this plugin method accept lxd audit bench results
//event include test data , description , audit, remediation and result
func LxdBenchAuditResultHook(lxdAuditResults models.LxdAuditResults) error {
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
