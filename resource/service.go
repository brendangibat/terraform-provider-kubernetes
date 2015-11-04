package resource

import (
	"github.com/brendangibat/terraform-provider-kubernetes/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func Service() *schema.Resource {
	return &schema.Resource{
		Create: operations.ServiceCreate,
		Read:   operations.ServiceRead,
		Update: operations.ServiceUpdate,
		Delete: operations.ServiceDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ServiceSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Metadata(),
			},
		},
	}
}

func ServiceSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ports": &schema.Schema{
				Type:     	schema.TypeList,
				Required: 	true,
				Elem:     	ServicePort(),
			},
			"selector": &schema.Schema{
				Type:     	schema.TypeMap,
				Optional: 	true,
			},
			"cluster_ip": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
				Computed:	true,
			},
			"type": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
			},
			"session_affinity": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
				Default:	"None",
			},
		},
	}
}

func ServicePort() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"node_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
