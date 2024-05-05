package brainfuck

import (
	"flag"
	"fmt"
)

func main() {
	var input = flag.String("i", "", "input file")
	flag.Parse()

	if len(*input) == 0 {
		fmt.Println("usage: emulator -i <filename>")
		return
	}

	var emulator = readInstruction(*input)
	emulator.execute()
}
