load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "node",
    srcs = [
        "doc.go",
        "generated_expansion.go",
        "node_client.go",
        "runtimeclass.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/kubernetes/typed/node/v1",
    importpath = "k8s.io/client-go/kubernetes/typed/node/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/node/v1:node",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/applyconfigurations/node/v1:node",
        "//vendor/k8s.io/client-go/gentype",
        "//vendor/k8s.io/client-go/kubernetes/scheme",
        "//vendor/k8s.io/client-go/rest",
    ],
)
