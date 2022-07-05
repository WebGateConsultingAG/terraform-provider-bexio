---
page_title: "bexio_contact Resource - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_contact resource lets you create a Contact.

---

# Resource `bexio_contact`
The bexio_contact resource allows you to perform CRUD-Operations on bexio contacts.

## Example usage
### Data and Resource Blocks
- `contact_email` - this field referes to an existing bexio contact. For clarification, the new contact is subordinated to that contact.


```terraform
data "bexio_contact" "this" {
  contact_email = "support@bexio.com"
}
data "bexio_salutation" "this" {
  name = "Herr" 
}
data "bexio_language" "this" {
  language = "de"  
}
data "bexio_country" "this" {
  country = "Schweiz" 
}
resource "bexio_contact" "create" {
  name_1             = "firstname"
  name_2             = "lastname"
  birthday           = "1975-09-23"
  address            = "Riedstrasse 3"
  postcode           = "8953"
  city               = "Dietikon"
  country_id         = data.bexio_country.this.id
  mail               = "wgc.contact@webgate.biz"
  mail_second        = "wgc.contact@external.biz"
  phone_fixed        = "+41 44 444 44 44"
  phone_fixed_second = "+41 44 444 44 55"
  phone_mobile       = "+41 79 333 33 33"
  fax                = "+41 44 333 33 33"
  url                = "https://www.webgate.biz"
  skype_name         = "webgate@skype.com"
  remarks            = "Created with Terraform Test"
  language_id        = data.bexio_language.this.id
  contact_group_ids  = ""
  contact_branch_ids = ""
  user_id            = data.bexio_contact.this.id
  owner_id           = data.bexio_contact.this.id
  salutation_id      = data.bexio_salutation.this.id
}
```

## Argument Reference
All not required attributes can be `null`.
The following arguments are supported:

* `name_1` - (Required) lastname
* `name_2` -  (Required) firstname
* `birthday` - (Optional) date, date-format `yyyy-mm-dd`
* `address` - (Optional) address
* `postcode` - (Optional) postcode
* `city` - (Optional) city
* `country_id` - we fetch `int` via [country search](../data-sources/country.md)-call
* `mail` - (Optional) email
* `mail_second` - (Optional) secondary email
* `phone_fixed` - (Optional) phonefixed
* `phone_fixed_second` - (Optional) phone fixed secondary
* `phone_mobile` - (Optional) phone mobile
* `fax` - (Optional) fax
* `url` - (Optional) url
* `skype_name` - (Optional) Skype name
* `remarks` - (Optional) remarks
* `language_id` -  we fetch `int` via [language search](../data-sources/language.md)-call
* `contact_group_ids`
* `contact_branch_ids`
* `user_id` - (Required) we fetch `int` via [contact search](../data-sources/contact.md)-call
* `owner_id` - (Required) we fetch `int` via [contact search](../data-sources/contact.md)-call
* `salutation_id` - we fetch `int` via [salutation search](../data-sources/salutation.md)-call


## Attribute Reference
### Contact
* `id` - The resource ID in terraform of contact.
* `nr`
* `contact_type_id` - `int`, will always be 2 for a contact
* `is_lead` - `boolean`
* `updated_at`
* `profile_image`

## Import

A Bexio Contact can be imported using the id, e.g.

```
$ terraform import bexio_contact.example <id>
```

## Referece
* [Bexio-API](https://docs.bexio.com/)
* [Bexio-API create contact section](https://docs.bexio.com/#tag/Contacts/operation/v2CreateContact)
* [Bexio Language Reference](https://docs.bexio.com/#tag/Languages/operation/v2SearchLanguages)


