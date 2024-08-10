package main

import (
	"interpreter/chunk"
	"interpreter/debug"
	"interpreter/vm"
)

func main() {
	vm := vm.InitVM()
	c := chunk.InitChunk();
	index := c.AddConstant(1.2);
	c.WriteChunk(chunk.OP_CONSTANT, 123);
	c.WriteChunk(index, 123);
	index = c.AddConstant(3.4);
	c.WriteChunk(chunk.OP_CONSTANT, 123);
	c.WriteChunk(index, 123);

	c.WriteChunk(chunk.OP_ADD, 123)

	index = c.AddConstant(5.6);
	c.WriteChunk(chunk.OP_CONSTANT, 123);
	c.WriteChunk(index, 123);

	c.WriteChunk(chunk.OP_DIVIDE, 123)
	c.WriteChunk(chunk.OP_NEGATE, 123)
	c.WriteChunk(chunk.OP_RETURN, 123)
	vm.Interpret(c)
	debug.DisassembleChunk(c, "chunk");
}