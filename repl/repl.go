package repl // read-eval-print loop

import (
	"bufio"
	"fmt"
	"io"
	"slow-lang/evaluator"
	"slow-lang/lexer"
	"slow-lang/object"
	"slow-lang/parser"
)

const SNAIL = `
     /^\    /^\
    {  O}  {  O}
     \ /    \ /
     //     //       _------_
    //     //     ./~        ~-_
   / ~----~/     /              \
 /         :   ./       _---_    ~-
|  \________) :       /~     ~\   |
|        /    |      |  :~~\  |   |
|       |     |      |  \___-~    |
|        \ __/^\______\.        ./
 \                     ~-______-~\.
 .|                                ~-_
/_____________________________________~~____`

const PROMPT = "\n>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Fprint(out, PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, SNAIL)
	io.WriteString(out, "\nYou got it wrong!!\n  parser errors: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
