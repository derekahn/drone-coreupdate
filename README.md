# Drone Coreupdate

[![Go Report Card](https://goreportcard.com/badge/github.com/derekahn/drone-coreupdate)](https://goreportcard.com/report/github.com/derekahn/drone-coreupdate) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/5a94272a12404567a895eda82d8c54cf)](https://www.codacy.com/app/git.derek/drone-coreupdate?utm_source=github.com&utm_medium=referral&utm_content=derekahn/drone-coreupdate&utm_campaign=Badge_Grade) [![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)

A drone plugin that syncs `yaml` file `version: ${VERSION}` and coreupdate's package `--version` with the latest github or gitlab tag .

> This is just a wrapper around [updateservicectl](https://github.com/coreos/updateservicectl)

## Required Envs

```bash
# Sets the version from the latest git repo's tag
# Github Example (need the latest version as first)
export PLUGIN_GIT_API=https://api.github.com/repos/derekahn/drone-coreupdate/tags
export PLUGIN_GIT_HEADER=Authorization
export PLUGIN_GIT_TOKEN=2a19zc584484ahb02b683bvcm1092db3za6p888l

# Sets the version from the latest git repo's tag
# Gitlab Example (query params 'order_by' required; need latest version as first)
export PLUGIN_GIT_API=https://gitlab.com/api/v4/projects/101/repository/tags?order_by=updated
export PLUGIN_GIT_HEADER=PRIVATE-TOKEN
export PLUGIN_GIT_TOKEN=N0maQBY8qss2L0NiLPhz

# These vars are to set the deployment.image sha256
# We end the URL at '=' so that we can dynamically set the tag at runtime
export PLUGIN_QUAY_API=https://quay.io/api/v1/repository/derekahn/autoapp/tag/?onlyActiveTags=true&specificTag=
export PLUGIN_QUAY_TOKEN=xTaZ12XHXXG0pXLxW67kCWV312J5iwBaMw2UAjOy

# Required for 'updateservicectl'
export PLUGIN_USER=human
export PLUGIN_KEY=x2g1eia2dg29gbkkkbz211c4a893e8e1
export PLUGIN_SERVER=https://coreupdate.com
export PLUGIN_APP_ID=14830zbd-40ee-uj38-4lhl-11205bdb820z

# Required for 'updateservicectl package [create || upload]'
export PLUGIN_PKG_FILE=some-project-name
export PLUGIN_PKG_SRC=directory_to_be_tarball

# Required for 'updateservicectl channel update'
export PLUGIN_CHANNEL=release-me
export PLUGIN_PUBLISH="false"
```

## Run 🐳 locally

```bash
make build

make run
```

## Make Commands

```bash
 Choose a command to run in drone-coreupdate:

  build     Creates a docker image of the app
  clean     Removes the recently built docker image
  install   Installs 🐹 dependencies
  run       Runs the current docker image
  shell     To be executed after `make run` to give you a shell into the running container
```

## How it works

This is an idea of what is basically happening under the hood. Not exactly but close enough.

### GET the latest tag from [github](https://developer.github.com/v3/repos/#list-tags) as `$version`

```bash
$ curl \
  -H "Authorization: 2a19zc584484ahb02b683bvcm1092db3za6p888l"  \
  https://api.github.com/repos/derekahn/drone-coreupdate/tags
```

or [gitlab](https://docs.gitlab.com/ee/api/tags.html)

```bash
$ curl \
  -H "PRIVATE-TOKEN: N0maQBY8qss2L0NiLPhz"  \
  https://gitlab.com/api/v4/projects/101/repository/tags
```

### Interpolates all files with `${VERSION}` to `$version` in `$PLUGIN_PKG_SRC`/\*\*

```bash
$ find $PLUGIN_PKG_SRC \( -type d \) \
  -o -type f  \
  -print0 | xargs -0 sed -i 's/${VERSION}/$version/g'
```

### Creates a tarball from `$PLUGIN_PKG_SRC`

```bash
$ tar -cvzf $PLUGIN_PKG_FILE.$version.tar $PLUGIN_PKG_SRC
```

### [Creates 📦](https://coreos.com/products/coreupdate/docs/latest/updatectl-client.html#package-management)

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

### Uploads 📦

```bash
$ updateservicectl
  --user=$PLUGIN_USER \
  --key=$PLUGIN_KEY \
  --server=$PLUGIN_SERVER \

  package upload \
  --file=$PLUGIN_PKG_FILE.$version.tar
```

### Updates channel

```bash
$ updateservicectl
  --user=$PLUGIN_USER \
  --key=$PLUGIN_KEY \
  --server=$PLUGIN_SERVER \

  channel update \
  --app-id=$PLUGIN_APP_ID	\
  --version=$LATEST_GIT_TAG \
  --channel=$PLUGIN_CHANNEL \
  --publish=$PLUGIN_PUBLISH
```

## Drone Usage

```yaml
---
kind: pipeline
name: default

steps:
  - name: upload
    image: derekahn/drone-coreupdate
    settings:
      app_id: 14830zbd-40ee-uj38-4lhl-11205bdb820z
      user: human
      key: x2g1eia2dg29gbkkkbz211c4a893e8e1
      server: https://coreupdate.com
      git_header: Authorization
      git_api: https://api.github.com/repos/derekahn/drone-coreupdate/tags
      git_token: 2a19zc584484ahb02b683bvcm1092db3za6p888l
      git_api: https://quay.io/api/v1/repository/derekahn/autoapp/tag/?onlyActiveTags=true&specificTag=
      git_token: xTaZ12XHXXG0pXLxW67kCWV312J5iwBaMw2UAjOy
      pkg_src: directory_to_be_tarball
      pkg_file: some-project-name
      channel: release-me
      publish: 'true'
```

### With Secrets

```yaml
---
kind: pipeline
name: default

steps:
  - name: upload
    image: derekahn/drone-coreupdate
    settings:
      git_header: Authorization
      git_api: https://api.github.com/repos/derekahn/drone-coreupdate/tags
      git_token:
        from_secret: git_token
      app_id:
        from_secret: ctl_app_id
      user:
        from_secret: ctl_user
      key:
        from_secret: ctl_key
      server:
        from_secret: ctl_server
      quay_api: https://quay.io/api/v1/repository/derekahn/autoapp/tag/?onlyActiveTags=true&specificTag=
      quay_token:
        from_secret: quay_token
      pkg_src: directory_to_be_tarball
      pkg_file: some-project-name
      channel: release-me
      publish: 'true'
```
