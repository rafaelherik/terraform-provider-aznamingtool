# Azure Naming Tool Provider

## Overview

The Azure Naming Tool provider allows you to manage and configure naming conventions for your Azure resources using the Azure Naming Tool API. This provider can be used to ensure consistent naming across your Azure environment, adhering to your organizational policies.

## Example Usage

```hcl
provider "azurenaming" {
  base_url      = "https://api.namingtool.example.com"
  api_key       = var.api_key
  admin_password = var.admin_password
}

data "azurenamingtool_resource_name" "example" {
  environment   = "prod"
  organization  = "my-org"
  location      = "westus"
  resource_type = "vm"
  project       = "project1"
  instance      = "001"
}

resource "azurenamingtool_resource" "example" {
  name = data.azurenamingtool_resource_name.example.name
}
```

## Argument Reference

* `base_url` - (Optional) The base URL of the Azure Naming Tool API. Defaults to the value of the `AZ_NAMINGTOOL_BASEURL` environment variable if not provided.

* `api_key` - (Optional, Sensitive) The API key used to authenticate with the Azure Naming Tool API. Defaults to the value of the `AZ_NAMINGTOOL_APIKEY` environment variable if not provided.

* `admin_password` - (Optional, Sensitive) The administrator password used for privileged operations in the Azure Naming Tool. Defaults to the value of the `AZ_NAMINGTOOL_ADMINPASSWORD` environment variable if not provided.

> Note: The administrator passowrd is a sensitive information, only generated name deletion requires this password for now. If you enable naming duplication in the configuration you can omit the password.

## Attribute Reference

* `id` - A unique identifier for the resource.
* `name` - The name of the resource as configured or generated.
* `created_at` - The timestamp when the resource was created, in ISO 8601 format (e.g., 2023-08-01T12:34:56Z).
