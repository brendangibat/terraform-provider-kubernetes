package resource

import (
	"github.com/brendangibat/terraform-provider-kubernetes/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func Pod() *schema.Resource {
	return &schema.Resource{
		Create: operations.PodCreate,
		Read:   operations.PodRead,
		Update: operations.PodUpdate,
		Delete: operations.PodDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     PodSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Metadata(),
			},
		},
	}
}

func PodTemplateSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Metadata(),
			},

			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     PodSpec(),
			},
		},
	}
}

func PodSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"volumes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     Volume(),
			},

			"containers": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     Container(),
			},

			"node_selector": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"service_account_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"host_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
			},

			"termination_grace_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"active_deadline_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"restart_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"image_pull_secrets": &schema.Schema{
				Type:     	schema.TypeList,
				Optional: 	true,
				Elem:		LocalObjectReference(),
			},
		},
	}
}
