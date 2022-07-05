package provider

import (
	"context"
	"strconv"
	"time"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFictionalUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFictionalUsersRead,
		Schema: map[string]*schema.Schema{
			"fictionalusers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
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
							Computed: true,
						},
						"title_id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}
func dataSourceFictionalUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics

	resp, err := c.ListFictionalUsers()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("fictionalusers", resp); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
