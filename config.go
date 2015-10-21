package main

import (
	"errors"
	"log"
	"k8s.io/kubernetes/pkg/client/unversioned"
)

const DefaultVersion = "v1"
const DefaultNamespace = "default"

type (
	Config struct {
		Endpoint     string `mapstructure:"endpoint"`
		User     	 string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		Namespace	 string	`mapstructure:"namespace"`
		Version	     string	`mapstructure:"version"`
	}

	KubeProviderClient struct {
		KubeClient  *client.Client
		Namespace	string
		Version		string
	}
)

// Client() returns a new client for accessing kubernetes.
func (c *Config) Client() (*KubeProviderClient, error) {
	if c.Endpoint == "" {
		return nil, errors.New("endpoint is required")
	}
	kubeProClient := &KubeProviderClient{}
	kubeConfig := &client.Config{
		Host:    c.Endpoint,
		Version: "v1",
	}

	if len(c.User) > 0 {
		kubeConfig.Username = c.User
	}

	if len(c.Password) > 0 {
		kubeConfig.Password = c.Password
	}

	if len(c.Version) > 0 {
		kubeConfig.Version = c.Version
	} else {
		kubeConfig.Version = DefaultVersion
	}

	kubeProClient.Version = kubeConfig.Version

	if len(c.Namespace) > 0 {
		kubeProClient.Namespace = c.Namespace
	} else {
		kubeProClient.Namespace = DefaultNamespace
	}

	kubeProClient.KubeClient = client.NewOrDie(kubeConfig)

	log.Printf("[INFO] Kubernetes Client configured with endpoint: '%s'", c.Endpoint)

	return kubeProClient, nil
}
