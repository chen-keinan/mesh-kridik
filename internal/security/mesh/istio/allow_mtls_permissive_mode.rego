package istio

policy_eval = {"namespace":namespace_name,"match":allow_policy} {
    count(input.items) > 0
 	namespace_name = input.namespace
	some i
	input.items[i].kind == "PeerAuthentication"
	mtlsMode := input.items[i].spec.mtls.mode
	allow_policy = mtlsMode ==  "STRICT"
  }

  policy_eval = {"namespace":namespace_name,"match":allow_policy} {
    count(input.items) == 0
 	namespace_name = input.namespace
    allow_policy = false
  }

