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

func ServiceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEDUG] ServiceCreate")

	kubeClient := meta.(*config.KubeProviderClient)

	srvc := build.Service(d, kubeClient.Version)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)
	srvcCreate, err := kubeServices.Create(srvc)

	if err != nil {
		return err
	}

	d.SetId(srvcCreate.ObjectMeta.Name)
	return ServiceRead(d, meta)
}

func ServiceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ServiceRead")

	kubeClient := meta.(*config.KubeProviderClient)
	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	service, err := kubeServices.Get(d.Id())

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

	update.Service(d, *service)

	return nil
}

func ServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ServiceUpdate")

	kubeClient := meta.(*config.KubeProviderClient)

	srvc := build.Service(d, kubeClient.Version)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)
	// This will most likely need deleted and recreated
	srvcCreate, err := kubeServices.Update(srvc)

	if err != nil {
		return err
	}

	d.SetId(srvcCreate.ObjectMeta.Name)
	return ServiceRead(d, meta)
}

func ServiceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] ServiceDelete")

	kubeClient := meta.(*config.KubeProviderClient)

	kubeServices := kubeClient.KubeClient.Services(kubeClient.Namespace)

	return kubeServices.Delete(d.Id())
}
