package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSearchSalutation(t *testing.T) {
	name := "Herr"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSalutationConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.bexio_salutation.this", "name", "Herr"),
				),
			},
		},
	})
}

func testAccSalutationConfigBasic(s1 string) string {
	return fmt.Sprintf(`
		data "bexio_salutation" "this" {
			name        = "%s"
		}
		`, s1)
}
