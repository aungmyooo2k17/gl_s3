package initializers

import (
	"gl_s3/internal/pkg/cloud/aws"
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
)

func InitAWS() *session.Session {
	// Create a session instance.
	ses, err := aws.New(aws.Config{
		Address: "http://localhost:4566",
		Region:  GetEnvWithKey("AWS_REGION"),
		Profile: "localstack",
		ID:      GetEnvWithKey("AWS_ACCESS_KEY_ID"),
		Secret:  GetEnvWithKey("AWS_SECRET_ACCESS_KEY"),
	})
	if err != nil {
		log.Fatalln(err)
	}

	return ses
}
