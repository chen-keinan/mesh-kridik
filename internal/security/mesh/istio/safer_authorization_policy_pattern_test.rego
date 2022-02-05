package istio

test_safer_authorization_allow_missing_from {
a:= policy_eval with input as {
     "apiVersion": "security.istio.io/v1beta1",
     "kind": "AuthorizationPolicy",
     "namespace": "test",
     "metadata": {
       "name": "foo"
     },
     "spec": {
       "action": "ALLOW",
       "rules": [
         {
           "to": [
             {
               "operation": {
                 "paths": [
                   "/public"
                 ]
               }
             }
           ]
         }
       ]
     }
   }
 not a.match
}

test_safer_authorization_allow_missing_to {
a:= policy_eval with input as {
     "apiVersion": "security.istio.io/v1beta1",
     "kind": "AuthorizationPolicy",
     "namespace": "test",
     "metadata": {
       "name": "foo"
     },
     "spec": {
       "action": "ALLOW",
       "rules": [
         {
           "from": [
             {
               "source": {
                 "paths": [
                   "/public"
                 ]
               }
             }
           ]
         }
       ]
     }
   }
 not a.match
}

test_safer_authorization_allow_missing_to_and_from {
a:= policy_eval with input as {
     "apiVersion": "security.istio.io/v1beta1",
     "kind": "AuthorizationPolicy",
     "namespace": "test",
     "metadata": {
       "name": "foo"
     },
     "spec": {
       "action": "ALLOW",
       "rules": [
       ]
     }
   }
 not a.match
}

test_safer_authorization_allow {
    a:= policy_eval with input as {
        "apiVersion": "security.istio.io/v1beta1",
        "namespace":"test",
        "items": [
          {
            "kind": "AuthorizationPolicy",
            "namespace": "test",
            "metadata": {
              "name": "foo"
            },
            "spec": {
              "action": "ALLOW",
              "rules": [
                {
                  "from": [
                    {
                      "source": {
                        "principals": [
                          "/public"
                        ]
                      }
                    }
                  ],
                  "to": [
                    {
                      "operation": {
                        "paths": [
                          "/public"
                        ]
                      }
                    }
                  ]
                }
              ]
            }
          }
        ]
      }
 a.match
}

test_safer_authorization_deny {
    a:= policy_eval with input as {
      "apiVersion": "security.istio.io/v1beta1",
      "namespace":"test",
      "items": [
        {
          "kind": "AuthorizationPolicy",
          "namespace": "test",
          "metadata": {
            "name": "foo"
          },
          "spec": {
            "action": "DENY",
            "rules": [
              {
                "from": [
                  {
                    "source": {
                      "notrequestPrincipals": [
                        "/public"
                      ]
                    }
                  }
                ],
                "to": [
                  {
                    "operation": {
                      "notpaths": [
                        "/public"
                      ]
                    }
                  }
                ]
              }
            ]
          }
        }
      ]
    }
 a.match
}