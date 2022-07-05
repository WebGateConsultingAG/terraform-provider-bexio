package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSearchcontact(t *testing.T) {
	mail := "support@bexio.com"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccContactConfigBasic(mail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.bexio_contact.this", "contact_email", "support@bexio.com"),
				),
			},
		},
	})
}

func testAccContactConfigBasic(s1 string) string {
	return fmt.Sprintf(`
		data "bexio_contact" "this" {
			contact_email        = "%s"
		}
		`, s1)
}
