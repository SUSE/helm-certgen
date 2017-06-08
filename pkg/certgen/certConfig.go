package certgen

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

//CertConfig cert.yaml file get read into this struct
type CertConfig struct {
	Name     string                      `yaml:"name"`
	Metadata MetadataInfo                `yaml:"metadata"`
	Spec     map[string]*CertificateInfo `yaml:"spec"`
}

//SecretInfo
type SecretInfo struct {
	Name string
}

//MetadataInfo
type MetadataInfo struct {
	Name   string                 `yaml:"name"`
	Lables map[string]interface{} `yaml:"lables"`
}

//CertificateInfo
type CertificateInfo struct {
	Name string `yaml:"name"`
	//References     []string `yaml:"references"`
	CSR CSRInfo `yaml:"csr"`
	//AlternameNames []string `yaml:"alternamenames"`
	Hosts []string `yaml:"hosts"`
	//KeyLenght      int      `yaml:"keylenght"`
}

// CSR request details
type CSRInfo struct {
	Country                string `yaml:"country"`
	State                  string `yaml:"state"`
	Locality               string `yaml:"locality"`
	OrganisationName       string `yaml:"organisationname"`
	OrganisationalUnitName string `yaml:"organisationalunitname"`
}

// GetCertConfig reads the cert.yaml and create CertConfig object
func GetCertConfig(filename string) (*CertConfig, error) {
	csFileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Failed to read the file @ ", filename)
		return nil, err
	}

	CertConfig := CertConfig{}
	err = yaml.Unmarshal(csFileData, &CertConfig)
	if err != nil {
		fmt.Println("Failed to parse yaml from file @ ", filename)
		return nil, err
	}
	return &CertConfig, nil
}

// GetCertificateObject converts CertConfig to CertificateDetails object
func (c *CertConfig) GetCertificateObjects() []*CertificateObject {
	certObjList := []*CertificateObject{}
	for _, value := range c.Spec {
		certObjList = append(certObjList, NewCertificateObject(value))
	}
	return certObjList

}
