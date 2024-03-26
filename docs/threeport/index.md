# Threeport

Threeport is an application orchestrator and software delivery control plane.
It allows a user to define a workload and declare its dependencies,
orchestrating the delivery of the workload with all of those dependencies
connected and available.

As an alternative to continuous delivery pipelines that use git as the source
of truth, Threeport stores state in a database instead of git. It leverages
software controllers that access the database and reconcile the desired state
to gracefully manage delivery, eliminating the need for sprawling DevOps tools
and configuration languages.

Furthermore, Threeport provides a unified, global control plane for workloads.
It functions as an orchestration system that manages cloud provider
infrastructure and the software utilizing it across any region, all through a
single, scalable control plane.

Threeport treats the following concerns as application dependencies:

* infrastructure
* container orchestration
* installed supporting services
* infra provider managed services

The Threeport control plane shown below consists of a RESTful API and a
collection of controllers that reconcile the state provided by users.  The
controllers perform reconciliation by interfacing with infrastructure service
provider APIs and the Kubernetes API.

![Threeport Stack](img/ThreeportStack.png)

## What Threeport Is Not

### Threeport is not a Kubernetes Distribution

Kubernetes distributions provide installation of Kubernetes clusters along with
supporting services, or cluster addons.  They provide a way to install
Kubernetes clusters that are ready to accept workloads.  The workflow is
generally as follows:

1. Platform engineers use the Kubernetes distro to install clusters and prepare
   them for use.
2. DevOps sets up CI/CD or GitOps pipelines to deploy into those clusters.
3. Developers push changes to config repos that trigger delivery of workloads to
   the clusters.

Threeport performs cluster install and preparation in response to application
deployments as needed.  The workflow is as follows:

1. Operations installs the Threeport control plane.
2. Operations creates dependency profiles that are not satisfied by the existing
   system defaults.  This could include profiles for clusters, managed services
   or installed supporting services.
3. Developers provide workload configs with dependency declarations to
   Threeport.  Threeport orchestrates the deployment of the application and its
   dependencies.

### Threeport is not a Continuous Integration System

Traditional CI includes automated testing and build processes for software.  The
existing tools and systems used by developers today for this are perfectly
adequate.  Threeport requires no change to these developer workflows.

In order to integrate Threeport, simply add a call to the Threeport control
plane to notify it of a new build of a container image at the end of your CI
process.  Threeport will perform the delivery of the new version into the
appropriate environment/s.

## Comparable Projects

Following is a comparison between Threeport and some other open source projects
to help illustrate where Threeport fits in.

### Radius

[Radius](https://radapp.io/) helps teams manage cloud native application
dependencies.

Similarities:

* Both Threeport and Radius have a strong emphasis on providing developer
  abstractions that allow workloads to be deployed _with_ their dependencies,
  such as managed services like AWS RDS and S3.  Radius' workload and dependency
  management capabilities are more mature than in Threeport.
* Both Threeport and Radius are fundamentally multi-cloud systems.  Threeport
  only supports AWS today - but it is designed to have other cloud provider
  support plugged in.  Radius offers support for AWS and Azure today.
* Both Threeport and Radius aim to provide a platform for collaboration between
  developers and other IT operators.  Developers need ways to smoothly leverage
  the expertise offered by other teams with minimal friction.

Differences:

* Radius is an extension of the Kubernetes control plane.  The Threeport control
  plane is a distinct control plane with its own APIs.  The Threeport control
  plane supports greater scalability and geo-redundancy than Kubernetes so as to
  serve as a global control plane for all clusters under management.
* Radius does not manage Kubernetes clusters.  To get started with Radius, you
  must have a Kubernetes cluster.  In contrast, Threeport manages Kubernetes
  clusters as runtime dependencies.
* Threeport manages support services that must be installed on Kubernetes as
  application dependencies.  Examples include network ingress routing, TLS
  termination and DNS management.  These common support services are installed
  and configured for tenant applications by Threeport.  With Radius, these
  services can be installed but aren't managed as dependencies, per se.

Radius and Threeport have very complimentary characteristics and could be
combined well.

### Crossplane

[Crossplane](https://www.crossplane.io/) provides a framework for building
customizations to the Kubernetes control plane.

Similarities:

* Both Threeport and Crossplane facilitate building custom application
  platforms.
* Threeport manages workload dependencies, such as managed services, as a
  primary function.  Similar functionality can be built out with Crossplane.

Differences:

* Crossplane aims to build custom Kubernetes control planes without needing to
  write code.  This is achieved with compositions that define new APIs with YAML.
  In contrast, platform engineers extend Threeport by writing code.  We believe
  that languages like Go are a better choice for implementing sophisticated
  software systems.  As such, we are working on an SDK that allows users to
  build their custom implementations with Go, rather than with compositions
  defined in YAML.
* Crossplane is an extension of the Kubernetes control plane.  The Threeport control
  plane is a distinct control plane with its own APIs.  The Threeport control
  plane supports greater scalability and geo-redundancy than Kubernetes so as to
  serve as a global control plane for all clusters under management.

Crossplane and Threeport could be used in conjunction by using Threeport to
provision and manage Kubernetes with Crossplane extensions.  However, there are a
lot of overlapping concerns between projects.  Building an application platform
using both projects would introduce more complexity and unclear boundaries.

### ArgoCD

[Argo CD](https://argoproj.github.io/cd/) is a modern Kubernetes-native
continuous delivery system.

Similarities:

* Both ArgoCD and Threeport manage software delivery.

Differences:

* ArgoCD supports various DevOps tools to be used in workflows to execute the
  steps needed to deliver software.  Threeport instead uses software
  controllers to manage software delivery.  With ArgoCD you can get a delivery
  pipeline up and running pretty quickly.  The challenge is maintainability when
  complexity increases.  When using Helm charts with Kustomize overlays for
  sophisticated distributed applications, the complexity overhead can become
  quite a burden.  Threeport advocates using code in a software controller
  instead of config languages in a pipeline.  This means more work up-front and
  changes to the delivery system are a bit more involved.  However, this
  approach improves the maintainability of complex delivery systems.
* ArgoCD generally pulls configuration from Git repos and applies them to
  Kubernetes clusters.  Threeport uses a relational database to store config
  which provides more efficient access to software controllers that need to both
  read and write configuration details.

ArgoCD and Threeport could be used in conjunction by using Threeport to
provision and manage Kubernetes clusters with ArgoCD.  However, similar to
Crossplane, there are a lot of overlapping concerns between the projects.  Using
Crossplane and ArgoCD together make far more sense than using Threeport with
either Crossplane or ArgoCD.

## Summary

Fundamentally, Threeport exists to reduce engineering toil and increase resource
consumption efficiency in delivering software to its users.  This leads to
greater development velocity as well as lowered engineering and infrastructure
costs.

It is designed and built upon the following principles:

* General purpose programming languages like Go are superior to DSLs and
  templates for defining the behavior of complex systems.  Threeport makes it as
  easy as possible to extend it with custom controllers.
* Use progressive disclosure in the abstractions available to users.  Dead
  simple use cases should be trivial to configure and execute.  However, complex
  use cases should be supported by allowing users greater level of
  configurability in the underlying systems when needed.
* Git repos are not great for storing the configuration of complex systems.  For
  a system that is driven by software controllers, a database is more efficient
  for both reads and writes by those controllers.

If you'd like to try it out, visit our [getting started
guide](guides/getting-started/).

If you'd like to learn about the architecture, check out our [architecture
overview](architecture/overview/).

