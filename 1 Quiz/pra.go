package main;

import (
	"fmt";
	"bufio";
	"encoding/csv";
	"io";
	"os";
)

func main() {
	file, _ := os.Open("some.txt"); // Returns the pointe to a File structure

	csv_reader := csv.NewReader(bufio.NewReader(file)); // Interface is used in the back end, since csv.NewReader takes any 
														// struct object with .read() methods
	fmt.Println("************Welcome to Nishans first go quiz************************");

	for {
		line, err := csv_reader.Read(); // Read each line
		if err == io.EOF {
			fmt.Println("END of file reached");
			break;
		}
		fmt.Print("What is the answer ",line[0]," :");
		reader := bufio.NewReader(os.Stdin);
		p := make([]byte, 8);
		answer,_ := reader.Read(p);
		fmt.Println(answer);
	}
}