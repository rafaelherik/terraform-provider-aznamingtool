# aznamingtool Provider

This is a terraform provider to consume the Azure Naming Tool API, to manage resource naming conventions for Azure Resource. This is an Open Source project supported by Microsoft and developed by the community.

## Example Usage

```hcl
terraform {
    required_providers {
         aznamingtool = {
            source = "registry.terrafrom.io/rafaelherik/aznamingtool"
        }
    }
}

provider "aznamingtool" {
  api_key = "YOUR_API_KEY"
  base_url = "http://localhost:8081"
  admin_password = "YOUR_ADMIN_PASSWORD" // THE ADMIN PASSWORD IS USED WHEN YOU WANT TO DELETE PREVIOUSLY GENERATED NAMES, IF IT'S NOT THE CASE YOU CAN SKIP THIS
}

```