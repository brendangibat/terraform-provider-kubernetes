package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateService(d *schema.ResourceData, service api.Service) {
	updateServiceSpec(d.Get("spec").(map[string]interface{}), service.Spec)
	updateMetadata(d.Get("metadata").(map[string]interface{}), service.ObjectMeta)
}

func updateServiceSpec(d map[string]interface{}, serviceSpec api.ServiceSpec) {
    d["cluster_ip"] = serviceSpec.ClusterIP
}
