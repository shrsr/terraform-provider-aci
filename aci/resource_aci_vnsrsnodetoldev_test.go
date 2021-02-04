package aci

import (
	"fmt"
	"testing"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAciRelationfromaAbsNodetoanLDev_Basic(t *testing.T) {
	var relationfroma_abs_nodetoan_l_dev models.RelationfromaAbsNodetoanLDev

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciRelationfromaAbsNodetoanLDevDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciRelationfromaAbsNodetoanLDevConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciRelationfromaAbsNodetoanLDevExists("aci_relationfroma_abs_nodetoan_l_dev.foorelationfroma_abs_nodetoan_l_dev", &relationfroma_abs_nodetoan_l_dev),
					testAccCheckAciRelationfromaAbsNodetoanLDevAttributes(&relationfroma_abs_nodetoan_l_dev),
				),
			},
		},
	})
}

func TestAccAciRelationfromaAbsNodetoanLDev_update(t *testing.T) {
	var relationfroma_abs_nodetoan_l_dev models.RelationfromaAbsNodetoanLDev

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciRelationfromaAbsNodetoanLDevDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciRelationfromaAbsNodetoanLDevConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciRelationfromaAbsNodetoanLDevExists("aci_relationfroma_abs_nodetoan_l_dev.foorelationfroma_abs_nodetoan_l_dev", &relationfroma_abs_nodetoan_l_dev),
					testAccCheckAciRelationfromaAbsNodetoanLDevAttributes(&relationfroma_abs_nodetoan_l_dev),
				),
			},
			{
				Config: testAccCheckAciRelationfromaAbsNodetoanLDevConfig_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciRelationfromaAbsNodetoanLDevExists("aci_relationfroma_abs_nodetoan_l_dev.foorelationfroma_abs_nodetoan_l_dev", &relationfroma_abs_nodetoan_l_dev),
					testAccCheckAciRelationfromaAbsNodetoanLDevAttributes(&relationfroma_abs_nodetoan_l_dev),
				),
			},
		},
	})
}

func testAccCheckAciRelationfromaAbsNodetoanLDevConfig_basic() string {
	return fmt.Sprintf(`

	resource "aci_relationfroma_abs_nodetoan_l_dev" "foorelationfroma_abs_nodetoan_l_dev" {
		  function_node_dn  = "uni/tn-check_tenantnk/AbsGraph-second/AbsNode-N1"
		  annotation  = "example"
		  t_dn  = "uni/tn-check_tenantnk/lDevVip-second"
		}
	`)
}

func testAccCheckAciRelationfromaAbsNodetoanLDevExists(name string, relationfroma_abs_nodetoan_l_dev *models.RelationfromaAbsNodetoanLDev) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Relation from a AbsNode to an LDev %s not found", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Relation from a AbsNode to an LDev dn was set")
		}

		client := testAccProvider.Meta().(*client.Client)

		cont, err := client.Get(rs.Primary.ID)
		if err != nil {
			return err
		}

		relationfroma_abs_nodetoan_l_devFound := models.RelationfromaAbsNodetoanLDevFromContainer(cont)
		if relationfroma_abs_nodetoan_l_devFound.DistinguishedName != rs.Primary.ID {
			return fmt.Errorf("Relation from a AbsNode to an LDev %s not found", rs.Primary.ID)
		}
		*relationfroma_abs_nodetoan_l_dev = *relationfroma_abs_nodetoan_l_devFound
		return nil
	}
}

func testAccCheckAciRelationfromaAbsNodetoanLDevDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*client.Client)

	for _, rs := range s.RootModule().Resources {

		if rs.Type == "aci_relationfroma_abs_nodetoan_l_dev" {
			cont, err := client.Get(rs.Primary.ID)
			relationfroma_abs_nodetoan_l_dev := models.RelationfromaAbsNodetoanLDevFromContainer(cont)
			if err == nil {
				return fmt.Errorf("Relation from a AbsNode to an LDev %s Still exists", relationfroma_abs_nodetoan_l_dev.DistinguishedName)
			}

		} else {
			continue
		}
	}

	return nil
}

func testAccCheckAciRelationfromaAbsNodetoanLDevAttributes(relationfroma_abs_nodetoan_l_dev *models.RelationfromaAbsNodetoanLDev) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if "example" != relationfroma_abs_nodetoan_l_dev.Annotation {
			return fmt.Errorf("Bad relationfroma_abs_nodetoan_l_dev annotation %s", relationfroma_abs_nodetoan_l_dev.Annotation)
		}

		if "uni/tn-check_tenantnk/lDevVip-second" != relationfroma_abs_nodetoan_l_dev.TDn {
			return fmt.Errorf("Bad relationfroma_abs_nodetoan_l_dev t_dn %s", relationfroma_abs_nodetoan_l_dev.TDn)
		}

		return nil
	}
}
