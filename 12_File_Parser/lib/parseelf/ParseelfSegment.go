package parseelf

import (
	"encoding/binary"
	"fmt"
)

var shType map[uint32]string
var shFlag map[uint32]string

func init() {
	shType = map[uint32]string{
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

	shFlag = map[uint32]string{
		0x1:        "SHF_WRITE",            // 	Writable
		0x2:        "SHF_ALLOC",            //Occupies memory during execution
		0x4:        "SHF_EXECINSTR",        // 	Executable
		0x10:       "SHF_MERGE",            // 	Might be merged
		0x20:       "SHF_STRINGS",          //	Contains nul-terminated strings
		0x40:       "SHF_INFO_LINK",        // 	'sh_info' contains SHT index
		0x80:       "SHF_LINK_ORDER",       //Preserve order after combining
		0x100:      "SHF_OS_NONCONFORMING", // 	Non-standard OS specific handling required
		0x200:      "SHF_GROUP",            // 	Section is member of a group
		0x400:      "SHF_TLS",              //Section hold thread-local data
		0x0ff00000: "SHF_MASKOS",           // 	OS-specific
		0xf0000000: "SHF_MASKPROC",         //Processor-specific
		0x4000000:  "SHF_ORDERED",          //	Special ordering requirement (Solaris)
		0x8000000:  "SHF_EXCLUDE",          //	Section is excluded unless referenced or allocated (Solaris)
	}
}

// SH32 ...
type SH32 struct {
	IndexOffset  uint32 // An offset to a string in the .shstrtab section that represents the name of this section
	Type         string //Identifies the type of this header.
	Flags        string // Identifies the attributes of the section.
	VirAddr      uint32 // Virtual address of the section in memory, for sections that are loaded.
	Offset       uint32 // Offset of the section in the file image.
	Size         uint32 //  	Size in bytes of the section in the file image. May be 0.
	SectionIndex uint32 //  	Contains the section index of an associated section. This field is used for several purposes, depending on the type of section.
	ExtraInfo    uint32 // Contains extra information about the section
	Alignment    uint32 // Contains the required alignment of the section
	EntrySize    uint32 // 	Contains the size, in bytes, of each entry, for sections that contain fixed-size entries. Otherwise, this field contains zero.
	Name         string
	Off          []uint32
}

// SH64 ...
type SH64 struct {
	IndexOffset uint32 // An offset to a string in the .shstrtab section that represents the name of this section
	Type        string //Identifies the type of this header.
	Flags       string // Identifies the attributes of the section.
	VirAddr     uint64 // Virtual address of the section in memory, for sections that are loaded.
	Offset      uint64 // Offset of the section in the file image.
	Size        uint64 //  	Size in bytes of the section in the file image. May be 0.
	Alignment   uint64 // Contains the required alignment of the section
	EntrySize   uint64 // 	Contains the size, in bytes, of each entry, for sections that contain fixed-size entries. Otherwise, this field contains zero.
	Name        string
	Off         []uint64
}

func (elf ElfHeader32) ParseSegments() ELF {

	// If no section just return
	if elf.SectionHeaderOffset == 0 {
		fmt.Println("No Section Header")
		return nil
	}

	begin := elf.SectionHeaderOffset

	/*fmt.Printf("%-10s%s\n", "Index", "Type")
	for i := elf.SectionHeaderNumberOfEntries; i > 0; i-- {
		fmt.Printf("%-10d%s\n", elf.FileContents[begin], shType[binary.LittleEndian.Uint32(elf.FileContents[(i+0x04):(i+0x08)])])
	}*/

	sliceOfSH32 := []SH32{}
	for i := elf.SectionHeaderNumberOfEntries; i > 0; i-- {
		sh := SH32{
			IndexOffset:  binary.LittleEndian.Uint32(elf.FileContents[(begin):(begin + 0x04)]),
			Type:         shType[binary.LittleEndian.Uint32(elf.FileContents[(begin+0x04):(begin+0x08)])],
			Flags:        shFlag[binary.LittleEndian.Uint32(elf.FileContents[(begin+0x08):(begin+0x0c)])],
			VirAddr:      binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x0c):(begin + 0x10)]),
			Offset:       binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x10):(begin + 0x14)]),
			Size:         binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x14):(begin + 0x18)]),
			SectionIndex: binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x18):(begin + 0x1c)]),
			ExtraInfo:    binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x1c):(begin + 20)]),
			Alignment:    binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x20):(begin + 0x24)]),
			EntrySize:    binary.LittleEndian.Uint32(elf.FileContents[(begin + 0x24):(begin + 0x28)]),
			Off:          []uint32{(begin + 0x10), (begin + 0x14)},
		}
		sliceOfSH32 = append(sliceOfSH32, sh)
		begin = begin + 0x28
	}
	elf.SectionHeaders = sliceOfSH32
	// Now adding Section Name
	sectionHeaderTable := elf.SectionHeaders[elf.IndexOfSectionHeaderTable]
	begin2 := sectionHeaderTable.Offset // This is where the section header table is located in file
	for key, value := range elf.SectionHeaders {

		i := value.IndexOffset
		start := begin2 + i + 1
		char := ""
		for {
			if elf.FileContents[start] == 0 {
				break
			}
			char = char + string(elf.FileContents[start])
			start = start + 1

		}
		elf.SectionHeaders[key].Name = char
	}
	return elf
} ////

func (elf ElfHeader64) ParseSegments() ELF {
	if elf.SectionHeaderOffset == 0 {
		fmt.Println("No Section Header")
		return nil
	}

	begin := elf.SectionHeaderOffset

	sliceOfSH64 := []SH64{}
	for i := elf.SectionHeaderNumberOfEntries; i > 0; i-- {
		sh := SH64{
			IndexOffset: binary.LittleEndian.Uint32(elf.FileContents[(begin):(begin + 0x04)]),
			Type:        shType[binary.LittleEndian.Uint32(elf.FileContents[(begin+0x04):(begin+0x08)])],
			Flags:       shFlag[binary.LittleEndian.Uint32(elf.FileContents[(begin+0x08):(begin+0x10)])],
			VirAddr:     binary.LittleEndian.Uint64(elf.FileContents[(begin + 0x10):(begin + 0x18)]),
			Offset:      binary.LittleEndian.Uint64(elf.FileContents[(begin + 0x18):(begin + 0x20)]),
			Size:        binary.LittleEndian.Uint64(elf.FileContents[(begin + 0x20):(begin + 0x28)]),
			Alignment:   binary.LittleEndian.Uint64(elf.FileContents[(begin + 0x30):(begin + 0x38)]),
			EntrySize:   binary.LittleEndian.Uint64(elf.FileContents[(begin + 0x38):(begin + 0x40)]),
			Off:         []uint64{(begin + 0x18), (begin + 0x20)},
		}
		sliceOfSH64 = append(sliceOfSH64, sh)
		begin = begin + 0x40
	}
	elf.SectionHeaders = sliceOfSH64

	// Now adding Section Name
	sectionHeaderTable := elf.SectionHeaders[elf.IndexOfSectionHeaderTable]
	begin2 := sectionHeaderTable.Offset // This is where the section header table is located in file
	for key, value := range elf.SectionHeaders {

		i := value.IndexOffset
		start := begin2 + uint64(i) + 1
		char := ""
		for {
			if elf.FileContents[start] == 0 {
				break
			}
			char = char + string(elf.FileContents[start])
			start = start + 1

		}
		elf.SectionHeaders[key].Name = char
	}
	return elf

}

func (elf ElfHeader32) DisplaySegments() {
	fmt.Printf("\n-----Section Header------\n")
	fmt.Printf("%-18s%-18s%-20s%-20s%-20s%-20s%-20s%-20s%-20s%-20s\n", "ShName", "ShType", "ShFlags", "VirtualAddress", "FileOffset",
		"SectionSize", "ShLink")

	//
	for _, value := range elf.SectionHeaders {
		fmt.Printf("%-18d%-18s%-18s0x%-18x%-18x%-18d%-18d%-18d\n",
			value.IndexOffset, value.Type, value.Flags, value.VirAddr, value.Offset, value.Size)
	}
	fmt.Printf("-----Section Header------\n\n")

}

func (elf ElfHeader64) DisplaySegments() {
	fmt.Printf("\n-----Section Header------\n")
	fmt.Printf("%-18s%-18s%-20s%-20s%-20s%-20s%-20s%-20s%-20s%-20s\n", "ShName", "ShType", "ShFlags", "VirtualAddress", "FileOffset",
		"SectionSize", "ShLink")

	//
	for _, value := range elf.SectionHeaders {
		fmt.Printf("%-18s%-18s%-18s0x%-18x%-18x%-18d%-18d%-18d\n",
			value.Name, value.Type, value.Flags, value.VirAddr, value.Offset, value.Size)

	}
	fmt.Printf("-----Section Header------\n\n")

}
