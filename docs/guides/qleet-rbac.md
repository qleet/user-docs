# Qleet RBAC

You can manage user RBAC via [qleetctl](/guides/qleet-rbac)

## Roles and Policies

You can assign user id's to RBAC Roles within Qleet.

These Roles can then be assigned to access policies to give you a fine grained control over what a user can do within the confines of the system.

1. To view all current Roles within your qleet account, you can use the following command:

    ```bash
    qleetctl get roles
    ```

    For a newly created account, you should see something akin to the following:

    ```
    NAME                MEMBERS
    admin               [bob@congobooks.com]
    account-viewer      []
    ```

    Please note that the roles admin and account-viewer are default roles for Qleet. While you can make changes to the members you cannot delete them entirely.

1. Now say you have a use case within your organization where you need a permissions group that can only view the user-invitations within the system

    You can call this role a user-invitation-viewer and can make it via the following command:

    ```bash
    qleetctl create role -n user-invitation-viewer -u alice@congobooks.com
    ```

1. Now that we have a role that is designated for users to view the user-invitations within the system.
It is important to ensure that it is configured with the correct policy.

    You can view all current policies within the system via:
    ```bash
    qleetctl get policies
    ```

    You should see something that resembles the below:
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

    Policies unlike Roles cannot be created by the user via qleetctl, these are instead managed by the system itself.

    In the above example, we can see that the role group admin has access to all API Objects for different HTTP verbs and API versions.

    We are interested in adding our newly created user-invitation-viewer role to the user-invitations policy that is specific to version v0 and HTTP verb GET.

    So we can run the following command to attach the role to our desired policy:

    ```bash
    qleetctl update policy -a v0 -v GET -n user-invitations -r user-invitation-viewer
    ```

    Now all users as part of this Role group will be authorized to access the v0 user-invitations GET endpoint.
