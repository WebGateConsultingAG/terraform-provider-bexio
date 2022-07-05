package provider

import (
	"context"
	"fmt"
	client "wgc_bexio_provider/client"
	"wgc_bexio_provider/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCreateContact() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceReadContact,
		CreateContext: resourceContactCreate,
		DeleteContext: resourceDeleteContact,
		UpdateContext: resourceUpdateContact,
		Schema: map[string]*schema.Schema{
			"contact_type_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name_1": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name_2": {
				Type:     schema.TypeString,
				Required: true,
			},
			"salutation_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"salutation_form": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"title_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"birthday": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"postcode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"city": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"country_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"mail": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mail_second": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone_fixed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone_fixed_second": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"phone_mobile": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fax": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skype_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remarks": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"language_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"is_lead": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"contact_group_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"contact_branch_ids": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"user_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"owner_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"updated_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"profile_image": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceContactCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	contactEmail := d.Get("mail").(string)
	contact := extractContact(d)

	search, err := c.Search(contactEmail, "mail", "=", "/2.0/contact/search/")
	if err != nil {
		return diag.FromErr(err)
	}
	if len(search) != 0 {
		return diag.FromErr(fmt.Errorf("ERROR: There's already a user with email " + contactEmail))
	}

	resp, err := c.Create(contact, "/2.0/contact/")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(resp["id"]))
	d.Set("nr", fmt.Sprint(resp["nr"]))
	return diags
}
func resourceReadContact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	contactId := d.Get("id").(string)

	contact, err := c.ReadContact(contactId)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("%f", contact["id"]))
	d, diags = client.SetFields(contact, d)
	return diags
}

func resourceDeleteContact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	var diags diag.Diagnostics
	contactId := d.Get("id").(string)

	err := c.Delete(contactId, "/2.0/contact/")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return diags
}

func resourceUpdateContact(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*client.API)
	contactId := d.Get("id").(string)
	contact := extractContact(d)

	d.Partial(true)
	if d.HasChange("mail") {
		contactEmail := d.Get("mail").(string)
		search, err := c.Search(contactEmail, "mail", "=", "/2.0/contact/search/")
		if err != nil {
			return diag.FromErr(err)
		}
		if len(search) != 0 {
			return diag.FromErr(fmt.Errorf("ERROR: There's already a user with email " + contactEmail))
		}
	}

	resp, err := c.Update(contactId, contact, "POST", "/2.0/contact/")
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprint(resp["id"]))
	d.Set("nr", fmt.Sprint(resp["nr"]))
	d.Partial(false)
	return resourceReadContact(ctx, d, m)
}
func extractContact(d *schema.ResourceData) models.ContactModel {
	contact := models.ContactModel{
		Contact_type_id:    2,
		Name_1:             d.Get("name_1").(string),
		Name_2:             d.Get("name_2").(string),
		Salutation_id:      1,
		Salutation_form:    d.Get("salutation_form").(string),
		Titel_id:           d.Get("title_id").(int),
		Birthday:           d.Get("birthday").(string),
		Address:            d.Get("address").(string),
		Postcode:           d.Get("postcode").(string),
		City:               d.Get("city").(string),
		Country_id:         d.Get("country_id").(int),
		Mail:               d.Get("mail").(string),
		Mail_second:        d.Get("mail_second").(string),
		Phone_fixed:        d.Get("phone_fixed").(string),
		Phone_fixed_second: d.Get("phone_fixed_second").(string),
		Phone_mobile:       d.Get("phone_mobile").(string),
		Fax:                d.Get("fax").(string),
		Url:                d.Get("url").(string),
		Skype_name:         d.Get("skype_name").(string),
		Remarks:            d.Get("remarks").(string),
		Language_id:        d.Get("language_id").(int),
		Contact_group_ids:  d.Get("contact_group_ids").(string),
		Contact_branch_ids: d.Get("contact_branch_ids").(string),
		User_id:            d.Get("user_id").(int),
		Owner_id:           d.Get("owner_id").(int),
	}
	return contact
}
