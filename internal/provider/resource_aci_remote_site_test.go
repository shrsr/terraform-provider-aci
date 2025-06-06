// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRemoteIdWithFvSiteAssociated(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "site_id", "0"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "site_id", "0"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test", "site_id", "0"),
					resource.TestCheckResourceAttr("aci_remote_site.allow_test_2", "site_id", "0"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "3.0(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "site_id", "0"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRemoteIdAllDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "description", "description_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name", "name_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name_alias", "name_alias_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_key", "owner_key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_tag", "owner_tag_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_pc_tag", "16386"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "site_id", "1"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRemoteIdMinDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "site_id", "0"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRemoteIdResetDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "site_id", "0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_vrf_pc_tag", "any"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_remote_site.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRemoteIdChildrenDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_pc_tag", "16387"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "site_id", "0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "description", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "name_alias", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_key", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "owner_tag", ""),
					resource.TestCheckResourceAttr("aci_remote_site.test", "remote_vrf_pc_tag", "any"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_remote_site.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRemoteIdChildrenRemoveFromConfigDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRemoteIdChildrenRemoveOneDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRemoteIdChildrenRemoveAllDependencyWithFvSiteAssociated,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_remote_site.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_remote_site.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testConfigFvRemoteIdMinDependencyWithFvSiteAssociatedAllowExisting = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "allow_test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
}
resource "aci_remote_site" "allow_test_2" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
  depends_on = [aci_remote_site.allow_test]
}
`

const testConfigFvRemoteIdMinDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
}
`

const testConfigFvRemoteIdAllDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  annotation = "annotation"
  description = "description_1"
  name = "name_1"
  name_alias = "name_alias_1"
  owner_key = "owner_key_1"
  owner_tag = "owner_tag_1"
  remote_pc_tag = "16386"
  remote_vrf_pc_tag = "any"
  site_id = "1"
}
`

const testConfigFvRemoteIdResetDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  annotation = "orchestrator:terraform"
  description = ""
  name = ""
  name_alias = ""
  owner_key = ""
  owner_tag = ""
  remote_pc_tag = "16387"
  remote_vrf_pc_tag = "any"
  site_id = "0"
}
`
const testConfigFvRemoteIdChildrenDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
  annotations = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
  tags = [
    {
      key = "key_0"
      value = "value_1"
    },
    {
      key = "key_1"
      value = "test_value"
    },
  ]
}
`

const testConfigFvRemoteIdChildrenRemoveFromConfigDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
}
`

const testConfigFvRemoteIdChildrenRemoveOneDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
  annotations = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
  tags = [ 
	{
	  key = "key_1"
	  value = "test_value"
	},
  ]
}
`

const testConfigFvRemoteIdChildrenRemoveAllDependencyWithFvSiteAssociated = testConfigFvSiteAssociatedMinDependencyWithFvBD + `
resource "aci_remote_site" "test" {
  parent_dn = aci_associated_site.test.id
  remote_pc_tag = "16387"
  site_id = "0"
  annotations = []
  tags = []
}
`
