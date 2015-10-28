package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUnitMetadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},

			"uid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"resource_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"creation_timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
