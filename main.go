package main

import "fmt"

type Brainfxck struct {
	dp      int
	ip      int
	memory  [16]int
	program string
}

func loadProgram(program string) *Brainfxck {
	var bf Brainfxck
	bf.dp = 0
	bf.ip = 0
	bf.program = program
	return &bf
}

func (bf *Brainfxck) execute() {
	for len(bf.program) > bf.ip {
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
}

func main() {
	var bf = loadProgram("++++++[>++++++<-]>.")
	bf.execute()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("ip      :", bf.ip)
	fmt.Println("dp      :", bf.dp)
	fmt.Println("program :", bf.program)
	fmt.Println("memory  :", bf.memory)
	return
}
