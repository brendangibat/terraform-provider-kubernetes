provider "kubernetes" {
    endpoint = "http://internal-kube-dev-kube-master-elb-1635047245.us-west-2.elb.amazonaws.com:8080"
}

resource "template_file" "user_data" {
  filename = "service.json"
  lifecycle { create_before_destroy = true }
}

resource "kubernetes_service" "udhtest" {
    name = "udhtest"
    config = "${template_file.user_data.rendered}"
}
