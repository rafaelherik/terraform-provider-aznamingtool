package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient/models"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/utils"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ResourceNameDataSource{}
	_ datasource.DataSourceWithConfigure = &ResourceNameDataSource{}
)

func NewResourceNameDataSource() datasource.DataSource {
	return &ResourceNameDataSource{}
}

// ResourceNameDataSource defines the data source implementation.
type ResourceNameDataSource struct {
	client *apiclient.APIClient
}

// ResourceNameDataSourceModel describes the data source data model.
type ResourceNameDataSourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	ResourceName     types.String `tfsdk:"resource_name"`
	ResourceTypeName types.String `tfsdk:"resource_type_name"`
	Components       types.Map    `tfsdk:"components"`
	CreatedOn        types.String `tfsdk:"created_on"`
}

// Metadata returns the data source type name.
func (d *ResourceNameDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "aznamingtool_resource_name"
}

// Schema defines the schema for the data source.
func (d *ResourceNameDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"resource_name": schema.StringAttribute{
				Computed: true,
			},
			"resource_type_name": schema.StringAttribute{
				Computed: true,
			},
			"components": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
			"created_on": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Configure prepares the struct.
func (d *ResourceNameDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*apiclient.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *apiclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

// Read handles reading the data source data.
func (d *ResourceNameDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state ResourceNameDataSourceModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if d.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The provider client has not been configured.")
		return
	}

	svc := apiclient.NewResourceNamingService(d.client)
	resource, err := svc.GetGeneratedName(state.ID.String())
	if err != nil {
		resp.Diagnostics.AddError("Error reading resource", err.Error())
		diags = resp.State.Set(ctx, state)
		resp.Diagnostics.Append(diags...)
		return
	}

	finalData, err := transformResponseToDataSourceSchema(resource)

	if err != nil {
		resp.Diagnostics.AddError("Failed to transform the response.", err.Error())
		return
	}
	state.ID = finalData.ID
	state.Components = finalData.Components
	state.CreatedOn = finalData.CreatedOn
	state.ResourceName = finalData.ResourceName
	state.ResourceTypeName = finalData.ResourceTypeName

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func transformResponseToDataSourceSchema(resource *models.ResourceGeneratedName) (*ResourceNameDataSourceModel, error) {
	componentsMap := make(map[string]attr.Value)
	for _, component := range resource.Components {
		if len(component) == 2 {
			snakeKey := utils.CamelToSnake(component[0])
			componentsMap[snakeKey] = types.StringValue(component[1])
		}
	}
	components, diags := types.MapValue(types.StringType, componentsMap)
	if diags.HasError() {
		return nil, fmt.Errorf("failed to transform components: %v", diags.Errors())
	}

	return &ResourceNameDataSourceModel{
		ID:               types.Int64Value(resource.Id),
		ResourceName:     types.StringValue(resource.ResourceName),
		ResourceTypeName: types.StringValue(resource.ResourceTypeName),
		Components:       components,
		CreatedOn:        types.StringValue(resource.CreatedOn),
	}, nil
}

var mapKeyToField = map[string]string{
	"resource_environment":  "ResourceEnvironment",
	"resource_function":     "ResourceFunction",
	"resource_instance":     "ResourceInstance",
	"resource_location":     "ResourceLocation",
	"resource_org":          "ResourceOrg",
	"resource_proj_app_svc": "ResourceProjAppSvc",
	"resource_type":         "ResourceType",
	"resource_unit_dept":    "ResourceUnitDept",
}
