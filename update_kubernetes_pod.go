package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func updatePod(d *schema.ResourceData, pod api.Pod) {
	if len(d.Get("metadata").([]interface{})) > 0 {
		updateMetadata(d.Get("metadata").([]interface{})[0].(map[string]interface{}), pod.ObjectMeta)
	}
}

func updatePodTemplateSpec(d map[string]interface{}, pts api.PodTemplateSpec) {
	if len(d["metadata"].([]interface{})) > 0 {
		updateMetadata(d["metadata"].([]interface{})[0].(map[string]interface{}), pts.ObjectMeta)
	}
}
