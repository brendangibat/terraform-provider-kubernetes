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
				ForceNew: true,
				Elem:     resourceUnitMetadata(),
			},

			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},

			"node_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
				ForceNew: true,
			},

			"active_deadline_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			"restart_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
