terraform {
  required_version = ">= 1.7.4"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=3.107.0"
    }
    modtm = {
      source  = "Azure/modtm"
      version = ">= 0.1.8, < 1.0"
    }
  }
}

provider "azurerm" {
  environment = "public"
  features {
    cognitive_account {
      purge_soft_delete_on_destroy = false
    }
  }
  subscription_id = var.secrets.subscription_id
  client_id       = var.secrets.client_id
  client_secret   = var.secrets.client_secret
  tenant_id       = var.secrets.tenant_id
}

provider "modtm" {
  // https://registry.terraform.io/providers/Azure/modtm/latest
  enabled = false
}
