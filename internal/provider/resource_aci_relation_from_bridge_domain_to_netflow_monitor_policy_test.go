// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccResourceFvRsBDToNetflowMonitorPolWithFvBD(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBDAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test_2", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test", "filter_type", "ipv4"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test_2", "filter_type", "ipv4"),
				),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "false")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:      testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBDAllowExisting,
				ExpectError: regexp.MustCompile("Object Already Exists"),
			},
		},
	})

	setEnvVariable(t, "ACI_ALLOW_EXISTING_ON_CREATE", "true")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBDAllowExisting,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test_2", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test_2", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test", "filter_type", "ipv4"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test_2", "filter_type", "ipv4"),
				),
			},
		},
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t, "both", "2.2(1k)-") },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create with minimum config and verify default APIC values
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "filter_type", "ipv4"),
				),
			},
			// Update with all config and verify default APIC values
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolAllDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotation", "annotation"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "filter_type", "ce"),
				),
			},
			// Update with minimum config and verify config is unchanged
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "filter_type", "ipv4"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
				),
			},
			// Update with empty strings config or default value
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolResetDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "filter_type", "ipv4"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
				),
			},
			// Import testing
			{
				ResourceName:      "aci_relation_from_bridge_domain_to_netflow_monitor_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolChildrenDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "filter_type", "ipv4"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "netflow_monitor_policy_name", "test_tn_netflow_monitor_pol_name"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotation", "orchestrator:terraform"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.1.value", "test_value"),
				),
			},
			// Refresh State before import testing to ensure that the state is up to date
			{
				RefreshState:       true,
				ExpectNonEmptyPlan: false,
			},
			// Import testing with children
			{
				ResourceName:      "aci_relation_from_bridge_domain_to_netflow_monitor_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update with children removed from config
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolChildrenRemoveFromConfigDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.#", "2"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.0.key", "key_0"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.0.value", "value_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.1.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.1.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.#", "2"),
				),
			},
			// Update with children first child removed
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolChildrenRemoveOneDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.#", "1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.0.key", "key_1"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.0.value", "test_value"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.#", "1"),
				),
			},
			// Update with all children removed
			{
				Config:             testConfigFvRsBDToNetflowMonitorPolChildrenRemoveAllDependencyWithFvBD,
				ExpectNonEmptyPlan: false,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "annotations.#", "0"),
					resource.TestCheckResourceAttr("aci_relation_from_bridge_domain_to_netflow_monitor_policy.test", "tags.#", "0"),
				),
			},
		},
		CheckDestroy: testCheckResourceDestroy,
	})
}

const testDependencyConfigFvRsBDToNetflowMonitorPol = `
resource "aci_netflow_monitor_policy" "test_netflow_monitor_policy_0" {
  parent_dn = aci_tenant.test.id
  name = "netflow_monitor_policy_name_1"
}
`

const testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBDAllowExisting = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "allow_test" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
}
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "allow_test_2" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
  depends_on = [aci_relation_from_bridge_domain_to_netflow_monitor_policy.allow_test]
}
`

const testConfigFvRsBDToNetflowMonitorPolMinDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
}
`

const testConfigFvRsBDToNetflowMonitorPolAllDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
  annotation = "annotation"
  filter_type = "ce"
}
`

const testConfigFvRsBDToNetflowMonitorPolResetDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
  annotation = "orchestrator:terraform"
  filter_type = "ipv4"
}
`
const testConfigFvRsBDToNetflowMonitorPolChildrenDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
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

const testConfigFvRsBDToNetflowMonitorPolChildrenRemoveFromConfigDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
}
`

const testConfigFvRsBDToNetflowMonitorPolChildrenRemoveOneDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
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

const testConfigFvRsBDToNetflowMonitorPolChildrenRemoveAllDependencyWithFvBD = testDependencyConfigFvRsBDToNetflowMonitorPol + testConfigFvBDMinDependencyWithFvTenant + `
resource "aci_relation_from_bridge_domain_to_netflow_monitor_policy" "test" {
  parent_dn = aci_bridge_domain.test.id
  filter_type = "ipv4"
  netflow_monitor_policy_name = "test_tn_netflow_monitor_pol_name"
  annotations = []
  tags = []
}
`
