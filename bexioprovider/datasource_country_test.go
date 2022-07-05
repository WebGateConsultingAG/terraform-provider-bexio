package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSearchCountry(t *testing.T) {
	countryName := "Schweiz"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCountryConfigBasic(countryName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.bexio_country.this", "id", "1"),
				),
			},
		},
	})
}

func testAccCountryConfigBasic(s1 string) string {
	return fmt.Sprintf(`
		data "bexio_country" "this" {
			name        = "%s"
		}
		`, s1)
}
