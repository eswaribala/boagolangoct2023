provider "azurerm" {
  version = "~> 1.23"
  use_msi = true
  subscription_id = var.subscription_id
  tenant_id = var.tenant_id
  client_id = var.client_id
  client_secret = var.client_secret
}

resource "azurerm_resource_group" "rg" {
  name     = "vebgroup2022"
  location = "Southeast Asia"
}
