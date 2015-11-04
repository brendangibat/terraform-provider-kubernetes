package resource

import (
	"github.com/brendangibat/terraform-provider-kubernetes/operations"
	"github.com/hashicorp/terraform/helper/schema"
)

func Namespace() *schema.Resource {
	return &schema.Resource{
		Create: operations.NamespaceCreate,
		Read:   operations.NamespaceRead,
		Update: operations.NamespaceUpdate,
		Delete: operations.NamespaceDelete,
		Schema: map[string]*schema.Schema{
			"spec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     NamespaceSpec(),
			},
			"metadata": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Metadata(),
			},
		},
	}
}

func NamespaceSpec() *schema.Resource {
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
