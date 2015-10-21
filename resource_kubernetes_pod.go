package main

import (
	"log"
	"encoding/json"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/errors"
)

func resourceKubernetesPod() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesPodCreate,
		Read:   resourceKubernetesPodRead,
		Update: resourceKubernetesPodUpdate,
		Delete: resourceKubernetesPodDelete,
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

func resourceKubernetesPodCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodCreate")

	kubeClient := meta.(*KubeProviderClient)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	pod := api.Pod{}

	err := json.Unmarshal([]byte(d.Get("config").(string)), &pod)
	if err != nil {
		return err
	}

	pod.APIVersion = kubeClient.Version
	pod.Kind = "Pod"
	pod.ObjectMeta.Name = d.Get("name").(string)

	podCreate,errCreate := kubePods.Create(&pod)

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

	pod,err := kubePods.Get(d.Id())

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

	serializedPod, err := json.Marshal(pod)
	if err != nil {
		return err
	}

	d.Set("config", serializedPod)

	return nil
}

func resourceKubernetesPodUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodUpdate")

	kubeClient := meta.(*KubeProviderClient)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	pod := api.Pod{}

	err := json.Unmarshal([]byte(d.Get("config").(string)), &pod)
	if err != nil {
		return err
	}

	pod.APIVersion = kubeClient.Version
	pod.Kind = "Pod"
	pod.ObjectMeta.Name = d.Get("name").(string)

	updatedServ,updateErr := kubePods.Update(&pod)

	if updateErr != nil {
		return updateErr
	}

	d.SetId(updatedServ.ObjectMeta.Name)
	return resourceKubernetesPodRead(d, meta)
}

func resourceKubernetesPodDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesPodDelete")

	kubeClient := meta.(*KubeProviderClient)

	kubePods := kubeClient.KubeClient.Pods(kubeClient.Namespace)

	return kubePods.Delete(d.Id(), nil)
}
