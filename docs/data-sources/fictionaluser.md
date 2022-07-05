---
page_title: "bexio_fictionaluserbyid Data Source - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_fictionaluserbyid data source allows you to fetch a fictional user by id. 
---

# Data Source `bexio_fictionaluserbyid`
The bexio_fictionaluserbyid data source allows you to fetch a fictional user by id. 

## Example Usage
### Data Block
```terraform
data "bexio_fictionaluserbyid" "this" {
  id = "100"
}
```

## Argument Reference
- `id` (string) Id of the fictional user.

## Attributes Reference
The following attributes are exported.
- `id` - `int`
- `salutation_type` - `string`
- `firstname` - `string`
- `lastname` - `string`
- `email` - `string`
- `title_id` - `int` or `null`

## Reference
- [Bexio-API](https://docs.bexio.com/)
- [Bexio Fictional User](https://docs.bexio.com/#tag/User-Management/operation/v3ShowFictionalUser)