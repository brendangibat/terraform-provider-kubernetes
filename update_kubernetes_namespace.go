package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateNamespace(d *schema.ResourceData, namespace api.Namespace) {
	if d.Get("metadata.#").(int) > 0 {
		updateMetadata(d.Get("metadata.0").(*schema.ResourceData), namespace.ObjectMeta)
	}
}
