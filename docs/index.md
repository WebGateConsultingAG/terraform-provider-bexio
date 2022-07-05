---
layout: "bexio"
page_title: "WGC Bexio Provider"
description: |-
  The Bexio provider is used to interact with Bexio resources
---
# WGC Bexio Provider
The Bexio provider is used to interact with the [Bexio API](https://docs.bexio.com/) in
order to create Bexio resources. 

Use the navigation to the left to read about the available resources and data sources.

## Example Usage

```hcl
provider "bexio" {
  url = "${var.url}"
  api_token = "${var.api-token}"
}
```

~> Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this
file ever be committed to a public version control system. See [Environment Variables](#environment-variables) for a
better alternative.
  


## Authentication
* Static credentials

* Environment variables  
  

### Static credentials

Static credentials can be provided by adding `api-token` in-line in the
bexio provider block:

Usage:

```hcl
provider "bexio" {
  api_token = "${var.api-token}"
}
```

### Environment variables

You can provide your credentials via `BEXIO_TOKEN`
environment variable, representing your Bexio API secret token.  

```hcl
provider "bexio" {}
```

Usage:

```shell
$ export BEXIO_TOKEN="astokensecret"
$ terraform plan
```


## Importing resources

To import Bexio resources, you will need to know their id. You will need the [Browser interface for Bexio](https://office.bexio.com/) to easily find your resource ids.




## Schema

### Optional

- `host` (String)
- `password` (String, Sensitive)
- `username` (String)

## Argument Reference

* `url` - (Optional) The Bexio API URL which should be used, defaults to 'https://api.bexio.com'. 
It can also be sourced from the `BEXIO_URL` environment variable.

* `api_token` - (Optional) Your [api access token](https://docs.bexio.com/#section/Authentication/API-Tokens).
It can also be sourced from the `BEXIO_TOKEN` environment variable. 

## Testing

Credentials must be provided via the `BEXIO_URL` and `BEXIO_TOKEN` environment variables in order to run acceptance tests.