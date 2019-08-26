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

// The main function of program
func main() {

	if len(os.Args) < 2 {
		log.Println("Not enough arguments")
		os.Exit(-1)
	}

	elf, err := parseelf.ParseFile(os.Args[1])
	errorHandle(err, "Error while parsing file")

	elf, err = elf.ParseProgramHeader()
	elf.DisplayELF()
	//elf.DisplayProgramHeader()
	//elf.DisplaySegments()
	//elf.Disassemble()

}
