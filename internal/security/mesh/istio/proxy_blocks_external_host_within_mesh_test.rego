package istio

test_proxy_blocks_external_host_within_mesh {
a:= policy_eval with input as {
    "apiVersion": "security.istio.io/v1beta1",
    "kind": "IstioOperator",
    "namespace": "test",
    "metadata": {
      "name": "foo"
    },
    "spec": {
      "meshConfig": {
        "outboundTrafficPolicy": {
          "mode": "REGISTRY_ONLY"
        }
      }
    }
  }
 a.match
}

test_proxy_blocks_external_host_not_within_mesh {
a:= policy_eval with input as {
    "apiVersion": "security.istio.io/v1beta1",
    "kind": "IstioOperator",
    "namespace": "test",
    "metadata": {
      "name": "foo"
    },
    "spec": {
      "meshConfig": {
        "outboundTrafficPolicy": {}
      }
    }
  }
 not a.match
}