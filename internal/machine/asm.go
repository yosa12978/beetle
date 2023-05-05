package machine

const (
	PUSH uint8 = iota + 1 //with operand
	POP
	DUP
	SWAP
	LOAD  //with operand
	STORE //with operand
	INC
	DEC
	ADD
	SUB
	MUL
	DIV
	AND
	OR
	XOR
	NOT
	SHR
	SHL
	JMP //with operand
	JZ  //with operand
	JNZ //with operand
	HALT
)
