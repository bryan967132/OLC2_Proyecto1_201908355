package utils

var plus = [][]Type{
	//	     INT    FLOAT  STR
	/*INT*/ {INT, FLOAT, NIL},
	/*FLT*/ {FLOAT, FLOAT, NIL},
	/*STR*/ {NIL, NIL, STRING},
}

var minus = [][]Type{
	//	     INT    FLOAT
	/*INT*/ {INT, FLOAT},
	/*FLT*/ {FLOAT, FLOAT},
}

var mult = [][]Type{
	//	     INT    FLOAT
	/*INT*/ {INT, FLOAT},
	/*FLT*/ {FLOAT, FLOAT},
}

var div = [][]Type{
	//	     INT    FLOAT
	/*INT*/ {INT, FLOAT},
	/*FLT*/ {FLOAT, FLOAT},
}

var mod = [][]Type{
	//	     INT
	/*INT*/ {INT},
}
