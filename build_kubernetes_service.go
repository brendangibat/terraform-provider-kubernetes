package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func buildService(d *schema.ResourceData, version string) *api.Service {
	srvc := &api.Service{
		Spec: buildServiceSpec(d.Get("spec").([]interface{})),
	}

	srvc.Kind = "Service"
	srvc.APIVersion = version

	populateMetadata(&srvc.ObjectMeta, d.Get("metadata").([]interface{}))

	return srvc
}

func buildServiceSpec(srvcSpecs []interface{}) api.ServiceSpec {
	if len(srvcSpecs) == 0 {
		return api.ServiceSpec{}
	}

	userServiceSpec := srvcSpecs[0].(map[string]interface{})

	srvcSpec := api.ServiceSpec{
		Ports: buildServicePorts(userServiceSpec["ports"].([]interface{})),
	}

	if _, ok := userServiceSpec["selector"]; ok {
		srvcSpec.Selector = convertMapTypeToStringMap(userServiceSpec["selector"].(map[string]interface{}))
	}

	if _, ok := userServiceSpec["cluster_ip"]; ok {
		srvcSpec.ClusterIP = userServiceSpec["cluster_ip"].(string)
	}

	if _, ok := userServiceSpec["type"]; ok {
		srvcSpec.Type = api.ServiceType(userServiceSpec["type"].(string))
	}

	if _, ok := userServiceSpec["session_affinity"]; ok {
		srvcSpec.SessionAffinity = api.ServiceAffinity(userServiceSpec["session_affinity"].(string))
	}

	return srvcSpec
}

func buildServicePorts(srvcPorts []interface{}) []api.ServicePort {
	if len(srvcPorts) == 0 {
		return nil
	}

	var ports []api.ServicePort

	for _, p := range srvcPorts {
		userPort := p.(map[string]interface{})
		srvcPort := api.ServicePort{}

		if _, ok := userPort["name"]; ok {
			srvcPort.Name = userPort["name"].(string)
		}

		if _, ok := userPort["protocol"]; ok {
			srvcPort.Protocol = api.Protocol(userPort["protocol"].(string))
		}

		if _, ok := userPort["port"]; ok {
			srvcPort.Port = userPort["port"].(int)
		}

		if _, ok := userPort["node_port"]; ok {
			srvcPort.NodePort = userPort["node_port"].(int)
		}

		ports = append(ports, srvcPort)
	}

	return ports
}
