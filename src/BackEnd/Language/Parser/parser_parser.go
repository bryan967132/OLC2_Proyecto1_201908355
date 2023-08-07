// Code generated from Parser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Parser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ParserParser struct {
	*antlr.BaseParser
}

var ParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func parserParserInit() {
	staticData := &ParserParserStaticData
	staticData.SymbolicNames = []string{
		"", "RW_Int", "RW_Float", "RW_String", "RW_Bool", "RW_Character", "RW_var",
		"RW_let", "RW_if", "RW_else", "RW_for", "RW_while", "RW_guard", "RW_switch",
		"RW_case", "RW_default", "RW_break", "RW_continue", "RW_return", "RW_true",
		"RW_false", "RW_nil", "RW_func", "RW_inout", "RW_in", "RW_append", "RW_removeLast",
		"RW_remove", "RW_at", "RW_isEmpty", "RW_count", "RW_repeating", "RW_print",
		"TK_prompt", "TK_under", "TK_char", "TK_string", "TK_int", "TK_float",
		"TK_id", "TK_add", "TK_sub", "TK_plus", "TK_minus", "TK_mult", "TK_div",
		"TK_mod", "TK_equequ", "TK_notequ", "TK_lessequ", "TK_moreequ", "TK_equ",
		"TK_less", "TK_more", "TK_and", "TK_or", "TK_not", "TK_lpar", "TK_rpar",
		"TK_lbrc", "TK_rbrc", "TK_lbrk", "TK_rbrk", "TK_dot", "TK_comma", "TK_colon",
		"TK_semicolon", "TK_question", "TK_amp",
	}
	staticData.RuleNames = []string{
		"init", "instsglobal", "instglobal", "callfunc", "listargs", "decvar",
		"deccst", "declfunc", "listparams", "ifstruct", "switchstruct", "envs",
		"casesdefault", "cases", "case", "default", "loopfor", "range", "loopwhile",
		"guard", "reasign", "addsub", "decvector", "defvector", "listexp", "simplevec",
		"funcvector", "reasignvector", "print", "env", "instructions", "instruction",
		"type", "exp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 68, 442, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 3, 0, 70, 8, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 78, 8, 1, 1, 2, 1, 2, 3, 2, 82, 8, 2, 1, 3, 1,
		3, 1, 3, 3, 3, 87, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 93, 8, 4, 1, 4,
		3, 4, 96, 8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 104, 8, 4, 1,
		4, 3, 4, 107, 8, 4, 1, 4, 3, 4, 110, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 129, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 3, 6, 142, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 148, 8, 7, 1,
		7, 1, 7, 1, 7, 3, 7, 153, 8, 7, 1, 7, 1, 7, 1, 8, 3, 8, 158, 8, 8, 1, 8,
		1, 8, 1, 8, 3, 8, 163, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 170, 8,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 175, 8, 8, 1, 8, 3, 8, 178, 8, 8, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 196, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 211, 8, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 3, 13, 217, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3,
		16, 233, 8, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 271, 8, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 276, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24,
		3, 24, 283, 8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 319, 8, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 3, 28, 331, 8,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 3, 29, 337, 8, 29, 1, 29, 1, 29, 1, 30,
		1, 30, 1, 30, 1, 30, 3, 30, 345, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 366, 8, 31, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 3, 33, 414, 8, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 33, 1, 33, 5, 33, 437, 8, 33, 10, 33, 12, 33, 440, 9, 33,
		1, 33, 0, 1, 66, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
		64, 66, 0, 8, 2, 0, 34, 34, 39, 39, 1, 0, 40, 41, 1, 0, 1, 5, 1, 0, 44,
		46, 1, 0, 42, 43, 1, 0, 49, 50, 1, 0, 52, 53, 1, 0, 47, 48, 482, 0, 69,
		1, 0, 0, 0, 2, 77, 1, 0, 0, 0, 4, 81, 1, 0, 0, 0, 6, 83, 1, 0, 0, 0, 8,
		109, 1, 0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 141, 1, 0, 0, 0, 14, 143, 1,
		0, 0, 0, 16, 177, 1, 0, 0, 0, 18, 195, 1, 0, 0, 0, 20, 197, 1, 0, 0, 0,
		22, 201, 1, 0, 0, 0, 24, 210, 1, 0, 0, 0, 26, 216, 1, 0, 0, 0, 28, 218,
		1, 0, 0, 0, 30, 223, 1, 0, 0, 0, 32, 227, 1, 0, 0, 0, 34, 236, 1, 0, 0,
		0, 36, 242, 1, 0, 0, 0, 38, 246, 1, 0, 0, 0, 40, 251, 1, 0, 0, 0, 42, 255,
		1, 0, 0, 0, 44, 259, 1, 0, 0, 0, 46, 275, 1, 0, 0, 0, 48, 282, 1, 0, 0,
		0, 50, 284, 1, 0, 0, 0, 52, 318, 1, 0, 0, 0, 54, 320, 1, 0, 0, 0, 56, 327,
		1, 0, 0, 0, 58, 334, 1, 0, 0, 0, 60, 344, 1, 0, 0, 0, 62, 365, 1, 0, 0,
		0, 64, 367, 1, 0, 0, 0, 66, 413, 1, 0, 0, 0, 68, 70, 3, 2, 1, 0, 69, 68,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 72, 5, 0, 0, 1,
		72, 1, 1, 0, 0, 0, 73, 74, 3, 4, 2, 0, 74, 75, 3, 2, 1, 0, 75, 78, 1, 0,
		0, 0, 76, 78, 3, 4, 2, 0, 77, 73, 1, 0, 0, 0, 77, 76, 1, 0, 0, 0, 78, 3,
		1, 0, 0, 0, 79, 82, 3, 62, 31, 0, 80, 82, 3, 14, 7, 0, 81, 79, 1, 0, 0,
		0, 81, 80, 1, 0, 0, 0, 82, 5, 1, 0, 0, 0, 83, 84, 5, 39, 0, 0, 84, 86,
		5, 57, 0, 0, 85, 87, 3, 8, 4, 0, 86, 85, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 89, 5, 58, 0, 0, 89, 7, 1, 0, 0, 0, 90, 91, 5,
		39, 0, 0, 91, 93, 5, 65, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0,
		93, 95, 1, 0, 0, 0, 94, 96, 5, 68, 0, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1,
		0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 3, 66, 33, 0, 98, 99, 5, 64, 0, 0,
		99, 100, 3, 8, 4, 0, 100, 110, 1, 0, 0, 0, 101, 102, 5, 39, 0, 0, 102,
		104, 5, 65, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 106,
		1, 0, 0, 0, 105, 107, 5, 68, 0, 0, 106, 105, 1, 0, 0, 0, 106, 107, 1, 0,
		0, 0, 107, 108, 1, 0, 0, 0, 108, 110, 3, 66, 33, 0, 109, 92, 1, 0, 0, 0,
		109, 103, 1, 0, 0, 0, 110, 9, 1, 0, 0, 0, 111, 112, 5, 6, 0, 0, 112, 113,
		5, 39, 0, 0, 113, 114, 5, 65, 0, 0, 114, 115, 3, 64, 32, 0, 115, 116, 5,
		51, 0, 0, 116, 117, 3, 66, 33, 0, 117, 129, 1, 0, 0, 0, 118, 119, 5, 6,
		0, 0, 119, 120, 5, 39, 0, 0, 120, 121, 5, 65, 0, 0, 121, 122, 3, 64, 32,
		0, 122, 123, 5, 67, 0, 0, 123, 129, 1, 0, 0, 0, 124, 125, 5, 6, 0, 0, 125,
		126, 5, 39, 0, 0, 126, 127, 5, 51, 0, 0, 127, 129, 3, 66, 33, 0, 128, 111,
		1, 0, 0, 0, 128, 118, 1, 0, 0, 0, 128, 124, 1, 0, 0, 0, 129, 11, 1, 0,
		0, 0, 130, 131, 5, 7, 0, 0, 131, 132, 5, 39, 0, 0, 132, 133, 5, 65, 0,
		0, 133, 134, 3, 64, 32, 0, 134, 135, 5, 51, 0, 0, 135, 136, 3, 66, 33,
		0, 136, 142, 1, 0, 0, 0, 137, 138, 5, 7, 0, 0, 138, 139, 5, 39, 0, 0, 139,
		140, 5, 51, 0, 0, 140, 142, 3, 66, 33, 0, 141, 130, 1, 0, 0, 0, 141, 137,
		1, 0, 0, 0, 142, 13, 1, 0, 0, 0, 143, 144, 5, 22, 0, 0, 144, 145, 5, 39,
		0, 0, 145, 147, 5, 57, 0, 0, 146, 148, 3, 16, 8, 0, 147, 146, 1, 0, 0,
		0, 147, 148, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 152, 5, 58, 0, 0, 150,
		151, 5, 33, 0, 0, 151, 153, 3, 64, 32, 0, 152, 150, 1, 0, 0, 0, 152, 153,
		1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 155, 3, 58, 29, 0, 155, 15, 1, 0,
		0, 0, 156, 158, 7, 0, 0, 0, 157, 156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0,
		158, 159, 1, 0, 0, 0, 159, 160, 5, 39, 0, 0, 160, 162, 5, 65, 0, 0, 161,
		163, 5, 23, 0, 0, 162, 161, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 165, 3, 64, 32, 0, 165, 166, 5, 64, 0, 0, 166, 167, 3,
		16, 8, 0, 167, 178, 1, 0, 0, 0, 168, 170, 7, 0, 0, 0, 169, 168, 1, 0, 0,
		0, 169, 170, 1, 0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 5, 39, 0, 0, 172,
		174, 5, 65, 0, 0, 173, 175, 5, 23, 0, 0, 174, 173, 1, 0, 0, 0, 174, 175,
		1, 0, 0, 0, 175, 176, 1, 0, 0, 0, 176, 178, 3, 64, 32, 0, 177, 157, 1,
		0, 0, 0, 177, 169, 1, 0, 0, 0, 178, 17, 1, 0, 0, 0, 179, 180, 5, 8, 0,
		0, 180, 181, 3, 66, 33, 0, 181, 182, 3, 58, 29, 0, 182, 183, 5, 9, 0, 0,
		183, 184, 3, 18, 9, 0, 184, 196, 1, 0, 0, 0, 185, 186, 5, 8, 0, 0, 186,
		187, 3, 66, 33, 0, 187, 188, 3, 58, 29, 0, 188, 189, 5, 9, 0, 0, 189, 190,
		3, 58, 29, 0, 190, 196, 1, 0, 0, 0, 191, 192, 5, 8, 0, 0, 192, 193, 3,
		66, 33, 0, 193, 194, 3, 58, 29, 0, 194, 196, 1, 0, 0, 0, 195, 179, 1, 0,
		0, 0, 195, 185, 1, 0, 0, 0, 195, 191, 1, 0, 0, 0, 196, 19, 1, 0, 0, 0,
		197, 198, 5, 13, 0, 0, 198, 199, 3, 66, 33, 0, 199, 200, 3, 22, 11, 0,
		200, 21, 1, 0, 0, 0, 201, 202, 5, 59, 0, 0, 202, 203, 3, 24, 12, 0, 203,
		204, 5, 60, 0, 0, 204, 23, 1, 0, 0, 0, 205, 206, 3, 26, 13, 0, 206, 207,
		3, 30, 15, 0, 207, 211, 1, 0, 0, 0, 208, 211, 3, 26, 13, 0, 209, 211, 3,
		30, 15, 0, 210, 205, 1, 0, 0, 0, 210, 208, 1, 0, 0, 0, 210, 209, 1, 0,
		0, 0, 211, 25, 1, 0, 0, 0, 212, 213, 3, 28, 14, 0, 213, 214, 3, 26, 13,
		0, 214, 217, 1, 0, 0, 0, 215, 217, 3, 28, 14, 0, 216, 212, 1, 0, 0, 0,
		216, 215, 1, 0, 0, 0, 217, 27, 1, 0, 0, 0, 218, 219, 5, 14, 0, 0, 219,
		220, 3, 66, 33, 0, 220, 221, 5, 65, 0, 0, 221, 222, 3, 60, 30, 0, 222,
		29, 1, 0, 0, 0, 223, 224, 5, 15, 0, 0, 224, 225, 5, 65, 0, 0, 225, 226,
		3, 60, 30, 0, 226, 31, 1, 0, 0, 0, 227, 228, 5, 10, 0, 0, 228, 229, 5,
		39, 0, 0, 229, 232, 5, 24, 0, 0, 230, 233, 3, 66, 33, 0, 231, 233, 3, 34,
		17, 0, 232, 230, 1, 0, 0, 0, 232, 231, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0,
		234, 235, 3, 58, 29, 0, 235, 33, 1, 0, 0, 0, 236, 237, 3, 66, 33, 0, 237,
		238, 5, 63, 0, 0, 238, 239, 5, 63, 0, 0, 239, 240, 5, 63, 0, 0, 240, 241,
		3, 66, 33, 0, 241, 35, 1, 0, 0, 0, 242, 243, 5, 11, 0, 0, 243, 244, 3,
		66, 33, 0, 244, 245, 3, 58, 29, 0, 245, 37, 1, 0, 0, 0, 246, 247, 5, 12,
		0, 0, 247, 248, 3, 66, 33, 0, 248, 249, 5, 9, 0, 0, 249, 250, 3, 58, 29,
		0, 250, 39, 1, 0, 0, 0, 251, 252, 5, 39, 0, 0, 252, 253, 5, 51, 0, 0, 253,
		254, 3, 66, 33, 0, 254, 41, 1, 0, 0, 0, 255, 256, 5, 39, 0, 0, 256, 257,
		7, 1, 0, 0, 257, 258, 3, 66, 33, 0, 258, 43, 1, 0, 0, 0, 259, 260, 5, 6,
		0, 0, 260, 261, 5, 39, 0, 0, 261, 262, 5, 65, 0, 0, 262, 263, 5, 61, 0,
		0, 263, 264, 3, 64, 32, 0, 264, 265, 5, 62, 0, 0, 265, 266, 5, 51, 0, 0,
		266, 267, 3, 46, 23, 0, 267, 45, 1, 0, 0, 0, 268, 270, 5, 61, 0, 0, 269,
		271, 3, 48, 24, 0, 270, 269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272,
		1, 0, 0, 0, 272, 276, 5, 62, 0, 0, 273, 276, 3, 50, 25, 0, 274, 276, 5,
		39, 0, 0, 275, 268, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 274, 1, 0, 0,
		0, 276, 47, 1, 0, 0, 0, 277, 278, 3, 66, 33, 0, 278, 279, 5, 64, 0, 0,
		279, 280, 3, 48, 24, 0, 280, 283, 1, 0, 0, 0, 281, 283, 3, 66, 33, 0, 282,
		277, 1, 0, 0, 0, 282, 281, 1, 0, 0, 0, 283, 49, 1, 0, 0, 0, 284, 285, 5,
		61, 0, 0, 285, 286, 3, 64, 32, 0, 286, 287, 5, 62, 0, 0, 287, 288, 5, 57,
		0, 0, 288, 289, 5, 31, 0, 0, 289, 290, 5, 65, 0, 0, 290, 291, 3, 66, 33,
		0, 291, 292, 5, 64, 0, 0, 292, 293, 5, 30, 0, 0, 293, 294, 5, 65, 0, 0,
		294, 295, 3, 66, 33, 0, 295, 296, 5, 58, 0, 0, 296, 51, 1, 0, 0, 0, 297,
		298, 5, 39, 0, 0, 298, 299, 5, 63, 0, 0, 299, 300, 5, 25, 0, 0, 300, 301,
		5, 57, 0, 0, 301, 302, 3, 66, 33, 0, 302, 303, 5, 58, 0, 0, 303, 319, 1,
		0, 0, 0, 304, 305, 5, 39, 0, 0, 305, 306, 5, 63, 0, 0, 306, 307, 5, 26,
		0, 0, 307, 308, 5, 57, 0, 0, 308, 319, 5, 58, 0, 0, 309, 310, 5, 39, 0,
		0, 310, 311, 5, 63, 0, 0, 311, 312, 5, 27, 0, 0, 312, 313, 5, 57, 0, 0,
		313, 314, 5, 28, 0, 0, 314, 315, 5, 65, 0, 0, 315, 316, 3, 66, 33, 0, 316,
		317, 5, 58, 0, 0, 317, 319, 1, 0, 0, 0, 318, 297, 1, 0, 0, 0, 318, 304,
		1, 0, 0, 0, 318, 309, 1, 0, 0, 0, 319, 53, 1, 0, 0, 0, 320, 321, 5, 39,
		0, 0, 321, 322, 5, 61, 0, 0, 322, 323, 3, 66, 33, 0, 323, 324, 5, 62, 0,
		0, 324, 325, 5, 51, 0, 0, 325, 326, 3, 66, 33, 0, 326, 55, 1, 0, 0, 0,
		327, 328, 5, 32, 0, 0, 328, 330, 5, 57, 0, 0, 329, 331, 3, 66, 33, 0, 330,
		329, 1, 0, 0, 0, 330, 331, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 333,
		5, 58, 0, 0, 333, 57, 1, 0, 0, 0, 334, 336, 5, 59, 0, 0, 335, 337, 3, 60,
		30, 0, 336, 335, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 338, 1, 0, 0, 0,
		338, 339, 5, 60, 0, 0, 339, 59, 1, 0, 0, 0, 340, 341, 3, 62, 31, 0, 341,
		342, 3, 60, 30, 0, 342, 345, 1, 0, 0, 0, 343, 345, 3, 62, 31, 0, 344, 340,
		1, 0, 0, 0, 344, 343, 1, 0, 0, 0, 345, 61, 1, 0, 0, 0, 346, 366, 3, 10,
		5, 0, 347, 366, 3, 12, 6, 0, 348, 366, 3, 18, 9, 0, 349, 366, 3, 20, 10,
		0, 350, 366, 3, 32, 16, 0, 351, 366, 3, 36, 18, 0, 352, 366, 3, 38, 19,
		0, 353, 366, 3, 40, 20, 0, 354, 366, 3, 42, 21, 0, 355, 366, 3, 44, 22,
		0, 356, 366, 3, 52, 26, 0, 357, 366, 3, 54, 27, 0, 358, 366, 3, 6, 3, 0,
		359, 366, 3, 56, 28, 0, 360, 361, 5, 18, 0, 0, 361, 366, 3, 66, 33, 0,
		362, 366, 5, 18, 0, 0, 363, 366, 5, 17, 0, 0, 364, 366, 5, 16, 0, 0, 365,
		346, 1, 0, 0, 0, 365, 347, 1, 0, 0, 0, 365, 348, 1, 0, 0, 0, 365, 349,
		1, 0, 0, 0, 365, 350, 1, 0, 0, 0, 365, 351, 1, 0, 0, 0, 365, 352, 1, 0,
		0, 0, 365, 353, 1, 0, 0, 0, 365, 354, 1, 0, 0, 0, 365, 355, 1, 0, 0, 0,
		365, 356, 1, 0, 0, 0, 365, 357, 1, 0, 0, 0, 365, 358, 1, 0, 0, 0, 365,
		359, 1, 0, 0, 0, 365, 360, 1, 0, 0, 0, 365, 362, 1, 0, 0, 0, 365, 363,
		1, 0, 0, 0, 365, 364, 1, 0, 0, 0, 366, 63, 1, 0, 0, 0, 367, 368, 7, 2,
		0, 0, 368, 65, 1, 0, 0, 0, 369, 370, 6, 33, -1, 0, 370, 371, 5, 43, 0,
		0, 371, 414, 3, 66, 33, 25, 372, 373, 5, 56, 0, 0, 373, 414, 3, 66, 33,
		19, 374, 375, 5, 1, 0, 0, 375, 376, 5, 57, 0, 0, 376, 377, 3, 66, 33, 0,
		377, 378, 5, 58, 0, 0, 378, 414, 1, 0, 0, 0, 379, 380, 5, 2, 0, 0, 380,
		381, 5, 57, 0, 0, 381, 382, 3, 66, 33, 0, 382, 383, 5, 58, 0, 0, 383, 414,
		1, 0, 0, 0, 384, 385, 5, 3, 0, 0, 385, 386, 5, 57, 0, 0, 386, 387, 3, 66,
		33, 0, 387, 388, 5, 58, 0, 0, 388, 414, 1, 0, 0, 0, 389, 390, 5, 39, 0,
		0, 390, 391, 5, 61, 0, 0, 391, 392, 3, 66, 33, 0, 392, 393, 5, 62, 0, 0,
		393, 414, 1, 0, 0, 0, 394, 395, 5, 39, 0, 0, 395, 396, 5, 63, 0, 0, 396,
		414, 5, 29, 0, 0, 397, 398, 5, 39, 0, 0, 398, 399, 5, 63, 0, 0, 399, 414,
		5, 30, 0, 0, 400, 414, 3, 6, 3, 0, 401, 414, 5, 21, 0, 0, 402, 414, 5,
		39, 0, 0, 403, 414, 5, 36, 0, 0, 404, 414, 5, 35, 0, 0, 405, 414, 5, 37,
		0, 0, 406, 414, 5, 38, 0, 0, 407, 414, 5, 19, 0, 0, 408, 414, 5, 20, 0,
		0, 409, 410, 5, 57, 0, 0, 410, 411, 3, 66, 33, 0, 411, 412, 5, 58, 0, 0,
		412, 414, 1, 0, 0, 0, 413, 369, 1, 0, 0, 0, 413, 372, 1, 0, 0, 0, 413,
		374, 1, 0, 0, 0, 413, 379, 1, 0, 0, 0, 413, 384, 1, 0, 0, 0, 413, 389,
		1, 0, 0, 0, 413, 394, 1, 0, 0, 0, 413, 397, 1, 0, 0, 0, 413, 400, 1, 0,
		0, 0, 413, 401, 1, 0, 0, 0, 413, 402, 1, 0, 0, 0, 413, 403, 1, 0, 0, 0,
		413, 404, 1, 0, 0, 0, 413, 405, 1, 0, 0, 0, 413, 406, 1, 0, 0, 0, 413,
		407, 1, 0, 0, 0, 413, 408, 1, 0, 0, 0, 413, 409, 1, 0, 0, 0, 414, 438,
		1, 0, 0, 0, 415, 416, 10, 24, 0, 0, 416, 417, 7, 3, 0, 0, 417, 437, 3,
		66, 33, 25, 418, 419, 10, 23, 0, 0, 419, 420, 7, 4, 0, 0, 420, 437, 3,
		66, 33, 24, 421, 422, 10, 22, 0, 0, 422, 423, 7, 5, 0, 0, 423, 437, 3,
		66, 33, 23, 424, 425, 10, 21, 0, 0, 425, 426, 7, 6, 0, 0, 426, 437, 3,
		66, 33, 22, 427, 428, 10, 20, 0, 0, 428, 429, 7, 7, 0, 0, 429, 437, 3,
		66, 33, 21, 430, 431, 10, 18, 0, 0, 431, 432, 5, 54, 0, 0, 432, 437, 3,
		66, 33, 19, 433, 434, 10, 17, 0, 0, 434, 435, 5, 55, 0, 0, 435, 437, 3,
		66, 33, 18, 436, 415, 1, 0, 0, 0, 436, 418, 1, 0, 0, 0, 436, 421, 1, 0,
		0, 0, 436, 424, 1, 0, 0, 0, 436, 427, 1, 0, 0, 0, 436, 430, 1, 0, 0, 0,
		436, 433, 1, 0, 0, 0, 437, 440, 1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 438,
		439, 1, 0, 0, 0, 439, 67, 1, 0, 0, 0, 440, 438, 1, 0, 0, 0, 33, 69, 77,
		81, 86, 92, 95, 103, 106, 109, 128, 141, 147, 152, 157, 162, 169, 174,
		177, 195, 210, 216, 232, 270, 275, 282, 318, 330, 336, 344, 365, 413, 436,
		438,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ParserParserInit initializes any static state used to implement ParserParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParserParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParserParserInit() {
	staticData := &ParserParserStaticData
	staticData.once.Do(parserParserInit)
}

// NewParserParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParserParser(input antlr.TokenStream) *ParserParser {
	ParserParserInit()
	this := new(ParserParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Parser.g4"

	return this
}

// ParserParser tokens.
const (
	ParserParserEOF           = antlr.TokenEOF
	ParserParserRW_Int        = 1
	ParserParserRW_Float      = 2
	ParserParserRW_String     = 3
	ParserParserRW_Bool       = 4
	ParserParserRW_Character  = 5
	ParserParserRW_var        = 6
	ParserParserRW_let        = 7
	ParserParserRW_if         = 8
	ParserParserRW_else       = 9
	ParserParserRW_for        = 10
	ParserParserRW_while      = 11
	ParserParserRW_guard      = 12
	ParserParserRW_switch     = 13
	ParserParserRW_case       = 14
	ParserParserRW_default    = 15
	ParserParserRW_break      = 16
	ParserParserRW_continue   = 17
	ParserParserRW_return     = 18
	ParserParserRW_true       = 19
	ParserParserRW_false      = 20
	ParserParserRW_nil        = 21
	ParserParserRW_func       = 22
	ParserParserRW_inout      = 23
	ParserParserRW_in         = 24
	ParserParserRW_append     = 25
	ParserParserRW_removeLast = 26
	ParserParserRW_remove     = 27
	ParserParserRW_at         = 28
	ParserParserRW_isEmpty    = 29
	ParserParserRW_count      = 30
	ParserParserRW_repeating  = 31
	ParserParserRW_print      = 32
	ParserParserTK_prompt     = 33
	ParserParserTK_under      = 34
	ParserParserTK_char       = 35
	ParserParserTK_string     = 36
	ParserParserTK_int        = 37
	ParserParserTK_float      = 38
	ParserParserTK_id         = 39
	ParserParserTK_add        = 40
	ParserParserTK_sub        = 41
	ParserParserTK_plus       = 42
	ParserParserTK_minus      = 43
	ParserParserTK_mult       = 44
	ParserParserTK_div        = 45
	ParserParserTK_mod        = 46
	ParserParserTK_equequ     = 47
	ParserParserTK_notequ     = 48
	ParserParserTK_lessequ    = 49
	ParserParserTK_moreequ    = 50
	ParserParserTK_equ        = 51
	ParserParserTK_less       = 52
	ParserParserTK_more       = 53
	ParserParserTK_and        = 54
	ParserParserTK_or         = 55
	ParserParserTK_not        = 56
	ParserParserTK_lpar       = 57
	ParserParserTK_rpar       = 58
	ParserParserTK_lbrc       = 59
	ParserParserTK_rbrc       = 60
	ParserParserTK_lbrk       = 61
	ParserParserTK_rbrk       = 62
	ParserParserTK_dot        = 63
	ParserParserTK_comma      = 64
	ParserParserTK_colon      = 65
	ParserParserTK_semicolon  = 66
	ParserParserTK_question   = 67
	ParserParserTK_amp        = 68
)

// ParserParser rules.
const (
	ParserParserRULE_init          = 0
	ParserParserRULE_instsglobal   = 1
	ParserParserRULE_instglobal    = 2
	ParserParserRULE_callfunc      = 3
	ParserParserRULE_listargs      = 4
	ParserParserRULE_decvar        = 5
	ParserParserRULE_deccst        = 6
	ParserParserRULE_declfunc      = 7
	ParserParserRULE_listparams    = 8
	ParserParserRULE_ifstruct      = 9
	ParserParserRULE_switchstruct  = 10
	ParserParserRULE_envs          = 11
	ParserParserRULE_casesdefault  = 12
	ParserParserRULE_cases         = 13
	ParserParserRULE_case          = 14
	ParserParserRULE_default       = 15
	ParserParserRULE_loopfor       = 16
	ParserParserRULE_range         = 17
	ParserParserRULE_loopwhile     = 18
	ParserParserRULE_guard         = 19
	ParserParserRULE_reasign       = 20
	ParserParserRULE_addsub        = 21
	ParserParserRULE_decvector     = 22
	ParserParserRULE_defvector     = 23
	ParserParserRULE_listexp       = 24
	ParserParserRULE_simplevec     = 25
	ParserParserRULE_funcvector    = 26
	ParserParserRULE_reasignvector = 27
	ParserParserRULE_print         = 28
	ParserParserRULE_env           = 29
	ParserParserRULE_instructions  = 30
	ParserParserRULE_instruction   = 31
	ParserParserRULE_type          = 32
	ParserParserRULE_exp           = 33
)

// IInitContext is an interface to support dynamic dispatch.
type IInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	Instsglobal() IInstsglobalContext

	// IsInitContext differentiates from other interfaces.
	IsInitContext()
}

type InitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitContext() *InitContext {
	var p = new(InitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
	return p
}

func InitEmptyInitContext(p *InitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_init
}

func (*InitContext) IsInitContext() {}

func NewInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitContext {
	var p = new(InitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_init

	return p
}

func (s *InitContext) GetParser() antlr.Parser { return s.parser }

func (s *InitContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParserParserEOF, 0)
}

func (s *InitContext) Instsglobal() IInstsglobalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstsglobalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstsglobalContext)
}

func (s *InitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInit(s)
	}
}

func (s *InitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInit(s)
	}
}

func (p *ParserParser) Init() (localctx IInitContext) {
	localctx = NewInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParserParserRULE_init)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&554055450048) != 0 {
		{
			p.SetState(68)
			p.Instsglobal()
		}

	}
	{
		p.SetState(71)
		p.Match(ParserParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstsglobalContext is an interface to support dynamic dispatch.
type IInstsglobalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instglobal() IInstglobalContext
	Instsglobal() IInstsglobalContext

	// IsInstsglobalContext differentiates from other interfaces.
	IsInstsglobalContext()
}

type InstsglobalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstsglobalContext() *InstsglobalContext {
	var p = new(InstsglobalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instsglobal
	return p
}

func InitEmptyInstsglobalContext(p *InstsglobalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instsglobal
}

func (*InstsglobalContext) IsInstsglobalContext() {}

func NewInstsglobalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstsglobalContext {
	var p = new(InstsglobalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_instsglobal

	return p
}

func (s *InstsglobalContext) GetParser() antlr.Parser { return s.parser }

func (s *InstsglobalContext) Instglobal() IInstglobalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstglobalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstglobalContext)
}

func (s *InstsglobalContext) Instsglobal() IInstsglobalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstsglobalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstsglobalContext)
}

func (s *InstsglobalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstsglobalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstsglobalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInstsglobal(s)
	}
}

func (s *InstsglobalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInstsglobal(s)
	}
}

func (p *ParserParser) Instsglobal() (localctx IInstsglobalContext) {
	localctx = NewInstsglobalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParserParserRULE_instsglobal)
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(73)
			p.Instglobal()
		}
		{
			p.SetState(74)
			p.Instsglobal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(76)
			p.Instglobal()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstglobalContext is an interface to support dynamic dispatch.
type IInstglobalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instruction() IInstructionContext
	Declfunc() IDeclfuncContext

	// IsInstglobalContext differentiates from other interfaces.
	IsInstglobalContext()
}

type InstglobalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstglobalContext() *InstglobalContext {
	var p = new(InstglobalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instglobal
	return p
}

func InitEmptyInstglobalContext(p *InstglobalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instglobal
}

func (*InstglobalContext) IsInstglobalContext() {}

func NewInstglobalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstglobalContext {
	var p = new(InstglobalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_instglobal

	return p
}

func (s *InstglobalContext) GetParser() antlr.Parser { return s.parser }

func (s *InstglobalContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstglobalContext) Declfunc() IDeclfuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclfuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclfuncContext)
}

func (s *InstglobalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstglobalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstglobalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInstglobal(s)
	}
}

func (s *InstglobalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInstglobal(s)
	}
}

func (p *ParserParser) Instglobal() (localctx IInstglobalContext) {
	localctx = NewInstglobalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParserParserRULE_instglobal)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_var, ParserParserRW_let, ParserParserRW_if, ParserParserRW_for, ParserParserRW_while, ParserParserRW_guard, ParserParserRW_switch, ParserParserRW_break, ParserParserRW_continue, ParserParserRW_return, ParserParserRW_print, ParserParserTK_id:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.Instruction()
		}

	case ParserParserRW_func:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Declfunc()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallfuncContext is an interface to support dynamic dispatch.
type ICallfuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	TK_rpar() antlr.TerminalNode
	Listargs() IListargsContext

	// IsCallfuncContext differentiates from other interfaces.
	IsCallfuncContext()
}

type CallfuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallfuncContext() *CallfuncContext {
	var p = new(CallfuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_callfunc
	return p
}

func InitEmptyCallfuncContext(p *CallfuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_callfunc
}

func (*CallfuncContext) IsCallfuncContext() {}

func NewCallfuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallfuncContext {
	var p = new(CallfuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_callfunc

	return p
}

func (s *CallfuncContext) GetParser() antlr.Parser { return s.parser }

func (s *CallfuncContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *CallfuncContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *CallfuncContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *CallfuncContext) Listargs() IListargsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListargsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListargsContext)
}

func (s *CallfuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallfuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallfuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCallfunc(s)
	}
}

func (s *CallfuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCallfunc(s)
	}
}

func (p *ParserParser) Callfunc() (localctx ICallfuncContext) {
	localctx = NewCallfuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParserParserRULE_callfunc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(84)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216182643362365454) != 0) || _la == ParserParserTK_amp {
		{
			p.SetState(85)
			p.Listargs()
		}

	}
	{
		p.SetState(88)
		p.Match(ParserParserTK_rpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListargsContext is an interface to support dynamic dispatch.
type IListargsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	TK_comma() antlr.TerminalNode
	Listargs() IListargsContext
	TK_id() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	TK_amp() antlr.TerminalNode

	// IsListargsContext differentiates from other interfaces.
	IsListargsContext()
}

type ListargsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListargsContext() *ListargsContext {
	var p = new(ListargsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listargs
	return p
}

func InitEmptyListargsContext(p *ListargsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listargs
}

func (*ListargsContext) IsListargsContext() {}

func NewListargsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListargsContext {
	var p = new(ListargsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listargs

	return p
}

func (s *ListargsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListargsContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ListargsContext) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *ListargsContext) Listargs() IListargsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListargsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListargsContext)
}

func (s *ListargsContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *ListargsContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *ListargsContext) TK_amp() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_amp, 0)
}

func (s *ListargsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListargsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListargsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListargs(s)
	}
}

func (s *ListargsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListargs(s)
	}
}

func (p *ParserParser) Listargs() (localctx IListargsContext) {
	localctx = NewListargsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParserParserRULE_listargs)
	var _la int

	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(92)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(90)
				p.Match(ParserParserTK_id)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(91)
				p.Match(ParserParserTK_colon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_amp {
			{
				p.SetState(94)
				p.Match(ParserParserTK_amp)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(97)
			p.exp(0)
		}
		{
			p.SetState(98)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Listargs()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(103)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(101)
				p.Match(ParserParserTK_id)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(102)
				p.Match(ParserParserTK_colon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_amp {
			{
				p.SetState(105)
				p.Match(ParserParserTK_amp)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(108)
			p.exp(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecvarContext is an interface to support dynamic dispatch.
type IDecvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_var() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Type_() ITypeContext
	TK_equ() antlr.TerminalNode
	Exp() IExpContext
	TK_question() antlr.TerminalNode

	// IsDecvarContext differentiates from other interfaces.
	IsDecvarContext()
}

type DecvarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecvarContext() *DecvarContext {
	var p = new(DecvarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decvar
	return p
}

func InitEmptyDecvarContext(p *DecvarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decvar
}

func (*DecvarContext) IsDecvarContext() {}

func NewDecvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecvarContext {
	var p = new(DecvarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_decvar

	return p
}

func (s *DecvarContext) GetParser() antlr.Parser { return s.parser }

func (s *DecvarContext) RW_var() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_var, 0)
}

func (s *DecvarContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DecvarContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *DecvarContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DecvarContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *DecvarContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *DecvarContext) TK_question() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_question, 0)
}

func (s *DecvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDecvar(s)
	}
}

func (s *DecvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDecvar(s)
	}
}

func (p *ParserParser) Decvar() (localctx IDecvarContext) {
	localctx = NewDecvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParserParserRULE_decvar)
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.Match(ParserParserRW_var)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(112)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Type_()
		}
		{
			p.SetState(115)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(ParserParserRW_var)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Type_()
		}
		{
			p.SetState(122)
			p.Match(ParserParserTK_question)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(ParserParserRW_var)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.exp(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeccstContext is an interface to support dynamic dispatch.
type IDeccstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_let() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Type_() ITypeContext
	TK_equ() antlr.TerminalNode
	Exp() IExpContext

	// IsDeccstContext differentiates from other interfaces.
	IsDeccstContext()
}

type DeccstContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeccstContext() *DeccstContext {
	var p = new(DeccstContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_deccst
	return p
}

func InitEmptyDeccstContext(p *DeccstContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_deccst
}

func (*DeccstContext) IsDeccstContext() {}

func NewDeccstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeccstContext {
	var p = new(DeccstContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_deccst

	return p
}

func (s *DeccstContext) GetParser() antlr.Parser { return s.parser }

func (s *DeccstContext) RW_let() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_let, 0)
}

func (s *DeccstContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DeccstContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *DeccstContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeccstContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *DeccstContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *DeccstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeccstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeccstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDeccst(s)
	}
}

func (s *DeccstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDeccst(s)
	}
}

func (p *ParserParser) Deccst() (localctx IDeccstContext) {
	localctx = NewDeccstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParserParserRULE_deccst)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(ParserParserRW_let)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(133)
			p.Type_()
		}
		{
			p.SetState(134)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(ParserParserRW_let)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(138)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(140)
			p.exp(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDeclfuncContext is an interface to support dynamic dispatch.
type IDeclfuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_func() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	TK_rpar() antlr.TerminalNode
	Env() IEnvContext
	Listparams() IListparamsContext
	TK_prompt() antlr.TerminalNode
	Type_() ITypeContext

	// IsDeclfuncContext differentiates from other interfaces.
	IsDeclfuncContext()
}

type DeclfuncContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclfuncContext() *DeclfuncContext {
	var p = new(DeclfuncContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_declfunc
	return p
}

func InitEmptyDeclfuncContext(p *DeclfuncContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_declfunc
}

func (*DeclfuncContext) IsDeclfuncContext() {}

func NewDeclfuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclfuncContext {
	var p = new(DeclfuncContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_declfunc

	return p
}

func (s *DeclfuncContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclfuncContext) RW_func() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_func, 0)
}

func (s *DeclfuncContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DeclfuncContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *DeclfuncContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *DeclfuncContext) Env() IEnvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvContext)
}

func (s *DeclfuncContext) Listparams() IListparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListparamsContext)
}

func (s *DeclfuncContext) TK_prompt() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_prompt, 0)
}

func (s *DeclfuncContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DeclfuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclfuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclfuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDeclfunc(s)
	}
}

func (s *DeclfuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDeclfunc(s)
	}
}

func (p *ParserParser) Declfunc() (localctx IDeclfuncContext) {
	localctx = NewDeclfuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParserParserRULE_declfunc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(ParserParserRW_func)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserParserTK_under || _la == ParserParserTK_id {
		{
			p.SetState(146)
			p.Listparams()
		}

	}
	{
		p.SetState(149)
		p.Match(ParserParserTK_rpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserParserTK_prompt {
		{
			p.SetState(150)
			p.Match(ParserParserTK_prompt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Type_()
		}

	}
	{
		p.SetState(154)
		p.Env()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListparamsContext is an interface to support dynamic dispatch.
type IListparamsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTK_id() []antlr.TerminalNode
	TK_id(i int) antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Type_() ITypeContext
	TK_comma() antlr.TerminalNode
	Listparams() IListparamsContext
	RW_inout() antlr.TerminalNode
	TK_under() antlr.TerminalNode

	// IsListparamsContext differentiates from other interfaces.
	IsListparamsContext()
}

type ListparamsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListparamsContext() *ListparamsContext {
	var p = new(ListparamsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listparams
	return p
}

func InitEmptyListparamsContext(p *ListparamsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listparams
}

func (*ListparamsContext) IsListparamsContext() {}

func NewListparamsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListparamsContext {
	var p = new(ListparamsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listparams

	return p
}

func (s *ListparamsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListparamsContext) AllTK_id() []antlr.TerminalNode {
	return s.GetTokens(ParserParserTK_id)
}

func (s *ListparamsContext) TK_id(i int) antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, i)
}

func (s *ListparamsContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *ListparamsContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *ListparamsContext) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *ListparamsContext) Listparams() IListparamsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListparamsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListparamsContext)
}

func (s *ListparamsContext) RW_inout() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_inout, 0)
}

func (s *ListparamsContext) TK_under() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_under, 0)
}

func (s *ListparamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListparamsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListparamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListparams(s)
	}
}

func (s *ListparamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListparams(s)
	}
}

func (p *ParserParser) Listparams() (localctx IListparamsContext) {
	localctx = NewListparamsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParserParserRULE_listparams)
	var _la int

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(157)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(156)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ParserParserTK_under || _la == ParserParserTK_id) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(159)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_inout {
			{
				p.SetState(161)
				p.Match(ParserParserRW_inout)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(164)
			p.Type_()
		}
		{
			p.SetState(165)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(166)
			p.Listparams()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(168)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ParserParserTK_under || _la == ParserParserTK_id) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		{
			p.SetState(171)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_inout {
			{
				p.SetState(173)
				p.Match(ParserParserRW_inout)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(176)
			p.Type_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfstructContext is an interface to support dynamic dispatch.
type IIfstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_if() antlr.TerminalNode
	Exp() IExpContext
	AllEnv() []IEnvContext
	Env(i int) IEnvContext
	RW_else() antlr.TerminalNode
	Ifstruct() IIfstructContext

	// IsIfstructContext differentiates from other interfaces.
	IsIfstructContext()
}

type IfstructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstructContext() *IfstructContext {
	var p = new(IfstructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_ifstruct
	return p
}

func InitEmptyIfstructContext(p *IfstructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_ifstruct
}

func (*IfstructContext) IsIfstructContext() {}

func NewIfstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstructContext {
	var p = new(IfstructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_ifstruct

	return p
}

func (s *IfstructContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstructContext) RW_if() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_if, 0)
}

func (s *IfstructContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *IfstructContext) AllEnv() []IEnvContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnvContext); ok {
			len++
		}
	}

	tst := make([]IEnvContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnvContext); ok {
			tst[i] = t.(IEnvContext)
			i++
		}
	}

	return tst
}

func (s *IfstructContext) Env(i int) IEnvContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvContext)
}

func (s *IfstructContext) RW_else() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_else, 0)
}

func (s *IfstructContext) Ifstruct() IIfstructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstructContext)
}

func (s *IfstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterIfstruct(s)
	}
}

func (s *IfstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitIfstruct(s)
	}
}

func (p *ParserParser) Ifstruct() (localctx IIfstructContext) {
	localctx = NewIfstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParserParserRULE_ifstruct)
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Match(ParserParserRW_if)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(180)
			p.exp(0)
		}
		{
			p.SetState(181)
			p.Env()
		}
		{
			p.SetState(182)
			p.Match(ParserParserRW_else)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Ifstruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(ParserParserRW_if)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(186)
			p.exp(0)
		}
		{
			p.SetState(187)
			p.Env()
		}
		{
			p.SetState(188)
			p.Match(ParserParserRW_else)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Env()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(191)
			p.Match(ParserParserRW_if)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.exp(0)
		}
		{
			p.SetState(193)
			p.Env()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchstructContext is an interface to support dynamic dispatch.
type ISwitchstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_switch() antlr.TerminalNode
	Exp() IExpContext
	Envs() IEnvsContext

	// IsSwitchstructContext differentiates from other interfaces.
	IsSwitchstructContext()
}

type SwitchstructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchstructContext() *SwitchstructContext {
	var p = new(SwitchstructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_switchstruct
	return p
}

func InitEmptySwitchstructContext(p *SwitchstructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_switchstruct
}

func (*SwitchstructContext) IsSwitchstructContext() {}

func NewSwitchstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchstructContext {
	var p = new(SwitchstructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_switchstruct

	return p
}

func (s *SwitchstructContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchstructContext) RW_switch() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_switch, 0)
}

func (s *SwitchstructContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *SwitchstructContext) Envs() IEnvsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvsContext)
}

func (s *SwitchstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterSwitchstruct(s)
	}
}

func (s *SwitchstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitSwitchstruct(s)
	}
}

func (p *ParserParser) Switchstruct() (localctx ISwitchstructContext) {
	localctx = NewSwitchstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParserParserRULE_switchstruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(ParserParserRW_switch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.exp(0)
	}
	{
		p.SetState(199)
		p.Envs()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnvsContext is an interface to support dynamic dispatch.
type IEnvsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrc() antlr.TerminalNode
	Casesdefault() ICasesdefaultContext
	TK_rbrc() antlr.TerminalNode

	// IsEnvsContext differentiates from other interfaces.
	IsEnvsContext()
}

type EnvsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvsContext() *EnvsContext {
	var p = new(EnvsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_envs
	return p
}

func InitEmptyEnvsContext(p *EnvsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_envs
}

func (*EnvsContext) IsEnvsContext() {}

func NewEnvsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvsContext {
	var p = new(EnvsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_envs

	return p
}

func (s *EnvsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvsContext) TK_lbrc() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrc, 0)
}

func (s *EnvsContext) Casesdefault() ICasesdefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesdefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesdefaultContext)
}

func (s *EnvsContext) TK_rbrc() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrc, 0)
}

func (s *EnvsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterEnvs(s)
	}
}

func (s *EnvsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitEnvs(s)
	}
}

func (p *ParserParser) Envs() (localctx IEnvsContext) {
	localctx = NewEnvsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParserParserRULE_envs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.Match(ParserParserTK_lbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(202)
		p.Casesdefault()
	}
	{
		p.SetState(203)
		p.Match(ParserParserTK_rbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasesdefaultContext is an interface to support dynamic dispatch.
type ICasesdefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Cases() ICasesContext
	Default_() IDefaultContext

	// IsCasesdefaultContext differentiates from other interfaces.
	IsCasesdefaultContext()
}

type CasesdefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesdefaultContext() *CasesdefaultContext {
	var p = new(CasesdefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_casesdefault
	return p
}

func InitEmptyCasesdefaultContext(p *CasesdefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_casesdefault
}

func (*CasesdefaultContext) IsCasesdefaultContext() {}

func NewCasesdefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesdefaultContext {
	var p = new(CasesdefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_casesdefault

	return p
}

func (s *CasesdefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesdefaultContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *CasesdefaultContext) Default_() IDefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultContext)
}

func (s *CasesdefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesdefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesdefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCasesdefault(s)
	}
}

func (s *CasesdefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCasesdefault(s)
	}
}

func (p *ParserParser) Casesdefault() (localctx ICasesdefaultContext) {
	localctx = NewCasesdefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParserParserRULE_casesdefault)
	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Cases()
		}
		{
			p.SetState(206)
			p.Default_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Cases()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Default_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Case_() ICaseContext
	Cases() ICasesContext

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_cases
	return p
}

func InitEmptyCasesContext(p *CasesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_cases
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) Case_() ICaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseContext)
}

func (s *CasesContext) Cases() ICasesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICasesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCases(s)
	}
}

func (p *ParserParser) Cases() (localctx ICasesContext) {
	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParserParserRULE_cases)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Case_()
		}
		{
			p.SetState(213)
			p.Cases()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Case_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseContext is an interface to support dynamic dispatch.
type ICaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_case() antlr.TerminalNode
	Exp() IExpContext
	TK_colon() antlr.TerminalNode
	Instructions() IInstructionsContext

	// IsCaseContext differentiates from other interfaces.
	IsCaseContext()
}

type CaseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseContext() *CaseContext {
	var p = new(CaseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_case
	return p
}

func InitEmptyCaseContext(p *CaseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_case
}

func (*CaseContext) IsCaseContext() {}

func NewCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseContext {
	var p = new(CaseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_case

	return p
}

func (s *CaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseContext) RW_case() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_case, 0)
}

func (s *CaseContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *CaseContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *CaseContext) Instructions() IInstructionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitCase(s)
	}
}

func (p *ParserParser) Case_() (localctx ICaseContext) {
	localctx = NewCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ParserParserRULE_case)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(ParserParserRW_case)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.exp(0)
	}
	{
		p.SetState(220)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Instructions()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultContext is an interface to support dynamic dispatch.
type IDefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_default() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Instructions() IInstructionsContext

	// IsDefaultContext differentiates from other interfaces.
	IsDefaultContext()
}

type DefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultContext() *DefaultContext {
	var p = new(DefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_default
	return p
}

func InitEmptyDefaultContext(p *DefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_default
}

func (*DefaultContext) IsDefaultContext() {}

func NewDefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultContext {
	var p = new(DefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_default

	return p
}

func (s *DefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultContext) RW_default() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_default, 0)
}

func (s *DefaultContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *DefaultContext) Instructions() IInstructionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDefault(s)
	}
}

func (p *ParserParser) Default_() (localctx IDefaultContext) {
	localctx = NewDefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ParserParserRULE_default)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(ParserParserRW_default)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Instructions()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopforContext is an interface to support dynamic dispatch.
type ILoopforContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_for() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	RW_in() antlr.TerminalNode
	Env() IEnvContext
	Exp() IExpContext
	Range_() IRangeContext

	// IsLoopforContext differentiates from other interfaces.
	IsLoopforContext()
}

type LoopforContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopforContext() *LoopforContext {
	var p = new(LoopforContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loopfor
	return p
}

func InitEmptyLoopforContext(p *LoopforContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loopfor
}

func (*LoopforContext) IsLoopforContext() {}

func NewLoopforContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopforContext {
	var p = new(LoopforContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loopfor

	return p
}

func (s *LoopforContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopforContext) RW_for() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_for, 0)
}

func (s *LoopforContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *LoopforContext) RW_in() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_in, 0)
}

func (s *LoopforContext) Env() IEnvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvContext)
}

func (s *LoopforContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *LoopforContext) Range_() IRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *LoopforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopforContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopforContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLoopfor(s)
	}
}

func (s *LoopforContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLoopfor(s)
	}
}

func (p *ParserParser) Loopfor() (localctx ILoopforContext) {
	localctx = NewLoopforContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ParserParserRULE_loopfor)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(ParserParserRW_for)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(ParserParserRW_in)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(230)
			p.exp(0)
		}

	case 2:
		{
			p.SetState(231)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(234)
		p.Env()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExp() []IExpContext
	Exp(i int) IExpContext
	AllTK_dot() []antlr.TerminalNode
	TK_dot(i int) antlr.TerminalNode

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_range
	return p
}

func InitEmptyRangeContext(p *RangeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_range
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *RangeContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *RangeContext) AllTK_dot() []antlr.TerminalNode {
	return s.GetTokens(ParserParserTK_dot)
}

func (s *RangeContext) TK_dot(i int) antlr.TerminalNode {
	return s.GetToken(ParserParserTK_dot, i)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitRange(s)
	}
}

func (p *ParserParser) Range_() (localctx IRangeContext) {
	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ParserParserRULE_range)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.exp(0)
	}
	{
		p.SetState(237)
		p.Match(ParserParserTK_dot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(238)
		p.Match(ParserParserTK_dot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(239)
		p.Match(ParserParserTK_dot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.exp(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopwhileContext is an interface to support dynamic dispatch.
type ILoopwhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_while() antlr.TerminalNode
	Exp() IExpContext
	Env() IEnvContext

	// IsLoopwhileContext differentiates from other interfaces.
	IsLoopwhileContext()
}

type LoopwhileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopwhileContext() *LoopwhileContext {
	var p = new(LoopwhileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loopwhile
	return p
}

func InitEmptyLoopwhileContext(p *LoopwhileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_loopwhile
}

func (*LoopwhileContext) IsLoopwhileContext() {}

func NewLoopwhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopwhileContext {
	var p = new(LoopwhileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_loopwhile

	return p
}

func (s *LoopwhileContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopwhileContext) RW_while() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_while, 0)
}

func (s *LoopwhileContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *LoopwhileContext) Env() IEnvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvContext)
}

func (s *LoopwhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopwhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopwhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterLoopwhile(s)
	}
}

func (s *LoopwhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitLoopwhile(s)
	}
}

func (p *ParserParser) Loopwhile() (localctx ILoopwhileContext) {
	localctx = NewLoopwhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ParserParserRULE_loopwhile)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(ParserParserRW_while)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.exp(0)
	}
	{
		p.SetState(244)
		p.Env()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGuardContext is an interface to support dynamic dispatch.
type IGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_guard() antlr.TerminalNode
	Exp() IExpContext
	RW_else() antlr.TerminalNode
	Env() IEnvContext

	// IsGuardContext differentiates from other interfaces.
	IsGuardContext()
}

type GuardContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuardContext() *GuardContext {
	var p = new(GuardContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_guard
	return p
}

func InitEmptyGuardContext(p *GuardContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_guard
}

func (*GuardContext) IsGuardContext() {}

func NewGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GuardContext {
	var p = new(GuardContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_guard

	return p
}

func (s *GuardContext) GetParser() antlr.Parser { return s.parser }

func (s *GuardContext) RW_guard() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_guard, 0)
}

func (s *GuardContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *GuardContext) RW_else() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_else, 0)
}

func (s *GuardContext) Env() IEnvContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnvContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnvContext)
}

func (s *GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterGuard(s)
	}
}

func (s *GuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitGuard(s)
	}
}

func (p *ParserParser) Guard() (localctx IGuardContext) {
	localctx = NewGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ParserParserRULE_guard)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(ParserParserRW_guard)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.exp(0)
	}
	{
		p.SetState(248)
		p.Match(ParserParserRW_else)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Env()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReasignContext is an interface to support dynamic dispatch.
type IReasignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	Exp() IExpContext

	// IsReasignContext differentiates from other interfaces.
	IsReasignContext()
}

type ReasignContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReasignContext() *ReasignContext {
	var p = new(ReasignContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_reasign
	return p
}

func InitEmptyReasignContext(p *ReasignContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_reasign
}

func (*ReasignContext) IsReasignContext() {}

func NewReasignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReasignContext {
	var p = new(ReasignContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_reasign

	return p
}

func (s *ReasignContext) GetParser() antlr.Parser { return s.parser }

func (s *ReasignContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *ReasignContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *ReasignContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ReasignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReasignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReasignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterReasign(s)
	}
}

func (s *ReasignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitReasign(s)
	}
}

func (p *ParserParser) Reasign() (localctx IReasignContext) {
	localctx = NewReasignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ParserParserRULE_reasign)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.exp(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAddsubContext is an interface to support dynamic dispatch.
type IAddsubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	Exp() IExpContext
	TK_add() antlr.TerminalNode
	TK_sub() antlr.TerminalNode

	// IsAddsubContext differentiates from other interfaces.
	IsAddsubContext()
}

type AddsubContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddsubContext() *AddsubContext {
	var p = new(AddsubContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_addsub
	return p
}

func InitEmptyAddsubContext(p *AddsubContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_addsub
}

func (*AddsubContext) IsAddsubContext() {}

func NewAddsubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddsubContext {
	var p = new(AddsubContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_addsub

	return p
}

func (s *AddsubContext) GetParser() antlr.Parser { return s.parser }

func (s *AddsubContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *AddsubContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *AddsubContext) TK_add() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_add, 0)
}

func (s *AddsubContext) TK_sub() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_sub, 0)
}

func (s *AddsubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddsubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddsubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterAddsub(s)
	}
}

func (s *AddsubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitAddsub(s)
	}
}

func (p *ParserParser) Addsub() (localctx IAddsubContext) {
	localctx = NewAddsubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ParserParserRULE_addsub)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(255)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParserParserTK_add || _la == ParserParserTK_sub) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(257)
		p.exp(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDecvectorContext is an interface to support dynamic dispatch.
type IDecvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_var() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	TK_lbrk() antlr.TerminalNode
	Type_() ITypeContext
	TK_rbrk() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	Defvector() IDefvectorContext

	// IsDecvectorContext differentiates from other interfaces.
	IsDecvectorContext()
}

type DecvectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecvectorContext() *DecvectorContext {
	var p = new(DecvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decvector
	return p
}

func InitEmptyDecvectorContext(p *DecvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decvector
}

func (*DecvectorContext) IsDecvectorContext() {}

func NewDecvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecvectorContext {
	var p = new(DecvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_decvector

	return p
}

func (s *DecvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *DecvectorContext) RW_var() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_var, 0)
}

func (s *DecvectorContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DecvectorContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *DecvectorContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *DecvectorContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DecvectorContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *DecvectorContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *DecvectorContext) Defvector() IDefvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefvectorContext)
}

func (s *DecvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDecvector(s)
	}
}

func (s *DecvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDecvector(s)
	}
}

func (p *ParserParser) Decvector() (localctx IDecvectorContext) {
	localctx = NewDecvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ParserParserRULE_decvector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(ParserParserRW_var)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(260)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(ParserParserTK_lbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
		p.Type_()
	}
	{
		p.SetState(264)
		p.Match(ParserParserTK_rbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(265)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Defvector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefvectorContext is an interface to support dynamic dispatch.
type IDefvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrk() antlr.TerminalNode
	TK_rbrk() antlr.TerminalNode
	Listexp() IListexpContext
	Simplevec() ISimplevecContext
	TK_id() antlr.TerminalNode

	// IsDefvectorContext differentiates from other interfaces.
	IsDefvectorContext()
}

type DefvectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefvectorContext() *DefvectorContext {
	var p = new(DefvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_defvector
	return p
}

func InitEmptyDefvectorContext(p *DefvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_defvector
}

func (*DefvectorContext) IsDefvectorContext() {}

func NewDefvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefvectorContext {
	var p = new(DefvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_defvector

	return p
}

func (s *DefvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *DefvectorContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *DefvectorContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *DefvectorContext) Listexp() IListexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListexpContext)
}

func (s *DefvectorContext) Simplevec() ISimplevecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplevecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplevecContext)
}

func (s *DefvectorContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DefvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDefvector(s)
	}
}

func (s *DefvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDefvector(s)
	}
}

func (p *ParserParser) Defvector() (localctx IDefvectorContext) {
	localctx = NewDefvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ParserParserRULE_defvector)
	var _la int

	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216182643362365454) != 0 {
			{
				p.SetState(269)
				p.Listexp()
			}

		}
		{
			p.SetState(272)
			p.Match(ParserParserTK_rbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(273)
			p.Simplevec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(274)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IListexpContext is an interface to support dynamic dispatch.
type IListexpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Exp() IExpContext
	TK_comma() antlr.TerminalNode
	Listexp() IListexpContext

	// IsListexpContext differentiates from other interfaces.
	IsListexpContext()
}

type ListexpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListexpContext() *ListexpContext {
	var p = new(ListexpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listexp
	return p
}

func InitEmptyListexpContext(p *ListexpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listexp
}

func (*ListexpContext) IsListexpContext() {}

func NewListexpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListexpContext {
	var p = new(ListexpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listexp

	return p
}

func (s *ListexpContext) GetParser() antlr.Parser { return s.parser }

func (s *ListexpContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ListexpContext) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *ListexpContext) Listexp() IListexpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListexpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListexpContext)
}

func (s *ListexpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListexpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListexpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListexp(s)
	}
}

func (s *ListexpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListexp(s)
	}
}

func (p *ParserParser) Listexp() (localctx IListexpContext) {
	localctx = NewListexpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ParserParserRULE_listexp)
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(277)
			p.exp(0)
		}
		{
			p.SetState(278)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Listexp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.exp(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimplevecContext is an interface to support dynamic dispatch.
type ISimplevecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrk() antlr.TerminalNode
	Type_() ITypeContext
	TK_rbrk() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	RW_repeating() antlr.TerminalNode
	AllTK_colon() []antlr.TerminalNode
	TK_colon(i int) antlr.TerminalNode
	AllExp() []IExpContext
	Exp(i int) IExpContext
	TK_comma() antlr.TerminalNode
	RW_count() antlr.TerminalNode
	TK_rpar() antlr.TerminalNode

	// IsSimplevecContext differentiates from other interfaces.
	IsSimplevecContext()
}

type SimplevecContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimplevecContext() *SimplevecContext {
	var p = new(SimplevecContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_simplevec
	return p
}

func InitEmptySimplevecContext(p *SimplevecContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_simplevec
}

func (*SimplevecContext) IsSimplevecContext() {}

func NewSimplevecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimplevecContext {
	var p = new(SimplevecContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_simplevec

	return p
}

func (s *SimplevecContext) GetParser() antlr.Parser { return s.parser }

func (s *SimplevecContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *SimplevecContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *SimplevecContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *SimplevecContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *SimplevecContext) RW_repeating() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_repeating, 0)
}

func (s *SimplevecContext) AllTK_colon() []antlr.TerminalNode {
	return s.GetTokens(ParserParserTK_colon)
}

func (s *SimplevecContext) TK_colon(i int) antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, i)
}

func (s *SimplevecContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *SimplevecContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *SimplevecContext) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *SimplevecContext) RW_count() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_count, 0)
}

func (s *SimplevecContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *SimplevecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplevecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimplevecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterSimplevec(s)
	}
}

func (s *SimplevecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitSimplevec(s)
	}
}

func (p *ParserParser) Simplevec() (localctx ISimplevecContext) {
	localctx = NewSimplevecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ParserParserRULE_simplevec)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(ParserParserTK_lbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Type_()
	}
	{
		p.SetState(286)
		p.Match(ParserParserTK_rbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Match(ParserParserRW_repeating)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(289)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.exp(0)
	}
	{
		p.SetState(291)
		p.Match(ParserParserTK_comma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(ParserParserRW_count)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(294)
		p.exp(0)
	}
	{
		p.SetState(295)
		p.Match(ParserParserTK_rpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncvectorContext is an interface to support dynamic dispatch.
type IFuncvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	TK_dot() antlr.TerminalNode
	RW_append() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	Exp() IExpContext
	TK_rpar() antlr.TerminalNode
	RW_removeLast() antlr.TerminalNode
	RW_remove() antlr.TerminalNode
	RW_at() antlr.TerminalNode
	TK_colon() antlr.TerminalNode

	// IsFuncvectorContext differentiates from other interfaces.
	IsFuncvectorContext()
}

type FuncvectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncvectorContext() *FuncvectorContext {
	var p = new(FuncvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_funcvector
	return p
}

func InitEmptyFuncvectorContext(p *FuncvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_funcvector
}

func (*FuncvectorContext) IsFuncvectorContext() {}

func NewFuncvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncvectorContext {
	var p = new(FuncvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_funcvector

	return p
}

func (s *FuncvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncvectorContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *FuncvectorContext) TK_dot() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_dot, 0)
}

func (s *FuncvectorContext) RW_append() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_append, 0)
}

func (s *FuncvectorContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *FuncvectorContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *FuncvectorContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *FuncvectorContext) RW_removeLast() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_removeLast, 0)
}

func (s *FuncvectorContext) RW_remove() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_remove, 0)
}

func (s *FuncvectorContext) RW_at() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_at, 0)
}

func (s *FuncvectorContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *FuncvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterFuncvector(s)
	}
}

func (s *FuncvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitFuncvector(s)
	}
}

func (p *ParserParser) Funcvector() (localctx IFuncvectorContext) {
	localctx = NewFuncvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ParserParserRULE_funcvector)
	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(298)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(299)
			p.Match(ParserParserRW_append)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(300)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(301)
			p.exp(0)
		}
		{
			p.SetState(302)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(304)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(305)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.Match(ParserParserRW_removeLast)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(307)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(308)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(309)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(310)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.Match(ParserParserRW_remove)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(313)
			p.Match(ParserParserRW_at)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(314)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(315)
			p.exp(0)
		}
		{
			p.SetState(316)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReasignvectorContext is an interface to support dynamic dispatch.
type IReasignvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	TK_lbrk() antlr.TerminalNode
	AllExp() []IExpContext
	Exp(i int) IExpContext
	TK_rbrk() antlr.TerminalNode
	TK_equ() antlr.TerminalNode

	// IsReasignvectorContext differentiates from other interfaces.
	IsReasignvectorContext()
}

type ReasignvectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReasignvectorContext() *ReasignvectorContext {
	var p = new(ReasignvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_reasignvector
	return p
}

func InitEmptyReasignvectorContext(p *ReasignvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_reasignvector
}

func (*ReasignvectorContext) IsReasignvectorContext() {}

func NewReasignvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReasignvectorContext {
	var p = new(ReasignvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_reasignvector

	return p
}

func (s *ReasignvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ReasignvectorContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *ReasignvectorContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *ReasignvectorContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ReasignvectorContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ReasignvectorContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *ReasignvectorContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *ReasignvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReasignvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReasignvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterReasignvector(s)
	}
}

func (s *ReasignvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitReasignvector(s)
	}
}

func (p *ParserParser) Reasignvector() (localctx IReasignvectorContext) {
	localctx = NewReasignvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ParserParserRULE_reasignvector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
		p.Match(ParserParserTK_lbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.exp(0)
	}
	{
		p.SetState(323)
		p.Match(ParserParserTK_rbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(325)
		p.exp(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrintContext is an interface to support dynamic dispatch.
type IPrintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_print() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	TK_rpar() antlr.TerminalNode
	Exp() IExpContext

	// IsPrintContext differentiates from other interfaces.
	IsPrintContext()
}

type PrintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintContext() *PrintContext {
	var p = new(PrintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_print
	return p
}

func InitEmptyPrintContext(p *PrintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_print
}

func (*PrintContext) IsPrintContext() {}

func NewPrintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintContext {
	var p = new(PrintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_print

	return p
}

func (s *PrintContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintContext) RW_print() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_print, 0)
}

func (s *PrintContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *PrintContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *PrintContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (p *ParserParser) Print_() (localctx IPrintContext) {
	localctx = NewPrintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ParserParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(ParserParserRW_print)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(328)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&216182643362365454) != 0 {
		{
			p.SetState(329)
			p.exp(0)
		}

	}
	{
		p.SetState(332)
		p.Match(ParserParserTK_rpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnvContext is an interface to support dynamic dispatch.
type IEnvContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrc() antlr.TerminalNode
	TK_rbrc() antlr.TerminalNode
	Instructions() IInstructionsContext

	// IsEnvContext differentiates from other interfaces.
	IsEnvContext()
}

type EnvContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnvContext() *EnvContext {
	var p = new(EnvContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_env
	return p
}

func InitEmptyEnvContext(p *EnvContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_env
}

func (*EnvContext) IsEnvContext() {}

func NewEnvContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnvContext {
	var p = new(EnvContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_env

	return p
}

func (s *EnvContext) GetParser() antlr.Parser { return s.parser }

func (s *EnvContext) TK_lbrc() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrc, 0)
}

func (s *EnvContext) TK_rbrc() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrc, 0)
}

func (s *EnvContext) Instructions() IInstructionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *EnvContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnvContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnvContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterEnv(s)
	}
}

func (s *EnvContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitEnv(s)
	}
}

func (p *ParserParser) Env() (localctx IEnvContext) {
	localctx = NewEnvContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ParserParserRULE_env)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(ParserParserTK_lbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&554051255744) != 0 {
		{
			p.SetState(335)
			p.Instructions()
		}

	}
	{
		p.SetState(338)
		p.Match(ParserParserTK_rbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instruction() IInstructionContext
	Instructions() IInstructionsContext

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instructions
	return p
}

func InitEmptyInstructionsContext(p *InstructionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instructions
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionsContext) Instructions() IInstructionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInstructions(s)
	}
}

func (p *ParserParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ParserParserRULE_instructions)
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.Instruction()
		}
		{
			p.SetState(341)
			p.Instructions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Instruction()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Decvar() IDecvarContext
	Deccst() IDeccstContext
	Ifstruct() IIfstructContext
	Switchstruct() ISwitchstructContext
	Loopfor() ILoopforContext
	Loopwhile() ILoopwhileContext
	Guard() IGuardContext
	Reasign() IReasignContext
	Addsub() IAddsubContext
	Decvector() IDecvectorContext
	Funcvector() IFuncvectorContext
	Reasignvector() IReasignvectorContext
	Callfunc() ICallfuncContext
	Print_() IPrintContext
	RW_return() antlr.TerminalNode
	Exp() IExpContext
	RW_continue() antlr.TerminalNode
	RW_break() antlr.TerminalNode

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Decvar() IDecvarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecvarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecvarContext)
}

func (s *InstructionContext) Deccst() IDeccstContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeccstContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeccstContext)
}

func (s *InstructionContext) Ifstruct() IIfstructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstructContext)
}

func (s *InstructionContext) Switchstruct() ISwitchstructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchstructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchstructContext)
}

func (s *InstructionContext) Loopfor() ILoopforContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopforContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopforContext)
}

func (s *InstructionContext) Loopwhile() ILoopwhileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopwhileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopwhileContext)
}

func (s *InstructionContext) Guard() IGuardContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuardContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuardContext)
}

func (s *InstructionContext) Reasign() IReasignContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReasignContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReasignContext)
}

func (s *InstructionContext) Addsub() IAddsubContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddsubContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddsubContext)
}

func (s *InstructionContext) Decvector() IDecvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecvectorContext)
}

func (s *InstructionContext) Funcvector() IFuncvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncvectorContext)
}

func (s *InstructionContext) Reasignvector() IReasignvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReasignvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReasignvectorContext)
}

func (s *InstructionContext) Callfunc() ICallfuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallfuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallfuncContext)
}

func (s *InstructionContext) Print_() IPrintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintContext)
}

func (s *InstructionContext) RW_return() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_return, 0)
}

func (s *InstructionContext) Exp() IExpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *InstructionContext) RW_continue() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_continue, 0)
}

func (s *InstructionContext) RW_break() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_break, 0)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *ParserParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ParserParserRULE_instruction)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)
			p.Decvar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Deccst()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(348)
			p.Ifstruct()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(349)
			p.Switchstruct()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(350)
			p.Loopfor()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(351)
			p.Loopwhile()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(352)
			p.Guard()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(353)
			p.Reasign()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(354)
			p.Addsub()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(355)
			p.Decvector()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(356)
			p.Funcvector()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(357)
			p.Reasignvector()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(358)
			p.Callfunc()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(359)
			p.Print_()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(360)
			p.Match(ParserParserRW_return)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(361)
			p.exp(0)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(362)
			p.Match(ParserParserRW_return)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(363)
			p.Match(ParserParserRW_continue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(364)
			p.Match(ParserParserRW_break)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_String() antlr.TerminalNode
	RW_Int() antlr.TerminalNode
	RW_Bool() antlr.TerminalNode
	RW_Character() antlr.TerminalNode
	RW_Float() antlr.TerminalNode

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) RW_String() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_String, 0)
}

func (s *TypeContext) RW_Int() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_Int, 0)
}

func (s *TypeContext) RW_Bool() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_Bool, 0)
}

func (s *TypeContext) RW_Character() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_Character, 0)
}

func (s *TypeContext) RW_Float() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_Float, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *ParserParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ParserParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(367)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&62) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetS returns the s token.
	GetS() antlr.Token

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// GetP returns the p token.
	GetP() antlr.Token

	// SetS sets the s token.
	SetS(antlr.Token)

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpContext

	// GetE returns the e rule contexts.
	GetE() IExpContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpContext)

	// SetE sets the e rule contexts.
	SetE(IExpContext)

	// Getter signatures
	TK_minus() antlr.TerminalNode
	AllExp() []IExpContext
	Exp(i int) IExpContext
	TK_not() antlr.TerminalNode
	RW_Int() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	TK_rpar() antlr.TerminalNode
	RW_Float() antlr.TerminalNode
	RW_String() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_lbrk() antlr.TerminalNode
	TK_rbrk() antlr.TerminalNode
	TK_dot() antlr.TerminalNode
	RW_isEmpty() antlr.TerminalNode
	RW_count() antlr.TerminalNode
	Callfunc() ICallfuncContext
	RW_nil() antlr.TerminalNode
	TK_string() antlr.TerminalNode
	TK_char() antlr.TerminalNode
	TK_int() antlr.TerminalNode
	TK_float() antlr.TerminalNode
	RW_true() antlr.TerminalNode
	RW_false() antlr.TerminalNode
	TK_mult() antlr.TerminalNode
	TK_div() antlr.TerminalNode
	TK_mod() antlr.TerminalNode
	TK_plus() antlr.TerminalNode
	TK_lessequ() antlr.TerminalNode
	TK_moreequ() antlr.TerminalNode
	TK_less() antlr.TerminalNode
	TK_more() antlr.TerminalNode
	TK_equequ() antlr.TerminalNode
	TK_notequ() antlr.TerminalNode
	TK_and() antlr.TerminalNode
	TK_or() antlr.TerminalNode

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	e1     IExpContext
	s      antlr.Token
	e2     IExpContext
	n      antlr.Token
	id     antlr.Token
	p      antlr.Token
	e      IExpContext
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_exp
	return p
}

func InitEmptyExpContext(p *ExpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_exp
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) GetS() antlr.Token { return s.s }

func (s *ExpContext) GetN() antlr.Token { return s.n }

func (s *ExpContext) GetId() antlr.Token { return s.id }

func (s *ExpContext) GetP() antlr.Token { return s.p }

func (s *ExpContext) SetS(v antlr.Token) { s.s = v }

func (s *ExpContext) SetN(v antlr.Token) { s.n = v }

func (s *ExpContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpContext) SetP(v antlr.Token) { s.p = v }

func (s *ExpContext) GetE1() IExpContext { return s.e1 }

func (s *ExpContext) GetE2() IExpContext { return s.e2 }

func (s *ExpContext) GetE() IExpContext { return s.e }

func (s *ExpContext) SetE1(v IExpContext) { s.e1 = v }

func (s *ExpContext) SetE2(v IExpContext) { s.e2 = v }

func (s *ExpContext) SetE(v IExpContext) { s.e = v }

func (s *ExpContext) TK_minus() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_minus, 0)
}

func (s *ExpContext) AllExp() []IExpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpContext); ok {
			len++
		}
	}

	tst := make([]IExpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpContext); ok {
			tst[i] = t.(IExpContext)
			i++
		}
	}

	return tst
}

func (s *ExpContext) Exp(i int) IExpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpContext) TK_not() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_not, 0)
}

func (s *ExpContext) RW_Int() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_Int, 0)
}

func (s *ExpContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *ExpContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *ExpContext) RW_Float() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_Float, 0)
}

func (s *ExpContext) RW_String() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_String, 0)
}

func (s *ExpContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *ExpContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *ExpContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *ExpContext) TK_dot() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_dot, 0)
}

func (s *ExpContext) RW_isEmpty() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_isEmpty, 0)
}

func (s *ExpContext) RW_count() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_count, 0)
}

func (s *ExpContext) Callfunc() ICallfuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallfuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallfuncContext)
}

func (s *ExpContext) RW_nil() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_nil, 0)
}

func (s *ExpContext) TK_string() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_string, 0)
}

func (s *ExpContext) TK_char() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_char, 0)
}

func (s *ExpContext) TK_int() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_int, 0)
}

func (s *ExpContext) TK_float() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_float, 0)
}

func (s *ExpContext) RW_true() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_true, 0)
}

func (s *ExpContext) RW_false() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_false, 0)
}

func (s *ExpContext) TK_mult() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_mult, 0)
}

func (s *ExpContext) TK_div() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_div, 0)
}

func (s *ExpContext) TK_mod() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_mod, 0)
}

func (s *ExpContext) TK_plus() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_plus, 0)
}

func (s *ExpContext) TK_lessequ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lessequ, 0)
}

func (s *ExpContext) TK_moreequ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_moreequ, 0)
}

func (s *ExpContext) TK_less() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_less, 0)
}

func (s *ExpContext) TK_more() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_more, 0)
}

func (s *ExpContext) TK_equequ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equequ, 0)
}

func (s *ExpContext) TK_notequ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_notequ, 0)
}

func (s *ExpContext) TK_and() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_and, 0)
}

func (s *ExpContext) TK_or() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_or, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *ParserParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *ParserParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, ParserParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(413)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(370)

			var _m = p.Match(ParserParserTK_minus)

			localctx.(*ExpContext).s = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(371)

			var _x = p.exp(25)

			localctx.(*ExpContext).e2 = _x
		}

	case 2:
		{
			p.SetState(372)

			var _m = p.Match(ParserParserTK_not)

			localctx.(*ExpContext).s = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(373)

			var _x = p.exp(19)

			localctx.(*ExpContext).e2 = _x
		}

	case 3:
		{
			p.SetState(374)
			p.Match(ParserParserRW_Int)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			p.exp(0)
		}
		{
			p.SetState(377)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(379)
			p.Match(ParserParserRW_Float)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.exp(0)
		}
		{
			p.SetState(382)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(384)
			p.Match(ParserParserRW_String)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(386)
			p.exp(0)
		}
		{
			p.SetState(387)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(389)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.exp(0)
		}
		{
			p.SetState(392)
			p.Match(ParserParserTK_rbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		{
			p.SetState(394)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(ParserParserRW_isEmpty)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		{
			p.SetState(397)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.Match(ParserParserRW_count)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		{
			p.SetState(400)
			p.Callfunc()
		}

	case 10:
		{
			p.SetState(401)

			var _m = p.Match(ParserParserRW_nil)

			localctx.(*ExpContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		{
			p.SetState(402)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*ExpContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		{
			p.SetState(403)

			var _m = p.Match(ParserParserTK_string)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		{
			p.SetState(404)

			var _m = p.Match(ParserParserTK_char)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		{
			p.SetState(405)

			var _m = p.Match(ParserParserTK_int)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		{
			p.SetState(406)

			var _m = p.Match(ParserParserTK_float)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		{
			p.SetState(407)

			var _m = p.Match(ParserParserRW_true)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		{
			p.SetState(408)

			var _m = p.Match(ParserParserRW_false)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		{
			p.SetState(409)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)

			var _x = p.exp(0)

			localctx.(*ExpContext).e = _x
		}
		{
			p.SetState(411)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(436)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(415)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(416)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).s = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123145302310912) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).s = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(417)

					var _x = p.exp(25)

					localctx.(*ExpContext).e2 = _x
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(418)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(419)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).s = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ParserParserTK_plus || _la == ParserParserTK_minus) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).s = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(420)

					var _x = p.exp(24)

					localctx.(*ExpContext).e2 = _x
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(421)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(422)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).s = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ParserParserTK_lessequ || _la == ParserParserTK_moreequ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).s = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(423)

					var _x = p.exp(23)

					localctx.(*ExpContext).e2 = _x
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(424)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(425)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).s = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ParserParserTK_less || _la == ParserParserTK_more) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).s = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(426)

					var _x = p.exp(22)

					localctx.(*ExpContext).e2 = _x
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(427)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
					goto errorExit
				}
				{
					p.SetState(428)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).s = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ParserParserTK_equequ || _la == ParserParserTK_notequ) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).s = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(429)

					var _x = p.exp(21)

					localctx.(*ExpContext).e2 = _x
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(430)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(431)

					var _m = p.Match(ParserParserTK_and)

					localctx.(*ExpContext).s = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(432)

					var _x = p.exp(19)

					localctx.(*ExpContext).e2 = _x
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(433)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(434)

					var _m = p.Match(ParserParserTK_or)

					localctx.(*ExpContext).s = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(435)

					var _x = p.exp(18)

					localctx.(*ExpContext).e2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *ParserParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 33:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParserParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 17)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
