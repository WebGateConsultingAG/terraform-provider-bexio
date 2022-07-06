terraform {
  required_providers {
    bexio = {
      source = "WebGateConsultingAG/bexio"d
    }
  }
}

data "bexio_contacts" "this" {
}
output contacts {
  value       = data.bexio_contacts.this
  description = "description"
}
data "bexio_contact" "this" {
  contact_email = "support@bexio.com"
}
output contact {
  value       = data.bexio_contact.this
  description = "description"
}
data "bexio_salutation" "this" {
  name = "Herr" 
}
output salutation {
  value       = data.bexio_salutation.this
  description = "description"
}
data "bexio_language" "this" {
  language = "fr" 
}
output language {
  value       = data.bexio_language.this
  description = "description"
}
data "bexio_country" "this" {
  name = "Schweiz" 
}
output country {
  value       = data.bexio_country.this
  description = "description"
}
resource "bexio_contact" "create" {
    name_1             = "Test"
    name_2             = "User"
    birthday           = "1975-09-23"
    address            = "Bahnhofsstrasse 10"
    postcode           = "8000"
    city               = "Zurich"
    country_id         = data.bexio_country.this.id
    mail               = "testuser@webgate.biz"
    mail_second        = "test.user@webgate.biz"
    phone_fixed        = "+41 44 444 44 44"
    phone_fixed_second = ""
    phone_mobile       = ""
    fax                = ""
    url                = ""
    skype_name         = ""
    remarks            = "Created with Terraform2"
    language_id        = data.bexio_language.this.id
    contact_group_ids  = ""
    contact_branch_ids = ""
    user_id            = data.bexio_contact.this.id
    owner_id           = data.bexio_contact.this.id
    salutation_id      = data.bexio_salutation.this.id
}
