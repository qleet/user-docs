# Qleet Users

Your company's Qleet account has individual user accounts associated with it.
This guide outlines how to manage those users.

## Admin User

You may recall when you created your Qleet account for the first time, you
supplied a list of user email addresses as part of the account creation process.
These users were assigned Admin roles within the context of your Qleet account.

The Admin role can invite other users to join your company's Qleet account and
manage those users' access with role-based-access-control (RBAC).

## Adding a new user

Follow these steps to add a new user to your Qleet account.

### Prerequisites

If you haven't already, [install qleetctl](/guides/install-qleetctl), the Qleet
command line tool, and [log in](qleet-authentication.md).

### User Invite

A current Admin of the Qleet account will first need to onboard the user id via the following command:

```bash
qleetctl invite user --userid bob@congobooks.com
```

## Next Steps

Once a successful user invite has been created, the new user should
[install qleetctl](install-qleetctl.md).  Next, they can register and verify
their user account, then log in.  New users can use the [Qleet Authentication
guide](qleet-authentication.md) to do so.

If you've added a new user and wish to manage their access, follow the [Qleet
RBAC guide](qleet-rbac.md).

If you're all done with managing users, move on to creating a Threeport control
plane by following our [Contol Plane guide](qleet-control-plane.md).
