package main

import (
	brainfxck "github.com/mshr-h/gobrainfxck/lib"
)

func main() {
	var code = "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++." // "Hello World!"
	var bf = brainfxck.LoadProgram(code)
	bf.Run()
	bf.Print()
}
