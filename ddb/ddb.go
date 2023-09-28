package ddb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

struct dynamodb type {
	
}

func setupDDB(port int) *dynamodb.DynamoDB {
	sessConf := &aws.Config{
		Credentials: credentials.NewStaticCredentials("STATIC", "DUMMY", "VALUE"),
		Region:      aws.String("local"),
	}
	request.WithRetryer(sessConf, nil)
	sess := session.Must(session.NewSession(sessConf))
	ddbConf := aws.Config{
		Endpoint: aws.String(fmt.Sprintf("http://localhost:%v", port)),
	}
	return dynamodb.New(sess, &ddbConf)
}

func 