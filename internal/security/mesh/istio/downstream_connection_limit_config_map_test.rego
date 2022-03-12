package istio

test_downstream_connection_limit_config_map{
a:= policy_eval with input as {
     "apiVersion": "v1",
     "data":
       {
         "custom_bootstrap.json": "sss"
       }
     ,
     "kind": "ConfigMap",
     "metadata": {
       "creationTimestamp": "2022-02-03T09:50:01Z",
       "name": "istio-custom-bootstrap-config",
       "namespace": "default"
     }
   }
 a.match
}

test_downstream_connection_no_limit_config_map{
a:= policy_eval with input as {
     "apiVersion": "v1",
     "data":
       {
         "ca": "sss"
       }
     ,
     "kind": "ConfigMap",
     "metadata": {
       "creationTimestamp": "2022-02-03T09:50:01Z",
       "name": "istio-custom-bootstrap-config",
       "namespace": "default"
     }
   }
 not a.match
}