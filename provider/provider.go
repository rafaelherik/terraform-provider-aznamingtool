package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/rafaelherik/terraform-provider-aznamingtool/apiclient"
)

var (
	_ provider.Provider = &AzureNamingToolProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AzureNamingToolProvider{
			version: version,
		}
	}
}

type AzureNamingToolProvider struct {
	version string
}

type AzureNamingToolProviderModel struct {
	ApiKey       types.String `tfsdk:"api_key"`
	BaseUrl      types.String `tfsdk:"base_url"`
	AdminPassord types.String `tfsdk:"admin_password"`
}

func (p *AzureNamingToolProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_url": schema.StringAttribute{
				Optional: true,
			},
			"api_key": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
			"admin_password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *AzureNamingToolProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "azurenaming"
	resp.Version = p.version
}

func (p *AzureNamingToolProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config AzureNamingToolProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Api Key value",
			"The provider cannot create the API client as there is an unknown configuration value for the API key. ",
		)
	}
	if config.BaseUrl.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("base_url"),
			"Unknown Base Url value",
			"The provider cannot create the API client as there is an unknown configuration value for the base url. ",
		)
	}
	if config.AdminPassord.IsUnknown() {
		resp.Diagnostics.AddAttributeWarning(
			path.Root("admin_password"),
			"Unknown Admin Password value",
			"The provider cannot delete generated names if the admin password is unknown. ",
		)
	}

	api_key := os.Getenv("AZ_NAMINGTOOL_APIKEY")
	base_url := os.Getenv("AZ_NAMINGTOOL_BASEURL")
	admin_password := os.Getenv("AZ_NAMINGTOOL_ADMINPASSWORD")

	if !config.BaseUrl.IsNull() {
		base_url = config.BaseUrl.ValueString()
	}

	if !config.ApiKey.IsNull() {
		api_key = config.ApiKey.ValueString()
	}

	if !config.AdminPassord.IsNull() {
		admin_password = config.AdminPassord.ValueString()
	}
	// Logging
	tflog.Info(ctx, "Configuring ApiClient")

	// Example of configuring the client
	client := apiclient.NewAPIClient(api_key, base_url, admin_password, nil)

	// Make the client available during DataSource and Resource type Configure methods
	resp.DataSourceData = client
	resp.ResourceData = client

}

func (p *AzureNamingToolProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewResourceNameDataSource,
	}
}

func (p *AzureNamingToolProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAzureNameResource,
	}
}
