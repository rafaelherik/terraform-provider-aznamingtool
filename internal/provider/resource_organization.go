package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient/models"
)

var (
	_ resource.Resource              = &AzureOrganiztionResource{}
	_ resource.ResourceWithConfigure = &AzureOrganiztionResource{}
)

type AzureOrganiztionResource struct {
	client *apiclient.APIClient
}

func NewAzureOrganizationResource() resource.Resource {
	return &AzureOrganiztionResource{}
}

type AzureOrganiztionResourceModel struct {
	ID        types.Int32  `tfsdk:"organization_id"`
	Name      types.String `tfsdk:"organization_name"`
	ShortName types.String `tfsdk:"organization_short_name"`
}

func (r *AzureOrganiztionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "aznamingtool_resource_name"
}

func (r *AzureOrganiztionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.Int64Attribute{
				Computed: true,
			},
			"organization_name": schema.StringAttribute{
				Computed: true,
			},
			"organization_short_name": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func (r *AzureOrganiztionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*apiclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ApiClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}
	r.client = client
}

// Create handles the creation of the resource.
func (r *AzureOrganiztionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AzureOrganiztionResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The provider client has not been configured.")
		return
	}

	request := models.ResourceOrganization{}

	resp.Diagnostics.AddWarning("Testing", fmt.Sprintf("Request data: %#v", request))

	svc := apiclient.NewResourceOrganizationService(r.client)
	result, _err := svc.CreateOrUpdateResourceOrganization(request)

	if _err != nil {
		resp.Diagnostics.AddError("Failed to request the name.", _err.Error())
		return
	} else {

		plan.ID = types.Int32Value(result.Id)
		plan.Name = types.StringValue(result.Name)
		plan.ShortName = types.StringValue(result.ShortName)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read handles reading the resource data.
func (r *AzureOrganiztionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AzureNameResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	svc := apiclient.NewResourceNamingService(r.client)

	result, _err := svc.GetGeneratedName(state.ID.String())

	if _err != nil {
		resp.Diagnostics.AddWarning("Failed to get the generated name.", result.Message)
		resp.Diagnostics.AddError("Failed to get the generated name.", _err.Error())
		return

	} else {
		state.ID = types.Int64Value(result.Id)
		state.Name = types.StringValue(result.ResourceName)
		state.ResourceType = types.StringValue(result.ResourceTypeName)

	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update handles updating the resource.
func (r *AzureOrganiztionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

}

// Delete handles deleting the resource.
func (r *AzureOrganiztionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Example API call to delete the resource
}

// ImportState handles importing the resource state.
func (r *AzureOrganiztionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Example import logic
}
