data "aws_caller_identity" "current" {}
data "archive_file" "createRaceResults" {
  type = "zip"

  source_file = "${path.module}/main"
  output_path = "${path.module}/main.zip"
}

resource "aws_lambda_function" "S3ToDyanamodb" {
function_name    = var.name
role             = aws_iam_role.this.arn
handler          = "main"
timeout          = 10
runtime          = "go1.x"
filename         = "module/main.zip"
environment {
   variables = {
    REGION = var.region
    TABLE_NAME = aws_dynamodb_table.this.name
    }
}
}

resource "aws_s3_bucket" "bucket" {
bucket = "cvsdatatolambda"
acl    = "private"
tags = {
  Name        = "csvdata"
 }
}

resource "aws_s3_bucket_notification" "aws-lambda-trigger" {
bucket = aws_s3_bucket.bucket.id
lambda_function {
lambda_function_arn = aws_lambda_function.S3ToDyanamodb.arn
events              = ["s3:ObjectCreated:*"]
}
}
resource "aws_lambda_permission" "test" {
statement_id  = "AllowS3Invoke"
action        = "lambda:InvokeFunction"
function_name = aws_lambda_function.S3ToDyanamodb.function_name
principal = "s3.amazonaws.com"
source_arn = "arn:aws:s3:::${aws_s3_bucket.bucket.id}"
}

output "arn" {
value = aws_lambda_function.S3ToDyanamodb.arn
}
