resource "azurerm_resource_group" "chatai_rg" {
  location = var.azure_rosources.location
  name     = var.azure_rosources.resource_group_name
}

module "openai" {
  source  = "Azure/openai/azurerm"
  version = "0.1.3"
  # source  = var.azure_openai.source
  # version = var.azure_openai.version
  account_name                  = var.azure_openai.account_name
  sku_name                      = var.azure_openai.sku_name
  environment                   = var.azure_openai.environment
  public_network_access_enabled = var.azure_openai.public_network_access_enabled
  resource_group_name           = azurerm_resource_group.chatai_rg.name
  location                      = azurerm_resource_group.chatai_rg.location
  deployment                    = var.azure_openai.deployment
  tags                          = var.azure_openai.tags
  depends_on = [
    azurerm_resource_group.chatai_rg
  ]
}
