package db

import (
	"log"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
  )  
func GetDyanmodbClient() (*dynamodb.DynamoDB) {
	aws_session := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_DEFAULT_REGION")),
	}))
	return dynamodb.New(aws_session)
}
func Check_varibale(){
	if os.Getenv("AWS_DEFAULT_REGION") == "" || os.Getenv("AWS_ACCESS_KEY_ID") == "" || os.Getenv("AWS_SECRET_ACCESS_KEY") == ""{
		log.Panicln("Check Env Varible AWS_DEFAULT_REGION or AWS_ACCESS_KEY_ID or AWS_SECRET_ACCESS_KEY")
		os.Exit(1)
	}
	
}
// func DbPing() {
//     client, err := GetMongoClient()
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 	  fmt.Println(err)
// 	}
//     fmt.Println("connected")
// }