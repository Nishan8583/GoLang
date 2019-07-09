/*
The Package contians function that can be used to push data from MySQL to elasticsearch
	pusher, err := New("127.0.0.1:3306", "root", "root", "lolta", "netflow", "https://127.0.0.1:9200", "namer", "admin", "admin", false, true)
	if (err) != nil {
		fmt.Println("Error getting pusher", err)
		return
	}
	pusher.PushToES()
*/
package MySQLtoES

import (
	"crypto/tls"
	"database/sql"
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"os"
	"net/http"
	"strings"
	"time"
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLToEs structure will hold the necessary infromation about the MySQL Database and elasticsearch
type MySQLToEs struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBTable    string
	DBIp       string
	Req        *http.Request
	Client     *http.Client
}

var err error
var dbip,dbuser,dbpassword,dbname,dbtablename,esip,esuser,espassword,esindex *string;
var usessl,useauth *bool;

// init() function is setting up the arguments that will be used by the compiled program
func init() {
	dbip = flag.String("dbip","127.0.0.1:3306",`Insert the Ipaddress of the MySQL Database\n Usage: -dbip 127.0.0.1:3306`);
	dbuser = flag.String("dbuser","root",`Insert Ex: -dbuser root`)
	dbpassword = flag.String("dbpassword","root",`Insert dbpassword : -dbpassword root`)
	dbname = flag.String("dbname", "test",`InsertDBName please Ex: -dbname test`)
	dbtablename = flag.String("dbtablename","test",`Inser DB TAble name please Ex: -dbtablename test`)
	esip = flag.String("esip","http://127.0.0.1:9200",`Insert The elasticsearch Ip address Ex: -esip http://127.0.0.1`)
	esuser = flag.String("esuser","",`Insert ES Username Ex: -esuer admin`)
	espassword = flag.String("espassword","",`Insert ES Password Ex: -espassword admin`)
	esindex = flag.String("esindex","test",`Insert elasticsearch Index to push data to Ex: -esindex test`)
	usessl = flag.Bool("usessl",false,`Specifiy whether to use ssl while communicating with ES Ex: -usessl false`)
	useauth = flag.Bool("useauth",false,`Specifiy whether to use authentication while communicating with ES Ex: -usessl false`)

	flag.Parse();
}

// New() A factory function that does the query
func New(DBIp, DBUser, DBPassword, DBName, DBTable, ESIp, ESIndex, ESUser, ESPassword string, ESSSLVerification, ESUseAuth bool) (MySQLToEs, error) {

	// Type to perform actions
	mysqlT0es := MySQLToEs{
		DBUser:     DBUser,
		DBPassword: DBPassword,
		DBName:     DBName,
		DBTable:    DBTable,
		DBIp:       DBIp,
	}

	// Generating Requests
	Req, err := http.NewRequest("POST", fmt.Sprintf(`%s/%s/_doc`, ESIp, ESIndex), nil)
	if err != nil {
		log.Println("Could not create request", err)
		return mysqlT0es, err
	}
	Req.Header.Add("Content-Type", "application/json")
	// If using authentication flag is provided
	if ESUseAuth {
		auth := b64.StdEncoding.EncodeToString([]byte("admin:admin"))
		auth = "Basic " + auth
		Req.Header.Add("Authorization", auth)
	}
	mysqlT0es.Req = Req

	// if SSLVerification is enable to true
	if ESSSLVerification {
		mysqlT0es.Client = &http.Client{}
	} else {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
		mysqlT0es.Client = &http.Client{Transport: tr}
	}
	return mysqlT0es, nil
}

// Function that connects to db and does stuff
func (mte *MySQLToEs) PushToES() error {

	// registering the database driver interface, now the function from mysql driver will work
	db, err := sql.Open("mysql", fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8`, mte.DBUser, mte.DBPassword, mte.DBIp, mte.DBName))
	if err != nil {
		fmt.Println("An error occured while trying to connect to database", err)
		return err
	}
	defer db.Close()
	fmt.Println("connection was successfull")

	// Read everything from the table
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM %s`, mte.DBTable))
	if err != nil {
		fmt.Println("Could not read data from table", err)
		return err
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
		mte.Req.Body = ioutil.NopCloser(strings.NewReader(final))

		res, err := mte.Client.Do(mte.Req)
		if err != nil {
			fmt.Println("Error Posting value to elasticsearch:", err)
			return err
		}
		wr, err := io.Copy(mte,res.Body)
		if (err != nil) {
			fmt.Println(err,wr);
		}
		/*
		content, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer res.Body.Close()
		fmt.Println(string(content))
	}*/
}
	return err
}

func (mte *MySQLToEs) Write(p []byte) (n int, err error) {
	fmt.Println(string(p));
	return 1,nil
}
// main() function of the program
func main() {
	if len(os.Args) < 2{
		fmt.Println("No arguments provided wil use default values. Please use ./MtoE --help for more information")
	}
	pusher, err := New(*dbip, *dbuser, *dbpassword, *dbname, *dbtablename, *esip, *esindex, *esuser, *esindex, *usessl, *useauth)

	//pusher, err := New("127.0.0.1:3306", "root", "root", "lolta", "netflow", "https://127.0.0.1:9200", "namer", "admin", "admin", false, true)
	if (err) != nil {
		fmt.Println("Error getting pusher", err)
		return
	}
	pusher.PushToES()
}

/* Sample Database Creation
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
