load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "value",
    srcs = [
        "allocator.go",
        "doc.go",
        "fields.go",
        "jsontagutil.go",
        "list.go",
        "listreflect.go",
        "listunstructured.go",
        "map.go",
        "mapreflect.go",
        "mapunstructured.go",
        "reflectcache.go",
        "scalar.go",
        "structreflect.go",
        "value.go",
        "valuereflect.go",
        "valueunstructured.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/sigs.k8s.io/structured-merge-diff/v4/value",
    importpath = "sigs.k8s.io/structured-merge-diff/v4/value",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/json-iterator/go",
        "//vendor/sigs.k8s.io/yaml/goyaml.v2:goyaml_v2",
    ],
)
