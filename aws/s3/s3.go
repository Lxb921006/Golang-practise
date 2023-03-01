package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Object struct {
	S3Sess *S3Sess
	Bucket string
}

func (p *S3Object) PutObject(src, dst string) (err error) {
	s3api, err := p.S3Sess.s3sess()
	if err != nil {
		return
	}

	of, err := os.Open(src)
	if err != nil {
		return
	}

	defer of.Close()

	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(of),
		Bucket: aws.String(p.Bucket),
		Key:    aws.String(dst),
	}

	_, err = s3api.PutObject(input)
	if err != nil {
		return
	}

	return
}
