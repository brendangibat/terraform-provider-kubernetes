package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updateService(d *schema.ResourceData, service api.Service) {
	if d.Get("spec.#").(int) > 0 {
    	updateServiceSpec(d.Get("spec.0").(*schema.ResourceData), service.Spec)
	}
	if d.Get("metadata.#").(int) > 0 {
    	updateMetadata(d.Get("metadata.0").(*schema.ResourceData), service.ObjectMeta)
	}
}

func updateServiceSpec(d *schema.ResourceData, serviceSpec api.ServiceSpec) {
    d.Set("cluster_ip", serviceSpec.ClusterIP)
}
