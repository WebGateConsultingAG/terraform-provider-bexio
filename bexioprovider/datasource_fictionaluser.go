package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFictionalUserById() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFictionalUserByIdRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"salutation_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firstname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lastname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}
func dataSourceFictionalUserByIdRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	id := d.Get("id").(string)
	resp, err := c.ReadFictionalUser(id)
	if err != nil {
		return diag.FromErr(err)
	}
	if len(resp) == 0 {
		return diag.FromErr(fmt.Errorf("ERROR: No fictional user with id %v  found", id))
	}

	d, diags = client.SetFields(resp, d)
	return diags
}
