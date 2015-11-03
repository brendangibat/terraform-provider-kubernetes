package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func resourceKubernetesNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesNamespaceCreate")

	kubeClient := meta.(*KubeProviderClient)

	namespace := buildNamespace(d, kubeClient.Version)

	kubeNamspaces := kubeClient.KubeClient.Namespaces()
	namespaceCreate, err := kubeNamspaces.Create(namespace)

	if err != nil {
		return err
	}

	d.SetId(namespaceCreate.ObjectMeta.Name)
	return resourceKubernetesNamespaceRead(d, meta)
}

func resourceKubernetesNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesNamespaceRead")

	kubeClient := meta.(*KubeProviderClient)
	kubeNamespaces := kubeClient.KubeClient.Namespaces()

	namespace, err := kubeNamespaces.Get(d.Id())

	if err != nil {
		switch err.(type) {
		case *errors.StatusError:
			if err.(*errors.StatusError).ErrStatus.Reason == api.StatusReasonNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	updateNamespace(d, *namespace)

	return nil
}

func resourceKubernetesNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesNamespaceUpdate")

	kubeClient := meta.(*KubeProviderClient)

	namespace := buildNamespace(d, kubeClient.Version)

	kubeNamespaces := kubeClient.KubeClient.Namespaces()
	// This will most likely need deleted and recreated
	updatedNamespace, err := kubeNamespaces.Update(namespace)

	if err != nil {
		return err
	}

	d.SetId(updatedNamespace.ObjectMeta.Name)
	return resourceKubernetesNamespaceRead(d, meta)
}

func resourceKubernetesNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesNamespaceDelete")

	kubeClient := meta.(*KubeProviderClient)

	kubeNamespaces := kubeClient.KubeClient.Namespaces()

	return kubeNamespaces.Delete(d.Id())
}
