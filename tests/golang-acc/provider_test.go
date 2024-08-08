package acctests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/rafaelherik/terraform-provider-aznamingtool/internal/provider"
)

const (
	// providerConfig is a shared configuration to combine with the actual
	// test configuration so the HashiCups client is properly configured.
	// It is also possible to use the HASHICUPS_ environment variables instead,
	// such as updating the Makefile and running the testing through that tool.
	providerConfig = `
	provider "aznamingtool" {
		base_url = "%s"
		api_key = "test"
		admin_password = "test123"
	  }
`
)

var (
	// testAccProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"hashicups": providerserver.NewProtocol6WithError(provider.NewProvider("test")()),
	}
)

func TestProviderResource(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/resource":
			if r.Method == http.MethodPost {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"id": "12345", "name": "test-resource"}`))
			} else if r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"id": "12345", "name": "test-resource"}`))
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer mockServer.Close()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceConfig(mockServer.URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("aznamingtool.test", "name", "test-resource"),
				),
			},
		},
	})
}

func testAccPreCheck(t *testing.T) {
	return
}

func testAccResourceConfig(serverURL string) string {
	return fmt.Sprintf(`
provider "aznamingtool" {
  base_url = "%s"
  api_key = "test"
  admin_password = "test123"
}

resource "aznamingtool_resource_name" "test" {
  name = "test-resource"
}
`, serverURL)
}
