package istio

  host_change_verbs(obj) {
  	has_key(obj,["create","delete","update"])
 }

  policy_eval = {"role":role_name,"match":match_policy} {
    role_name = input.metadata.name
    s:=can_change_gateway
    s == false
    has_resource(input.rules)
    match_policy = false
  }

   policy_eval = {"role":role_name,"match":match_policy} {
    role_name = input.metadata.name
    s:=can_change_gateway
    s == false
    not has_resource(input.rules)
    match_policy = true
  }

  policy_eval = {"role":role_name,"match":match_policy} {
    role_name = input.metadata.name
    s:=can_change_gateway
    s == true
    has_resource(input.rules)
    match_policy = true
  }

 has_resource(obj){
 	some i
 	 obj[i]["resources"]
 }

 default can_change_gateway = false
 can_change_gateway {
     input.kind == "ClusterRole"
     some i
     some p
	 input.rules[i].resources[p] == "gateway"
     not host_change_verbs(input.rules[i].verbs)
 }

 can_change_gateway {
     input.kind == "ClusterRole"
     some i
     some p
	 input.rules[i].resources[p] != "gateway"
 }

  has_key(x, arr_verbs) {
	some r
    some t
    permission_val:= x[r]
    verb = arr_verbs[t]
    permission_val == verb
   }
