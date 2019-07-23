package parseelf

import (
	"encoding/binary"
	"fmt"
)

var shNameMap map[uint32]string

func init() {
	shNameMap = map[uint32]string{
		0x0:        "SHT_NULL",          //Section header table entry unused
		0x1:        "SHT_PROGBITS",      //Program data
		0x2:        "SHT_SYMTAB",        //Symbol table
		0x3:        "SHT_STRTAB",        //String table
		0x4:        "SHT_RELA",          //Relocation entries with addends
		0x5:        "SHT_HASH",          //Symbol hash table
		0x6:        "SHT_DYNAMIC",       //Dynamic linking information
		0x7:        "SHT_NOTE",          //Notes
		0x8:        "SHT_NOBITS",        //Program space with no data (bss)
		0x9:        "SHT_REL",           //Relocation entries, no addends
		0x0A:       "SHT_SHLIB",         //Reserved
		0x0B:       "SHT_DYNSYM",        //Dynamic linker symbol table
		0x0E:       "SHT_INIT_ARRAY",    //Array of constructors
		0x0F:       "SHT_FINI_ARRAY",    //Array of destructors
		0x10:       "SHT_PREINIT_ARRAY", //Array of pre-constructors
		0x11:       "SHT_GROUP",         //Section group
		0x12:       "SHT_SYMTAB_SHNDX",  //Extended section indices
		0x13:       "SHT_NUM",           //Number of defined types.
		0x60000000: "SHT_LOOS",          //Start OS-specific.
	}
}
func (elf elfHeader32) ParseSegments() {
	if elf.SectionHeaderOffset == 0 {
		fmt.Println("No Section Header")
		return
	}
	fmt.Printf("%-10s%s\n", "Index", "Type")
	for i := elf.SectionHeaderNumberOfEntries; i > 0; i-- {
		fmt.Printf("%-10d%s\n", elf.FileContents[i], shNameMap[binary.LittleEndian.Uint32(elf.FileContents[(i+0x04):(i+0x08)])])
	}
} //

func (elf elfHeader64) ParseSegments() {
	if elf.SectionHeaderOffset == 0 {
		fmt.Println("No Section Header")
		return
	}
	fmt.Printf("%-10s%s\n", "Index", "Type")
	for i := elf.SectionHeaderNumberOfEntries; i > 0; i-- {
		fmt.Printf("%-10d%s\n", elf.FileContents[i], shNameMap[binary.LittleEndian.Uint32(elf.FileContents[(i+0x04):(i+0x08)])])
	}
}
