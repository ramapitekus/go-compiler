package main

import (
	"fmt"
	"interpreter/vm"
	"os"
)


func repl(){
}

func runFile(){

}

func main() {
	args := os.Args
	vm := vm.InitVM()
	if (len(args) == 1){
		repl()
	} else if (len(args) == 2) {
		runFile()
	} else {
		fmt.Fprintf(os.Stderr, "Usage: clox [path]\n")
		os.Exit(64)
	}
	// vm.FreeVM()
	// vm.Interpret(c)
}