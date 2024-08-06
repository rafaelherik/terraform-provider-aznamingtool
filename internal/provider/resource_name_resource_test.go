// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("scaffolding_example.test", "configurable_attribute", "one"),
					resource.TestCheckResourceAttr("scaffolding_example.test", "defaulted", "example value when not configured"),
					resource.TestCheckResourceAttr("scaffolding_example.test", "id", "example-id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "azurenamingtool_resource_name.test",
				ImportState:       false,
				ImportStateVerify: false,
			},
		},
	})
}
