package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSearchLanguage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSearchLanguageRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_format_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"decimal_point": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"thousands_separator": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_639_1": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSearchLanguageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	languageName := d.Get("language").(string)

	resp, err := c.Search(languageName, "iso_639_1", "=", "/2.0/language/search")
	if err != nil {
		return diag.FromErr(err)
	}
	if len(resp) == 0 {
		return diag.FromErr(fmt.Errorf("ERROR: No language with iso " + languageName + " found"))
	}

	d, diags = client.SetFields(resp[0], d)
	return diags
}
