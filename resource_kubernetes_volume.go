package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUnitVolume() *schema.Resource {
	return shallowResourceSchemaMerge(&schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}, resourceUnitVolumeSource())
}

func resourceUnitVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"host_path": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitHostPathVolumeSource(),
			},

			"empty_dir": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitEmptyDirVolumeSource(),
			},

			"gce_persistent_disk": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitGCEPersistentDiskVolumeSource(),
			},

			"aws_elastic_block_store": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitAWSElasticBlockStoreVolumeSource(),
			},

			"git_repo": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitGitRepoVolumeSource(),
			},

			"secret": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitSecretVolumeSource(),
			},

			"nfs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitNFSVolumeSource(),
			},

			"iscsi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitISCSIVolumeSource(),
			},

			"glusterfs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitGlusterFSVolumeSource(),
			},

			"persistent_volume_claim": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitPersistentVolumeClaimVolumeSource(),
			},

			"rbd": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     resourceUnitRBDVolumeSource(),
			},
		},
	}
}

func resourceUnitHostPathVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUnitEmptyDirVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"medium": &schema.Schema{
				Type:     	schema.TypeString,
				Required:	true,
			},
		},
	}
}

func resourceUnitGCEPersistentDiskVolumeSource() *schema.Resource {
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

func resourceUnitAWSElasticBlockStoreVolumeSource() *schema.Resource {
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

func resourceUnitGitRepoVolumeSource() *schema.Resource {
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

func resourceUnitSecretVolumeSource() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"secret_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceUnitNFSVolumeSource() *schema.Resource {
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

func resourceUnitISCSIVolumeSource() *schema.Resource {
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

func resourceUnitGlusterFSVolumeSource() *schema.Resource {
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

func resourceUnitRBDVolumeSource() *schema.Resource {
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
				Elem:		resourceUnitLocalObjectReference(),
			},

			"read_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: false,
			},
		},
	}
}

func resourceUnitPersistentVolumeClaimVolumeSource() *schema.Resource {
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

func resourceUnitVolumeMount() *schema.Resource {
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
