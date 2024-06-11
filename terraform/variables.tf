variable "azure_rosources" {
  description = "Azure resources description"
  default = {
    resource_group_name = "chatai"
    location            = "westeurope"
  }
}

variable "azure_openai" {
  description = "Azure OpenAI resources"
  default = {
    source       = "Azure/openai/azurerm"
    version      = "0.1.3"
    account_name = "chatai"
    # location = "West Europe"
    sku_name                      = "S0"
    environment                   = "dev"
    public_network_access_enabled = true
    deployment = {
      gpt-35-turbo = {
        name          = "gpt-35-turbo"
        model_format  = "OpenAI"
        model_name    = "gpt-35-turbo"
        model_version = "0301"
        scale_type    = "Standard"
      },
    },
    tags = {
      terraform = true
    }
  }
}
