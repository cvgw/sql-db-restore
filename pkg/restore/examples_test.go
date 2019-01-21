package restore_test

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cvgw/sql-db-restore/pkg/restore"
)

func ExampleRestoreSQLFile() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/", "foo", "bar"))
	if err != nil {
		log.Fatal("could not open connection to sql db")
	}

	restore.RestoreSQLFile(db, "dump.sql")
}
