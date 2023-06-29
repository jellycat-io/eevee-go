package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/TwiN/go-color"
	"github.com/jellycat-io/eevee/evaluator"
	"github.com/jellycat-io/eevee/lexer"
	"github.com/jellycat-io/eevee/parser"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, color.InBold(PROMPT))
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)

		if evaluated != nil {
			io.WriteString(out, color.InBold(color.InGreen(evaluated.Inspect())))
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, color.InBold(color.InRed("parser errors:\n")))
	for _, msg := range errors {
		io.WriteString(out, color.InRed("\t"+msg+"\n"))
	}
}
