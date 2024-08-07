package chunk

type OpCode uint8
type Value float32

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
)

type Chunk struct {
	Code      []OpCode
	Constants []Value
}

func (chunk *Chunk) WriteChunk(ch OpCode) {
	chunk.Code = append(chunk.Code, ch)
}

func (chunk *Chunk) AddConstant(constant Value) OpCode {
	chunk.Constants = append(chunk.Constants, constant)
	return OpCode(uint8(len(chunk.Constants) - 1))
}

func InitChunk() *Chunk {
	return &Chunk{Code: []OpCode{}}
}