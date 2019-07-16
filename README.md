# Drone Coreupdate

[![Go Report Card](https://goreportcard.com/badge/github.com/derekahn/drone-coreupdate)](https://goreportcard.com/report/github.com/derekahn/drone-coreupdate)

A drone plugin that syncs `yaml` file `version: ${VERSION}` and coreupdate's package `--version` with the latest github or gitlab tag .

> This is just a wrapper around [updateservicectl](https://github.com/coreos/updateservicectl)

## Required Envs

```bash
# Sets the version from the latest git repo's tag
# PLUGIN_GIT_TOKEN can be a secret

# Github Example
export PLUGIN_GIT_API=https://api.github.com/repos/derekahn/drone-coreupdate/tags
export PLUGIN_GIT_HEADER="Authorization"
export PLUGIN_GIT_TOKEN="token 2a19zc584484ahb02b683bvcm1092db3za6p888l"

# Gitlab Example
export PLUGIN_GIT_API=https://gitlab.com/api/v4/projects/101/repository/tags
export PLUGIN_GIT_HEADER=PRIVATE-TOKEN
export PLUGIN_GIT_TOKEN=N0maQBY8qss2L0NiLPhz

# Required for 'updateservicectl'
# These can be set as secrets
export PLUGIN_APP_ID=14830zbd-40ee-uj38-4lhl-11205bdb820z
export PLUGIN_USER=human
export PLUGIN_KEY=x2g1eia2dg29gbkkkbz211c4a893e8e1
export PLUGIN_SERVER=https://coreupdate.com

# Required for 'updateservicectl package [create || upload]'
export PLUGIN_PKG_FILE=some-project-name
export PLUGIN_PKG_SRC=directory_to_be_tarball
```

## Run 🐳 locally

```bash
$ make build

$ make run
```

## Commands

```bash
 Choose a command to run in drone-coreupdate:

  build     Creates a docker image of the app
  clean     Removes the recently built docker image
  install   Installs 🐹 dependencies
  run       Runs the current docker image
  shell     To be executed after `make run` to give you a shell into the running container
```

## How does it work?

This is an idea of what is basically happening under the hood. Not exactly but close enough.

### GET the latest tag from [github](https://developer.github.com/v3/repos/#list-tags) as `$version`:

```bash
$ curl \
  -H "Authorization: token 2a19zc584484ahb02b683bvcm1092db3za6p888l"  \
  https://api.github.com/repos/derekahn/drone-coreupdate/tags
```

or [gitlab](https://docs.gitlab.com/ee/api/tags.html)

```bash
$ curl \
  -H "PRIVATE-TOKEN: N0maQBY8qss2L0NiLPhz"  \
  https://gitlab.com/api/v4/projects/101/repository/tags
```

### Interpolates all files with `${VERSION}` to `$version` in `$PLUGIN_PKG_SRC`/\*\*:

```bash
$ find $PLUGIN_PKG_SRC \( -type d \) \
  -o -type f  \
  -print0 | xargs -0 sed -i 's/${VERSION}/$version/g'
```

### Creates a tarball from `$PLUGIN_PKG_SRC`:

```bash
$ tar -cvzf $PLUGIN_PKG_FILE.$version.tar $PLUGIN_PKG_SRC
```

### [creates 📦](https://coreos.com/products/coreupdate/docs/latest/updatectl-client.html#package-management):

```bash
$ updateservicectl
  --user=$PLUGIN_USER \
  --key=$PLUGIN_KEY \
  --server=$PLUGIN_SERVER \

  package create \
  --app-id=$PLUGIN_APP_ID	\
  --version=$version	\
  --url=$PLUGIN_SERVER/packages/$PLUGIN_PKG_FILE.$version.tar  \
  --file=$PLUGIN_PKG_FILE.$version.tar
```

### uploads 📦:

```bash
$ updateservicectl
  --user=$PLUGIN_USER \
  --key=$PLUGIN_KEY \
  --server=$PLUGIN_SERVER \

  package upload \
  --file=$PLUGIN_PKG_FILE.$version.tar
```
