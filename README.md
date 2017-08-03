# helm-certgen

[![Build Status](https://travis-ci.org/SUSE/helm-certgen.svg?branch=master)](https://travis-ci.org/SUSE/helm-certgen)

This repository is for `certgen` plugin for [Kubernetes Helm](https://github.com/kubernetes/helm) CLI. This plugin enables TLS certificate generation for helm charts. 

To use the helm-certgen plugin, the helm charts should have a cert.yaml file which outlines the requirements for certificates (the CSR details, alias names etc.,). The plugin will use these information to generate the certificates and place them as a Kubernetes secret in the appropriate namespace.

More details about what happens behind the scene can be found @ [how certgen plugin works](docs/how-it-works.md) 

Since the plugin uses Kubernetes certificate API, we suggest that you use it against a Kubernetes version 1.6.4 or above.

If you are using Minikube then please follow the [minikube configuration](docs/minikube-configuration.md) guide to make sure that your minikube is setup correctly for approving CSR and issuing the certificates.

## Installing the plugin
Download the released version of the plugin for your operating system from the [https://github.com/SUSE/helm-certgen/releases](https://github.com/SUSE/helm-certgen/releases) page. After downloading the package files, extract the archieve locally and then run the helm plugin install command on that path.

```
helm plugin install /tmp/certgen-darwin-amd64-1-0-0-1501786067-53c4719
```
If you would like to build the plugin yourself locally follow the instructions from this @ [readme](docs/build-plugin.md) 

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