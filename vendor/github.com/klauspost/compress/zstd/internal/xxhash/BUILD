load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "xxhash",
    srcs = [
        "xxhash.go",
        "xxhash_amd64.s",
        "xxhash_arm64.s",
        "xxhash_asm.go",
        "xxhash_other.go",
        "xxhash_safe.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/klauspost/compress/zstd/internal/xxhash",
    importpath = "github.com/klauspost/compress/zstd/internal/xxhash",
    visibility = ["//vendor/github.com/klauspost/compress/zstd:__subpackages__"],
)
