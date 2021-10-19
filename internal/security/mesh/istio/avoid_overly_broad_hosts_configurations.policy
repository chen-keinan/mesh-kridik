package istio

  host_overlay(obj) {
  	has_key(obj,"*")
 }

  policy_eval = {"gateway":gateway_name,"namespace":namespace_name,"match":match_policy} {
    namespace_name = input.metadata.namespace
    gateway_name = input.metadata.name
    match_policy = is_host_overlayed
  }

 default is_host_overlayed = false
 is_host_overlayed {
     input.kind == "Gateway"
     some i
     host_overlay(input.spec.servers[i].hosts)
 }

  has_key(x, a) {
	some p
    a != x[p]
   }