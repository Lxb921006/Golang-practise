package s3

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Object struct {
	S3Sess
	Bucket string
}

func (p *S3Object) PutObject(src, dst string) (err error) {
	s3api, err := p.s3sess()
	if err != nil {
		return
	}

	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("C:/Users/Administrator/Desktop/TaskcplconfigController.class.php")),
		Bucket: aws.String("db-backup-huawen"),
		Key:    aws.String("truco/TaskcplconfigController.class.php"),
	}

	_, err = s3api.PutObject(input)
	if err != nil {
		return
	}

	return
}
