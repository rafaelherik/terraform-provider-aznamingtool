terraform {
  required_providers {
    aznamingtool = {
        source = "registry.terraform.io/rafaelherik/aznamingtool"
    }
  }
}

provider "aznamingtool" {}



variable "project_configuration" {
  type = map(string)
  default = {
    resource_environment = "dev"
    resource_location = "euw"
    resource_proj_app_svc = "tnp"
  } 
}
resource "aznamingtool_resource_name" "my-resource-group" {  
  components = merge(var.project_configuration, {    
    resource_type= "rg"   
    resource_instance = "1" 
  })
}