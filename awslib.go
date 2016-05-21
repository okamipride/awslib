package awslib

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	US_EAST_VA       = "us-east-1"
	US_WEST_OR       = "us-west-2"
	US_WEST_CA       = "us-west-1"
	EU_IR            = "eu-west-1"
	EU_FK            = "eu-central-1"
	ASIA_SG          = "ap-southeast-1"
	ASIA_SD          = "ap-southeast-2"
	ASIA_TK          = "ap-northeast-1"
	SOUTH_AMERICA_SP = "sa-east-1"
)

const (
	access_Key_ID     = ""
	secret_Access_Key = ""
)

func getCredential(accessKey, secretKey string) *credentials.Credentials {
	return credentials.NewStaticCredentials(accessKey, secretKey, ``)
}
