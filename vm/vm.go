package vm

import (
	"fmt"
	"interpreter/chunk"
	"interpreter/debug"
)

type InterpretResult uint8

const (
	INTERPRET_OK InterpretResult = iota
	INTERPRET_COMPILE_ERROR
	INTERPRET_RUNTIME_ERROR
)

type VM struct {
	Chunk *chunk.Chunk
	Ip int
}
func (vm *VM) Interpret(c *chunk.Chunk) InterpretResult{
	vm.Chunk = c
	// vm.Ip = 0
	return vm.run()
}

func (vm *VM) run() InterpretResult{
	instructions := vm.Chunk.Code
	for instructionIndex := 0; instructionIndex < len(instructions); instructionIndex++ {
		opcode := instructions[instructionIndex]
		switch opcode {
		case chunk.OP_CONSTANT:
			instructionIndex++
			constantIndex := vm.Chunk.Code[instructionIndex];
			debug.PrintValue(vm.Chunk.Constants[constantIndex])
			fmt.Printf("\n")
		case chunk.OP_RETURN:
			return INTERPRET_OK
		}
	}
	return INTERPRET_RUNTIME_ERROR
}

func InitVM() *VM {
	return &VM{}
}

func FreeVM(vm *VM) {

}
