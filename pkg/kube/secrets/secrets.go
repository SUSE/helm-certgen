package secrets

import (
	"k8s.io/kubernetes/pkg/api"
)

type SecretObj struct {
	*api.Secret
}

func New(name, namespace string) *SecretObj {
	sec := SecretObj{
		Secret: &api.Secret{},
	}
	sec.Name = name
	sec.Namespace = namespace
	sec.Data = map[string][]byte{}
	sec.Labels = map[string]string{}
	return &sec
}

func (s *SecretObj) AddLables(key string, value string) {
	s.ObjectMeta.Labels[key] = value
}

func (s *SecretObj) AddSecret(key string, value []byte) {
	s.Data[key] = value
}
