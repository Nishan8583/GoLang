package main;

import (
	"fmt";
	"bufio";
	"os";
)

func main() {
	reader := bufio.NewReader(os.Stdin);
	s,l,_ := reader.ReadLine()
	fmt.Println(string(s),l);
}