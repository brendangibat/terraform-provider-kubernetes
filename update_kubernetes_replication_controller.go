package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateReplicationController(d *schema.ResourceData, rc api.ReplicationController) {
	if d.Get("spec.#").(int) > 0 {
    	updateReplicationControllerSpec(d.Get("spec.0").(*schema.ResourceData), rc.Spec)
	}

	if d.Get("metadata.#").(int) > 0 {
    	updateMetadata(d.Get("metadata.0").(*schema.ResourceData), rc.ObjectMeta)
	}
}

func updateReplicationControllerSpec(d *schema.ResourceData, rcs api.ReplicationControllerSpec) {
	if rcs.Template != nil {
		if d.Get("template.#").(int) > 0 {
			updatePodTemplateSpec(d.Get("template.0").(*schema.ResourceData), *rcs.Template)
		}
	}
}
