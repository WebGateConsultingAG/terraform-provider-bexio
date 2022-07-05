package provider

import (
	"context"

	client "wgc_bexio_provider/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	defaultBaseURL = "https://api.bexio.com"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BEXIO_URL", nil),
				Description: "The url of the Bexio API which should be used",
			},
			"api_token": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("BEXIO_TOKEN", nil),
				Description: "The personal access token which should be used.",
				Sensitive:   true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"bexio_contact":       resourceCreateContact(),
			"bexio_fictionaluser": resourceFictionalUser(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"bexio_contacts":          dataSourceContacts(),
			"bexio_contact":           dataSourceSearchContact(),
			"bexio_salutation":        dataSourceSearchSalutation(),
			"bexio_language":          dataSourceSearchLanguage(),
			"bexio_country":           dataSourceSearchCountry(),
			"bexio_fictionalusers":    dataSourceFictionalUsers(),
			"bexio_fictionaluserbyid": dataSourceFictionalUserById(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("api_token").(string)
	url := d.Get("url").(string)
	var diags diag.Diagnostics
	if url == "" {
		url = defaultBaseURL
	}

	c, err := client.NewClient(url, token)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Bexio client",
			Detail:   "Unable to validate api token",
		})
		return nil, diags
	}

	return c, diags
}
