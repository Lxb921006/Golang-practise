package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// aws s3 putObject example: aws s3api put-object --bucket db-backup-huawen --key truco/rpmlist.txt(上传到bucket的位置) --body ./rpmlist.txt(本地文件)
func main() {

	sess, err := session.NewSession(&aws.Config{
		MaxRetries:  aws.Int(3),
		Credentials: credentials.NewSharedCredentials("C:/Users/Administrator/Desktop/aws.ini", "aws-huawen-root"),
		Region:      aws.String("sa-east-1"),
	})

	if err != nil {
		return
	}

	s3api := s3.New(sess)

	input := &s3.PutObjectInput{
		Body:   aws.ReadSeekCloser(strings.NewReader("C:/Users/Administrator/Desktop/TaskcplconfigController.class.php")),
		Bucket: aws.String("db-backup-huawen"),
		Key:    aws.String("truco/TaskcplconfigController.class.php"),
	}

	res, err := s3api.PutObject(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

}
