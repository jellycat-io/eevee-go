package repl

import (
	"bufio"
	"fmt"
	"io"
	"os/user"

	"github.com/TwiN/go-color"
	"github.com/jellycat-io/eevee/evaluator"
	"github.com/jellycat-io/eevee/lexer"
	"github.com/jellycat-io/eevee/object"
	"github.com/jellycat-io/eevee/parser"
	"github.com/jellycat-io/eevee/util"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	fmt.Print(color.InBold(color.InBlue(fmt.Sprintf("\nEevee REPL 0.1.0 - Welcome %s\n", user.Username))))

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
			util.PrintParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, color.InBold(color.InGreen(evaluated.Inspect())))
			io.WriteString(out, "\n")
		}
	}
}
