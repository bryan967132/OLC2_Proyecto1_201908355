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
		"RW_remove", "RW_at", "RW_isEmpty", "RW_count", "RW_repeating", "RW_struct",
		"RW_mutating", "RW_self", "RW_print", "TK_prompt", "TK_under", "TK_char",
		"TK_string", "TK_int", "TK_float", "TK_id", "TK_add", "TK_sub", "TK_plus",
		"TK_minus", "TK_mult", "TK_div", "TK_mod", "TK_equequ", "TK_notequ",
		"TK_lessequ", "TK_moreequ", "TK_equ", "TK_less", "TK_more", "TK_and",
		"TK_or", "TK_not", "TK_lpar", "TK_rpar", "TK_lbrc", "TK_rbrc", "TK_lbrk",
		"TK_rbrk", "TK_dot", "TK_comma", "TK_colon", "TK_semicolon", "TK_question",
		"TK_amp",
	}
	staticData.RuleNames = []string{
		"init", "instsglobal", "instglobal", "callfunc", "listargs", "decvar",
		"deccst", "declfunc", "listparams", "ifstruct", "switchstruct", "envs",
		"casesdefault", "cases", "case", "default", "loopfor", "range", "loopwhile",
		"guard", "reasign", "addsub", "decvector", "defvector", "listexp", "simplevec",
		"funcvector", "decmatrix", "typematrix", "defmatrix", "listvector",
		"listvector2", "simplematrix", "reasignvector", "dims", "defstruct",
		"listattribs", "attrib", "decstruct", "listdupla", "useattribs", "obj",
		"useattribs1", "print", "env", "instructions", "instruction", "type",
		"exp",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 71, 724, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 1, 0, 3, 0, 100, 8, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 108, 8, 1, 1, 2, 1, 2, 3, 2, 112, 8, 2, 1, 3, 1, 3, 1, 3,
		3, 3, 117, 8, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 123, 8, 4, 1, 4, 3, 4, 126,
		8, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 134, 8, 4, 1, 4, 3, 4,
		137, 8, 4, 1, 4, 3, 4, 140, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 159,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 172, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 178, 8, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 183, 8, 7, 1, 7, 1, 7, 1, 8, 3, 8, 188, 8, 8, 1, 8, 1, 8, 1, 8,
		3, 8, 193, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 200, 8, 8, 1, 8, 1,
		8, 1, 8, 3, 8, 205, 8, 8, 1, 8, 3, 8, 208, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 226, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 241, 8, 12, 1, 13, 1, 13, 1,
		13, 1, 13, 3, 13, 247, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 263, 8,
		16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 3, 23, 301, 8, 23, 1, 23, 1, 23, 1,
		23, 3, 23, 306, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 313, 8,
		24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 3, 26, 349, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 3,
		27, 355, 8, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 3, 28, 368, 8, 28, 1, 29, 1, 29, 3, 29, 372, 8, 29,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 3, 31, 381, 8, 31, 1,
		31, 1, 31, 1, 31, 5, 31, 386, 8, 31, 10, 31, 12, 31, 389, 9, 31, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		3, 32, 413, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 3, 34, 429, 8, 34, 1, 35,
		1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36, 439, 8, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 3, 36, 445, 8, 36, 3, 36, 447, 8, 36, 1, 37, 1,
		37, 1, 37, 1, 37, 3, 37, 453, 8, 37, 1, 37, 1, 37, 3, 37, 457, 8, 37, 1,
		37, 3, 37, 460, 8, 37, 1, 37, 3, 37, 463, 8, 37, 1, 38, 1, 38, 1, 38, 1,
		38, 3, 38, 469, 8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 475, 8, 38, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 482, 8, 38, 1, 38, 1, 38, 1, 38,
		1, 38, 3, 38, 488, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1,
		39, 1, 39, 1, 39, 3, 39, 499, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40,
		1, 40, 1, 40, 3, 40, 508, 8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 3, 41, 516, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 523, 8,
		42, 1, 43, 1, 43, 1, 43, 3, 43, 528, 8, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		3, 44, 534, 8, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 542,
		8, 45, 1, 46, 1, 46, 3, 46, 546, 8, 46, 1, 46, 1, 46, 3, 46, 550, 8, 46,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 3, 46, 559, 8, 46, 1,
		46, 1, 46, 3, 46, 563, 8, 46, 1, 46, 1, 46, 3, 46, 567, 8, 46, 1, 46, 1,
		46, 3, 46, 571, 8, 46, 1, 46, 1, 46, 3, 46, 575, 8, 46, 1, 46, 1, 46, 3,
		46, 579, 8, 46, 1, 46, 1, 46, 3, 46, 583, 8, 46, 1, 46, 1, 46, 3, 46, 587,
		8, 46, 1, 46, 1, 46, 3, 46, 591, 8, 46, 1, 46, 1, 46, 3, 46, 595, 8, 46,
		1, 46, 1, 46, 3, 46, 599, 8, 46, 1, 46, 1, 46, 3, 46, 603, 8, 46, 1, 46,
		1, 46, 3, 46, 607, 8, 46, 1, 46, 1, 46, 3, 46, 611, 8, 46, 1, 46, 1, 46,
		3, 46, 615, 8, 46, 1, 46, 1, 46, 3, 46, 619, 8, 46, 1, 46, 1, 46, 1, 46,
		3, 46, 624, 8, 46, 1, 46, 1, 46, 3, 46, 628, 8, 46, 1, 46, 1, 46, 3, 46,
		632, 8, 46, 1, 46, 1, 46, 3, 46, 636, 8, 46, 3, 46, 638, 8, 46, 1, 47,
		1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3,
		48, 672, 8, 48, 1, 48, 1, 48, 1, 48, 3, 48, 677, 8, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 3, 48, 683, 8, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 3, 48, 696, 8, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 5, 48, 719, 8,
		48, 10, 48, 12, 48, 722, 9, 48, 1, 48, 0, 2, 62, 96, 49, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		82, 84, 86, 88, 90, 92, 94, 96, 0, 9, 2, 0, 37, 37, 42, 42, 1, 0, 43, 44,
		1, 0, 6, 7, 2, 0, 1, 5, 42, 42, 1, 0, 47, 49, 1, 0, 45, 46, 1, 0, 52, 53,
		1, 0, 55, 56, 1, 0, 50, 51, 801, 0, 99, 1, 0, 0, 0, 2, 107, 1, 0, 0, 0,
		4, 111, 1, 0, 0, 0, 6, 113, 1, 0, 0, 0, 8, 139, 1, 0, 0, 0, 10, 158, 1,
		0, 0, 0, 12, 171, 1, 0, 0, 0, 14, 173, 1, 0, 0, 0, 16, 207, 1, 0, 0, 0,
		18, 225, 1, 0, 0, 0, 20, 227, 1, 0, 0, 0, 22, 231, 1, 0, 0, 0, 24, 240,
		1, 0, 0, 0, 26, 246, 1, 0, 0, 0, 28, 248, 1, 0, 0, 0, 30, 253, 1, 0, 0,
		0, 32, 257, 1, 0, 0, 0, 34, 266, 1, 0, 0, 0, 36, 272, 1, 0, 0, 0, 38, 276,
		1, 0, 0, 0, 40, 281, 1, 0, 0, 0, 42, 285, 1, 0, 0, 0, 44, 289, 1, 0, 0,
		0, 46, 305, 1, 0, 0, 0, 48, 312, 1, 0, 0, 0, 50, 314, 1, 0, 0, 0, 52, 348,
		1, 0, 0, 0, 54, 350, 1, 0, 0, 0, 56, 367, 1, 0, 0, 0, 58, 371, 1, 0, 0,
		0, 60, 373, 1, 0, 0, 0, 62, 380, 1, 0, 0, 0, 64, 412, 1, 0, 0, 0, 66, 414,
		1, 0, 0, 0, 68, 428, 1, 0, 0, 0, 70, 430, 1, 0, 0, 0, 72, 446, 1, 0, 0,
		0, 74, 462, 1, 0, 0, 0, 76, 487, 1, 0, 0, 0, 78, 498, 1, 0, 0, 0, 80, 507,
		1, 0, 0, 0, 82, 515, 1, 0, 0, 0, 84, 522, 1, 0, 0, 0, 86, 524, 1, 0, 0,
		0, 88, 531, 1, 0, 0, 0, 90, 541, 1, 0, 0, 0, 92, 637, 1, 0, 0, 0, 94, 639,
		1, 0, 0, 0, 96, 695, 1, 0, 0, 0, 98, 100, 3, 2, 1, 0, 99, 98, 1, 0, 0,
		0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 102, 5, 0, 0, 1, 102,
		1, 1, 0, 0, 0, 103, 104, 3, 4, 2, 0, 104, 105, 3, 2, 1, 0, 105, 108, 1,
		0, 0, 0, 106, 108, 3, 4, 2, 0, 107, 103, 1, 0, 0, 0, 107, 106, 1, 0, 0,
		0, 108, 3, 1, 0, 0, 0, 109, 112, 3, 92, 46, 0, 110, 112, 3, 14, 7, 0, 111,
		109, 1, 0, 0, 0, 111, 110, 1, 0, 0, 0, 112, 5, 1, 0, 0, 0, 113, 114, 5,
		42, 0, 0, 114, 116, 5, 60, 0, 0, 115, 117, 3, 8, 4, 0, 116, 115, 1, 0,
		0, 0, 116, 117, 1, 0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 5, 61, 0, 0,
		119, 7, 1, 0, 0, 0, 120, 121, 5, 42, 0, 0, 121, 123, 5, 68, 0, 0, 122,
		120, 1, 0, 0, 0, 122, 123, 1, 0, 0, 0, 123, 125, 1, 0, 0, 0, 124, 126,
		5, 71, 0, 0, 125, 124, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 128, 3, 96, 48, 0, 128, 129, 5, 67, 0, 0, 129, 130, 3, 8, 4,
		0, 130, 140, 1, 0, 0, 0, 131, 132, 5, 42, 0, 0, 132, 134, 5, 68, 0, 0,
		133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0, 134, 136, 1, 0, 0, 0, 135,
		137, 5, 71, 0, 0, 136, 135, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 138,
		1, 0, 0, 0, 138, 140, 3, 96, 48, 0, 139, 122, 1, 0, 0, 0, 139, 133, 1,
		0, 0, 0, 140, 9, 1, 0, 0, 0, 141, 142, 5, 6, 0, 0, 142, 143, 5, 42, 0,
		0, 143, 144, 5, 68, 0, 0, 144, 145, 3, 94, 47, 0, 145, 146, 5, 54, 0, 0,
		146, 147, 3, 96, 48, 0, 147, 159, 1, 0, 0, 0, 148, 149, 5, 6, 0, 0, 149,
		150, 5, 42, 0, 0, 150, 151, 5, 68, 0, 0, 151, 152, 3, 94, 47, 0, 152, 153,
		5, 70, 0, 0, 153, 159, 1, 0, 0, 0, 154, 155, 5, 6, 0, 0, 155, 156, 5, 42,
		0, 0, 156, 157, 5, 54, 0, 0, 157, 159, 3, 96, 48, 0, 158, 141, 1, 0, 0,
		0, 158, 148, 1, 0, 0, 0, 158, 154, 1, 0, 0, 0, 159, 11, 1, 0, 0, 0, 160,
		161, 5, 7, 0, 0, 161, 162, 5, 42, 0, 0, 162, 163, 5, 68, 0, 0, 163, 164,
		3, 94, 47, 0, 164, 165, 5, 54, 0, 0, 165, 166, 3, 96, 48, 0, 166, 172,
		1, 0, 0, 0, 167, 168, 5, 7, 0, 0, 168, 169, 5, 42, 0, 0, 169, 170, 5, 54,
		0, 0, 170, 172, 3, 96, 48, 0, 171, 160, 1, 0, 0, 0, 171, 167, 1, 0, 0,
		0, 172, 13, 1, 0, 0, 0, 173, 174, 5, 22, 0, 0, 174, 175, 5, 42, 0, 0, 175,
		177, 5, 60, 0, 0, 176, 178, 3, 16, 8, 0, 177, 176, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 182, 5, 61, 0, 0, 180, 181, 5, 36,
		0, 0, 181, 183, 3, 94, 47, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0,
		0, 183, 184, 1, 0, 0, 0, 184, 185, 3, 88, 44, 0, 185, 15, 1, 0, 0, 0, 186,
		188, 7, 0, 0, 0, 187, 186, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 190, 5, 42, 0, 0, 190, 192, 5, 68, 0, 0, 191, 193, 5,
		23, 0, 0, 192, 191, 1, 0, 0, 0, 192, 193, 1, 0, 0, 0, 193, 194, 1, 0, 0,
		0, 194, 195, 3, 94, 47, 0, 195, 196, 5, 67, 0, 0, 196, 197, 3, 16, 8, 0,
		197, 208, 1, 0, 0, 0, 198, 200, 7, 0, 0, 0, 199, 198, 1, 0, 0, 0, 199,
		200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 5, 42, 0, 0, 202, 204,
		5, 68, 0, 0, 203, 205, 5, 23, 0, 0, 204, 203, 1, 0, 0, 0, 204, 205, 1,
		0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 208, 3, 94, 47, 0, 207, 187, 1, 0,
		0, 0, 207, 199, 1, 0, 0, 0, 208, 17, 1, 0, 0, 0, 209, 210, 5, 8, 0, 0,
		210, 211, 3, 96, 48, 0, 211, 212, 3, 88, 44, 0, 212, 213, 5, 9, 0, 0, 213,
		214, 3, 18, 9, 0, 214, 226, 1, 0, 0, 0, 215, 216, 5, 8, 0, 0, 216, 217,
		3, 96, 48, 0, 217, 218, 3, 88, 44, 0, 218, 219, 5, 9, 0, 0, 219, 220, 3,
		88, 44, 0, 220, 226, 1, 0, 0, 0, 221, 222, 5, 8, 0, 0, 222, 223, 3, 96,
		48, 0, 223, 224, 3, 88, 44, 0, 224, 226, 1, 0, 0, 0, 225, 209, 1, 0, 0,
		0, 225, 215, 1, 0, 0, 0, 225, 221, 1, 0, 0, 0, 226, 19, 1, 0, 0, 0, 227,
		228, 5, 13, 0, 0, 228, 229, 3, 96, 48, 0, 229, 230, 3, 22, 11, 0, 230,
		21, 1, 0, 0, 0, 231, 232, 5, 62, 0, 0, 232, 233, 3, 24, 12, 0, 233, 234,
		5, 63, 0, 0, 234, 23, 1, 0, 0, 0, 235, 236, 3, 26, 13, 0, 236, 237, 3,
		30, 15, 0, 237, 241, 1, 0, 0, 0, 238, 241, 3, 26, 13, 0, 239, 241, 3, 30,
		15, 0, 240, 235, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 239, 1, 0, 0, 0,
		241, 25, 1, 0, 0, 0, 242, 243, 3, 28, 14, 0, 243, 244, 3, 26, 13, 0, 244,
		247, 1, 0, 0, 0, 245, 247, 3, 28, 14, 0, 246, 242, 1, 0, 0, 0, 246, 245,
		1, 0, 0, 0, 247, 27, 1, 0, 0, 0, 248, 249, 5, 14, 0, 0, 249, 250, 3, 96,
		48, 0, 250, 251, 5, 68, 0, 0, 251, 252, 3, 90, 45, 0, 252, 29, 1, 0, 0,
		0, 253, 254, 5, 15, 0, 0, 254, 255, 5, 68, 0, 0, 255, 256, 3, 90, 45, 0,
		256, 31, 1, 0, 0, 0, 257, 258, 5, 10, 0, 0, 258, 259, 5, 42, 0, 0, 259,
		262, 5, 24, 0, 0, 260, 263, 3, 96, 48, 0, 261, 263, 3, 34, 17, 0, 262,
		260, 1, 0, 0, 0, 262, 261, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 265,
		3, 88, 44, 0, 265, 33, 1, 0, 0, 0, 266, 267, 3, 96, 48, 0, 267, 268, 5,
		66, 0, 0, 268, 269, 5, 66, 0, 0, 269, 270, 5, 66, 0, 0, 270, 271, 3, 96,
		48, 0, 271, 35, 1, 0, 0, 0, 272, 273, 5, 11, 0, 0, 273, 274, 3, 96, 48,
		0, 274, 275, 3, 88, 44, 0, 275, 37, 1, 0, 0, 0, 276, 277, 5, 12, 0, 0,
		277, 278, 3, 96, 48, 0, 278, 279, 5, 9, 0, 0, 279, 280, 3, 88, 44, 0, 280,
		39, 1, 0, 0, 0, 281, 282, 5, 42, 0, 0, 282, 283, 5, 54, 0, 0, 283, 284,
		3, 96, 48, 0, 284, 41, 1, 0, 0, 0, 285, 286, 5, 42, 0, 0, 286, 287, 7,
		1, 0, 0, 287, 288, 3, 96, 48, 0, 288, 43, 1, 0, 0, 0, 289, 290, 5, 6, 0,
		0, 290, 291, 5, 42, 0, 0, 291, 292, 5, 68, 0, 0, 292, 293, 5, 64, 0, 0,
		293, 294, 3, 94, 47, 0, 294, 295, 5, 65, 0, 0, 295, 296, 5, 54, 0, 0, 296,
		297, 3, 46, 23, 0, 297, 45, 1, 0, 0, 0, 298, 300, 5, 64, 0, 0, 299, 301,
		3, 48, 24, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 302, 1,
		0, 0, 0, 302, 306, 5, 65, 0, 0, 303, 306, 3, 50, 25, 0, 304, 306, 5, 42,
		0, 0, 305, 298, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 304, 1, 0, 0, 0,
		306, 47, 1, 0, 0, 0, 307, 308, 3, 96, 48, 0, 308, 309, 5, 67, 0, 0, 309,
		310, 3, 48, 24, 0, 310, 313, 1, 0, 0, 0, 311, 313, 3, 96, 48, 0, 312, 307,
		1, 0, 0, 0, 312, 311, 1, 0, 0, 0, 313, 49, 1, 0, 0, 0, 314, 315, 5, 64,
		0, 0, 315, 316, 3, 94, 47, 0, 316, 317, 5, 65, 0, 0, 317, 318, 5, 60, 0,
		0, 318, 319, 5, 31, 0, 0, 319, 320, 5, 68, 0, 0, 320, 321, 3, 96, 48, 0,
		321, 322, 5, 67, 0, 0, 322, 323, 5, 30, 0, 0, 323, 324, 5, 68, 0, 0, 324,
		325, 3, 96, 48, 0, 325, 326, 5, 61, 0, 0, 326, 51, 1, 0, 0, 0, 327, 328,
		5, 42, 0, 0, 328, 329, 5, 66, 0, 0, 329, 330, 5, 25, 0, 0, 330, 331, 5,
		60, 0, 0, 331, 332, 3, 96, 48, 0, 332, 333, 5, 61, 0, 0, 333, 349, 1, 0,
		0, 0, 334, 335, 5, 42, 0, 0, 335, 336, 5, 66, 0, 0, 336, 337, 5, 26, 0,
		0, 337, 338, 5, 60, 0, 0, 338, 349, 5, 61, 0, 0, 339, 340, 5, 42, 0, 0,
		340, 341, 5, 66, 0, 0, 341, 342, 5, 27, 0, 0, 342, 343, 5, 60, 0, 0, 343,
		344, 5, 28, 0, 0, 344, 345, 5, 68, 0, 0, 345, 346, 3, 96, 48, 0, 346, 347,
		5, 61, 0, 0, 347, 349, 1, 0, 0, 0, 348, 327, 1, 0, 0, 0, 348, 334, 1, 0,
		0, 0, 348, 339, 1, 0, 0, 0, 349, 53, 1, 0, 0, 0, 350, 351, 5, 6, 0, 0,
		351, 354, 5, 42, 0, 0, 352, 353, 5, 68, 0, 0, 353, 355, 3, 56, 28, 0, 354,
		352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 357,
		5, 54, 0, 0, 357, 358, 3, 58, 29, 0, 358, 55, 1, 0, 0, 0, 359, 360, 5,
		64, 0, 0, 360, 361, 3, 56, 28, 0, 361, 362, 5, 65, 0, 0, 362, 368, 1, 0,
		0, 0, 363, 364, 5, 64, 0, 0, 364, 365, 3, 94, 47, 0, 365, 366, 5, 65, 0,
		0, 366, 368, 1, 0, 0, 0, 367, 359, 1, 0, 0, 0, 367, 363, 1, 0, 0, 0, 368,
		57, 1, 0, 0, 0, 369, 372, 3, 60, 30, 0, 370, 372, 3, 64, 32, 0, 371, 369,
		1, 0, 0, 0, 371, 370, 1, 0, 0, 0, 372, 59, 1, 0, 0, 0, 373, 374, 5, 64,
		0, 0, 374, 375, 3, 62, 31, 0, 375, 376, 5, 65, 0, 0, 376, 61, 1, 0, 0,
		0, 377, 378, 6, 31, -1, 0, 378, 381, 3, 60, 30, 0, 379, 381, 3, 48, 24,
		0, 380, 377, 1, 0, 0, 0, 380, 379, 1, 0, 0, 0, 381, 387, 1, 0, 0, 0, 382,
		383, 10, 3, 0, 0, 383, 384, 5, 67, 0, 0, 384, 386, 3, 60, 30, 0, 385, 382,
		1, 0, 0, 0, 386, 389, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1, 0,
		0, 0, 388, 63, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 390, 391, 3, 56, 28, 0,
		391, 392, 5, 60, 0, 0, 392, 393, 5, 31, 0, 0, 393, 394, 5, 68, 0, 0, 394,
		395, 3, 64, 32, 0, 395, 396, 5, 67, 0, 0, 396, 397, 5, 30, 0, 0, 397, 398,
		5, 68, 0, 0, 398, 399, 3, 96, 48, 0, 399, 400, 5, 61, 0, 0, 400, 413, 1,
		0, 0, 0, 401, 402, 3, 56, 28, 0, 402, 403, 5, 60, 0, 0, 403, 404, 5, 31,
		0, 0, 404, 405, 5, 68, 0, 0, 405, 406, 3, 96, 48, 0, 406, 407, 5, 67, 0,
		0, 407, 408, 5, 30, 0, 0, 408, 409, 5, 68, 0, 0, 409, 410, 3, 96, 48, 0,
		410, 411, 5, 61, 0, 0, 411, 413, 1, 0, 0, 0, 412, 390, 1, 0, 0, 0, 412,
		401, 1, 0, 0, 0, 413, 65, 1, 0, 0, 0, 414, 415, 5, 42, 0, 0, 415, 416,
		3, 68, 34, 0, 416, 417, 5, 54, 0, 0, 417, 418, 3, 96, 48, 0, 418, 67, 1,
		0, 0, 0, 419, 420, 5, 64, 0, 0, 420, 421, 3, 96, 48, 0, 421, 422, 5, 65,
		0, 0, 422, 423, 3, 68, 34, 0, 423, 429, 1, 0, 0, 0, 424, 425, 5, 64, 0,
		0, 425, 426, 3, 96, 48, 0, 426, 427, 5, 65, 0, 0, 427, 429, 1, 0, 0, 0,
		428, 419, 1, 0, 0, 0, 428, 424, 1, 0, 0, 0, 429, 69, 1, 0, 0, 0, 430, 431,
		5, 32, 0, 0, 431, 432, 5, 42, 0, 0, 432, 433, 5, 62, 0, 0, 433, 434, 3,
		72, 36, 0, 434, 435, 5, 63, 0, 0, 435, 71, 1, 0, 0, 0, 436, 438, 3, 74,
		37, 0, 437, 439, 5, 69, 0, 0, 438, 437, 1, 0, 0, 0, 438, 439, 1, 0, 0,
		0, 439, 440, 1, 0, 0, 0, 440, 441, 3, 72, 36, 0, 441, 447, 1, 0, 0, 0,
		442, 444, 3, 74, 37, 0, 443, 445, 5, 69, 0, 0, 444, 443, 1, 0, 0, 0, 444,
		445, 1, 0, 0, 0, 445, 447, 1, 0, 0, 0, 446, 436, 1, 0, 0, 0, 446, 442,
		1, 0, 0, 0, 447, 73, 1, 0, 0, 0, 448, 449, 7, 2, 0, 0, 449, 452, 5, 42,
		0, 0, 450, 451, 5, 68, 0, 0, 451, 453, 3, 94, 47, 0, 452, 450, 1, 0, 0,
		0, 452, 453, 1, 0, 0, 0, 453, 456, 1, 0, 0, 0, 454, 455, 5, 54, 0, 0, 455,
		457, 3, 96, 48, 0, 456, 454, 1, 0, 0, 0, 456, 457, 1, 0, 0, 0, 457, 463,
		1, 0, 0, 0, 458, 460, 5, 33, 0, 0, 459, 458, 1, 0, 0, 0, 459, 460, 1, 0,
		0, 0, 460, 461, 1, 0, 0, 0, 461, 463, 3, 14, 7, 0, 462, 448, 1, 0, 0, 0,
		462, 459, 1, 0, 0, 0, 463, 75, 1, 0, 0, 0, 464, 465, 7, 2, 0, 0, 465, 468,
		5, 42, 0, 0, 466, 467, 5, 68, 0, 0, 467, 469, 5, 42, 0, 0, 468, 466, 1,
		0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 470, 1, 0, 0, 0, 470, 471, 5, 54, 0,
		0, 471, 472, 5, 42, 0, 0, 472, 474, 5, 60, 0, 0, 473, 475, 3, 78, 39, 0,
		474, 473, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476,
		488, 5, 61, 0, 0, 477, 478, 7, 2, 0, 0, 478, 481, 5, 42, 0, 0, 479, 480,
		5, 68, 0, 0, 480, 482, 5, 42, 0, 0, 481, 479, 1, 0, 0, 0, 481, 482, 1,
		0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 5, 54, 0, 0, 484, 485, 5, 42,
		0, 0, 485, 486, 5, 60, 0, 0, 486, 488, 5, 61, 0, 0, 487, 464, 1, 0, 0,
		0, 487, 477, 1, 0, 0, 0, 488, 77, 1, 0, 0, 0, 489, 490, 5, 42, 0, 0, 490,
		491, 5, 68, 0, 0, 491, 492, 3, 96, 48, 0, 492, 493, 5, 67, 0, 0, 493, 494,
		3, 78, 39, 0, 494, 499, 1, 0, 0, 0, 495, 496, 5, 42, 0, 0, 496, 497, 5,
		68, 0, 0, 497, 499, 3, 96, 48, 0, 498, 489, 1, 0, 0, 0, 498, 495, 1, 0,
		0, 0, 499, 79, 1, 0, 0, 0, 500, 501, 3, 82, 41, 0, 501, 502, 3, 84, 42,
		0, 502, 508, 1, 0, 0, 0, 503, 504, 3, 82, 41, 0, 504, 505, 5, 66, 0, 0,
		505, 506, 3, 6, 3, 0, 506, 508, 1, 0, 0, 0, 507, 500, 1, 0, 0, 0, 507,
		503, 1, 0, 0, 0, 508, 81, 1, 0, 0, 0, 509, 510, 5, 42, 0, 0, 510, 511,
		5, 64, 0, 0, 511, 512, 3, 96, 48, 0, 512, 513, 5, 65, 0, 0, 513, 516, 1,
		0, 0, 0, 514, 516, 5, 42, 0, 0, 515, 509, 1, 0, 0, 0, 515, 514, 1, 0, 0,
		0, 516, 83, 1, 0, 0, 0, 517, 518, 5, 66, 0, 0, 518, 519, 5, 42, 0, 0, 519,
		523, 3, 84, 42, 0, 520, 521, 5, 66, 0, 0, 521, 523, 5, 42, 0, 0, 522, 517,
		1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 523, 85, 1, 0, 0, 0, 524, 525, 5, 35,
		0, 0, 525, 527, 5, 60, 0, 0, 526, 528, 3, 96, 48, 0, 527, 526, 1, 0, 0,
		0, 527, 528, 1, 0, 0, 0, 528, 529, 1, 0, 0, 0, 529, 530, 5, 61, 0, 0, 530,
		87, 1, 0, 0, 0, 531, 533, 5, 62, 0, 0, 532, 534, 3, 90, 45, 0, 533, 532,
		1, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 536, 5, 63,
		0, 0, 536, 89, 1, 0, 0, 0, 537, 538, 3, 92, 46, 0, 538, 539, 3, 90, 45,
		0, 539, 542, 1, 0, 0, 0, 540, 542, 3, 92, 46, 0, 541, 537, 1, 0, 0, 0,
		541, 540, 1, 0, 0, 0, 542, 91, 1, 0, 0, 0, 543, 545, 3, 10, 5, 0, 544,
		546, 5, 69, 0, 0, 545, 544, 1, 0, 0, 0, 545, 546, 1, 0, 0, 0, 546, 638,
		1, 0, 0, 0, 547, 549, 3, 12, 6, 0, 548, 550, 5, 69, 0, 0, 549, 548, 1,
		0, 0, 0, 549, 550, 1, 0, 0, 0, 550, 638, 1, 0, 0, 0, 551, 638, 3, 18, 9,
		0, 552, 638, 3, 20, 10, 0, 553, 638, 3, 32, 16, 0, 554, 638, 3, 36, 18,
		0, 555, 638, 3, 38, 19, 0, 556, 557, 5, 34, 0, 0, 557, 559, 5, 66, 0, 0,
		558, 556, 1, 0, 0, 0, 558, 559, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560,
		562, 3, 40, 20, 0, 561, 563, 5, 69, 0, 0, 562, 561, 1, 0, 0, 0, 562, 563,
		1, 0, 0, 0, 563, 638, 1, 0, 0, 0, 564, 565, 5, 34, 0, 0, 565, 567, 5, 66,
		0, 0, 566, 564, 1, 0, 0, 0, 566, 567, 1, 0, 0, 0, 567, 568, 1, 0, 0, 0,
		568, 570, 3, 42, 21, 0, 569, 571, 5, 69, 0, 0, 570, 569, 1, 0, 0, 0, 570,
		571, 1, 0, 0, 0, 571, 638, 1, 0, 0, 0, 572, 574, 3, 44, 22, 0, 573, 575,
		5, 69, 0, 0, 574, 573, 1, 0, 0, 0, 574, 575, 1, 0, 0, 0, 575, 638, 1, 0,
		0, 0, 576, 578, 3, 52, 26, 0, 577, 579, 5, 69, 0, 0, 578, 577, 1, 0, 0,
		0, 578, 579, 1, 0, 0, 0, 579, 638, 1, 0, 0, 0, 580, 581, 5, 34, 0, 0, 581,
		583, 5, 66, 0, 0, 582, 580, 1, 0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 584,
		1, 0, 0, 0, 584, 586, 3, 66, 33, 0, 585, 587, 5, 69, 0, 0, 586, 585, 1,
		0, 0, 0, 586, 587, 1, 0, 0, 0, 587, 638, 1, 0, 0, 0, 588, 590, 3, 54, 27,
		0, 589, 591, 5, 69, 0, 0, 590, 589, 1, 0, 0, 0, 590, 591, 1, 0, 0, 0, 591,
		638, 1, 0, 0, 0, 592, 594, 3, 70, 35, 0, 593, 595, 5, 69, 0, 0, 594, 593,
		1, 0, 0, 0, 594, 595, 1, 0, 0, 0, 595, 638, 1, 0, 0, 0, 596, 598, 3, 76,
		38, 0, 597, 599, 5, 69, 0, 0, 598, 597, 1, 0, 0, 0, 598, 599, 1, 0, 0,
		0, 599, 638, 1, 0, 0, 0, 600, 601, 5, 34, 0, 0, 601, 603, 5, 66, 0, 0,
		602, 600, 1, 0, 0, 0, 602, 603, 1, 0, 0, 0, 603, 604, 1, 0, 0, 0, 604,
		606, 3, 80, 40, 0, 605, 607, 5, 69, 0, 0, 606, 605, 1, 0, 0, 0, 606, 607,
		1, 0, 0, 0, 607, 638, 1, 0, 0, 0, 608, 609, 5, 34, 0, 0, 609, 611, 5, 66,
		0, 0, 610, 608, 1, 0, 0, 0, 610, 611, 1, 0, 0, 0, 611, 612, 1, 0, 0, 0,
		612, 614, 3, 6, 3, 0, 613, 615, 5, 69, 0, 0, 614, 613, 1, 0, 0, 0, 614,
		615, 1, 0, 0, 0, 615, 638, 1, 0, 0, 0, 616, 618, 3, 86, 43, 0, 617, 619,
		5, 69, 0, 0, 618, 617, 1, 0, 0, 0, 618, 619, 1, 0, 0, 0, 619, 638, 1, 0,
		0, 0, 620, 621, 5, 18, 0, 0, 621, 623, 3, 96, 48, 0, 622, 624, 5, 69, 0,
		0, 623, 622, 1, 0, 0, 0, 623, 624, 1, 0, 0, 0, 624, 638, 1, 0, 0, 0, 625,
		627, 5, 18, 0, 0, 626, 628, 5, 69, 0, 0, 627, 626, 1, 0, 0, 0, 627, 628,
		1, 0, 0, 0, 628, 638, 1, 0, 0, 0, 629, 631, 5, 17, 0, 0, 630, 632, 5, 69,
		0, 0, 631, 630, 1, 0, 0, 0, 631, 632, 1, 0, 0, 0, 632, 638, 1, 0, 0, 0,
		633, 635, 5, 16, 0, 0, 634, 636, 5, 69, 0, 0, 635, 634, 1, 0, 0, 0, 635,
		636, 1, 0, 0, 0, 636, 638, 1, 0, 0, 0, 637, 543, 1, 0, 0, 0, 637, 547,
		1, 0, 0, 0, 637, 551, 1, 0, 0, 0, 637, 552, 1, 0, 0, 0, 637, 553, 1, 0,
		0, 0, 637, 554, 1, 0, 0, 0, 637, 555, 1, 0, 0, 0, 637, 558, 1, 0, 0, 0,
		637, 566, 1, 0, 0, 0, 637, 572, 1, 0, 0, 0, 637, 576, 1, 0, 0, 0, 637,
		582, 1, 0, 0, 0, 637, 588, 1, 0, 0, 0, 637, 592, 1, 0, 0, 0, 637, 596,
		1, 0, 0, 0, 637, 602, 1, 0, 0, 0, 637, 610, 1, 0, 0, 0, 637, 616, 1, 0,
		0, 0, 637, 620, 1, 0, 0, 0, 637, 625, 1, 0, 0, 0, 637, 629, 1, 0, 0, 0,
		637, 633, 1, 0, 0, 0, 638, 93, 1, 0, 0, 0, 639, 640, 7, 3, 0, 0, 640, 95,
		1, 0, 0, 0, 641, 642, 6, 48, -1, 0, 642, 643, 5, 46, 0, 0, 643, 696, 3,
		96, 48, 26, 644, 645, 5, 59, 0, 0, 645, 696, 3, 96, 48, 20, 646, 647, 5,
		1, 0, 0, 647, 648, 5, 60, 0, 0, 648, 649, 3, 96, 48, 0, 649, 650, 5, 61,
		0, 0, 650, 696, 1, 0, 0, 0, 651, 652, 5, 2, 0, 0, 652, 653, 5, 60, 0, 0,
		653, 654, 3, 96, 48, 0, 654, 655, 5, 61, 0, 0, 655, 696, 1, 0, 0, 0, 656,
		657, 5, 3, 0, 0, 657, 658, 5, 60, 0, 0, 658, 659, 3, 96, 48, 0, 659, 660,
		5, 61, 0, 0, 660, 696, 1, 0, 0, 0, 661, 662, 5, 42, 0, 0, 662, 696, 3,
		68, 34, 0, 663, 664, 5, 42, 0, 0, 664, 665, 5, 66, 0, 0, 665, 696, 5, 29,
		0, 0, 666, 667, 5, 42, 0, 0, 667, 668, 5, 66, 0, 0, 668, 696, 5, 30, 0,
		0, 669, 670, 5, 34, 0, 0, 670, 672, 5, 66, 0, 0, 671, 669, 1, 0, 0, 0,
		671, 672, 1, 0, 0, 0, 672, 673, 1, 0, 0, 0, 673, 696, 3, 80, 40, 0, 674,
		675, 5, 34, 0, 0, 675, 677, 5, 66, 0, 0, 676, 674, 1, 0, 0, 0, 676, 677,
		1, 0, 0, 0, 677, 678, 1, 0, 0, 0, 678, 696, 3, 6, 3, 0, 679, 696, 5, 21,
		0, 0, 680, 681, 5, 34, 0, 0, 681, 683, 5, 66, 0, 0, 682, 680, 1, 0, 0,
		0, 682, 683, 1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684, 696, 5, 42, 0, 0, 685,
		696, 5, 39, 0, 0, 686, 696, 5, 38, 0, 0, 687, 696, 5, 40, 0, 0, 688, 696,
		5, 41, 0, 0, 689, 696, 5, 19, 0, 0, 690, 696, 5, 20, 0, 0, 691, 692, 5,
		60, 0, 0, 692, 693, 3, 96, 48, 0, 693, 694, 5, 61, 0, 0, 694, 696, 1, 0,
		0, 0, 695, 641, 1, 0, 0, 0, 695, 644, 1, 0, 0, 0, 695, 646, 1, 0, 0, 0,
		695, 651, 1, 0, 0, 0, 695, 656, 1, 0, 0, 0, 695, 661, 1, 0, 0, 0, 695,
		663, 1, 0, 0, 0, 695, 666, 1, 0, 0, 0, 695, 671, 1, 0, 0, 0, 695, 676,
		1, 0, 0, 0, 695, 679, 1, 0, 0, 0, 695, 682, 1, 0, 0, 0, 695, 685, 1, 0,
		0, 0, 695, 686, 1, 0, 0, 0, 695, 687, 1, 0, 0, 0, 695, 688, 1, 0, 0, 0,
		695, 689, 1, 0, 0, 0, 695, 690, 1, 0, 0, 0, 695, 691, 1, 0, 0, 0, 696,
		720, 1, 0, 0, 0, 697, 698, 10, 25, 0, 0, 698, 699, 7, 4, 0, 0, 699, 719,
		3, 96, 48, 26, 700, 701, 10, 24, 0, 0, 701, 702, 7, 5, 0, 0, 702, 719,
		3, 96, 48, 25, 703, 704, 10, 23, 0, 0, 704, 705, 7, 6, 0, 0, 705, 719,
		3, 96, 48, 24, 706, 707, 10, 22, 0, 0, 707, 708, 7, 7, 0, 0, 708, 719,
		3, 96, 48, 23, 709, 710, 10, 21, 0, 0, 710, 711, 7, 8, 0, 0, 711, 719,
		3, 96, 48, 22, 712, 713, 10, 19, 0, 0, 713, 714, 5, 57, 0, 0, 714, 719,
		3, 96, 48, 20, 715, 716, 10, 18, 0, 0, 716, 717, 5, 58, 0, 0, 717, 719,
		3, 96, 48, 19, 718, 697, 1, 0, 0, 0, 718, 700, 1, 0, 0, 0, 718, 703, 1,
		0, 0, 0, 718, 706, 1, 0, 0, 0, 718, 709, 1, 0, 0, 0, 718, 712, 1, 0, 0,
		0, 718, 715, 1, 0, 0, 0, 719, 722, 1, 0, 0, 0, 720, 718, 1, 0, 0, 0, 720,
		721, 1, 0, 0, 0, 721, 97, 1, 0, 0, 0, 722, 720, 1, 0, 0, 0, 80, 99, 107,
		111, 116, 122, 125, 133, 136, 139, 158, 171, 177, 182, 187, 192, 199, 204,
		207, 225, 240, 246, 262, 300, 305, 312, 348, 354, 367, 371, 380, 387, 412,
		428, 438, 444, 446, 452, 456, 459, 462, 468, 474, 481, 487, 498, 507, 515,
		522, 527, 533, 541, 545, 549, 558, 562, 566, 570, 574, 578, 582, 586, 590,
		594, 598, 602, 606, 610, 614, 618, 623, 627, 631, 635, 637, 671, 676, 682,
		695, 718, 720,
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
	ParserParserRW_struct     = 32
	ParserParserRW_mutating   = 33
	ParserParserRW_self       = 34
	ParserParserRW_print      = 35
	ParserParserTK_prompt     = 36
	ParserParserTK_under      = 37
	ParserParserTK_char       = 38
	ParserParserTK_string     = 39
	ParserParserTK_int        = 40
	ParserParserTK_float      = 41
	ParserParserTK_id         = 42
	ParserParserTK_add        = 43
	ParserParserTK_sub        = 44
	ParserParserTK_plus       = 45
	ParserParserTK_minus      = 46
	ParserParserTK_mult       = 47
	ParserParserTK_div        = 48
	ParserParserTK_mod        = 49
	ParserParserTK_equequ     = 50
	ParserParserTK_notequ     = 51
	ParserParserTK_lessequ    = 52
	ParserParserTK_moreequ    = 53
	ParserParserTK_equ        = 54
	ParserParserTK_less       = 55
	ParserParserTK_more       = 56
	ParserParserTK_and        = 57
	ParserParserTK_or         = 58
	ParserParserTK_not        = 59
	ParserParserTK_lpar       = 60
	ParserParserTK_rpar       = 61
	ParserParserTK_lbrc       = 62
	ParserParserTK_rbrc       = 63
	ParserParserTK_lbrk       = 64
	ParserParserTK_rbrk       = 65
	ParserParserTK_dot        = 66
	ParserParserTK_comma      = 67
	ParserParserTK_colon      = 68
	ParserParserTK_semicolon  = 69
	ParserParserTK_question   = 70
	ParserParserTK_amp        = 71
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
	ParserParserRULE_decmatrix     = 27
	ParserParserRULE_typematrix    = 28
	ParserParserRULE_defmatrix     = 29
	ParserParserRULE_listvector    = 30
	ParserParserRULE_listvector2   = 31
	ParserParserRULE_simplematrix  = 32
	ParserParserRULE_reasignvector = 33
	ParserParserRULE_dims          = 34
	ParserParserRULE_defstruct     = 35
	ParserParserRULE_listattribs   = 36
	ParserParserRULE_attrib        = 37
	ParserParserRULE_decstruct     = 38
	ParserParserRULE_listdupla     = 39
	ParserParserRULE_useattribs    = 40
	ParserParserRULE_obj           = 41
	ParserParserRULE_useattribs1   = 42
	ParserParserRULE_print         = 43
	ParserParserRULE_env           = 44
	ParserParserRULE_instructions  = 45
	ParserParserRULE_instruction   = 46
	ParserParserRULE_type          = 47
	ParserParserRULE_exp           = 48
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
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4453885754816) != 0 {
		{
			p.SetState(98)
			p.Instsglobal()
		}

	}
	{
		p.SetState(101)
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
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Instglobal()
		}
		{
			p.SetState(104)
			p.Instsglobal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_var, ParserParserRW_let, ParserParserRW_if, ParserParserRW_for, ParserParserRW_while, ParserParserRW_guard, ParserParserRW_switch, ParserParserRW_break, ParserParserRW_continue, ParserParserRW_return, ParserParserRW_struct, ParserParserRW_self, ParserParserRW_print, ParserParserTK_id:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Instruction()
		}

	case ParserParserRW_func:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
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
		p.SetState(113)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1729461164053102606) != 0) || _la == ParserParserTK_amp {
		{
			p.SetState(115)
			p.Listargs()
		}

	}
	{
		p.SetState(118)
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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(122)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(120)
				p.Match(ParserParserTK_id)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(121)
				p.Match(ParserParserTK_colon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_amp {
			{
				p.SetState(124)
				p.Match(ParserParserTK_amp)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(127)
			p.exp(0)
		}
		{
			p.SetState(128)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Listargs()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(133)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
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

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_amp {
			{
				p.SetState(135)
				p.Match(ParserParserTK_amp)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(138)
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
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.Match(ParserParserRW_var)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(142)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(144)
			p.Type_()
		}
		{
			p.SetState(145)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(146)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(ParserParserRW_var)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(150)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Type_()
		}
		{
			p.SetState(152)
			p.Match(ParserParserTK_question)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(154)
			p.Match(ParserParserRW_var)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(157)
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
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(160)
			p.Match(ParserParserRW_let)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(161)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(163)
			p.Type_()
		}
		{
			p.SetState(164)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.exp(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(ParserParserRW_let)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(170)
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
		p.SetState(173)
		p.Match(ParserParserRW_func)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserParserTK_under || _la == ParserParserTK_id {
		{
			p.SetState(176)
			p.Listparams()
		}

	}
	{
		p.SetState(179)
		p.Match(ParserParserTK_rpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserParserTK_prompt {
		{
			p.SetState(180)
			p.Match(ParserParserTK_prompt)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Type_()
		}

	}
	{
		p.SetState(184)
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

	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(187)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(186)
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
			p.SetState(189)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(190)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_inout {
			{
				p.SetState(191)
				p.Match(ParserParserRW_inout)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(194)
			p.Type_()
		}
		{
			p.SetState(195)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(196)
			p.Listparams()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(198)
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
			p.SetState(201)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(202)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_inout {
			{
				p.SetState(203)
				p.Match(ParserParserRW_inout)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(206)
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
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Match(ParserParserRW_if)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(210)
			p.exp(0)
		}
		{
			p.SetState(211)
			p.Env()
		}
		{
			p.SetState(212)
			p.Match(ParserParserRW_else)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Ifstruct()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(ParserParserRW_if)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(216)
			p.exp(0)
		}
		{
			p.SetState(217)
			p.Env()
		}
		{
			p.SetState(218)
			p.Match(ParserParserRW_else)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Env()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(221)
			p.Match(ParserParserRW_if)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(222)
			p.exp(0)
		}
		{
			p.SetState(223)
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
		p.SetState(227)
		p.Match(ParserParserRW_switch)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(228)
		p.exp(0)
	}
	{
		p.SetState(229)
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
		p.SetState(231)
		p.Match(ParserParserTK_lbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Casesdefault()
	}
	{
		p.SetState(233)
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
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(235)
			p.Cases()
		}
		{
			p.SetState(236)
			p.Default_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Cases()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
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
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Case_()
		}
		{
			p.SetState(243)
			p.Cases()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)
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
		p.SetState(248)
		p.Match(ParserParserRW_case)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.exp(0)
	}
	{
		p.SetState(250)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
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
		p.SetState(253)
		p.Match(ParserParserRW_default)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(254)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(255)
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
		p.SetState(257)
		p.Match(ParserParserRW_for)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(259)
		p.Match(ParserParserRW_in)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(260)
			p.exp(0)
		}

	case 2:
		{
			p.SetState(261)
			p.Range_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(264)
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
		p.SetState(266)
		p.exp(0)
	}
	{
		p.SetState(267)
		p.Match(ParserParserTK_dot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
		p.Match(ParserParserTK_dot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.Match(ParserParserTK_dot)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(270)
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
		p.SetState(272)
		p.Match(ParserParserRW_while)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.exp(0)
	}
	{
		p.SetState(274)
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
		p.SetState(276)
		p.Match(ParserParserRW_guard)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.exp(0)
	}
	{
		p.SetState(278)
		p.Match(ParserParserRW_else)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(279)
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
		p.SetState(281)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
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
		p.SetState(285)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ParserParserTK_add || _la == ParserParserTK_sub) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(287)
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
		p.SetState(289)
		p.Match(ParserParserRW_var)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(ParserParserTK_lbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
		p.Type_()
	}
	{
		p.SetState(294)
		p.Match(ParserParserTK_rbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(295)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1729461164053102606) != 0 {
			{
				p.SetState(299)
				p.Listexp()
			}

		}
		{
			p.SetState(302)
			p.Match(ParserParserTK_rbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(303)
			p.Simplevec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(304)
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
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.exp(0)
		}
		{
			p.SetState(308)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.Listexp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)
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
		p.SetState(314)
		p.Match(ParserParserTK_lbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)
		p.Type_()
	}
	{
		p.SetState(316)
		p.Match(ParserParserTK_rbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(317)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(318)
		p.Match(ParserParserRW_repeating)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(319)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(320)
		p.exp(0)
	}
	{
		p.SetState(321)
		p.Match(ParserParserTK_comma)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(ParserParserRW_count)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(323)
		p.Match(ParserParserTK_colon)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.exp(0)
	}
	{
		p.SetState(325)
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
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(329)
			p.Match(ParserParserRW_append)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(331)
			p.exp(0)
		}
		{
			p.SetState(332)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(334)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(ParserParserRW_removeLast)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(338)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(339)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.Match(ParserParserRW_remove)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(342)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			p.Match(ParserParserRW_at)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(344)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(345)
			p.exp(0)
		}
		{
			p.SetState(346)
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

// IDecmatrixContext is an interface to support dynamic dispatch.
type IDecmatrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_var() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	Defmatrix() IDefmatrixContext
	TK_colon() antlr.TerminalNode
	Typematrix() ITypematrixContext

	// IsDecmatrixContext differentiates from other interfaces.
	IsDecmatrixContext()
}

type DecmatrixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecmatrixContext() *DecmatrixContext {
	var p = new(DecmatrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decmatrix
	return p
}

func InitEmptyDecmatrixContext(p *DecmatrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decmatrix
}

func (*DecmatrixContext) IsDecmatrixContext() {}

func NewDecmatrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecmatrixContext {
	var p = new(DecmatrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_decmatrix

	return p
}

func (s *DecmatrixContext) GetParser() antlr.Parser { return s.parser }

func (s *DecmatrixContext) RW_var() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_var, 0)
}

func (s *DecmatrixContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DecmatrixContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *DecmatrixContext) Defmatrix() IDefmatrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefmatrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefmatrixContext)
}

func (s *DecmatrixContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *DecmatrixContext) Typematrix() ITypematrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypematrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypematrixContext)
}

func (s *DecmatrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecmatrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecmatrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDecmatrix(s)
	}
}

func (s *DecmatrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDecmatrix(s)
	}
}

func (p *ParserParser) Decmatrix() (localctx IDecmatrixContext) {
	localctx = NewDecmatrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ParserParserRULE_decmatrix)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(ParserParserRW_var)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParserParserTK_colon {
		{
			p.SetState(352)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(353)
			p.Typematrix()
		}

	}
	{
		p.SetState(356)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(357)
		p.Defmatrix()
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

// ITypematrixContext is an interface to support dynamic dispatch.
type ITypematrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrk() antlr.TerminalNode
	Typematrix() ITypematrixContext
	TK_rbrk() antlr.TerminalNode
	Type_() ITypeContext

	// IsTypematrixContext differentiates from other interfaces.
	IsTypematrixContext()
}

type TypematrixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypematrixContext() *TypematrixContext {
	var p = new(TypematrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_typematrix
	return p
}

func InitEmptyTypematrixContext(p *TypematrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_typematrix
}

func (*TypematrixContext) IsTypematrixContext() {}

func NewTypematrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypematrixContext {
	var p = new(TypematrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_typematrix

	return p
}

func (s *TypematrixContext) GetParser() antlr.Parser { return s.parser }

func (s *TypematrixContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *TypematrixContext) Typematrix() ITypematrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypematrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypematrixContext)
}

func (s *TypematrixContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *TypematrixContext) Type_() ITypeContext {
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

func (s *TypematrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypematrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypematrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterTypematrix(s)
	}
}

func (s *TypematrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitTypematrix(s)
	}
}

func (p *ParserParser) Typematrix() (localctx ITypematrixContext) {
	localctx = NewTypematrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ParserParserRULE_typematrix)
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(359)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(360)
			p.Typematrix()
		}
		{
			p.SetState(361)
			p.Match(ParserParserTK_rbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(363)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(364)
			p.Type_()
		}
		{
			p.SetState(365)
			p.Match(ParserParserTK_rbrk)
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

// IDefmatrixContext is an interface to support dynamic dispatch.
type IDefmatrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Listvector() IListvectorContext
	Simplematrix() ISimplematrixContext

	// IsDefmatrixContext differentiates from other interfaces.
	IsDefmatrixContext()
}

type DefmatrixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefmatrixContext() *DefmatrixContext {
	var p = new(DefmatrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_defmatrix
	return p
}

func InitEmptyDefmatrixContext(p *DefmatrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_defmatrix
}

func (*DefmatrixContext) IsDefmatrixContext() {}

func NewDefmatrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefmatrixContext {
	var p = new(DefmatrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_defmatrix

	return p
}

func (s *DefmatrixContext) GetParser() antlr.Parser { return s.parser }

func (s *DefmatrixContext) Listvector() IListvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListvectorContext)
}

func (s *DefmatrixContext) Simplematrix() ISimplematrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplematrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplematrixContext)
}

func (s *DefmatrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefmatrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefmatrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDefmatrix(s)
	}
}

func (s *DefmatrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDefmatrix(s)
	}
}

func (p *ParserParser) Defmatrix() (localctx IDefmatrixContext) {
	localctx = NewDefmatrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ParserParserRULE_defmatrix)
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.Listvector()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.Simplematrix()
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

// IListvectorContext is an interface to support dynamic dispatch.
type IListvectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrk() antlr.TerminalNode
	Listvector2() IListvector2Context
	TK_rbrk() antlr.TerminalNode

	// IsListvectorContext differentiates from other interfaces.
	IsListvectorContext()
}

type ListvectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListvectorContext() *ListvectorContext {
	var p = new(ListvectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listvector
	return p
}

func InitEmptyListvectorContext(p *ListvectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listvector
}

func (*ListvectorContext) IsListvectorContext() {}

func NewListvectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListvectorContext {
	var p = new(ListvectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listvector

	return p
}

func (s *ListvectorContext) GetParser() antlr.Parser { return s.parser }

func (s *ListvectorContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *ListvectorContext) Listvector2() IListvector2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListvector2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListvector2Context)
}

func (s *ListvectorContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *ListvectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListvectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListvectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListvector(s)
	}
}

func (s *ListvectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListvector(s)
	}
}

func (p *ParserParser) Listvector() (localctx IListvectorContext) {
	localctx = NewListvectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ParserParserRULE_listvector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(ParserParserTK_lbrk)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(374)
		p.listvector2(0)
	}
	{
		p.SetState(375)
		p.Match(ParserParserTK_rbrk)
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

// IListvector2Context is an interface to support dynamic dispatch.
type IListvector2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Listvector() IListvectorContext
	Listexp() IListexpContext
	Listvector2() IListvector2Context
	TK_comma() antlr.TerminalNode

	// IsListvector2Context differentiates from other interfaces.
	IsListvector2Context()
}

type Listvector2Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListvector2Context() *Listvector2Context {
	var p = new(Listvector2Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listvector2
	return p
}

func InitEmptyListvector2Context(p *Listvector2Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listvector2
}

func (*Listvector2Context) IsListvector2Context() {}

func NewListvector2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Listvector2Context {
	var p = new(Listvector2Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listvector2

	return p
}

func (s *Listvector2Context) GetParser() antlr.Parser { return s.parser }

func (s *Listvector2Context) Listvector() IListvectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListvectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListvectorContext)
}

func (s *Listvector2Context) Listexp() IListexpContext {
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

func (s *Listvector2Context) Listvector2() IListvector2Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListvector2Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListvector2Context)
}

func (s *Listvector2Context) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *Listvector2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listvector2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Listvector2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListvector2(s)
	}
}

func (s *Listvector2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListvector2(s)
	}
}

func (p *ParserParser) Listvector2() (localctx IListvector2Context) {
	return p.listvector2(0)
}

func (p *ParserParser) listvector2(_p int) (localctx IListvector2Context) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewListvector2Context(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListvector2Context = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ParserParserRULE_listvector2, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(380)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserTK_lbrk:
		{
			p.SetState(378)
			p.Listvector()
		}

	case ParserParserRW_Int, ParserParserRW_Float, ParserParserRW_String, ParserParserRW_true, ParserParserRW_false, ParserParserRW_nil, ParserParserRW_self, ParserParserTK_char, ParserParserTK_string, ParserParserTK_int, ParserParserTK_float, ParserParserTK_id, ParserParserTK_minus, ParserParserTK_not, ParserParserTK_lpar:
		{
			p.SetState(379)
			p.Listexp()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListvector2Context(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_listvector2)
			p.SetState(382)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(383)
				p.Match(ParserParserTK_comma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(384)
				p.Listvector()
			}

		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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

// ISimplematrixContext is an interface to support dynamic dispatch.
type ISimplematrixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Typematrix() ITypematrixContext
	TK_lpar() antlr.TerminalNode
	RW_repeating() antlr.TerminalNode
	AllTK_colon() []antlr.TerminalNode
	TK_colon(i int) antlr.TerminalNode
	Simplematrix() ISimplematrixContext
	TK_comma() antlr.TerminalNode
	RW_count() antlr.TerminalNode
	AllExp() []IExpContext
	Exp(i int) IExpContext
	TK_rpar() antlr.TerminalNode

	// IsSimplematrixContext differentiates from other interfaces.
	IsSimplematrixContext()
}

type SimplematrixContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimplematrixContext() *SimplematrixContext {
	var p = new(SimplematrixContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_simplematrix
	return p
}

func InitEmptySimplematrixContext(p *SimplematrixContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_simplematrix
}

func (*SimplematrixContext) IsSimplematrixContext() {}

func NewSimplematrixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimplematrixContext {
	var p = new(SimplematrixContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_simplematrix

	return p
}

func (s *SimplematrixContext) GetParser() antlr.Parser { return s.parser }

func (s *SimplematrixContext) Typematrix() ITypematrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypematrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypematrixContext)
}

func (s *SimplematrixContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *SimplematrixContext) RW_repeating() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_repeating, 0)
}

func (s *SimplematrixContext) AllTK_colon() []antlr.TerminalNode {
	return s.GetTokens(ParserParserTK_colon)
}

func (s *SimplematrixContext) TK_colon(i int) antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, i)
}

func (s *SimplematrixContext) Simplematrix() ISimplematrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimplematrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimplematrixContext)
}

func (s *SimplematrixContext) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *SimplematrixContext) RW_count() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_count, 0)
}

func (s *SimplematrixContext) AllExp() []IExpContext {
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

func (s *SimplematrixContext) Exp(i int) IExpContext {
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

func (s *SimplematrixContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *SimplematrixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimplematrixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimplematrixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterSimplematrix(s)
	}
}

func (s *SimplematrixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitSimplematrix(s)
	}
}

func (p *ParserParser) Simplematrix() (localctx ISimplematrixContext) {
	localctx = NewSimplematrixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ParserParserRULE_simplematrix)
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(390)
			p.Typematrix()
		}
		{
			p.SetState(391)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(ParserParserRW_repeating)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Simplematrix()
		}
		{
			p.SetState(395)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(396)
			p.Match(ParserParserRW_count)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(398)
			p.exp(0)
		}
		{
			p.SetState(399)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			p.Typematrix()
		}
		{
			p.SetState(402)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(403)
			p.Match(ParserParserRW_repeating)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(404)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.exp(0)
		}
		{
			p.SetState(406)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(407)
			p.Match(ParserParserRW_count)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(408)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			p.exp(0)
		}
		{
			p.SetState(410)
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
	Dims() IDimsContext
	TK_equ() antlr.TerminalNode
	Exp() IExpContext

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

func (s *ReasignvectorContext) Dims() IDimsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimsContext)
}

func (s *ReasignvectorContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *ReasignvectorContext) Exp() IExpContext {
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
	p.EnterRule(localctx, 66, ParserParserRULE_reasignvector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(414)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(415)
		p.Dims()
	}
	{
		p.SetState(416)
		p.Match(ParserParserTK_equ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(417)
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

// IDimsContext is an interface to support dynamic dispatch.
type IDimsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_lbrk() antlr.TerminalNode
	Exp() IExpContext
	TK_rbrk() antlr.TerminalNode
	Dims() IDimsContext

	// IsDimsContext differentiates from other interfaces.
	IsDimsContext()
}

type DimsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDimsContext() *DimsContext {
	var p = new(DimsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_dims
	return p
}

func InitEmptyDimsContext(p *DimsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_dims
}

func (*DimsContext) IsDimsContext() {}

func NewDimsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimsContext {
	var p = new(DimsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_dims

	return p
}

func (s *DimsContext) GetParser() antlr.Parser { return s.parser }

func (s *DimsContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *DimsContext) Exp() IExpContext {
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

func (s *DimsContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *DimsContext) Dims() IDimsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimsContext)
}

func (s *DimsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDims(s)
	}
}

func (s *DimsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDims(s)
	}
}

func (p *ParserParser) Dims() (localctx IDimsContext) {
	localctx = NewDimsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ParserParserRULE_dims)
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(419)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(420)
			p.exp(0)
		}
		{
			p.SetState(421)
			p.Match(ParserParserTK_rbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(422)
			p.Dims()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(424)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(425)
			p.exp(0)
		}
		{
			p.SetState(426)
			p.Match(ParserParserTK_rbrk)
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

// IDefstructContext is an interface to support dynamic dispatch.
type IDefstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_struct() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	TK_lbrc() antlr.TerminalNode
	Listattribs() IListattribsContext
	TK_rbrc() antlr.TerminalNode

	// IsDefstructContext differentiates from other interfaces.
	IsDefstructContext()
}

type DefstructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefstructContext() *DefstructContext {
	var p = new(DefstructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_defstruct
	return p
}

func InitEmptyDefstructContext(p *DefstructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_defstruct
}

func (*DefstructContext) IsDefstructContext() {}

func NewDefstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefstructContext {
	var p = new(DefstructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_defstruct

	return p
}

func (s *DefstructContext) GetParser() antlr.Parser { return s.parser }

func (s *DefstructContext) RW_struct() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_struct, 0)
}

func (s *DefstructContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *DefstructContext) TK_lbrc() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrc, 0)
}

func (s *DefstructContext) Listattribs() IListattribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListattribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListattribsContext)
}

func (s *DefstructContext) TK_rbrc() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrc, 0)
}

func (s *DefstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDefstruct(s)
	}
}

func (s *DefstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDefstruct(s)
	}
}

func (p *ParserParser) Defstruct() (localctx IDefstructContext) {
	localctx = NewDefstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ParserParserRULE_defstruct)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		p.Match(ParserParserRW_struct)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(431)
		p.Match(ParserParserTK_id)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(432)
		p.Match(ParserParserTK_lbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(433)
		p.Listattribs()
	}
	{
		p.SetState(434)
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

// IListattribsContext is an interface to support dynamic dispatch.
type IListattribsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Attrib() IAttribContext
	Listattribs() IListattribsContext
	TK_semicolon() antlr.TerminalNode

	// IsListattribsContext differentiates from other interfaces.
	IsListattribsContext()
}

type ListattribsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListattribsContext() *ListattribsContext {
	var p = new(ListattribsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listattribs
	return p
}

func InitEmptyListattribsContext(p *ListattribsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listattribs
}

func (*ListattribsContext) IsListattribsContext() {}

func NewListattribsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListattribsContext {
	var p = new(ListattribsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listattribs

	return p
}

func (s *ListattribsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListattribsContext) Attrib() IAttribContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttribContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *ListattribsContext) Listattribs() IListattribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListattribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListattribsContext)
}

func (s *ListattribsContext) TK_semicolon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_semicolon, 0)
}

func (s *ListattribsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListattribsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListattribsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListattribs(s)
	}
}

func (s *ListattribsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListattribs(s)
	}
}

func (p *ParserParser) Listattribs() (localctx IListattribsContext) {
	localctx = NewListattribsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ParserParserRULE_listattribs)
	var _la int

	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(436)
			p.Attrib()
		}
		p.SetState(438)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(437)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(440)
			p.Listattribs()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Attrib()
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(443)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	RW_let() antlr.TerminalNode
	RW_var() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Type_() ITypeContext
	TK_equ() antlr.TerminalNode
	Exp() IExpContext
	Declfunc() IDeclfuncContext
	RW_mutating() antlr.TerminalNode

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_attrib
	return p
}

func InitEmptyAttribContext(p *AttribContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_attrib
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *AttribContext) RW_let() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_let, 0)
}

func (s *AttribContext) RW_var() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_var, 0)
}

func (s *AttribContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *AttribContext) Type_() ITypeContext {
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

func (s *AttribContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *AttribContext) Exp() IExpContext {
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

func (s *AttribContext) Declfunc() IDeclfuncContext {
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

func (s *AttribContext) RW_mutating() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_mutating, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *ParserParser) Attrib() (localctx IAttribContext) {
	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ParserParserRULE_attrib)
	var _la int

	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParserParserRW_var, ParserParserRW_let:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ParserParserRW_var || _la == ParserParserRW_let) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(449)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_colon {
			{
				p.SetState(450)
				p.Match(ParserParserTK_colon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(451)
				p.Type_()
			}

		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_equ {
			{
				p.SetState(454)
				p.Match(ParserParserTK_equ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(455)
				p.exp(0)
			}

		}

	case ParserParserRW_func, ParserParserRW_mutating:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_mutating {
			{
				p.SetState(458)
				p.Match(ParserParserRW_mutating)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(461)
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

// IDecstructContext is an interface to support dynamic dispatch.
type IDecstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTK_id() []antlr.TerminalNode
	TK_id(i int) antlr.TerminalNode
	TK_equ() antlr.TerminalNode
	TK_lpar() antlr.TerminalNode
	TK_rpar() antlr.TerminalNode
	RW_let() antlr.TerminalNode
	RW_var() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Listdupla() IListduplaContext

	// IsDecstructContext differentiates from other interfaces.
	IsDecstructContext()
}

type DecstructContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecstructContext() *DecstructContext {
	var p = new(DecstructContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decstruct
	return p
}

func InitEmptyDecstructContext(p *DecstructContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_decstruct
}

func (*DecstructContext) IsDecstructContext() {}

func NewDecstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecstructContext {
	var p = new(DecstructContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_decstruct

	return p
}

func (s *DecstructContext) GetParser() antlr.Parser { return s.parser }

func (s *DecstructContext) AllTK_id() []antlr.TerminalNode {
	return s.GetTokens(ParserParserTK_id)
}

func (s *DecstructContext) TK_id(i int) antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, i)
}

func (s *DecstructContext) TK_equ() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_equ, 0)
}

func (s *DecstructContext) TK_lpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lpar, 0)
}

func (s *DecstructContext) TK_rpar() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rpar, 0)
}

func (s *DecstructContext) RW_let() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_let, 0)
}

func (s *DecstructContext) RW_var() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_var, 0)
}

func (s *DecstructContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *DecstructContext) Listdupla() IListduplaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListduplaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListduplaContext)
}

func (s *DecstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterDecstruct(s)
	}
}

func (s *DecstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitDecstruct(s)
	}
}

func (p *ParserParser) Decstruct() (localctx IDecstructContext) {
	localctx = NewDecstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ParserParserRULE_decstruct)
	var _la int

	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ParserParserRW_var || _la == ParserParserRW_let) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(465)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(468)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_colon {
			{
				p.SetState(466)
				p.Match(ParserParserTK_colon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(467)
				p.Match(ParserParserTK_id)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(470)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(471)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(474)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_id {
			{
				p.SetState(473)
				p.Listdupla()
			}

		}
		{
			p.SetState(476)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ParserParserRW_var || _la == ParserParserRW_let) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(478)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_colon {
			{
				p.SetState(479)
				p.Match(ParserParserTK_colon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(480)
				p.Match(ParserParserTK_id)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(483)
			p.Match(ParserParserTK_equ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(484)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(485)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(486)
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

// IListduplaContext is an interface to support dynamic dispatch.
type IListduplaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	TK_colon() antlr.TerminalNode
	Exp() IExpContext
	TK_comma() antlr.TerminalNode
	Listdupla() IListduplaContext

	// IsListduplaContext differentiates from other interfaces.
	IsListduplaContext()
}

type ListduplaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListduplaContext() *ListduplaContext {
	var p = new(ListduplaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listdupla
	return p
}

func InitEmptyListduplaContext(p *ListduplaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_listdupla
}

func (*ListduplaContext) IsListduplaContext() {}

func NewListduplaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListduplaContext {
	var p = new(ListduplaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_listdupla

	return p
}

func (s *ListduplaContext) GetParser() antlr.Parser { return s.parser }

func (s *ListduplaContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *ListduplaContext) TK_colon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_colon, 0)
}

func (s *ListduplaContext) Exp() IExpContext {
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

func (s *ListduplaContext) TK_comma() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_comma, 0)
}

func (s *ListduplaContext) Listdupla() IListduplaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListduplaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListduplaContext)
}

func (s *ListduplaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListduplaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListduplaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterListdupla(s)
	}
}

func (s *ListduplaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitListdupla(s)
	}
}

func (p *ParserParser) Listdupla() (localctx IListduplaContext) {
	localctx = NewListduplaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ParserParserRULE_listdupla)
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(491)
			p.exp(0)
		}
		{
			p.SetState(492)
			p.Match(ParserParserTK_comma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(493)
			p.Listdupla()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(495)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(496)
			p.Match(ParserParserTK_colon)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(497)
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

// IUseattribsContext is an interface to support dynamic dispatch.
type IUseattribsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Obj() IObjContext
	Useattribs1() IUseattribs1Context
	TK_dot() antlr.TerminalNode
	Callfunc() ICallfuncContext

	// IsUseattribsContext differentiates from other interfaces.
	IsUseattribsContext()
}

type UseattribsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseattribsContext() *UseattribsContext {
	var p = new(UseattribsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_useattribs
	return p
}

func InitEmptyUseattribsContext(p *UseattribsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_useattribs
}

func (*UseattribsContext) IsUseattribsContext() {}

func NewUseattribsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseattribsContext {
	var p = new(UseattribsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_useattribs

	return p
}

func (s *UseattribsContext) GetParser() antlr.Parser { return s.parser }

func (s *UseattribsContext) Obj() IObjContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjContext)
}

func (s *UseattribsContext) Useattribs1() IUseattribs1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseattribs1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseattribs1Context)
}

func (s *UseattribsContext) TK_dot() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_dot, 0)
}

func (s *UseattribsContext) Callfunc() ICallfuncContext {
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

func (s *UseattribsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseattribsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseattribsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterUseattribs(s)
	}
}

func (s *UseattribsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitUseattribs(s)
	}
}

func (p *ParserParser) Useattribs() (localctx IUseattribsContext) {
	localctx = NewUseattribsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ParserParserRULE_useattribs)
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(500)
			p.Obj()
		}
		{
			p.SetState(501)
			p.Useattribs1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(503)
			p.Obj()
		}
		{
			p.SetState(504)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(505)
			p.Callfunc()
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

// IObjContext is an interface to support dynamic dispatch.
type IObjContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_id() antlr.TerminalNode
	TK_lbrk() antlr.TerminalNode
	Exp() IExpContext
	TK_rbrk() antlr.TerminalNode

	// IsObjContext differentiates from other interfaces.
	IsObjContext()
}

type ObjContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjContext() *ObjContext {
	var p = new(ObjContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_obj
	return p
}

func InitEmptyObjContext(p *ObjContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_obj
}

func (*ObjContext) IsObjContext() {}

func NewObjContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjContext {
	var p = new(ObjContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_obj

	return p
}

func (s *ObjContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *ObjContext) TK_lbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_lbrk, 0)
}

func (s *ObjContext) Exp() IExpContext {
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

func (s *ObjContext) TK_rbrk() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_rbrk, 0)
}

func (s *ObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterObj(s)
	}
}

func (s *ObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitObj(s)
	}
}

func (p *ParserParser) Obj() (localctx IObjContext) {
	localctx = NewObjContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ParserParserRULE_obj)
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(509)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(510)
			p.Match(ParserParserTK_lbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(511)
			p.exp(0)
		}
		{
			p.SetState(512)
			p.Match(ParserParserTK_rbrk)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(514)
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

// IUseattribs1Context is an interface to support dynamic dispatch.
type IUseattribs1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TK_dot() antlr.TerminalNode
	TK_id() antlr.TerminalNode
	Useattribs1() IUseattribs1Context

	// IsUseattribs1Context differentiates from other interfaces.
	IsUseattribs1Context()
}

type Useattribs1Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseattribs1Context() *Useattribs1Context {
	var p = new(Useattribs1Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_useattribs1
	return p
}

func InitEmptyUseattribs1Context(p *Useattribs1Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParserParserRULE_useattribs1
}

func (*Useattribs1Context) IsUseattribs1Context() {}

func NewUseattribs1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Useattribs1Context {
	var p = new(Useattribs1Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParserParserRULE_useattribs1

	return p
}

func (s *Useattribs1Context) GetParser() antlr.Parser { return s.parser }

func (s *Useattribs1Context) TK_dot() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_dot, 0)
}

func (s *Useattribs1Context) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
}

func (s *Useattribs1Context) Useattribs1() IUseattribs1Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseattribs1Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseattribs1Context)
}

func (s *Useattribs1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Useattribs1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Useattribs1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.EnterUseattribs1(s)
	}
}

func (s *Useattribs1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParserListener); ok {
		listenerT.ExitUseattribs1(s)
	}
}

func (p *ParserParser) Useattribs1() (localctx IUseattribs1Context) {
	localctx = NewUseattribs1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ParserParserRULE_useattribs1)
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(517)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(518)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(519)
			p.Useattribs1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(521)
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
	p.EnterRule(localctx, 86, ParserParserRULE_print)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Match(ParserParserRW_print)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(525)
		p.Match(ParserParserTK_lpar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1729461164053102606) != 0 {
		{
			p.SetState(526)
			p.exp(0)
		}

	}
	{
		p.SetState(529)
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
	p.EnterRule(localctx, 88, ParserParserRULE_env)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(531)
		p.Match(ParserParserTK_lbrc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(533)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4453881560512) != 0 {
		{
			p.SetState(532)
			p.Instructions()
		}

	}
	{
		p.SetState(535)
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
	p.EnterRule(localctx, 90, ParserParserRULE_instructions)
	p.SetState(541)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(537)
			p.Instruction()
		}
		{
			p.SetState(538)
			p.Instructions()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(540)
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
	TK_semicolon() antlr.TerminalNode
	Deccst() IDeccstContext
	Ifstruct() IIfstructContext
	Switchstruct() ISwitchstructContext
	Loopfor() ILoopforContext
	Loopwhile() ILoopwhileContext
	Guard() IGuardContext
	Reasign() IReasignContext
	RW_self() antlr.TerminalNode
	TK_dot() antlr.TerminalNode
	Addsub() IAddsubContext
	Decvector() IDecvectorContext
	Funcvector() IFuncvectorContext
	Reasignvector() IReasignvectorContext
	Decmatrix() IDecmatrixContext
	Defstruct() IDefstructContext
	Decstruct() IDecstructContext
	Useattribs() IUseattribsContext
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

func (s *InstructionContext) TK_semicolon() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_semicolon, 0)
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

func (s *InstructionContext) RW_self() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_self, 0)
}

func (s *InstructionContext) TK_dot() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_dot, 0)
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

func (s *InstructionContext) Decmatrix() IDecmatrixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecmatrixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecmatrixContext)
}

func (s *InstructionContext) Defstruct() IDefstructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefstructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefstructContext)
}

func (s *InstructionContext) Decstruct() IDecstructContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDecstructContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDecstructContext)
}

func (s *InstructionContext) Useattribs() IUseattribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseattribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseattribsContext)
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
	p.EnterRule(localctx, 92, ParserParserRULE_instruction)
	var _la int

	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 73, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(543)
			p.Decvar()
		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(544)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(547)
			p.Deccst()
		}
		p.SetState(549)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(548)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(551)
			p.Ifstruct()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(552)
			p.Switchstruct()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(553)
			p.Loopfor()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(554)
			p.Loopwhile()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(555)
			p.Guard()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		p.SetState(558)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(556)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(557)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(560)
			p.Reasign()
		}
		p.SetState(562)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(561)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		p.SetState(566)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(564)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(565)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(568)
			p.Addsub()
		}
		p.SetState(570)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(569)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(572)
			p.Decvector()
		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(573)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(576)
			p.Funcvector()
		}
		p.SetState(578)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(577)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		p.SetState(582)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(580)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(581)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(584)
			p.Reasignvector()
		}
		p.SetState(586)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(585)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(588)
			p.Decmatrix()
		}
		p.SetState(590)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(589)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(592)
			p.Defstruct()
		}
		p.SetState(594)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(593)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(596)
			p.Decstruct()
		}
		p.SetState(598)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(597)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		p.SetState(602)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(600)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(601)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(604)
			p.Useattribs()
		}
		p.SetState(606)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(605)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		p.SetState(610)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(608)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(609)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(612)
			p.Callfunc()
		}
		p.SetState(614)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(613)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(616)
			p.Print_()
		}
		p.SetState(618)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(617)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(620)
			p.Match(ParserParserRW_return)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.exp(0)
		}
		p.SetState(623)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(622)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(625)
			p.Match(ParserParserRW_return)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(627)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(626)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(629)
			p.Match(ParserParserRW_continue)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(631)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(630)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(633)
			p.Match(ParserParserRW_break)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(635)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserTK_semicolon {
			{
				p.SetState(634)
				p.Match(ParserParserTK_semicolon)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
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
	TK_id() antlr.TerminalNode

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

func (s *TypeContext) TK_id() antlr.TerminalNode {
	return s.GetToken(ParserParserTK_id, 0)
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
	p.EnterRule(localctx, 94, ParserParserRULE_type)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(639)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4398046511166) != 0) {
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
	Dims() IDimsContext
	TK_dot() antlr.TerminalNode
	RW_isEmpty() antlr.TerminalNode
	RW_count() antlr.TerminalNode
	Useattribs() IUseattribsContext
	RW_self() antlr.TerminalNode
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

func (s *ExpContext) Dims() IDimsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDimsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDimsContext)
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

func (s *ExpContext) Useattribs() IUseattribsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseattribsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseattribsContext)
}

func (s *ExpContext) RW_self() antlr.TerminalNode {
	return s.GetToken(ParserParserRW_self, 0)
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
	_startState := 96
	p.EnterRecursionRule(localctx, 96, ParserParserRULE_exp, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(695)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 77, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(642)

			var _m = p.Match(ParserParserTK_minus)

			localctx.(*ExpContext).s = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(643)

			var _x = p.exp(26)

			localctx.(*ExpContext).e2 = _x
		}

	case 2:
		{
			p.SetState(644)

			var _m = p.Match(ParserParserTK_not)

			localctx.(*ExpContext).s = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(645)

			var _x = p.exp(20)

			localctx.(*ExpContext).e2 = _x
		}

	case 3:
		{
			p.SetState(646)
			p.Match(ParserParserRW_Int)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(647)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(648)
			p.exp(0)
		}
		{
			p.SetState(649)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(651)
			p.Match(ParserParserRW_Float)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(652)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(653)
			p.exp(0)
		}
		{
			p.SetState(654)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(656)
			p.Match(ParserParserRW_String)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(657)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(658)
			p.exp(0)
		}
		{
			p.SetState(659)
			p.Match(ParserParserTK_rpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		{
			p.SetState(661)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(662)
			p.Dims()
		}

	case 7:
		{
			p.SetState(663)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(664)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(665)
			p.Match(ParserParserRW_isEmpty)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		{
			p.SetState(666)
			p.Match(ParserParserTK_id)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(667)
			p.Match(ParserParserTK_dot)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(668)
			p.Match(ParserParserRW_count)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.SetState(671)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(669)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(670)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(673)
			p.Useattribs()
		}

	case 10:
		p.SetState(676)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(674)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(675)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(678)
			p.Callfunc()
		}

	case 11:
		{
			p.SetState(679)

			var _m = p.Match(ParserParserRW_nil)

			localctx.(*ExpContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.SetState(682)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ParserParserRW_self {
			{
				p.SetState(680)
				p.Match(ParserParserRW_self)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(681)
				p.Match(ParserParserTK_dot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(684)

			var _m = p.Match(ParserParserTK_id)

			localctx.(*ExpContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		{
			p.SetState(685)

			var _m = p.Match(ParserParserTK_string)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		{
			p.SetState(686)

			var _m = p.Match(ParserParserTK_char)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		{
			p.SetState(687)

			var _m = p.Match(ParserParserTK_int)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		{
			p.SetState(688)

			var _m = p.Match(ParserParserTK_float)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 17:
		{
			p.SetState(689)

			var _m = p.Match(ParserParserRW_true)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 18:
		{
			p.SetState(690)

			var _m = p.Match(ParserParserRW_false)

			localctx.(*ExpContext).p = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 19:
		{
			p.SetState(691)
			p.Match(ParserParserTK_lpar)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(692)

			var _x = p.exp(0)

			localctx.(*ExpContext).e = _x
		}
		{
			p.SetState(693)
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
	p.SetState(720)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 79, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(718)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 78, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(697)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
					goto errorExit
				}
				{
					p.SetState(698)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpContext).s = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&985162418487296) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpContext).s = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(699)

					var _x = p.exp(26)

					localctx.(*ExpContext).e2 = _x
				}

			case 2:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(700)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
					goto errorExit
				}
				{
					p.SetState(701)

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
					p.SetState(702)

					var _x = p.exp(25)

					localctx.(*ExpContext).e2 = _x
				}

			case 3:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(703)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
					goto errorExit
				}
				{
					p.SetState(704)

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
					p.SetState(705)

					var _x = p.exp(24)

					localctx.(*ExpContext).e2 = _x
				}

			case 4:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(706)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
					goto errorExit
				}
				{
					p.SetState(707)

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
					p.SetState(708)

					var _x = p.exp(23)

					localctx.(*ExpContext).e2 = _x
				}

			case 5:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(709)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
					goto errorExit
				}
				{
					p.SetState(710)

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
					p.SetState(711)

					var _x = p.exp(22)

					localctx.(*ExpContext).e2 = _x
				}

			case 6:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(712)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
					goto errorExit
				}
				{
					p.SetState(713)

					var _m = p.Match(ParserParserTK_and)

					localctx.(*ExpContext).s = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(714)

					var _x = p.exp(20)

					localctx.(*ExpContext).e2 = _x
				}

			case 7:
				localctx = NewExpContext(p, _parentctx, _parentState)
				localctx.(*ExpContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ParserParserRULE_exp)
				p.SetState(715)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
					goto errorExit
				}
				{
					p.SetState(716)

					var _m = p.Match(ParserParserTK_or)

					localctx.(*ExpContext).s = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(717)

					var _x = p.exp(19)

					localctx.(*ExpContext).e2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(722)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 79, p.GetParserRuleContext())
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
	case 31:
		var t *Listvector2Context = nil
		if localctx != nil {
			t = localctx.(*Listvector2Context)
		}
		return p.Listvector2_Sempred(t, predIndex)

	case 48:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParserParser) Listvector2_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ParserParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 18)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
