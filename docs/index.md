# Qleet

Qleet is a managed [Threeport](https://threeport.io/) offering.

This site provides user documentation specific to the Qleet platform.  Since it
provides fully managed Threeport control planes, the [Threeport
Documentation](https://docs.threeport.io/) is still applicable and should be
referenced for using it.

### Note

The open source Threeport project provides the `tptctl` command line tool for
interacting with Threeport.  Qleet uses an extension of `tptctl` called `qleetctl`
which includes all the functionality of `tptctl` with added features specific to
the Qleet platform.

Qleet users should always use `qleetctl` instead of `tptctl`.  Any command examples
in the Threeport docs that provide commands for `tptctl`, just substitute qleetctl
and use the same subcommands.

For example...

```bash
tptctl create workload -c my-workload.yaml`
```

simply becomes...

```bash
qleetctl create workload -c my-workload.yaml
```

