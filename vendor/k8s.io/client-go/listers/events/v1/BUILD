load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "events",
    srcs = [
        "event.go",
        "expansion_generated.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/listers/events/v1",
    importpath = "k8s.io/client-go/listers/events/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/events/v1:events",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/client-go/listers",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
