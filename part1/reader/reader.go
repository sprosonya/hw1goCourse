package reader

import (
	"bufio"
	"fmt"
	"io"
)

func Reader(input io.Reader) ([]string, error) {
	in := bufio.NewScanner(input)
	var masOfLines []string = make([]string, 0)
	for in.Scan() {
		txt := in.Text()
		masOfLines = append(masOfLines, txt)
	}
	if len(masOfLines) == 0 {
		return nil, fmt.Errorf("Empty stream")
	} else {
		return masOfLines, nil
	}
}
