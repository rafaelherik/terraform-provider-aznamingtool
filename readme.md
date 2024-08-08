# Azure Naming Tool - Terraform provider

[![Release](https://github.com/rafaelherik/terraform-provider-aznamingtool/actions/workflows/release.yml/badge.svg)](https://github.com/rafaelherik/terraform-provider-aznamingtool/actions/workflows/release.yml)

This provider is an implementation to enable the declarative API to interact with the [Azure Naming Tool](https://github.com/mspnp/AzureNamingTool)

> [!NOTE]
>The Azure Naming Tool was created to help administrators define and manage their naming conventions, while providing a simple interface for users to generate a compliant name. The tool was developed using a naming pattern based on Microsoft's best practices. Once an administrator has defined the organizational components, users can use the tool to generate a name for the desired Azure resource.


>[!IMPORTANT]
>The AzureNamingTool manages its data using JSON files, which can cause concurrency issues when multiple users or processes attempt to read or write to these files simultaneously. This can lead to data corruption or conflicts, making it challenging to maintain consistency and integrity in the stored data. To avoid this issue, the API client for this provider implements a simple queue system. Be aware that in large environments, this may still cause issues, as the queue system is implemented only on the client instance.


## Usage


### Provider intialization

```hcl copy 
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

## Compatibility 

There are the compatibility matrix between the Terraform Provider and the AzureNamingTool releases:

| rafaelherik/aznamingtool  | AzureNamingTool        |
|---------------------------|----------------------------|
| [1.0.0](https://registry.terraform.io/providers/rafaelherik/aznamingtool/1.0.0)    |  [4.20](https://github.com/mspnp/AzureNamingTool/releases/tag/v4.2.0) , [4.21](https://github.com/mspnp/AzureNamingTool/releases/tag/v4.2.1)      |


## Provider Documentatation

[Documentation](https://registry.terraform.io/providers/rafaelherik/aznamingtool/latest/docs)