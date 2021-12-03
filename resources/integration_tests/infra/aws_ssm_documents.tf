resource "aws_ssm_document" "ssm_document" {
  name          = "${var.test_prefix}doc${var.test_suffix}"
  document_type = "Command"

  content = <<DOC
  {
    "schemaVersion": "1.2",
    "description": "Check ip configuration of a Linux instance.",
    "parameters": {

    },
    "runtimeConfig": {
      "aws:runShellScript": {
        "properties": [
          {
            "id": "0.aws:runShellScript",
            "runCommand": ["ifconfig"]
          }
        ]
      }
    }
  }
DOC

  tags = {
    Name = "${var.test_prefix}doc${var.test_suffix}"
  }
}