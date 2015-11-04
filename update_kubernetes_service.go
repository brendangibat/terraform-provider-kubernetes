package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateService(d *schema.ResourceData, service api.Service) {
	if len(d.Get("spec").([]interface{})) > 0 {
		updateServiceSpec(d.Get("spec").([]interface{})[0].(map[string]interface{}), service.Spec)
	}
	if len(d.Get("metadata").([]interface{})) > 0 {
		updateMetadata(d.Get("metadata").([]interface{})[0].(map[string]interface{}), service.ObjectMeta)
	}
}

func updateServiceSpec(d map[string]interface{}, serviceSpec api.ServiceSpec) {
    d["cluster_ip"] = serviceSpec.ClusterIP
}
