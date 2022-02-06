package istio

test_detect_by_protocol_exist {
a:= policy_eval with input as {
            "kind": "Service",
            "metadata": {
              "name": "myservice",
              "namespace": "test"
            },
            "spec": {
              "ports": [
                {
                  "number": 3306,
                  "name": "database",
                  "appProtocol": "https"
                },
                {
                  "number": 80,
                  "name": "http-web"
                }
              ]
            }
          }
  a.match
}

test_detect_by_protocol_not_exist {
a:= policy_eval with input as {
            "kind": "Service",
            "metadata": {
              "name": "myservice",
              "namespace": "test"
            },
            "spec": {
              "ports": [
                {
                  "number": 3306,
                  "name": "database"
                },
                {
                  "number": 80,
                  "name": "http-web"
                }
              ]
            }
          }
  not a.match
}