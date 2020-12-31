package curd

import (
	"fmt"
	"strconv"
	"dynamodb-crud/db"
	"dynamodb-crud/middleware"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

)

var dynamo *dynamodb.DynamoDB
func init() {
	dynamo = db.GetDyanmodbClient()
}

type Person struct {
	Id int
	Name string
	Website string
}


func Routescall(router *mux.Router) {
	// router.HandleFunc("/tables", ListTables).Methods("GET")
	router.HandleFunc("/table/{table}/id/{id}", GetItem).Methods("GET")
	router.HandleFunc("/table/{table}/id/{id}", middleware.CheckBodyJson(UpdateItem)).Methods("PUT")
	router.HandleFunc("/table/{table}", middleware.CheckBodyJson(CreateItem)).Methods("POST")
	router.HandleFunc("/table/{table}/id/{id}", DeleteItem).Methods("DELETE")
}

func ListTables(w http.ResponseWriter, r *http.Request) {
	result, err := dynamo.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
	  if aerr, ok := err.(awserr.Error); ok {
		fmt.Println(aerr.Error())
	  }
	}
	json.NewEncoder(w).Encode(result.TableNames)
}
func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
		  "Id": {
			N: aws.String(vars["id"]),
		  },
		},
		TableName: aws.String(vars["table"]),
	  })
	
	  if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
		  fmt.Println(aerr.Error())
		}
	  }
	  var person Person
	  err = dynamodbattribute.UnmarshalMap(result.Item, &person)
	  if err != nil {
		panic(err)
	  }
	  if person.Id == 0 {
		json.NewEncoder(w).Encode("Id not Found")
	  } else {
		json.NewEncoder(w).Encode(person)
	  }
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var person Person
	req := r.Context().Value("req")
	reqBody, _ := json.Marshal(req)
	json.Unmarshal(reqBody, &person)
	vars := mux.Vars(r)
	if CheckItem(person.Id, vars["table"]) {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue {
		  "Id": {
			N: aws.String(strconv.Itoa(person.Id)),
		  },
		  "Name": {
			S: aws.String(person.Name),
		  },
		  "Website": {
			S: aws.String(person.Website),
		  },
		},
		TableName: aws.String(vars["table"]),
	  })
	
	  if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
		  fmt.Println(aerr.Error())
		}
		json.NewEncoder(w).Encode("Somthing went Wrong")
		return
	  }
		json.NewEncoder(w).Encode("Successfully Added")
	  } else {
		  json.NewEncoder(w).Encode("Id already Exist")
	  }
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	var person Person
	req := r.Context().Value("req")
	reqBody, _ := json.Marshal(req)
	json.Unmarshal(reqBody, &person)
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Somthing went wrong")
	}
	if CheckItem(id, vars["table"]) {
		json.NewEncoder(w).Encode("Id not exist")
	} else {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string {
		  "#N": aws.String("Name"),
		  "#W": aws.String("Website"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
		  ":Name": {
			S: aws.String(person.Name),
		  },
		  ":Website": {
			S: aws.String(person.Website),
		  },
		},
		Key: map[string]*dynamodb.AttributeValue{
		  "Id": {
			N: aws.String(vars["id"]),
		  },
		},
		TableName: aws.String(vars["table"]),
		UpdateExpression: aws.String("SET #N = :Name, #W = :Website"),
	  })
	
	  if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
		  fmt.Println(aerr.Error())
		}
		json.NewEncoder(w).Encode("Somthing went Wrong")
		return
	  }
		json.NewEncoder(w).Encode("Successfully Updated")
	}
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id,err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode("Somthing went wrong")
	}
	if CheckItem(id, vars["table"]) {
		json.NewEncoder(w).Encode("Id not exist")
	} else {	
	_, err := dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
		  "Id": {
			N: aws.String(vars["id"]),
		  },
		},
		TableName: aws.String(vars["table"]),
	  })
	
	  if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
		  fmt.Println(aerr.Error())
		}
		json.NewEncoder(w).Encode("Something went wrong")
		return
	  }
		json.NewEncoder(w).Encode("Successfully Deleted")
    }
}

func CheckItem(id int ,table string) bool {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
		  "Id": {
			N: aws.String(strconv.Itoa(id)),
		  },
		},
		TableName: aws.String(table),
	  })
	
	  if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
		  fmt.Println(aerr.Error())
		}
	  }
	  var person Person
	  err = dynamodbattribute.UnmarshalMap(result.Item, &person)
	  if err != nil {
		panic(err)
	  }
	  if person.Id == 0 {
		return true
	  } else {
		return false
	  }
}