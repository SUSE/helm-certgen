package certgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCertConfig(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/cert.yaml")
	assert.Nil(t, err, "No error was expected")
	assert.Equal(t, "go-demo-api-certificates", gc.Name, "Expected name not returned")
}

func TestGetCertConfigFileNotFound(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/cert1.yaml")
	assert.Contains(t, err.Error(), "no such file or directory", "file not found error expected")
	assert.Nil(t, gc, "Nil object expected to be returned")
}

func TestGetCertConfigInvalidYaml(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/invalidYaml.yaml")
	assert.Contains(t, err.Error(), "yaml: line", "expecting an yaml error back")
	assert.Nil(t, gc, "Nil object expected to be returned")
}

func TestGetCertConfigMultiCert(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/multiCert.yaml")
	assert.Nil(t, err, "No error was expected")
	assert.Equal(t, "go-demo-api-certificates", gc.Name, "Expected name not returned")
	assert.Equal(t, 2, len(gc.Spec), "Two spec object expected")
}

func TestGetCertificateObjectsSingleSpec(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/cert.yaml")
	assert.Nil(t, err, "No error was expected")
	cObjList := gc.GetCertificateObjects()
	assert.Equal(t, 1, len(cObjList), "Exepcted one object to be returned from GetCertificateObjects")
}
func TestGetCertificateObjectsMultiSpec(t *testing.T) {
	gc, err := GetCertConfig("../../testdata/certconfig/multiCert.yaml")
	assert.Nil(t, err, "No error was expected")
	assert.Equal(t, "go-demo-api-certificates", gc.Name, "Expected name not returned")
	assert.Equal(t, 2, len(gc.Spec), "Two spec object expected")
	cObjList := gc.GetCertificateObjects()
	assert.Equal(t, 2, len(cObjList), "Exepcted two object to be returned from GetCertificateObjects")
}
