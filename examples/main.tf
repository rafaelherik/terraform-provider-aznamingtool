terraform {
  required_providers {
    aznamingtool = {
        source = "registry.terrafrom.io/rafaelherik/aznamingtool"
        version = "1.0.0-beta"
    }        
  }
}


provider "aznamingtool" {
  api_key = "603a01da-b170-4a0f-8d55-f809332faacd"
  base_url = "http://localhost:8081"
  admin_password = "1q2w3e$R%T" 
}


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