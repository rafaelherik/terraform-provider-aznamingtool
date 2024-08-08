terraform {
  required_providers {
    aznamingtool = {
        source = "registry.terrafrom.io/rafaelherik/aznamingtool"
        version = "1.0.0-beta"
    }        
   azurerm = {
      source = "hashicorp/azurerm"
      version = "2.46.0"
    }
  }
}

variable "project_configuration" {
  type = map(string)
  default = {
    resource_environment = "dev"
    resource_location = "euw"
    resource_proj_app_svc = "tnp"
    resource_function = "func"
  } 
}


provider "aznamingtool" {
  api_key = "603a01da-b170-4a0f-8d55-f809332faacd"
  base_url = "http://localhost:8081"
  admin_password = "1q2w3e$R%T" 
}

resource "aznamingtool_resource_name" "my-resource-group" {  
  components = {
    resource_environment = var.project_configuration["resource_environment"]
    resource_location =  var.project_configuration["resource_location"]
    resource_proj_app_svc = var
    resource_function = var.project_configuration["resource_function"]  
    resource_type= "rg"    
  }
}

resource "aznamingtool_resource_name" "my-virutal-machine" { 
  resource_type_id = 85 
  components = {
    resource_environment = var.project_configuration["resource_environment"]
    resource_location =  var.project_configuration["resource_location"]
    resource_proj_app_svc = var
    resource_function = var.project_configuration["resource_function"]
    resource_instance = "1"
  }
}



reso


output "vm-linux-name" {
  value = aznamingtool_resource_name.vm-linux-name.name
}