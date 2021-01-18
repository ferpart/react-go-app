variable "AWS_ACCESS_KEY_ID" {}
variable "AWS_SECRET_KEY" {}
variable "AWS_DEFAULT_REGION" {}


resource "github_actions_secret" "AWS_ACCESS_KEY_ID" {
    repository       = "react-go-app"
    secret_name      = "AWS_ACCESS_KEY_ID"
    plaintext_value  = var.AWS_ACCESS_KEY_ID
}

resource "github_actions_secret" "AWS_SECRET_ACCESS_KEY" {
    repository       = "react-go-app"
    secret_name      = "AWS_SECRET_ACCESS_KEY"
    plaintext_value  = var.AWS_SECRET_KEY
}

resource "github_actions_secret" "INSTANCE_ID" {
    repository       = "react-go-app"
    secret_name      = "INSTANCE_ID"
    plaintext_value  = aws_instance.web-server-instance.id
    depends_on = [ aws_instance.web-server-instance ]
}

resource "github_actions_secret" "AWS_DEFAULT_REGION" {
    repository       = "react-go-app"
    secret_name      = "AWS_DEFAULT_REGION"
    plaintext_value  = var.AWS_DEFAULT_REGION
}