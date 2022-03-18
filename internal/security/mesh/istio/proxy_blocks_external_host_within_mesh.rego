package istio

policy_eval = {"match":allow_policy} {
 	input.kind == "IstioOperator"
  	allow_policy = input.spec.meshConfig.outboundTrafficPolicy.mode ==  "REGISTRY_ONLY"
}


policy_eval = {"match":allow_policy} {
    input.kind == "IstioOperator"
    not input.spec.meshConfig.outboundTrafficPolicy.mode
    allow_policy= false
}