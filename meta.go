package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

type MetaData struct {
	mdata  *ec2metadata.EC2Metadata
	config *aws.Config
}

func MetaDataInit() MetaData {
	log.Println("MetaDataInit")
	cred := getCredential(access_Key_ID, secret_Access_Key)
	config := aws.NewConfig().WithCredentials(cred).WithRegion(US_WEST_OR)
	return MetaData{config: config}
}

func (m *MetaData) MetaNewSess() {
	log.Println("MetaGetSess")
	m.mdata = ec2metadata.New(session.New())
}

func (m *MetaData) GetPublicIP() (string, error) {
	return m.mdata.GetMetadata("latest/public-ipv4")
}
