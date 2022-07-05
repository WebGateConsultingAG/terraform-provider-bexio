---
page_title: "bexio_fictionalusers Data Source - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_fictionalusers data source allows you list all your fictional users. 
---

# Data Source `bexio_fictionalusers`
The bexio_fictionalusers data source allows you list all your fictional users. 

## Example Usage
### Data Block
```terraform
data "bexio_fictionalusers" "this" {}
```


## Attributes Reference
The following attributes are exported as an array of objects.
- `id` - `int`
- `salutation_type` - `string`
- `firstname` - `string`
- `lastname` - `string`
- `email` - `string`
- `title_id` - `int` or `null`

## Reference
- [Bexio-API](https://docs.bexio.com/)
- [Bexio List Fictional Users](https://docs.bexio.com/#tag/User-Management/operation/v3ListFictionalUsers)