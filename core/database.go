package core

import (
	"database/sql"
	"fmt"
)

// DB is type of database
type DB struct {
	DBName   string
	Username string
	Password string
}

// Databases is a set of db config
var Databases = map[string]DB{}

// MasterDB is master db
var MasterDB *sql.DB

// var ClientDatabase string

// Connect connect to database
func (db DB) Connect() *sql.DB {
	dbInfo := fmt.Sprintf("user=%s dbname=%s password=%s", db.Username, db.DBName, db.Password)
	d, err := sql.Open("postgres", dbInfo)

	CheckErr(err, "Can not connect to  master database ")

	return d
}
