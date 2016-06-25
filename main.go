package main

import (
	brainfxck "github.com/mshr-h/gobrainfxck/lib"
)

func main() {
	var bf = brainfxck.LoadProgram("++++++[>++++++<-]>.")
	bf.Run()
	bf.Print()
}
