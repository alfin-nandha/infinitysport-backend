package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var theSession *session.Session

//GetConfig Initiatilize config in singleton way
func GetSession() *session.Session {

	if theSession == nil {
		theSession = initSession()
	}

	return theSession
}

func initSession() *session.Session {
	newSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Config.S3.Region),
		Credentials: credentials.NewStaticCredentials(Config.S3.Key, Config.S3.Secret, ""),
	}))
	return newSession
}
