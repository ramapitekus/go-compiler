package main

import (
	"interpreter/chunk"
	"interpreter/vm"
)

func main() {
	vm := vm.InitVM()
	c := chunk.InitChunk();
	c.WriteChunk(chunk.OP_RETURN, 123);
	index := c.AddConstant(1.2);
	c.WriteChunk(chunk.OP_CONSTANT, 123);
	c.WriteChunk(index, 123);
	vm.Interpret(c)
	// fmt.Printf("%4d ", c.Lines[1].Occurences);
	// debug.DisassembleChunk(c, "chunk");
	// vm.FreeVM()
}