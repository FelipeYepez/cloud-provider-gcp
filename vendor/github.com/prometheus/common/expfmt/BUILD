load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "expfmt",
    srcs = [
        "decode.go",
        "encode.go",
        "expfmt.go",
        "openmetrics_create.go",
        "text_create.go",
        "text_parse.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/prometheus/common/expfmt",
    importpath = "github.com/prometheus/common/expfmt",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/munnerz/goautoneg",
        "//vendor/github.com/prometheus/client_model/go",
        "//vendor/github.com/prometheus/common/model",
        "//vendor/google.golang.org/protobuf/encoding/protodelim",
        "//vendor/google.golang.org/protobuf/encoding/prototext",
        "//vendor/google.golang.org/protobuf/proto",
        "//vendor/google.golang.org/protobuf/types/known/timestamppb",
    ],
)
