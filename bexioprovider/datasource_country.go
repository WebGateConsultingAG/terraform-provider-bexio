package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSearchCountry() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSearchCountryRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_short": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iso_3166_alpha2": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSearchCountryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	country := d.Get("name").(string)

	resp, err := c.Search(country, "name", "=", "/2.0/country/search/")
	if err != nil {
		return diag.FromErr(err)
	}
	if len(resp) == 0 {
		return diag.FromErr(fmt.Errorf("ERROR: No country with name " + country + " found"))
	}

	d, diags = client.SetFields(resp[0], d)
	return diags
}
