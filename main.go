package main

import (
	"fmt"
	"os"

	"github.com/MuradyanArtem/go-hw/tree/making-hw-1/hw-1.2/rpn"
)

func read() (out string) {
	if _, err := fmt.Scanln(&out); err != nil {
		fmt.Fprintf(os.Stderr, "error-> %v\n", err)
		os.Exit(1)
	}
	return
}

func main() {
	expression, err := rpn.New(read())
	if err != nil {
		fmt.Fprintf(os.Stderr, "error-> %v\n", err)
		os.Exit(1)
	}

	res, err := rpn.Calculate(expression)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error-> %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, res)
}
