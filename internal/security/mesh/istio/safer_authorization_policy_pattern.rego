package istio

  has_allow_operation(obj) {
  has_key(obj,["hosts","paths","methods","ports"])
 }
   has_allow_source(obj) {
  has_key(obj,["principals","requestPrincipals","namespaces","ipBlocks","remoteIpBlocks"])
 }

  has_deny_source(obj) {
  	has_key(obj,["notprincipals","notrequestPrincipals","notnamespaces","notipBlocks","notremoteIpBlocks"])
 }

  has_deny_operation(obj) {
  	has_key(obj,["nothosts","notpaths","notmethods","notPorts"])
 }

  policy_eval = {"namespace":namespace_name,"match":allow_policy} {
  	namespace_name = input.namespace
    allow_policy = valid_authz
  }

  policy_eval = {"namespace":namespace_name,"match":allow_policy} {
    namespace_name = input.namespace
    allow_policy = valid_authz
  }

 default valid_authz = false
 valid_authz {
     count(input.items) > 0
     some i
     input.items[i].kind == "AuthorizationPolicy"
     input.items[i].spec.action == "ALLOW"
     some k
     some p
     count(input.items[i].spec.rules[k].to[p].operation) > 0
     has_allow_operation(input.items[i].spec.rules[k].to[p].operation)
     some t
     some r
     count(input.items[i].spec.rules[t].from[r].source) > 0
     has_allow_source(input.items[i].spec.rules[t].from[r].source)
 }

valid_authz {
    count(input.items) > 0
    some i
    input.items[i].kind == "AuthorizationPolicy"
    input.items[i].spec.action == "DENY"
    some k
    some p
    count(input.items[i].spec.rules[k].to[p].operation) > 0
    has_deny_operation(input.items[i].spec.rules[k].to[p].operation)
    some t
    some r
    count(input.items[i].spec.rules[t].from[r].source) > 0
    has_deny_source(input.items[i].spec.rules[t].from[r].source)
}


  has_key(x, a) {
	some p
    s:=a[p]
    x[s]
    }
