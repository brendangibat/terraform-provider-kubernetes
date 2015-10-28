package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUnitPodTemplateSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitMetadata(),
			},

			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitPodSpec(),
			},
		},
	}
}

func resourceUnitPodSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"containers": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem:     resourceUnitContainer(),
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
				Default: false,
				ForceNew: true,
			},

			"termination_grace_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default: 1,
			},

			"active_deadline_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default: 1,
			},

			"restart_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
