# Getting Started

In this guide we'll install the qleetctl CLI tool on your local machine and then
install the QleetOS control plane locally using qleetctl.  Then we'll install a
sample app using QleetOS.

## Install qleetctl

You can install qleetctl using Homebrew or by downloading the binary release from
Github.

### Homebrew

[Homebrew](https://brew.sh/) offers the simplest install for Mac and Linux:

```bash
brew tap qleet/tap
brew install qleet/tap/qleetctl
```

### Binary Install

Currently, qleetctl requires that you have the following tools installed on your
local machine.  If you use Homebrew to install, these dependencies will be
handled for you.  Otherwise, ensure these tools are installed first:

* [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
* [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
* [curl](https://help.ubidots.com/en/articles/2165289-learn-how-to-install-run-curl-on-windows-macosx-linux)
* [wget](https://www.gnu.org/software/wget/)

Then install qleetctl:

```bash
VERSION=$(curl -sL https://github.com/qleet/qleetctl/releases/ | xmllint -html -xpath '//a[contains(@href, "releases")]/text()' - 2> /dev/null | grep -P '^v' | head -n1)
wget https://github.com/qleet/qleetctl/releases/download/${VERSION}/qleetctl_${VERSION}_$(echo $(uname))_$(uname -m).tar.gz -O - |\
    tar -xz && sudo mv qleetctl /usr/local/bin/qleetctl
```

Usage info for qleetctl can be seen as follows:

```bash
qleetctl help
```

## Install QleetOS

To install the QleetOS control plane locally:

```bash
qleetctl install
```

This will create a local kind Kubernetes cluster and install all of the control
plane components.  It will also register the same kind cluster as the default
compute space cluster for tenant workloads.

To view the pods that constitute the QleetOS control plane:

```bash
kubectl get po -n threeport-control-plane
```

Note: Threeport is the name of the control plane.  You can think of Threeport as
the kernel and QleetOS as a distribution of the operating system.

The QleetOS API is now available at localhost:1323.  Ensure that it is up and
running by opening the Swagger API docs at:
[http://localhost:1323/swagger/index.html](http://localhost:1323/swagger/index.html).

## Deploy A Workload

To get a feel for how the QleetOS works, let's deploy a sample workload using
curl to make calls to the QleetOS API.

First we need to create a workload definition for the sample app:

```bash
curl \
    http://localhost:1323/v0/workload_definitions \
    --data '{"Name":"web3-sample-app","YAMLDocument":"\napiVersion: v1\nkind: Namespace\nmetadata:\n  name: sample-app\n---\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: web3-sample-app-config\n  namespace: sample-app\ndata:\n  rpc.endpoint: https://compatible-greatest-energy.discover.quiknode.pro/47ac872f53b4c944c4000778f004280c9172eda8/\n---\napiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: web3-sample-app\n  namespace: sample-app\nspec:\n  selector:\n    matchLabels:\n      app: web3-sample\n  replicas: 2\n  template:\n    metadata:\n      labels:\n        app: web3-sample\n    spec:\n      containers:\n        - name: web3-sample-app\n          image: ghcr.io/qleet/web3-sample-app:v0.0.8\n          imagePullPolicy: IfNotPresent\n          env:\n            - name: PORT\n              value: '8080'\n            - name: VITE_RPC_ENDPOINT\n              valueFrom:\n                configMapKeyRef:\n                  name: web3-sample-app-config\n                  key: rpc.endpoint\n          ports:\n            - containerPort: 8080\n          resources:\n            requests:\n              cpu: '1m'\n              memory: '6Mi'\n            limits:\n              cpu: '3m'\n              memory: '8Mi'\n      restartPolicy: Always\n---\napiVersion: v1\nkind: Service\nmetadata:\n  name: web3-sample-app\n  namespace: sample-app\nspec:\n  selector:\n    app: web3-sample-app\n  ports:\n    - protocol: TCP\n      port: 80\n      targetPort: 8080\n\n","UserID":1}' \
    --header "Content-Type: application/json" \
    --request POST
```

Next we need to create an instance of the sample app using that definition:

```bash
curl \
    http://localhost:1323/v0/workload_instances \
    --data '{"Name":"web3-sample-app","WorkloadClusterID":1,"WorkloadDefinitionID":1}' \
    --header "Content-Type: application/json" \
    --request POST
```

Now, you can query the Kubernetes API for the local cluster to see the pods for
the sample app running:

```bash
kubectl get po -n sample-app
```

## Uninstall QleetOS

To uninstall the QleetOS control plane locally:

```bash
qleetctl uninstall
```