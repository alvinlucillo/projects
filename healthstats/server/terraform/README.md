## Getting started
This simple project demonstrates basic operations in managing your AWS infrastructure using Terraform. More on this by visiting this Medium article:
https://medium.com/nullifying-the-null/getting-started-with-terraform-and-aws-751d762f31fc

### Project structure
There are different files in this project:

1. `Makefile` — contains command shortcuts to execute Terraform commands. For example, **make apply** executes Terraform's apply command.
2. `main.tf` — defines the resource blocks and provider configuration
3. `variables.tf` — defines the variables used in main.tf
4. `values.tfvars` — assigns values to the variables in variables.tf

### Prerequisites
1. AWS CLI - https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
    
    * Terraform uses the configured credentials when you set up AWS CLI
    * You'll you have access to your AWS resources via CLI when you successfully executed `aws sts get-caller-identity`

2. Terraform CLI - https://developer.hashicorp.com/terraform/tutorials/aws-get-started/install-cli

### Environment setup
1. Create an S3 bucket that will contain the Terraform state file.
2. Populate Makefile variables with your AWS region and S3 bucket name respectively: **region** and **s3bucket**
3. Execute `make init`. You should see this on your terminal: *Terraform has been successfully initialized!*
