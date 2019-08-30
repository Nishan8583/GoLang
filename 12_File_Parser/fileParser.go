package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"./lib/parseelf"
)

var sh, ph, elfh, dis bool
var filename string

func init() {
	flag.BoolVar(&sh, "section-header", false, "./fileparser --file /path/to/file --section-header")
	flag.BoolVar(&ph, "program-header", false, "./fileparser --file /path/to/file --program-header")
	flag.StringVar(&filename, "file", "", "./fileparser --file /path/to/file")
	flag.BoolVar(&elfh, "elf-header", false, "./fileparser --file /path/to/file --elf-header")
	flag.BoolVar(&dis, "disassemble", false, "./fileparser --file /path/to/file --dissassemble")


}

// Handle ERROR
func errorHandle(err error, msg string) {
	if err != nil {
		log.Println("Error in stage", msg, err)
		os.Exit(-1)
	}
}

// The main function of program
func main() {
	flag.Parse()
	if len(filename) == 0 {
		fmt.Println("please provide file name")
		os.Exit(-1)
	}

	elf, err := parseelf.ParseFile(filename)
	errorHandle(err, "Error while parsing file")

	elf = elf.ParseProgramHeader()
	elf = elf.ParseSegments()

	if (ph) {
		elf.DisplayProgramHeader()
	}
	if (sh) {
		elf.DisplaySegments()
	}
	if (dis) {
		elf.Disassemble()
	}
	if (elfh) {
		elf.DisplayELF()
	}

}

