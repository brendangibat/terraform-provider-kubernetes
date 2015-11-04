package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateNamespace(d *schema.ResourceData, namespace api.Namespace) {
	if len(d.Get("metadata").([]interface{})) > 0 {
		updateMetadata(d.Get("metadata").([]interface{})[0].(map[string]interface{}), namespace.ObjectMeta)
	}
}
