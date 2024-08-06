terraform {
  required_providers {
    aznamingtool = {
        source = "registry.terrafrom.io/rafaelherik/aznamingtool"
    }
  }
}


provider "aznamingtool" {
  api_key = "603a01da-b170-4a0f-8d55-f809332faacd"
  base_url = "http://localhost:8081"
  admin_password = "1q2w3e$R%T" 
}

resource "aznamingtool_resource_name" "example" {
  environment = "example"
  organization = "example"
  location = "eastus"
  resource_type = "example"
  project = "example"
}