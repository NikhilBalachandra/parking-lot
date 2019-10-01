package main

import (
	"fmt"
	"io"
	"os"
	"parking_lot/dao"
	"parking_lot/parser"
	"parking_lot/processor"
	"sync"
)

func main() {
	allocator := processor.NewNearestAllocator()
	storage := dao.InMemoryStorage{}
	mu := sync.Mutex{}

	argsWithoutProg := os.Args[1:]
	var fileArgument string

	if len(argsWithoutProg) > 1 {
		fmt.Fprintf(os.Stderr, "%s", "Incorrect usage")
		os.Exit(-1)
	}

	if len(argsWithoutProg) == 1 {
		fileArgument = argsWithoutProg[0]
	}

	if fileArgument != "" {
		inputFile, err := os.OpenFile(fileArgument, os.O_RDONLY, os.ModePerm)
		if err != nil {
			os.Exit(-1)
		}
		tokenizer := parser.NewTokenizer(inputFile)
		runNonInteractive(&tokenizer, &allocator, &storage, &mu)
	} else {
		tokenizer := parser.NewTokenizer(os.Stdin)
		runInteractive(&tokenizer, &allocator, &storage, &mu)
	}
}

// runInteractive inits the program in the interactive mode.
func runInteractive(tokenizer *parser.Tokenizer, allocator processor.Allocator, storage dao.Storage, mu *sync.Mutex) {
	for {
		fmt.Printf("%s", "$ ") // Print prompt
		out, err := processor.Process(tokenizer, mu, allocator, storage)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
		fmt.Printf("%s", out)
	}
}

// runInteractive inits the program in the non-interactive mode. In non-interactive mode,
// any errors processing the input will terminate the program.
func runNonInteractive(tokenizer *parser.Tokenizer, allocator processor.Allocator, storage dao.Storage, mu *sync.Mutex) {
	for {
		out, err := processor.Process(tokenizer, mu, allocator, storage)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(-2)
		}
		fmt.Print(out)
	}
}
