
resource "aws_s3_bucket" "sample_bucket" {
  bucket = "ikari-s3-bucket-sample"
  acl    = "private"
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.sample_bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.sample_lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "sample/"
  }
}
