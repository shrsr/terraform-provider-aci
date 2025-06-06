---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Networking"
layout: "aci"
page_title: "ACI: aci_relation_from_vrf_to_address_family_ospf_timers"
sidebar_current: "docs-aci-resource-aci_relation_from_vrf_to_address_family_ospf_timers"
description: |-
  Manages ACI Relation From VRF To Address Family OSPF Timers
---

# aci_relation_from_vrf_to_address_family_ospf_timers #

Manages ACI Relation From VRF To Address Family OSPF Timers



## API Information ##

* Class: [fvRsCtxToOspfCtxPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCtxToOspfCtxPol/overview)

* Supported in ACI versions: 1.1(1j) and later.

* Distinguished Name Format: `uni/tn-{name}/ctx-{name}/rsctxToOspfCtxPol-[{tnOspfCtxPolName}]-{af}`

## GUI Information ##

* Location: `Tenants -> Networking -> VRFs -> Policy -> OSPF Context Per Address Family`

## Example Usage ##

The configuration snippet below creates a Relation From VRF To Address Family OSPF Timers with only required attributes.

```hcl

resource "aci_relation_from_vrf_to_address_family_ospf_timers" "example_vrf" {
  parent_dn        = aci_vrf.example.id
  address_family   = "ipv4-ucast"
  ospf_timers_name = aci_ospf_timers.example.name
}

```
The configuration snippet below shows all possible attributes of the Relation From VRF To Address Family OSPF Timers.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_relation_from_vrf_to_address_family_ospf_timers" "full_example_vrf" {
  parent_dn        = aci_vrf.example.id
  address_family   = "ipv4-ucast"
  annotation       = "annotation"
  ospf_timers_name = aci_ospf_timers.example.name
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
}

```

All examples for the Relation From VRF To Address Family OSPF Timers resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_relation_from_vrf_to_address_family_ospf_timers) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_vrf](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/vrf) ([fvCtx](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCtx/overview))
* `address_family` (af) - (string) The type of address family for the Relation From VRF To Address Family OSPF Timers.
  - Valid Values: `ipv4-ucast`, `ipv6-ucast`.
* `ospf_timers_name` (tnOspfCtxPolName) - (string) The name of the OSPF timers policy associated with this object. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/ospf_timers) with `aci_ospf_timers.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/ospf_timers) with `data.aci_ospf_timers.example.name`.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation From VRF To Address Family OSPF Timers object.

### Optional ###

* `annotation` (annotation) - (string) The annotation of the Relation From VRF To Address Family OSPF Timers object. This attribute is supported in ACI versions: 3.2(1l) and later.
  - Default: `orchestrator:terraform`
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing Relation From VRF To Address Family OSPF Timers can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_relation_from_vrf_to_address_family_ospf_timers.example_vrf uni/tn-{name}/ctx-{name}/rsctxToOspfCtxPol-[{tnOspfCtxPolName}]-{af}
```

Starting in Terraform version 1.5, an existing Relation From VRF To Address Family OSPF Timers can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ctx-{name}/rsctxToOspfCtxPol-[{tnOspfCtxPolName}]-{af}"
  to = aci_relation_from_vrf_to_address_family_ospf_timers.example_vrf
}
```
