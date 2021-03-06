[![Go Report Card](https://goreportcard.com/badge/github.com/chen-keinan/mesh-kridik)](https://goreportcard.com/report/github.com/chen-keinan/mesh-kridik)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/chen-keinan/mesh-kridik/blob/main/LICENSE)
[![Go](https://github.com/chen-keinan/mesh-kridik/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/chen-keinan/mesh-kridik/actions/workflows/go.yml)
<img src="./pkg/img/coverage_badge.png" alt="test coverage badge">
[![Gitter](https://badges.gitter.im/beacon-sec/mesh-kridik.svg)](https://gitter.im/beacon-sec/mesh-kridik?source=orgpage)
<br><img src="./pkg/img/mesh_kridik_img.png" width="350" alt="mesh-kridik logo"><br>

# mesh-kridik
Enhance your Kubernetes service mesh security !!

mesh-kridik is an open-source security checker that performs various security checks on a Kubernetes cluster with istio service mesh and outputs a security report

The security checks tests are the full implementation of [istio security best practices](https://istio.io/latest/docs/ops/best-practices/security/) <br>

The security checks performed on a Kubernetes cluster with istio service mesh and is leveraged by OPA (Open Policy Agent) to enforce security rules, and the output audit report includes:
the root cause of the security issue  and proposed remediation for the security issue.

![mesh kridik](./pkg/img/mesh-kridik-repo.gif)

* [Requirements](#requirements)
* [Installation](#installation)
* [Quick Start](#quick-start)
* [Istio Security Checks](#istio-security-checks)
* [User Plugin Usage](#user-plugin-usage)
* [Supported Spec](#supported-spec)
* [Contribution](#Contribution)

## Requirements
- Go 1.16+
- jq
- istio

## Installation

```shell
git clone https://github.com/chen-keinan/mesh-kridik
cd mesh-kridik
make build
```

- Note: mesh-kridik require root user to be executed

## Quick Start

Execute Mesh-Kridik without any flags , execute all tests
```shell
 ./mesh-kridik 

```

Execute mesh-kridik  with flags , execute test on demand

```shell
Usage: mesh-kridik [--version] [--help] <command> [<args>]

Available commands are:
  -r , --report : run security checks and generate remediation report
  -i , --include: execute only specific security check, example -i=1.1
  -e , --exclude: ignore specific security check,  example -e=1.1,2.0
 ```

Execute tests and generate failure tests report and it remediation's


```
./mesh-kridik -r
``` 

## Istio Security Checks
<table style="width:600px; font-size:10px;">
<tr>
    <th style="width:100px">Name</th>
    <th style="width:200px">Description</th>
    <th style="width:300px">Impact</th>
</tr>
<tr>
    <td> Mutual TLS </td>
    <td> Istio  Mutual TLS proxies are configured in permissive mode by default </td>
    <td> proxies will accept both mutual TLS and plaintext traffic</td>
</tr>
<tr>
    <td>Istio Safer Authorization Policy Patterns</td>
    <td> Use ALLOW-with-positive-matching or DENY-with-negative-match patterns</td>
    <td>These authorization policy patterns are safer because the worst result in the case of policy mismatch is an unexpected 403 rejection instead of an authorization policy bypass.</td>
</tr>
<tr>
    <td>path normalization in authorization policy</td>
    <td>The enforcement point for authorization policies is the Envoy proxy instead of the usual resource access point in the backend application</td>
    <td>A mismatch can lead to either unexpected rejection or a policy bypass</td>
</tr>
<tr>
    <td>TLS origination for egress traffic</td>
    <td>Use of DestinationRule on service ServiceEntry for egress traffic</td>
    <td>Not using TLS origination for egress traffic to an external service will be send with plain/text</td>
</tr>
<tr>
    <td>Protocol detection</td>
    <td>explicitly declare the service protocol</td>
    <td>miss detection may result in unexpected traffic behavior</td>
</tr>
<tr>
    <td>CNI support</td>
    <td>istio transparent traffic capture</td>
    <td>not al net traffic will not be capture</td>
</tr>
<tr>
    <td>overly broad hosts</td>
    <td>avoid overly broad hosts settings in Gateway</td>
    <td>may cause potential exposure of unexpected domains</td>
</tr>
<tr>
    <td>Restrict Gateway creation privileges</td>
    <td>restrict creation of Gateway resources to trusted cluster administrators</td>
    <td>may cause  creation of gateway by untrusted users</td>
</tr>
<tr>
    <td>Configure a limit on downstream connections</td>
    <td>Update global_downstream_max_connections in the config map according to the number of concurrent connections needed by individual gateway instances in your deployment. Once the limit is reached, Envoy will start rejecting tcp connections</td>
    <td>no limit on the number of downstream connections can cause exploited by a malicious actor</td>
</tr>
<tr>
    <td>Configure third party service account tokens</td>
    <td>It is recommended to configure 3rd party tokens Because the properties of the first party token are less secure</td>
    <td>first party token properties are less secure and might cause authentication breach</td>
</tr>
<tr>
    <td>Control Plane</td>
    <td>Istiod exposes a few unauthenticated plaintext ports for convenience by default</td>
    <td>exposes the XDS service port 15010 and debug port 8080 over unauthenticated  plaintext</td>
</tr>
<tr>
    <td>Data Plane</td>
    <td>The proxy exposes a variety of ports</td>
    <td>The applications running in the same pod as the proxy have access; there is no trust boundary between the sidecar and application</td>
</tr>
<tr>
    <td>Understand traffic capture limitations</td>
    <td>Securing egress traffic by setting the meshConfig.outboundTrafficPolicy.mode </td>
    <td>external service access will not be controlled</td>
</tr>
</table>

## User Plugin Usage (via go plugins)
The Kube-kridik expose a hook for user plugins [Example](https://github.com/chen-keinan/mesh-kridik/tree/master/examples/plugins) :
- **MeshSecurityCheckResultHook** - this hook accepts k8s service mesh security checks results

##### Compile user plugin
```shell
go build -buildmode=plugin -o=~/<plugin folder>/<plugin>.so ~/<plugin folder>/<plugin>.go
```
##### Copy plugin to folder (.kube-kridik folder is created on the 1st startup)
```shell
cp ~/<plugin folder>/<plugin>.so ~/.kube-kridik/plugins/compile/<plugin>.so
```
## Supported Spec
The Kube-kridik support this specifications and can be easily extended:
- The full Istio service mesh best practices [istio security best practices](https://github.com/chen-keinan/mesh-kridik/tree/master/internal/security/mesh/istio)

this specs can be easily extended by amended the spec files under ```~/.mesh-kridik/security/mesh/istio``` folder

## Contribution
- code contribution is welcome !! , contribution with tests and passing linter is more than welcome :)
- /.dev folder include vagrantfile to be used for development : [Dev Instruction](https://github.com/chen-keinan/mesh-kridik/tree/master/.dev)
