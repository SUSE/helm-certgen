package certgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCertificate(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/cert.yaml")
	assert.Nil(t, err, "No error was expected")
	cObjList := gc.GetCertificateObjects()
	for _, cObj := range cObjList {
		cObj.CreateCertificate(cObj.CertificateInfoObj.Name)
		assert.Equal(t, cObj.CertificateInfoObj.Name, cObj.CSR.Name, "Expected name on CSR not found")
	}
}

func TestUpdatedApproval(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/cert.yaml")
	assert.Nil(t, err, "No error was expected")
	cObjList := gc.GetCertificateObjects()
	for _, cObj := range cObjList {
		cObj.CreateCertificate(cObj.CertificateInfoObj.Name)
		cObj.UpdatedApproval()
		assert.Equal(t, "KubectlApprove", cObj.CSR.Status.Conditions[0].Reason, "Expected name on CSR not found")
	}
}
