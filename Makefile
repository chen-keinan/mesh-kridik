SHELL := /bin/bash

GOCMD=go
MOVESANDBOX=mv ~/vms/mesh-kridikmesh-kridik ~/vms-local/mesh-kridik
GOMOD=$(GOCMD) mod
GOMOCKS=$(GOCMD) generate ./...
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
BINARY_NAME=mesh-kridik
METALINTER=golangci-lint run -v  > lint.xml
GOCOPY=cp mesh-kridik ~/vagrant_file/.

all:test lint build

fmt:
	$(GOCMD) fmt ./...
lint:
	$(GOCMD) get -d github.com/golang/mock/mockgen@v1.6.0
	$(GOCMD) install -v github.com/golang/mock/mockgen
	export GOPATH=/Users/chen.keinan/go
	export PATH=$GOPATH/bin:$PATH
	export PATH=$PATH:/root/go/bin
	(go env GOPATH)/bin generate ./...
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.27.0
	./scripts/lint.sh
tidy:
	$(GOMOD) tidy -v
test:
	$(GOCMD) get -d github.com/golang/mock/mockgen@v1.6.0
	$(GOCMD) install -v github.com/golang/mock/mockgen && export PATH=$GOPATH/bin:$PATH;
	$(GOMOCKS)
	$(GOTEST) ./... -coverprofile coverage.md fmt
	$(GOCMD) tool cover -html=coverage.md -o coverage.html
	$(GOCMD) tool cover  -func coverage.md
test_rego:
	curl -L -o opa https://openpolicyagent.org/downloads/latest/opa_darwin_arm64
	chmod 755 ./opa
	./opa test ./internal/security/mesh/istio/allow_mtls_permissive_mode_test.rego ./internal/security/mesh/istio/allow_mtls_permissive_mode.rego -v
	./opa test ./internal/security/mesh/istio/safer_authorization_policy_pattern_test.rego ./internal/security/mesh/istio/safer_authorization_policy_pattern.rego -v
	./opa test ./internal/security/mesh/istio/pod_capabilities_exist_test.rego ./internal/security/mesh/istio/pod_capabilities_exist.rego -v
	./opa test ./internal/security/mesh/istio/restrict_gateway_creation_privileges_test.rego ./internal/security/mesh/istio/restrict_gateway_creation_privileges.rego -v
	./opa test ./internal/security/mesh/istio/istio_using_3rd_party_tokens_test.rego ./internal/security/mesh/istio/istio_using_3rd_party_tokens.rego -v
	./opa test ./internal/security/mesh/istio/ingress_gateway_patched_downstream_connection_limit_test.rego ./internal/security/mesh/istio/ingress_gateway_patched_downstream_connection_limit.rego -v
	./opa test ./internal/security/mesh/istio/path_normalization_in_authorization_test.rego ./internal/security/mesh/istio/path_normalization_in_authorization.rego -v
	./opa test ./internal/security/mesh/istio/avoid_overly_broad_hosts_configurations_test.rego ./internal/security/mesh/istio/avoid_overly_broad_hosts_configurations.rego -v
	./opa test ./internal/security/mesh/istio/close_port_8008_as_unauthenticate_plaintext_test.rego ./internal/security/mesh/istio/close_port_8008_as_unauthenticate_plaintext.rego -v
	./opa test ./internal/security/mesh/istio/close_port_15010_as_unauthenticate_plaintext_test.rego ./internal/security/mesh/istio/close_port_15010_as_unauthenticate_plaintext.rego -v
	./opa test ./internal/security/mesh/istio/downstream_connection_limit_config_map_test.rego ./internal/security/mesh/istio/downstream_connection_limit_config_map.rego -v

test_rego_actions:
	curl -L -o opa https://openpolicyagent.org/downloads/v0.37.1/opa_linux_amd64_static
	chmod 755 ./opa
	./opa test ./internal/security/mesh/istio/allow_mtls_permissive_mode_test.rego ./internal/security/mesh/istio/allow_mtls_permissive_mode.rego -v
	./opa test ./internal/security/mesh/istio/safer_authorization_policy_pattern_test.rego ./internal/security/mesh/istio/safer_authorization_policy_pattern.rego -v
	./opa test ./internal/security/mesh/istio/pod_capabilities_exist_test.rego ./internal/security/mesh/istio/pod_capabilities_exist.rego -v
	./opa test ./internal/security/mesh/istio/restrict_gateway_creation_privileges_test.rego ./internal/security/mesh/istio/restrict_gateway_creation_privileges.rego -v
	./opa test ./internal/security/mesh/istio/istio_using_3rd_party_tokens_test.rego ./internal/security/mesh/istio/istio_using_3rd_party_tokens.rego -v
	./opa test ./internal/security/mesh/istio/ingress_gateway_patched_downstream_connection_limit_test.rego ./internal/security/mesh/istio/ingress_gateway_patched_downstream_connection_limit.rego -v
	./opa test ./internal/security/mesh/istio/path_normalization_in_authorization_test.rego ./internal/security/mesh/istio/path_normalization_in_authorization.rego -v
	./opa test ./internal/security/mesh/istio/avoid_overly_broad_hosts_configurations_test.rego ./internal/security/mesh/istio/avoid_overly_broad_hosts_configurations.rego -v
	./opa test ./internal/security/mesh/istio/close_port_8008_as_unauthenticate_plaintext_test.rego ./internal/security/mesh/istio/close_port_8008_as_unauthenticate_plaintext.rego -v
	./opa test ./internal/security/mesh/istio/close_port_15010_as_unauthenticate_plaintext_test.rego ./internal/security/mesh/istio/close_port_15010_as_unauthenticate_plaintext.rego -v
	./opa test ./internal/security/mesh/istio/downstream_connection_limit_config_map_test.rego ./internal/security/mesh/istio/downstream_connection_limit_config_map.rego -v

build:
	export PATH=$GOPATH/bin:$PATH;
	export PATH=$PATH:/home/vagrant/go/bin
	export PATH=$PATH:/home/root/go/bin
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v ./cmd/mesh-kridik;
build_local:
	export PATH=$GOPATH/bin:$PATH;
	export PATH=$PATH:/home/vagrant/go/bin
	export PATH=$PATH:/home/root/go/bin
	$(GOBUILD) ./cmd/mesh-kridik;
build_travis:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v ./cmd/mesh-kridik;
build_remote:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v ./cmd/mesh-kridik
	mv mesh-kridik ~/boxes/basic_box/mesh-kridik

build_docker_local:
	docker build -t chenkeinan/mesh-kridik:3 .
	docker push chenkeinan/mesh-kridik:3
dlv:
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./mesh-kridik
build_beb:
	$(GOPACKR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -v -gcflags='-N -l' cmd/mesh-kridik/mesh-kridik.go
	scripts/deb.sh
.PHONY: all build install test
