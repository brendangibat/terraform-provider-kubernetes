package main

import (
	"log"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func resourceKubernetesService() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesServiceCreate,
		Read:   resourceKubernetesServiceRead,
		Update: resourceKubernetesServiceUpdate,
		Delete: resourceKubernetesServiceDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required:  true,
			},

			"config": &schema.Schema{
				Type:     schema.TypeString,
				Required:  true,
			},
		},
	}
}

func resourceKubernetesServiceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceCreate")

	kubeClient := meta.(*KubeProviderClient)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	serv := api.Service{}

	err := json.Unmarshal([]byte(d.Get("config").(string)), &serv)
	if err != nil {
		return err
	}

	serv.APIVersion = kubeClient.Version
	serv.Kind = "Service"
	serv.ObjectMeta.Name = d.Get("name").(string)

	servCreate,errCreate := kubeServices.Create(&serv)

	if errCreate != nil {
		return errCreate
	}

	d.SetId(servCreate.ObjectMeta.Name)
	return resourceKubernetesServiceRead(d, meta)
}

func resourceKubernetesServiceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceRead")

	kubeClient := meta.(*KubeProviderClient)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	serv,err := kubeServices.Get(d.Id())

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

	serializedService, err := json.Marshal(serv)
	if err != nil {
		return err
	}

	d.Set("config", serializedService)

	return nil
}

func resourceKubernetesServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceUpdate")

	kubeClient := meta.(*KubeProviderClient)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	serv := api.Service{}

	err := json.Unmarshal([]byte(d.Get("config").(string)), &serv)
	if err != nil {
		return err
	}

	serv.APIVersion = kubeClient.Version
	serv.Kind = "Service"
	serv.ObjectMeta.Name = d.Get("name").(string)

	updatedServ,updateErr := kubeServices.Update(&serv)

	if updateErr != nil {
		return updateErr
	}

	d.SetId(updatedServ.ObjectMeta.Name)
	return resourceKubernetesServiceRead(d, meta)
}

func resourceKubernetesServiceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceDelete")

	kubeClient := meta.(*KubeProviderClient)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	return kubeServices.Delete(d.Id())
}
