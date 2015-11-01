# terraform-provider-kubernetes

Currently supports pushing serialized JSON to kubernetes of :
 * Replication Controllers
 * Pods
 * Services

Operations supported:
 * Create
 * Read
 * Update
 * Delete

TODO: Pull inner objects of kubernetes definitions down in as terraform resource.

Example usage:

```
provider "kubernetes" {
    endpoint = "http://kube.domain.test:8080"
}

resource "kubernetes_pod" "blog-example" {
  spec {
    containers {
      image = "quay.io/kelcecil/kelcecil-com"
      name = "blog"
    }
  }
  metadata {
    labels {
      name = "blog"
      region = "us-east-1"
    }
    name = "blog"
  }
}

resource "kubernetes_replication_controller" "rc-service-example" {
    spec {
      replicas = 2
      template {
        spec {
          containers {
              image = "quay.io/kelcecil/chucksay"
              name = "chuck-as-a-service"
            }
        }
        metadata {
          labels {
            "k8s-app" = "chucksay"
            "name" = "chuck-as-a-service"
          }
          "name" = "chuck-as-a-service"
        }
      }
      selector {
        "k8s-app" = "chucksay"
        "name" = "chuck-as-a-service"
      }
    }
    metadata {
      labels {
        "k8s-app" = "chucksay"
      }
      name = "chuck-as-a-service"
    }
}

resource "kubernetes_service" "service-example" {
    spec {
      selector {
        "app" = "MyApp"
      }
      ports {
        protocol = "TCP"
        port = 80
      }
    }
    metadata {
      labels {
        "environment" = "dev"
      }
      name = "my-service"
    }
}
```
