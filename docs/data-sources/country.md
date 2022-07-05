---
page_title: "bexio_country Data Source - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_country data source lets you fetch the country-id via a string.
---

# Data Source `bexio_country`
The bexio_country data source allows you to get the bexio country-id. 

## Example Usage
```terraform
data "bexio_country" "this" {
    name = "Schweiz"
}
```
## Argument Reference
- `name` (String) The country of which you want the id.

## Attributes Reference
The following attributes are exported.
- `id` - Id of the bexio country

## Reference
- [Bexio-API](https://docs.bexio.com/)
- [Bexio Search country](https://docs.bexio.com/#tag/Countries/operation/v2ListCountries)