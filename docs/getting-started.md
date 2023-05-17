# Getting Started

In this guide we'll install the tptctl CLI tool on your local machine and then
install the Threeport control plane locally using tptctl.  Then we'll install a
sample app using Threeport.

## Install tptctl

In order to run Threeport locally, you must first have [Docker
Desktop](https://docs.docker.com/desktop/install/mac-install/) installed if on a
Mac or [Docker Engine](https://docs.docker.com/engine/install/) on Linux.

If you are on Ubuntu you can install and add your user to the docker group as
follows:

```bash
sudo apt-get install gcc docker.io
sudo usermod -aG docker $USER
```

Once docker is installed, you can install tptctl by downloading the binary release from Github.

### Binary Install

Currently, tptctl requires that you have the following tools installed on your
local machine.  If you use Homebrew to install, these dependencies will be
handled for you.  Otherwise, ensure these tools are installed first:

* [kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
* [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
* [curl](https://help.ubidots.com/en/articles/2165289-learn-how-to-install-run-curl-on-windows-macosx-linux)
* [wget](https://www.gnu.org/software/wget/)
* [jq](https://github.com/stedolan/jq/wiki/Installation)

Then install tptctl:

```bash
VERSION=$(curl --silent "https://api.github.com/repos/qleet/tptctl/releases/latest" | jq '.tag_name' -r)
wget https://github.com/qleet/tptctl/releases/download/${VERSION}/tptctl_${VERSION}_$(echo $(uname))_$(uname -m).tar.gz -O - |\
    tar -xz && sudo mv tptctl /usr/local/bin/tptctl
```

Usage info for tptctl can be seen as follows:

```bash
tptctl help
```

## Install Threeport

Threeport can be deployed on one of two providers: Kind (for local development) and EKS
(for AWS deployment).

### Kind
To install the Threeport control plane locally:

```bash
tptctl create control-plane --provider kind --name test
```

This will create a local kind Kubernetes cluster and install all of the control
plane components.  It will also register the same kind cluster as the default
compute space cluster for tenant workloads.

### EKS

This section assumes you already have AWS credentials configured on your local machine
with a profile named "threeport".  Follow their
[quickstart page](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html)
for steps on how to do this.

With credentials configured, run the following to install the Threeport control plane in EKS:

```bash
tptctl create control-plane --provider eks --aws-config-profile threeport --name test
```

This will create a remote EKS Kubernetes cluster and install all of the control plane
components.  It will also register the same EKS cluster as the default compute space
cluster for tenant workloads.

### Validate Deployment
To view the pods that constitute the Threeport control plane:

```bash
kubectl get pods -n threeport-control-plane
```

## Deploy A Workload

To deploy a workload using Threeport, you minimally need to create two API objects: a
`WorkloadDefinition` and a `WorkloadInstance`.  We can create both resources with a single
configuration file.

First, create a workspace on your local filesystem:

```bash
mkdir threeport-test
cd threeport-test
```

Download a sample workload config as follows:

```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/wordpress-workload.yaml
```

You now have the workload config on your local filesystem.  If you open the file you'll
see it has a configuration for two resources. Let's dig into what each of them represent:

### Workload Definition

The `WorkloadDefinition` is what it sounds like: a definition for a workload
that can be deployed as many times as you like.  It includes a field
`YAMLDocument` that refers to a file on your filesystem.  Let's download that
file:

```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/wordpress-manifest.yaml
```

That file contains the Kubernetes manifest for the resources required to deploy an
instance of Wordpress.

### Workload Instance
The `WorkloadInstance` refers to the workload definition and actually deploys
the instance of the workload.  It also refers to the cluster which is set up as
the default when we created Threeport above.


### Create Workload
We can now create the workload as follows:

```bash
tptctl create workload --config wordpress-workload.yaml
```

This command calls the the Threeport API to create those two Workload objects.
The API notifies the workload controller via the message broker.  The workload
controller processes the workload definition and creates the workload by calling
the Kubernetes API.

We can use `tptctl` to view deployed workloads:

```bash
tptctl get workloads
```

We can also use `kubectl` to query the Kubernetes API directly. First, set a local
environment variable to the appropriate namespace for the Wordpress application:

```bash
NAMESPACE=$(kubectl get namespace -l app.kubernetes.io/name=wordpress -o=jsonpath='{.items[0].metadata.name}')
```

Confirm the Wordpress application is running with:

```bash
kubectl get pods -l app.kubernetes.io/instance=getting-started -n $NAMESPACE
```

You can now visit the Wordpress application by forwarding a local port to it with this command:

```bash
kubectl port-forward svc/getting-started-wordpress 8080:80 -n $NAMESPACE
```

Now visit the app [here](http://localhost:8080).  It will display the welcome screen of
the Wordpress application.

## Summary

This diagram illustrates the relationships between components introduced in this
guide.

![Threeport Getting Started](img/ThreeportGettingStartedWordpress.png)

When we installed Threeport using `tptctl create threeport` we created a new
control plane on a local kind Kubernetes cluster.

When we installed the sample app using `tptctl create workload` we called the API to
create the two workload objects: a definition and an instance.  The reconciliation for
these objects was carried out by the workload controller which created the necessary
Kubernetes resources via the Kubernetes control plane.

## Clean Up

To delete a workload:
```bash
tptctl delete workload --config wordpress-workload.yaml
```

To uninstall the Threeport control plane locally:

```bash
tptctl delete threeport -n test
```

Remove the test configs from you filesystem:

```bash
cd ../
rm -rf threeport-test
```

