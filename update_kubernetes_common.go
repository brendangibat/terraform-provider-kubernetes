package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateMetadata(d *schema.ResourceData, objectMeta api.ObjectMeta) {
    d.Set("uid", string(objectMeta.UID))
	d.Set("resource_version", objectMeta.ResourceVersion)
	d.Set("creation_timestamp", objectMeta.CreationTimestamp.String())
}
