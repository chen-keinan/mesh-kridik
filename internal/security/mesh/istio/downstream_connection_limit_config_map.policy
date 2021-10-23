package istio

policy_eval = {"match":allow_policy} {
 	allow_policy = downstream_connections_config_map_exist
  }

  default downstream_connections_config_map_exist = false
  downstream_connections_config_map_exist {
  	input.kind == "ConfigMap"
	input.metadata.name == "istio-custom-bootstrap-config"
 	input.data["custom_bootstrap.json"]
  }
