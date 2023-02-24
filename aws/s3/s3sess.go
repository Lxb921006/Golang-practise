package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Sess struct {
	AccessFile    string
	AccessSection string
	Region        string
}

func (p *S3Sess) s3sess() (s3api *s3.S3, err error) {
	sess, err := session.NewSession(&aws.Config{
		MaxRetries:  aws.Int(3),
		Credentials: credentials.NewSharedCredentials(p.AccessFile, p.AccessSection),
		Region:      &p.Region,
	})

	if err != nil {
		return
	}

	s3api = s3.New(sess)

	return
}

func NewS3Sess(parms ...string) *S3Sess {
	return &S3Sess{
		AccessFile:    parms[0],
		AccessSection: parms[1],
		Region:        parms[2],
	}
}
