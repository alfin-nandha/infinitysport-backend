package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func S3Config()(*aws.Config){
	return &aws.Config{
		Region : aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("AKIA3JBST3XEWGFLCPYY", "CNWNe7ZuXs9PsJwJxmAnxblCt7gAnO6qnppsVtrJ", ""),
	}
	
}