package main

import (
	"log"

	"k8s.io/kubernetes/pkg/api/errors"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
)

func resourceKubernetesServiceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEDUG] resourceKubernetesServiceCreate")

	kubeClient := meta.(*KubeProviderClient)

	srvc := buildService(d, kubeClient.Version)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)
	srvcCreate, err := kubeServices.Create(srvc)

	if err != nil {
		return err
	}

	d.SetId(srvcCreate.ObjectMeta.Name)
	return resourceKubernetesServiceRead(d, meta)
}

func resourceKubernetesServiceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceRead")

	kubeClient := meta.(*KubeProviderClient)
	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	_, err := kubeServices.Get(d.Id())

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

	//TODO: populate resourcedata with service response

	return nil
}

func resourceKubernetesServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceUpdate")

	kubeClient := meta.(*KubeProviderClient)

	srvc := buildService(d, kubeClient.Version)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)
	// This will most likely need deleted and recreated
	srvcCreate, err := kubeServices.Update(srvc)

	if err != nil {
		return err
	}

	d.SetId(srvcCreate.ObjectMeta.Name)
	return resourceKubernetesServiceRead(d, meta)
}

func resourceKubernetesServiceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceDelete")

	kubeClient := meta.(*KubeProviderClient)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	return kubeServices.Delete(d.Id())
}
