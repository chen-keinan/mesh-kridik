package istio

  has_protocol(obj) {
  	has_key(obj,["http","http2","https","tcp","tls","grpc","grpc-web","mongo","mysql","redis"])
 }

  policy_eval = {"service":service_name,"namespace":namespace_name,"match":match_policy} {
    namespace_name = input.metadata.namespace
    service_name = input.metadata.name
    match_policy = protocol_detected
  }

 default protocol_detected = false
 protocol_detected {
     input.kind == "Service"
     some i
     has_protocol(input.spec.ports[i])
 }

  has_key(x, a) {
	some p
    s:=a[p]
    s == x["appProtocol"]
   }