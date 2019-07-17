package main

import (
	"log"
	"os"

	"./lib/parseelf"
)

// Handle ERROR
func errorHandle(err error, msg string) {
	if err != nil {
		log.Println("Error in stage", msg, err)
		os.Exit(-1)
	}
}
func main() {
	elf, err := parseelf.ParseFile(os.Args[1])
	errorHandle(err, "Error while parsing file")
	elf.DisplayELF()
	elf.ParseProgramHeader()
}
