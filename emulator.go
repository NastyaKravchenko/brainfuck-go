package brainfuck

import "fmt"

func (e Emulator) execute() {
	for e.ip < uint(len(e.app)) {
		e.step()
	}
}

func (e *Emulator) step() error {
	switch command := e.app[e.ip]; command {
	case '>':
		if e.mp > MEM_MAX {
			return fmt.Errorf("(op >) memory pointer must be <= %d", MEM_MAX)
		}
		e.mp += 1
	case '<':
		if e.mp == 0 {
			return fmt.Errorf("(op <) memory pointer must be >= 0")
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
			return fmt.Errorf("error reading")
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
		return fmt.Errorf("invalid token")
	}
	e.ip += 1
	return nil
}
