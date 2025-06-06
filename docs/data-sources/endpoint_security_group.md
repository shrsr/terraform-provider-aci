---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_endpoint_security_group"
sidebar_current: "docs-aci-data-source-aci_endpoint_security_group"
description: |-
  Data source for ACI Endpoint Security Group
---

# aci_endpoint_security_group #

Data source for ACI Endpoint Security Group

## API Information ##

* Class: [fvESg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvESg/overview)

* Supported in ACI versions: 5.0(1k) and later.

* Distinguished Name Format: `uni/tn-{name}/ap-{name}/esg-{name}`

## GUI Information ##

* Location: `Tenants -> Application Profiles -> Endpoint Security Groups`

## Example Usage ##

```hcl

data "aci_endpoint_security_group" "example_application_profile" {
  parent_dn = aci_application_profile.example.id
  name      = "test_name"
}

```

## Schema ##

### Required ###

* `parent_dn` - (string) The distinguished name (DN) of the parent object, possible resources:
  - [aci_application_profile](https://registry.terraform.io/providers/CiscoDevNet/aci/latest/docs/resources/application_profile) ([fvAp](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvAp/overview))
* `name` (name) - (string) The name of the Endpoint Security Group object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Endpoint Security Group object.
* `annotation` (annotation) - (string) The annotation of the Endpoint Security Group object.
* `description` (descr) - (string) The description of the Endpoint Security Group object.
* `exception_tag` (exceptionTag) - (string) Contract Exception Tag.
* `match_criteria` (matchT) - (string) The provider label match criteria.
* `name_alias` (nameAlias) - (string) The name alias of the Endpoint Security Group object.
* `intra_esg_isolation` (pcEnfPref) - (string) Parameter used to determine whether communication between endpoints within the ESG is blocked.
* `pc_tag` (pcTag) - (string) The classification tag used for policy enforcement and zoning.
* `preferred_group_member` (prefGrMemb) - (string) Parameter used to determine whether the ESG is part of the preferred group. Members of this group are allowed to communicate without contracts.
* `scope` (scope) - (string) The scope ID (L3-VNI) of the Endpoint Security Group object.
* `admin_state` (shutdown) - (string) Withdraw the ESG configuration from all nodes in the fabric. This attribute is supported in ACI versions: 5.2(1g) and later.
* `relation_to_consumed_contracts` - (list) A list of Relation To Consumed Contracts (ACI object [fvRsCons](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsCons/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)). This attribute is supported in ACI versions: 1.0(1e) and later.
    * `annotation` (annotation) - (string) The annotation of the Relation To Consumed Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
    * `contract_name` (tnVzBrCPName) - (string) The consumer contract name.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_imported_contracts` - (list) A list of Relation To Imported Contracts (ACI object [fvRsConsIf](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsConsIf/overview)) pointing to Imported Contract (ACI Object [vzCPIf](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzCPIf/overview)). This attribute is supported in ACI versions: 1.0(1e) and later.
    * `annotation` (annotation) - (string) The annotation of the Relation To Imported Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
    * `imported_contract_name` (tnVzCPIfName) - (string) The contract interface name.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_intra_epg_contracts` - (list) A list of Relation To Intra EPG Contracts (ACI object [fvRsIntraEpg](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsIntraEpg/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)). This attribute is supported in ACI versions: 3.0(1k) and later.
    * `annotation` (annotation) - (string) The annotation of the Relation To Intra EPG Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `contract_name` (tnVzBrCPName) - (string) The contract name.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_provided_contracts` - (list) A list of Relation To Provided Contracts (ACI object [fvRsProv](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsProv/overview)) pointing to Contract (ACI Object [vzBrCP](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/vzBrCP/overview)). This attribute is supported in ACI versions: 1.0(1e) and later.
    * `annotation` (annotation) - (string) The annotation of the Relation To Provided Contract object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `match_criteria` (matchT) - (string) The provider label match criteria.
    * `priority` (prio) - (string) The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.
    * `contract_name` (tnVzBrCPName) - (string) The provider contract name.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_vrf` - (map) A map of Relation To VRF (ACI object [fvRsScope](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsScope/overview)) pointing to VRF (ACI Object [fvCtx](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvCtx/overview)).
    * `annotation` (annotation) - (string) The annotation of the Relation To VRF object.
    * `vrf_name` (tnFvCtxName) - (string) The name of the VRF object.
    * `annotations` - (list) A list of Annotations (ACI object [tagAnnotation](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagAnnotation/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
    * `tags` - (list) A list of Tags (ACI object [tagTag](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/tagTag/overview)). This attribute is supported in ACI versions: 3.2(1l) and later.
        * `key` (key) - (string) The key used to uniquely identify this configuration object.
        * `value` (value) - (string) The value of the property.
* `relation_to_contract_masters` - (list) A list of Relation To Contract Masters (ACI object [fvRsSecInherited](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsSecInherited/overview)) pointing to . This attribute is supported in ACI versions: 2.3(1e) and later.
    * `annotation` (annotation) - (string) The annotation of the Relation To Contract Master object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `target_dn` (tDn) - (string) The distinguished name of the target.
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
