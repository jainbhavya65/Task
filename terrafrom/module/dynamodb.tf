resource "aws_dynamodb_table" "this" {
  name           = "people1"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "Id"

  attribute {
    name = "Id"
    type = "N"
  }
}