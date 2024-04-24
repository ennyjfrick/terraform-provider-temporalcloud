terraform {
  required_providers {
    temporalcloud = {
      source = "temporalio/temporalcloud"
    }
  }
}

provider "temporalcloud" {

}

resource "temporalcloud_account_metrics" "terraform" {
  enabled            = true
  accepted_client_ca = base64encode(file("${path.module}/ca.pem"))
}