package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"
	"wgc_bexio_provider/models"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFictionalUser() *schema.Resource {
	return &schema.Resource{
		ReadContext:   dataSourceFictionalUserByIdRead,
		CreateContext: resourceFictionalUserCreate,
		DeleteContext: resourceDeleteFictionalUser,
		UpdateContext: resourceUpdateFictionalUser,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"salutation_type": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: isMaleOrFemale,
			},
			"firstname": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lastname": {
				Type:     schema.TypeString,
				Required: true,
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFictionalUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	fictionalUser := extractFictionalUser(d)

	resp, err := c.Create(fictionalUser, "/3.0/fictional_users/")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(resp["id"]))
	d.Set("email", fmt.Sprint(resp["email"]))
	return diags
}

func resourceDeleteFictionalUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	fictionalUserId := d.Get("id").(string)

	err := c.Delete(fictionalUserId, "/3.0/fictional_users/")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func resourceUpdateFictionalUser(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	fictionalUserId := d.Get("id").(string)
	fictionalUser := extractFictionalUser(d)

	d.Partial(true)
	resp, err := c.Update(fictionalUserId, fictionalUser, "PATCH", "/3.0/fictional_users/")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(resp["id"]))
	d.Set("email", fmt.Sprint(resp["email"]))
	d.Partial(false)
	return dataSourceFictionalUserByIdRead(ctx, d, m)
}
func extractFictionalUser(d *schema.ResourceData) models.FictionalUserModel {
	fictionalUser := models.FictionalUserModel{
		Salutation_type: d.Get("salutation_type").(string),
		Firstname:       d.Get("firstname").(string),
		Lastname:        d.Get("lastname").(string),
		Email:           d.Get("email").(string),
	}
	return fictionalUser
}
func isMaleOrFemale(v interface{}, p cty.Path) diag.Diagnostics {
	salutation := v.(string)
	if salutation != "male" && salutation != "female" {
		return diag.Diagnostics{
			diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "wrong value",
				Detail:   "is not 'male' or 'female', got: " + salutation,
			},
		}
	} else {
		return diag.Diagnostics{}
	}
}
