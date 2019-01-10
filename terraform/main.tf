variable "region" {}
variable "owner" {}
variable "identifier" {}
variable "environment" {}

provider "aws" {
  region	      = "${var.region}"
  allowed_account_ids = ["724781030999"]
}

resource "aws_s3_bucket" "sql_files" {
  acl		= "private"
  bucket_prefix = "${format("%s-sql_files-", var.identifier)}"
  force_destroy = true

  tags {
    Owner       = "${var.owner}"
    Environment = "${var.environment}"
  }
}


