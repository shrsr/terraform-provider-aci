// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &InfraHPathSResource{}
var _ resource.ResourceWithImportState = &InfraHPathSResource{}

func NewInfraHPathSResource() resource.Resource {
	return &InfraHPathSResource{}
}

// InfraHPathSResource defines the resource implementation.
type InfraHPathSResource struct {
	client *client.Client
}

// InfraHPathSResourceModel describes the resource data model.
type InfraHPathSResourceModel struct {
	Id                      types.String `tfsdk:"id"`
	ParentDn                types.String `tfsdk:"parent_dn"`
	Annotation              types.String `tfsdk:"annotation"`
	Descr                   types.String `tfsdk:"description"`
	Name                    types.String `tfsdk:"name"`
	NameAlias               types.String `tfsdk:"name_alias"`
	OwnerKey                types.String `tfsdk:"owner_key"`
	OwnerTag                types.String `tfsdk:"owner_tag"`
	InfraRsHPathAtt         types.Set    `tfsdk:"relation_to_host_path"`
	InfraRsPathToAccBaseGrp types.Set    `tfsdk:"relation_to_access_interface_policy_group"`
	TagAnnotation           types.Set    `tfsdk:"annotations"`
	TagTag                  types.Set    `tfsdk:"tags"`
}

func getEmptyInfraHPathSResourceModel() *InfraHPathSResourceModel {
	return &InfraHPathSResourceModel{
		Id:         basetypes.NewStringNull(),
		ParentDn:   basetypes.NewStringNull(),
		Annotation: basetypes.NewStringNull(),
		Descr:      basetypes.NewStringNull(),
		Name:       basetypes.NewStringNull(),
		NameAlias:  basetypes.NewStringNull(),
		OwnerKey:   basetypes.NewStringNull(),
		OwnerTag:   basetypes.NewStringNull(),
		InfraRsHPathAtt: types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"annotation": types.StringType,
				"target_dn":  types.StringType,
			},
		}),
		InfraRsPathToAccBaseGrp: types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"annotation": types.StringType,
				"target_dn":  types.StringType,
			},
		}),
		TagAnnotation: types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":   types.StringType,
				"value": types.StringType,
			},
		}),
		TagTag: types.SetNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"key":   types.StringType,
				"value": types.StringType,
			},
		}),
	}
}

// InfraRsHPathAttInfraHPathSResourceModel describes the resource data model for the children without relation ships.
type InfraRsHPathAttInfraHPathSResourceModel struct {
	Annotation types.String `tfsdk:"annotation"`
	TDn        types.String `tfsdk:"target_dn"`
}

func getEmptyInfraRsHPathAttInfraHPathSResourceModel() InfraRsHPathAttInfraHPathSResourceModel {
	return InfraRsHPathAttInfraHPathSResourceModel{
		Annotation: basetypes.NewStringNull(),
		TDn:        basetypes.NewStringNull(),
	}
}

// InfraRsPathToAccBaseGrpInfraHPathSResourceModel describes the resource data model for the children without relation ships.
type InfraRsPathToAccBaseGrpInfraHPathSResourceModel struct {
	Annotation types.String `tfsdk:"annotation"`
	TDn        types.String `tfsdk:"target_dn"`
}

func getEmptyInfraRsPathToAccBaseGrpInfraHPathSResourceModel() InfraRsPathToAccBaseGrpInfraHPathSResourceModel {
	return InfraRsPathToAccBaseGrpInfraHPathSResourceModel{
		Annotation: basetypes.NewStringNull(),
		TDn:        basetypes.NewStringNull(),
	}
}

// TagAnnotationInfraHPathSResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationInfraHPathSResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagAnnotationInfraHPathSResourceModel() TagAnnotationInfraHPathSResourceModel {
	return TagAnnotationInfraHPathSResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

// TagTagInfraHPathSResourceModel describes the resource data model for the children without relation ships.
type TagTagInfraHPathSResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

func getEmptyTagTagInfraHPathSResourceModel() TagTagInfraHPathSResourceModel {
	return TagTagInfraHPathSResourceModel{
		Key:   basetypes.NewStringNull(),
		Value: basetypes.NewStringNull(),
	}
}

type InfraHPathSIdentifier struct {
	Name types.String
}

func (r *InfraHPathSResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *InfraHPathSResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setInfraHPathSId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "infraHPathS", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *InfraHPathSResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_access_interface_override")
	resp.TypeName = req.ProviderTypeName + "_access_interface_override"
	tflog.Debug(ctx, "End metadata of resource: aci_access_interface_override")
}

func (r *InfraHPathSResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_access_interface_override")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The access_interface_override resource for the 'infraHPathS' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Access Interface Override object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("uni/infra"),
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the Access Interface Override object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the Access Interface Override object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the Access Interface Override object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the Access Interface Override object.`,
			},
			"owner_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"relation_to_host_path": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Set{
					setvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The annotation of the Relation To Host Path object.`,
						},
						"target_dn": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The distinguished name of the target.`,
						},
					},
				},
			},
			"relation_to_access_interface_policy_group": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Set{
					setvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The annotation of the Relation To Access Interface Policy Group object.`,
						},
						"target_dn": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The distinguished name of the target.`,
						},
					},
				},
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_access_interface_override")
}

func (r *InfraHPathSResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_access_interface_override")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, "End configure of resource: aci_access_interface_override")
}

func (r *InfraHPathSResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_access_interface_override")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *InfraHPathSResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setInfraHPathSId(ctx, stateData)
	}
	getAndSetInfraHPathSAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The infraHPathS object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *InfraHPathSResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setInfraHPathSId(ctx, data)
	}

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))

	var infraRsHPathAttPlan, infraRsHPathAttState []InfraRsHPathAttInfraHPathSResourceModel
	data.InfraRsHPathAtt.ElementsAs(ctx, &infraRsHPathAttPlan, false)
	stateData.InfraRsHPathAtt.ElementsAs(ctx, &infraRsHPathAttState, false)
	var infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState []InfraRsPathToAccBaseGrpInfraHPathSResourceModel
	data.InfraRsPathToAccBaseGrp.ElementsAs(ctx, &infraRsPathToAccBaseGrpPlan, false)
	stateData.InfraRsPathToAccBaseGrp.ElementsAs(ctx, &infraRsPathToAccBaseGrpState, false)
	var tagAnnotationPlan, tagAnnotationState []TagAnnotationInfraHPathSResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagInfraHPathSResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getInfraHPathSCreateJsonPayload(ctx, &resp.Diagnostics, true, data, infraRsHPathAttPlan, infraRsHPathAttState, infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetInfraHPathSAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))
}

func (r *InfraHPathSResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_access_interface_override")
	var data *InfraHPathSResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))

	getAndSetInfraHPathSAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *InfraHPathSResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))
}

func (r *InfraHPathSResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_access_interface_override")
	var data *InfraHPathSResourceModel
	var stateData *InfraHPathSResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))

	var infraRsHPathAttPlan, infraRsHPathAttState []InfraRsHPathAttInfraHPathSResourceModel
	data.InfraRsHPathAtt.ElementsAs(ctx, &infraRsHPathAttPlan, false)
	stateData.InfraRsHPathAtt.ElementsAs(ctx, &infraRsHPathAttState, false)
	var infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState []InfraRsPathToAccBaseGrpInfraHPathSResourceModel
	data.InfraRsPathToAccBaseGrp.ElementsAs(ctx, &infraRsPathToAccBaseGrpPlan, false)
	stateData.InfraRsPathToAccBaseGrp.ElementsAs(ctx, &infraRsPathToAccBaseGrpState, false)
	var tagAnnotationPlan, tagAnnotationState []TagAnnotationInfraHPathSResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagInfraHPathSResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getInfraHPathSCreateJsonPayload(ctx, &resp.Diagnostics, false, data, infraRsHPathAttPlan, infraRsHPathAttState, infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetInfraHPathSAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))
}

func (r *InfraHPathSResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_access_interface_override")
	var data *InfraHPathSResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "infraHPathS", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_access_interface_override with id '%s'", data.Id.ValueString()))
}

func (r *InfraHPathSResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_access_interface_override")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *InfraHPathSResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_access_interface_override with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_access_interface_override")
}

func getAndSetInfraHPathSAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *InfraHPathSResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "infraHPathS,infraRsHPathAtt,infraRsPathToAccBaseGrp,tagAnnotation,tagTag"), "GET", nil)

	*data = *getEmptyInfraHPathSResourceModel()

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("infraHPathS").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("infraHPathS").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setInfraHPathSParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					data.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					data.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerKey" {
					data.OwnerKey = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerTag" {
					data.OwnerTag = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			InfraRsHPathAttInfraHPathSList := make([]InfraRsHPathAttInfraHPathSResourceModel, 0)
			InfraRsPathToAccBaseGrpInfraHPathSList := make([]InfraRsPathToAccBaseGrpInfraHPathSResourceModel, 0)
			TagAnnotationInfraHPathSList := make([]TagAnnotationInfraHPathSResourceModel, 0)
			TagTagInfraHPathSList := make([]TagTagInfraHPathSResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "infraRsHPathAtt" {
							InfraRsHPathAttInfraHPathS := getEmptyInfraRsHPathAttInfraHPathSResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "annotation" {
									InfraRsHPathAttInfraHPathS.Annotation = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "tDn" {
									InfraRsHPathAttInfraHPathS.TDn = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							InfraRsHPathAttInfraHPathSList = append(InfraRsHPathAttInfraHPathSList, InfraRsHPathAttInfraHPathS)
						}
						if childClassName == "infraRsPathToAccBaseGrp" {
							InfraRsPathToAccBaseGrpInfraHPathS := getEmptyInfraRsPathToAccBaseGrpInfraHPathSResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "annotation" {
									InfraRsPathToAccBaseGrpInfraHPathS.Annotation = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "tDn" {
									InfraRsPathToAccBaseGrpInfraHPathS.TDn = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							InfraRsPathToAccBaseGrpInfraHPathSList = append(InfraRsPathToAccBaseGrpInfraHPathSList, InfraRsPathToAccBaseGrpInfraHPathS)
						}
						if childClassName == "tagAnnotation" {
							TagAnnotationInfraHPathS := getEmptyTagAnnotationInfraHPathSResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationInfraHPathS.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationInfraHPathS.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationInfraHPathSList = append(TagAnnotationInfraHPathSList, TagAnnotationInfraHPathS)
						}
						if childClassName == "tagTag" {
							TagTagInfraHPathS := getEmptyTagTagInfraHPathSResourceModel()
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagInfraHPathS.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagInfraHPathS.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagInfraHPathSList = append(TagTagInfraHPathSList, TagTagInfraHPathS)
						}
					}
				}
			}
			infraRsHPathAttSet, _ := types.SetValueFrom(ctx, data.InfraRsHPathAtt.ElementType(ctx), InfraRsHPathAttInfraHPathSList)
			data.InfraRsHPathAtt = infraRsHPathAttSet
			infraRsPathToAccBaseGrpSet, _ := types.SetValueFrom(ctx, data.InfraRsPathToAccBaseGrp.ElementType(ctx), InfraRsPathToAccBaseGrpInfraHPathSList)
			data.InfraRsPathToAccBaseGrp = infraRsPathToAccBaseGrpSet
			tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationInfraHPathSList)
			data.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, data.TagTag.ElementType(ctx), TagTagInfraHPathSList)
			data.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'infraHPathS'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getInfraHPathSRn(ctx context.Context, data *InfraHPathSResourceModel) string {
	rn := "hpaths-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setInfraHPathSParentDn(ctx context.Context, dn string, data *InfraHPathSResourceModel) {
	bracketIndex := 0
	rnIndex := 0
	for i := len(dn) - 1; i >= 0; i-- {
		if string(dn[i]) == "]" {
			bracketIndex = bracketIndex + 1
		} else if string(dn[i]) == "[" {
			bracketIndex = bracketIndex - 1
		} else if string(dn[i]) == "/" && bracketIndex == 0 {
			rnIndex = i
			break
		}
	}
	data.ParentDn = basetypes.NewStringValue(dn[:rnIndex])
}

func setInfraHPathSId(ctx context.Context, data *InfraHPathSResourceModel) {
	rn := getInfraHPathSRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getInfraHPathSInfraRsHPathAttChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *InfraHPathSResourceModel, infraRsHPathAttPlan, infraRsHPathAttState []InfraRsHPathAttInfraHPathSResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.InfraRsHPathAtt.IsUnknown() {
		for _, infraRsHPathAtt := range infraRsHPathAttPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !infraRsHPathAtt.Annotation.IsUnknown() && !infraRsHPathAtt.Annotation.IsNull() {
				childMap["attributes"]["annotation"] = infraRsHPathAtt.Annotation.ValueString()
			} else {
				childMap["attributes"]["annotation"] = globalAnnotation
			}
			if !infraRsHPathAtt.TDn.IsUnknown() && !infraRsHPathAtt.TDn.IsNull() {
				childMap["attributes"]["tDn"] = infraRsHPathAtt.TDn.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"infraRsHPathAtt": childMap})
		}
		if len(infraRsHPathAttPlan) == 0 && len(infraRsHPathAttState) == 1 {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			childMap["attributes"]["status"] = "deleted"
			childMap["attributes"]["tDn"] = infraRsHPathAttState[0].TDn.ValueString()
			childPayloads = append(childPayloads, map[string]interface{}{"infraRsHPathAtt": childMap})
		}
	} else {
		data.InfraRsHPathAtt = types.SetNull(data.InfraRsHPathAtt.ElementType(ctx))
	}

	return childPayloads
}
func getInfraHPathSInfraRsPathToAccBaseGrpChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *InfraHPathSResourceModel, infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState []InfraRsPathToAccBaseGrpInfraHPathSResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.InfraRsPathToAccBaseGrp.IsUnknown() {
		for _, infraRsPathToAccBaseGrp := range infraRsPathToAccBaseGrpPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !infraRsPathToAccBaseGrp.Annotation.IsUnknown() && !infraRsPathToAccBaseGrp.Annotation.IsNull() {
				childMap["attributes"]["annotation"] = infraRsPathToAccBaseGrp.Annotation.ValueString()
			} else {
				childMap["attributes"]["annotation"] = globalAnnotation
			}
			if !infraRsPathToAccBaseGrp.TDn.IsUnknown() && !infraRsPathToAccBaseGrp.TDn.IsNull() {
				childMap["attributes"]["tDn"] = infraRsPathToAccBaseGrp.TDn.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"infraRsPathToAccBaseGrp": childMap})
		}
		if len(infraRsPathToAccBaseGrpPlan) == 0 && len(infraRsPathToAccBaseGrpState) == 1 {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			childMap["attributes"]["status"] = "deleted"
			childPayloads = append(childPayloads, map[string]interface{}{"infraRsPathToAccBaseGrp": childMap})
		}
	} else {
		data.InfraRsPathToAccBaseGrp = types.SetNull(data.InfraRsPathToAccBaseGrp.ElementType(ctx))
	}

	return childPayloads
}
func getInfraHPathSTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *InfraHPathSResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationInfraHPathSResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotation := range tagAnnotationPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagAnnotation.Key.IsUnknown() && !tagAnnotation.Key.IsNull() {
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
			}
			if !tagAnnotation.Value.IsUnknown() && !tagAnnotation.Value.IsNull() {
				childMap["attributes"]["value"] = tagAnnotation.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotation.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationState {
			delete := true
			for _, tagAnnotationIdentifier := range tagAnnotationIdentifiers {
				if tagAnnotationIdentifier.Key == tagAnnotation.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			}
		}
	} else {
		data.TagAnnotation = types.SetNull(data.TagAnnotation.ElementType(ctx))
	}

	return childPayloads
}
func getInfraHPathSTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *InfraHPathSResourceModel, tagTagPlan, tagTagState []TagTagInfraHPathSResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTag := range tagTagPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagTag.Key.IsUnknown() && !tagTag.Key.IsNull() {
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
			}
			if !tagTag.Value.IsUnknown() && !tagTag.Value.IsNull() {
				childMap["attributes"]["value"] = tagTag.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTag.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagState {
			delete := true
			for _, tagTagIdentifier := range tagTagIdentifiers {
				if tagTagIdentifier.Key == tagTag.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			}
		}
	} else {
		data.TagTag = types.SetNull(data.TagTag.ElementType(ctx))
	}

	return childPayloads
}

func getInfraHPathSCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *InfraHPathSResourceModel, infraRsHPathAttPlan, infraRsHPathAttState []InfraRsHPathAttInfraHPathSResourceModel, infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState []InfraRsPathToAccBaseGrpInfraHPathSResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationInfraHPathSResourceModel, tagTagPlan, tagTagState []TagTagInfraHPathSResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	InfraRsHPathAttchildPayloads := getInfraHPathSInfraRsHPathAttChildPayloads(ctx, diags, data, infraRsHPathAttPlan, infraRsHPathAttState)
	if InfraRsHPathAttchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, InfraRsHPathAttchildPayloads...)

	InfraRsPathToAccBaseGrpchildPayloads := getInfraHPathSInfraRsPathToAccBaseGrpChildPayloads(ctx, diags, data, infraRsPathToAccBaseGrpPlan, infraRsPathToAccBaseGrpState)
	if InfraRsPathToAccBaseGrpchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, InfraRsPathToAccBaseGrpchildPayloads...)

	TagAnnotationchildPayloads := getInfraHPathSTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getInfraHPathSTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}
	if !data.OwnerKey.IsNull() && !data.OwnerKey.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerKey"] = data.OwnerKey.ValueString()
	}
	if !data.OwnerTag.IsNull() && !data.OwnerTag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerTag"] = data.OwnerTag.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"infraHPathS": payloadMap})
	if err != nil {
		diags.AddError(
			"Marshalling of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := container.ParseJSON(payload)

	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}
