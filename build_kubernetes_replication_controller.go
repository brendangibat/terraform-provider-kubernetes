package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func buildReplicationController(d *schema.ResourceData, version string) *api.ReplicationController {

	rc := &api.ReplicationController{
		Spec: buildReplicationControllerSpec(d.Get("spec").([]interface{})),
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
		Replicas: spec["replicas"].(int),
		Selector: convertMapTypeToStringMap(spec["selector"].(map[string]interface{})),
		Template: buildPodTemplateSpec(spec["template"].([]interface{})),
	}
}

func buildPodTemplateSpec(templates []interface{}) *api.PodTemplateSpec {
	if len(templates) == 0 {
		return nil
	}
	template := templates[0].(map[string]interface{})

	pts := &api.PodTemplateSpec{
		Spec: buildPodSpec(template["spec"].([]interface{})),
	}

	populateMetadata(&pts.ObjectMeta, template["metadata"].([]interface{}))

	return pts
}

func buildContainers(userContainers []interface{}) []api.Container {
	if len(userContainers) == 0 {
		return nil
	}

	var containers []api.Container

	for _, c := range userContainers {
		userContainer := c.(map[string]interface{})
		container := api.Container{
			Image: userContainer["image"].(string),
			Name:  userContainer["name"].(string),
		}

		if _, ok := userContainer["args"]; ok {
			container.Args = convertListToStringArray(userContainer["args"].([]interface{}))
		}
		if _, ok := userContainer["command"]; ok {
			container.Command = convertListToStringArray(userContainer["command"].([]interface{}))
		}
		if _, ok := userContainer["working_dir"]; ok {
			container.WorkingDir = userContainer["working_dir"].(string)
		}
		if _, ok := userContainer["ports"]; ok {
			container.Ports = buildContainerPorts(userContainer["ports"].([]interface{}))
		}
		if _, ok := userContainer["env"]; ok {
			container.Env = buildEnvVar(userContainer["env"].([]interface{}))
		}

		if _, ok := userContainer["volume_mounts"]; ok {
			container.VolumeMounts = buildVolumeMounts(userContainer["volume_mounts"].([]interface{}))
		}

		if _, ok := userContainer["termination_message_path"]; ok {
			container.TerminationMessagePath = userContainer["termination_message_path"].(string)
		}

		if _, ok := userContainer["image_pull_policy"]; ok {
			container.ImagePullPolicy = api.PullPolicy(userContainer["image_pull_policy"].(string))
		}

// TODO: populate these fields:
// resources
// liveness_probe
// readiness_probe
// lifecycle
// security_context

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
			ContainerPort: userPort["container_port"].(int),
		}

		if _, ok := userPort["host_port"]; ok {
			port.HostPort = userPort["host_port"].(int)
		}
		if _, ok := userPort["name"]; ok {
			port.Name = userPort["name"].(string)
		}
		if _, ok := userPort["protocol"]; ok {
			port.Protocol = api.Protocol(userPort["protocol"].(string))
		}
		if _, ok := userPort["host_ip"]; ok {
			port.HostIP = userPort["host_ip"].(string)
		}

		ports = append(ports, port)
	}
	return ports
}
