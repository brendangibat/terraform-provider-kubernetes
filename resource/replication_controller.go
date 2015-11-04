package resource

import (
	"github.com/brendangibat/terraform-provider-kubernetes/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func ReplicationController() *schema.Resource {
	return &schema.Resource{
		Create: operations.ReplicationControllerCreate,
		Read:   operations.ReplicationControllerRead,
		Update: operations.ReplicationControllerUpdate,
		Delete: operations.ReplicationControllerDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ReplicationControllerSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Metadata(),
			},
		},
	}
}

func ReplicationControllerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"replicas": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},

			"template": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     PodTemplateSpec(),
			},

			"selector": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
			},
		},
	}
}
