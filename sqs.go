package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type Sqs struct {
	SqsSrv *sqs.SQS
	Config *aws.Config
}

func SqsNewSess() *Sqs {
	log.Println("SqsNewSess")
	//cred := getCredential(access_Key_ID, secret_Access_Key)
	//config := aws.NewConfig().WithCredentials(cred).WithRegion(US_WEST_OR)
	config := aws.NewConfig().WithRegion(US_WEST_OR)
	srv := sqs.New(session.New(config), config)
	return &Sqs{SqsSrv: srv, Config: config}
}

func (s *Sqs) SendMessage(par *sqs.SendMessageInput) {
	resp, err := s.SqsSrv.SendMessage(par)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return
	}
	// Pretty-print the response data.
	log.Println(resp)
}

func (s *Sqs) RecieveMessage(par *sqs.ReceiveMessageInput) *sqs.ReceiveMessageOutput {
	resp, err := s.SqsSrv.ReceiveMessage(par)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return nil
	}
	// Pretty-print the response data.
	log.Println(resp)
	return resp
}

func (s *Sqs) DeleteMessage(par *sqs.DeleteMessageInput) {
	resp, err := s.SqsSrv.DeleteMessage(par)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return
	}
	// Pretty-print the response data.
	log.Println(resp)
}
