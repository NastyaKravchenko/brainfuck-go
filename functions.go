package brainfuck

import "strings"

func (e Emulator) execute() {
	for e.ip < uint(len(e.app)) {
		e.step()
	}
}

func pop(a []uint) (uint, []uint) {
	return a[len(a)-1], a[:len(a)-1]
}

func IsValidToken(t byte) bool {
	return strings.Contains("+-><,[].,", string(t))
}
