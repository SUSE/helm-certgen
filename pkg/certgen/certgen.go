package certgen

import (
	"github.com/SUSE/helm-certgen/pkg/kube"

	"github.com/SUSE/helm-certgen/pkg/kube/secrets"
	"k8s.io/kubernetes/pkg/apis/certificates"
)

//CertGen type
type CertGen struct {
	Namespace string
	config    Config
	client    *kube.Client
}

//New  creates new CertGen object
func New(namespace string) *CertGen {
	certGen := CertGen{
		Namespace: namespace,
	}
	certGen.client = kube.New(kube.GetConfig(""))
	return &certGen
}

//GenerateCertificate Generates certificates as per the cert.yaml configuration
// It performes following steps:
// 1. Generates local certificate key and CSR
// 2. Sends CSR to kubernetes certificate services
// 3. Tries to approve that CSR
// 4. Once its approved, downloads the certificate
// 5. Creates the secret in provided namespace
func (c *CertGen) GenerateCertificate(cs *CertConfig) {
	// generate certificate key and CSR
	certObjList := cs.GetCertificateObjects()
	for _, certObj := range certObjList {
		// send and approve CSR from kubernetes
		csrResponse := c.createAndApproveCSR(certObj)
		certObj.Cert = csrResponse.Status.Certificate

		// create secret in provided namespace
		c.createSecrets(certObj, cs)
	}
}

func (c *CertGen) createAndApproveCSR(certObj *CertificateObject) *certificates.CertificateSigningRequest {
	certObj.CreateCertificate(certObj.CertificateInfoObj.Name)
	c.client.CreateCSR(certObj.CSR)
	certObj.UpdatedApproval()
	csrResponse := c.client.ApproveCSR(certObj.CSR)
	return csrResponse
}

func (c *CertGen) createSecrets(certObj *CertificateObject, cs *CertConfig) {
	sec := secrets.New(cs.Name, c.Namespace)
	for key, value := range cs.Metadata.Lables {
		sec.AddLables(key, value.(string))
	}
	sec.AddSecret("cert", certObj.Cert)
	sec.AddSecret("key", certObj.CertKey)
	c.client.CheckNamespace(c.Namespace)
	c.client.CreateSecret(sec)
}
