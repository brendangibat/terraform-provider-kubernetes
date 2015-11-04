package operations

import (
	"log"

	"github.com/brendangibat/terraform-provider-kubernetes/config"
	"github.com/brendangibat/terraform-provider-kubernetes/update"
	"github.com/brendangibat/terraform-provider-kubernetes/build"
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func PodCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] PodCreate")

	kubeClient := meta.(*config.KubeProviderClient)

	pod := build.Pod(d, kubeClient.Version)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)
	podCreate, errCreate := kubePods.Create(pod)

	if errCreate != nil {
		return errCreate
	}

	d.SetId(podCreate.ObjectMeta.Name)
	return PodRead(d, meta)
}

func PodRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] PodRead")

	kubeClient := meta.(*config.KubeProviderClient)
	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	pod, err := kubePods.Get(d.Id())

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

	update.Pod(d, *pod)

	return nil
}

func PodUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] PodUpdate")

	kubeClient := meta.(*config.KubeProviderClient)

	pod := build.Pod(d, kubeClient.Version)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)
	// This might have to be killed and recreated.
	updatedPod, updateErr := kubePods.Update(pod)

	if updateErr != nil {
		return updateErr
	}

	d.SetId(updatedPod.ObjectMeta.Name)
	return PodRead(d, meta)
}

func PodDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] PodDelete")

	kubeClient := meta.(*config.KubeProviderClient)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	return kubePods.Delete(d.Id(), nil)
}
