package main

import "os"

func main() {
	if err := ParseArguments(); err != nil {
		Eprintln(err)

		os.Exit(1)
	}
}
