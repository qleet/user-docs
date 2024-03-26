# Frequently Asked Questions

## Should I use separate Threeport control planes for dev, staging and prod?

In most cases, we don't recommend this.  It makes promoting changes through the
tiers to production more troublesome.  Managing each tier through a single
control plane provides a smoother experience.

If your organization has different departments or lines of business that you'd
like to segregate and that don't overlap in the workloads they manage, separate
Threeport control planes makes more sense in that case.

## Can I add Extensions to my Qleet-managed Threeport control plane?

For security reasons, we cannot allow our users to run Threeport extensions that
they build on Qleet.  However, we can build extensions for you and install them
on your existing control plane/s.  Shoot us an email at <support@qleet.io> if
you want to explore that option.

Your other option is to use Qleet Enterprise.  This product involves running the
entire Qleet service on your AWS account with full support from Qleet.  In that
case, you are free to internally develop Threeport extensions and add them to
your control plane.  To find out more, email us at <sales@qleet.io>.

More information about Threeport extensions can be found in our [SDK
Introduction document](../threeport/sdk/sdk-intro).

