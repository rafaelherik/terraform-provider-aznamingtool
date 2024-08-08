
resource "aznamingtool_resource_name" "aznt-rg" {  
  components = merge(var.project_configuration, {
    resource_type= "rg"    
  })
}

resource "aznamingtool_resource_name" "aznt-vm" { 
  resource_type_id = 85 
  components = merge(var.project_configuration, {
    resource_instance = "1"
    resource_type = "vm"    
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