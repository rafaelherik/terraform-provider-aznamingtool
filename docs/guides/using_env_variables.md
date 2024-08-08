# Configuration with Environment Variables

To configure the AzureNamingTool provider using environment variables, follow these steps:

## Set the Environment Variables:

Define the following environment variables in your system:

```shell
export AZ_NAMINGTOOL_APIKEY="your_api_key"
export AZ_NAMINGTOOL_BASEURL="http://localhost:8081"
export AZ_NAMINGTOOL_ADMINPASSWORD="your_admin_password" # Mandatory to destroy names
```
Then you can keep the provider configuration empty:

```hcl
provider "aznamingtool" {
  api_key        = var.aznamingtool_api_key
  base_url       = var.aznamingtool_base_url
  admin_password = var.aznamingtool_admin_password // Optional
}


```