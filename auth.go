package awslib

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	access_Key_ID     = ""
	secret_Access_Key = ""
)

func getCredential(accessKey, secretKey string) *credentials.Credentials {
	return credentials.NewStaticCredentials(accessKey, secretKey, ``)
}
