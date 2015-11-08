package update

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func ReplicationController(d *schema.ResourceData, rc api.ReplicationController) {
	if len(d.Get("spec").([]interface{})) > 0 {
		specList := d.Get("spec").([]interface{})
		ReplicationControllerSpec(specList[0].(map[string]interface{}), rc.Spec)
		d.Set("spec", specList)
	}

	if len(d.Get("metadata").([]interface{})) > 0 {
		metadataList := d.Get("metadata").([]interface{})
		Metadata(metadataList[0].(map[string]interface{}), rc.ObjectMeta)
		d.Set("metadata", metadataList)
	}
}

func ReplicationControllerSpec(d map[string]interface{}, rcs api.ReplicationControllerSpec) {
	if rcs.Template != nil {
		if len(d["template"].([]interface{})) > 0 {
			PodTemplateSpec(d["template"].([]interface{})[0].(map[string]interface{}), *rcs.Template)
		}
	}
}
