package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateReplicationController(d *schema.ResourceData, rc api.ReplicationController) {
	if len(d.Get("spec").([]interface{})) > 0 {
		updateReplicationControllerSpec(d.Get("spec").([]interface{})[0].(map[string]interface{}), rc.Spec)
	}

	if len(d.Get("metadata").([]interface{})) > 0 {
		updateMetadata(d.Get("metadata").([]interface{})[0].(map[string]interface{}), rc.ObjectMeta)
	}
}

func updateReplicationControllerSpec(d map[string]interface{}, rcs api.ReplicationControllerSpec) {
	if rcs.Template != nil {
		if len(d["template"].([]interface{})) > 0 {
			updatePodTemplateSpec(d["template"].([]interface{})[0].(map[string]interface{}), *rcs.Template)
		}
	}
}
