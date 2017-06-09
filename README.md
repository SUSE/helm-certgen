# helm-certgen

[![Build Status](https://travis-ci.org/SUSE/helm-certgen.svg?branch=master)](https://travis-ci.org/SUSE/helm-certgen)

This repository is for `certgen` plugin for [Kubernetes Helm](https://github.com/kubernetes/helm) CLI. This plugin enables TLS certificate generation for helm charts. 

To use the helm-certgen plugin, the helm charts should have a cert.yaml file which outlines the requirements for certificates (the CSR details, alias names etc.,). The plugin will use these information to generate the certificates and place them as a Kubernetes secret in the appropriate namespace.

More details about what happens behind the scene can be found @ [how certgen plugin works](docs/how-it-works.md) 

Since the plugin uses Kubernetes certificate API, we suggest that you use it against a Kubernetes version 1.6.4 or above.

If you are using Minikube then please follow the [minikube configuration](docs/minikube-configuration.md) guide to make sure that your minikube is setup correctly for approving CSR and issuing the certificates.


## Installing certgen plugin
There are three steps for installing this plugin successfully.

### 1. Setting up environment
#### Go runtime
Currently to use the plugin you need to have **`go`** (1.8) run time installed on your machine so that you can build the plugin binary. To verify that you can run the following command and you should see the expected output
```
# go version
go version go1.8 darwin/amd64
```
If you run this on Linux machine, you will see `go version go1.8 linux/amd64`

#### Environment variables
Make sure GOBIN path is set and is included in your PATH
```
export GOBIN=$GOPATH/bin
export PATH=$PATH:${GOBIN}
```

### 2. Build plugin binary
The following command will generate `helm-certgen` binary in GOBIN path
```
mkdir -p ${GOPATH}/src/github.com/SUSE
cd ${GOPATH}/src/github.com/SUSE
git clone https://github.com/SUSE/helm-certgen.git
cd helm-certgen
make build
```

### 3. Add certgen to helm CLI
To install the plugin, we will have to copy the wrapper scripts from project's plugin directory to HELM_HOME, so run following commands

```
HELM_HOME=$(helm home)
mkdir -p ${HELM_HOME}/plugins/certgen
cp $GOPATH/src/github.com/SUSE/helm-certgen/plugin/* $HELM_HOME/plugins/certgen/
```

## Using Plugin
To understand how to use the plugin, we will try to deploy sample/go-demo-app using this plugin and helm cli.

More details on how sample/go-demo-app can be found @ [sample/go-demo-app](sample/go-demo-app/README.md) 


### Deploying the application

To start with, make sure you have helm API (tiller) enabled on your Kubernetes cluster (using `helm init`). Also, make sure you have installed the `certgen` plugin as per [these instructions](#installing-certgen-plugin).

Now let's generate the certificates. For this, we can run the following command:
```
cd $GOPATH/src/github.com/SUSE/helm-certgen
helm certgen generate sample/go-demo-app --namespace sample-go-demo-app
```

This invokes the certgen plugin with `generate` command and provides the path of the of chart directory along with the namespace in which to deploy this chart. Once we do this, the certificates will be generated, and a secret will be created in the specified namespace.

Now we should be able to deploy the application with the help of helm CLI
```
helm install sample/go-demo-app --namespace sample-go-demo-app
```
Once the application goes to running state, you should be able to invoke the application API.

Remember to follow the steps in the minikube configuration guide to make sure that you have correct /etc/hosts entries and also the CA cert is added to your trust store.

## Contribute

This project is still under active development, so you'll likely encounter [issues](https://github.com/SUSE/helm-certgen/issue).

Please participate by filing issues or contributing a pull request!