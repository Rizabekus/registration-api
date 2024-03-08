package storage

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateTables(db *sql.DB) error {
	migrationPath := os.Getenv("MIGRATION_PATH")
	file, err := os.ReadFile(migrationPath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(file))
	if err != nil {
		fmt.Println("WHATWHY")
		return err
	}
	// requests := strings.Split(string(file), ";")
	// for _, request := range requests {
	// 	fmt.Println("REQUEST ", request)
	// 	_, err := db.Exec(request)
	// 	if err != nil {
	// 		fmt.Println("WHATWHY")
	// 		return err
	// 	}
	// }
	return nil
}
