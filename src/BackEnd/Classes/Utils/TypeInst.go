package utils

type TypeInst int

const (
	PRINT         TypeInst = 0
	INIT_ID       TypeInst = 1
	ASIGN_ID      TypeInst = 2
	INIT_FUNCTION TypeInst = 3
	MAIN          TypeInst = 4
	BLOCK_INST    TypeInst = 5
	IF            TypeInst = 6
	LOOP_FOR      TypeInst = 7
	LOOP_WHILE    TypeInst = 8
	LOOP_DOWHILE  TypeInst = 9
	SWITCH        TypeInst = 10
	CASE          TypeInst = 11
	BREAK         TypeInst = 12
	CONTINUE      TypeInst = 13
	ADD           TypeInst = 14
	SUB           TypeInst = 15
)
