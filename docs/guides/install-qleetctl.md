# Install qleetctl

Use the following instructions to install the `qleetctl` command line tool for
using Qleet.

1. Visit the [releases page](https://github.com/qleet/resources/releases) on
   GitHub and download:

    1. `checksums.txt`

    1. the latest pacakge for your computer's architecture

1. Verify the integrity of the downloaded package.
   ```bash
   sha256sum -c --ignore-missing checksums.txt
   ```
1. Extract `qleetctl`.  The examples here are for version v0.1.0 and the Darwin ARM64
   architecture.  Adjust as necessary for the version and architecture you're
   using.
   ```bash
   tar xf qleetctl_v0.1.0_Darwin_arm64.tar.gz
   ```
1. Move `qleetctl` to your path.
   ```bash
   sudo mv qleetctl_v0.1.0_Darwin_arm64/qleetctl /usr/local/bin/
   ```
1. Check the version installed.
   ```bash
   qleetctl version
   ```
1. View help info.
   ```bash
   qleetctl help
   ```
1. Clean up.
   ```bash
   rm checksums.txt qleetctl_v0.1.0_Darwin_arm64.tar.gz
   rm -rf qleetctl_v0.1.0_Darwin_arm64
   ```

## Next Steps

Once you have `qleetctl` installed, you should next follow our [Qleet Account
guide](qleet-account.md) to get your account set up.

