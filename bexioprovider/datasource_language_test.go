package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSearchLanguage(t *testing.T) {
	language := "fr"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccLanguageConfigBasic(language),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.bexio_language.this", "language", "fr"),
				),
			},
		},
	})
}
func testAccLanguageConfigBasic(s1 string) string {
	return fmt.Sprintf(`
		data "bexio_language" "this" {
			language        = "%s"
		}
		`, s1)
}
