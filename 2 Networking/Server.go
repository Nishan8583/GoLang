package main;

import (
	"net";
	"fmt";
	"bufio";
	"log";
	"os";
)

func main() {
	ln, err := net.Listen("tcp",":8080");
	if err != nil {
		log.Fatalln("Could not start the server ERROR: ",err);
	}
	defer func(){
		ln.Close();
		}();

	fmt.Println("I am ready to recieve one connection from Client");

	conn, err := ln.Accept();
	if err != nil {
		fmt.Println("An error occured while listening for connection: ",err);
		return;
	}
	server_reader := bufio.NewReader(os.Stdin);
	fmt.Printf("ME<Server>: ");
	server_msg,err := server_reader.ReadString('\n');
	if err != nil {
		fmt.Println("Error reading my own input ERROR: ",err);
	}
	conn.Write([]byte(server_msg));

	for {
		
		message, err := bufio.NewReader(conn).ReadString('\n');
		if err != nil {
			fmt.Println("Error occured while recieving message from the client ERROR: ",err);
			return;
		}
		if message == "quit\n" {
			fmt.Println("Client does not want to talk anymore. Whats the purpose in life anymore. I EXIT");
			os.Exit(1);
		}
		fmt.Println("<Client> : ",message);
		fmt.Printf("ME<Server>: ");
		server_msg,_ = server_reader.ReadString('\n');
		conn.Write([]byte(server_msg));
	}
}