package resource

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func shallowResourceSchemaMerge(to *schema.Resource, from *schema.Resource) *schema.Resource {
	for key, value := range from.Schema {
		to.Schema[key] = value
	}
	return to
}

func Metadata() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     	schema.TypeString,
				Required: 	true,
				ForceNew: 	true,
			},

			"labels": &schema.Schema{
				Type:     	schema.TypeMap,
				Optional: 	true,
			},

			"annotations": &schema.Schema{
				Type:		schema.TypeMap,
				Optional:	true,
			},

			"uid": &schema.Schema{
				Type:     	schema.TypeString,
				Computed: 	true,
			},

			"namespace": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
				ForceNew: 	true,
			},

			"resource_version": &schema.Schema{
				Type:     	schema.TypeString,
				Computed:	true,
			},

			"creation_timestamp": &schema.Schema{
				Type:     	schema.TypeString,
				Computed: 	true,
			},
		},
	}
}

func LocalObjectReference() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func Probe() *schema.Resource {
	return shallowResourceSchemaMerge(&schema.Resource{
		Schema: map[string]*schema.Schema{
			"initial_delay_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},

			"timeout_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}, Handler())
}

func Handler() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"exec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ExecAction(),
			},

			"http_get": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     HTTPGetAction(),
			},

			"tcp_socket": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     TCPSocketAction(),
			},
		},
	}
}

func ExecAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
		},
	}
}

func HTTPGetAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"host": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"uri_scheme": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func TCPSocketAction() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func ResourceRequirements() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"limits": &schema.Schema{
				Type:     	schema.TypeMap,
				Optional: 	true,
				Elem:		Quantity(),
			},

			"requests": &schema.Schema{
				Type:     	schema.TypeMap,
				Optional: 	true,
				Elem:		Quantity(),
			},
		},
	}
}

func Quantity() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"unscaled": &schema.Schema{
				Type:     	schema.TypeInt,
				Required: 	true,
			},

			"scaled": &schema.Schema{
				Type:     	schema.TypeInt,
				Required: 	true,
			},
		},
	}
}
