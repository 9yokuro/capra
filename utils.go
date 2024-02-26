package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Eprintln(err error) {
	error_red := color.RedString("error")

	fmt.Fprintf(os.Stderr, "%s: %s\n", error_red, err)
}
