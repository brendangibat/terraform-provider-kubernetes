package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func buildPod(d *schema.ResourceData, version string) *api.Pod {
	pod := &api.Pod{
		Spec: buildPodSpec(d.Get("spec").([]interface{})),
	}

	pod.Kind = "Pod"
	pod.APIVersion = version

	populateMetadata(&pod.ObjectMeta, d.Get("metadata").([]interface{}))

	return pod
}

func buildPodSpec(podSpecs []interface{}) api.PodSpec {
	if len(podSpecs) == 0 {
		return api.PodSpec{}
	}
	userPodSpec := podSpecs[0].(map[string]interface{})

	podSpec := api.PodSpec{
		Containers:                    buildContainers(userPodSpec["containers"].([]interface{})),
		ActiveDeadlineSeconds:         nil,
		TerminationGracePeriodSeconds: nil,
	}

	if _, ok := userPodSpec["node_selector"]; ok {
		podSpec.NodeSelector = convertMapTypeToStringMap(userPodSpec["node_selector"].(map[string]interface{}))
	}
	if _, ok := userPodSpec["node_name"]; ok {
		podSpec.NodeName = userPodSpec["node_name"].(string)
	}
	if _, ok := userPodSpec["service_account_name"]; ok {
		podSpec.ServiceAccountName = userPodSpec["service_account_name"].(string)
	}
	if _, ok := userPodSpec["host_network"]; ok {
		podSpec.HostNetwork = userPodSpec["host_network"].(bool)
	}
	if _, ok := userPodSpec["termination_grace_period"]; ok {
		helper := int64(userPodSpec["termination_grace_period"].(int))
		podSpec.TerminationGracePeriodSeconds = &helper
		if helper > 0 {
			podSpec.TerminationGracePeriodSeconds = &helper
		} else {
			podSpec.TerminationGracePeriodSeconds = nil
		}
	} else {
		podSpec.TerminationGracePeriodSeconds = nil
	}
	if _, ok := userPodSpec["active_deadline_seconds"]; ok {
		helper := int64(userPodSpec["active_deadline_seconds"].(int))
		if helper > 0 {
			podSpec.ActiveDeadlineSeconds = &helper
		} else {
			podSpec.ActiveDeadlineSeconds = nil
		}
	} else {
		podSpec.ActiveDeadlineSeconds = nil
	}
	if _, ok := userPodSpec["restart_policy"]; ok {
		podSpec.RestartPolicy = api.RestartPolicy(userPodSpec["restart_policy"].(string))
	}

	return podSpec
}
