// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"fmt"

	customTypes "github.com/CiscoDevNet/terraform-provider-aci/v2/internal/custom_types"
	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &MgmtInstPDataSource{}

func NewMgmtInstPDataSource() datasource.DataSource {
	return &MgmtInstPDataSource{}
}

// MgmtInstPDataSource defines the data source implementation.
type MgmtInstPDataSource struct {
	client *client.Client
}

func (d *MgmtInstPDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_external_management_network_instance_profile")
	resp.TypeName = req.ProviderTypeName + "_external_management_network_instance_profile"
	tflog.Debug(ctx, "End metadata of datasource: aci_external_management_network_instance_profile")
}

func (d *MgmtInstPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_external_management_network_instance_profile")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The external_management_network_instance_profile datasource for the 'mgmtInstP' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the External Management Network Instance Profile object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the External Management Network Instance Profile object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the External Management Network Instance Profile object.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the External Management Network Instance Profile object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the External Management Network Instance Profile object.`,
			},
			"priority": schema.StringAttribute{
				CustomType:          customTypes.MgmtInstPPrioStringType{},
				Computed:            true,
				MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
			},
			"relation_to_consumed_out_of_band_contracts": schema.SetNestedAttribute{
				MarkdownDescription: `An external management entity instance profile to an out-of-band binary contract profile. The instance profiles of external management entities can communicate with nodes that are part of out-of-band management endpoint group. To enable this communication, a contract is required between the instance profile and the out-of-band management endpoint group.`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The annotation of the Relation To Consumed Out Of Band Contract object.`,
						},
						"priority": schema.StringAttribute{
							CustomType:          customTypes.MgmtRsOoBConsPrioStringType{},
							Computed:            true,
							MarkdownDescription: `The Quality of Service (QoS) priority class ID. QoS refers to the capability of a network to provide better service to selected network traffic over various technologies. The primary goal of QoS is to provide priority including dedicated bandwidth, controlled jitter and latency (required by some real-time and interactive traffic), and improved loss characteristics. You can configure the bandwidth of each QoS level using QoS profiles.`,
						},
						"out_of_band_contract_name": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The name of the Out Of Band Contract object.`,
						},
						"annotations": schema.SetNestedAttribute{
							MarkdownDescription: ``,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: `The key used to uniquely identify this configuration object.`,
									},
									"value": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: `The value of the property.`,
									},
								},
							},
						},
						"tags": schema.SetNestedAttribute{
							MarkdownDescription: ``,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: `The key used to uniquely identify this configuration object.`,
									},
									"value": schema.StringAttribute{
										Computed:            true,
										MarkdownDescription: `The value of the property.`,
									},
								},
							},
						},
					},
				},
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of datasource: aci_external_management_network_instance_profile")
}

func (d *MgmtInstPDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_external_management_network_instance_profile")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
	tflog.Debug(ctx, "End configure of datasource: aci_external_management_network_instance_profile")
}

func (d *MgmtInstPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_external_management_network_instance_profile")
	var data *MgmtInstPResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setMgmtInstPId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetMgmtInstPAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_external_management_network_instance_profile with id '%s'", data.Id.ValueString()))

	getAndSetMgmtInstPAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_external_management_network_instance_profile data source",
			fmt.Sprintf("The aci_external_management_network_instance_profile data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_external_management_network_instance_profile with id '%s'", data.Id.ValueString()))
}
