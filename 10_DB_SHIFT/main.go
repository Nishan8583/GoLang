package main

import (
	"database/sql"
	"fmt"
	//"errors"
	_ "github.com/go-sql-driver/mysql"
)

func Conenctdb(username, password string) {
	var db *sql.DB
	var err error
	// If no password then just use no password
	if len(password) == 0 {
		fmt.Println("No password was given so using default empty password")
		db, err = sql.Open("mysql", fmt.Sprintf(`%s:@tcp(127.0.0.1:3306)/`, username))
		if err != nil {
			fmt.Println("Could not connect to Database ", err)
			return
		}
	} else { // if password was given
		db, err = sql.Open("mysql", fmt.Sprintf(`%s:%s@tcp(127.0.0.1:3306`, username, password))
		if err != nil {
			fmt.Println("Could not conenct to database ", err)
		}

	}

	// Creating a sample database;
	_, err = db.Exec("CREATE DATABASE testdb")
	if err != nil {
		fmt.Println("Could not create dataabase", err)
		return
	}
	fmt.Println("Creation of database succefull")

	// Selecting DATABASE
	_, err = db.Exec("USE testdb")
	if err != nil {
		fmt.Println("Could connect to dataabase", err)
		return
	}
	fmt.Println("Connecting to database succefull")

	// Putting data
	stmt, err := db.Prepare(`CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));`)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}

}

func main() {
	Conenctdb("root", "")
}
