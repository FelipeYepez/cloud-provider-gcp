load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "packages",
    srcs = [
        "doc.go",
        "external.go",
        "golist.go",
        "golist_overlay.go",
        "loadmode_string.go",
        "packages.go",
        "visit.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/golang.org/x/tools/go/packages",
    importpath = "golang.org/x/tools/go/packages",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/sync/errgroup",
        "//vendor/golang.org/x/tools/go/gcexportdata",
        "//vendor/golang.org/x/tools/internal/gocommand",
        "//vendor/golang.org/x/tools/internal/packagesinternal",
        "//vendor/golang.org/x/tools/internal/typesinternal",
    ],
)
