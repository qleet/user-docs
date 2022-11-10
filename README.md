# qleetctl

Install and manage instances of the QleetOS.

## Requirements
* [go 1.18](https://go.dev/doc/install)
* [curl](https://help.ubidots.com/en/articles/2165289-learn-how-to-install-run-curl-on-windows-macosx-linux)
* [wget](https://www.gnu.org/software/wget/)
* [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
* [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
* [homebrew](https://brew.sh/) MacOS users

## Install

### Installation from public repository

Download pre-compiled binaries:
set `VERSION` environment variable to latest
```bash
VERSION=$(curl -sL https://github.com/qleet/qleetctl/releases/ | xmllint -html -xpath '//a[contains(@href, "releases")]/text()' - 2> /dev/null | grep -P '^v' | head -n1)
```
or a specific version
```bash
VERSION=v0.0.2
```
then download and install binaries
```bash
wget https://github.com/qleet/qleetctl/releases/download/${VERSION}/qleetctl_${VERSION}_$(echo $(uname))_$(uname -m).tar.gz -O - |\
    tar -xz && sudo mv qleetctl /usr/local/bin/qleetctl
```

### Homebrew
Available for MacOS and Linux.

```bash
brew tap qleet/tap
brew install qleet/tap/qleetctl
```

## Release
Run `release` target
```bash
make release
```

## Quickstart

Install the QleetOS control plane:

```bash
qleetctl install
```

Remove the QleetOS control plane:

```bash
qleetctl uninstall
```

### Help

```text
$ make
Usage: make COMMAND
Commands :
help                - List available tasks
test                - Run tests
build               - Build workload controller binary
get                 - Download and install dependency packages
install             - Install the qleetctl CLI
release             - Create and push a new tag
test-release-local  - Build binaries locally without publishing
update              - Update dependencies to latest versions
version             - Print current version(tag)
```
