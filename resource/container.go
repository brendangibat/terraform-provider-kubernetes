package resource

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func Container() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"args": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},

			"command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},

			"working_dir": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     ContainerPort(),
			},

			"env": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     EnvVar(),
			},

			"volume_mounts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem:     VolumeMount(),
			},

			"termination_message_path": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
				Default:	api.TerminationMessagePathDefault,
			},

			// "Always", "Never", "IfNotPresent" are valid values
			"image_pull_policy": &schema.Schema{
				Type:     	schema.TypeString,
				Optional: 	true,
				Default:	"Always",
			},

			// TODO: use these when we're ready to populate them
			//
			// "resources": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	Elem:     ResourceRequirements(),
			// },
			//
			// "security_context": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	Elem:     SecurityContext(),
			// },
			// "liveness_probe": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	Elem:     Probe(),
			// },
			//
			// "readiness_probe": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	Elem:     Probe(),
			// },
			//
			// "lifecycle": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Optional: true,
			// 	ForceNew: true,
			// 	Elem:     Lifecycle(),
			// },
		},
	}
}

func ContainerPort() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"host_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},

			"container_port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},

			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"host_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func EnvVar() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func Lifecycle() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"post_start": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Handler(),
			},

			"pre_stop": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Handler(),
			},
		},
	}
}

func SecurityContext() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"capabilities": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     Capability(),
			},

			"privileged": &schema.Schema{
				Type:     	schema.TypeBool,
				Optional: 	true,
				Default:	false,
			},

			"se_linux_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     SELinuxOptions(),
			},

			"run_as_user": &schema.Schema{
				Type:     	schema.TypeInt,
				Optional: 	true,
			},
		},
	}
}

func Capability() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"add": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},

			"drop": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema {
					Type: schema.TypeString,
				},
			},
		},
	}
}

func SELinuxOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}
