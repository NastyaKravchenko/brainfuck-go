package brainfuck

import "fmt"

func (e *Emulator) step() {
	switch command := e.app[e.ip]; command {
	case '>':
		if e.mp > MEM_MAX {
			errmsg := fmt.Sprintf("(op >) memory pointer must be <= %d", MEM_MAX)
			panic(errmsg)
		}
		e.mp += 1
	case '<':
		if e.mp == 0 {
			panic("(op <) memory pointer must be >= 0")
		}
		e.mp -= 1
	case '+':
		e.mem[e.mp] += 1
	case '-':
		e.mem[e.mp] -= 1
	case '.':
		fmt.Printf("%c", e.mem[e.mp])
	case ',':
		char := make([]byte, 1)
		_, err := e.input.Read(char)
		if err != nil {
			panic(err)
		}
		e.mem[e.mp] = char[0]
	case '[':
		if e.mem[e.mp] == 0 {
			e.ip = e.jumps[e.ip]
		}
	case ']':
		if e.mem[e.mp] != 0 {
			e.ip = e.jumps[e.ip]
		}
	default:
	}
	e.ip += 1
}
