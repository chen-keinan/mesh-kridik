package istio

test_avoid_overly_non_broad_hosts_configurations {
a:= policy_eval with input as {
    "kind": "Gateway",
    "metadata": {
      "name": "guestgateway",
      "namespace": "guestgateway"
    },
    "spec": {
      "selector": {
        "istio": "ingressgateway"
      },
      "servers": [
        {
          "port": {
            "number": 443,
            "name": "https",
            "protocol": "HTTPS"
          },
          "hosts": [
            "*.example.com"
          ],
          "tls": {
            "mode": "SIMPLE"
          }
        }
      ]
    }
  }
 a.match
}


test_avoid_overly_broad_hosts_configurations {
a:= policy_eval with input as {
    "kind": "Gateway",
    "metadata": {
      "name": "guestgateway",
      "namespace": "guestgateway"
    },
    "spec": {
      "selector": {
        "istio": "ingressgateway"
      },
      "servers": [
        {
          "port": {
            "number": 443,
            "name": "https",
            "protocol": "HTTPS"
          },
          "hosts": [
            "*"
          ],
          "tls": {
            "mode": "SIMPLE"
          }
        }
      ]
    }
  }
 not a.match
}
