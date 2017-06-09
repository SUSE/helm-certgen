# go-demo-app
go-demo-app is simple "Hello World" web application written in GO.


## Introduction
This is a helm chart for go-demo-app, this chart deployes the go-demo-app application using
kubernetes templates. When this chart is deployed via helm cli (helm install command), it creates
a deployment and a service endpoint. 

The chart definition also contains a cert.yaml which is used  to configure the certificate requirements.
[certgen](https://github.com/SUSE/helm-certgen) plugin looks at the cert.yaml and using that information it generates appropriate certificates with the help of Kubernete's certificate service api. Once the certificate is generated the plugin then
creates a secret which is then used by the application deployment.


