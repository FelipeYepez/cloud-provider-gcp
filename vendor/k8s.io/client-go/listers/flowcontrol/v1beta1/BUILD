load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1beta1",
    srcs = [
        "expansion_generated.go",
        "flowschema.go",
        "prioritylevelconfiguration.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/listers/flowcontrol/v1beta1",
    importpath = "k8s.io/client-go/listers/flowcontrol/v1beta1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/flowcontrol/v1beta1",
        "//vendor/k8s.io/apimachinery/pkg/labels",
        "//vendor/k8s.io/client-go/listers",
        "//vendor/k8s.io/client-go/tools/cache",
    ],
)
