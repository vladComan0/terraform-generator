provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "EC2Instance1" {
  ami             = "ami-12345678"
  instance_type   = "t2.micro"
  key_name        = "my-keypair"
  security_groups = ["sg-12345678"]
  tags = {
    "Environment" = "Development"
    "Project" = "MyProject"
  }
}
resource "aws_s3_bucket" "MyS3Bucket" {
  bucket        = "my-bucket"
  acl           = "private"
  region        = "us-west-1"
  force_destroy = true
  tags = {
    "Environment" = "Acceptance"
    "Project" = "MyProject"
  }
}