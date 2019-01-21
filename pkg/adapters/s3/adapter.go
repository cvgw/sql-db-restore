package s3

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//GetObj retrieves the specified S3 object using the supplied IAM profile
func GetObj(bucket, key, iamProfile string) (string, error) {
	log.Println("pulling sql file from s3")

	defaultRegion := "us-west-2"

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(defaultRegion)},
		Profile: "dev",
	}))

	downloader := s3manager.NewDownloader(sess)

	filename := "s3_sql_out.sql"
	// Create a file to write the S3 Object contents to.
	f, err := os.Create(filename)
	if err != nil {
		return "", fmt.Errorf("failed to create file %q, %v", filename, err)
	}
	defer f.Close()

	// Write the contents of S3 Object to the file
	_, err = downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return "", fmt.Errorf("failed to download file, %v", err)
	}

	return filename, nil
}
