package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFictionalUsers(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFictionalUsersConfigBasic(),
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func testAccFictionalUsersConfigBasic() string {
	return fmt.Sprintf(`
	data "bexio_fictionalusers" "this" {
	}
	`)
}
