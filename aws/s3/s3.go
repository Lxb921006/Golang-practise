package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Object struct {
	S3Sess *S3Sess
	Bucket string
}

func (p *Object) PutObject(src, dst string) (err error) {
	s3api, err := p.S3Sess.s3sess()
	if err != nil {
		return
	}

	of, err := os.Open(src)
	if err != nil {
		return
	}

	defer func(of *os.File) {
		err := of.Close()
		if err != nil {
			return
		}
	}(of)

	input := &s3.PutObjectInput{
		//Body:   aws.ReadSeekCloser(of),
		Body:   of,
		Bucket: aws.String(p.Bucket),
		Key:    aws.String(dst),
	}

	_, err = s3api.PutObject(input)
	if err != nil {
		return
	}

	return
}
