resource "aws_iam_role" "lambda_role" {
  name               = "go-lambda-slack-notifier-role"
  assume_role_policy = file("policies/lambda-assume-role.json")
}

resource "aws_iam_policy" "lambda_policy" {
  name   = "go-lambda-slack-notifier-policy"
  policy = file("policies/lambda-policy.json")
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attachment" {
  role       = aws_iam_role.lambda_role.name
  policy_arn = aws_iam_policy.lambda_policy.arn
}
