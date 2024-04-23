# Qleet Control Plane

Qleet is a managed threeport offering and therefore qleetctl can also be used in a way similiar to tptctl for your control plane infrastructure.

As part of private preview, please reach out to a member of an engineering team after your account is setup to help kickstart the process. Once we reach a more mature stage, the ability to create a new control plane should be just a qleetctl command away!

## Qleetctl initialization for a control plane

1. Once you get the name for the control plane instance, you can set your current control plane instance via:

    ```bash
    qleetctl config current-control-plane --name=$controlPlaneName
    ```