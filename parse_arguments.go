package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

var options struct {
	format       string
	no_highlight bool
	style        string
	version      bool
}

func ParseArguments() error {
	pflag.StringVarP(&options.format, "format", "f", "json", "Set the output format")

	pflag.BoolVarP(&options.no_highlight, "no-highlight", "H", false, "Disable syntax highlighting")

	pflag.StringVarP(&options.style, "style", "s", "monokai", "Set the highlight style")

	pflag.BoolVarP(&options.version, "version", "v", false, "Print version")

	pflag.Parse()

	if options.version {
		fmt.Println("capra 0.1.1")

		return nil
	}

	maps := ToMaps(pflag.Args())

	bytes, err := ToBytes(maps, options.format)

	if err != nil {
		return err
	}

	if options.no_highlight {
		Print(bytes)
	} else {
		HighLightedPrint(bytes, options.format, options.style)
	}

	return nil
}
