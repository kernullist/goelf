package elf

import (
	"unsafe"
)

// http://lxr.free-electrons.com/source/arch/x86/include/asm/elf.h#L16
type ELFGReg uint

const ELF_NGREG = unsafe.Sizeof(UserRegs{}) / unsafe.Sizeof(ELFGReg(0))

type ELFGRegSet [ELF_NGREG]ELFGReg
