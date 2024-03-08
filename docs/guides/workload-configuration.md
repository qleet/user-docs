# Workload Configuration

One of the core purposes of Threeport is to manage the delivery of application
dependencies:

* Infrastructure: The compute, network and storage resources needed to run your
  application and its dependencies.  For example, on AWS these include EC2
  instances, VPCs and EBS volumes.
* Runtime: The software the provides the environment for your software to run
  in.  For example, we treat Kubernetes as a runtime environment and manage it
  to support your applications.
* Workload Dependencies: The software that runs on the infrastructure and
  runtimes to support your workloads.
  * Support Services: Those services installed on the runtime next to your
    workloads that provide important services.  For example, we use
    [cert-manager](https://cert-manager.io/) to provision and manage TLS assets
    for your workload.  Cert-manager runs in the runtime where your workloads
    reside.
  * Managed Services: Those services managed by a 3rd party that your
    workloads rely on.  If you use AWS Relational Database Service to run a
    database for your workload, this falls into this category.
  * Workloads: In a distributed or microservices architecture, you may develop
    many different services that, collectively, serve your end users.  As such,
    you can define other workloads as dependencies so they be managed
    holistically.

### Note

The terms "application" and "workload" are often used somewhat interchangeably.
We use the term "workload" to refer to a single piece of software that can be
deployed independently.  In contrast, an "application" refers to all the
workloads that constitute the software that serves the end user.  This is most
applicable with distributed systems where multiple workloads constitute a
distributed application.  But even most "monolithic" applications have a
distinct database workload that runs separately from the core service.

