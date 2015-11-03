package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceKubernetesService() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesServiceCreate,
		Read:   resourceKubernetesServiceRead,
		Update: resourceKubernetesServiceUpdate,
		Delete: resourceKubernetesServiceDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitServiceSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitMetadata(),
			},
		},
	}
}

func resourceUnitServiceSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ports": &schema.Schema{
				Type:     	schema.TypeList,
				Required: 	true,
				Elem:     	resourceUnitServicePort(),
			},
			"selector": &schema.Schema{
				Type:     	schema.TypeMap,
				Optional: 	true,
			},
			"cluster_ip": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
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

func resourceUnitServicePort() *schema.Resource {
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
