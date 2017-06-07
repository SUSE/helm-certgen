package kube

import (
	"fmt"

	"strings"

	"github.com/SUSE/helm-certgen/pkg/kube/namespaces"
	"github.com/SUSE/helm-certgen/pkg/kube/secrets"
	"github.com/SUSE/helm-certgen/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/kubernetes/pkg/apis/certificates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

// Client represents a client capable of communicating with the Kubernetes API.
type Client struct {
	cmdutil.Factory
	// SchemaCacheDir is the path for loading cached schema.
	SchemaCacheDir string
}

// New create a new Client
func New(config clientcmd.ClientConfig) *Client {
	return &Client{
		Factory:        cmdutil.NewFactory(config),
		SchemaCacheDir: clientcmd.RecommendedSchemaFile,
	}
}

func (c *Client) CreateCSR(req *certificates.CertificateSigningRequest) {
	client, _ := c.ClientSet()
	_, err := client.Certificates().CertificateSigningRequests().Create(req)
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) ApproveCSR(req *certificates.CertificateSigningRequest) *certificates.CertificateSigningRequest {
	client, _ := c.ClientSet()
	_, err := client.Certificates().CertificateSigningRequests().UpdateApproval(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	c.waitForCertToBeIssued(req)
	resp, err := client.Certificates().CertificateSigningRequests().Get(req.Name, metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return resp
}

func (c *Client) waitForCertToBeIssued(req *certificates.CertificateSigningRequest) {
	client, _ := c.ClientSet()
	utils.RetryOperation(30, 1, "Check_Certificate", func() error {
		resp, err := client.Certificates().CertificateSigningRequests().Get(req.Name, metav1.GetOptions{})
		if err != nil {
			fmt.Println(err)
			return err
		}

		if len(resp.Status.Certificate) > 0 {
			return nil
		} else {
			return fmt.Errorf("Certificate not ready yet")
		}
	})
}

func (c *Client) CheckNamespace(namespace string) {
	client, _ := c.ClientSet()
	_, err := client.Core().Namespaces().Get(namespace, metav1.GetOptions{})
	if err != nil &&
		strings.HasPrefix(err.Error(), "namespaces") &&
		strings.HasSuffix(err.Error(), "not found") {
		// namespace doesn't exist already
		// lets create one
		ns := namespaces.New(namespace)

		_, err := client.Core().Namespaces().Create(&ns.Namespace)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}

func (c *Client) CreateSecret(req *secrets.SecretObj) {
	client, _ := c.ClientSet()
	resp, err := client.Core().Secrets(req.Namespace).Create(req.Secret)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Name, " Created successfully")
	}
}
