package main;

import (
	"net";
	"fmt";
	"log";
	"sync";
	"time";
	"bufio";
)


func server() {
	ln, err := net.Listen("tcp",":8080");  // Returns *IPConn struct type
	if err != nil {
		log.Fatalln(err);
	}
	defer func(){
		ln.Close();
		wg.Done();
		}();
	fmt.Println(ln, " A server was successfully started");

	conn,err := ln.Accept(); // Accepting a conneciton

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n');
		if message == "quit\n" {
			fmt.Println("Client said to quit");
			return
		} else {
			fmt.Println("CLient: ",message);
			conn.Write([]byte("Hello client"+"\n"));
		}
	}
}

func client() {
	time.Sleep(3);
	conn, err := net.Dial("tcp",":8080");
	if err != nil {
		log.Fatalln(err);
	}
	defer func(){
		wg.Done();
		}();	
	fmt.Println(conn, " connection successfully");
	conn.Write([]byte("hello nibba"+"\n"));
	message,_ := bufio.NewReader(conn).ReadString('\n');
	fmt.Println("Server Sama: ",message);
	conn.Write([]byte("quit"+"\n"))
}

var wg sync.WaitGroup;
func main() {
	wg.Add(2);

	go server();
	go client();

	wg.Wait();
}