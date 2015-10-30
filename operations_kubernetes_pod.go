package main

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func resourceKubernetesPodCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodCreate")

	kubeClient := meta.(*KubeProviderClient)

	pod := buildPod(d, kubeClient.Version)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)
	podCreate, errCreate := kubePods.Create(pod)

	if errCreate != nil {
		return errCreate
	}

	d.SetId(podCreate.ObjectMeta.Name)
	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodRead")

	kubeClient := meta.(*KubeProviderClient)
	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	_, err := kubePods.Get(d.Id())

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

	//TODO: populate resourcedata with pod response

	return nil
}

func resourceKubernetesPodUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodUpdate")

	kubeClient := meta.(*KubeProviderClient)

	pod := buildPod(d, kubeClient.Version)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)
	// This might have to be killed and recreated.
	updatedPod, updateErr := kubePods.Update(pod)

	if updateErr != nil {
		return updateErr
	}

	d.SetId(updatedPod.ObjectMeta.Name)
	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodDelete")

	kubeClient := meta.(*KubeProviderClient)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	return kubePods.Delete(d.Id(), nil)
}
