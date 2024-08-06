# Azure Naming Tool - Terraform provider

[![Release](https://github.com/rafaelherik/terraform-provider-aznamingtool/actions/workflows/release.yml/badge.svg)](https://github.com/rafaelherik/terraform-provider-aznamingtool/actions/workflows/release.yml)

This provider is an implementation to enable the declarative API to interact with the [Azure Naming Tool](https://github.com/mspnp/AzureNamingTool)

"The Azure Naming Tool was created to help administrators define and manage their naming conventions, while providing a simple interface for users to generate a compliant name. The tool was developed using a naming pattern based on Microsoft's best practices. Once an administrator has defined the organizational components, users can use the tool to generate a name for the desired Azure resource."


## Usage


### Provider intialization

```hcl
terraform {
    required_providers {
         aznamingtool = {
            source = "registry.terraform.io/rafaelherik/aznamingtool"
        }
    }
}

provider "aznamingtool" {
  api_key = "YOUR_API_KEY"
  base_url = "http://localhost:8081"
  admin_password = "YOUR_ADMIN_PASSWORD" // THE ADMIN PASSWORD IS USED WHEN YOU WANT TO DELETE PREVIOUSLY GENERATED NAMES, IF IT'S NOT THE CASE YOU CAN SKIP THIS
}

```

[Documentation](https://registry.terraform.io/providers/rafaelherik/aznamingtool/latest/docs)