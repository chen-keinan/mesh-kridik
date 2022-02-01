package istio

policy_eval = {"match":allow_policy} {
 	allow_policy = downstream_connections_exist
  }

  default downstream_connections_exist = false
  downstream_connections_exist {
  	input.kind == "Deployment"
	input.metadata.name == "istio-ingressgateway"
    some i
 	input.spec.template.spec.volumes[i].configMap.name ==  "istio-custom-bootstrap-config"
  }

