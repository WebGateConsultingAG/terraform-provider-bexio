package provider

import (
	"fmt"
	"regexp"
	"testing"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCreateContactAndImport(t *testing.T) {
	lastname := "Bexio"
	firstname := "WGC"
	birthday := "2003-09-23"
	address := "Riedstrasse 3"
	postcode := "8953"
	city := "Dietikon"
	country_id := 1
	email := "wgc.bexio@webgate.biz"
	emailprivate := "wgc.bexio@private.biz"
	phone_fixed := "+41 44 111 11 11"
	phone_fixed_second := "+41 44 222 22 22"
	phone_mobile := "+41 79 333 33 33"
	fax := "+41 44 333 33 33"
	url := "https://www.webgate.biz"
	skype_name := "webgate@skype.com"
	remarks := "Created with Terraform Test"
	language_id := 1
	contact_group_ids := ""
	contact_branch_ids := ""
	user_id := 1
	owner_id := 1
	salutation_id := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContactResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccContactCreateConfigBasic(lastname, firstname, birthday, address, postcode, city, country_id, email, emailprivate, phone_fixed, phone_fixed_second, phone_mobile, fax, url, skype_name, remarks, language_id, contact_group_ids, contact_branch_ids, user_id, owner_id, salutation_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bexio_contact.create", "name_1", lastname),
					resource.TestCheckResourceAttr("bexio_contact.create", "name_2", firstname),
					resource.TestCheckResourceAttr("bexio_contact.create", "birthday", birthday),
					resource.TestCheckResourceAttr("bexio_contact.create", "address", address),
					resource.TestCheckResourceAttr("bexio_contact.create", "postcode", postcode),
					resource.TestCheckResourceAttr("bexio_contact.create", "city", city),
					resource.TestCheckResourceAttr("bexio_contact.create", "mail", email),
					resource.TestCheckResourceAttr("bexio_contact.create", "mail_second", emailprivate),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_fixed", phone_fixed),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_fixed_second", phone_fixed_second),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_mobile", phone_mobile),
					resource.TestCheckResourceAttr("bexio_contact.create", "fax", fax),
					resource.TestCheckResourceAttr("bexio_contact.create", "url", url),
					resource.TestCheckResourceAttr("bexio_contact.create", "skype_name", skype_name),
					resource.TestCheckResourceAttr("bexio_contact.create", "remarks", remarks),
				),
			},
			{
				ResourceName:      "bexio_contact.create",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccUpdateContact(t *testing.T) {
	lastname := "Bexio"
	lastname2 := "WGC"
	firstname := "WGC"
	firstname2 := "Updated"
	birthday := "2003-09-23"
	address := "Riedstrasse 3"
	postcode := "8953"
	city := "Dietikon"
	country_id := 1
	email := "wgc.bexio@webgate.biz"
	emailprivate := "wgc.bexio@private.biz"
	phone_fixed := "+41 44 111 11 11"
	phone_fixed_second := "+41 44 222 22 22"
	phone_mobile := "+41 79 333 33 33"
	fax := "+41 44 333 33 33"
	url := "https://www.webgate.biz"
	skype_name := "webgate@skype.com"
	remarks := "Created with Terraform Test"
	language_id := 1
	contact_group_ids := ""
	contact_branch_ids := ""
	user_id := 1
	owner_id := 1
	salutation_id := 1

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccContactCreateConfigBasic(lastname, firstname, birthday, address, postcode, city, country_id, email, emailprivate, phone_fixed, phone_fixed_second, phone_mobile, fax, url, skype_name, remarks, language_id, contact_group_ids, contact_branch_ids, user_id, owner_id, salutation_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bexio_contact.create", "name_1", lastname),
					resource.TestCheckResourceAttr("bexio_contact.create", "name_2", firstname),
					resource.TestCheckResourceAttr("bexio_contact.create", "birthday", birthday),
					resource.TestCheckResourceAttr("bexio_contact.create", "address", address),
					resource.TestCheckResourceAttr("bexio_contact.create", "postcode", postcode),
					resource.TestCheckResourceAttr("bexio_contact.create", "city", city),
					resource.TestCheckResourceAttr("bexio_contact.create", "mail", email),
					resource.TestCheckResourceAttr("bexio_contact.create", "mail_second", emailprivate),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_fixed", phone_fixed),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_fixed_second", phone_fixed_second),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_mobile", phone_mobile),
					resource.TestCheckResourceAttr("bexio_contact.create", "fax", fax),
					resource.TestCheckResourceAttr("bexio_contact.create", "url", url),
					resource.TestCheckResourceAttr("bexio_contact.create", "skype_name", skype_name),
					resource.TestCheckResourceAttr("bexio_contact.create", "remarks", remarks),
				),
			},
			{
				Config: testAccContactCreateConfigBasic(lastname2, firstname2, birthday, address, postcode, city, country_id, email, emailprivate, phone_fixed, phone_fixed_second, phone_mobile, fax, url, skype_name, remarks, language_id, contact_group_ids, contact_branch_ids, user_id, owner_id, salutation_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("bexio_contact.create", "name_1", lastname2),
					resource.TestCheckResourceAttr("bexio_contact.create", "name_2", firstname2),
					resource.TestCheckResourceAttr("bexio_contact.create", "birthday", birthday),
					resource.TestCheckResourceAttr("bexio_contact.create", "address", address),
					resource.TestCheckResourceAttr("bexio_contact.create", "postcode", postcode),
					resource.TestCheckResourceAttr("bexio_contact.create", "city", city),
					resource.TestCheckResourceAttr("bexio_contact.create", "mail", email),
					resource.TestCheckResourceAttr("bexio_contact.create", "mail_second", emailprivate),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_fixed", phone_fixed),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_fixed_second", phone_fixed_second),
					resource.TestCheckResourceAttr("bexio_contact.create", "phone_mobile", phone_mobile),
					resource.TestCheckResourceAttr("bexio_contact.create", "fax", fax),
					resource.TestCheckResourceAttr("bexio_contact.create", "url", url),
					resource.TestCheckResourceAttr("bexio_contact.create", "skype_name", skype_name),
					resource.TestCheckResourceAttr("bexio_contact.create", "remarks", remarks),
				),
			},
		},
	})
}
func testAccContactCreateConfigBasic(v1, v2, v3, v4, v5, v6 string, v7 int, v8, v9, v10, v11, v12, v13, v14, v15, v16 string, v17 int, v18, v19 string, v20, v21, v22 int) string {
	return fmt.Sprintf(`
	resource "bexio_contact" "create" {
		name_1             = "%s"
		name_2             = "%s"
		birthday           = "%s"
		address            = "%s"
		postcode           = "%s"
		city               = "%s"
		country_id         = "%d"
		mail               = "%s"
		mail_second        = "%s"
		phone_fixed        = "%s"
		phone_fixed_second = "%s"
		phone_mobile       = "%s"
		fax                = "%s"
		url                = "%s"
		skype_name         = "%s"
		remarks            = "%s"
		language_id        = "%d"
		contact_group_ids  = "%s"
		contact_branch_ids = "%s"
		user_id            = "%d"
		owner_id           = "%d"
		salutation_id      = "%d"
	  }
	`, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22)
}
func testAccCheckContactResourceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.API)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "bexio_contact" {
			continue
		}

		_, err := c.ReadContact(rs.Primary.ID)

		if err == nil {
			return fmt.Errorf("Contact user still exists")
		}
		notFoundErr := "Page not found"
		expectedErr := regexp.MustCompile(notFoundErr)
		if !expectedErr.Match([]byte(err.Error())) {
			return fmt.Errorf("expected %s, got %s", notFoundErr, err)
		}
	}
	return nil
}
