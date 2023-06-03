
resource "aws_s3_bucket" "sample_bucket" {
  bucket = "ikari-s3-bucket-sample"
}

resource "aws_s3_bucket_ownership_controls" "firelens" {
  bucket = aws_s3_bucket.sample_bucket.bucket

  rule {
    object_ownership = "BucketOwnerEnforced"
  }
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.sample_bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "sample/"
  }
}
