resource "google_cloud_run_service" "microserviceurl" {
  name     = var.serviceName
  location = var.region

  template {
    spec {
      containers {
        image = var.imageMicroserviceurl
      }
    }
  }

  autogenerate_revision_name = true
}
