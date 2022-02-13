package istio

test_with_path_normalization_in_authorization {
a:= policy_eval with input as {
     "apiVersion": "security.istio.io/v1beta1",
     "kind": "IstioOperator",
     "namespace": "test",
     "metadata": {
       "name": "operator_name"
     },
     "spec": {
       "meshConfig": {
       "pathNormalization":{
       "normalization":"DECODE_AND_MERGE_SLASHES"
       }
       }
     }
   }
 a.match
}


test_no_path_normalization_in_authorization {
a:= policy_eval with input as {
     "apiVersion": "security.istio.io/v1beta1",
     "kind": "IstioOperator",
     "namespace": "test",
     "metadata": {
       "name": "operator_name"
     },
     "spec": {
       "meshConfig": {
       }
     }
   }
 not a.match
}
