# resource_name Resource

The ```resource_name``` resource is a representation of the generated name from the Azure Naming Tool.

## Example Usage

```hcl
resource "aznamingtool_resource_name" "example" {
  environment = "dev"
  organization = "MyOrg"
  location = "eastus"
  resource_type = "Compute/Virtual Machine/Linux"
  project = "myProject"
  instance = 1
}

```

The Default Policy configuration requires all the arguments to be able to return a new name.