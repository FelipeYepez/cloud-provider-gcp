load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "instrumentation",
    srcs = [
        "doc.go",
        "library.go",
        "scope.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/go.opentelemetry.io/otel/sdk/instrumentation",
    importpath = "go.opentelemetry.io/otel/sdk/instrumentation",
    visibility = ["//visibility:public"],
    deps = ["//vendor/go.opentelemetry.io/otel/attribute"],
)
