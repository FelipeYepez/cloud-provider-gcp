load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "extensions",
    srcs = [
        "extension.pb.go",
        "extensions.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/google/gnostic-models/extensions",
    importpath = "github.com/google/gnostic-models/extensions",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/google.golang.org/protobuf/proto",
        "//vendor/google.golang.org/protobuf/reflect/protoreflect",
        "//vendor/google.golang.org/protobuf/runtime/protoimpl",
        "//vendor/google.golang.org/protobuf/types/known/anypb",
    ],
)
