terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

variable "do_token" {}
variable "pvt_key" {}
variable "pub_key" {}
variable "key_name" {}

provider "digitalocean" {
  token = var.do_token
}

data "digitalocean_ssh_key" "do_ssh_key" {
  name = var.key_name
}