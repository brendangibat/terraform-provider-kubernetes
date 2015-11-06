package build

import (
	"log"

	"k8s.io/kubernetes/pkg/api"
)

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
	if _, ok := metadata["labels"]; ok {
		obj.Labels = convertMapTypeToStringMap(metadata["labels"].(map[string]interface{}))
	}
	if _, ok := metadata["annotations"]; ok {
		obj.Annotations = convertMapTypeToStringMap(metadata["annotations"].(map[string]interface{}))
	}
}

func EnvVar(userEnvVars []interface{}) []api.EnvVar {
	if len(userEnvVars) == 0 {
		return nil
	}

	var envVars []api.EnvVar

	for _, e := range userEnvVars {
		userEnvVar := e.(map[string]interface{})

		envVar := api.EnvVar{
			Name: userEnvVar["name"].(string),
		}

		if _, ok := userEnvVar["value"]; ok {
			log.Printf("envvar value : %s", userEnvVar["value"].(string))
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
	for po, val := range list {
		ret[po] = val.(string)
	}
	return ret
}

func convertMapTypeToStringMap(userConfig map[string]interface{}) map[string]string {
	config := make(map[string]string)
	for k, v := range userConfig {
		config[k] = v.(string)
	}
	return config
}

func convertNameValueListToStringMap(userConfig []interface{}) map[string]string {
	config := make(map[string]string)
	for _, userNameValuePair := range userConfig {
		nameValue := userNameValuePair.(map[string]interface{})
		var value string
		if _, ok := nameValue["value"]; ok {
			value = nameValue["value"].(string)
		}
		config[nameValue["name"].(string)] = value
	}
	return config
}
