# Qleet User Guide

Your Qleet account has users associated with it. This guide outlines how to manage those users.

## Admin User

You may recall when you created your Qleet account for the first time, you supplied a list of user ids as part of the account creation process. These users were assigned Admin roles within the context of your Qleet account.

The admin role can manage other users within the system via RBAC. They can also invite other users to be able to register a user account within your Qleet account.

## Adding a new user
Follow these steps to add a new user to your Qleet account.

### Prerequisites

If you haven't already, [install qleetctl](/guides/install-qleetctl), the Qleet
command line tool.

### User Invite

A current admin of the Qleet account will first need to onboard the user id via the following command:

```bash
qleetctl invite user --userid bob@congobooks.com
```

Once a successful user invite has been created, the user can then use [qleetctl]((/guides/install-qleetctl)) to register there user id.

