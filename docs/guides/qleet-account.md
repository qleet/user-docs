# Qleet Account

## Onboarding

We are currently in private beta and are looking for potential partners to work with.

Please reach out to cameron@qleet.io if you are interested in using Qleet.

We want to ensure that we have the ability to support your use case or otherwise see how we can we can achieve it

Once we determine there is a fit, you will recieve an account invite code.

Now all thats left for you is to [install qleetctl](/guides/install-qleetctl) and setup your account!


## Setup

The following steps create your Qleet account.

1. Run the following command to initialize qleetctl. Note: this needs to be only done once on a fresh qleetctl install.

    ```bash
    qleetctl init
    ```

2. Once initialized, you can run the following command to create your Qleet account.
Use the account name and code from the onboarding process. Lets assume your account is named congobookstore and the invite code you recieved is "7fgz3weqd".

    ```bash
    qleetctl create account -n congobookstore -c 7fgz3weqd -u tim@congobooks.com,allen@congobooks.com
    ```
The additional user id's we provided as a comma seperated list to the command serves as the initial admin users for your Qleet Account. So please ensure you delegate these user id's appropiately. For more information on how to manage users for your Qleet account please refer to the [Qleet User](/guides/qleet-user) guide

3. Once your account is successfully created, you are now all set to use Qleet!
Refer to the other guides on how to manage Qleet Users or reach out to us to get you started on a managed control plane.
