package istio

test_close_port_15010_as_unauthenticate_plaintext_true{
 a:= policy_eval with input as {
        "apiVersion": "apps/v1",
        "kind": "Deployment",
        "metadata": {
          "name": "nginx-deployment",
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
              "containers": [
                {
                  "name": "nginx",
                  "image": "nginx:1.14.2",
                  "args": [
                    "--grpcAddr="
                  ],
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


test_close_port_15010_as_unauthenticate_plaintext_fasle{
 a:= policy_eval with input as {
        "apiVersion": "apps/v1",
        "kind": "Deployment",
        "metadata": {
          "name": "nginx-deployment",
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

