# QleetOS User Documentation

User docs for QleetOS.

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
