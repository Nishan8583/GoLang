package parseelf

import (
	"fmt"

	capstone "github.com/bnagy/gapstone"
)

// Disassmble... the function that disassembles the code for 64 bit programs
func (elf ElfHeader64) Disassemble() {
	for _, value := range elf.SectionHeaders {
		if value.Type == "SHT_PROGBITS" {
			engine, err := capstone.New(
				capstone.CS_ARCH_X86, // x86 intel
				capstone.CS_MODE_64,  // 64bit mode
			)

			if err != nil {
				fmt.Println("Could not create engine", err)
				return
			}
			defer engine.Close()

			fmt.Printf("\n\nDisassembly For %s section\n\n", value.Name)

			inss, err := engine.Disasm(elf.FileContents[value.Offset:(value.Offset+value.Size)], value.VirAddr, 0)
			if err != nil {
				fmt.Println("Error while disassembling ", err)
				continue
			}

			for _, insn := range inss {
				fmt.Printf("0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
			}

		}
	}
}

/*
// Disassmble... the function that disassembles the code for 32 bit programs
func (elf ElfHeader32) Disassemble() {
	for _, value := range elf.SectionHeaders {
		if value.Type == "SHT_PROGBITS" {
			engine, err := capstone.New(
				capstone.CS_ARCH_X86, // x86 intel
				capstone.CS_MODE_32,  // 32bit mode
			)

			if err != nil {
				fmt.Println("Could not create engine", err)
				return
			}
			defer engine.Close()

			fmt.Printf("\n\nDisassembly For %s section\n\n", value.Name)

			inss, err := engine.Disasm(elf.FileContents[value.Offset:(value.Offset+value.Size)], value.VirAddr, 0)
			if err != nil {
				fmt.Println("Error while disassembling ", err)
				continue
			}

			for _, insn := range inss {
				fmt.Printf("0x%x:\t%s\t\t%s\n", insn.Address, insn.Mnemonic, insn.OpStr)
			}

		}
	}
}
*/
