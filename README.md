# Qleet User Documentation

User docs for Qleet hosted Threeport service.

## Local Development

Prerequisistes:

* [python 3](https://docs.python-guide.org/starting/installation/)
* [mkdocs](https://www.mkdocs.org/getting-started/#installation)
* [mkdocs-material theme](https://squidfunk.github.io/mkdocs-material/getting-started/#installation)

Run the server locally:

```bash
mkdocs serve
```

View the site at [http://127.0.0.1:8000](http://127.0.0.1:8000/)


## Release
Run `release` target
```bash
make release
```

### Help

```text
$ make
Usage: make COMMAND
Commands :
help    - List available tasks
deps    - Install dependencies
run     - Run mkdocs
stop    - Stop mkdocs
release - Create and push a new tag
version - Print current version(tag)
```

## Merge Threeport Docs

In order to have the open source Threeport user docs nested in the Qleet docs we
do the following:

* Keep the upstream Threeport docs as a git submodule of Qleet docs
* Use a merge tool to add the Threeport docs into the Qleet docs.  The source
  code for this tool is in the `threeport-merge` directory.

The `threeport-merge` tool does the following:

* Copies all Threeport markdown docs into the `docs/threeport` directory.
* Copies all images from Threeport docs into the `docs/img/threeport` directory.
* Updates all instances of `tptctl` in markdown docs with `qleetctl`.
* Updates all paths to images in markdown docs to point to its new home in Qleet
  docs.
* Updates the `.nav` in Qleet's `mkdocs.yml` config to include the Threeport
  docs `.nav` with updated paths to markdown docs.
* Removes any files found in the threeport-merge config file under `.exclude`.

### threeport-merge Config

Sample config:

```yaml
exclude:
  - docs/threeport/guides/getting-started.md
  - docs/threeport/guides/install-tptctl.md
  - docs/threeport/guides/install-threeport-local.md
  - docs/threeport/guides/install-threeport-aws.md
  - docs/threeport/guides/deploy-workload-local.md
```

If there are any documents in the Threeport user docs that aren't appropriate in
Qleet user docs and should be excluded, add the path to the `.exclude` list and
they will be omitted.  The path provided must be the path to the Threeport
document in the Qleet docs where it is found post-merge, i.e. the path to the
document in the `docs/threeport` directory after `threeport-merge` is run.

### Run Merge

If there are changes in Threeport docs that need to be pulled into the Threeport
user-docs submodule, first pull those changes.

```bash
git submodule update --remote
```

Then run `threeport-merge` with a make target.

```bash
make threeport-merge
```

This make target will:

* Build threeport-merge from source.
* Run threeport-merge and reference the `threeport-merge/merge-config.yaml`
  config.

