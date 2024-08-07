package debug

import (
	"fmt"
	"interpreter/chunk"
)

func DisassembleChunk(chunk *chunk.Chunk, name string) {
	fmt.Printf("== %s ==\n", name);
	
	offset := 0;
	for offset < len(chunk.Code){
		offset = DisassembleInstruction(chunk, offset);
	}
}

func DisassembleInstruction(c *chunk.Chunk, offset int) int {
	fmt.Printf("%04d ", offset);
	instruction := c.Code[offset];
	switch instruction {
	case chunk.OP_RETURN:
		return simpleInstruction("OP_RETURN", offset);
	case chunk.OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", offset, c);
	default:
		fmt.Printf("Unknown opcode %d\n", instruction);
		return offset + 1;
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1;
}

func constantInstruction(name string, offset int, c *chunk.Chunk) int {
	index := c.Code[offset + 1] // +1 because it's 2 bytes
	fmt.Printf("%-16s %4d '", name, index);
	printValue(c.Constants[index]);
	fmt.Printf("'\n");
	return offset + 2;
}

func printValue(value chunk.Value) {
	fmt.Printf("%g", value);
}