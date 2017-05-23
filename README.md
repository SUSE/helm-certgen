# helm-certgen

[![Build Status](https://travis-ci.org/SUSE/helm-certgen.svg?branch=master)](https://travis-ci.org/SUSE/helm-certgen)

This is a plugin for [Kubernetes Helm](https://github.com/kubernetes/helm) CLI.
This plugin enables TLS certificate generation for helm charts.

## Plugin installation
There are three steps for installing this plugin successfully.

### 1. Setting up environment
i. Currently in order to use the plugin you need to have **`go`** (1.8) run time installed on your machine so that you can build the plugin binary. To verify that you can run following command and you should see the expected output
```
# go version
go version go1.8 darwin/amd64
```
If you run this on linux machine you will see `go version go1.8 linux/amd64`

ii. Make sure GOBIN path is set, and is included in your PATH
```
export GOBIN=$GOPATH/bin
export PATH=$PATH:${GOBIN}
```

### 2. Build plugin binary
Following command will generate `helm-certgen` binary in GOBIN path
```
go get github.com/SUSE/helm-certgen
```

### 3. Add certgen to helm CLI
In order to install the plugin, we will have to place the wrapper script from plugin directory to HELM_HOME, so run following commands

```
HELM_HOME=$(helm home)
mkdir -p "$HELM_HOME/plugins/certgen"
cp $GOPATH/github.com/SUSE/helm-certgen/plugin/* $HELM_HOME/plugins/certgen/
```