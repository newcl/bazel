# load the http_archive rule itself
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

# load rules_nixpkgs
http_archive(
    name = "io_tweag_rules_nixpkgs",
    sha256 = "980edfceef2e59e1122d9be6c52413bc298435f0a3d452532b8a48d7562ffd67",
    strip_prefix = "rules_nixpkgs-0.10.0",
    urls = ["https://github.com/tweag/rules_nixpkgs/releases/download/v0.10.0/rules_nixpkgs-0.10.0.tar.gz"],
)

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
    ],
)

http_archive(
    name = "bazel_skylib",
    sha256 = "74d544d96f4a5bb630d465ca8bbcfe231e3594e5aae57e1edbf17a6eb3ca2506",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.3.0/bazel-skylib-1.3.0.tar.gz",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
    ],
)

# http_archive(
#     name = "com_google_protobuf",
#     sha256 = "540200ef1bb101cf3f86f257f7947035313e4e485eea1f7eed9bc99dd0e2cb68",
#     strip_prefix = "protobuf-3.25.0",
#     # latest, as of 2023-11-02
#     urls = [
#         "https://github.com/protocolbuffers/protobuf/archive/v3.25.0.tar.gz",
#     ],
# )

http_archive(
    name = "com_google_protobuf",
    sha256 = "a79d19dcdf9139fa4b81206e318e33d245c4c9da1ffed21c87288ed4380426f9",
    strip_prefix = "protobuf-3.11.4",
    # latest, as of 2020-02-21
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.11.4.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.11.4.tar.gz",
    ],
)

# docker/container

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "b1e80761a8a8243d03ebca8845e9cc1ba6c82ce7c5179ce2b295cd36f7e394bf",
    urls = ["https://github.com/bazelbuild/rules_docker/releases/download/v0.25.0/rules_docker-v0.25.0.tar.gz"],
)



# http_archive(
#     name = "com_google_absl",
#     urls = ["https://github.com/abseil/abseil-cpp/archive/273292d1cfc0a94a65082ee350509af1d113344d.zip"],
#     strip_prefix = "abseil-cpp-273292d1cfc0a94a65082ee350509af1d113344d",
#   )

# load("@bazel_tools//tools/jdk:local_java_repository.bzl", "local_java_repository")

# local_java_repository(
#   name = "additionaljdk",          # Can be used with --java_runtime_version=additionaljdk, --java_runtime_version=11 or --java_runtime_version=additionaljdk_11
#   version = "11",                    # Optional, if not set it is autodetected
#   java_home = "$JAVA_HOME",  # Path to directory containing bin/java
# )


load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
load("@io_tweag_rules_nixpkgs//nixpkgs:repositories.bzl", "rules_nixpkgs_dependencies")

rules_nixpkgs_dependencies()

load("@io_tweag_rules_nixpkgs//nixpkgs:nixpkgs.bzl", "nixpkgs_cc_configure", "nixpkgs_git_repository", "nixpkgs_package")
load("@io_tweag_rules_nixpkgs//nixpkgs:toolchains/go.bzl", "nixpkgs_go_configure")  # optional
load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

############################################################
# Define your own dependencies here using go_repository.
# Else, dependencies declared by rules_go/gazelle will be used.
# The first declaration of an external repository "wins".
############################################################

load("//:deps.bzl", "go_dependencies")

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

go_rules_dependencies()

go_register_toolchains(version = "1.20.5")

gazelle_dependencies()

load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

# http_archive(
#     name = "io_bazel_rules_k8s",
#     strip_prefix = "rules_k8s-0.5",
#     urls = ["https://github.com/bazelbuild/rules_k8s/archive/v0.5.tar.gz"],
#     sha256 = "773aa45f2421a66c8aa651b8cecb8ea51db91799a405bd7b913d77052ac7261a",
# )
# load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")


# load("@io_bazel_rules_k8s//toolchains/kubectl:kubectl_configure.bzl", "kubectl_configure")

# http_file(
#     name="k8s_binary",
#     downloaded_file_path = "kubectl",
#     sha256="0e03ab096163f61ab610b33f37f55709d3af8e16e4dcc1eb682882ef80f96fd5",
#     executable=True,
#     urls=["https://dl.k8s.io/release/v1.29.0/bin/linux/amd64/kubectl"],
# )
# kubectl_configure(name="k8s_config", kubectl_path="@k8s_binary//file")
# k8s_repositories()

# kubectl_configure(name="k8s_config", build_srcs=True)

# kubectl_configure(name="k8s_config", build_srcs=True,
#     k8s_commit = "v1.28.5",
#     # Run wget https://github.com/kubernetes/kubernetes/archive/v1.13.1.tar.gz
#     # to download v1.13.1.tar.gz and run sha256sum on the downloaded archive
#     # to get the value of this attribute.
#     # k8s_sha256 = "677d2a5021c3826a9122de5a9c8827fed4f28352c6abacb336a1a5a007e434b7",
#     k8s_sha256 = "dd6a1faa4566995792e3fabe36617092729d2d83d0450eeb6d395afcd03a3885",
#     # Open the archive downloaded from https://github.com/kubernetes/kubernetes/archive/v1.13.1.tar.gz.
#     # This attribute is the name of the top level directory in that archive.
#     k8s_prefix = "kubernetes-1.28.5"
# )

# k8s_repositories()

#load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
# load("@io_bazel_rules_go//go:deps.bzl", "go_rules_dependencies", "go_register_toolchains")

# go_rules_dependencies()
# go_register_toolchains()

# kubectl_configure(name="k8s_config", build_srcs=True,
#     k8s_commit = "v1.13.1",
#     # Run wget https://github.com/kubernetes/kubernetes/archive/v1.13.1.tar.gz
#     # to download v1.13.1.tar.gz and run sha256sum on the downloaded archive
#     # to get the value of this attribute.
#     k8s_sha256 = "677d2a5021c3826a9122de5a9c8827fed4f28352c6abacb336a1a5a007e434b7",
#     # Open the archive downloaded from https://github.com/kubernetes/kubernetes/archive/v1.13.1.tar.gz.
#     # This attribute is the name of the top level directory in that archive.
#     k8s_prefix = "kubernetes-1.13.1"
# )

# k8s_repositories()

# load("@io_bazel_rules_k8s//k8s:k8s_go_deps.bzl", k8s_go_deps = "deps")

# k8s_go_deps()