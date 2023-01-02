# QleetOS

QleetOS is a distributed operating system for managing complex software
deployments.

If you're managing containerized software with any of the following
requirements, QleetOS will help reduce operational toil, improve developer
velocity and reduce production costs:

* layered dependencies
* autoscaling
* multi-region
* multi-cloud

Developers and DevOps engineers interact with the control plane which deploys
and manages workloads into the compute space.

![QleetOS Architecture](img/QleetOSArchitecture.png)

## Control Plane

The QleetOS control plane exposes a RESTful API that clients use to make changes
to the system.  QleetOS controllers are responsible for managing the state of
the system in response to changes made by clients.  When changes are made to
running software deployments, the controllers connect to the compute space and
make the appropriate updates there.

![QleetOS Control Plane](img/QleetOSControlPlaneArchitecture.png)

This diagram illustrates Workload Controller operations at a high level.
After a client makes a change to a workload, the API sends a notification to the
message broker.  The message broker relays that message to the Workload
Controller.  The Workload Controller does one or more of these three things as
needed:

1. Makes changes to the existing state of the system by connecting to the
   compute space and managing Kubernetes resources there.
2. Makes updates in the API to reflect those changes.
3. Re-queues notifications with the message broker when subsequent
   reconciliation is needed for some part of the system.

## Compute Space

The QleetOS compute space is populated by Kubernetes clusters.  This is where
the software runs.  Those clusters can be run on whichever infrastructure
providers that are supported by QleetOS in as many regions as needed to meet
your apps' requirements.

![QleetOS Compute Space](img/QleetOSComputeSpaceArchitecture.png)

This example shows a Web3 application that has the following dependencies:

* A cluster to run on
* Ingress routing for end-user requests
* TLS termination of incoming HTTPS requests
* An RPC node to gain access to a blockchain

QleetOS is responsible for managing these dependencies.  In other words, you can
deploy an application with no infrastructure in place.  QleetOS will provision a
cluster as needed, install the supporting services like ingress traffic routing
and TLS asset management, install direct dependencies like the RPC node, and
finally deploy the workload itself.

