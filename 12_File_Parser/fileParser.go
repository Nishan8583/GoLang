package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Some Metainfo Help
var cpu map[byte]string
var endianness map[byte]string
var osVersion map[byte]string
var fileType map[int]string
var machineType map[byte]string

//init() initalizes the map that pooints the values to their respective
func init() {
	cpu = map[byte]string{
		1: "0x32",
		2: "0x64",
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

// elfHeader will contin the ELF Header part
type elfHeader struct {
	Magic                        []byte // The first 4 bytes
	Class                        string // 32 Or 64 Bit
	Data                         string // 1 or 2, little or Big
	Version                      string // ELF version
	OSVersion                    string // Target OS
	ABIVersion                   int    // ABI Version
	FileType                     string // The type of ELF File
	MachineType                  string // The Target ISA
	EVersion                     byte   // Elf Version
	EntryPoint                   []byte // Entry postring Address
	ProgramHeaderOffset          []byte // Offset to program header
	SectionHeaderOffset          []byte // Offset to Section Header
	ELFHeaderSize                []byte // Size Of this ELF Header
	ProgramHeaderSize            []byte // SizeOfProgramHeader
	ProgramHeaderNumberOfEntries []byte // NumberOfEntriesInProgramHeaders
	SectionHeaderSize            []byte // SizeOfSectionHeader
	SectionHeaderNumberOfEntries []byte // Number of Entries in Section Headers
	IndexOfSectionHeaderTable    byte   // Index to e_shstrndx
}

func displayELF32(elf elfHeader) {
	fmt.Println("Magic Byte:", elf.Magic)
	fmt.Println("Class:", elf.Class)
	fmt.Println("Data:", elf.Data)
	fmt.Println("OSVersion:", elf.OSVersion)
	fmt.Println("ABI Version", elf.ABIVersion)
	fmt.Println("FileType:", elf.FileType)
	fmt.Println("ISA:", elf.MachineType)
	fmt.Println("ELF Version:", elf.EVersion)
	fmt.Printf("EntryPoint: 0x%x\n", binary.LittleEndian.Uint32(elf.EntryPoint))
	fmt.Printf("ProgramHeaderOffset: 0x%x\n", binary.LittleEndian.Uint32(elf.ProgramHeaderOffset))
}

//Unmarshal([]byte) will take a slice of byte and return elfHeader type
func elfUnmarshal(cont []byte) elfHeader {
	elf := elfHeader{
		Magic:       cont[0:4],
		Class:       cpu[cont[4]],
		Data:        endianness[cont[5]],
		Version:     osVersion[cont[7]],
		FileType:    fileType[int(cont[0x10])],
		MachineType: machineType[cont[0x12]],
		EVersion:    cont[0x14],
		EntryPoint:  cont[0x18:0x1c],
	}

	// If 32 bit
	if cont[4] == 1 {
		elf.ProgramHeaderOffset = cont[0x1c:0x20]
		elf.SectionHeaderOffset = cont[0x20:0x28]
		elf.ELFHeaderSize = cont[0x28:0x2a]
		elf.ProgramHeaderSize = cont[0x2a:0x2c]
		elf.ProgramHeaderNumberOfEntries = cont[0x2C:0x2e]
		elf.SectionHeaderSize = cont[0x2e:0x30]
		elf.SectionHeaderNumberOfEntries = cont[0x30:0x32]
		elf.IndexOfSectionHeaderTable = cont[0x32]
	}
	return elf
}

// Handle ERROR
func ErrorHandle(err error, msg string) {
	if err != nil {
		log.Println("Error in stage", msg, err)
		os.Exit(-1)
	}
}
func ParseFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	ErrorHandle(err, "Opening File ERROR:")

	if bytes.Equal(content[0:4], []byte{0x7f, 0x45, 0x4c, 0x46}) {
		fmt.Println("The File is probably ELF, now Unmarshalling")
		elf := elfUnmarshal(content)
		displayELF32(elf)
	} else {
		fmt.Println("I don't understand the file format")
	}

}

func main() {
	ParseFile(os.Args[1])
}
