package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceKubernetesReplicationController() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesReplicationControllerCreate,
		Read:   resourceKubernetesReplicationControllerRead,
		Update: resourceKubernetesReplicationControllerUpdate,
		Delete: resourceKubernetesReplicationControllerDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     resourceUnitReplicationControllerSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     resourceUnitMetadata(),
			},
		},
	}
}

func resourceUnitReplicationControllerSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"replicas": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"template": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     resourceUnitPodTemplateSpec(),
			},

			"selector": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
