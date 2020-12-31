before run 
    - export AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY> 
    - export AWS_DEFAULT_REGION=<AWS_DEFAULT_REGION> 
    - export AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>
Note :- Aws user should have proper access to s3, dynamodb, cloudwatch

Go into lamba-fuction folder and run command
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o main & zip main.zip main
and copy main.zip and main file into folder terraform/module

To run 
-> terraform login 
-> terraform init
-> terraform plan 
-> terraform apply
