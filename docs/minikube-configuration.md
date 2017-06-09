# Minikube configuration

Minikube is an awesome tool to get started with Kubernetes and to run Kubernetes for the development environment. It is also very powerful because it is easy to configure various services.

## Starting Minikube
As specified in Kubernetes documentation, to configure certificate signer correctly we need to pass set of command line parameters to Kubernetes controller. To set these parameters on Minikube, you can specify following command line parameters to minikube start command

```
minikube start  --extra-config=controller-manager.ClusterSigningCertFile="<UserHome>/.minikube/certs/ca.pem" --extra-config=controller-manager.ClusterSigningKeyFile="<UserHome>/.minikube/certs/ca-key.pem"
```

Make sure to set an appropriate path for UserHome in above command, because that is where .minikube directory will get created and certs will be available inside that directory.

## Adding CA to the trust store

Once your Minikube starts, if you plan on using the certgen plugin to generate TLS certificates, all those certificates will be signed by the CA that we specified via above command. So the services that you deploy with the help of helm cli and the certgen plugin, if you plan to access those services from your host, then it will be a good idea to add that CA to your host's trust store.

If you are on the mac, you can double click the ca.pem file from above path and then start the keychain application and trust the certificate.

If you are on Linux then add this to your trust store by running `update-ca-certificates` command after placing ca.pem file in the ca-certificate path.

## DNS names 

The certgen plugin reads cert.yaml file from chart's directory, if cert.yaml includes hostnames which are external DNS, then we need to make sure those DNS names resolve correctly to the appropriate IP addresses. If you are using load balancers on your Kubernetes deployment and you have a real DNS server on which the domain is configured, then you can create a new host entry for that DNS name.

If you are using Minikube then most likely you would like to do this locally by updating your /etc/hosts entries. You can determine the IP address of your Minikube by running command `minikube ip`. Whatever IP you get back from this, add it to your /etc/hosts files against the DNS names that you want/configured in your cert.yaml file.

For example, if 192.168.99.100 was the ip and `godemoapp.com` is the host entry in your cert.yaml file then your /etc/hosts file should have a line like this
```
192.168.99.100      godemoapp.com
```