---
page_title: "bexio_salutation Data Source - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_salutation data source lets you fetch the salutation-id via a string.
---

# Data Source `bexio_salutation`
The bexio_salutation data source allows you to get the bexio salutation-id. 

## Example Usage
```terraform
data "bexio_salutation" "this" {
    name = "Herr"
}
```
## Argument Reference
- `name` (String) The salutation of which you want the id.

## Attributes Reference
The following attributes are exported.
- `id` - Id of the bexio salutation

## Reference
- [Bexio-API](https://docs.bexio.com/)
- [Bexio Search Salutation](https://docs.bexio.com/#tag/Salutations/operation/v2ListSalutations)