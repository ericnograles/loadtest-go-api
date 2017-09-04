package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func parseDatabaseURL(databaseSourceURL string) string {
	// This function assumes a standard Postgres URI format with credentials at the beginning
	// e.g. for localhost, postgres://postgres@127.0.0.1:5432/some_db
	databaseURL := strings.Replace(databaseSourceURL, "postgres://", "", -1)
	parsedURL := strings.Split(databaseURL, "@")
	credentials := strings.Split(parsedURL[0], ":")
	dbParams := strings.Split(parsedURL[1], "/")
	hostAndPort := strings.Split(dbParams[0], ":")
	sslMode := "require"

	if strings.Contains(databaseURL, "127.0.0.1") || strings.Contains(databaseURL, "localhost") {
		sslMode = "disable"
	}

	postgresConnection := ""
	if len(credentials) == 2 {
		host, port, user, password, dbname := hostAndPort[0], hostAndPort[1], credentials[0], credentials[1], dbParams[1]
		postgresConnection = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v", host, port, user, password, dbname, sslMode)
	} else {
		host, port, user, dbname := hostAndPort[0], hostAndPort[1], credentials[0], dbParams[1]
		postgresConnection = fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=%v", host, port, user, dbname, sslMode)
	}

	return postgresConnection
}

func getDatabase(databaseSourceURL string) *gorm.DB {
	db, err := gorm.Open("postgres", parseDatabaseURL(databaseSourceURL))
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func migrate(db *gorm.DB) {
	db.DropTableIfExists(&User{})
	db.AutoMigrate(&User{})
	db.Create(&User{FirstName: "Test", LastName: "User", EmailAddress: "test@test.com"})
	db.Create(&User{FirstName: "Test", LastName: "User2", EmailAddress: "test2@test.com"})
}
