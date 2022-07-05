package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSearchSalutation() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSearchSalutationRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func dataSourceSearchSalutationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	salutationName := d.Get("name").(string)
	resp, err := c.Search(salutationName, "name", "=", "/2.0/salutation/search")
	if err != nil {
		return diag.FromErr(err)
	}
	if len(resp) == 0 {
		return diag.FromErr(fmt.Errorf("ERROR: No salutation " + salutationName + " found"))
	}
	d, diags = client.SetFields(resp[0], d)
	return diags
}
