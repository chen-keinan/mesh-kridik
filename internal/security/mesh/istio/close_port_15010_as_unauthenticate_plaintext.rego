package istio

policy_eval = {"match":allow_policy} {
 	allow_policy = expose_port_15010_plaintext
  }

  default expose_port_15010_plaintext = false
  expose_port_15010_plaintext {
  	input.kind == "Deployment"
    some i
    some k
    input.spec.template.spec.containers[i].args[k] == "--grpcAddr="
  }
