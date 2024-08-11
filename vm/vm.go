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

type Stack struct {
	Stack [256]chunk.Value
	StackTopIndex int
}
func (stack *Stack) Push(value chunk.Value) {
	stack.Stack[stack.StackTopIndex] = value
	stack.StackTopIndex += 1
}

func (stack *Stack) Pop() chunk.Value{
	stack.StackTopIndex -= 1
	return stack.Stack[stack.StackTopIndex]
}

type VM struct {
	chunk *chunk.Chunk
	Stack *Stack
}
func (vm *VM) Interpret(c *chunk.Chunk) InterpretResult {
	vm.chunk = c
	return vm.run()
}

func (vm *VM)printStack(){
	fmt.Printf(" ")
	for _, value := range(vm.Stack.Stack[:vm.Stack.StackTopIndex]) {
		fmt.Printf("[ ")
		debug.PrintValue(value)
		fmt.Printf(" ]")
	}
	fmt.Printf("\n")
}

func (vm *VM) run() InterpretResult {
	instructions := vm.chunk.Code
	for instructionIndex := 0; instructionIndex < len(instructions); instructionIndex++ {
		opcode := instructions[instructionIndex]
		vm.printStack()  // print the current stack after each instruction
		switch opcode {
		case chunk.OP_CONSTANT:
			instructionIndex++
			constantIndex := vm.chunk.Code[instructionIndex];
			vm.Stack.Push(vm.chunk.Constants[constantIndex])
		case chunk.OP_RETURN:
			debug.PrintValue(vm.Stack.Pop())
			fmt.Printf("\n")
			return INTERPRET_OK
		case chunk.OP_NEGATE:
			vm.Stack.Push(-vm.Stack.Pop())
		case chunk.OP_ADD:
			vm.Stack.Push(vm.Stack.Pop() + vm.Stack.Pop())
		case chunk.OP_MULTIPLY:
			vm.Stack.Push(vm.Stack.Pop() * vm.Stack.Pop())
		case chunk.OP_SUBTRACT:
			vm.Stack.Push(vm.Stack.Pop() - vm.Stack.Pop())
		case chunk.OP_DIVIDE:
			vm.Stack.Push(vm.Stack.Pop() / vm.Stack.Pop())
		}
	}

	return INTERPRET_RUNTIME_ERROR
}

func InitVM() *VM {
	stack := Stack{StackTopIndex: 0}
	return &VM{Stack: &stack}
}

func FreeVM(vm *VM) {

}
