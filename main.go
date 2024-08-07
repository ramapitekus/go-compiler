package main

import (
	"interpreter/chunk"
	"interpreter/debug"
)

func main() {
	c := chunk.InitChunk();
	c.WriteChunk(chunk.OP_RETURN);
	index := c.AddConstant(1.2);
	c.WriteChunk(chunk.OP_CONSTANT);
	c.WriteChunk(index);
	debug.DisassembleChunk(c, "chunk");
}