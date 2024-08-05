package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rafaelherik/terraform-provider-aznamingtool/apiclient"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ResourceNameDataSource{}
	_ datasource.DataSourceWithConfigure = &ResourceNameDataSource{}
)

func NewResourceNameDataSource() datasource.DataSource {
	return &ResourceNameDataSource{}
}

// ExampleDataSource defines the data source implementation.
type ResourceNameDataSource struct {
	client *apiclient.APIClient
}

// ExampleDataSourceModel describes the data source data model.
type ResourceNameDataSourceModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// Metadata returns the data source type name.
func (d *ResourceNameDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "azure_naming_data_source"
}

// Schema defines the schema for the data source.
func (d *ResourceNameDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
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
			fmt.Sprintf("Expected *session.Session, got: %T. Please report this issue to the provider developers.", req.ProviderData),
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

	// Example API call to read the data source
	// Assuming the API response contains the name of the resource

	// Example response processing
	state.Name = types.StringValue("example-name")

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}
