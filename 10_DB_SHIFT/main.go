package main

// s
import (
	"database/sql"
	"fmt"
	//"errors"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	fmt.Println(err)
}

// Function that connects to db and does stuff
func ConnectDB() {

	// registering the database driver interface, now the function from mysql driver will work
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testdb?charset=utf8")
	if err != nil {
		fmt.Println("An error occured while trying to connect to database", err)
		return
	}
	defer db.Close()
	fmt.Println("connection was successfull")

	// Statement to be executed is being prepared, ? is the placeholder
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

  // Actually executing the statement
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

  // Read everything from the table
	rows, err := db.Query(`SELECT * FROM userinfo`)
	checkErr(err)
	defer rows.Close();

	// Get the information about the columntypes, that will be used for later reason
	coltypes, err := rows.ColumnTypes(); // Reeturns slice of pointer to Column/ColumnTypes
	for _, value := range coltypes {
		fmt.Println(value.Name(),value.ScanType());
	}
	checkErr(err);

	for rows.Next() {  // next() prepares the next result that will be used by the scan
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)  // Actually read form the dataset
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func main() {
	ConnectDB()
}
