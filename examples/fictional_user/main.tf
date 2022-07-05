terraform {
  required_providers {
    bexio = {
      source = "WebGateConsultingAG/bexio"
    }
  }
}

/* data "bexio_fictionalusers" "this" {
}
output "bexio_fictionalusers_output" {
  description = "Bexio FictionalUsers"
  value       = data.bexio_fictionalusers.this
}
output "bexio_fictionalusers_output_sub" {
  description = "Bexio FictionalUsers"
  value       = data.bexio_fictionalusers.this.fictionalusers[0].email
} */
data "bexio_fictionaluserbyid" "this" {
  id = 9
}
output "fictionaluserbyid_output" {
  description = "Bexio FictionalUser By Id"
  value       = data.bexio_fictionaluserbyid.this
}
resource "bexio_fictionaluser" "create" {
    salutation_type             = "male"
    firstname             = "Klaus"
    lastname           = "Fictional2"
    email            = "klaus.fictional2@webgate.biz"
}