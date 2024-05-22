# Add AWS Account

## Prerequisites

If you haven't already, [install qleetctl](/guides/install-qleetctl), the Qleet
command line tool.

## Add Account

In order to allow Cross-Account access for Qleet to your AWS Account. You will
need the Qleet Account ID which can be exported before any relevant commands
via:

```bash
export qleetAwsAccountId=983530947477
```

You will need this account for specifying the `--aws-account-id` parameter where
required using the [Advanced AWS Setup Guide](../threeport/aws/advanced-aws-setup.md).

## Next Steps

Now that you have your AWS account connected to your Threeport control plane,
you can spin up a new Kubernetes cluster in your AWS account using our [Remote
Kubernetes Runtime guide](../threeport/kubernetes-runtime/remote-kubernetes-runtime.md).

