package istio

test_close_port_8008_as_unauthenticate_plaintext_false{
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
               "env": [
                 {
                   "name": "ENABLE_DEBUG_ON_HTTP",
                   "value": false
                 }
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


test_close_port_8008_as_unauthenticate_plaintext_true{
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
               "env": [
                 {
                   "name": "ENABLE_DEBUG_ON_HTTP",
                   "value": true
                 }
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
 not a.match
}