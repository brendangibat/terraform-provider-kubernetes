package update

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func Service(d *schema.ResourceData, service api.Service) {
	if len(d.Get("spec").([]interface{})) > 0 {
		specList := d.Get("spec").([]interface{})
		ServiceSpec(specList[0].(map[string]interface{}), service.Spec)
		d.Set("spec", specList)
	}
	if len(d.Get("metadata").([]interface{})) > 0 {
		metadataList := d.Get("metadata").([]interface{})
		Metadata(metadataList[0].(map[string]interface{}), service.ObjectMeta)
		d.Set("metadata", metadataList)
	}
}

func ServiceSpec(d map[string]interface{}, serviceSpec api.ServiceSpec) {
    d["cluster_ip"] = serviceSpec.ClusterIP
}
