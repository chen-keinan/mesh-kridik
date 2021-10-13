[![Go Report Card](https://goreportcard.com/badge/github.com/chen-keinan/mesh-kridik)](https://goreportcard.com/report/github.com/chen-keinan/mesh-kridik)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/chen-keinan/lxd-probe/blob/main/LICENSE)
[![Build Status](https://travis-ci.com/chen-keinan/mesh-kridik.svg?branch=master)](https://travis-ci.com/chen-keinan/mesh-kridik)
<img src="./pkg/img/coverage_badge.png" alt="test coverage badge">
# mesh-kridik
Scan your Kubernetes service mesh security !!

mesh-kridik is an open-source security scanner that performs various security checks on a Kubernetes istio service mesh and outputs a security report.

The security checks tests are the full implementation of [istio security best practices](https://istio.io/latest/docs/ops/best-practices/security/) <br>

The security checks performed on a Kubernetes cluster with istio service mesh, and the output audit report includes:
the root cause of the security issue  and proposed remediation for the security issue