package provider

import (
	"fmt"
	"regexp"
	"testing"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCreateFictionalUserAndImport(t *testing.T) {
	salutation := "male"
	firstname := "WGC"
	lastname := "Bexio"
	email := "wgc.bexio@webgate.biz"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContactCreateFicitonalUserConfigBasic(salutation, firstname, lastname, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "salutation_type", salutation),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "firstname", firstname),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "lastname", lastname),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "email", email),
				),
			},
			{
				ResourceName:      "bexio_fictionaluser.create",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccUpdateFictionalUser(t *testing.T) {
	salutation := "male"
	firstname := "WGC"
	firstname2 := "UPDATED"
	lastname := "Bexio"
	email := "wgc.bexio@webgate.biz"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContactCreateFicitonalUserConfigBasic(salutation, firstname, lastname, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "salutation_type", salutation),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "firstname", firstname),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "lastname", lastname),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "email", email),
				),
			},
			{
				Config: testAccContactCreateFicitonalUserConfigBasic(salutation, firstname2, lastname, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "salutation_type", salutation),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "firstname", firstname2),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "lastname", lastname),
					resource.TestCheckResourceAttr("bexio_fictionaluser.create", "email", email),
				),
			},
		},
	})
}
func testAccContactCreateFicitonalUserConfigBasic(v1, v2, v3, v4 string) string {
	return fmt.Sprintf(`
	resource "bexio_fictionaluser" "create" {
		salutation_type         = "%s"
		firstname               = "%s"
		lastname                = "%s"
		email                   = "%s"
	  }
	`, v1, v2, v3, v4)
}
func testAccCheckExampleResourceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "bexio_fictionaluser" {
			continue
		}

		_, err := c.ReadFictionalUser(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Functional user still exists")
		}
		notFoundErr := "Page not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}
	return nil
}
