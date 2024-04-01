load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/loganasherjones/gazelle-nested-embed
gazelle(name = "gazelle")

go_library(
    name = "gazelle-nested-embed_lib",
    srcs = ["main.go"],
    importpath = "github.com/loganasherjones/gazelle-nested-embed",
    visibility = ["//visibility:private"],
    deps = [
        "//broken",
        "//working",
    ],
)

go_binary(
    name = "gazelle-nested-embed",
    embed = [":gazelle-nested-embed_lib"],
    visibility = ["//visibility:public"],
)
