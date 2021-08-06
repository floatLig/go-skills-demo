package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

}

var DB *sql.DB

func mysqlInit() error {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", "root", "root", "127.0.0.1", "3306", "entryTask")
	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
