load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1beta1",
    srcs = [
        "batch_client.go",
        "cronjob.go",
        "doc.go",
        "generated_expansion.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/kubernetes/typed/batch/v1beta1",
    importpath = "k8s.io/client-go/kubernetes/typed/batch/v1beta1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/batch/v1beta1",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/applyconfigurations/batch/v1beta1",
        "//vendor/k8s.io/client-go/gentype",
        "//vendor/k8s.io/client-go/kubernetes/scheme",
        "//vendor/k8s.io/client-go/rest",
    ],
)
