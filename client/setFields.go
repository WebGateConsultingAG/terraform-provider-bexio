package bexioclient

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SetFields(data map[string]interface{}, d *schema.ResourceData) (*schema.ResourceData, diag.Diagnostics) {
	var diags diag.Diagnostics
	for k, v := range data {
		if k != "id" {
			str, ok := v.(string)
			if ok {
				err := d.Set(k, str)
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Unable to set field as string: " + k,
						Detail:   "Unable to set field as string: " + k,
					})
				}
			} else {
				err := d.Set(k, v)
				if err != nil {
					diags = append(diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Unable to set field as integer: " + k,
						Detail:   "Unable to set field as integer:" + k,
					})
				}
			}
		} else {
			d.SetId(fmt.Sprint(v))
		}
	}
	return d, diags
}
