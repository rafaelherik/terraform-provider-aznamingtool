package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rafaelherik/terraform-provider-aznamingtool/apiclient"
	"github.com/rafaelherik/terraform-provider-aznamingtool/apiclient/models"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_                      resource.Resource              = &AzureNameResource{}
	_                      resource.ResourceWithConfigure = &AzureNameResource{}
	validResourceTypes     *[]models.ResourceType
	validResourceTypeNames []string
)

type AzureResourceType struct {
}

func NewAzureNameResource() resource.Resource {
	return &AzureNameResource{}
}

// ExampleResource defines the resource implementation.
type AzureNameResource struct {
	client *apiclient.APIClient
}

// ExampleResourceModel describes the resource data model.
type AzureNameResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	ResourceType   types.String `tfsdk:"resource_type"`
	Organization   types.String `tfsdk:"organization"`
	BusinessUnit   types.String `tfsdk:"business_unit"`
	Project        types.String `tfsdk:"project"`
	Location       types.String `tfsdk:"location"`
	Environment    types.String `tfsdk:"environment"`
	ResourceTypeId types.Int64  `tfsdk:"resource_id"`
	CreatedAt      types.String `tfsdk:"created_at"`
}

// Metadata returns the resource type name.
func (r *AzureNameResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "aznamingtool_resource_name"
}

// Schema defines the schema for the resource.
func (r *AzureNameResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"organization": schema.StringAttribute{
				Required: true,
			},
			"business_unit": schema.StringAttribute{
				Optional: true,
			},
			"project": schema.StringAttribute{
				Required: true,
			},
			"resource_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					validators.resourceTypeValidator(validResourceTypes),
				},
			},
			"location": schema.StringAttribute{
				Required: true,
			},
			"environment": schema.StringAttribute{
				Required: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Configure prepares the struct.
func (r *AzureNameResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = req.ProviderData.(*apiclient.APIClient)
	validResourceTypes, _ = _GetAllResourceTypes(r.client)
	validResourceTypeNames = func(resourceType *[]models.ResourceType) []string {
		resources := make([]string, len(*resourceType))
		for i, resourceType := range *resourceType {
			resources[i] = resourceType.Resource
		}
		return resources
	}(validResourceTypes)

}

// Create handles the creation of the resource.
func (r *AzureNameResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AzureNameResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	request := models.ResourceNameRequest{
		ResourceOrg:         plan.Organization.String(),
		ResourceUnitDept:    plan.BusinessUnit.String(),
		ResourceProjAppSvc:  plan.Project.String(),
		ResourceType:        plan.ResourceType.String(),
		ResourceLocation:    plan.Location.String(),
		ResourceEnvironment: plan.Environment.String(),
		ResourceId:          plan.ResourceTypeId.ValueInt64(),
	}
	svc := apiclient.NewResourceNamingService(r.client)
	result, _err := svc.RequestName(request)

	if _err != nil {
		resp.Diagnostics.AddError("Failed to request the name.", _err.Error())
	} else {

		plan.ID = types.Int64Value(result.ResourceNameDetails.Id)
		plan.Name = types.StringValue(result.ResourceNameDetails.ResourceName)
		plan.ResourceType = types.StringValue(result.ResourceNameDetails.ResourceTypeName)
		plan.Organization = types.StringValue(request.ResourceOrg)
		plan.BusinessUnit = types.StringValue(request.ResourceUnitDept)
		plan.Project = types.StringValue(request.ResourceProjAppSvc)
		plan.Location = types.StringValue(request.ResourceLocation)
		plan.Environment = types.StringValue(request.ResourceEnvironment)
		plan.ResourceTypeId = types.Int64Value(request.ResourceId)
		plan.CreatedAt = types.StringValue(result.ResourceNameDetails.CreatedOn)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Read handles reading the resource data.
func (r *AzureNameResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AzureNameResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	svc := apiclient.NewResourceNamingService(r.client)

	result, _err := svc.GetGeneratedName(state.ID.String())

	if _err != nil {
		resp.Diagnostics.AddError("Failed to get the generated name.", _err.Error())

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
func (r *AzureNameResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AzureNameResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Example API call to update the resource

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

// Delete handles deleting the resource.
func (r *AzureNameResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Example API call to delete the resource
}

// ImportState handles importing the resource state.
func (r *AzureNameResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Example import logic
}

func _GetAllResourceTypes(c *apiclient.APIClient) (*[]models.ResourceType, error) {
	var response *[]models.ResourceType
	svc := apiclient.NewResourceTypeService(c)
	var err error
	response, err = svc.GetAllResourceTypes()
	if err != nil {
		return nil, err
	}

	return response, nil
}
