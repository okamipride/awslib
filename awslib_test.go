package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/sqs"
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

func TestSqsSendMessage(t *testing.T) {
	sqsurl := "https://sqs.us-west-2.amazonaws.com/432825407307/practice"
	params := &sqs.SendMessageInput{
		MessageBody:  aws.String("boto3"), // Required
		QueueUrl:     aws.String(sqsurl),  // Required
		DelaySeconds: aws.Int64(1),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Author": { // Required
				DataType:    aws.String("String"), // Required
				StringValue: aws.String("mary"),
			},
		},
	}
	mysqs := SqsNewSess()
	mysqs.SendMessage(params)
}

func TestSqsFetchMessage(t *testing.T) {
	sqsurl := "https://sqs.us-west-2.amazonaws.com/432825407307/practice"
	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(sqsurl), // Required
		AttributeNames: []*string{
			aws.String("boto3"), // Required.
		},
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String("Author"), // Required
		},
		VisibilityTimeout: aws.Int64(1),
		WaitTimeSeconds:   aws.Int64(1),
	}
	mysqs := SqsNewSess()
	input := mysqs.RecieveMessage(params)
	numberRecieve := len(input.Messages)

	for i := 0; i < numberRecieve; i++ {
		handle := input.Messages[i].ReceiptHandle
		delparams := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(sqsurl),  // Required
			ReceiptHandle: aws.String(*handle), // Required
		}
		mysqs.DeleteMessage(delparams)
	}
	//Delete Message

}
