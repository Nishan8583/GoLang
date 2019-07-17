package parseelf

import (
	"fmt"
	"io/ioutil"
)

var pType map[int]string

type PH struct {
	PType string
}

func init() {
	pType = map[int]string{
		0x00000000: "PT_NULL",    // Program header table entry unused
		0x00000001: "PT_LOAD",    // Loadable segment
		0x00000002: "PT_DYNAMIC", // 	Dynamic linking information
		0x00000003: "PT_INTERP",  //Interpreter information
		0x00000004: "PT_NOTE",    //  	Auxiliary information
		0x00000005: "PT_SHLIB",   //reserved
		0x00000006: "PT_PHDR",    //segment containing program header table itself
		0x60000000: "PT_LOOS",
		0x6FFFFFFF: "PT_HIOS",
		0x70000000: "PT_LOPROC",
		0x7FFFFFFF: "PT_HIPROC",
	}
}
func (elf elfHeader32) ParseProramHeader() error {

	cont, err := ioutil.ReadFile(elf.Filename)
	if err != nil {
		fmt.Println("Error ")
		return err
	}

	firstOff := elf.ProgramHeaderOffset
	//count := elf.ProgramHeaderNumberOfEntries

	//for count != 0; count-- {
	fmt.Println(cont[int(firstOff)])
	//}
	return nil
}
