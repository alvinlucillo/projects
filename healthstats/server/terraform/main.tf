# The empty block is also called partial configuration which
#  inform Terraform that backend configuration is defined dynamically.
#  This project uses -backend-configuration (see Makefile)
terraform {
  backend "s3" {}
}

provider "aws" {
  region = var.region
}

resource "aws_s3_bucket" "s3_buckets" {
  for_each = var.s3_map

  bucket   = "${each.key}"
}

resource "aws_s3_bucket_versioning" "s3_bucket_versioning" {
  for_each = aws_s3_bucket.s3_buckets

  bucket             = each.value.id
  versioning_configuration {
    status = var.s3_map[each.key].versioning ? "Enabled" : "Suspended"
  } 
}