package main

import (
	"fmt"
	"strings"
)

type Brainfxck struct {
	dp      int
	ip      int
	memory  []int
	program string
}

func loadProgram(program string) *Brainfxck {
	var bf Brainfxck
	bf.dp = 0
	bf.ip = 0
	bf.memory = make([]int, 256)
	bf.program = program
	return &bf
}

func (bf *Brainfxck) execute() {
	switch bf.program[bf.ip] {
	case '+':
		bf.memory[bf.dp]++
	case '-':
		bf.memory[bf.dp]--
	case '>':
		bf.dp++
	case '<':
		bf.dp--
	case '.':
		fmt.Printf("%c", bf.memory[bf.dp])
	case '[':
		{
			if bf.memory[bf.dp] == 0 {
				var count = 1
				for count > 0 {
					bf.ip++
					switch bf.program[bf.ip] {
					case '[':
						count++
					case ']':
						count--
					}
				}
			}
		}
	case ']':
		{
			var count = 1
			for count > 0 {
				bf.ip--
				switch bf.program[bf.ip] {
				case '[':
					count--
				case ']':
					count++
				}
			}
			bf.ip--
		}
	}
	bf.ip++
}

func (bf *Brainfxck) run() {
	for bf.ip < len(bf.program) {
		bf.execute()
	}
}

func (bf *Brainfxck) print() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("dp      :", bf.dp)
	fmt.Println("program :", bf.program)
	fmt.Println("         ", strings.Repeat(" ", bf.ip-2), "^")
	fmt.Println("address         content")
	printMemory(bf.memory)
}

func printMemory(memory []int) {

	/*
		0x0000 00 24 00 00 00 00 00 00
	*/
	fmt.Printf("      ")
	for i := 0; i < 16; i++ {
		fmt.Printf(" %2X", i)
	}
	fmt.Println()

	for i := 0; i < len(memory); i += 16 {
		fmt.Printf("0x%.4X ", i)
		for _, v := range memory[i : i+16] {
			fmt.Printf("%.2X ", v)
		}
		fmt.Println()
	}
}

func main() {
	var bf = loadProgram("++++++[>++++++<-]>.")
	bf.run()
	bf.print()

	return
}
