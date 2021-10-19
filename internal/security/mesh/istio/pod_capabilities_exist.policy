package istio

  has_capabilities(obj) {
  	has_key(obj,"NET_RAW")
    has_key(obj,"NET_ADMIN")
 }

  policy_eval = {"pod":pod_name,"namespace":namespace_name,"match":match_policy} {
    namespace_name = input.metadata.namespace
    pod_name = input.metadata.name
    match_policy = capabilities_detected
  }

 default capabilities_detected = false
 capabilities_detected {
     input.kind == "Pod"
     some i
     has_capabilities(input.spec.containers[i].securityContext.capabilities)
 }

  has_key(x, a) {
	some p
    a == x.add[p]
   }