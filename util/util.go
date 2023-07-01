package util

import (
	"io"

	"github.com/TwiN/go-color"
)

func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, color.InBold(color.InRed("parser errors:\n")))
	for _, msg := range errors {
		io.WriteString(out, color.InRed("\t"+msg+"\n"))
	}
}
