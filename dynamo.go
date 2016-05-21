package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
)

type Dynamo struct {
	DB     *dynamodb.DynamoDB
	Config *aws.Config
}

func DynamoNewSess() *Dynamo {
	log.Println("DynamoDBNewSess")
	//cred := getCredential(access_Key_ID, secret_Access_Key)
	//config := aws.NewConfig().WithCredentials(cred).WithRegion(US_WEST_OR)
	config := aws.NewConfig().WithRegion(US_WEST_OR)
	srv := dynamodb.New(session.New(config), config)
	return &Dynamo{DB: srv, Config: config}
}

func (d *Dynamo) ScanDB(par *dynamodb.ScanInput) {
	resp, err := d.DB.Scan(par)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	log.Println(resp)

}
