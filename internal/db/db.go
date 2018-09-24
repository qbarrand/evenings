package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/golog"
)

var dbPath string

func GetDb() (db *gorm.DB) {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		golog.Fatal(err)
	}

	db.SingularTable(true)

	return
}

// func RunQueryAndClose(statement string) (rows *sql.Rows) {
// 	db := GetDb()

// 	rows, err := db.Query(statement)
// 	if err != nil {
// 		golog.Fatal(err)
// 	}

// 	return
// }

func SetDbPath(path string) {
	dbPath = path
}
