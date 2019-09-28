package parser

import (
	"bufio"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestTokenizer_NextToken(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		want        []byte
		wantErr     bool
		wantErrType error
	}{
		{
			name: "Empty String returns io.EOF", input: "",
			want: nil, wantErr: true, wantErrType: io.EOF,
		},
		{
			name: "Returns characters without newline", input: "park KA-01-HH-1234\n",
			want: []byte("park KA-01-HH-1234"), wantErr: false,
		},
		{
			name: "Returns left over characters on EOF (without newline)", input: "park KA-02-HH-1234",
			want: []byte("park KA-02-HH-1234"), wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Tokenizer{
				scanner: bufio.NewScanner(strings.NewReader(tt.input)),
			}
			got, err := p.NextToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("NextToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && (err != nil) && err != tt.wantErrType {
				t.Errorf("NextToken() got = %v, want %v", err, tt.wantErrType)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NextToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTokenizer_NextTokenInvocationAfterEOFResultsInEOF(t *testing.T) {
	p := &Tokenizer{
		scanner: bufio.NewScanner(strings.NewReader("status\n")),
	}
	_, _ = p.NextToken()
	_, err := p.NextToken()

	if err != io.EOF {
		t.Errorf("NextToken() got = %v, want %v", err, io.EOF)
	}
}
