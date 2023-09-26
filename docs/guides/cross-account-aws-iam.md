# Cross-Account AWS IAM Permissions

Appropriate permissions must be configured before Qleet can manage resources in your AWS account
on your behalf. These steps walk you through how to do so via the AWS CLI based on [documented best-practice](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-cross-account-resource-access.html).

## Create a Policy

Download the policy document:

```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/threeport-admin-iam-policy.json
```

Create the policy in AWS:

```bash
aws iam create-policy \
    --policy-name QleetServiceAccount \
    --policy-document file://threeport-admin-iam-policy.json \
    --description "Allow creation and deletion of resources for Threeport"
```

Take note of the `Arn` that is returned to you when you create the policy.
We'll use that to attach the policy to a role.


## Create a Role

Download the trust policy document:

```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/resource-manager-threeport-trust-policy.json
```

```bash
aws iam create-role \
    --role-name resource-manager-threeport \
    --assume-role-policy-document file://resource-manager-threeport-trust-policy.json
```

## Attach Policy to Role

Substitute the policy ARN that was returned to you when you created the policy above:

```bash
aws iam attach-role-policy \
    --role-name resource-manager-threeport \
    --policy-arn [policy ARN]
```

Qleet is now able to manage resources within your AWS account on your behalf.