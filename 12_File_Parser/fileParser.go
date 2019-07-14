package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Some Metainfo Help
var cpu map[byte]string
var endianness map[byte]string
var osVersion map[byte]string

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
}

// elfHeader will contin the ELF Header part
type elfHeader struct {
	Magic      []byte // The first 4 bytes
	Class      string // 32 Or 64 Bit
	Data       string // 1 or 2, little or Big
	Version    string // ELF version
	OSVersion  string // Target OS
	ABIVersion int    // ABI Version
}

//Unmarshal([]byte) will take a slice of byte and return elfHeader type
func Unmarshal(cont []byte) elfHeader {
	elf := elfHeader{
		Magic:   cont[0:4],
		Class:   cpu[cont[5]],
		Data:    endianness[cont[6]],
		Version: osVersion[cont[7]],
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
		elf := Unmarshal(content)
		fmt.Println(elf)
	} else {
		fmt.Println("I don't understand the file format")
	}

}

func main() {
	ParseFile(os.Args[1])
}
