package kubernetes

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"k8s.io/kubernetes/pkg/client"
	"k8s.io/kubernetes/pkg/runtime"
)

func resourceKubernetesService() *schema.Resource {
	return &schema.Resource{
		Create: resourceKubernetesServiceCreate,
		Read:   resourceKubernetesServiceRead,
		Schema: map[string]*schema.Schema{
			"endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceKubernetesServiceCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG] resourceKubernetesServiceCreate")
	return nil
}

func resourceKubernetesServiceRead(d *schema.ResourceData, meta interface{}) error {

	client := client.NewRESTClient("http://internal-kube-dev-kube-master-elb-1635047245.us-west-2.elb.amazonaws.com:8080",
		"v1.0.6",
		runtime.Codec.Codec,
		0,
		0)
	client.Services("default").
	log.Printf("[DEBUG] resourceKubernetesServiceRead")
	return nil
}
