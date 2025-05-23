---
# Documentation generated by "gen/generator.go"; DO NOT EDIT.
# In order to regenerate this file execute `go generate` from the repository root.
# More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).
subcategory: "Application Management"
layout: "aci"
page_title: "ACI: aci_tenant"
sidebar_current: "docs-aci-data-source-aci_tenant"
description: |-
  Data source for ACI Tenant
---

# aci_tenant #

Data source for ACI Tenant

## API Information ##

* Class: [fvTenant](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvTenant/overview)

* Supported in ACI versions: 1.0(1e) and later.

* Distinguished Name Format: `uni/tn-{name}`

## GUI Information ##

* Location: `Tenants`

## Example Usage ##

```hcl

data "aci_tenant" "example" {
  name = "test_name"
}

```

## Schema ##

### Required ###

* `name` (name) - (string) The name of the Tenant object.

### Read-Only ###

* `id` - (string) The distinguished name (DN) of the Tenant object.
* `annotation` (annotation) - (string) The annotation of the Tenant object. This attribute is supported in ACI versions: 3.2(1l) and later.
* `description` (descr) - (string) The description of the Tenant object.
* `name_alias` (nameAlias) - (string) The name alias of the Tenant object. This attribute is supported in ACI versions: 2.2(1k) and later.
* `owner_key` (ownerKey) - (string) The key for enabling clients to own their data for entity correlation.
* `owner_tag` (ownerTag) - (string) A tag for enabling clients to add their own data. For example, to indicate who created this object.
* `relation_to_monitoring_policy` - (map) A map of Relation From Tenant To Monitoring Policy (ACI object [fvRsTenantMonPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/fvRsTenantMonPol/overview)) pointing to Monitoring Policy (ACI Object [monEPGPol](https://pubhub.devnetcloud.com/media/model-doc-latest/docs/app/index.html#/objects/monEPGPol/overview)).
    * `annotation` (annotation) - (string) The annotation of the Relation From Tenant To Monitoring Policy object. This attribute is supported in ACI versions: 3.2(1l) and later.
    * `monitoring_policy_name` (tnMonEPGPolName) - (string) The name of the monitoring policy.
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
