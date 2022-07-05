package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSearchContact() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSearchContactRead,
		Schema: map[string]*schema.Schema{
			"contact_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact_type_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name_1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"salutation_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"salutation_form": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"title_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"birthday": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"postcode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mail": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mail_second": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_fixed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_fixed_second": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone_mobile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fax": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"skype_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remarks": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"language_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_lead": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"contact_group_ids": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact_branch_ids": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"owner_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
func dataSourceSearchContactRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	contactEmail := d.Get("contact_email").(string)
	resp, err := c.Search(contactEmail, "mail", "=", "/2.0/contact/search/")
	if err != nil {
		return diag.FromErr(err)
	}
	if len(resp) == 0 {
		return diag.FromErr(fmt.Errorf("ERROR: No contact with email " + contactEmail + " found"))
	}

	d, diags = client.SetFields(resp[0], d)
	return diags

}
