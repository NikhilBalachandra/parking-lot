package parser

import (
	"bufio"
	"io"
)

// Tokenizer splits input into commands and it's arguments.
type Tokenizer struct {
	scanner *bufio.Scanner
}

// NewTokenizer builds a new Tokenizer to split input into the tokens.
func NewTokenizer(r io.Reader) Tokenizer {
	p := Tokenizer{
		scanner: bufio.NewScanner(r),
	}
	p.scanner.Split(bufio.ScanLines)
	return p
}

// NextToken returns the next token (Line entry) from the input.
// Next invocation of this method will result in error io.EOF.
func (p *Tokenizer) NextToken() ([]byte, error) {
	tokenAvailable := p.scanner.Scan()
	err := p.scanner.Err()
	if tokenAvailable && err == nil {
		token := p.scanner.Bytes()
		return token, err
	} else if p.scanner.Err() == nil {
		return nil, io.EOF
	}
	return nil, err
}
