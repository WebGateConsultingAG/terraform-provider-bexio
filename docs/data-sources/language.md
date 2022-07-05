---
page_title: "bexio_language Data Source - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_language data source lets you fetch the language-id via an iso_639_1 language string.
---

# Data Source `bexio_language`
The bexio_language data source allows you to get the bexio language-id. 

## Example Usage
```terraform
data "bexio_language" "this" {
    language = "fr"
}
```
## Argument Reference
- `language` The language of which you want the id. -> <iso_639_1>

## Attributes Reference
The following attributes are exported.
* `id` - Id of the bexio language
* `date_format` - "d.m.Y"
* `decimal_point` - i.e. "."
* `iso_639_1` - iso code for language
* `language`
* `name` - Fullname of language
* `thousands_separator` - i.e.  "'"

## Reference
- [Bexio-API](https://docs.bexio.com/)
- [Bexio Search Language](https://docs.bexio.com/#tag/Languages/operation/v2SearchLanguages)