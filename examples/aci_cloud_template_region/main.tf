terraform {
  required_providers {
    aci = {
      source = "ciscodevnet/aci"
    }
  }
}

provider "aci" {
  username = ""
  password = ""
  url      = ""
  insecure = true
}

# In order to add the cloud subnets in the azure cloud networking, hub networking needs to be disabled
resource "aci_cloud_template_region_detail" "hub_network" {
  parent_dn      = "uni/tn-infra/infranetwork-default/intnetwork-default/provider-azure-region-westus"
  hub_networking = "disabled"
}

data "aci_cloud_context_profile" "hub_vnet" {
  tenant_dn  = "uni/tn-infra"
  name       = "ct_ctxprofile_westus"
}

resource "aci_cloud_cidr_pool" "azure_cloud_cidr_pool" {
  cloud_context_profile_dn = data.aci_cloud_context_profile.hub_vnet.id
  addr                     = "10.8.0.0/25"
  primary                  = "no"
}

resource "aci_cloud_subnet" "azure_subnets" {
  cloud_cidr_pool_dn              = aci_cloud_cidr_pool.cloud_cidr_pool.id
  name                            = "test"
  ip                              = "10.8.0.3/25"
  relation_cloud_rs_subnet_to_ctx = "uni/tn-infra/ctx-ave-ctrl"
}