load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "proto",
    srcs = [
        "checkinit.go",
        "decode.go",
        "decode_gen.go",
        "doc.go",
        "encode.go",
        "encode_gen.go",
        "equal.go",
        "extension.go",
        "merge.go",
        "messageset.go",
        "proto.go",
        "proto_methods.go",
        "reset.go",
        "size.go",
        "size_gen.go",
        "wrapperopaque.go",
        "wrappers.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/protobuf/proto",
    importpath = "google.golang.org/protobuf/proto",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/google.golang.org/protobuf/encoding/protowire",
        "//vendor/google.golang.org/protobuf/internal/encoding/messageset",
        "//vendor/google.golang.org/protobuf/internal/errors",
        "//vendor/google.golang.org/protobuf/internal/flags",
        "//vendor/google.golang.org/protobuf/internal/genid",
        "//vendor/google.golang.org/protobuf/internal/order",
        "//vendor/google.golang.org/protobuf/internal/pragma",
        "//vendor/google.golang.org/protobuf/internal/strs",
        "//vendor/google.golang.org/protobuf/reflect/protoreflect",
        "//vendor/google.golang.org/protobuf/reflect/protoregistry",
        "//vendor/google.golang.org/protobuf/runtime/protoiface",
    ],
)
