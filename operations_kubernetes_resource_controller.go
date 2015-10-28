package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func resourceKubernetesReplicationControllerCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerCreate")

	kubeClient := meta.(*KubeProviderClient)

	rc := buildReplicationController(d, kubeClient.Version)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)
	rcCreate,errCreate := kubeRepControllers.Create(rc)

	if errCreate != nil {
		return errCreate
	}

	d.SetId(rcCreate.ObjectMeta.Name)
	return resourceKubernetesReplicationControllerRead(d, meta)
}

func resourceKubernetesReplicationControllerRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerRead")

	kubeClient := meta.(*KubeProviderClient)
	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	_,err := kubeRepControllers.Get(d.Id())

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

	//TODO: populate resourcedata with rc response

	return nil
}

func resourceKubernetesReplicationControllerUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerUpdate")

	kubeClient := meta.(*KubeProviderClient)

	rc := buildReplicationController(d, kubeClient.Version)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)
	updatedRepController,updateErr := kubeRepControllers.Update(rc)

	if updateErr != nil {
		return updateErr
	}

	d.SetId(updatedRepController.ObjectMeta.Name)
	return resourceKubernetesReplicationControllerRead(d, meta)
}

func resourceKubernetesReplicationControllerDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerDelete")

	kubeClient := meta.(*KubeProviderClient)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	return kubeRepControllers.Delete(d.Id())
}
