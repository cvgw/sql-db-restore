package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cvgw/sql-db-restore/pkg/adapters/s3"
	"github.com/cvgw/sql-db-restore/pkg/restore"
)

const (
	dbHostVar          = "DB_HOST"
	dbPortVar          = "DB_PORT"
	sqlFilePathVar     = "SQL_FILE_PATH"
	dbUserNameVar      = "DB_USER_NAME"
	dbUserPassVar      = "DB_USER_PASS"
	sqlFileS3BucketVar = "SQL_FILE_S3_BUCKET"
	sqlFileS3KeyVar    = "SQL_FILE_S3_KEY"
	iamProfileVar      = "AWS_IAM_PROFILE"
)

func main() {
	var (
		sqlFileS3Key    string
		sqlFileS3Bucket string
		iamProfile      string
	)

	validate := make([][]string, 0)

	sqlFilePath := os.Getenv(sqlFilePathVar)

	dbHost := os.Getenv(dbHostVar)
	validate = append(validate, []string{
		dbHost, dbHostVar,
	})

	dbPort := os.Getenv(dbPortVar)
	validate = append(validate, []string{
		dbPort, dbPortVar,
	})

	dbUserName := os.Getenv(dbUserNameVar)
	validate = append(validate, []string{
		dbUserName, dbUserNameVar,
	})

	dbUserPass := os.Getenv(dbUserPassVar)
	validate = append(validate, []string{
		dbUserPass, dbUserPassVar,
	})

	if sqlFilePath == "" {
		sqlFileS3Key = os.Getenv(sqlFileS3KeyVar)
		validate = append(validate, []string{
			sqlFileS3Key, sqlFileS3KeyVar,
		})

		sqlFileS3Bucket = os.Getenv(sqlFileS3BucketVar)
		validate = append(validate, []string{
			sqlFileS3Bucket, sqlFileS3BucketVar,
		})

		iamProfile := os.Getenv(iamProfileVar)
		validate = append(validate, []string{
			iamProfile, iamProfileVar,
		})
	}

	for _, validatePair := range validate {
		if validatePair[0] == "" {
			log.Fatalf("%s cannot be blank", validatePair[1])
		}
	}

	if sqlFilePath == "" {
		var err error
		sqlFilePath, err = s3.GetObj(sqlFileS3Bucket, sqlFileS3Key, iamProfile)
		if err != nil {
			log.Fatalf("could not pull SQL file from s3: %s", err)
		}
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%s/", dbUserName, dbUserPass, dbHost, dbPort))
	if err != nil {
		log.Fatal("could not open connection to sql db")
	}

	restore.RestoreSQLFile(db, sqlFilePath)
}
