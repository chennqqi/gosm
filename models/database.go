package models

import (
	"database/sql"
	"encoding/json"
	"os"

	"github.com/jmoiron/sqlx"
	// Add sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

const (
	createTableSQL = `CREATE TABLE services (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    protocol TEXT NOT NULL,
    host TEXT NOT NULL,
    port TEXT
);`
)

// Database The sqlite3 database
var Database *sqlx.DB

// Connect Connects to the sqlite3 database, and creates the database if it does not already exist
func Connect(dbpath string) {
	var needsSetup = false
	if _, err := os.Stat(dbpath); os.IsNotExist(err) {
		needsSetup = true
		f, err := os.Create(dbpath)
		if err != nil {
			panic(err)
		}
		f.Close()
	}
	database, err := sqlx.Open("sqlite3", dbpath)
	if err != nil {
		panic(err)
	}
	Database = database
	if needsSetup {
		_, err = Database.Exec(createTableSQL)
		if err != nil {
			panic(err)
		}
		//		database, err = sqlx.Open("sqlite3", dbpath)
		//		if err != nil {
		//			panic(err)
		//		}
		//		Database = database
	}
}

type jsonNullInt64 struct {
	sql.NullInt64
}

func (v jsonNullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}
	return json.Marshal(nil)
}

func (v *jsonNullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}
	return nil
}
