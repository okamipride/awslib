package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
)

func EC2GetSession() *ec2.EC2 {
	log.Println("EC2GetSession")
	cred := getCredential(access_Key_ID, secret_Access_Key)
	srv := ec2.New(session.New(aws.NewConfig().WithCredentials(cred).WithRegion(US_WEST_OR)))
	return srv
}
