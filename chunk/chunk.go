package chunk

type OpCode uint8
type Value float32

const (
	OP_RETURN OpCode = iota
	OP_CONSTANT
	OP_NEGATE
	OP_ADD
	OP_SUBTRACT
	OP_MULTIPLY
	OP_DIVIDE
)

type LinesEncoded struct {
	Occurences int
	Line       int
}

type Chunk struct {
	Code      []OpCode
	Lines     []LinesEncoded
	Constants []Value
}

func (chunk *Chunk) WriteChunk(ch OpCode, line int) {
	chunk.Code = append(chunk.Code, ch)
	chunk.addLineInfo(line)
}

func (chunk *Chunk) addLineInfo(line int) {
	if len(chunk.Lines) == 0 {
		chunk.Lines = append(chunk.Lines, LinesEncoded{Occurences: 1, Line: line})
		return
	}

	prevLineIndex := len(chunk.Lines) - 1
	if chunk.Lines[prevLineIndex].Line == line {
		chunk.Lines[prevLineIndex].Occurences += 1
	} else {
		chunk.Lines = append(chunk.Lines, LinesEncoded{Occurences: 1, Line: line})
	}
}

func (chunk *Chunk) AddConstant(constant Value) OpCode {
	chunk.Constants = append(chunk.Constants, constant)
	return OpCode(uint8(len(chunk.Constants) - 1))
}

func InitChunk() *Chunk {
	return &Chunk{}
}