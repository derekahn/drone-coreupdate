# Drone Coreupdate

[![Go Report Card](https://goreportcard.com/badge/github.com/derekahn/drone-coreupdate)](https://goreportcard.com/report/github.com/derekahn/drone-coreupdate)

A drone plugin for **creating** and **uploading** packages to coreos coreupdate.


> This is just a wrapper around [updateservicectl](https://github.com/coreos/updateservicectl)

## Required Envs

```bash
# Required for 'updateservicectl'
# These can be configured as drone secrets
export PLUGIN_USER=human
export PLUGIN_KEY=x2g1eia2dg29gbkkkbz211c4a893e8e1
export PLUGIN_SERVER=https://coreupdate.com
export PLUGIN_APP_ID=01468bca-70db-2d5d-9cef-81063caa049x

# Required for 'updateservicectl package [create || upload]'
export PLUGIN_PKG_SRC=directory_to_be_tarball
export PLUGIN_PKG_FILE=some-project
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
