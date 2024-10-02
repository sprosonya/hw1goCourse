package writer

import (
	"bufio"
	"io"
)

func Writer(masOfLines []string, output io.Writer) error {
	out := bufio.NewWriter(output)
	for i := 0; i < len(masOfLines); i++ {
		if _, err := out.WriteString(masOfLines[i] + "\n"); err != nil {
			return err
		}
	}
	err := out.Flush()
	return err
}
