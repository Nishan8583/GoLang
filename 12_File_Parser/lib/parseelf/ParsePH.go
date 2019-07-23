package parseelf

import (
	"encoding/binary"
	"fmt"
)

var pType map[int]string // Maps the integer to its meaning

// PH32 is prorgram header type for 32 bit
type PH32 struct {
	PType     string
	POffset   uint32
	Pvaddr    uint32
	PAddr     uint32
	PFileSize uint32
	PMemSize  uint32
	PFlag     uint32
	PAlign    uint32
}

// PH64 is program header type for 64 bit
type PH64 struct {
	PType     string
	PFlag     uint32
	POffset   uint64
	Pvaddr    uint64
	PAddr     uint64
	PFileSize uint64
	PMemSize  uint64
	PAlign    uint64
}

// init() ... initalizes the program header type to their meaning
func init() {
	pType = map[int]string{
		0x00000000: "PT_NULL",    // Program header table entry unused
		0x00000001: "PT_LOAD",    // Loadable segment
		0x00000002: "PT_DYNAMIC", // 	Dynamic linking information
		0x00000003: "PT_INTERP",  //Interpreter information
		0x00000004: "PT_NOTE",    //  	Auxiliary information
		0x00000005: "PT_SHLIB",   //reserved
		0x00000006: "PT_PHDR",    //segment elf.FileContentsaining program header table itself
		0x60000000: "PT_LOOS",
		0x6FFFFFFF: "PT_HIOS",
		0x70000000: "PT_LOPROC",
		0x7FFFFFFF: "PT_HIPROC",
	}
}

// (elf elfHeader32) ParseProgramHeader() ... Parses the program header section and return an ELF type for 32 bit
// Since the interface methods works with value symantec, a local copy is created, so we need to return it.
func (elf elfHeader32) ParseProgramHeader() (ELF, error) {
	var pHeaders []PH32

	firstOff := elf.ProgramHeaderOffset       // The first progrma header
	count := elf.ProgramHeaderNumberOfEntries // numer of program header secitons

	// Loop for each program header
	for counter := count; counter != 0; counter-- {
		offset := firstOff
		pHeaders = append(pHeaders, PH32{
			PType:     pType[int(elf.FileContents[offset])],
			POffset:   binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x04):(offset + 0x08)]),
			Pvaddr:    binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x08):(offset + 0x0c)]),
			PAddr:     binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x0c):(offset + 0x10)]),
			PFileSize: binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x10):(offset + 0x14)]),
			PMemSize:  binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x14):(offset + 0x18)]),
			PFlag:     binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x18):(offset + 0x1c)]),
			PAlign:    binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x1c):(offset + 0x20)]),
		})
		firstOff = firstOff + 0x20 // For 32 bit the program header size is 0x20
	}
	elf.ProgramHeader = pHeaders
	return elf, nil
}

// (elf elfHeader32) ParseProgramHeader() ... Parses the program header section and return an ELF type for 64 bit
// Since the interface methods works with value symantec, a local copy is created, so we need to return it.
func (elf elfHeader64) ParseProgramHeader() (ELF, error) {
	var pHeaders []PH64

	firstOff := elf.ProgramHeaderOffset
	count := elf.ProgramHeaderNumberOfEntries

	//for count != 0; count-- {
	for counter := count; counter != 0; counter-- {
		offset := firstOff
		pHeaders = append(pHeaders, PH64{
			PType:     pType[int(elf.FileContents[offset])],
			PFlag:     binary.LittleEndian.Uint32(elf.FileContents[(offset + 0x04):(offset + 0x08)]),
			POffset:   binary.LittleEndian.Uint64(elf.FileContents[(offset + 0x08):(offset + 0x10)]),
			Pvaddr:    binary.LittleEndian.Uint64(elf.FileContents[(offset + 0x10):(offset + 0x18)]),
			PAddr:     binary.LittleEndian.Uint64(elf.FileContents[(offset + 0x18):(offset + 0x20)]),
			PFileSize: binary.LittleEndian.Uint64(elf.FileContents[(offset + 0x20):(offset + 0x28)]),
			PMemSize:  binary.LittleEndian.Uint64(elf.FileContents[(offset + 0x28):(offset + 0x30)]),
			PAlign:    binary.LittleEndian.Uint64(elf.FileContents[(offset + 0x30):(offset + 0x38)]),
		})
		firstOff = firstOff + 0x38 // For 32 bit the program header size is 0x20
	}
	elf.ProgramHeader = pHeaders
	return elf, nil
}

func (elf elfHeader32) DisplayProgramHeader() {
	fmt.Printf("%-18s%-18s%-10s%-18s%-18s%-18s%-18s%-18s\n", "Flag", "SegmentType", "Offset", "VirtualAddress", "PhysicalAddress", "SegmentSize", "SegmentSizeInMemory", "Alignment")

	for _, value := range elf.ProgramHeader {
		fmt.Printf("%-18d%-18s%-18d0x%-18x0x%-18x%-18d%-18d%-18d\n",
			value.PFlag, value.PType, value.POffset, value.Pvaddr, value.PAddr, value.PFileSize, value.PMemSize, value.PAlign)
	}
}

func (elf elfHeader64) DisplayProgramHeader() {
	for _, value := range elf.ProgramHeader {
		fmt.Printf("%-18d%-18s%-18d0x%-18x0x%-18x%-18d%-18d%-18d\n",
			value.PFlag, value.PType, value.POffset, value.Pvaddr, value.PAddr, value.PFileSize, value.PMemSize, value.PAlign)
	}
}
