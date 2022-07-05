[![Bexio-Provider](https://img.shields.io/badge/%20Bexio--Provider-0.5.0-blue)](https://github.com/WebGateConsultingAG/terraform-provider-bexio) [![Go](https://img.shields.io/badge/Go-v1.18.2-blue)](https://go.dev/doc/) [![PkgGoDev](https://pkg.go.dev/badge/github.com/hashicorp/terraform-plugin-sdk/v2)](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk/v2)
# WGC BEXIO PROVIDER # 

WGC bexio provider is a custom terraform provider to interact with [Bexio](https://www.bexio.com) resources.  
Right now only "Contacts" and "Fictional Users" maybe generated.

### Table of Contents
- [Descripton](#descripton)
- [Technologies](#technologies)
- [Setup](#setup)
- [How to](#how-to)
- [How to contribute](#how-to-contribute)

## Descripton
With this provider you can interact with the Bexio [API](https://docs.bexio.com/). You can search for users via their email address, and you can read all users from your Organisation. This provider helps you create a contact or fictional user on bexio without going through the create new user wizard on the bexio website.
As basis for our code we used [Go](https://go.dev/doc/) and [terraform](https://www.terraform.io/docs). Very good help is the following tutorial [Perform CRUD Operations with Providers](https://learn.hashicorp.com/tutorials/terraform/provider-use). 
## Technologies
Project is created with:
- terraform version: v1.2.2
- Go version: go1.18.3
- hashicorps terraform-plugin-sdk version: v2
Imported packages:
- github.com/hashicorp/terraform-plugin-sdk/v2/diag
- github.com/hashicorp/terraform-plugin-sdk/v2/helper
## Setup
1. Customize the Makefile according to your OS Architecture
2. Delete the go.mod file
3. ```
   $ go mod init wgc_bexio_provider
   $ go mod tidy
   $ make install
   ```
 Type these three commands in a terminal in your Project folder. &#8594; this will build the provider and moves it to a location where it can be used by terraform.

4. Test the provider by executing `$ terraform init ` &#8594; this will initialize the provider for terraform, then to Test it run the following command ` $ terraform plan ` and to apply it run `$ terraform apply `. If everything worked as expected delete the resource with the command `$ terraform destroy `.

## Acceptance Testing
Run all Tests via:  
`make testacc TEST=./bexioprovider`  
or a single test with:  
`make testacc TEST=./bexioprovider TESTARGS="-run=TestAccUpdateFictionalUser"`

## How To
For a tutorial/documentation of the provider go to [Bexio Hashicorp](inset link for hashicorp site where docu iss)

## How to Contribute
Please read the [CONTRIBUTION](CONTRIBUTION.md)-File
  
    

## Initial Contributors - first Release
* [David Buergler](https://github.com/Davy42)
* [Klaus Bild](https://github.com/kbild)
* [Nathan Fahrni](https://github.com/NKFahrni)