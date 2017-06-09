# helm-certgen

[![Build Status](https://travis-ci.org/SUSE/helm-certgen.svg?branch=master)](https://travis-ci.org/SUSE/helm-certgen)

This repository is for `certgen` plugin for [Kubernetes Helm](https://github.com/kubernetes/helm) CLI.
This plugin enables TLS certificate generation for helm charts. 

If a chart comes with cert.yaml file, which describes the requirements around certificates,
this plugin can be used to generate those certificates using Kubernetes certificates API.
More details on how certgen plugin works can be found [here](docs/how-it-works.md) 

## Installing certgen plugin
There are three steps for installing this plugin successfully.

### 1. Setting up environment
#### Go runtime
Currently in order to use the plugin you need to have **`go`** (1.8) run time installed on your machine so that you can build the plugin binary. To verify that you can run following command and you should see the expected output
```
# go version
go version go1.8 darwin/amd64
```
If you run this on linux machine you will see `go version go1.8 linux/amd64`

#### Environment variables
Make sure GOBIN path is set, and is included in your PATH
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
cp $GOPATH/src/github.com/SUSE/helm-certgen/plugin/* $HELM_HOME/plugins/certgen/
```


## Using Plugin
To understand how to use the plugin we will try to deploy sample/go-demo-app using this plugin and helm cli.

### sample/go-demo-app:
This is a helm chart for go-demo-app, this chart deployes thgo-demo-app application using
kubernetes templates. When this chart is deployed via helm cli (helm install command), it creates
a deployment and a service endpoint. The applicaiton kubernetes template depend on a secret which is created

The chart definition also contains a cert.yaml which is used  to configure the certificate requirements.
certgen plugin looks at the cert.yaml and using that information it generates appropriate certificates with the
help of Kubernete's certificate service api, and then using the generated certificates it create a secret which is
then used by the application deployment.


go-demo-app is simple "Hello World" web application written in GO


### deploying the application

To start with make sure you have helm api enabled on your kubernetes cluster (using `helm init`).
Also make sure you have installed the certgen plugin as per these instructions.

Now lets generate the certificates for this we can run following commands:
```
helm certgen generate sample/go-demo-app --namespace sample-go-demo-app
```
This command invokes the certgen plugin with generate (subcommand) and provides the path of the of
chart directory long with the namespace in which to deploy this chart.

Once we do this the certificates will be generated and a secret will be created in the specified namespace.

Now we should be able to deploy the application with the help of helm CLI
```
helm install sample/go-demo-app --namespace sample-go-demo-app
```
Once the application goes to running state you should be able to invoke the application API.

Remember to follow the steps in the minikube configuration guide to make sure that you have 
correct /etc/hosts entries and also the CA cert is added to your trust store.

## Contribute

This project is still under active development, so you'll likely encounter [issues](https://github.com/helm/monocular/issues).

Please participate by filing issues or contributing a pull request!