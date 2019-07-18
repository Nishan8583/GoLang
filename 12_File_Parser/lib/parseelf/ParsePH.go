package parseelf

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

var pType map[int]string

type PH32 struct {
	PType    string
	POffset  uint32
	Pvaddr   uint32
	PAddr    uint32
	PFileZ   uint32
	PMemSize uint32
	PFlag    uint32
	PAlign   uint32
}

type PH64 struct {
	PType    string
	PFlag    uint32
	POffset  uint64
	Pvaddr   uint64
	PAddr    uint64
	PFileZ   uint64
	PMemSize uint64
	PAlign   uint64
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

func (elf elfHeader32) ParseProgramHeader() error {
	var pHeaders []PH32

	cont, err := ioutil.ReadFile(elf.Filename)
	if err != nil {
		fmt.Println("Error ", err)
		return err
	}

	firstOff := elf.ProgramHeaderOffset
	count := elf.ProgramHeaderNumberOfEntries

	//for count != 0; count-- {
	for counter := count; counter != 0; counter-- {
		offset := firstOff
		pHeaders = append(pHeaders, PH32{
			PType:    pType[int(cont[offset])],
			POffset:  binary.LittleEndian.Uint32(cont[(offset + 0x04):(offset + 0x08)]),
			Pvaddr:   binary.LittleEndian.Uint32(cont[(offset + 0x08):(offset + 0x0c)]),
			PAddr:    binary.LittleEndian.Uint32(cont[(offset + 0x0c):(offset + 0x10)]),
			PFileZ:   binary.LittleEndian.Uint32(cont[(offset + 0x10):(offset + 0x14)]),
			PMemSize: binary.LittleEndian.Uint32(cont[(offset + 0x14):(offset + 0x18)]),
			PFlag:    binary.LittleEndian.Uint32(cont[(offset + 0x18):(offset + 0x1c)]),
			PAlign:   binary.LittleEndian.Uint32(cont[(offset + 0x1c):(offset + 0x20)]),
		})
		firstOff = firstOff + 0x20
	}
	fmt.Println(pHeaders)
	return nil
}

func (elf elfHeader64) ParseProgramHeader() error {
	var pHeaders []PH64
	cont, err := ioutil.ReadFile(elf.Filename)
	if err != nil {
		fmt.Println("Error ", err)
		return err
	}

	firstOff := elf.ProgramHeaderOffset
	count := elf.ProgramHeaderNumberOfEntries

	//for count != 0; count-- {
	for counter := count; counter != 0; counter-- {
		offset := firstOff
		pHeaders = append(pHeaders, PH64{
			PType:    pType[int(cont[offset])],
			PFlag:    binary.LittleEndian.Uint32(cont[(offset + 0x04):(offset + 0x08)]),
			POffset:  binary.LittleEndian.Uint64(cont[(offset + 0x08):(offset + 0x10)]),
			Pvaddr:   binary.LittleEndian.Uint64(cont[(offset + 0x10):(offset + 0x18)]),
			PAddr:    binary.LittleEndian.Uint64(cont[(offset + 0x18):(offset + 0x20)]),
			PFileZ:   binary.LittleEndian.Uint64(cont[(offset + 0x20):(offset + 0x28)]),
			PMemSize: binary.LittleEndian.Uint64(cont[(offset + 0x28):(offset + 0x30)]),
			PAlign:   binary.LittleEndian.Uint64(cont[(offset + 0x30):(offset + 0x38)]),
		})
		firstOff = firstOff + 0x38
	}
	fmt.Println(pHeaders)

	//}
	return nil
}
