provider "aws" {
  region = "us-east-1"
}

resource "aws_instance" "ec2-instance" {
  ami           = "ami-test"
  instance_type = "t2.micro"
}