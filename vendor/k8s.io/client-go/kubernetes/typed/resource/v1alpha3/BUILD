load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "v1alpha3",
    srcs = [
        "deviceclass.go",
        "doc.go",
        "generated_expansion.go",
        "podschedulingcontext.go",
        "resource_client.go",
        "resourceclaim.go",
        "resourceclaimtemplate.go",
        "resourceslice.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/kubernetes/typed/resource/v1alpha3",
    importpath = "k8s.io/client-go/kubernetes/typed/resource/v1alpha3",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/resource/v1alpha3",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/types",
        "//vendor/k8s.io/apimachinery/pkg/watch",
        "//vendor/k8s.io/client-go/applyconfigurations/resource/v1alpha3",
        "//vendor/k8s.io/client-go/gentype",
        "//vendor/k8s.io/client-go/kubernetes/scheme",
        "//vendor/k8s.io/client-go/rest",
    ],
)
