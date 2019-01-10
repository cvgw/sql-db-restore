package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/cvgw/sql-db-restore/pkg/restore"
)

const (
	sqlFilePathVar     = "SQL_FILE_PATH"
	dbUserNameVar      = "DB_USER_NAME"
	dbUserPassVar      = "DB_USER_PASS"
	sqlFileS3BucketVar = "SQL_FILE_S3_BUCKET"
	sqlFileS3KeyVar    = "SQL_FILE_S3_KEY"
	iamProfileVar      = "AWS_IAM_PROFILE"
)

func main() {
	validate := make([][]string, 0)

	sqlFilePath := os.Getenv(sqlFilePathVar)
	validate = append(validate, []string{
		sqlFilePath, sqlFilePathVar,
	})

	dbUserName := os.Getenv(dbUserNameVar)
	validate = append(validate, []string{
		dbUserName, dbUserNameVar,
	})

	dbUserPass := os.Getenv(dbUserPassVar)
	validate = append(validate, []string{
		dbUserPass, dbUserPassVar,
	})

	sqlFileS3Bucket := os.Getenv(sqlFileS3BucketVar)
	validate = append(validate, []string{
		sqlFileS3Bucket, sqlFileS3BucketVar,
	})

	sqlFileS3Key := os.Getenv(sqlFileS3KeyVar)
	validate = append(validate, []string{
		sqlFileS3Key, sqlFileS3KeyVar,
	})

	iamProfile := os.Getenv(iamProfileVar)
	validate = append(validate, []string{
		iamProfile, iamProfileVar,
	})

	for _, validatePair := range validate {
		if validatePair[0] == "" {
			log.Fatalf("%s cannot be blank", validatePair[1])
		}
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/", dbUserName, dbUserPass))
	if err != nil {
		log.Fatal("could not open connection to sql db")
	}

	restore.RestoreSQLFile(db, sqlFilePath)
}
