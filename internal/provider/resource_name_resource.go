package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/apiclient/models"
	"github.com/rafaelherik/terraform-provider-aznamingtool/tools/utils"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &AzureNameResource{}
	_ resource.ResourceWithConfigure = &AzureNameResource{}
)

func NewAzureNameResource() resource.Resource {
	return &AzureNameResource{}
}

// ExampleResource defines the resource implementation.
type AzureNameResource struct {
	client *apiclient.APIClient
}

// ExampleResourceModel describes the resource data model.
type AzureNameResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	ResourceName     types.String `tfsdk:"resource_name"`
	ResourceTypeId   types.Int64  `tfsdk:"resource_type_id"`
	ResourceTypeName types.String `tfsdk:"resource_type_name"`
	Components       types.Map    `tfsdk:"components"`
	CreatedOn        types.String `tfsdk:"created_on"`
}

// Metadata returns the resource type name.
func (r *AzureNameResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "aznamingtool_resource_name"
}

// Schema defines the schema for the resource.
func (r *AzureNameResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed: true,
			},
			"resource_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_type_id": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"resource_type_name": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"components": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
					mapplanmodifier.UseStateForUnknown(),
				},
			},
			"created_on": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

// Configure prepares the struct.
func (r *AzureNameResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *AzureNameResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan *AzureNameResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if r.client == nil {
		resp.Diagnostics.AddError("Client not configured", "The provider client has not been configured.")
		return
	}

	request, _err := plan.ToResourceRequest()

	if _err != nil {
		resp.Diagnostics.AddError("Failed to transform the request.", _err.Error())
		return
	}

	svc := apiclient.NewResourceNamingService(r.client)
	result, _err := svc.RequestName(request)

	if _err != nil {
		resp.Diagnostics.AddError("Failed to request the name.", _err.Error())
		return
	}
	plan.ID = types.Int64Value(result.ResourceNameDetails.Id)
	// Call the read function to update the state
	plan, _err = _ReadFromAPI(r.client, plan.ID.String())

	if request.ResourceId != 0 {
		plan.ResourceTypeId = types.Int64Value(request.ResourceId)
	}

	if _err != nil {
		resp.Diagnostics.AddError("Failed to read the resource.", _err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)

}

func _ReadFromAPI(client *apiclient.APIClient, id string) (*AzureNameResourceModel, error) {
	svc := apiclient.NewResourceNamingService(client)
	result, err := svc.GetGeneratedName(id)
	if err != nil {
		return nil, err
	}
	return transformResponseToSchema(result)

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
		resp.Diagnostics.AddWarning("Failed to get the generated name.", result.Message)
		resp.Diagnostics.AddError("Failed to get the generated name.", _err.Error())
		return

	} else {
		adminSvc := apiclient.NewResourceNamingService(r.client)
		generatedName, _err := adminSvc.GetGeneratedName(state.ID.String())

		if _err != nil {
			resp.Diagnostics.AddError("Failed to get the generated name.", _err.Error())
			return
		}
		finalData, err := transformResponseToSchema(generatedName)

		if err != nil {
			resp.Diagnostics.AddError("Failed to transform the response.", err.Error())
			return
		}

		state.ID = finalData.ID
		state.Components = finalData.Components
		state.CreatedOn = finalData.CreatedOn
		state.ResourceName = finalData.ResourceName
		state.ResourceTypeName = finalData.ResourceTypeName
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// Update handles updating the resource.
func (r *AzureNameResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AzureNameResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	resp.Diagnostics.AddError("Update not supported", "This resource does not support updates.")
}

// Delete handles deleting the resource.
func (r *AzureNameResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AzureNameResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	svc := apiclient.NewResourceNamingService(r.client)
	_err := svc.DeleteGeneratedName(state.ID.String())
	if _err != nil {
		resp.Diagnostics.AddError("Failed to get the generated name.", _err.Error())
		return

	}
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// ImportState handles importing the resource state.
func (r *AzureNameResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Example import logic
}

func (r *AzureNameResourceModel) ToResourceRequest() (*models.ResourceNameRequest, error) {
	request := &models.ResourceNameRequest{
		CustomComponents: make(map[string]string),
	}

	if !r.ResourceTypeId.IsNull() && !r.ResourceTypeId.IsUnknown() {
		request.ResourceId = r.ResourceTypeId.ValueInt64()
	}

	for key, value := range r.Components.Elements() {
		stringValue := value.(types.String).ValueString()
		mappedKey := strings.ToLower(key)

		if fieldName, exists := mapKeyToField[mappedKey]; exists {
			switch fieldName {
			case "ResourceEnvironment":
				request.ResourceEnvironment = stringValue
			case "ResourceFunction":
				request.ResourceFunction = stringValue
			case "ResourceInstance":
				request.ResourceInstance = stringValue
			case "ResourceLocation":
				request.ResourceLocation = stringValue
			case "ResourceOrg":
				request.ResourceOrg = stringValue
			case "ResourceProjAppSvc":
				request.ResourceProjAppSvc = stringValue
			case "ResourceType":
				request.ResourceType = stringValue
			case "ResourceUnitDept":
				request.ResourceUnitDept = stringValue
			}
		} else {
			request.CustomComponents[key] = stringValue
		}
	}

	return request, nil
}

func (plan *AzureNameResourceModel) FromResourceName(request *models.ResourceNameRequest, result *models.ResourceGeneratedName) (*AzureNameResourceModel, error) {

	return transformResponseToSchema(result)
}

func transformResponseToSchema(resource *models.ResourceGeneratedName) (*AzureNameResourceModel, error) {
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

	return &AzureNameResourceModel{
		ID:               types.Int64Value(resource.Id),
		ResourceName:     types.StringValue(resource.ResourceName),
		ResourceTypeName: types.StringValue(resource.ResourceTypeName),
		Components:       components,
		CreatedOn:        types.StringValue(resource.CreatedOn),
	}, nil
}
