package parseelf

import (
	"encoding/binary"
	"fmt"
)

func (elf elfHeader32) ParseSegments() {
	for _, ph := range elf.ProgramHeader {
		fmt.Println("Index", elf.FileContents[ph.POffset])
		fmt.Printf("Type: 0x%x\n", binary.LittleEndian.Uint32(elf.FileContents[(ph.POffset+0x04):(ph.POffset+0x08)]))
		fmt.Println()
	}
}

func (elf elfHeader64) ParseSegments() {
	for _, ph := range elf.ProgramHeader {
		fmt.Println("Index", elf.FileContents[ph.POffset])
		fmt.Printf("Type:0x%x\n", elf.FileContents[ph.POffset+0x04])
		fmt.Println()
	}
}
