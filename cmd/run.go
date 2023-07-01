/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/TwiN/go-color"
	"github.com/jellycat-io/eevee/evaluator"
	"github.com/jellycat-io/eevee/lexer"
	"github.com/jellycat-io/eevee/object"
	"github.com/jellycat-io/eevee/parser"
	"github.com/jellycat-io/eevee/util"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Executes file at given path",
	Long:  `This command takes a filepath as argument`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		out := os.Stdout

		filepath := args[0]
		if _, err := os.Stat(filepath); err != nil {
			fmt.Printf(color.InRed("invalid filepath. got=%q"), filepath)
			os.Exit(1)
		}

		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			fmt.Printf(color.InRed("cannot read file: %q"), filepath)
		}
		source := string(buf)
		env := object.NewEnvironment()

		l := lexer.New(source)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			util.PrintParserErrors(out, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)

		if evaluated != nil {
			io.WriteString(out, color.InBold(color.InGreen(evaluated.Inspect())))
			io.WriteString(out, "\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
