package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func buildVolumes(userVolumes []interface{}) []api.Volume {
	volumes := make([]api.Volume, len(userVolumes))
	if len(userVolumes) == 0 {
		return volumes
	}

	for index, userVolume := range userVolumes {
		volumes[index] = buildVolume(userVolume.(map[string]interface{}))
	}

	return volumes
}

func buildVolume(userVolume map[string]interface{}) api.Volume {
	volume := api.Volume{
		Name:	userVolume["name"].(string),
	}
	volume.VolumeSource.populateVolumeSource(userVolume)
}

func (volume api.VolumeSource) populateVolumeSource(userVolumeSource map[string]interface{}) {

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
