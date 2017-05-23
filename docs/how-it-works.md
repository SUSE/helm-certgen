
## How it works:
You can invoke the plugin with a command like this:
```
helm certgen sample/cert-demo --namespace cert-demo
```
On invoking this command plugin does following
1. It looks for cert.yaml configuration file inside the sample/cert-demo chart definition.
2. As per the configuration in cert.yaml it creates a TLS cert and a CSR (certificate signing request).
3. It then sends the CSR to kubernetes certificate singing service.
4. Initially the CSR request will be in 'Pending' status and waiting for approval. 
4. Once the singing request is approved on kubernetes, it will download 
