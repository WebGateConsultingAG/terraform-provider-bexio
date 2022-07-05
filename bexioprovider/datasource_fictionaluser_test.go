//toDo: add creation first, check and destroy
package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFictionalUserById(t *testing.T) {
	var fictionaluserid int = 9
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccFictionalUserConfigBasic(fictionaluserid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.bexio_fictionaluserbyid.this", "id", fmt.Sprintf("%d", fictionaluserid)),
				),
			},
		},
	})
}

func testAccFictionalUserConfigBasic(d1 int) string {
	return fmt.Sprintf(`
	data "bexio_fictionaluserbyid" "this" {
		id        = %d
	}
	`, d1)
}
