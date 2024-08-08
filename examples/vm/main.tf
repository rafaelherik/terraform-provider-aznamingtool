terraform {
  required_providers {
    aznamingtool = {
        source = "registry.terraform.io/rafaelherik/aznamingtool"
        version = "1.0.0-beta"
    }        
   azurerm = {
      source = "hashicorp/azurerm"     
    }
  }
}

provider "aznamingtool" {}

provider "azurerm" {
  features {}
}



