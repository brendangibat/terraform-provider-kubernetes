package update

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func Namespace(d *schema.ResourceData, namespace api.Namespace) {
	if len(d.Get("metadata").([]interface{})) > 0 {
		metadataList := d.Get("metadata").([]interface{})
		Metadata(metadataList[0].(map[string]interface{}), namespace.ObjectMeta)
		d.Set("metadata", metadataList)
	}
}
