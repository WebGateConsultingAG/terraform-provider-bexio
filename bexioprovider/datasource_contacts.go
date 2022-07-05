package provider

import (
	"context"
	"strconv"
	"time"
	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceContacts() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceContactsRead,
		Schema: map[string]*schema.Schema{
			"contacts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
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
				},
			},
		},
	}
}
func dataSourceContactsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics

	resp, err := c.ListContacts()
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("contacts", resp); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
