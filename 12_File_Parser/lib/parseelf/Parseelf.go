package parseelf

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
)

// Some Metainfo Help, Converts the number to their actual meaning
var cpu map[byte]int            // Maps to CPU architecture
var endianness map[byte]string  // Maps to the endianess of the
var osVersion map[byte]string   // Maps to number to OS version
var fileType map[int]string     // Maps number to file type
var machineType map[byte]string // Maps number to machine type

//init() initalizes the map that pooints the values to their respective
func init() {
	cpu = map[byte]int{
		1: 0x32,
		2: 0x64,
	}
	endianness = map[byte]string{
		1: "LittleEndian",
		2: "BigEndian",
	}
	osVersion = map[byte]string{
		0x00: "SystemV",
		0x01: "HP-UX",
		0x02: "NETBSD",
		0x03: "LINUX",
		0x04: "GNU Hurd",
		0x06: "Solaris",
		0x07: "AIX",
		0x08: "IRIX",
		0x09: "FreeBSD",
		0x0A: "Tru64",
		0x0B: "Novell Modesto",
		0x0c: "OpenBSD",
		0x0D: "OpenVMS",
		0x0E: "NonStop Kernel",
		0x0F: "AROS",
		0x10: "Fentix OS",
		0x11: "CloudABI",
	}
	fileType = map[int]string{
		0x00:   "ET_NONE",
		0x01:   "ET_REL",
		0x02:   "ET_EXEC",
		0x03:   "ET_DYN",
		0x04:   "ET_CORE",
		0xfe00: "ET_LOOS",
		0xfeff: "ET_HIOS",
		0xff00: "ET_LOPROC",
		0xffff: "ET_HIPROC",
	}
	machineType = map[byte]string{
		0x00: "No specific instruction set",
		0x02: "SPARC",
		0x03: "x86",
		0x08: "MIPS",
		0x14: "PowerPC",
		0x16: "S390",
		0x28: "ARM",
		0x2A: "SuperH",
		0x32: "IA-64",
		0x3E: "x86-64",
		0xB7: "AArch64",
		0xF3: "RISC-V",
	}
}

// ELF interface is used to integrate both 32 and 64 bit
type ELF interface {
	DisplayELF()
	ParseProgramHeader() error
}

// Both of the 32 bit and 64 bit has same field values but it will later be helpful for calling methods
// elfHeader for 32 bit will contin the ELF Header part
type elfHeader32 struct {
	Magic                        []byte // The first 4 bytes
	Class                        int    // 32 Or 64 Bit
	Data                         string // 1 or 2, little or Big
	Version                      byte   // ELF version
	OSVersion                    string // Target OS
	ABIVersion                   int    // ABI Version
	FileType                     string // The type of ELF File
	MachineType                  string // The Target ISA
	EVersion                     byte   // Elf Version
	EntryPoint                   uint32 // Entry postring Address
	ProgramHeaderOffset          uint32 // Offset to program header
	SectionHeaderOffset          uint32 // Offset to Section Header
	ELFHeaderSize                uint16 // Size Of this ELF Header
	ProgramHeaderSize            uint16 // SizeOfProgramHeader
	ProgramHeaderNumberOfEntries uint16 // NumberOfEntriesInProgramHeaders
	SectionHeaderSize            uint16 // SizeOfSectionHeader
	SectionHeaderNumberOfEntries uint16 // Number of Entries in Section Headers
	IndexOfSectionHeaderTable    byte   // Index to e_shstrndx
	Filename                     string // The filename of the file
}

// Elf Header for 64 bit
type elfHeader64 struct {
	Magic                        []byte // The first 4 bytes
	Class                        int    // 32 Or 64 Bit
	Data                         string // 1 or 2, little or Big
	Version                      byte   // ELF version
	OSVersion                    string // Target OS
	ABIVersion                   int    // ABI Version
	FileType                     string // The type of ELF File
	MachineType                  string // The Target ISA
	EVersion                     byte   // Elf Version
	EntryPoint                   uint64 // Entry postring Address
	ProgramHeaderOffset          uint64 // Offset to program header
	SectionHeaderOffset          uint64 // Offset to Section Header
	ELFHeaderSize                uint16 // Size Of this ELF Header
	ProgramHeaderSize            uint16 // SizeOfProgramHeader
	ProgramHeaderNumberOfEntries uint16 // NumberOfEntriesInProgramHeaders
	SectionHeaderSize            uint16 // SizeOfSectionHeader
	SectionHeaderNumberOfEntries uint16 // Number of Entries in Section Headers
	IndexOfSectionHeaderTable    byte   // Index to e_shstrndx
	Filename                     string // The filename
}

// DisplayELF() will display the elf Header Values
func (elf elfHeader32) DisplayELF() {
	fmt.Println("Magic Byte:", elf.Magic)
	fmt.Printf("Class: 0x%x\n", elf.Class)
	fmt.Println("Data:", elf.Data)
	fmt.Println("OSVersion:", elf.OSVersion)
	fmt.Println("ABI Version", elf.ABIVersion)
	fmt.Println("FileType:", elf.FileType)
	fmt.Println("ISA:", elf.MachineType)
	fmt.Println("ELF Version:", elf.EVersion)
	fmt.Printf("EntryPoint: 0x%x\n", elf.EntryPoint)
	fmt.Printf("ProgramHeaderOffset: 0x%x\n", elf.ProgramHeaderOffset)
	fmt.Printf("SectionHeaderOffset: 0x%x\n", elf.SectionHeaderOffset)
	fmt.Printf("ELF Header Size: 0x%x\n", elf.ELFHeaderSize)
	fmt.Printf("Program Header Size: 0x%x\n", elf.ProgramHeaderSize)
	fmt.Printf("Program Header Number Of Entries: 0x%x\n", elf.ProgramHeaderNumberOfEntries)
	fmt.Printf("Section Header Size: 0x%x\n", elf.SectionHeaderSize)
	fmt.Printf("Section Header Number Of Entries: 0x%x\n", elf.SectionHeaderNumberOfEntries)
	fmt.Printf("Index Section: 0x%x\n", elf.IndexOfSectionHeaderTable)
}

// DisplayELF() will display the elf Header Values
func (elf elfHeader64) DisplayELF() {
	fmt.Println("Magic Byte:", elf.Magic)
	fmt.Printf("Class: 0x%x\n", elf.Class)
	fmt.Println("Data:", elf.Data)
	fmt.Println("OSVersion:", elf.OSVersion)
	fmt.Println("ABI Version", elf.ABIVersion)
	fmt.Println("FileType:", elf.FileType)
	fmt.Println("ISA:", elf.MachineType)
	fmt.Println("ELF Version:", elf.EVersion)
	fmt.Printf("EntryPoint: 0x%x\n", elf.EntryPoint)
	fmt.Printf("ProgramHeaderOffset: 0x%x\n", elf.ProgramHeaderOffset)
	fmt.Printf("SectionHeaderOffset: 0x%x\n", elf.SectionHeaderOffset)
	fmt.Printf("ELF Header Size: %d\n", elf.ELFHeaderSize)
	fmt.Printf("Program Header Size: %d\n", elf.ProgramHeaderSize)
	fmt.Printf("Program Header Number Of Entries: %d\n", elf.ProgramHeaderNumberOfEntries)
	fmt.Printf("Section Header Size: %d\n", elf.SectionHeaderSize)
	fmt.Printf("Section Header Number Of Entries: %d\n", elf.SectionHeaderNumberOfEntries)
	fmt.Printf("Index Section: %d\n", elf.IndexOfSectionHeaderTable)
}

//ElfUnmarshal ... ELF will take a slice of byte and return elfHeader type
func ElfUnmarshal(cont []byte, filename string) (ELF, error) {

	// If 32 bit
	if cont[4] == 1 {
		elf := elfHeader32{
			Magic:                        cont[0:4],
			Class:                        cpu[cont[4]],
			Data:                         endianness[cont[5]],
			Version:                      cont[6],
			OSVersion:                    osVersion[cont[7]],
			FileType:                     fileType[int(cont[0x10])],
			MachineType:                  machineType[cont[0x12]],
			EVersion:                     cont[0x14],
			EntryPoint:                   binary.LittleEndian.Uint32(cont[0x18:0x1c]),
			ProgramHeaderOffset:          binary.LittleEndian.Uint32(cont[0x1c:0x20]),
			SectionHeaderOffset:          binary.LittleEndian.Uint32(cont[0x20:0x24]),
			ELFHeaderSize:                binary.LittleEndian.Uint16(cont[0x28:0x2a]),
			ProgramHeaderSize:            binary.LittleEndian.Uint16(cont[0x2a:0x2c]),
			ProgramHeaderNumberOfEntries: binary.LittleEndian.Uint16(cont[0x2C:0x2e]),
			SectionHeaderSize:            binary.LittleEndian.Uint16(cont[0x2e:0x30]),
			SectionHeaderNumberOfEntries: binary.LittleEndian.Uint16(cont[0x30:0x32]),
			IndexOfSectionHeaderTable:    cont[0x32],
			Filename:                     filename,
		}
		return elf, nil
	}
	elf := elfHeader64{ // If the file is 64 bit
		Magic:                        cont[0:4],
		Class:                        cpu[cont[4]],
		Data:                         endianness[cont[5]],
		Version:                      cont[6],
		OSVersion:                    osVersion[cont[7]],
		FileType:                     fileType[int(cont[0x10])],
		MachineType:                  machineType[cont[0x12]],
		EVersion:                     cont[0x14],
		EntryPoint:                   binary.LittleEndian.Uint64(cont[0x18:0x20]),
		ProgramHeaderOffset:          binary.LittleEndian.Uint64(cont[0x20:0x28]),
		SectionHeaderOffset:          binary.LittleEndian.Uint64(cont[0x28:0x30]),
		ELFHeaderSize:                binary.LittleEndian.Uint16(cont[0x34:0x36]),
		ProgramHeaderSize:            binary.LittleEndian.Uint16(cont[0x36:0x38]),
		ProgramHeaderNumberOfEntries: binary.LittleEndian.Uint16(cont[0x38:0x3a]),
		SectionHeaderSize:            binary.LittleEndian.Uint16(cont[0x3a:0x3c]),
		SectionHeaderNumberOfEntries: binary.LittleEndian.Uint16(cont[0x3c:0x3e]),
		IndexOfSectionHeaderTable:    cont[0x3e],
		Filename:                     filename,
	}
	return elf, nil
}

// ParseFile ... will parse the elf header and display the content
func ParseFile(filename string) (ELF, error) {

	var elf ELF

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return elf, errors.New("Error while Reading File ERROR:" + err.Error())
	}

	// Checking if the file is actually a valid ELF file
	if bytes.Equal(content[0:4], []byte{0x7f, 0x45, 0x4c, 0x46}) {
		elf, err = ElfUnmarshal(content, filename)
		if err != nil {
			return nil, errors.New("Error while calling the parsing function" + err.Error())
		}
		return elf, nil
	}
	return nil, errors.New("Not a valid ELF file")

}
