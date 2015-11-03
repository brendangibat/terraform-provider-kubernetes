package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateNamespace(d *schema.ResourceData, namespace api.Namespace) {
		updateMetadata(d.Get("metadata").(map[string]interface{}), namespace.ObjectMeta)
}
