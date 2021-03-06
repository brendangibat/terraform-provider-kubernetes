package operations

import (
	"log"

	"github.com/brendangibat/terraform-provider-kubernetes/config"
	"github.com/brendangibat/terraform-provider-kubernetes/update"
	"github.com/brendangibat/terraform-provider-kubernetes/build"
	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/fields"
)

func ReplicationControllerCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ReplicationControllerCreate")

	kubeClient := meta.(*config.KubeProviderClient)

	rc := build.ReplicationController(d, kubeClient.Version)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)
	rcCreate,errCreate := kubeRepControllers.Create(rc)

	if errCreate != nil {
		return errCreate
	}

	d.SetId(rcCreate.ObjectMeta.Name)
	return ReplicationControllerRead(d, meta)
}

func ReplicationControllerRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ReplicationControllerRead")

	kubeClient := meta.(*config.KubeProviderClient)
	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	rc,err := kubeRepControllers.Get(d.Id())

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

	update.ReplicationController(d, *rc)

	return nil
}

func ReplicationControllerUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ReplicationControllerUpdate")

	kubeClient := meta.(*config.KubeProviderClient)

	rc := build.ReplicationController(d, kubeClient.Version)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	originalRepController, err := kubeRepControllers.Get(d.Id())

	if err != nil {
		log.Printf("Error getting original replication controller: %v", err)
		return err
	}

	originalRepController.Spec.Replicas = 0
	_, err = kubeRepControllers.Update(originalRepController)

	if err != nil {
		log.Printf("Error updating replication controller replica count to 0: %v", err)
		return err
	}

	// Just to be sure lets delete any pods matching the selectors for the RC

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	podsList, podGetErr := kubePods.List(
		labels.SelectorFromSet(
			originalRepController.Spec.Selector),
			fields.Everything())

	if podGetErr != nil {
		log.Printf("Error listing pods for old RC: %v", podGetErr)
	} else if podsList != nil {
		if len(podsList.Items) > 0 {
			for _,pod := range podsList.Items {
				podDeleteErr := kubePods.Delete(pod.Name, nil)
				if podDeleteErr != nil {
					log.Printf("Error deleting pod (%s) for old RC: %v", pod.ObjectMeta.Name, podDeleteErr)
				}
			}
		}
	}

	updatedRepController, updateErr := kubeRepControllers.Update(rc)

	if updateErr != nil {
		return updateErr
	}

	d.SetId(updatedRepController.ObjectMeta.Name)
	return ReplicationControllerRead(d, meta)
}

func ReplicationControllerDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ReplicationControllerDelete")

	kubeClient := meta.(*config.KubeProviderClient)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	rc, err := kubeRepControllers.Get(d.Id())

	if err != nil {
		log.Printf("Error getting replication controller in delete: %v", err)
		return err
	}

	rc.Spec.Replicas = 0
	_, updateErr := kubeRepControllers.Update(rc)

	if updateErr != nil {
		log.Printf("Error updating replication controller replica count to 0: %v", updateErr)
		return updateErr
	}

	// Just to be sure lets delete any pods matching the selectors for the RC

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	podsList, listErr := kubePods.List(
		labels.SelectorFromSet(
			rc.Spec.Selector),
			fields.Everything())

	if listErr != nil {
		log.Printf("Error listing pods for RC (%s): %v", d.Id(), listErr)
	} else if podsList != nil {
		if len(podsList.Items) > 0 {
			for _, pod := range podsList.Items {
				deleteErr := kubePods.Delete(pod.Name, nil)
				if deleteErr != nil {
					log.Printf("Error deleting pod (%s) for RC (%s): %v",
						pod.Name,
						d.Id(),
						deleteErr)
				}
			}
		}
	}

	return kubeRepControllers.Delete(d.Id())
}
