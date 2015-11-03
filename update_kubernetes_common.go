package main

import (
	"k8s.io/kubernetes/pkg/api"
)

func updateMetadata(d map[string]interface{}, objectMeta api.ObjectMeta) {
    d["uid"] = string(objectMeta.UID)
	d["resource_version"] = objectMeta.ResourceVersion
	d["creation_timestamp"] = objectMeta.CreationTimestamp.String()
}
