---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Node Management"
layout: "aci"
page_title: "ACI: aci_external_management_network_instance_profile"
sidebar_current: "docs-aci-data-source-aci_external_management_network_instance_profile"
description: |-
  Data source for ACI External Management Network Instance Profile
---

# aci_external_management_network_instance_profile #

Data source for ACI External Management Network Instance Profile

## API Information ##

* Class: [mgmtInstP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtInstP/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-mgmt/extmgmt-default/instp-{name}`

## GUI Information ##

* Location: `Tenants (mgmt) -> External Management Network Instance Profiles`

## Example Usage ##

```hcl

data "aci_external_management_network_instance_profile" "example" {
  name = "test_name"
}

```

## Schema ##

### Required ###

* `name` (name) - (string) The name of the External Management Network Instance Profile object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the External Management Network Instance Profile object.
* `annotation` (annotation) - (string) The annotation of the External Management Network Instance Profile object. This attribute is supported in ACI versions: 3.2(1l) and later.
* `description` (descr) - (string) The description of the External Management Network Instance Profile object.
* `name_alias` (nameAlias) - (string) The name alias of the External Management Network Instance Profile object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
* `relation_to_consumed_out_of_band_contracts` - (list) A list of Relation To Consumed Out Of Band Contracts (ACI object [mgmtRsOoBCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/mgmtRsOoBCons/overview)) pointing to Out Of Band Contract (ACI Object [vzOOBBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzOOBBrCP/overview)).
    * `annotation` (annotation) - (string) The annotation of the Relation To Consumed Out Of Band Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
    * `out_of_band_contract_name` (tnVzOOBBrCPName) - (string) The name of the Out Of Band Contract object.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
* `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
    * `key` (key) - (string) The key used to uniquely identify this configuration object.
    * `value` (value) - (string) The value of the property.
