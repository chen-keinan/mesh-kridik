package istio

  has_tls_setting(obj,fields) {
  some i
  has_key(obj,fields[i])
 }


  policy_eval = {"serviceEntry":namespace_name,"match":allow_policy} {
  	namespace_name = input.serviceEntry
    allow_policy = valid_tls_setting
  }

  policy_eval = {"serviceEntry":namespace_name,"match":allow_policy} {
    namespace_name = input.serviceEntry
    allow_policy = valid_tls_setting
  }

 default valid_tls_setting = false
 valid_tls_setting {
     count(input.items) > 0
     some i
     input.items[i].kind == "DestinationRule"
     some k
     count(input.items[i].spec.trafficPolicy.portLevelSettings[k]) > 0
     has_tls_setting(input.items[i].spec.trafficPolicy.portLevelSettings[k].tls,["caCertificates"])
 }

  has_key(x, a) {
    x[a]
    }
