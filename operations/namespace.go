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

func NamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] NamespaceCreate")

	kubeClient := meta.(*config.KubeProviderClient)

	namespace := build.Namespace(d, kubeClient.Version)

	kubeNamspaces := kubeClient.KubeClient.Namespaces()
	namespaceCreate, err := kubeNamspaces.Create(namespace)

	if err != nil {
		return err
	}

	d.SetId(namespaceCreate.ObjectMeta.Name)
	return NamespaceRead(d, meta)
}

func NamespaceRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] NamespaceRead")

	kubeClient := meta.(*config.KubeProviderClient)
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

	update.Namespace(d, *namespace)

	return nil
}

func NamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] NamespaceUpdate")

	kubeClient := meta.(*config.KubeProviderClient)

	namespace := build.Namespace(d, kubeClient.Version)

	kubeNamespaces := kubeClient.KubeClient.Namespaces()
	// This will most likely need deleted and recreated
	updatedNamespace, err := kubeNamespaces.Update(namespace)

	if err != nil {
		return err
	}

	d.SetId(updatedNamespace.ObjectMeta.Name)
	return NamespaceRead(d, meta)
}

func NamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] NamespaceDelete")

	kubeClient := meta.(*config.KubeProviderClient)

	kubeNamespaces := kubeClient.KubeClient.Namespaces()

	return kubeNamespaces.Delete(d.Id())
}
