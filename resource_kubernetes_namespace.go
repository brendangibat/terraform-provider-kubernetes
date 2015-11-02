package main

import "github.com/hashicorp/terraform/helper/schema"

func resourceKubernetesNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesNamespaceCreate,
		Read:   resourceKubernetesNamespaceRead,
		Update: resourceKubernetesNamespaceUpdate,
		Delete: resourceKubernetesNamespaceDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitNamespaceSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitMetadata(),
			},
		},
	}
}

func resourceUnitNamespaceSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"finalizers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}
