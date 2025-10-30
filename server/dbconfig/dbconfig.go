package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var Database *sql.DB

func ConnectDB() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=tenacious_data_db sslmode=disable"

	Database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	err = Database.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	fmt.Println("✅ PostgreSQL Connected Successfully!")
}


// PSQL user name : postgres and password : 12345
// user name : Poosu and password : 12345