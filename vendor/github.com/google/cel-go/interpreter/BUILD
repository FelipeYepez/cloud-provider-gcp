load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "interpreter",
    srcs = [
        "activation.go",
        "attribute_patterns.go",
        "attributes.go",
        "decorators.go",
        "dispatcher.go",
        "evalstate.go",
        "interpretable.go",
        "interpreter.go",
        "optimizations.go",
        "planner.go",
        "prune.go",
        "runtimecost.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/google/cel-go/interpreter",
    importpath = "github.com/google/cel-go/interpreter",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/google/cel-go/common",
        "//vendor/github.com/google/cel-go/common/ast",
        "//vendor/github.com/google/cel-go/common/containers",
        "//vendor/github.com/google/cel-go/common/functions",
        "//vendor/github.com/google/cel-go/common/operators",
        "//vendor/github.com/google/cel-go/common/overloads",
        "//vendor/github.com/google/cel-go/common/types",
        "//vendor/github.com/google/cel-go/common/types/ref",
        "//vendor/github.com/google/cel-go/common/types/traits",
    ],
)
