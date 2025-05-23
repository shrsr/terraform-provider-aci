---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Networking"
layout: "aci"
page_title: "ACI: aci_vrf_fallback_route"
sidebar_current: "docs-aci-resource-aci_vrf_fallback_route"
description: |-
  Manages ACI VRF Fallback Route
---

# aci_vrf_fallback_route #

Manages ACI VRF Fallback Route

  -> This resource should not be used in combination with the `vrf_fallback_routes` nested attribute of the `aci_vrf_fallback_route_group` resource. Doing so will result in unexpected behaviour.


## API Information ##

* Class: [fvFBRoute](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvFBRoute/overview)

* Supported in ACI versions: 5.2(4d) and later.

* Distinguished Name Format: `uni/tn-{name}/ctx-{name}/fbrg-{name}/pfx-[{fbrPrefix}]`

## GUI Information ##

* Location: `Tenants -> Networking -> VRFs -> Policy -> Fallback Route Group -> Fallback Routes`

## Example Usage ##

The configuration snippet below creates a VRF Fallback Route with only required attributes.

```hcl

resource "aci_vrf_fallback_route" "example_vrf_fallback_route_group" {
  parent_dn      = aci_vrf_fallback_route_group.example.id
  prefix_address = "2.2.2.3/24"
}

```
The configuration snippet below shows all possible attributes of the VRF Fallback Route.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_vrf_fallback_route" "full_example_vrf_fallback_route_group" {
  parent_dn      = aci_vrf_fallback_route_group.example.id
  annotation     = "annotation"
  description    = "description_1"
  prefix_address = "2.2.2.3/24"
  name           = "name_1"
  name_alias     = "name_alias_1"
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

All examples for the VRF Fallback Route resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_vrf_fallback_route) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_vrf_fallback_route_group](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/vrf_fallback_route_group) ([fvFBRGroup](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvFBRGroup/overview))
* `prefix_address` (fbrPrefix) - (string) The prefix address of the VRF Fallback Route object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the VRF Fallback Route object.

### Optional ###

* `annotation` (annotation) - (string) The annotation of the VRF Fallback Route object.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the VRF Fallback Route object.
* `name` (name) - (string) The name of the VRF Fallback Route object.
* `name_alias` (nameAlias) - (string) The name alias of the VRF Fallback Route object.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing VRF Fallback Route can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_vrf_fallback_route.example_vrf_fallback_route_group uni/tn-{name}/ctx-{name}/fbrg-{name}/pfx-[{fbrPrefix}]
```

Starting in Terraform version 1.5, an existing VRF Fallback Route can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ctx-{name}/fbrg-{name}/pfx-[{fbrPrefix}]"
  to = aci_vrf_fallback_route.example_vrf_fallback_route_group
}
```
