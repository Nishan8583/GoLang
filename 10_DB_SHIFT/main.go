package main

import (
	"crypto/tls"
	"database/sql"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Req is used to creat post data to elasticsearch
var Req *http.Request

// Client is used to actualy perform the request to elasticsearch
var Client *http.Client

// error type
var err error

func init() {
	Req, err = http.NewRequest("POST", "https://127.0.0.1:9200/lolta/_doc", nil)
	if err != nil {
		fmt.Println("Error creating request object", err)
		return
	}

	// adding Headers
	Req.Header.Add("Content-Type", "application/json")

	auth := b64.StdEncoding.EncodeToString([]byte("admin:admin"))
	auth = "Basic " + auth

	Req.Header.Add("Authorization", auth)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	Client = &http.Client{Transport: tr}

}

/* CREATE DATABASE testdb;
mysql> CREATE DATABASE testdb;
Query OK, 1 row affected (0.01 sec)

mysql> use testdb;
Database changed
mysql> CREATE TABLE userinfo (username varchar(20) not null, departname varchar(20) not null,age int);
Query OK, 0 rows affected (0.09 sec)
CREATE DATABASE lolta
CREATE TABLE netflow (source_ip varchar(20) not null, destination_ip varchar(20) not null,souorce_port int,destinatination_port int);
INSERT INTO netflow (source_ip,destination_ip,souorce_port,destinatination_port) VALUES ("192.168.0.1","192.168.0.3",56,60);
INSERT INTO netflow (source_ip,destination_ip,souorce_port,destinatination_port) VALUES ("192.168.0.1","192.168.0.3",56,60);
INSERT INTO netflow (source_ip,destination_ip,souorce_port,destinatination_port) VALUES ("192.168.0.2","192.168.0.3",56,60);
INSERT INTO netflow (source_ip,destination_ip,souorce_port,destinatination_port) VALUES ("192.168.0.6","192.168.0.3",56,60);
INSERT INTO netflow (source_ip,destination_ip,souorce_port,destinatination_port) VALUES ("192.168.0.56","192.168.0.3",56,60);

mysql> INSERT INTO userinfo (username,departname,age) VALUES ("nishan","R&D",50);
Query OK, 1 row affected (0.01 sec)

mysql> INSERT INTO userinfo (username,departname,age) VALUES ("nishu","R&D",50);
Query OK, 1 row affected (0.06 sec)

mysql> INSERT INTO userinfo (username,departname,age) VALUES ("lolta","R&D",50);
Query OK, 1 row affected (0.00 sec)
stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,age=?")
checkErr(err)

// Actually executing the statement
res, err := stmt.Exec("newguy", "SOC", 80)
checkErr(err)

id, err := res.LastInsertId()
checkErr(err)
fmt.Println(id)
*/

// Function that connects to db and does stuff
func ConnectDB() {

	// registering the database driver interface, now the function from mysql driver will work
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lolta?charset=utf8")
	if err != nil {
		fmt.Println("An error occured while trying to connect to database", err)
		return
	}
	defer db.Close()
	fmt.Println("connection was successfull")

	// Read everything from the table
	fmt.Println("Now reading from the table netflow")
	rows, err := db.Query(`SELECT * FROM netflow`)
	if err != nil {
		fmt.Println("Could not read data from table", err)
		return
	}
	defer rows.Close()

	cols, err := rows.Columns()            // Get slice of columns name
	vals := make([]interface{}, len(cols)) // Making a slice of any type of length number of columns
	for i := 0; i < len(cols); i++ {
		vals[i] = new(sql.RawBytes) // Each element is of pointer to sql.RawBytes, everything can be interpreted as this
	}

	for rows.Next() { // next() prepares the next result that will be used by the scan
		err = rows.Scan(vals...) // Get the row data and unmarshal it in vals
		final := `{`

		for i, v := range vals {
			switch s := v.(type) {
			default:
				_ = s
				final = final + fmt.Sprintf(`"%s":"%s",`, cols[i], string(*v.(*sql.RawBytes)))
			}

		}
		final = final + fmt.Sprintf(`"timestamp":"%s"`, time.Now().Format(time.RFC3339)) + `}`
		fmt.Println(final)
		Req.Body = ioutil.NopCloser(strings.NewReader(final))

		res, err := Client.Do(Req)
		if err != nil {
			fmt.Println("Error Posting value to elasticsearch:", err)
			return
		}
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		fmt.Println(string(content))
	}

}

func main() {
	ConnectDB()
}
