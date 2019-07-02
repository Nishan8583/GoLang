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

/* CREATE DATABASE testdb;
mysql> CREATE DATABASE testdb;
Query OK, 1 row affected (0.01 sec)

mysql> use testdb;
Database changed
mysql> CREATE TABLE userinfo (username varchar(20) not null, departname varchar(20) not null,age int);
Query OK, 0 rows affected (0.09 sec)

mysql> INSERT INTO userinfo (username,departname,age) VALUES ("nishan","R&D",50);
Query OK, 1 row affected (0.01 sec)

mysql> INSERT INTO userinfo (username,departname,age) VALUES ("nishu","R&D",50);
Query OK, 1 row affected (0.06 sec)

mysql> INSERT INTO userinfo (username,departname,age) VALUES ("lolta","R&D",50);
Query OK, 1 row affected (0.00 sec)

*/

// Function that connects to db and does stuff
func ConnectDB() {

        // registering the database driver interface, now the function from mysql driver will work
        db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testdb?charset=utf8")
        if err != nil {
                fmt.Println("An error occured while trying to connect to database", err)
                return
        }
        defer db.Close()
        fmt.Println("connection was successfull")

        // Statement to be executed is being prepared, ? is the placeholder
        stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,age=?")
        checkErr(err)

  // Actually executing the statement
        res, err := stmt.Exec("astaxie", "SOC", 80)
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)
        fmt.Println(id)

  // Read everything from the table
        rows, err := db.Query(`SELECT * FROM userinfo`)
        checkErr(err)
        defer rows.Close();


        for rows.Next() {  // next() prepares the next result that will be used by the scan
                var username string
                var department string
                var age int
                err = rows.Scan(&username, &department, &age)  // Actually read form the dataset
                checkErr(err)
                fmt.Println(username)
                fmt.Println(department)
                fmt.Println(age)
        }
}

func main() {
        ConnectDB()
}
