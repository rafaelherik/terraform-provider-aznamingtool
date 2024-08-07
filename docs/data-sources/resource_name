# aznamingtool_resource_name_data_source Data Source

The `aznamingtool_resource_name` data source retrieves the details of a generated resource name from the Azure Naming Tool.

## Example Usage

```hcl
data "aznamingtool_resource_name_data_source" "example" {
  id = "12345"
}
```

## Argument Reference

* `id` - (Required) The unique identifier for the resource.


## Attributes Reference

* `id` - The unique identifier for the resource.
* `resource_name` - The generated name of the resource as configured or generated. This is often a combination of other attributes to create a unique and meaningful name.
* `resource_type_name` - The name of the resource type as determined by the resource type ID.
* `components` - A map containing the various parts of the resource name as key-value pairs. This includes the environment, function, instance, location, organization, project application service, unit department, and any custom components.
* `created_on` - The timestamp when the resource was created. This is typically in ISO 8601 format (e.g., 2023-08-01T12:34:56Z), and is useful for auditing and management purposes.

## Import

Data sources cannot be imported, but you can reference the id in your configuration to retrieve the existing resource details.

```hcl
data "aznamingtool_resource_name_data_source" "example" {
  id = "12345"
}

```

> Note: Replace 12345 with the actual ID of the resource you wish to retrieve.