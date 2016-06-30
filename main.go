package main

import (
	"os"

	brainfxck "github.com/mshr-h/gobrainfxck/lib"
)

func main() {
	var code = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++." // "Hello World!"
	var bf = brainfxck.LoadProgram(os.Stdin, os.Stdout, code)
	bf.Run()
	bf.Print()
}
