package update

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func Pod(d *schema.ResourceData, pod api.Pod) {
	if len(d.Get("metadata").([]interface{})) > 0 {
		Metadata(d.Get("metadata").([]interface{})[0].(map[string]interface{}), pod.ObjectMeta)
	}
}

func PodTemplateSpec(d map[string]interface{}, pts api.PodTemplateSpec) {
	if len(d["metadata"].([]interface{})) > 0 {
		Metadata(d["metadata"].([]interface{})[0].(map[string]interface{}), pts.ObjectMeta)
	}
}
