package utils

type TypeExp int

const (
	PRIMITIVE TypeExp = iota
	ARITHMETIC_OP
	LOGIC_OP
	RELATIONAL_OP
	ACCESS_ID
	INC
	DEC
	PARAMETER
	CALL_FUNC
	RETURN
	TERNARY
)
