package main;

import (
	"net";
	"fmt";
	"bufio";
	"log";
	"os";
)

func main() {
	conn, err := net.Dial("tcp", ":8080");
	if err != nil {
		log.Fatalln("A fucking unrecoverable ERROR occured, ERROR: ",err);
	}
	message, err := bufio.NewReader(conn).ReadString('\n');
	if err != nil {
		fmt.Println("Server sends whatttt!!!!!!!!!!! ERROR: ",err);
		return;
	}
	fmt.Println("<Server> : ",message);
	my_reader := bufio.NewReader(os.Stdin);
	for {
		fmt.Println("<Waiting for server to repy>")
		fmt.Printf("ME<Client>: ");
		msg,err := my_reader.ReadString('\n');
		if err != nil {
			fmt.Println("Error reading my own message ");
			return;
		}
		if msg == "quit\n" {
			fmt.Println("Server does not want to talk with me anymore, Whats the purpose in life. I EXIT");
			os.Exit(1);
		}
		conn.Write([]byte(msg));
		message, err := bufio.NewReader(conn).ReadString('\n');
			if err != nil {
			fmt.Println("Server sends whatttt!!!!!!!!!!! ERROR: ",err);
			return;
		}
		if msg == "quit\n" {
			fmt.Println("Server does not want to talk with me anymore, Whats the purpose in life. I EXIT");
			os.Exit(1);
		}
		fmt.Println("<Server> : ",message);
	}

}