package namespaces

import (
	"k8s.io/kubernetes/pkg/api"
)

type NamespaceObj struct {
	api.Namespace
}

func New(namespace string) *NamespaceObj {
	na := NamespaceObj{
		Namespace: api.Namespace{},
	}
	na.Namespace.Name = namespace
	return &na
}
