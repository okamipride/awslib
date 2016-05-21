package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println(" ---- awslib testing start----")
	retcode := m.Run()
	os.Exit(retcode)
}

func TestMetaData(t *testing.T) {
	log.Println("TestMetaData")
	meta := MetaNewSess()
	ip, err := meta.GetPublicIP()
	log.Println("My Ip is:", ip, " err:", err)
}

func TestDynamoScan(t *testing.T) {
	tb := "composit"
	/*
		params1 := &dynamodb.ScanInput{
			TableName: aws.String(tb), // Required
			AttributesToGet: []*string{
				aws.String("ip"), // Required
				aws.String("login_time"),
				aws.String("user"),
				aws.String("firstname"),
				aws.String("lastname"),
			},
			ConsistentRead: aws.Bool(true),
		}
	*/
	params2 := &dynamodb.ScanInput{
		TableName:            aws.String(tb),
		ProjectionExpression: aws.String("ip,lastname"),
	}
	dynamo := DynamoNewSess()
	dynamo.ScanDB(params2)
}
