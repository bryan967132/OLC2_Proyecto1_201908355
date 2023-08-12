package utils

type TypeExp int

const (
	PRIMITIVE     TypeExp = 0
	ARITHMETIC_OP TypeExp = 1
	LOGIC_OP      TypeExp = 2
	RELATIONAL_OP TypeExp = 3
	ACCESS_ID     TypeExp = 4
	INC           TypeExp = 5
	DEC           TypeExp = 6
	PARAMETER     TypeExp = 7
	CALL_FUNC     TypeExp = 8
	RETURN        TypeExp = 9
	TERNARY       TypeExp = 10
)
