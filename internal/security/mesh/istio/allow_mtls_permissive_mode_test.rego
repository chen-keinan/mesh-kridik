package istio

test_allow_mts_permission_STRICT {
 a:= policy_eval with input as {
    "apiVersion": "security.istio.io/v1beta1",
    "namespace": "test",
    "items": [
      {
        "kind": "PeerAuthentication",
        "metadata": {
          "name": "default"
        },
        "spec": {
          "mtls": {
            "mode": "STRICT"
          }
        }
      }
    ]
  }
 a.match
}

test_allow_mts_permission_PERMISSIVE {
 a:= policy_eval with input as {
    "apiVersion": "security.istio.io/v1beta1",
    "namespace": "test",
    "items": [
      {
        "kind": "PeerAuthentication",
        "metadata": {
          "name": "default"
        },
        "spec": {
          "mtls": {
            "mode": "PERMISSIVE"
          }
        }
      }
    ]
  }
 not a.match
}

test_allow_mts_permission_no_items {
 a:= policy_eval with input as {
    "apiVersion": "security.istio.io/v1beta1",
    "namespace": "test",
    "items": [
    ]
  }
 not a.match
}