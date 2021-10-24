package istio

policy_eval = {"match":allow_policy} {
 	allow_policy = expose_port_8080_plaintext
  }

  default expose_port_8080_plaintext = false
  expose_port_8080_plaintext {
  	input.kind == "Deployment"
    some i
    some k
 	input.spec.template.spec.containers[i].env[k].name ==  "ENABLE_DEBUG_ON_HTTP"
    input.spec.template.spec.containers[i].env[k].value ==  false
  }
