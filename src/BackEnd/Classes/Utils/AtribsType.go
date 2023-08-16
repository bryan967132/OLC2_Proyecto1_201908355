package utils

type AttribsType struct {
	Line   int
	Column int
	Value  interface{}
}

func NewAttribsType(line, column int, value interface{}) *AttribsType {
	return &AttribsType{line, column, value}
}
