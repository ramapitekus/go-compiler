package compiler

import "interpreter/scanner"

func Compile(source string) {
	scanner.InitScanner(source)
}