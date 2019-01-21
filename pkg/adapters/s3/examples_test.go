package s3_test

import (
	"log"

	"github.com/cvgw/sql-db-restore/pkg/adapters/s3"
)

func ExampleGetObj() {
	sqlFilePath, err := s3.GetObj("bucket-foo", "dump.sql", "dev")
	if err != nil {
		log.Fatalf("could not pull SQL file from s3: %s", err)
	}

	log.Println(sqlFilePath)
}
