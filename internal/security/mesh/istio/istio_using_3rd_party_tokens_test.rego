package istio

test_using_3rd_party_token_with {
a:= policy_eval with input as {
    "kind": "TokenRequest",
    "namespaced": true,
    "name": "serviceaccounts/token",
    "group": "authentication.k8s.io",
    "verbs": [
      "create"
    ]
  }
  a.match
}

test_using_3rd_party_token_with {
a:= policy_eval with input as {
    "kind": "TokenRequest",
    "namespaced": true,
    "name": "serviceaccounts/token",
    "group": "authentication.k8s.io",
    "verbs": [
      "create"
    ]
  }
  a.match
}

test_using_3rd_party_token_with_no_name {
a:= policy_eval with input as {
    "kind": "TokenRequest",
    "namespaced": true,
     "group": "authentication.k8s.io",
    "verbs": [
      "create"
    ]
  }
  not a.match
}

test_using_3rd_party_token_not_namespaced {
a:= policy_eval with input as {
    "kind": "TokenRequest",
     "name": "serviceaccounts/token",
    "group": "authentication.k8s.io",
    "verbs": [
      "create"
    ]
  }
  not a.match
}

test_using_3rd_party_token_with_no_create {
a:= policy_eval with input as {
    "kind": "TokenRequest",
    "namespaced": true,
    "name": "serviceaccounts/token",
    "group": "authentication.k8s.io",
    "verbs": [
      "view"
    ]
  }
  not a.match
}