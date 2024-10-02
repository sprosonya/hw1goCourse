package main

import (
	"fmt"
	"os"
	"uniq/reader"
	"uniq/settings"
	"uniq/uniq"
	"uniq/writer"
)

func main() {

	var mas []string
	var err error

	options, err := settings.GetOptions()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	in, err := settings.SetInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := settings.SetOutput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mas, err = reader.Reader(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res := uniq.Uniq(mas, options)

	err = writer.Writer(res, out)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
