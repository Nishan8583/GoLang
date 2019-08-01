package main

import (
	"fmt"
	"log"
	"os"

	"./lib/parseelf"
)

// ELF interface is used to integrate both 32 and 64 bit
type ELF interface {
	DisplayELF()
	ParseProgramHeader() (ELF, error)
	DisplayProgramHeader()
}

// Handle ERROR
func errorHandle(err error, msg string) {
	if err != nil {
		log.Println("Error in stage", msg, err)
		os.Exit(-1)
	}
}
func main() {
	if len(os.Args) < 2 {
		log.Println("Not enough arguments")
		os.Exit(-1)
	}
	elf, err := parseelf.ParseFile(os.Args[1])
	errorHandle(err, "Error while parsing file")
	elf, err = elf.ParseProgramHeader()
	elf.DisplayELF()
	fmt.Println("****")
	elf.DisplayProgramHeader()
	fmt.Println("****")
	s := elf.ParseSegments()
	t := s.(parseelf.ElfHeader64)
	//fmt.Println(s.(parseelf.ElfHeader64).SectionHeaders)
	t.DisplaySegments()
	fmt.Println(t.SectionHeaders[t.IndexOfSectionHeaderTable].Offset)
	fmt.Println("trying to get sectio name")

	for key, value := range t.SectionHeaders {
		if value.Name == "init" {
			fmt.Println("Got it ", key, value)
		}
	}

}
