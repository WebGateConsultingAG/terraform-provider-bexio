---
page_title: "bexio_contacts Data Source - terraform-provider-bexio"
subcategory: ""
description: |-
  The bexio_contact data source allows you list all your contacts. 
---

# Data Source `bexio_contacts`
The bexio_contact data source allows you list all your contacts. 

## Example Usage
```terraform
data "bexio_contacts" "this" {}
```

## Attributes Reference
The following attributes are exported as contacts list.
* `id` - The resource ID in terraform of contact.
* `nr`
* `contact_type_id` - `int`, will always be 2 for a contact
* `is_lead` - `boolean`
* `updated_at`
* `profile_image`
* `name_1` - lastname
* `name_2` -  firstname
* `birthday` - date, date-format `yyyy-mm-dd`
* `address` - address
* `postcode` - postcode
* `city` - city
* `country_id` - country id
* `mail` - email
* `mail_second` - secondary email
* `phone_fixed` - phonefixed
* `phone_fixed_second` - phone fixed secondary
* `phone_mobile` - phone mobile
* `fax` - fax
* `url` - url
* `skype_name` - Skype name
* `remarks` - remarks
* `language_id` -  language id
* `contact_group_ids`
* `contact_branch_ids`
* `user_id` - user id
* `owner_id` - owner id
* `salutation_id` - salutation id


## Reference
- [Bexio-API](https://docs.bexio.com/)