load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "src_lib",
    srcs = ["main.go"],
    importpath = "github.com/newcl/bazel/src",
    visibility = ["//visibility:private"],
)

go_binary(
    name = "src",
    embed = [":src_lib"],
    visibility = ["//visibility:public"],
)
