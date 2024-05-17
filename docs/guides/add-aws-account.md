# Add AWS Account

## Prerequisites

If you haven't already, [install qleetctl](/guides/install-qleetctl), the Qleet
command line tool.

## Add Account

The following steps register and configure your AWS account in Qleet.  Once
registered, you can spin up Kubernetes Runtimes and Workloads in your AWS account
using Qleet.

1. Set user environment variables by creating a file called `config-env-var`

    ```bash
    export region=<region>          # your default AWS region
    export email=<email>            # your user email
    export profile=<profile>        # your AWS profile name
    export accountName=<account>    # your Qleet account name

    # provided by Qleet
    export controlPlaneName=dev
    export qleetAwsAccountId=983530947477
    ```


1. Register your account and configure AWS:

    ```bash
    qleetctl config aws-account \
        --aws-account-name $accountName \
        --aws-region $region \
        --aws-profile $profile \
        --aws-account-id $qleetAwsAccountId \
        --external-runtime-manager-role-name resource-manager-threeport-$controlPlaneName-$accountName
    ```

## Next Steps

Now that you have your AWS account connected to your Threeport control plane,
you can spin up a new Kubernetes cluster in your AWS account using our [Remote
Kubernetes Runtime guide](../threeport/kubernetes-runtime/remote-kubernetes-runtime.md).

