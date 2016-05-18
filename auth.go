package awslib

import (
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
	access_Key_ID     = "AKIAIKSEZHUELPUWO7DA"
	secret_Access_Key = "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"
)

func getCredential(accessKey, secretKey string) *credentials.Credentials {
	return credentials.NewStaticCredentials(accessKey, secretKey, ``)
}
