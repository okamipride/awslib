package awslib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
)

type MetaData struct {
	Mdata  *ec2metadata.EC2Metadata
	Config *aws.Config
}

func MetaNewSess() *MetaData {
	log.Println("MetaDataInit")
	cred := getCredential(access_Key_ID, secret_Access_Key)
	config := aws.NewConfig().WithCredentials(cred).WithRegion(US_WEST_OR)
	ec2data := ec2metadata.New(session.New(config), config)
	log.Println("MetaGetSess")
	return &MetaData{Mdata: ec2data, Config: config}

}

func (m *MetaData) GetPublicIP() (string, error) {
	return m.Mdata.GetMetadata("public-ipv4")
}
