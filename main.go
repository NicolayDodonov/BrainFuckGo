package main

import (
	"fmt"
	"os"
)

func main() {
	var array [30000]int // Массив
	var pArray int = 0   // Указатель коретки
	var pProgram int = 0 // Программный указатель
	var program []byte   // Массив с символами программы
	loadFile(&program)
	lengProg := len(string(program))

	for pProgram <= lengProg-1 {
		switch rune(program[pProgram]) {
		case '>':
			pArray++
		case '<':
			pArray--
		case '+':
			array[pArray]++
		case '-':
			array[pArray]--
		case '.':
			fmt.Print(string(array[pArray]))
		case ',':
			array[pArray], _ = fmt.Scan()
		case '[':
			fmt.Print('|')
		case ']':
			fmt.Print('|')
		case 10:
		case 13:
			//Игнорируем символы переноса и возврата коретки
		default:
			fmt.Print("\nНеопознанный символ: ")
			fmt.Println(string(program[pProgram]))
			panic(-1)
		}
		pProgram++
	}
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
