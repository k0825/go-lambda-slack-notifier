resource "aws_lambda_function" "lambda" {
  function_name    = "go-lambda-slack-notifier"
  filename         = "./lambda/archive/main.zip"
  role             = aws_iam_role.lambda_role.arn
  handler          = "main"
  runtime          = "go1.x"
  source_code_hash = data.archive_file.lambda.output_base64sha256
}

resource "null_resource" "default" {
  triggers = {
    always_run = timestamp()
  }
  provisioner "local-exec" {
    command = "cd ./lambda/src/go-lambda-slack-notifier && GOOS=linux GOARCH=amd64 go build -o ../../build/main main.go"
  }
}

data "archive_file" "lambda" {
  type        = "zip"
  source_file = "./lambda/build/main"
  output_path = "./lambda/archive/main.zip"

  depends_on = [null_resource.default]
}
