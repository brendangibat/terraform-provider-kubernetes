package kubernetes

import (
	"log"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func resourceKubernetesReplicationController() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesReplicationControllerCreate,
		Read:   resourceKubernetesReplicationControllerRead,
		Update: resourceKubernetesReplicationControllerUpdate,
		Delete: resourceKubernetesReplicationControllerDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required:  true,
				ForceNew: true,
			},

			"config": &schema.Schema{
				Type:     schema.TypeString,
				Required:  true,
				ForceNew: true,
			},
		},
	}
}

func resourceKubernetesReplicationControllerCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerCreate")

	kubeClient := meta.(*KubeProviderClient)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	rc := api.ReplicationController{}

	err := json.Unmarshal([]byte(d.Get("config").(string)), &rc)
	if err != nil {
		return err
	}

	rc.APIVersion = kubeClient.Version
	rc.Kind = "ReplicationController"
	rc.ObjectMeta.Name = d.Get("name").(string)

	rcCreate,errCreate := kubeRepControllers.Create(&rc)

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

	serializedRepController, err := json.Marshal(rc)
	if err != nil {
		return err
	}

	d.Set("config", serializedRepController)

	return nil
}

func resourceKubernetesReplicationControllerUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesReplicationControllerUpdate")

	kubeClient := meta.(*KubeProviderClient)

	kubeRepControllers := kubeClient.KubeClient.ReplicationControllers(kubeClient.Namespace)

	rc := api.ReplicationController{}

	err := json.Unmarshal([]byte(d.Get("config").(string)), &rc)
	if err != nil {
		return err
	}

	rc.APIVersion = kubeClient.Version
	rc.Kind = "ReplicationController"
	rc.ObjectMeta.Name = d.Get("name").(string)

	updatedRepController,updateErr := kubeRepControllers.Update(&rc)

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
