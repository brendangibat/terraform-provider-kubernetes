package build

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func Namespace(d *schema.ResourceData, version string) *api.Namespace {
	namespace := &api.Namespace{
		Spec: NamespaceSpec(d.Get("spec").([]interface{})),
	}

	namespace.Kind = "Namespace"
	namespace.APIVersion = version

	populateMetadata(&namespace.ObjectMeta, d.Get("metadata").([]interface{}))

	return namespace
}

func NamespaceSpec(namespaceSpecs []interface{}) api.NamespaceSpec {
	if len(namespaceSpecs) == 0 {
		return api.NamespaceSpec{}
	}
	userNamespaceSpec := namespaceSpecs[0].(map[string]interface{})

	namespaceSpec := api.NamespaceSpec{}

	if _, ok := userNamespaceSpec["finalizers"]; ok {
		strFinalizers := convertListToStringArray(userNamespaceSpec["finalizers"].([]interface{}))
		var finalizers []api.FinalizerName
		for _, f := range strFinalizers {
			finalizers = append(finalizers, api.FinalizerName(f))
		}
		namespaceSpec.Finalizers = finalizers
	}

	return namespaceSpec
}
