package aci

import (
	"fmt"
	"log"

	"github.com/ciscoecosystem/aci-go-client/client"
	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceAciEndpointSecurityGroupSelector() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciEndpointSecurityGroupSelectorCreate,
		Update: resourceAciEndpointSecurityGroupSelectorUpdate,
		Read:   resourceAciEndpointSecurityGroupSelectorRead,
		Delete: resourceAciEndpointSecurityGroupSelectorDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAciEndpointSecurityGroupSelectorImport,
		},

		SchemaVersion: 1,
		Schema: AppendBaseAttrSchema(AppendNameAliasAttrSchema(map[string]*schema.Schema{
			"endpoint_security_group_dn": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"match_expression": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		})),
	}
}

func getRemoteEndpointSecurityGroupSelector(client *client.Client, dn string) (*models.EndpointSecurityGroupSelector, error) {
	fvEPSelectorCont, err := client.Get(dn)
	if err != nil {
		return nil, err
	}
	fvEPSelector := models.EndpointSecurityGroupSelectorFromContainer(fvEPSelectorCont)
	if fvEPSelector.DistinguishedName == "" {
		return nil, fmt.Errorf("EndpointSecurityGroupSelector %s not found", fvEPSelector.DistinguishedName)
	}
	return fvEPSelector, nil
}

func setEndpointSecurityGroupSelectorAttributes(fvEPSelector *models.EndpointSecurityGroupSelector, d *schema.ResourceData) *schema.ResourceData {
	d.SetId(fvEPSelector.DistinguishedName)
	d.Set("description", fvEPSelector.Description)
	fvEPSelectorMap, _ := fvEPSelector.ToMap()
	d.Set("endpoint_security_group_dn", GetParentDn(fvEPSelector.DistinguishedName, fmt.Sprintf("/epselector-[%s]")))
	d.Set("annotation", fvEPSelectorMap["annotation"])
	d.Set("match_expression", fvEPSelectorMap["matchExpression"])
	d.Set("name", fvEPSelectorMap["name"])
	d.Set("name_alias", fvEPSelectorMap["nameAlias"])
	return d
}

func resourceAciEndpointSecurityGroupSelectorImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] %s: Beginning Import", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	fvEPSelector, err := getRemoteEndpointSecurityGroupSelector(aciClient, dn)
	if err != nil {
		return nil, err
	}
	schemaFilled := setEndpointSecurityGroupSelectorAttributes(fvEPSelector, d)
	log.Printf("[DEBUG] %s: Import finished successfully", d.Id())
	return []*schema.ResourceData{schemaFilled}, nil
}

func resourceAciEndpointSecurityGroupSelectorCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] EndpointSecurityGroupSelector: Beginning Creation")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	matchExpression := d.Get("match_expression").(string)
	EndpointSecurityGroupDn := d.Get("endpoint_security_group_dn").(string)

	fvEPSelectorAttr := models.EndpointSecurityGroupSelectorAttributes{}
	nameAlias := ""
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		nameAlias = NameAlias.(string)
	}
	if Annotation, ok := d.GetOk("annotation"); ok {
		fvEPSelectorAttr.Annotation = Annotation.(string)
	} else {
		fvEPSelectorAttr.Annotation = "{}"
	}

	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		fvEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}

	if Name, ok := d.GetOk("name"); ok {
		fvEPSelectorAttr.Name = Name.(string)
	}
	fvEPSelector := models.NewEndpointSecurityGroupSelector(fmt.Sprintf("epselector-[%s]", matchExpression), EndpointSecurityGroupDn, desc, nameAlias, fvEPSelectorAttr)

	err := aciClient.Save(fvEPSelector)
	if err != nil {
		return err
	}
	d.Partial(true)
	d.SetPartial("name")
	d.Partial(false)

	d.SetId(fvEPSelector.DistinguishedName)
	log.Printf("[DEBUG] %s: Creation finished successfully", d.Id())
	return resourceAciEndpointSecurityGroupSelectorRead(d, m)
}

func resourceAciEndpointSecurityGroupSelectorUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] EndpointSecurityGroupSelector: Beginning Update")
	aciClient := m.(*client.Client)
	desc := d.Get("description").(string)
	matchExpression := d.Get("match_expression").(string)
	EndpointSecurityGroupDn := d.Get("endpoint_security_group_dn").(string)
	fvEPSelectorAttr := models.EndpointSecurityGroupSelectorAttributes{}
	nameAlias := ""
	if NameAlias, ok := d.GetOk("name_alias"); ok {
		nameAlias = NameAlias.(string)
	}

	if Annotation, ok := d.GetOk("annotation"); ok {
		fvEPSelectorAttr.Annotation = Annotation.(string)
	} else {
		fvEPSelectorAttr.Annotation = "{}"
	}

	if MatchExpression, ok := d.GetOk("match_expression"); ok {
		fvEPSelectorAttr.MatchExpression = MatchExpression.(string)
	}

	if Name, ok := d.GetOk("name"); ok {
		fvEPSelectorAttr.Name = Name.(string)
	}
	fvEPSelector := models.NewEndpointSecurityGroupSelector(fmt.Sprintf("epselector-[%s]", matchExpression), EndpointSecurityGroupDn, desc, nameAlias, fvEPSelectorAttr)

	fvEPSelector.Status = "modified"
	err := aciClient.Save(fvEPSelector)
	if err != nil {
		return err
	}
	d.Partial(true)
	d.SetPartial("name")
	d.Partial(false)

	d.SetId(fvEPSelector.DistinguishedName)
	log.Printf("[DEBUG] %s: Update finished successfully", d.Id())
	return resourceAciEndpointSecurityGroupSelectorRead(d, m)
}

func resourceAciEndpointSecurityGroupSelectorRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Read", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	fvEPSelector, err := getRemoteEndpointSecurityGroupSelector(aciClient, dn)
	if err != nil {
		d.SetId("")
		return err
	}
	setEndpointSecurityGroupSelectorAttributes(fvEPSelector, d)

	log.Printf("[DEBUG] %s: Read finished successfully", d.Id())
	return nil
}

func resourceAciEndpointSecurityGroupSelectorDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] %s: Beginning Destroy", d.Id())
	aciClient := m.(*client.Client)
	dn := d.Id()
	err := aciClient.DeleteByDn(dn, "fvEPSelector")
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %s: Destroy finished successfully", d.Id())
	d.SetId("")
	return err
}
