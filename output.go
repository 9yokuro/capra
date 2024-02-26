package main

import (
	"fmt"
	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
	"os"
)

func Print(data [][]byte) {
	for _, d := range data {
		fmt.Println(string(d))
	}
}

func HighLightedPrint(data [][]byte, format string, style string) {
	for _, d := range data {
		lexer := lexers.Get(format)

		if lexer == nil {
			lexer = lexers.Fallback
		}

		lexer = chroma.Coalesce(lexer)

		style := styles.Get(style)

		if style == nil {
			style = styles.Fallback
		}

		formatter := formatters.TTY256

		if formatter == nil {
			formatter = formatters.Fallback
		}

		iterator, err := lexer.Tokenise(nil, string(d))

		if err != nil {
			Eprintln(err)
		}

		if err := formatter.Format(os.Stdout, style, iterator); err != nil {
			Eprintln(err)
		}

		fmt.Println()
	}
}
