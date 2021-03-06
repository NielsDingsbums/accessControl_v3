package dbInteractions

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Inserting data into specified table
func InsertData(table string, data []interface{}) bool {

	// Establish connection
	db, err := sql.Open("mysql", DataSorceString)
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return false
	}

	// Actual insert
	switch table {
	case "accounts":
		stmt, err := db.Prepare(fmt.Sprintf("INSERT %s SET uid=?,name=?,tag_name=?,permission=?", table))
		checkErr(err)

		_, err = stmt.Exec(data[0], data[1], data[2], data[3])
		checkErr(err)
		return true

	case "statusLog":
		stmt, err := db.Prepare(fmt.Sprintf("INSERT %s SET status=?,creator=?,publishDate=?", table))
		checkErr(err)

		dateString := data[2].(time.Time).String()[0:19]

		fmt.Println(dateString)

		_, err = stmt.Exec(data[0], data[1], dateString)
		checkErr(err)
		return true

	case "accessLog":
		stmt, err := db.Prepare(fmt.Sprintf("INSERT %s SET uid=?,name=?,publishDate=?", table))
		checkErr(err)

		_, err = stmt.Exec(data[0], data[1], data[2])
		checkErr(err)
		return true
	}
	return true
}
