package awslib

import (
	"log"
	"testing"
)

func TestMetaData(t *testing.T) {
	meta := MetaDataInit()
	meta.MetaNewSess()
	ip, err := meta.GetPublicIP()
	log.Println("My Ip is:", ip, " err:", err)
}
