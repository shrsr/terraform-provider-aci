---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_endpoint_security_group"
sidebar_current: "docs-aci-resource-aci_endpoint_security_group"
description: |-
  Manages ACI Endpoint Security Group
---

# aci_endpoint_security_group #

Manages ACI Endpoint Security Group


  !> This resource has been migrated to the terraform plugin protocol version 6, refer to the [migration guide](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/guides/migration) for more details and implications for already managed resources.

## API Information ##

* Class: [fvESg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvESg/overview)

* Supported in ACI versions: 5.0(1k) and later.

* Distinguished Name Format: `uni/tn-{name}/ap-{name}/esg-{name}`

## GUI Information ##

* Location: `Tenants -> Application Profiles -> Endpoint Security Groups`

## Example Usage ##

The configuration snippet below creates a Endpoint Security Group with only required attributes.

```hcl

resource "aci_endpoint_security_group" "example_application_profile" {
  parent_dn = aci_application_profile.example.id
  name      = "test_name"
}

```
The configuration snippet below shows all possible attributes of the Endpoint Security Group.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_endpoint_security_group" "full_example_application_profile" {
  parent_dn              = aci_application_profile.example.id
  annotation             = "annotation"
  description            = "description_1"
  exception_tag          = "exception_tag_1"
  match_criteria         = "All"
  name                   = "test_name"
  name_alias             = "name_alias_1"
  intra_esg_isolation    = "enforced"
  preferred_group_member = "exclude"
  admin_state            = "no"
  relation_to_consumed_contracts = [
    {
      annotation    = "annotation_1"
      priority      = "level1"
      contract_name = aci_contract.example.name
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
  ]
  relation_to_imported_contracts = [
    {
      annotation             = "annotation_1"
      priority               = "level1"
      imported_contract_name = aci_imported_contract.example.name
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
  ]
  relation_to_intra_epg_contracts = [
    {
      annotation    = "annotation_1"
      contract_name = aci_contract.example.name
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
  ]
  relation_to_provided_contracts = [
    {
      annotation     = "annotation_1"
      match_criteria = "All"
      priority       = "level1"
      contract_name  = aci_contract.example.name
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
  ]
  relation_to_vrf = {
    annotation = "annotation_1"
    vrf_name   = aci_vrf.example.name
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
  relation_to_contract_masters = [
    {
      annotation = "annotation_1"
      target_dn  = aci_endpoint_security_group.test_endpoint_security_group_0.id
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
  ]
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

All examples for the Endpoint Security Group resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_endpoint_security_group) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_application_profile](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_profile) ([fvAp](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAp/overview))
* `name` (name) - (string) The name of the Endpoint Security Group object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Endpoint Security Group object.
* `pc_tag` (pcTag) - (string) The classification tag used for policy enforcement and zoning.
* `scope` (scope) - (string) The scope ID (L3-VNI) of the Endpoint Security Group object.
  - Default: `0`

### Optional ###

* `annotation` (annotation) - (string) The annotation of the Endpoint Security Group object.
  - Default: `orchestrator:terraform`
* `description` (descr) - (string) The description of the Endpoint Security Group object.
* `exception_tag` (exceptionTag) - (string) Contract Exception Tag.
* `match_criteria` (matchT) - (string) The provider label match criteria.
  - Default: `AtleastOne`
  - Valid Values: `All`, `AtleastOne`, `AtmostOne`, `None`.
* `name_alias` (nameAlias) - (string) The name alias of the Endpoint Security Group object.
* `intra_esg_isolation` (pcEnfPref) - (string) Parameter used to determine whether communication between endpoints within the ESG is blocked.
  - Default: `unenforced`
  - Valid Values: `enforced`, `unenforced`.
* `preferred_group_member` (prefGrMemb) - (string) Parameter used to determine whether the ESG is part of the preferred group. Members of this group are allowed to communicate without contracts.
  - Default: `exclude`
  - Valid Values: `exclude`, `include`.
* `admin_state` (shutdown) - (string) Withdraw the ESG configuration from all nodes in the fabric. This attribute is supported in ACI versions: 5.2(1g) and later.
  - Default: `no`
  - Valid Values: `no`, `yes`.
* `relation_to_consumed_contracts` - (list) A list of Relation To Consumed Contracts (ACI object [fvRsCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCons/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)) which can be configured using the [aci_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) resource. This attribute is supported in ACI versions: 1.0(1e) and later.
  #### Required ####
  
    * `contract_name` (tnVzBrCPName) - (string) The consumer contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.
  #### Optional ####
    
    * `annotation` (annotation) - (string) The annotation of the Relation To Consumed Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
      - Default: `orchestrator:terraform`
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
      - Default: `unspecified`
      - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_imported_contracts` - (list) A list of Relation To Imported Contracts (ACI object [fvRsConsIf](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsConsIf/overview)) pointing to Imported Contract (ACI Object [vzCPIf](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzCPIf/overview)) which can be configured using the [aci_imported_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/imported_contract) resource. This attribute is supported in ACI versions: 1.0(1e) and later.
  #### Required ####
  
    * `imported_contract_name` (tnVzCPIfName) - (string) The contract interface name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/imported_contract) with `aci_imported_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/imported_contract) with `data.aci_imported_contract.example.name`.
  #### Optional ####
    
    * `annotation` (annotation) - (string) The annotation of the Relation To Imported Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
      - Default: `orchestrator:terraform`
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
      - Default: `unspecified`
      - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_intra_epg_contracts` - (list) A list of Relation To Intra EPG Contracts (ACI object [fvRsIntraEpg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsIntraEpg/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)) which can be configured using the [aci_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) resource. This attribute is supported in ACI versions: 3.0(1k) and later.
  #### Required ####
  
    * `contract_name` (tnVzBrCPName) - (string) The contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.
  #### Optional ####
    
    * `annotation` (annotation) - (string) The annotation of the Relation To Intra EPG Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
      - Default: `orchestrator:terraform`
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_provided_contracts` - (list) A list of Relation To Provided Contracts (ACI object [fvRsProv](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsProv/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)) which can be configured using the [aci_contract](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) resource. This attribute is supported in ACI versions: 1.0(1e) and later.
  #### Required ####
  
    * `contract_name` (tnVzBrCPName) - (string) The provider contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.
  #### Optional ####
    
    * `annotation` (annotation) - (string) The annotation of the Relation To Provided Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
      - Default: `orchestrator:terraform`
    * `match_criteria` (matchT) - (string) The provider label match criteria.
      - Default: `AtleastOne`
      - Valid Values: `All`, `AtleastOne`, `AtmostOne`, `None`.
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
      - Default: `unspecified`
      - Valid Values: `level1`, `level2`, `level3`, `level4`, `level5`, `level6`, `unspecified`.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_vrf` - (map) A map of Relation To VRF (ACI object [fvRsScope](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsScope/overview)) pointing to VRF (ACI Object [fvCtx](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCtx/overview)) which can be configured using the [aci_vrf](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/vrf) resource.
  #### Optional ####
    
    * `annotation` (annotation) - (string) The annotation of the Relation To VRF object.
      - Default: `orchestrator:terraform`
    * `vrf_name` (tnFvCtxName) - (string) The name of the VRF object.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_contract_masters` - (list) A list of Relation To Contract Masters (ACI object [fvRsSecInherited](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsSecInherited/overview)) pointing to . This attribute is supported in ACI versions: 2.3(1e) and later.
  #### Required ####
  
    * `target_dn` (tDn) - (string) The distinguished name of the target.
  #### Optional ####
    
    * `annotation` (annotation) - (string) The annotation of the Relation To Contract Master object. This attribute is supported in ACI versions: 3.2(1l) and later.
      - Default: `orchestrator:terraform`
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
      #### Required ####
  
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). Annotations can also be configured using a separate [aci_annotation](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/annotation) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). Tags can also be configured using a separate [aci_tag](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/tag) resource. This attribute is supported in ACI versions: 3.2(1l) and later.
  #### Required ####
  
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.

## Importing

An existing Endpoint Security Group can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_endpoint_security_group.example_application_profile uni/tn-{name}/ap-{name}/esg-{name}
```

Starting in Terraform version 1.5, an existing Endpoint Security Group can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ap-{name}/esg-{name}"
  to = aci_endpoint_security_group.example_application_profile
}
```
