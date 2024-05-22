# Qleet RBAC

You can manage user RBAC via [qleetctl](install-qleetctl.md).

## Roles and Policies

You can assign user emails to RBAC Roles within Qleet.

These Roles can then be assigned to access policies to give you fine grained
control over what a user can do within the system.

1. To view all current Roles within your qleet account, you can use the following command.

    ```bash
    qleetctl get roles
    ```

    For a newly created account, you should see something like this.

    ```
    NAME                MEMBERS
    admin               [bob@congobooks.com]
    account-viewer      []
    ```

    Please note that `admin` and `account-viewer` are default roles for Qleet.
    While you can make changes to the members you cannot delete these roles
    entirely.

1. Now say you have a use case within your organization where you need a
permissions group that can only view the user-invitations within the system

    You can call this role `user-invitation-viewer` and can create it via the
    following command.  With this command we're also assigning the user with
    email `alice@congobooks.com` to that role.

    ```bash
    qleetctl create role -n user-invitation-viewer -u alice@congobooks.com
    ```

1. Now that we have a role that is designated for users to view the user-invitations
within the system, it is important to ensure that it is configured with the correct policy.

    You can view all current policies within the system as follows.

    ```bash
    qleetctl get policies
    ```

    You should see something that resembles this.

    ```bash
    NAME                                 API_VERSION     HTTP_VERB     MEMBERS                     CONTROLPLANE
    accounts                             v0              PATCH         [admin]                     
    accounts                             v0              GET           [admin account-viewer]      
    accounts                             v0              PUT           [admin]                     
    policies                             v0              GET           [admin]                     
    ...
    ...                   
    roles                                v0              GET           [admin]                     
    roles                                v0              POST          [admin]                     
    roles                                v0              PUT           [admin]                     
    user-invitations                     v0              GET           [admin]                     
    user-invitations                     v0              POST          [admin]                     
    user-invitations                     v0              PUT           [admin]                     
    user-invitations                     v0              DELETE        [admin]                     
    user-invitations                     v0              PATCH         [admin]          
    ```

    Unlike Roles, Policies cannot be created by the user via qleetctl,these are
    instead managed by the system itself.

    In the above example, we can see that the `admin` role is a member of each
    policy which gives that role access to all API Objects for each HTTP
    verb and API version.

    For this guide we will add our newly created user-invitation-viewer role
    to the user-invitations policy that is specific to version `v0` and HTTP verb
    `GET`.

    The following command will attach the `user-invitation-viewer` role to the
    `user-invitations` policy.

    ```bash
    qleetctl update policy -a v0 -v GET -n user-invitations -r user-invitation-viewer
    ```

    Now, all users assigned to the `user-invitation-viewer` role will be
    authorized to access the v0 user-invitations API endpoint with GET requests.

## Next Steps

Now that you have your users' access control set up, you can create a Threeport
control plane using our [Control Plane guide](qleet-control-plane.md).

