# AWS Permissions

Before standing up a threeport control plane instance in AWS using the `eks`
provider, ensure you have the required IAM permissions.  In some cases, it may
prevent `tptctl` from properly removing resources which can lead to dangling
resources in your AWS account which we _really_ want to prevent.

If your AWS user has the `AdministratorAccess` policy attached, you are all
good.  However, if the AWS account you're using is not your own, and you have a
more restrictive set of permissions, you should follow the instructions below to
ensure you can managed the required resources in AWS.

These instructions assume you have the access necessary to create IAM roles and
user groups.  If you don't, you'll need to contact your account admin to help with
this.

The steps below require you have the [AWS CLI](https://aws.amazon.com/cli/)
installed.

## Create a Policy

Download the policy document:

```bash
curl -O https://raw.githubusercontent.com/threeport/releases/main/samples/threeport-admin-iam-policy.json
```

Create the policy in AWS:

```bash
aws iam create-policy \
    --policy-name ThreeportAdmin \
    --policy-document file://threeport-admin-iam-policy.json \
    --description "Allow creation and deletion of resources for Threeport"
```

Take note of the `Arn` that is returned to you when you create the policy.
We'll use that to attach the policy to a user group.

## Create a User Group

Create a group for managing threeport instances:

```bash
aws iam create-group --group-name threeport-admin
```

## Attach Policy to User Group

Substitue the policy ARN that was returned to you when you created the policy
above:

```bash
aws iam attach-group-policy \
    --group-name threeport-admin \
    --policy-arn [policy ARN]
```

## Add Your User to the Group

If you're unsure what your user name is in AWS you can retrive it with:

```bash
aws sts get-caller-identity
```

That will return the ARN for your user which includes your user name.  For
example if the response is:

```json
{
    "UserId": "AIDAJDPLRKLG7UEXAMPLE",
    "Account": "112233445566",
    "Arn": "arn:aws:iam::112233445566:user/JohnDoe"
}
```

In this case the user name is `JohnDoe`.

Now add the user to the `threeport-admin` group:

```bash
aws iam add-user-to-group --group-name threeport-admin --user-name JohnDoe  # substitute your user name
```

Now, provided the local AWS profile you use is for that user, you'll be all set
to stand up the infra for a threeport control plane instance in AWS.

