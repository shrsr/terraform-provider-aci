---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_relation_to_consumed_contract"
sidebar_current: "docs-aci-resource-aci_relation_to_consumed_contract"
description: |-
  Manages ACI Relation To Consumed Contract
---

# aci_relation_to_consumed_contract #

Manages ACI Relation To Consumed Contract



## API Information ##

* Class: [fvRsCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCons/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Formats:
  - Too many DN formats to display, see model documentation for all possible parents of [fvRsCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCons/overview).
  - `uni/tn-{name}/ap-{name}/epg-{name}/rscons-{tnVzBrCPName}`
  - `uni/tn-{name}/ap-{name}/esg-{name}/rscons-{tnVzBrCPName}`
  - `uni/tn-{name}/l2out-{name}/instP-{name}/rscons-{tnVzBrCPName}`
  - `uni/tn-{name}/out-{name}/instP-{name}/rscons-{tnVzBrCPName}`

## GUI Information ##

* Locations:
  - `Tenants -> Application Profiles -> Application EPGs -> Contracts`
  - `Tenants -> Application Profiles -> Endpoint Security Groups -> Contracts`
  - `Tenants -> Networking -> L3Outs -> External EPGs -> Contracts`
  - `Tenants -> Networking -> L2Outs -> External EPGs -> Contracts`

## Example Usage ##

The configuration snippet below creates a Relation To Consumed Contract with only required attributes.

```hcl

resource "aci_relation_to_consumed_contract" "example_application_epg" {
  parent_dn     = aci_application_epg.example.id
  contract_name = aci_contract.example.name
}

resource "aci_relation_to_consumed_contract" "example_endpoint_security_group" {
  parent_dn     = aci_endpoint_security_group.example.id
  contract_name = aci_contract.example.name
}

```
The configuration snippet below shows all possible attributes of the Relation To Consumed Contract.

!> This example might not be valid configuration and is only used to show all possible attributes.

```hcl

resource "aci_relation_to_consumed_contract" "full_example_application_epg" {
  parent_dn     = aci_application_epg.example.id
  annotation    = "annotation"
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

resource "aci_relation_to_consumed_contract" "full_example_endpoint_security_group" {
  parent_dn     = aci_endpoint_security_group.example.id
  annotation    = "annotation"
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

```

All examples for the Relation To Consumed Contract resource can be found in the [examples](https://github.com/CiscoDevNet/terraform-provider-aci/tree/master/examples/resources/aci_relation_to_consumed_contract) folder.

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_cloud_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/cloud_epg) ([cloudEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/cloudEPg/overview))
  - [aci_cloud_external_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/cloud_external_epg) ([cloudExtEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/cloudExtEPg/overview))
  - [aci_cloud_service_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/cloud_service_epg) ([cloudSvcEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/cloudSvcEPg/overview))
  - [aci_application_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_epg) ([fvAEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAEPg/overview))
  - [aci_endpoint_security_group](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/endpoint_security_group) ([fvESg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvESg/overview))
  - [aci_l2out_extepg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/l2out_extepg) ([l2extInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l2extInstP/overview))
  - [aci_external_network_instance_profile](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/external_network_instance_profile) ([l3extInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l3extInstP/overview))
  - [aci_node_mgmt_epg](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/node_mgmt_epg) ([mgmtInB](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtInB/overview))
  - The distinguished name (DN) of classes below can be used but currently there is no available resource for it:
    - [cloudISvcEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/cloudISvcEPg/overview)
    - [dhcpCRelPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/dhcpCRelPg/overview)
    - [dhcpPRelPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/dhcpPRelPg/overview)
    - [fvIntEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvIntEPg/overview)
    - [fvTnlEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTnlEPg/overview)
    - [infraCEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraCEPg/overview)
    - [infraPEPg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/infraPEPg/overview)
    - [l3extInstPDef](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/l3extInstPDef/overview)
    - [vnsEPpInfo](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vnsEPpInfo/overview)
    - [vnsREPpInfo](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vnsREPpInfo/overview)
    - [vnsSDEPpInfo](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vnsSDEPpInfo/overview)
    - [vnsSHEPpInfo](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vnsSHEPpInfo/overview)

* `contract_name` (tnVzBrCPName) - (string) The consumer contract name. This attribute can be referenced from a [resource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/contract) with `aci_contract.example.name` or from a [datasource](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/data-sources/contract) with `data.aci_contract.example.name`.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Relation To Consumed Contract object.

### Optional ###

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

## Importing

An existing Relation To Consumed Contract can be [imported](https://www.terraform.io/docs/import/index.html) into this resource with its distinguished name (DN), via the following command:

```
terraform import aci_relation_to_consumed_contract.example_application_epg uni/tn-{name}/ap-{name}/epg-{name}/rscons-{tnVzBrCPName}
```

Starting in Terraform version 1.5, an existing Relation To Consumed Contract can be imported
using [import blocks](https://developer.hashicorp.com/terraform/language/import) via the following configuration:

```
import {
  id = "uni/tn-{name}/ap-{name}/epg-{name}/rscons-{tnVzBrCPName}"
  to = aci_relation_to_consumed_contract.example_application_epg
}
```
