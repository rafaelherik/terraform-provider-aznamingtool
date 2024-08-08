variable "project_configuration" {
  type = map(string)
  default = {
    resource_environment = "dev"
    resource_location = "euw"
    resource_proj_app_svc = "tnp"    
  } 

  
}