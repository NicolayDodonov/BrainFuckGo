package main

import (
	"fmt"
	"os"
)

func main() {
	var array [30000]int // Массив
	var pArray int = 0   // Указатель коретки
	var stack [10]int
	var pStack int = 0
	var pProgram int = 0 // Программный указатель
	var program []byte   // Массив с символами программы
	loadFile(&program)
	lengProg := len(string(program))

	for pProgram <= lengProg-1 {
		switch rune(program[pProgram]) {
		case '>':
			pArray++
			if pArray >= 30000 {
				pArray = 0
			}
		case '<':
			pArray--
			if pArray <= 0 {
				pArray = 29999
			}
		case '+':
			array[pArray]++
		case '-':
			array[pArray]--
		case '.':
			fmt.Print(string(array[pArray]))
		case ',':
			array[pArray], _ = fmt.Scanln()
		case '[':
			if array[pArray] != 0 {
				stack[pStack] = pProgram
				pStack++
			} else {
				var b = 1
				for b > 0 {
					pProgram++
					if rune(program[pProgram]) == '[' {
						b++
					}
					if rune(program[pProgram]) == ']' {
						b--
					}
				}
			}
		case ']':
			if array[pArray] != 0 {
				pProgram = stack[pStack-1]
			} else {
				pStack--
			}
		default:
		}
		pProgram++
	}
	_, _ = fmt.Scanln()
}

func loadFile(program *[]byte) {
	var data []byte
	data, err := os.ReadFile("program.bf")

	if err != nil {
		fmt.Println("Cannot open file")
		os.Exit(-1)
	}
	*program = data
}
