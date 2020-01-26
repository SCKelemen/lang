package scanner

// Scanner is responsible for lexing source files
type Scanner struct {
	input   string
	head    int
	read    int
	current rune
}

func New(input string) *Scanner {
	s := &Scanner{input: input}
	s.readChar()
	return s
}

func (s *Scanner) readChar() {

}
