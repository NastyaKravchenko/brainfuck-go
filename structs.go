package brainfuck

import "io"

type Emulator struct {
	app   []byte
	mem   [MEM_MAX]byte
	jumps []uint

	ip uint
	mp uint

	input io.Reader
}
