package istio

test_restrict_gateway_creation_privileges_no_gateway {
 a:= policy_eval with input as {
     "apiVersion": "rbac.authorization.k8s.io/v1",
     "kind": "ClusterRole",
     "metadata": {
       "name": "secret-reader"
     },
     "rules": [
       {
         "apiGroups": [
           ""
         ],
         "resources": [
           "gateway1"
         ],
         "verbs": [
           "create",
           "update",
           "delete"
         ]
       }
     ]
   }
 a.match
}


test_restrict_gateway_creation_privileges_no_mutation_verb {
 a:= policy_eval with input as {
     "apiVersion": "rbac.authorization.k8s.io/v1",
     "kind": "ClusterRole",
     "metadata": {
       "name": "secret-reader"
     },
     "rules": [
       {
         "apiGroups": [
           ""
         ],
         "resources": [
           "gateway"
         ],
         "verbs": [
           "aaa",
           "ccc",
           "bbb"
         ]
       }
     ]
   }
 a.match
}