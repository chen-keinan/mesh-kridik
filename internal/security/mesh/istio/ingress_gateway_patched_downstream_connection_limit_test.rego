package istio

test_downstream_connection_limit {
a:= policy_eval with input as {
      "apiVersion": "apps/v1",
      "kind": "Deployment",
      "metadata": {
        "name": "istio-ingressgateway",
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "replicas": 3,
        "selector": {
          "matchLabels": {
            "app": "nginx"
          }
        },
        "template": {
          "metadata": {
            "labels": {
              "app": "nginx"
            }
          },
          "spec": {
            "volumes": [
              {
                "configMap": {
                  "name": "istio-custom-bootstrap-config"
                }
              }
            ],
            "containers": [
              {
                "name": "nginx",
                "image": "nginx:1.14.2",
                "ports": [
                  {
                    "containerPort": 80
                  }
                ]
              }
            ]
          }
        }
      }
    }
 a.match
}


test_downstream_connection_limit_not_define {
a:= policy_eval with input as {
      "apiVersion": "apps/v1",
      "kind": "Deployment",
      "metadata": {
        "name": "istio-ingressgateway",
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "replicas": 3,
        "selector": {
          "matchLabels": {
            "app": "nginx"
          }
        },
        "template": {
          "metadata": {
            "labels": {
              "app": "nginx"
            }
          },
          "spec": {
            "volumes": [
              {
                "configMap": {
                }
              }
            ],
            "containers": [
              {
                "name": "nginx",
                "image": "nginx:1.14.2",
                "ports": [
                  {
                    "containerPort": 80
                  }
                ]
              }
            ]
          }
        }
      }
    }
 not a.match
}
