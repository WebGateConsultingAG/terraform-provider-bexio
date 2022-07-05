---
page_title: "bexio_fictionaluser Resource - terraform-provider-bexio"
subcategory: ""
description: |-
The bexio_fictionaluser resource lets you create a Fictional User
  
---

# Resource `bexio_fictionaluser`
The bexio_fictionaluser resource allows you to perform CRUD-Operations on bexio fictional users.

## Example usage
### Data and Resource Blocks


```terraform
resource "bexio_fictionaluser" "create" {
  salutation_type = "male"
  firstname       = "David"
  lastname        = "Buergler"
  email           = "david.buergler@webgate.biz"
}

```

## Argument Reference
### Create contact
All not required attributes can be `null`.
The following arguments are supported:
* `salutation_type` - (Required), must be "female" or "male"
* `firstname` - (Required) firstname 
* `lastname` - (Required) lastname 
* `email` - address 



## Attribute Reference
### Contact
* `id` -  The resource ID in terraform of fictional user.

## Import

A Bexio Fictional User can be imported using the id, e.g.

```
$ terraform import bexio_fictionaluser.example <id>
```

## Referece
* [Bexio-API](https://docs.bexio.com/)
* [Bexio-API create fictional user section](https://docs.bexio.com/#tag/User-Management/operation/v3CreateFictionalUser)


