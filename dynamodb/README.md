Build a code using command:
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o main

Build Docker image using command:
docker build -t dynamodb-crud .
run on Docker
docker run -p 8080:8080 -e -e AWS_DEFAULT_REGION=<AWS_DEFAULT_REGION> -e AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID> -e AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY> dynamodb-crud
Note:- make sure AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY should have proper access of dynamodb

run on local server:
export AWS_SECRET_ACCESS_KEY=<AWS_SECRET_ACCESS_KEY> 
export AWS_DEFAULT_REGION=<AWS_DEFAULT_REGION> 
export AWS_ACCESS_KEY_ID=<AWS_ACCESS_KEY_ID>
#{PWD}/main

Get jwt token
curl -X POST http://localhost:8080/signin -H 'Content-Type: application/json' -d '{"username": "admin","password": "password"}'
Set Token into cookie
Create Item
curl -X POST http://localhost:8080/api/v1/table/{table} -H 'Content-Type: application/json' -d '{ "id": "109", "Name": "hello-test","Website": "test"}' -b "token=<token>"
Get Item 
curl -X GET  http://localhost:8080/api/v1/table/{table}/id/{id} -b "token=<token>"
Update Item
curl -X PUT http://localhost:8080/api/v1/table/{table}/id/{id} -H 'Content-Type: application/json' -d '{"Name": "hello-test","Website": "link"}' -b "token=<token>"
Delete Item
curl -X DELETE http://localhost:8080/api/v1/table/{table}/id/{id} -H 'Content-Type: application/json' -b "token=<token>"

Note:- replace {table} with table name
       replace {id} with id no.