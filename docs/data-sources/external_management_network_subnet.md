---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Node Management"
layout: "aci"
page_title: "ACI: aci_external_management_network_subnet"
sidebar_current: "docs-aci-data-source-aci_external_management_network_subnet"
description: |-
  Data source for ACI External Management Network Subnet
---

# aci_external_management_network_subnet #

Data source for ACI External Management Network Subnet

## API Information ##

* Class: [mgmtSubnet](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtSubnet/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-mgmt/extmgmt-default/instp-{name}/subnet-[{ip}]`

## GUI Information ##

* Location: `Tenants (mgmt) -> External Management Network Instance Profiles -> Subnets`

## Example Usage ##

```hcl

data "aci_external_management_network_subnet" "example_external_management_network_instance_profile" {
  parent_dn = aci_external_management_network_instance_profile.example.id
  ip        = "1.1.1.0/24"
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_external_management_network_instance_profile](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/external_management_network_instance_profile) ([mgmtInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtInstP/overview))
* `ip` (ip) - (string) The external subnet IP address and subnet mask. This IP address is used for creating an external management entity. The subnet mask for the IP address to be imported from the outside into the fabric. The contracts associated with its parent instance profile (l3ext:InstP) are applied to the subnet.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the External Management Network Subnet object.
* `annotation` (annotation) - (string) The annotation of the External Management Network Subnet object. This attribute is supported in ACI versions: 3.2(1l) and later.
* `description` (descr) - (string) The description of the External Management Network Subnet object.
* `name` (name) - (string) The name of the External Management Network Subnet object.
* `name_alias` (nameAlias) - (string) The name alias of the External Management Network Subnet object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
