
## How it works:
You can invoke the plugin with a command like this:
```
helm certgen sample/cert-demo --namespace cert-demo
```
On invoking this command, plugin does following
1. It looks for cert.yaml configuration file inside the sample/cert-demo chart definition.
2. As per the configuration in the cert.yaml it creates a TLS cert and a CSR (certificate signing request).
3. It then sends the CSR to Kubernetes certificate singing service.
4. Initially, the CSR request will be in 'Pending' status and be waiting for approval. 
5. Once the CSR request is approved, and the certificate is issued, the plugin downloads the certificate.
6. Then it puts the downloaded certificate and the key that it generated into a secret in specified Kubernetes namespace.