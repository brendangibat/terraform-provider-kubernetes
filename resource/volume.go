package resource

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Volume() *schema.Resource {
	return shallowResourceSchemaMerge(&schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}, VolumeSource())
}

func VolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_path": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     HostPathVolumeSource(),
			},

			"empty_dir": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     EmptyDirVolumeSource(),
			},

			"gce_persistent_disk": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     GCEPersistentDiskVolumeSource(),
			},

			"aws_elastic_block_store": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     AWSElasticBlockStoreVolumeSource(),
			},

			"git_repo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     GitRepoVolumeSource(),
			},

			"secret": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     SecretVolumeSource(),
			},

			"persistent_volume_claim": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     PersistentVolumeClaimVolumeSource(),
			},
		},
	}
}

func HostPathVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func EmptyDirVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"medium": &schema.Schema{
				Type:     	schema.TypeString,
				Required:	true,
			},
		},
	}
}

func GCEPersistentDiskVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pd_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"fs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"partition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default: false,
			},
		},
	}
}

func AWSElasticBlockStoreVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"fs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"partition": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default: false,
			},
		},
	}
}

func GitRepoVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"repository": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"revision": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func SecretVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
/*
func NFSVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},
		},
	}
}

func ISCSIVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"target_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"iqn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"lun": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"fs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},
		},
	}
}

func GlusterFSVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"endpoints": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},
		},
	}
}

func RBDVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ceph_monitors": &schema.Schema{
				Type:		schema.TypeList,
				Required:	true,
				Elem: 		&schema.Schema {
					Type: schema.TypeString,
				},
			},

			"rbd_image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"fs_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"rbd_pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Default: "rbd",
			},

			"rados_user": &schema.Schema{
				Type:    	schema.TypeString,
				Required:	true,
				Default:	"admin",
			},

			"keyring": &schema.Schema{
				Type:    	schema.TypeString,
				Optional:	true,
				Default:	"/etc/ceph/keyring",
			},

			"secret_ref": &schema.Schema{
				Type:    	schema.TypeString,
				Required:	true,
				Elem:		LocalObjectReference(),
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},
		},
	}
}
*/

func PersistentVolumeClaimVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"claim_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},
		},
	}
}

func VolumeMount() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},

			"mount_path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
