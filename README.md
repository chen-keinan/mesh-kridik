[![Go Report Card](https://goreportcard.com/badge/github.com/chen-keinan/mesh-kridik)](https://goreportcard.com/report/github.com/chen-keinan/mesh-kridik)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/chen-keinan/lxd-probe/blob/main/LICENSE)
[![Build Status](https://travis-ci.com/chen-keinan/mesh-kridik.svg?branch=master)](https://travis-ci.com/chen-keinan/mesh-kridik)
<img src="./pkg/img/coverage_badge.png" alt="test coverage badge">
# mesh-kridik
Enhance your Kubernetes service mesh security !!

mesh-kridik is an open-source security scanner that performs various security checks on a Kubernetes istio service mesh and outputs a security report.

The security checks tests are the full implementation of [istio security best practices](https://istio.io/latest/docs/ops/best-practices/security/) <br>

The security checks performed on a Kubernetes cluster with istio service mesh and leveraged by OPA(Open Policy Agent) to validate results, and the output audit report includes:
the root cause of the security issue  and proposed remediation for the security issue


* [Installation](#installation)
* [Quick Start](#quick-start)
* [Istio Security Checks](#istio-security-checks)



## Installation

```shell
git clone https://github.com/chen-keinan/mesh-kridik
cd mesh-kridik
make build
```

- Note: kube-beacon require root user to be executed

## Quick Start

Execute Mesh-Kridik without any flags , execute all tests
```shell
 ./mesh-kridik 

```

Execute mesh-kridik  with flags , execute test on demand

```shell
Usage: mesh-kridik [--version] [--help] <command> [<args>]

Available commands are:
  -r , --report :  run audit tests and generate remediation report
 ```

Execute tests and generate failure tests report and it remediation's


```
./mesh-kridik -r
```

## Istio Security Checks
<table style="width:600px">
<tr>
    <th style="width:100px">Name</th>
    <th style="width:200px">Description</th>
    <th style="width:300px">impact</th>
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
    <td>Gateways support</td>
    <td>Avoid overly broad hosts configurations</td>
    <td>may cause potential exposure of unexpected domains</td>
</tr>
</table>
