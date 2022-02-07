package istio

test_pod_capabilities_not_exist {
a:= policy_eval with input as {
    "apiVersion": "v1",
    "kind": "Pod",
    "metadata": {
      "name": "security-context-demo-4",
       "namespace":"test"
    },
    "spec": {
      "containers": [
        {
          "name": "sec-ctx-4",
          "image": "gcr.io/google-samples/node-hello:1.0",
          "securityContext": {
            "capabilities": {
              "add": [
                "NET_ADMIN",
                "SYS_TIME"
              ]
            }
          }
        }
      ]
    }
  }
 not a.match
}

test_pod_capabilities_exist {
a:= policy_eval with input as {
    "apiVersion": "v1",
    "kind": "Pod",
    "metadata": {
      "name": "security-context-demo-4",
       "namespace":"test"
    },
    "spec": {
      "containers": [
        {
          "name": "sec-ctx-4",
          "image": "gcr.io/google-samples/node-hello:1.0",
          "securityContext": {
            "capabilities": {
              "add": [
                "NET_ADMIN",
                "NET_RAW"
              ]
            }
          }
        }
      ]
    }
  }
 not a.match
}