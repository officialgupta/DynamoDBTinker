package ddb

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Ddb struct {
	db *dynamodb.DynamoDB
}

func SetupDDB(port int) Ddb {
	sessConf := &aws.Config{
		Credentials: credentials.NewStaticCredentials("STATIC", "DUMMY", "VALUE"),
		Region:      aws.String("local"),
	}
	request.WithRetryer(sessConf, nil)
	sess := session.Must(session.NewSession(sessConf))
	ddbConf := aws.Config{
		Endpoint: aws.String(fmt.Sprintf("http://localhost:%v", port)),
	}
	return Ddb{
		db: dynamodb.New(sess, &ddbConf),
	}
}

func (d *Ddb) ListTables() (*[]string, error) {
	remaining := true
	tables := []string{}
	for remaining {
		resp, err := d.db.ListTables(&dynamodb.ListTablesInput{})
		if err != nil {
			return nil, fmt.Errorf("could not list tables: %w", err)
		}
		if resp.LastEvaluatedTableName == nil {
			remaining = false
		}
		for _, table := range resp.TableNames {
			tables = append(tables, *table)
		}
	}

	return &tables, nil
}

func (d *Ddb) CreateTable(tableName string, primaryKey string) error {
	_, err := d.db.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(primaryKey),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(primaryKey),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String(tableName),
	})
	if err != nil {
		return fmt.Errorf("could not create table: %w", err)
	}

	return nil
}

func (d *Ddb) DeleteTable(tableName string) error {
	_, err := d.db.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	if err != nil {
		return fmt.Errorf("could not delete table: %w", err)
	}

	return nil
}
