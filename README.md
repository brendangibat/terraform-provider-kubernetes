# terraform-provider-kubernetes

Currently supports pushing serialized JSON to kubernetes of :
 * Replication Controllers

Coming soon:
 * Services
 * Pods

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


resource "kubernetes_replication_controller" "logstash-test" {
    spec {
      replicas = 2
      template {
        spec {
          containers {
              image = "docker.moveaws.com/logstash-logging:latest"
              args = ["-e", "input {redis {host => 'redis.test' key => 'logging' data_type => 'list' type => 'redis'}}", "-e", "output { if [options][_logging_index] {elasticsearch {hosts => [\"es.test\"] index => '%{[options][_logging_index]}-%{+YYYY.MM.dd}'}} else {elasticsearch {hosts => ['es.test'] index => 'logging-%{+YYYY.MM.dd}'}}}"]
              name = "logstash-logging"
            }
        }
        metadata {
          labels {
            "k8s-app" = "logstash-logging"
            "name" = "logstash-logging"
          }
          "name" = "logstash-logging"
        }
      }
      selector {
        "k8s-app" = "logstash-logging"
        "name" = "logstash-logging"
      }
    }
    metadata {
      labels {
        "k8s-app" = "logstash-logging"
      }
      name = "logstash-logging"
    }
}
```
