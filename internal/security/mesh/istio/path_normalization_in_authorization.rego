package istio

policy_eval = {"match":allow_policy,"operator":operator_name} {
 	input.kind == "IstioOperator"
    operator_name = input.metadata.name
 	allow_policy = input.spec.meshConfig.pathNormalization.normalization ==  "DECODE_AND_MERGE_SLASHES"
}

policy_eval = {"match":allow_policy,"operator":operator_name} {
    input.kind == "IstioOperator"
    operator_name = input.metadata.name
    not input.spec.meshConfig.pathNormalization.normalization
    allow_policy= false
}

