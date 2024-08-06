---
layout: inline
---

# Azure Naming Tool - Terraform provider

[![Release](https://github.com/rafaelherik/terraform-provider-aznamingtool/actions/workflows/release.yml/badge.svg)](https://github.com/rafaelherik/terraform-provider-aznamingtool/actions/workflows/release.yml)

This provider is an implementation to enable the declarative API to interact with the [Azure Naming Tool](https://github.com/mspnp/AzureNamingTool)

> [!NOTE]
>The Azure Naming Tool was created to help administrators define and manage their naming conventions, while providing a simple interface for users to generate a compliant name. The tool was developed using a naming pattern based on Microsoft's best practices. Once an administrator has defined the organizational components, users can use the tool to generate a name for the desired Azure resource.


## Usage


### Provider intialization

```hcl annotate 
# 
terraform {
    required_providers {
         aznamingtool = {
            # The azure naming tool provider registry URL
            source = "registry.terraform.io/rafaelherik/aznamingtool"
        }
    }
}

provider "aznamingtool" {
  # The API Key provided by your hosted [Azure Naming tool](https://github.com/mspnp/AzureNamingTool/wiki/Using-the-API)
  api_key = "YOUR_API_KEY"
  # The URL of you published Azure Naming Tool application
  base_url = "http://localhost:8081"
  # The `admin_password` is used when you want to delete previosly generated names. 'Optional'
  admin_password = "YOUR_ADMIN_PASSWORD" 
}

```

[Documentation](https://registry.terraform.io/providers/rafaelherik/aznamingtool/latest/docs)