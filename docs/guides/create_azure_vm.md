# Creating a Virtual Machine with names coming from aznamingtool provider

### Prerequisites

Ensure you have configured the AzureNamingTool provider using environment variables as detailed in the previous documentation.

## Steps to Create a Virtual Machine

1. Define the providers:

```hcl 
terraform {
  required_providers {
    aznamingtool = {
        source = "registry.terraform.io/rafaelherik/aznamingtool"
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

```

2. Create a variable file:

```hcl

variable "project_configuration" {
  type = map(string)
  default = {} 
}

```

3. Configure a .tfvars file to use:

Create a .tfvars file to define naming convention properties. This file allows you to maintain consistent naming conventions across different projects and environments by reusing attributes. You can have multiple configurations tailored to specific needs, ensuring a standardized approach to naming resources.

```tfvars

project_configuration = {
    resource_environment = "dev" #The environment
    resource_location = "euw" #The Location - West Europe
    resource_proj_app_svc = "tnp" #The Project 
}

```

4. Configure the required resources:

```hcl

resource "aznamingtool_resource_name" "aznt-rg" {  
  components = merge(var.project_configuration, {
    resource_type= "rg"    
    resource_instance = "1"
  })
}

resource "aznamingtool_resource_name" "aznt-vm" { 
  resource_type_id = 85 
  components = merge(var.project_configuration, {
    resource_instance = "1"
    resource_type = "vm"        
    resource_function = "func"
  })
}

resource "aznamingtool_resource_name" "aznt-vnet" {   
  components = merge(var.project_configuration, {   
    resource_type = "vnet"
    resource_instance = "1"
  })
}

resource "aznamingtool_resource_name" "aznt-vnet-snet" {   
  components = merge(var.project_configuration, {
    resource_instance = "1"
    resource_type = "snet"
  })
}

resource "aznamingtool_resource_name" "aznt-vm-nic" {   
  components = merge(var.project_configuration, {
    resource_instance = "1"
    resource_type = "nic"
  })
}

resource "aznamingtool_resource_name" "aznt-osdisk" {   
  components = merge(var.project_configuration, {
    resource_instance = "1"
    resource_type = "osdisk"    
    resource_function = "func"
  })
}


resource "azurerm_resource_group" "az-rg" {
  name     = aznamingtool_resource_name.aznt-rg.resource_name
  location = "West Europe"
}

resource "azurerm_virtual_network" "az-vnet" {
  name                = aznamingtool_resource_name.aznt-vnet.resource_name
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.az-rg.location
  resource_group_name = aznamingtool_resource_name.aznt-rg.resource_name
}

resource "azurerm_subnet" "az-subnet" {
  name                 = "example-subnet"
  resource_group_name  = azurerm_resource_group.az-rg.name
  virtual_network_name = azurerm_virtual_network.az-vnet.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_network_interface" "az-nic" {
  name                = "example-nic"
  location            = azurerm_resource_group.az-rg.location
  resource_group_name = azurerm_resource_group.az-rg.name

  ip_configuration {
    name                          = "internal"
    subnet_id                     = azurerm_subnet.az-subnet.id
    private_ip_address_allocation = "Dynamic"
  }
}

resource "azurerm_virtual_machine" "az-vm" {
  name                  = "example-machine"
  location              = azurerm_resource_group.az-rg.location
  resource_group_name   = azurerm_resource_group.az-rg.name
  network_interface_ids = [azurerm_network_interface.az-nic.id]
  vm_size               = "Standard_DS1_v2"

  storage_os_disk {
    name              = aznamingtool_resource_name.aznt-osdisk.resource_name
    caching           = "ReadWrite"
    create_option     = "FromImage"
    managed_disk_type = "Standard_LRS"
  }

  storage_image_reference {
    publisher = "Canonical"
    offer     = "UbuntuServer"
    sku       = "18.04-LTS"
    version   = "latest"
  }

  os_profile {
    computer_name  = "examplevm"
    admin_username = "adminuser"
    admin_password = "Password1234!"
  }

  os_profile_linux_config {
    disable_password_authentication = false
  }
}
```

After the configuration you can run the terraform commands to create the infra:

```shell

terraform init
terraform plan
```

Your plan should be like that:

```bash
Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # aznamingtool_resource_name.aznt-osdisk will be created
  + resource "aznamingtool_resource_name" "aznt-osdisk" {
      + components         = {
          + "resource_environment"  = "dev"
          + "resource_function"     = "func"
          + "resource_instance"     = "1"
          + "resource_location"     = "euw"
          + "resource_proj_app_svc" = "tnp"
          + "resource_type"         = "osdisk"
        }
      + created_on         = (known after apply)
      + id                 = (known after apply)
      + resource_name      = (known after apply)
      + resource_type_name = (known after apply)
    }

  # aznamingtool_resource_name.aznt-rg will be created
  + resource "aznamingtool_resource_name" "aznt-rg" {
      + components         = {
          + "resource_environment"  = "dev"
          + "resource_instance"     = "1"
          + "resource_location"     = "euw"
          + "resource_proj_app_svc" = "tnp"
          + "resource_type"         = "rg"
        }
      + created_on         = (known after apply)
      + id                 = (known after apply)
      + resource_name      = (known after apply)
      + resource_type_name = (known after apply)
    }

  # aznamingtool_resource_name.aznt-vm will be created
  + resource "aznamingtool_resource_name" "aznt-vm" {
      + components         = {
          + "resource_environment"  = "dev"
          + "resource_function"     = "func"
          + "resource_instance"     = "1"
          + "resource_location"     = "euw"
          + "resource_proj_app_svc" = "tnp"
          + "resource_type"         = "vm"
        }
      + created_on         = (known after apply)
      + id                 = (known after apply)
      + resource_name      = (known after apply)
      + resource_type_id   = 85
      + resource_type_name = (known after apply)
    }

  # aznamingtool_resource_name.aznt-vm-nic will be created
  + resource "aznamingtool_resource_name" "aznt-vm-nic" {
      + components         = {
          + "resource_environment"  = "dev"
          + "resource_instance"     = "1"
          + "resource_location"     = "euw"
          + "resource_proj_app_svc" = "tnp"
          + "resource_type"         = "nic"
        }
      + created_on         = (known after apply)
      + id                 = (known after apply)
      + resource_name      = (known after apply)
      + resource_type_name = (known after apply)
    }

  # aznamingtool_resource_name.aznt-vnet will be created
  + resource "aznamingtool_resource_name" "aznt-vnet" {
      + components         = {
          + "resource_environment"  = "dev"
          + "resource_instance"     = "1"
          + "resource_location"     = "euw"
          + "resource_proj_app_svc" = "tnp"
          + "resource_type"         = "vnet"
        }
      + created_on         = (known after apply)
      + id                 = (known after apply)
      + resource_name      = (known after apply)
      + resource_type_name = (known after apply)
    }

  # aznamingtool_resource_name.aznt-vnet-snet will be created
  + resource "aznamingtool_resource_name" "aznt-vnet-snet" {
      + components         = {
          + "resource_environment"  = "dev"
          + "resource_instance"     = "1"
          + "resource_location"     = "euw"
          + "resource_proj_app_svc" = "tnp"
          + "resource_type"         = "snet"
        }
      + created_on         = (known after apply)
      + id                 = (known after apply)
      + resource_name      = (known after apply)
      + resource_type_name = (known after apply)
    }

  # azurerm_network_interface.az-nic will be created
  + resource "azurerm_network_interface" "az-nic" {
      + accelerated_networking_enabled = (known after apply)
      + applied_dns_servers            = (known after apply)
      + dns_servers                    = (known after apply)
      + enable_accelerated_networking  = (known after apply)
      + enable_ip_forwarding           = (known after apply)
      + id                             = (known after apply)
      + internal_domain_name_suffix    = (known after apply)
      + ip_forwarding_enabled          = (known after apply)
      + location                       = "westeurope"
      + mac_address                    = (known after apply)
      + name                           = "example-nic"
      + private_ip_address             = (known after apply)
      + private_ip_addresses           = (known after apply)
      + resource_group_name            = (known after apply)
      + virtual_machine_id             = (known after apply)

      + ip_configuration {
          + gateway_load_balancer_frontend_ip_configuration_id = (known after apply)
          + name                                               = "internal"
          + primary                                            = (known after apply)
          + private_ip_address                                 = (known after apply)
          + private_ip_address_allocation                      = "Dynamic"
          + private_ip_address_version                         = "IPv4"
          + subnet_id                                          = (known after apply)
        }
    }

  # azurerm_resource_group.az-rg will be created
  + resource "azurerm_resource_group" "az-rg" {
      + id       = (known after apply)
      + location = "westeurope"
      + name     = (known after apply)
    }

  # azurerm_subnet.az-subnet will be created
  + resource "azurerm_subnet" "az-subnet" {
      + address_prefixes                               = [
          + "10.0.2.0/24",
        ]
      + default_outbound_access_enabled                = true
      + enforce_private_link_endpoint_network_policies = (known after apply)
      + enforce_private_link_service_network_policies  = (known after apply)
      + id                                             = (known after apply)
      + name                                           = "example-subnet"
      + private_endpoint_network_policies              = (known after apply)
      + private_endpoint_network_policies_enabled      = (known after apply)
      + private_link_service_network_policies_enabled  = (known after apply)
      + resource_group_name                            = (known after apply)
      + virtual_network_name                           = (known after apply)
    }

  # azurerm_virtual_machine.az-vm will be created
  + resource "azurerm_virtual_machine" "az-vm" {
      + availability_set_id              = (known after apply)
      + delete_data_disks_on_termination = false
      + delete_os_disk_on_termination    = false
      + id                               = (known after apply)
      + license_type                     = (known after apply)
      + location                         = "westeurope"
      + name                             = "example-machine"
      + network_interface_ids            = (known after apply)
      + resource_group_name              = (known after apply)
      + vm_size                          = "Standard_DS1_v2"

      + os_profile {
          # At least one attribute in this block is (or was) sensitive,
          # so its contents will not be displayed.
        }

      + os_profile_linux_config {
          + disable_password_authentication = false
        }

      + storage_image_reference {
            id        = null
          + offer     = "UbuntuServer"
          + publisher = "Canonical"
          + sku       = "18.04-LTS"
          + version   = "latest"
        }

      + storage_os_disk {
          + caching                   = "ReadWrite"
          + create_option             = "FromImage"
          + disk_size_gb              = (known after apply)
          + managed_disk_id           = (known after apply)
          + managed_disk_type         = "Standard_LRS"
          + name                      = (known after apply)
          + os_type                   = (known after apply)
          + write_accelerator_enabled = false
        }
    }

  # azurerm_virtual_network.az-vnet will be created
  + resource "azurerm_virtual_network" "az-vnet" {
      + address_space       = [
          + "10.0.0.0/16",
        ]
      + dns_servers         = (known after apply)
      + guid                = (known after apply)
      + id                  = (known after apply)
      + location            = "westeurope"
      + name                = (known after apply)
      + resource_group_name = (known after apply)
      + subnet              = (known after apply)
    }

Plan: 11 to add, 0 to change, 0 to destroy.
```

After this you can apply the plan:

```shell
terraform apply --auto-approve
```

The result:

```shell
Plan: 11 to add, 0 to change, 0 to destroy.
aznamingtool_resource_name.aznt-vnet: Creating...
aznamingtool_resource_name.aznt-vm: Creating...
aznamingtool_resource_name.aznt-rg: Creating...
aznamingtool_resource_name.aznt-vm-nic: Creating...
aznamingtool_resource_name.aznt-osdisk: Creating...
aznamingtool_resource_name.aznt-vnet-snet: Creating...
aznamingtool_resource_name.aznt-vnet: Creation complete after 1s
aznamingtool_resource_name.aznt-osdisk: Creation complete after 1s
aznamingtool_resource_name.aznt-vm-nic: Creation complete after 1s
aznamingtool_resource_name.aznt-vm: Creation complete after 1s
aznamingtool_resource_name.aznt-rg: Creation complete after 1s
aznamingtool_resource_name.aznt-vnet-snet: Creation complete after 1s
azurerm_resource_group.az-rg: Creating...
azurerm_resource_group.az-rg: Creation complete after 10s [id=/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-tnp-dev-euw-1]
azurerm_virtual_network.az-vnet: Creating...
azurerm_virtual_network.az-vnet: Creation complete after 5s [id=/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-tnp-dev-euw-1/providers/Microsoft.Network/virtualNetworks/vnet-tnp-dev-euw-1]
azurerm_subnet.az-subnet: Creating...
azurerm_subnet.az-subnet: Creation complete after 4s [id=/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-tnp-dev-euw-1/providers/Microsoft.Network/virtualNetworks/vnet-tnp-dev-euw-1/subnets/example-subnet]
azurerm_network_interface.az-nic: Creating...
azurerm_network_interface.az-nic: Still creating... [10s elapsed]
azurerm_network_interface.az-nic: Creation complete after 11s [id=/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-tnp-dev-euw-1/providers/Microsoft.Network/networkInterfaces/example-nic]
azurerm_virtual_machine.az-vm: Creating...
azurerm_virtual_machine.az-vm: Still creating... [10s elapsed]
azurerm_virtual_machine.az-vm: Creation complete after 20s [id=/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg-tnp-dev-euw-1/providers/Microsoft.Compute/virtualMachines/example-machine]

Apply complete! Resources: 11 added, 0 changed, 0 destroyed.

```

There is a sample content from the final state file:

```json
{
  "version": 4,
  "terraform_version": "1.8.0",
  "serial": 38,
  "lineage": "3d5be281-6b49-7de6-da89-76abd5ccced0",
  "outputs": {},
  "resources": [
    {
      "mode": "managed",
      "type": "aznamingtool_resource_name",
      "name": "aznt-osdisk",
      "provider": "provider[\"registry.terraform.io/rafaelherik/aznamingtool\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "components": {
              "resource_environment": "dev",
              "resource_function": "func",
              "resource_instance": "1",
              "resource_location": "euw",
              "resource_proj_app_svc": "tnp",
              "resource_type": "osdisk"
            },
            "created_on": "2024-08-08T11:51:10.6962877+00:00",
            "id": 4,
            "resource_name": "osdisktnpfuncdeveuw1",
            "resource_type_id": null,
            "resource_type_name": "Compute/disks - OS Disk"
          },
          "sensitive_attributes": []
        }
      ]
    },
    ....
}

```