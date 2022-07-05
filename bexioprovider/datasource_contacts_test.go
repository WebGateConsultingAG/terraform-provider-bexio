package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccContacts(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccContactsConfigBasic(),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccContactsConfigBasic() string {
	return fmt.Sprintf(`
	data "bexio_contacts" "this" {
	}
	`)
}
