package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func buildReplicationController(d *schema.ResourceData, version string) *api.ReplicationController {

	rc := &api.ReplicationController{
		Spec : buildReplicationControllerSpec(d.Get("spec").([]interface{})),
	}

	rc.Kind = "ReplicationController"
	rc.APIVersion = version

	populateMetadata(&rc.ObjectMeta, d.Get("metadata").([]interface{}))

	return rc
}

func buildReplicationControllerSpec(specs []interface{}) api.ReplicationControllerSpec {
	if len(specs) == 0 {
		return api.ReplicationControllerSpec{}
	}
	spec := specs[0].(map[string]interface{})

	return api.ReplicationControllerSpec{
		Replicas : spec["replicas"].(int),
		Selector : convertMapTypeToStringMap(spec["selector"].(map[string]interface{})),
		Template : buildPodTemplateSpec(spec["template"].([]interface{})),
	}
}

func populateMetadata(obj *api.ObjectMeta, metadatas []interface{}) {
	if len(metadatas) == 0 {
		return
	}
	metadata := metadatas[0].(map[string]interface{})

	if _, ok := metadata["name"]; ok {
		obj.Name = metadata["name"].(string)
	}
	if _, ok := metadata["namespace"]; ok {
		obj.Namespace = metadata["namespace"].(string)
	}
	if _,ok := metadata["resource_version"]; ok {
		obj.ResourceVersion = metadata["resource_version"].(string)
	}
	if _,ok := metadata["labels"]; ok {
		obj.Labels = convertMapTypeToStringMap(metadata["labels"].(map[string]interface{}))
	}
}

func buildPodTemplateSpec(templates []interface{}) *api.PodTemplateSpec {
	if len(templates) == 0 {
		return nil
	}
	template := templates[0].(map[string]interface{})

	pts := &api.PodTemplateSpec{
		Spec : buildPodSpec(template["spec"].([]interface{})),
	}

	populateMetadata(&pts.ObjectMeta, template["metadata"].([]interface{}))

	return pts
}


func buildPodSpec(podSpecs []interface{}) api.PodSpec {
	if len(podSpecs) == 0 {
		return api.PodSpec{}
	}
	userPodSpec := podSpecs[0].(map[string]interface{})

	podSpec := api.PodSpec{
		Containers : buildContainers(userPodSpec["containers"].([]interface{})),
	}

	if _,ok := userPodSpec["node_selector"]; ok {
		podSpec.NodeSelector = convertMapTypeToStringMap(userPodSpec["node_selector"].(map[string]interface{}))
	}
	if _,ok := userPodSpec["node_name"]; ok {
		podSpec.NodeName = userPodSpec["node_name"].(string)
	}
	if _,ok := userPodSpec["service_account_name"]; ok {
		podSpec.ServiceAccountName = userPodSpec["service_account_name"].(string)
	}
	if _,ok := userPodSpec["host_network"]; ok {
		podSpec.HostNetwork = userPodSpec["host_network"].(bool)
	}
	if _,ok := userPodSpec["termination_grace_period"]; ok {
		helper := int64(userPodSpec["termination_grace_period"].(int))
		podSpec.TerminationGracePeriodSeconds = &helper
	}
	if _,ok := userPodSpec["active_deadline_seconds"]; ok {
		helper := int64(userPodSpec["active_deadline_seconds"].(int))
		podSpec.ActiveDeadlineSeconds = &helper
	}
	if _,ok := userPodSpec["restart_policy"]; ok {
		podSpec.RestartPolicy = api.RestartPolicy(userPodSpec["restart_policy"].(string))
	}

	return podSpec
}

func buildContainers(userContainers []interface{}) []api.Container {
	if len(userContainers) == 0 {
		return nil
	}

	var containers []api.Container

	for _,c := range userContainers {
		userContainer := c.(map[string]interface{})
		container := api.Container{
			Image : userContainer["image"].(string),
			Name : userContainer["name"].(string),
		}

		if _,ok := userContainer["args"]; ok {
			container.Args = convertListToStringArray(userContainer["args"].([]interface{}))
		}
		if _,ok := userContainer["command"]; ok {
			container.Command = convertListToStringArray(userContainer["command"].([]interface{}))
		}
		if _,ok := userContainer["working_dir"]; ok {
			container.WorkingDir = userContainer["working_dir"].(string)
		}
		if _,ok := userContainer["ports"]; ok {
			container.Ports = buildContainerPorts(userContainer["ports"].([]interface{}))
		}
		if _,ok := userContainer["env"]; ok {
			container.Env = buildEnvVar(userContainer["env"].([]interface{}))
		}
		containers = append(containers, container)
	}

	return containers
}

func buildContainerPorts(userPorts []interface{}) []api.ContainerPort {
	if len(userPorts) == 0 {
		return nil
	}

	var ports []api.ContainerPort

	for _, p := range userPorts {
		userPort := p.(map[string]interface{})

		port := api.ContainerPort{
			ContainerPort : userPort["container_port"].(int),
		}

		if _,ok := userPort["host_port"]; ok {
			port.HostPort = userPort["host_port"].(int)
		}
		if _,ok := userPort["name"]; ok {
			port.Name = userPort["name"].(string)
		}
		if _,ok := userPort["protocol"]; ok {
			port.Protocol = api.Protocol(userPort["protocol"].(string))
		}
		if _,ok := userPort["host_ip"]; ok {
			port.HostIP = userPort["host_ip"].(string)
		}

		ports = append(ports, port)
	}
	return ports
}

func buildEnvVar(userEnvVars []interface{}) []api.EnvVar {
	if len(userEnvVars) == 0 {
		return nil
	}

	var envVars []api.EnvVar

	for _, e := range userEnvVars {
		userEnvVar := e.(map[string]interface{})

		envVar := api.EnvVar{
			Name : userEnvVar["name"].(string),
		}

		if _,ok := userEnvVar["value"]; ok {
			envVar.Value = userEnvVar["value"].(string)
		}

		envVars = append(envVars, envVar)
	}
	return envVars
}

func convertListToStringArray(list []interface{}) []string {
	if list == nil || len(list) == 0 {
		return nil
	}
	ret := make([]string, len(list))
	for po,val := range list {
		ret[po] = val.(string)
	}
	return ret
}

func convertMapTypeToStringMap(userConfig map[string]interface{}) map[string]string {
	config := make(map[string]string)
	for k,v := range userConfig {
		config[k] = v.(string)
	}
	return config
}
