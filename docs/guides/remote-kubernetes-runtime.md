# Remote Kubernetes Runtime

Threeport supports the management of Kubernetes clusters as remote runtimes.
This pattern may be used to run workloads on separate Kubernetes clusters from the one
that the Threeport API is deployed to.

## Prerequisites

An instance of the Threeport API is required to get started. Follow the [getting started
guide](../getting-started/) to set one up if you have not already done so.

Note that EKS clusters are currently the only supported type of remote runtime.

First, create a workspace on your local filesystem:

```bash
mkdir threeport-test
cd threeport-test
```

To get started, a valid `AwsAccount` object must be created. Here is an example of what this config looks like:

```yaml
AwsAccount:
  Name: default-account
  AccountID: "555555555555"
  DefaultAccount: true

  # option 1: provide explicit configs/credentials
  #DefaultRegion: some-region
  #AccessKeyID: "ASDF"
  #SecretAccessKey: "asdf"

  # option 2: use local AWS configs/credentials
  LocalConfig: /path/to/local/.aws/config
  LocalCredentials: /path/to/local/.aws/credentials
  LocalProfile: default
```

Paste the following command to download `aws-account.yaml`. Open the file and update `AccountID`,
`LocalConfig`, and `LocalCredentials` to the appropriate values for your environment.

```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/aws-account.yaml
```

Once `aws-account.yaml` is prepared, run the following command to create the `AwsAccount`
object in the Threeport API:
```bash
tptctl create aws-account --config aws-account.yaml
```

## Deployment

Kubernetes clusters are represented as `KubernetesRuntime` objects in the Threeport API.
To deploy a new instance of one, we begin by configuring one as follows:

```yaml
KubernetesRuntime:
  Name: eks-remote
  InfraProvider: eks
  InfraProviderAccountName: default-account
  HighAvailability: false
  Location: NorthAmerica:NewYork
```

Paste the following command to download `k8s-runtime.yaml`.
```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/k8s-runtime.yaml
```

Create a `KubernetesRuntime` instance:
```bash
tptctl create kubernetes-runtime --config k8s-runtime.yaml
```

View the status of the deployed kubernetes runtime instance:
```bash
tptctl get kubernetes-runtime-instances
```

Note: if you would like to use
[kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl)
against the cluter where Threeport is
running and you have the [AWS CLI](https://aws.amazon.com/cli/)
installed you can update your kubeconfig
with:

```bash
aws eks update-kubeconfig --name threeport-test
```

## Cleanup


Run the following command to delete the remote kubernetes runtime instance:
```bash
tptctl delete kubernetes-runtime-instance --name eks-remote
```

Clean up the downloaded config files:
```bash
rm aws-account.yaml
rm k8s-runtime.yaml
```