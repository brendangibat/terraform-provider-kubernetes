package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updatePod(d *schema.ResourceData, pod api.Pod) {
	if d.Get("metadata.#").(int) > 0 {
    	updateMetadata(d.Get("metadata.0").(*schema.ResourceData), pod.ObjectMeta)
	}
}

func updatePodTemplateSpec(d *schema.ResourceData, pts api.PodTemplateSpec) {
	if d.Get("metadata.#").(int) > 0 {
    	updateMetadata(d.Get("metadata.0").(*schema.ResourceData), pts.ObjectMeta)
	}
}
