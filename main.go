package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/MuradyanArtem/go-hw/tree/making-hw-1/hw-1.2/rpn"
)

func write(res int) {
	writer := bufio.NewWriter(os.Stdout)

	_, err := writer.WriteString(strconv.Itoa(res))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error-> %v\n", err)
		os.Exit(1)
	}
	writer.Flush()
}

func read() (out string) {
	_, err := fmt.Scanln(&out)
	if err != nil {
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
	write(res)
}
