package brainfuck

import "testing"

func TestIsValid(t *testing.T) {

	var tests = []struct {
		name  string
		input byte
		want  bool
	}{
		{"ValidToken", byte('.'), true},
		{"NotValidToken", byte('7'), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := IsValidToken(tt.input)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

func TestCompilation(t *testing.T) {
	var err error
	var emul = readInstruction("examples/hello.bf")
	for emul.ip < uint(len(emul.app)) {
		err = emul.step()
	}
	if err != nil {
		t.Errorf("there's a mistake in code")
	}
}
