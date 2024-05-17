# Qleet Account

## Onboarding

We are currently in private beta and are looking for potential partners to work with.

Please reach out to [info@qleet.io](mailto:info@qleet.io) if you are interested in using Qleet.

We want to ensure that we have the ability to support your use case or otherwise
see how we can we can achieve it.  Once we determine there is a fit, you will
recieve an account invite code.

If you have an invite code, you're ready to [install qleetctl](/guides/install-qleetctl)
and set up your account using the instructions below.

## Setup

Following are the steps create your Qleet account.

1. Run the following command to initialize qleetctl. Note: this needs to be only
done once on a fresh qleetctl install.

    ```bash
    qleetctl init
    ```

1. Once initialized, you can run the following command to create your Qleet account.
Use the account name and invite code from the onboarding process. Let's assume
your account is named `congobookstore` and the invite code you recieved is `7fgz3weqd`.

    ```bash
    qleetctl create account -n congobookstore -c 7fgz3weqd -u tim@congobooks.com,allen@congobooks.com
    ```
The additional user email addresses are provided as a comma seperated list.  The email
addresses are your user IDs.  These emails are for the initial admin
users for your company's Qleet account. So please ensure you delegate these appropiately.
For more information on how to manage users for your Qleet account please refer
to the [Qleet User guide](/guides/qleet-user).

1. Once your account is successfully created, you are now all set to use Qleet!
Refer to the other guides on how to manage Qleet Users or reach out to us to get
you started on a managed Threeport control plane.

## Next Steps

Once your Qleet account is created, you can register and verify your user with
Qleet and log in using the [Qleet Authentication guide](qleet-authentication.md)

