package brainfuck

import (
	"os"
	"strings"
)

func readInstruction(filename string) Emulator {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var e = Emulator{ip: 0, mp: 0, input: os.Stdin}

	for _, char := range data {
		if IsValidToken(char) {
			e.app = append(e.app, char)
		}
	}

	e.jumps = make([]uint, len(e.app))

	var stack []uint
	for index, char := range e.app {
		if char == '[' {
			stack = append(stack, uint(index))
		}
		if char == ']' {
			var last uint
			last, stack = pop(stack)
			e.jumps[last] = uint(index)
			e.jumps[index] = last
		}
	}

	return e
}

func pop(a []uint) (uint, []uint) {
	return a[len(a)-1], a[:len(a)-1]
}

func IsValidToken(t byte) bool {
	return strings.Contains("+-><,[].,", string(t))
}
