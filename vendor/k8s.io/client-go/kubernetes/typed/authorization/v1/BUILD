load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "authorization",
    srcs = [
        "authorization_client.go",
        "doc.go",
        "generated_expansion.go",
        "localsubjectaccessreview.go",
        "selfsubjectaccessreview.go",
        "selfsubjectrulesreview.go",
        "subjectaccessreview.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/kubernetes/typed/authorization/v1",
    importpath = "k8s.io/client-go/kubernetes/typed/authorization/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/authorization/v1:authorization",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/client-go/gentype",
        "//vendor/k8s.io/client-go/kubernetes/scheme",
        "//vendor/k8s.io/client-go/rest",
    ],
)
