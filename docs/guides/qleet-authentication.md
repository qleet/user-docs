# Qleet Authentication

You can use [qleetctl](/guides/install-qleetctl) to authenticate to your [Qleet user account](/guides/qleet-user).



## User Registration

1. Set user password by creating a file called `/tmp/qleet-password-env`:

    ```bash
    echo "export password=<password>" > /tmp/qleet-password-env
    ```

1. Source password by running `source` on the above file

    ```bash
    source /tmp/qleet-password-env
    ```

1. Set user environment variables by creating a file called `config-env-var`

    ```bash
    export region=<region>     # your default AWS region
    export email=<email>       # your user email

    # provided by Qleet
    export controlPlaneName=dev
    export accountName=pos-tech
    ```

1. Source config by running `source` on the above file

    ```bash
    source config-env-var
    ```

1. Register Qleet user

    ```bash
    qleetctl register user --account $accountName --username $email --password $password
    ```

## User Verification
You can verify a successfully registered user as follows:

1. Check your email (and potentially spam folder) for an email from `haris@qleet.io` that looks like this

    ```bash
    Hi, please verify your account by using the following code: 894367, URL: http://localhost:31500/kratos/self-service/verification?code=894367&flow=b7182149-5f2a-40e2-b7e9-be59ce608171
    ```

1. Update your environment variables with the `code` and `flow` values from the email

    ```bash
    export code=894367
    export flowId=b7182149-5f2a-40e2-b7e9-be59ce608171
    ```

1. Verify Qleet account

    ```bash
    qleetctl verify user --account $accountName --code $code --flow $flowId
    ```

1. You can alternatively use the link in the email as is via the following command:

    ```bash
    qleetctl verify user --link "http://localhost:31500/kratos/self-service/verification?code=894367&flow=b7182149-5f2a-40e2-b7e9-be59ce608171"
    ```

## User Login
You can login to a succesfully verified account via the following command:

```bash
qleetctl login user --account $accountName --username $email --password $password
```