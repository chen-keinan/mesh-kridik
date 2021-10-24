package istio

policy_eval = {"match":allow_policy} {
 	allow_policy = using_3rd_party_token
  }

  default using_3rd_party_token = false
  using_3rd_party_token {
  	input.kind == "TokenRequest"
	input.name == "serviceaccounts/token"
    input.group == "authentication.k8s.io"
    input.namespaced == true
    some i
 	input.verbs[i] ==  "create"
  }
