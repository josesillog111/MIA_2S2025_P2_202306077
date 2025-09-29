// Code generated from parser/GoDiskGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package godisk // GoDiskGrammar
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

type GoDiskGrammar struct {
	*antlr.BaseParser
}

var GoDiskGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func godiskgrammarParserInit() {
	staticData := &GoDiskGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "'='", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "INT_LIT", "FLOAT_LIT", "STRING_LIT", "CHAR_LIT", "MKDISK", "RMDISK",
		"FDISK", "MOUNT", "MOUNTED", "MKFS", "CAT", "LOGIN", "LOGOUT", "MKGRP",
		"RMGRP", "MKUSR", "RMUSR", "CHGRP", "MKFILE", "MKDIR", "REP", "ID_TEXT",
		"SIZE", "FIT", "UNIT", "PATH", "TYPE", "NAME", "FILE", "USER", "PASS",
		"GRP", "CONT", "PATH_FILE_LS", "R", "P", "ASSIGN", "MINUS", "ID", "UNQUOTED_TEXT",
		"LINE_COMMENT", "BLOCK_COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "stmt", "size", "fit", "unit", "type", "id_text", "path", "name",
		"filen", "user", "pass", "grp", "cont", "path_file_ls", "p", "r", "mkdisk_params",
		"fdisk_params", "mount_params", "mkfs_params", "login_params", "cat_params",
		"mkusr_params", "chgrp_params", "mkfile_params", "mkdir_params", "rep_params",
		"mkdisk_param", "rmdisk_param", "fdisk_param", "mount_param", "mkfs_param",
		"cat_param", "login_param", "mkgrp_param", "rmgrp_param", "mkusr_param",
		"rmusr_param", "chgrp_param", "mkfile_param", "mkdir_param", "rep_param",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 43, 318, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 1, 0, 5, 0, 88, 8, 0, 10, 0, 12, 0, 91, 9, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 127, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		16, 1, 16, 1, 16, 1, 17, 4, 17, 202, 8, 17, 11, 17, 12, 17, 203, 1, 18,
		4, 18, 207, 8, 18, 11, 18, 12, 18, 208, 1, 19, 4, 19, 212, 8, 19, 11, 19,
		12, 19, 213, 1, 20, 4, 20, 217, 8, 20, 11, 20, 12, 20, 218, 1, 21, 4, 21,
		222, 8, 21, 11, 21, 12, 21, 223, 1, 22, 4, 22, 227, 8, 22, 11, 22, 12,
		22, 228, 1, 23, 4, 23, 232, 8, 23, 11, 23, 12, 23, 233, 1, 24, 4, 24, 237,
		8, 24, 11, 24, 12, 24, 238, 1, 25, 4, 25, 242, 8, 25, 11, 25, 12, 25, 243,
		1, 26, 4, 26, 247, 8, 26, 11, 26, 12, 26, 248, 1, 27, 4, 27, 252, 8, 27,
		11, 27, 12, 27, 253, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 260, 8, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 270, 8, 30,
		1, 31, 1, 31, 3, 31, 274, 8, 31, 1, 32, 1, 32, 3, 32, 278, 8, 32, 1, 33,
		1, 33, 1, 34, 1, 34, 1, 34, 3, 34, 285, 8, 34, 1, 35, 1, 35, 1, 36, 1,
		36, 1, 37, 1, 37, 1, 37, 3, 37, 294, 8, 37, 1, 38, 1, 38, 1, 39, 1, 39,
		3, 39, 300, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 306, 8, 40, 1, 41,
		1, 41, 3, 41, 310, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 316, 8, 42,
		1, 42, 0, 0, 43, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
		30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
		66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 0, 1, 2, 0, 3, 3, 40, 40, 324,
		0, 89, 1, 0, 0, 0, 2, 126, 1, 0, 0, 0, 4, 128, 1, 0, 0, 0, 6, 133, 1, 0,
		0, 0, 8, 138, 1, 0, 0, 0, 10, 143, 1, 0, 0, 0, 12, 148, 1, 0, 0, 0, 14,
		153, 1, 0, 0, 0, 16, 158, 1, 0, 0, 0, 18, 163, 1, 0, 0, 0, 20, 169, 1,
		0, 0, 0, 22, 174, 1, 0, 0, 0, 24, 179, 1, 0, 0, 0, 26, 184, 1, 0, 0, 0,
		28, 189, 1, 0, 0, 0, 30, 194, 1, 0, 0, 0, 32, 197, 1, 0, 0, 0, 34, 201,
		1, 0, 0, 0, 36, 206, 1, 0, 0, 0, 38, 211, 1, 0, 0, 0, 40, 216, 1, 0, 0,
		0, 42, 221, 1, 0, 0, 0, 44, 226, 1, 0, 0, 0, 46, 231, 1, 0, 0, 0, 48, 236,
		1, 0, 0, 0, 50, 241, 1, 0, 0, 0, 52, 246, 1, 0, 0, 0, 54, 251, 1, 0, 0,
		0, 56, 259, 1, 0, 0, 0, 58, 261, 1, 0, 0, 0, 60, 269, 1, 0, 0, 0, 62, 273,
		1, 0, 0, 0, 64, 277, 1, 0, 0, 0, 66, 279, 1, 0, 0, 0, 68, 284, 1, 0, 0,
		0, 70, 286, 1, 0, 0, 0, 72, 288, 1, 0, 0, 0, 74, 293, 1, 0, 0, 0, 76, 295,
		1, 0, 0, 0, 78, 299, 1, 0, 0, 0, 80, 305, 1, 0, 0, 0, 82, 309, 1, 0, 0,
		0, 84, 315, 1, 0, 0, 0, 86, 88, 3, 2, 1, 0, 87, 86, 1, 0, 0, 0, 88, 91,
		1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0,
		91, 89, 1, 0, 0, 0, 92, 93, 5, 0, 0, 1, 93, 1, 1, 0, 0, 0, 94, 95, 5, 5,
		0, 0, 95, 127, 3, 34, 17, 0, 96, 97, 5, 6, 0, 0, 97, 127, 3, 58, 29, 0,
		98, 99, 5, 7, 0, 0, 99, 127, 3, 36, 18, 0, 100, 101, 5, 8, 0, 0, 101, 127,
		3, 38, 19, 0, 102, 127, 5, 9, 0, 0, 103, 104, 5, 10, 0, 0, 104, 127, 3,
		40, 20, 0, 105, 106, 5, 11, 0, 0, 106, 127, 3, 44, 22, 0, 107, 108, 5,
		12, 0, 0, 108, 127, 3, 42, 21, 0, 109, 127, 5, 13, 0, 0, 110, 111, 5, 14,
		0, 0, 111, 127, 3, 70, 35, 0, 112, 113, 5, 15, 0, 0, 113, 127, 3, 72, 36,
		0, 114, 115, 5, 16, 0, 0, 115, 127, 3, 46, 23, 0, 116, 117, 5, 17, 0, 0,
		117, 127, 3, 76, 38, 0, 118, 119, 5, 18, 0, 0, 119, 127, 3, 48, 24, 0,
		120, 121, 5, 19, 0, 0, 121, 127, 3, 50, 25, 0, 122, 123, 5, 20, 0, 0, 123,
		127, 3, 52, 26, 0, 124, 125, 5, 21, 0, 0, 125, 127, 3, 54, 27, 0, 126,
		94, 1, 0, 0, 0, 126, 96, 1, 0, 0, 0, 126, 98, 1, 0, 0, 0, 126, 100, 1,
		0, 0, 0, 126, 102, 1, 0, 0, 0, 126, 103, 1, 0, 0, 0, 126, 105, 1, 0, 0,
		0, 126, 107, 1, 0, 0, 0, 126, 109, 1, 0, 0, 0, 126, 110, 1, 0, 0, 0, 126,
		112, 1, 0, 0, 0, 126, 114, 1, 0, 0, 0, 126, 116, 1, 0, 0, 0, 126, 118,
		1, 0, 0, 0, 126, 120, 1, 0, 0, 0, 126, 122, 1, 0, 0, 0, 126, 124, 1, 0,
		0, 0, 127, 3, 1, 0, 0, 0, 128, 129, 5, 38, 0, 0, 129, 130, 5, 23, 0, 0,
		130, 131, 5, 37, 0, 0, 131, 132, 5, 1, 0, 0, 132, 5, 1, 0, 0, 0, 133, 134,
		5, 38, 0, 0, 134, 135, 5, 24, 0, 0, 135, 136, 5, 37, 0, 0, 136, 137, 5,
		39, 0, 0, 137, 7, 1, 0, 0, 0, 138, 139, 5, 38, 0, 0, 139, 140, 5, 25, 0,
		0, 140, 141, 5, 37, 0, 0, 141, 142, 5, 39, 0, 0, 142, 9, 1, 0, 0, 0, 143,
		144, 5, 38, 0, 0, 144, 145, 5, 27, 0, 0, 145, 146, 5, 37, 0, 0, 146, 147,
		5, 39, 0, 0, 147, 11, 1, 0, 0, 0, 148, 149, 5, 38, 0, 0, 149, 150, 5, 22,
		0, 0, 150, 151, 5, 37, 0, 0, 151, 152, 5, 39, 0, 0, 152, 13, 1, 0, 0, 0,
		153, 154, 5, 38, 0, 0, 154, 155, 5, 26, 0, 0, 155, 156, 5, 37, 0, 0, 156,
		157, 7, 0, 0, 0, 157, 15, 1, 0, 0, 0, 158, 159, 5, 38, 0, 0, 159, 160,
		5, 28, 0, 0, 160, 161, 5, 37, 0, 0, 161, 162, 7, 0, 0, 0, 162, 17, 1, 0,
		0, 0, 163, 164, 5, 38, 0, 0, 164, 165, 5, 29, 0, 0, 165, 166, 5, 1, 0,
		0, 166, 167, 5, 37, 0, 0, 167, 168, 7, 0, 0, 0, 168, 19, 1, 0, 0, 0, 169,
		170, 5, 38, 0, 0, 170, 171, 5, 30, 0, 0, 171, 172, 5, 37, 0, 0, 172, 173,
		7, 0, 0, 0, 173, 21, 1, 0, 0, 0, 174, 175, 5, 38, 0, 0, 175, 176, 5, 31,
		0, 0, 176, 177, 5, 37, 0, 0, 177, 178, 7, 0, 0, 0, 178, 23, 1, 0, 0, 0,
		179, 180, 5, 38, 0, 0, 180, 181, 5, 32, 0, 0, 181, 182, 5, 37, 0, 0, 182,
		183, 7, 0, 0, 0, 183, 25, 1, 0, 0, 0, 184, 185, 5, 38, 0, 0, 185, 186,
		5, 33, 0, 0, 186, 187, 5, 37, 0, 0, 187, 188, 7, 0, 0, 0, 188, 27, 1, 0,
		0, 0, 189, 190, 5, 38, 0, 0, 190, 191, 5, 34, 0, 0, 191, 192, 5, 37, 0,
		0, 192, 193, 7, 0, 0, 0, 193, 29, 1, 0, 0, 0, 194, 195, 5, 38, 0, 0, 195,
		196, 5, 36, 0, 0, 196, 31, 1, 0, 0, 0, 197, 198, 5, 38, 0, 0, 198, 199,
		5, 35, 0, 0, 199, 33, 1, 0, 0, 0, 200, 202, 3, 56, 28, 0, 201, 200, 1,
		0, 0, 0, 202, 203, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203, 204, 1, 0, 0,
		0, 204, 35, 1, 0, 0, 0, 205, 207, 3, 60, 30, 0, 206, 205, 1, 0, 0, 0, 207,
		208, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0, 0, 209, 37, 1,
		0, 0, 0, 210, 212, 3, 62, 31, 0, 211, 210, 1, 0, 0, 0, 212, 213, 1, 0,
		0, 0, 213, 211, 1, 0, 0, 0, 213, 214, 1, 0, 0, 0, 214, 39, 1, 0, 0, 0,
		215, 217, 3, 64, 32, 0, 216, 215, 1, 0, 0, 0, 217, 218, 1, 0, 0, 0, 218,
		216, 1, 0, 0, 0, 218, 219, 1, 0, 0, 0, 219, 41, 1, 0, 0, 0, 220, 222, 3,
		68, 34, 0, 221, 220, 1, 0, 0, 0, 222, 223, 1, 0, 0, 0, 223, 221, 1, 0,
		0, 0, 223, 224, 1, 0, 0, 0, 224, 43, 1, 0, 0, 0, 225, 227, 3, 66, 33, 0,
		226, 225, 1, 0, 0, 0, 227, 228, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228,
		229, 1, 0, 0, 0, 229, 45, 1, 0, 0, 0, 230, 232, 3, 74, 37, 0, 231, 230,
		1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 231, 1, 0, 0, 0, 233, 234, 1, 0,
		0, 0, 234, 47, 1, 0, 0, 0, 235, 237, 3, 78, 39, 0, 236, 235, 1, 0, 0, 0,
		237, 238, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239,
		49, 1, 0, 0, 0, 240, 242, 3, 80, 40, 0, 241, 240, 1, 0, 0, 0, 242, 243,
		1, 0, 0, 0, 243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 51, 1, 0,
		0, 0, 245, 247, 3, 82, 41, 0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0,
		0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 53, 1, 0, 0, 0, 250,
		252, 3, 84, 42, 0, 251, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253, 251,
		1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 55, 1, 0, 0, 0, 255, 260, 3, 4,
		2, 0, 256, 260, 3, 6, 3, 0, 257, 260, 3, 8, 4, 0, 258, 260, 3, 14, 7, 0,
		259, 255, 1, 0, 0, 0, 259, 256, 1, 0, 0, 0, 259, 257, 1, 0, 0, 0, 259,
		258, 1, 0, 0, 0, 260, 57, 1, 0, 0, 0, 261, 262, 3, 14, 7, 0, 262, 59, 1,
		0, 0, 0, 263, 270, 3, 4, 2, 0, 264, 270, 3, 6, 3, 0, 265, 270, 3, 8, 4,
		0, 266, 270, 3, 14, 7, 0, 267, 270, 3, 10, 5, 0, 268, 270, 3, 16, 8, 0,
		269, 263, 1, 0, 0, 0, 269, 264, 1, 0, 0, 0, 269, 265, 1, 0, 0, 0, 269,
		266, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 268, 1, 0, 0, 0, 270, 61, 1,
		0, 0, 0, 271, 274, 3, 14, 7, 0, 272, 274, 3, 16, 8, 0, 273, 271, 1, 0,
		0, 0, 273, 272, 1, 0, 0, 0, 274, 63, 1, 0, 0, 0, 275, 278, 3, 12, 6, 0,
		276, 278, 3, 10, 5, 0, 277, 275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278,
		65, 1, 0, 0, 0, 279, 280, 3, 18, 9, 0, 280, 67, 1, 0, 0, 0, 281, 285, 3,
		20, 10, 0, 282, 285, 3, 22, 11, 0, 283, 285, 3, 12, 6, 0, 284, 281, 1,
		0, 0, 0, 284, 282, 1, 0, 0, 0, 284, 283, 1, 0, 0, 0, 285, 69, 1, 0, 0,
		0, 286, 287, 3, 16, 8, 0, 287, 71, 1, 0, 0, 0, 288, 289, 3, 16, 8, 0, 289,
		73, 1, 0, 0, 0, 290, 294, 3, 20, 10, 0, 291, 294, 3, 22, 11, 0, 292, 294,
		3, 24, 12, 0, 293, 290, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 292, 1,
		0, 0, 0, 294, 75, 1, 0, 0, 0, 295, 296, 3, 20, 10, 0, 296, 77, 1, 0, 0,
		0, 297, 300, 3, 20, 10, 0, 298, 300, 3, 24, 12, 0, 299, 297, 1, 0, 0, 0,
		299, 298, 1, 0, 0, 0, 300, 79, 1, 0, 0, 0, 301, 306, 3, 14, 7, 0, 302,
		306, 3, 32, 16, 0, 303, 306, 3, 4, 2, 0, 304, 306, 3, 26, 13, 0, 305, 301,
		1, 0, 0, 0, 305, 302, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 304, 1, 0,
		0, 0, 306, 81, 1, 0, 0, 0, 307, 310, 3, 30, 15, 0, 308, 310, 3, 14, 7,
		0, 309, 307, 1, 0, 0, 0, 309, 308, 1, 0, 0, 0, 310, 83, 1, 0, 0, 0, 311,
		316, 3, 16, 8, 0, 312, 316, 3, 14, 7, 0, 313, 316, 3, 12, 6, 0, 314, 316,
		3, 28, 14, 0, 315, 311, 1, 0, 0, 0, 315, 312, 1, 0, 0, 0, 315, 313, 1,
		0, 0, 0, 315, 314, 1, 0, 0, 0, 316, 85, 1, 0, 0, 0, 23, 89, 126, 203, 208,
		213, 218, 223, 228, 233, 238, 243, 248, 253, 259, 269, 273, 277, 284, 293,
		299, 305, 309, 315,
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

// GoDiskGrammarInit initializes any static state used to implement GoDiskGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoDiskGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoDiskGrammarInit() {
	staticData := &GoDiskGrammarParserStaticData
	staticData.once.Do(godiskgrammarParserInit)
}

// NewGoDiskGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewGoDiskGrammar(input antlr.TokenStream) *GoDiskGrammar {
	GoDiskGrammarInit()
	this := new(GoDiskGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoDiskGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "GoDiskGrammar.g4"

	return this
}

// GoDiskGrammar tokens.
const (
	GoDiskGrammarEOF           = antlr.TokenEOF
	GoDiskGrammarINT_LIT       = 1
	GoDiskGrammarFLOAT_LIT     = 2
	GoDiskGrammarSTRING_LIT    = 3
	GoDiskGrammarCHAR_LIT      = 4
	GoDiskGrammarMKDISK        = 5
	GoDiskGrammarRMDISK        = 6
	GoDiskGrammarFDISK         = 7
	GoDiskGrammarMOUNT         = 8
	GoDiskGrammarMOUNTED       = 9
	GoDiskGrammarMKFS          = 10
	GoDiskGrammarCAT           = 11
	GoDiskGrammarLOGIN         = 12
	GoDiskGrammarLOGOUT        = 13
	GoDiskGrammarMKGRP         = 14
	GoDiskGrammarRMGRP         = 15
	GoDiskGrammarMKUSR         = 16
	GoDiskGrammarRMUSR         = 17
	GoDiskGrammarCHGRP         = 18
	GoDiskGrammarMKFILE        = 19
	GoDiskGrammarMKDIR         = 20
	GoDiskGrammarREP           = 21
	GoDiskGrammarID_TEXT       = 22
	GoDiskGrammarSIZE          = 23
	GoDiskGrammarFIT           = 24
	GoDiskGrammarUNIT          = 25
	GoDiskGrammarPATH          = 26
	GoDiskGrammarTYPE          = 27
	GoDiskGrammarNAME          = 28
	GoDiskGrammarFILE          = 29
	GoDiskGrammarUSER          = 30
	GoDiskGrammarPASS          = 31
	GoDiskGrammarGRP           = 32
	GoDiskGrammarCONT          = 33
	GoDiskGrammarPATH_FILE_LS  = 34
	GoDiskGrammarR             = 35
	GoDiskGrammarP             = 36
	GoDiskGrammarASSIGN        = 37
	GoDiskGrammarMINUS         = 38
	GoDiskGrammarID            = 39
	GoDiskGrammarUNQUOTED_TEXT = 40
	GoDiskGrammarLINE_COMMENT  = 41
	GoDiskGrammarBLOCK_COMMENT = 42
	GoDiskGrammarWS            = 43
)

// GoDiskGrammar rules.
const (
	GoDiskGrammarRULE_prog          = 0
	GoDiskGrammarRULE_stmt          = 1
	GoDiskGrammarRULE_size          = 2
	GoDiskGrammarRULE_fit           = 3
	GoDiskGrammarRULE_unit          = 4
	GoDiskGrammarRULE_type          = 5
	GoDiskGrammarRULE_id_text       = 6
	GoDiskGrammarRULE_path          = 7
	GoDiskGrammarRULE_name          = 8
	GoDiskGrammarRULE_filen         = 9
	GoDiskGrammarRULE_user          = 10
	GoDiskGrammarRULE_pass          = 11
	GoDiskGrammarRULE_grp           = 12
	GoDiskGrammarRULE_cont          = 13
	GoDiskGrammarRULE_path_file_ls  = 14
	GoDiskGrammarRULE_p             = 15
	GoDiskGrammarRULE_r             = 16
	GoDiskGrammarRULE_mkdisk_params = 17
	GoDiskGrammarRULE_fdisk_params  = 18
	GoDiskGrammarRULE_mount_params  = 19
	GoDiskGrammarRULE_mkfs_params   = 20
	GoDiskGrammarRULE_login_params  = 21
	GoDiskGrammarRULE_cat_params    = 22
	GoDiskGrammarRULE_mkusr_params  = 23
	GoDiskGrammarRULE_chgrp_params  = 24
	GoDiskGrammarRULE_mkfile_params = 25
	GoDiskGrammarRULE_mkdir_params  = 26
	GoDiskGrammarRULE_rep_params    = 27
	GoDiskGrammarRULE_mkdisk_param  = 28
	GoDiskGrammarRULE_rmdisk_param  = 29
	GoDiskGrammarRULE_fdisk_param   = 30
	GoDiskGrammarRULE_mount_param   = 31
	GoDiskGrammarRULE_mkfs_param    = 32
	GoDiskGrammarRULE_cat_param     = 33
	GoDiskGrammarRULE_login_param   = 34
	GoDiskGrammarRULE_mkgrp_param   = 35
	GoDiskGrammarRULE_rmgrp_param   = 36
	GoDiskGrammarRULE_mkusr_param   = 37
	GoDiskGrammarRULE_rmusr_param   = 38
	GoDiskGrammarRULE_chgrp_param   = 39
	GoDiskGrammarRULE_mkfile_param  = 40
	GoDiskGrammarRULE_mkdir_param   = 41
	GoDiskGrammarRULE_rep_param     = 42
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarEOF, 0)
}

func (s *ProgContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoDiskGrammarRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4194272) != 0 {
		{
			p.SetState(86)
			p.Stmt()
		}

		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(92)
		p.Match(GoDiskGrammarEOF)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LOGOUTContext struct {
	StmtContext
}

func NewLOGOUTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LOGOUTContext {
	var p = new(LOGOUTContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *LOGOUTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LOGOUTContext) LOGOUT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarLOGOUT, 0)
}

func (s *LOGOUTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterLOGOUT(s)
	}
}

func (s *LOGOUTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitLOGOUT(s)
	}
}

func (s *LOGOUTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitLOGOUT(s)

	default:
		return t.VisitChildren(s)
	}
}

type MKGRPContext struct {
	StmtContext
}

func NewMKGRPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MKGRPContext {
	var p = new(MKGRPContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MKGRPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MKGRPContext) MKGRP() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMKGRP, 0)
}

func (s *MKGRPContext) Mkgrp_param() IMkgrp_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkgrp_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkgrp_paramContext)
}

func (s *MKGRPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMKGRP(s)
	}
}

func (s *MKGRPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMKGRP(s)
	}
}

func (s *MKGRPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMKGRP(s)

	default:
		return t.VisitChildren(s)
	}
}

type RMGRPContext struct {
	StmtContext
}

func NewRMGRPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RMGRPContext {
	var p = new(RMGRPContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *RMGRPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RMGRPContext) RMGRP() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarRMGRP, 0)
}

func (s *RMGRPContext) Rmgrp_param() IRmgrp_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmgrp_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmgrp_paramContext)
}

func (s *RMGRPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRMGRP(s)
	}
}

func (s *RMGRPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRMGRP(s)
	}
}

func (s *RMGRPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRMGRP(s)

	default:
		return t.VisitChildren(s)
	}
}

type MKFSContext struct {
	StmtContext
}

func NewMKFSContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MKFSContext {
	var p = new(MKFSContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MKFSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MKFSContext) MKFS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMKFS, 0)
}

func (s *MKFSContext) Mkfs_params() IMkfs_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfs_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfs_paramsContext)
}

func (s *MKFSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMKFS(s)
	}
}

func (s *MKFSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMKFS(s)
	}
}

func (s *MKFSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMKFS(s)

	default:
		return t.VisitChildren(s)
	}
}

type MKDIRContext struct {
	StmtContext
}

func NewMKDIRContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MKDIRContext {
	var p = new(MKDIRContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MKDIRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MKDIRContext) MKDIR() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMKDIR, 0)
}

func (s *MKDIRContext) Mkdir_params() IMkdir_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdir_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdir_paramsContext)
}

func (s *MKDIRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMKDIR(s)
	}
}

func (s *MKDIRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMKDIR(s)
	}
}

func (s *MKDIRContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMKDIR(s)

	default:
		return t.VisitChildren(s)
	}
}

type CHGRPContext struct {
	StmtContext
}

func NewCHGRPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CHGRPContext {
	var p = new(CHGRPContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *CHGRPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CHGRPContext) CHGRP() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCHGRP, 0)
}

func (s *CHGRPContext) Chgrp_params() IChgrp_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChgrp_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChgrp_paramsContext)
}

func (s *CHGRPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCHGRP(s)
	}
}

func (s *CHGRPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCHGRP(s)
	}
}

func (s *CHGRPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCHGRP(s)

	default:
		return t.VisitChildren(s)
	}
}

type RMDISKContext struct {
	StmtContext
}

func NewRMDISKContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RMDISKContext {
	var p = new(RMDISKContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *RMDISKContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RMDISKContext) RMDISK() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarRMDISK, 0)
}

func (s *RMDISKContext) Rmdisk_param() IRmdisk_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmdisk_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmdisk_paramContext)
}

func (s *RMDISKContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRMDISK(s)
	}
}

func (s *RMDISKContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRMDISK(s)
	}
}

func (s *RMDISKContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRMDISK(s)

	default:
		return t.VisitChildren(s)
	}
}

type MKUSRContext struct {
	StmtContext
}

func NewMKUSRContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MKUSRContext {
	var p = new(MKUSRContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MKUSRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MKUSRContext) MKUSR() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMKUSR, 0)
}

func (s *MKUSRContext) Mkusr_params() IMkusr_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusr_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkusr_paramsContext)
}

func (s *MKUSRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMKUSR(s)
	}
}

func (s *MKUSRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMKUSR(s)
	}
}

func (s *MKUSRContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMKUSR(s)

	default:
		return t.VisitChildren(s)
	}
}

type CATContext struct {
	StmtContext
}

func NewCATContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CATContext {
	var p = new(CATContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *CATContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CATContext) CAT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCAT, 0)
}

func (s *CATContext) Cat_params() ICat_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICat_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICat_paramsContext)
}

func (s *CATContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCAT(s)
	}
}

func (s *CATContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCAT(s)
	}
}

func (s *CATContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCAT(s)

	default:
		return t.VisitChildren(s)
	}
}

type FDISKContext struct {
	StmtContext
}

func NewFDISKContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FDISKContext {
	var p = new(FDISKContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *FDISKContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FDISKContext) FDISK() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarFDISK, 0)
}

func (s *FDISKContext) Fdisk_params() IFdisk_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdisk_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFdisk_paramsContext)
}

func (s *FDISKContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFDISK(s)
	}
}

func (s *FDISKContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFDISK(s)
	}
}

func (s *FDISKContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFDISK(s)

	default:
		return t.VisitChildren(s)
	}
}

type MOUNTEDContext struct {
	StmtContext
}

func NewMOUNTEDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MOUNTEDContext {
	var p = new(MOUNTEDContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MOUNTEDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MOUNTEDContext) MOUNTED() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMOUNTED, 0)
}

func (s *MOUNTEDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMOUNTED(s)
	}
}

func (s *MOUNTEDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMOUNTED(s)
	}
}

func (s *MOUNTEDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMOUNTED(s)

	default:
		return t.VisitChildren(s)
	}
}

type LOGINContext struct {
	StmtContext
}

func NewLOGINContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LOGINContext {
	var p = new(LOGINContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *LOGINContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LOGINContext) LOGIN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarLOGIN, 0)
}

func (s *LOGINContext) Login_params() ILogin_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogin_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILogin_paramsContext)
}

func (s *LOGINContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterLOGIN(s)
	}
}

func (s *LOGINContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitLOGIN(s)
	}
}

func (s *LOGINContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitLOGIN(s)

	default:
		return t.VisitChildren(s)
	}
}

type MOUNTContext struct {
	StmtContext
}

func NewMOUNTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MOUNTContext {
	var p = new(MOUNTContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MOUNTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MOUNTContext) MOUNT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMOUNT, 0)
}

func (s *MOUNTContext) Mount_params() IMount_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMount_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMount_paramsContext)
}

func (s *MOUNTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMOUNT(s)
	}
}

func (s *MOUNTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMOUNT(s)
	}
}

func (s *MOUNTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMOUNT(s)

	default:
		return t.VisitChildren(s)
	}
}

type MKFILEContext struct {
	StmtContext
}

func NewMKFILEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MKFILEContext {
	var p = new(MKFILEContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MKFILEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MKFILEContext) MKFILE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMKFILE, 0)
}

func (s *MKFILEContext) Mkfile_params() IMkfile_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfile_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkfile_paramsContext)
}

func (s *MKFILEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMKFILE(s)
	}
}

func (s *MKFILEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMKFILE(s)
	}
}

func (s *MKFILEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMKFILE(s)

	default:
		return t.VisitChildren(s)
	}
}

type REPContext struct {
	StmtContext
}

func NewREPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *REPContext {
	var p = new(REPContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *REPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *REPContext) REP() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarREP, 0)
}

func (s *REPContext) Rep_params() IRep_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRep_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRep_paramsContext)
}

func (s *REPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterREP(s)
	}
}

func (s *REPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitREP(s)
	}
}

func (s *REPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitREP(s)

	default:
		return t.VisitChildren(s)
	}
}

type MKDISKContext struct {
	StmtContext
}

func NewMKDISKContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MKDISKContext {
	var p = new(MKDISKContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MKDISKContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MKDISKContext) MKDISK() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMKDISK, 0)
}

func (s *MKDISKContext) Mkdisk_params() IMkdisk_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdisk_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMkdisk_paramsContext)
}

func (s *MKDISKContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMKDISK(s)
	}
}

func (s *MKDISKContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMKDISK(s)
	}
}

func (s *MKDISKContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMKDISK(s)

	default:
		return t.VisitChildren(s)
	}
}

type RMUSRContext struct {
	StmtContext
}

func NewRMUSRContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RMUSRContext {
	var p = new(RMUSRContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *RMUSRContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RMUSRContext) RMUSR() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarRMUSR, 0)
}

func (s *RMUSRContext) Rmusr_param() IRmusr_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRmusr_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRmusr_paramContext)
}

func (s *RMUSRContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRMUSR(s)
	}
}

func (s *RMUSRContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRMUSR(s)
	}
}

func (s *RMUSRContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRMUSR(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoDiskGrammarRULE_stmt)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoDiskGrammarMKDISK:
		localctx = NewMKDISKContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.Match(GoDiskGrammarMKDISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(95)
			p.Mkdisk_params()
		}

	case GoDiskGrammarRMDISK:
		localctx = NewRMDISKContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(GoDiskGrammarRMDISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(97)
			p.Rmdisk_param()
		}

	case GoDiskGrammarFDISK:
		localctx = NewFDISKContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.Match(GoDiskGrammarFDISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(99)
			p.Fdisk_params()
		}

	case GoDiskGrammarMOUNT:
		localctx = NewMOUNTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(100)
			p.Match(GoDiskGrammarMOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.Mount_params()
		}

	case GoDiskGrammarMOUNTED:
		localctx = NewMOUNTEDContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(102)
			p.Match(GoDiskGrammarMOUNTED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoDiskGrammarMKFS:
		localctx = NewMKFSContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(103)
			p.Match(GoDiskGrammarMKFS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(104)
			p.Mkfs_params()
		}

	case GoDiskGrammarCAT:
		localctx = NewCATContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(105)
			p.Match(GoDiskGrammarCAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Cat_params()
		}

	case GoDiskGrammarLOGIN:
		localctx = NewLOGINContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(107)
			p.Match(GoDiskGrammarLOGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Login_params()
		}

	case GoDiskGrammarLOGOUT:
		localctx = NewLOGOUTContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(109)
			p.Match(GoDiskGrammarLOGOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoDiskGrammarMKGRP:
		localctx = NewMKGRPContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(110)
			p.Match(GoDiskGrammarMKGRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.Mkgrp_param()
		}

	case GoDiskGrammarRMGRP:
		localctx = NewRMGRPContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(112)
			p.Match(GoDiskGrammarRMGRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Rmgrp_param()
		}

	case GoDiskGrammarMKUSR:
		localctx = NewMKUSRContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(114)
			p.Match(GoDiskGrammarMKUSR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(115)
			p.Mkusr_params()
		}

	case GoDiskGrammarRMUSR:
		localctx = NewRMUSRContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(116)
			p.Match(GoDiskGrammarRMUSR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.Rmusr_param()
		}

	case GoDiskGrammarCHGRP:
		localctx = NewCHGRPContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(118)
			p.Match(GoDiskGrammarCHGRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Chgrp_params()
		}

	case GoDiskGrammarMKFILE:
		localctx = NewMKFILEContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(120)
			p.Match(GoDiskGrammarMKFILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.Mkfile_params()
		}

	case GoDiskGrammarMKDIR:
		localctx = NewMKDIRContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(122)
			p.Match(GoDiskGrammarMKDIR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Mkdir_params()
		}

	case GoDiskGrammarREP:
		localctx = NewREPContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(124)
			p.Match(GoDiskGrammarREP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Rep_params()
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

// ISizeContext is an interface to support dynamic dispatch.
type ISizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	SIZE() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode

	// IsSizeContext differentiates from other interfaces.
	IsSizeContext()
}

type SizeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizeContext() *SizeContext {
	var p = new(SizeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_size
	return p
}

func InitEmptySizeContext(p *SizeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_size
}

func (*SizeContext) IsSizeContext() {}

func NewSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeContext {
	var p = new(SizeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_size

	return p
}

func (s *SizeContext) GetParser() antlr.Parser { return s.parser }

func (s *SizeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *SizeContext) SIZE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSIZE, 0)
}

func (s *SizeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *SizeContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarINT_LIT, 0)
}

func (s *SizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterSize(s)
	}
}

func (s *SizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitSize(s)
	}
}

func (s *SizeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitSize(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Size() (localctx ISizeContext) {
	localctx = NewSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoDiskGrammarRULE_size)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(129)
		p.Match(GoDiskGrammarSIZE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(GoDiskGrammarINT_LIT)
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

// IFitContext is an interface to support dynamic dispatch.
type IFitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	FIT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsFitContext differentiates from other interfaces.
	IsFitContext()
}

type FitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFitContext() *FitContext {
	var p = new(FitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fit
	return p
}

func InitEmptyFitContext(p *FitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fit
}

func (*FitContext) IsFitContext() {}

func NewFitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FitContext {
	var p = new(FitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_fit

	return p
}

func (s *FitContext) GetParser() antlr.Parser { return s.parser }

func (s *FitContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *FitContext) FIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarFIT, 0)
}

func (s *FitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *FitContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *FitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFit(s)
	}
}

func (s *FitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFit(s)
	}
}

func (s *FitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Fit() (localctx IFitContext) {
	localctx = NewFitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoDiskGrammarRULE_fit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.Match(GoDiskGrammarFIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(135)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(GoDiskGrammarID)
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

// IUnitContext is an interface to support dynamic dispatch.
type IUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	UNIT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsUnitContext differentiates from other interfaces.
	IsUnitContext()
}

type UnitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnitContext() *UnitContext {
	var p = new(UnitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_unit
	return p
}

func InitEmptyUnitContext(p *UnitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_unit
}

func (*UnitContext) IsUnitContext() {}

func NewUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnitContext {
	var p = new(UnitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_unit

	return p
}

func (s *UnitContext) GetParser() antlr.Parser { return s.parser }

func (s *UnitContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *UnitContext) UNIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNIT, 0)
}

func (s *UnitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *UnitContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *UnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterUnit(s)
	}
}

func (s *UnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitUnit(s)
	}
}

func (s *UnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Unit() (localctx IUnitContext) {
	localctx = NewUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoDiskGrammarRULE_unit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Match(GoDiskGrammarUNIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(GoDiskGrammarID)
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	TYPE() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

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
	p.RuleIndex = GoDiskGrammarRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *TypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarTYPE, 0)
}

func (s *TypeContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoDiskGrammarRULE_type)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.Match(GoDiskGrammarTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.Match(GoDiskGrammarID)
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

// IId_textContext is an interface to support dynamic dispatch.
type IId_textContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	ID_TEXT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsId_textContext differentiates from other interfaces.
	IsId_textContext()
}

type Id_textContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyId_textContext() *Id_textContext {
	var p = new(Id_textContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_id_text
	return p
}

func InitEmptyId_textContext(p *Id_textContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_id_text
}

func (*Id_textContext) IsId_textContext() {}

func NewId_textContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Id_textContext {
	var p = new(Id_textContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_id_text

	return p
}

func (s *Id_textContext) GetParser() antlr.Parser { return s.parser }

func (s *Id_textContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *Id_textContext) ID_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID_TEXT, 0)
}

func (s *Id_textContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *Id_textContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *Id_textContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Id_textContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Id_textContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterId_text(s)
	}
}

func (s *Id_textContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitId_text(s)
	}
}

func (s *Id_textContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitId_text(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Id_text() (localctx IId_textContext) {
	localctx = NewId_textContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoDiskGrammarRULE_id_text)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(GoDiskGrammarID_TEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(GoDiskGrammarID)
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

// IPathContext is an interface to support dynamic dispatch.
type IPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	PATH() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsPathContext differentiates from other interfaces.
	IsPathContext()
}

type PathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPathContext() *PathContext {
	var p = new(PathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_path
	return p
}

func InitEmptyPathContext(p *PathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_path
}

func (*PathContext) IsPathContext() {}

func NewPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PathContext {
	var p = new(PathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_path

	return p
}

func (s *PathContext) GetParser() antlr.Parser { return s.parser }

func (s *PathContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *PathContext) PATH() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarPATH, 0)
}

func (s *PathContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *PathContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *PathContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *PathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterPath(s)
	}
}

func (s *PathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitPath(s)
	}
}

func (s *PathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Path() (localctx IPathContext) {
	localctx = NewPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoDiskGrammarRULE_path)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(154)
		p.Match(GoDiskGrammarPATH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(155)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	NAME() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_name
	return p
}

func InitEmptyNameContext(p *NameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_name
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarNAME, 0)
}

func (s *NameContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *NameContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *NameContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoDiskGrammarRULE_name)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(159)
		p.Match(GoDiskGrammarNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IFilenContext is an interface to support dynamic dispatch.
type IFilenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	FILE() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsFilenContext differentiates from other interfaces.
	IsFilenContext()
}

type FilenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilenContext() *FilenContext {
	var p = new(FilenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_filen
	return p
}

func InitEmptyFilenContext(p *FilenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_filen
}

func (*FilenContext) IsFilenContext() {}

func NewFilenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilenContext {
	var p = new(FilenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_filen

	return p
}

func (s *FilenContext) GetParser() antlr.Parser { return s.parser }

func (s *FilenContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *FilenContext) FILE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarFILE, 0)
}

func (s *FilenContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarINT_LIT, 0)
}

func (s *FilenContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *FilenContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *FilenContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *FilenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFilen(s)
	}
}

func (s *FilenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFilen(s)
	}
}

func (s *FilenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFilen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Filen() (localctx IFilenContext) {
	localctx = NewFilenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoDiskGrammarRULE_filen)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Match(GoDiskGrammarFILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(165)
		p.Match(GoDiskGrammarINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IUserContext is an interface to support dynamic dispatch.
type IUserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	USER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsUserContext differentiates from other interfaces.
	IsUserContext()
}

type UserContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserContext() *UserContext {
	var p = new(UserContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_user
	return p
}

func InitEmptyUserContext(p *UserContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_user
}

func (*UserContext) IsUserContext() {}

func NewUserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserContext {
	var p = new(UserContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_user

	return p
}

func (s *UserContext) GetParser() antlr.Parser { return s.parser }

func (s *UserContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *UserContext) USER() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUSER, 0)
}

func (s *UserContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *UserContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *UserContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *UserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterUser(s)
	}
}

func (s *UserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitUser(s)
	}
}

func (s *UserContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitUser(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) User() (localctx IUserContext) {
	localctx = NewUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GoDiskGrammarRULE_user)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(GoDiskGrammarUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IPassContext is an interface to support dynamic dispatch.
type IPassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	PASS() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsPassContext differentiates from other interfaces.
	IsPassContext()
}

type PassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassContext() *PassContext {
	var p = new(PassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_pass
	return p
}

func InitEmptyPassContext(p *PassContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_pass
}

func (*PassContext) IsPassContext() {}

func NewPassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassContext {
	var p = new(PassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_pass

	return p
}

func (s *PassContext) GetParser() antlr.Parser { return s.parser }

func (s *PassContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *PassContext) PASS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarPASS, 0)
}

func (s *PassContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *PassContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *PassContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *PassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterPass(s)
	}
}

func (s *PassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitPass(s)
	}
}

func (s *PassContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitPass(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Pass() (localctx IPassContext) {
	localctx = NewPassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoDiskGrammarRULE_pass)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(GoDiskGrammarPASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IGrpContext is an interface to support dynamic dispatch.
type IGrpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	GRP() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsGrpContext differentiates from other interfaces.
	IsGrpContext()
}

type GrpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGrpContext() *GrpContext {
	var p = new(GrpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_grp
	return p
}

func InitEmptyGrpContext(p *GrpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_grp
}

func (*GrpContext) IsGrpContext() {}

func NewGrpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GrpContext {
	var p = new(GrpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_grp

	return p
}

func (s *GrpContext) GetParser() antlr.Parser { return s.parser }

func (s *GrpContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *GrpContext) GRP() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarGRP, 0)
}

func (s *GrpContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *GrpContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *GrpContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *GrpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GrpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GrpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterGrp(s)
	}
}

func (s *GrpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitGrp(s)
	}
}

func (s *GrpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitGrp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Grp() (localctx IGrpContext) {
	localctx = NewGrpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoDiskGrammarRULE_grp)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(180)
		p.Match(GoDiskGrammarGRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(182)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IContContext is an interface to support dynamic dispatch.
type IContContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	CONT() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsContContext differentiates from other interfaces.
	IsContContext()
}

type ContContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContContext() *ContContext {
	var p = new(ContContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_cont
	return p
}

func InitEmptyContContext(p *ContContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_cont
}

func (*ContContext) IsContContext() {}

func NewContContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContContext {
	var p = new(ContContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_cont

	return p
}

func (s *ContContext) GetParser() antlr.Parser { return s.parser }

func (s *ContContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *ContContext) CONT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCONT, 0)
}

func (s *ContContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *ContContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *ContContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *ContContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCont(s)
	}
}

func (s *ContContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCont(s)
	}
}

func (s *ContContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCont(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Cont() (localctx IContContext) {
	localctx = NewContContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoDiskGrammarRULE_cont)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(GoDiskGrammarCONT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(186)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IPath_file_lsContext is an interface to support dynamic dispatch.
type IPath_file_lsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	PATH_FILE_LS() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsPath_file_lsContext differentiates from other interfaces.
	IsPath_file_lsContext()
}

type Path_file_lsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPath_file_lsContext() *Path_file_lsContext {
	var p = new(Path_file_lsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_path_file_ls
	return p
}

func InitEmptyPath_file_lsContext(p *Path_file_lsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_path_file_ls
}

func (*Path_file_lsContext) IsPath_file_lsContext() {}

func NewPath_file_lsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Path_file_lsContext {
	var p = new(Path_file_lsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_path_file_ls

	return p
}

func (s *Path_file_lsContext) GetParser() antlr.Parser { return s.parser }

func (s *Path_file_lsContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *Path_file_lsContext) PATH_FILE_LS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarPATH_FILE_LS, 0)
}

func (s *Path_file_lsContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *Path_file_lsContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *Path_file_lsContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *Path_file_lsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Path_file_lsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Path_file_lsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterPath_file_ls(s)
	}
}

func (s *Path_file_lsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitPath_file_ls(s)
	}
}

func (s *Path_file_lsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitPath_file_ls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Path_file_ls() (localctx IPath_file_lsContext) {
	localctx = NewPath_file_lsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoDiskGrammarRULE_path_file_ls)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
		p.Match(GoDiskGrammarPATH_FILE_LS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(191)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoDiskGrammarSTRING_LIT || _la == GoDiskGrammarUNQUOTED_TEXT) {
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

// IPContext is an interface to support dynamic dispatch.
type IPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	P() antlr.TerminalNode

	// IsPContext differentiates from other interfaces.
	IsPContext()
}

type PContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPContext() *PContext {
	var p = new(PContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_p
	return p
}

func InitEmptyPContext(p *PContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_p
}

func (*PContext) IsPContext() {}

func NewPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PContext {
	var p = new(PContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_p

	return p
}

func (s *PContext) GetParser() antlr.Parser { return s.parser }

func (s *PContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *PContext) P() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarP, 0)
}

func (s *PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterP(s)
	}
}

func (s *PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitP(s)
	}
}

func (s *PContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) P() (localctx IPContext) {
	localctx = NewPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoDiskGrammarRULE_p)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(195)
		p.Match(GoDiskGrammarP)
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

// IRContext is an interface to support dynamic dispatch.
type IRContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	R() antlr.TerminalNode

	// IsRContext differentiates from other interfaces.
	IsRContext()
}

type RContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRContext() *RContext {
	var p = new(RContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_r
	return p
}

func InitEmptyRContext(p *RContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_r
}

func (*RContext) IsRContext() {}

func NewRContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RContext {
	var p = new(RContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_r

	return p
}

func (s *RContext) GetParser() antlr.Parser { return s.parser }

func (s *RContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *RContext) R() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarR, 0)
}

func (s *RContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterR(s)
	}
}

func (s *RContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitR(s)
	}
}

func (s *RContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitR(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) R() (localctx IRContext) {
	localctx = NewRContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoDiskGrammarRULE_r)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(GoDiskGrammarR)
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

// IMkdisk_paramsContext is an interface to support dynamic dispatch.
type IMkdisk_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMkdisk_param() []IMkdisk_paramContext
	Mkdisk_param(i int) IMkdisk_paramContext

	// IsMkdisk_paramsContext differentiates from other interfaces.
	IsMkdisk_paramsContext()
}

type Mkdisk_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdisk_paramsContext() *Mkdisk_paramsContext {
	var p = new(Mkdisk_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdisk_params
	return p
}

func InitEmptyMkdisk_paramsContext(p *Mkdisk_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdisk_params
}

func (*Mkdisk_paramsContext) IsMkdisk_paramsContext() {}

func NewMkdisk_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkdisk_paramsContext {
	var p = new(Mkdisk_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkdisk_params

	return p
}

func (s *Mkdisk_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkdisk_paramsContext) AllMkdisk_param() []IMkdisk_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMkdisk_paramContext); ok {
			len++
		}
	}

	tst := make([]IMkdisk_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMkdisk_paramContext); ok {
			tst[i] = t.(IMkdisk_paramContext)
			i++
		}
	}

	return tst
}

func (s *Mkdisk_paramsContext) Mkdisk_param(i int) IMkdisk_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdisk_paramContext); ok {
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

	return t.(IMkdisk_paramContext)
}

func (s *Mkdisk_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkdisk_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkdisk_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkdisk_params(s)
	}
}

func (s *Mkdisk_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkdisk_params(s)
	}
}

func (s *Mkdisk_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkdisk_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkdisk_params() (localctx IMkdisk_paramsContext) {
	localctx = NewMkdisk_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoDiskGrammarRULE_mkdisk_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(200)
			p.Mkdisk_param()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IFdisk_paramsContext is an interface to support dynamic dispatch.
type IFdisk_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFdisk_param() []IFdisk_paramContext
	Fdisk_param(i int) IFdisk_paramContext

	// IsFdisk_paramsContext differentiates from other interfaces.
	IsFdisk_paramsContext()
}

type Fdisk_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFdisk_paramsContext() *Fdisk_paramsContext {
	var p = new(Fdisk_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fdisk_params
	return p
}

func InitEmptyFdisk_paramsContext(p *Fdisk_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fdisk_params
}

func (*Fdisk_paramsContext) IsFdisk_paramsContext() {}

func NewFdisk_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fdisk_paramsContext {
	var p = new(Fdisk_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_fdisk_params

	return p
}

func (s *Fdisk_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Fdisk_paramsContext) AllFdisk_param() []IFdisk_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFdisk_paramContext); ok {
			len++
		}
	}

	tst := make([]IFdisk_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFdisk_paramContext); ok {
			tst[i] = t.(IFdisk_paramContext)
			i++
		}
	}

	return tst
}

func (s *Fdisk_paramsContext) Fdisk_param(i int) IFdisk_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFdisk_paramContext); ok {
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

	return t.(IFdisk_paramContext)
}

func (s *Fdisk_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fdisk_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fdisk_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFdisk_params(s)
	}
}

func (s *Fdisk_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFdisk_params(s)
	}
}

func (s *Fdisk_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFdisk_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Fdisk_params() (localctx IFdisk_paramsContext) {
	localctx = NewFdisk_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoDiskGrammarRULE_fdisk_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(205)
			p.Fdisk_param()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMount_paramsContext is an interface to support dynamic dispatch.
type IMount_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMount_param() []IMount_paramContext
	Mount_param(i int) IMount_paramContext

	// IsMount_paramsContext differentiates from other interfaces.
	IsMount_paramsContext()
}

type Mount_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMount_paramsContext() *Mount_paramsContext {
	var p = new(Mount_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mount_params
	return p
}

func InitEmptyMount_paramsContext(p *Mount_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mount_params
}

func (*Mount_paramsContext) IsMount_paramsContext() {}

func NewMount_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mount_paramsContext {
	var p = new(Mount_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mount_params

	return p
}

func (s *Mount_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Mount_paramsContext) AllMount_param() []IMount_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMount_paramContext); ok {
			len++
		}
	}

	tst := make([]IMount_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMount_paramContext); ok {
			tst[i] = t.(IMount_paramContext)
			i++
		}
	}

	return tst
}

func (s *Mount_paramsContext) Mount_param(i int) IMount_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMount_paramContext); ok {
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

	return t.(IMount_paramContext)
}

func (s *Mount_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mount_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mount_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMount_params(s)
	}
}

func (s *Mount_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMount_params(s)
	}
}

func (s *Mount_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMount_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mount_params() (localctx IMount_paramsContext) {
	localctx = NewMount_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoDiskGrammarRULE_mount_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(210)
			p.Mount_param()
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMkfs_paramsContext is an interface to support dynamic dispatch.
type IMkfs_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMkfs_param() []IMkfs_paramContext
	Mkfs_param(i int) IMkfs_paramContext

	// IsMkfs_paramsContext differentiates from other interfaces.
	IsMkfs_paramsContext()
}

type Mkfs_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfs_paramsContext() *Mkfs_paramsContext {
	var p = new(Mkfs_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfs_params
	return p
}

func InitEmptyMkfs_paramsContext(p *Mkfs_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfs_params
}

func (*Mkfs_paramsContext) IsMkfs_paramsContext() {}

func NewMkfs_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkfs_paramsContext {
	var p = new(Mkfs_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkfs_params

	return p
}

func (s *Mkfs_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkfs_paramsContext) AllMkfs_param() []IMkfs_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMkfs_paramContext); ok {
			len++
		}
	}

	tst := make([]IMkfs_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMkfs_paramContext); ok {
			tst[i] = t.(IMkfs_paramContext)
			i++
		}
	}

	return tst
}

func (s *Mkfs_paramsContext) Mkfs_param(i int) IMkfs_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfs_paramContext); ok {
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

	return t.(IMkfs_paramContext)
}

func (s *Mkfs_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkfs_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkfs_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkfs_params(s)
	}
}

func (s *Mkfs_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkfs_params(s)
	}
}

func (s *Mkfs_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkfs_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkfs_params() (localctx IMkfs_paramsContext) {
	localctx = NewMkfs_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoDiskGrammarRULE_mkfs_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(215)
			p.Mkfs_param()
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ILogin_paramsContext is an interface to support dynamic dispatch.
type ILogin_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllLogin_param() []ILogin_paramContext
	Login_param(i int) ILogin_paramContext

	// IsLogin_paramsContext differentiates from other interfaces.
	IsLogin_paramsContext()
}

type Login_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogin_paramsContext() *Login_paramsContext {
	var p = new(Login_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_login_params
	return p
}

func InitEmptyLogin_paramsContext(p *Login_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_login_params
}

func (*Login_paramsContext) IsLogin_paramsContext() {}

func NewLogin_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Login_paramsContext {
	var p = new(Login_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_login_params

	return p
}

func (s *Login_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Login_paramsContext) AllLogin_param() []ILogin_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILogin_paramContext); ok {
			len++
		}
	}

	tst := make([]ILogin_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILogin_paramContext); ok {
			tst[i] = t.(ILogin_paramContext)
			i++
		}
	}

	return tst
}

func (s *Login_paramsContext) Login_param(i int) ILogin_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILogin_paramContext); ok {
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

	return t.(ILogin_paramContext)
}

func (s *Login_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Login_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Login_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterLogin_params(s)
	}
}

func (s *Login_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitLogin_params(s)
	}
}

func (s *Login_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitLogin_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Login_params() (localctx ILogin_paramsContext) {
	localctx = NewLogin_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoDiskGrammarRULE_login_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(220)
			p.Login_param()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ICat_paramsContext is an interface to support dynamic dispatch.
type ICat_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCat_param() []ICat_paramContext
	Cat_param(i int) ICat_paramContext

	// IsCat_paramsContext differentiates from other interfaces.
	IsCat_paramsContext()
}

type Cat_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCat_paramsContext() *Cat_paramsContext {
	var p = new(Cat_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_cat_params
	return p
}

func InitEmptyCat_paramsContext(p *Cat_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_cat_params
}

func (*Cat_paramsContext) IsCat_paramsContext() {}

func NewCat_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cat_paramsContext {
	var p = new(Cat_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_cat_params

	return p
}

func (s *Cat_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Cat_paramsContext) AllCat_param() []ICat_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICat_paramContext); ok {
			len++
		}
	}

	tst := make([]ICat_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICat_paramContext); ok {
			tst[i] = t.(ICat_paramContext)
			i++
		}
	}

	return tst
}

func (s *Cat_paramsContext) Cat_param(i int) ICat_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICat_paramContext); ok {
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

	return t.(ICat_paramContext)
}

func (s *Cat_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cat_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cat_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCat_params(s)
	}
}

func (s *Cat_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCat_params(s)
	}
}

func (s *Cat_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCat_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Cat_params() (localctx ICat_paramsContext) {
	localctx = NewCat_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoDiskGrammarRULE_cat_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(225)
			p.Cat_param()
		}

		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMkusr_paramsContext is an interface to support dynamic dispatch.
type IMkusr_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMkusr_param() []IMkusr_paramContext
	Mkusr_param(i int) IMkusr_paramContext

	// IsMkusr_paramsContext differentiates from other interfaces.
	IsMkusr_paramsContext()
}

type Mkusr_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkusr_paramsContext() *Mkusr_paramsContext {
	var p = new(Mkusr_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkusr_params
	return p
}

func InitEmptyMkusr_paramsContext(p *Mkusr_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkusr_params
}

func (*Mkusr_paramsContext) IsMkusr_paramsContext() {}

func NewMkusr_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkusr_paramsContext {
	var p = new(Mkusr_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkusr_params

	return p
}

func (s *Mkusr_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkusr_paramsContext) AllMkusr_param() []IMkusr_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMkusr_paramContext); ok {
			len++
		}
	}

	tst := make([]IMkusr_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMkusr_paramContext); ok {
			tst[i] = t.(IMkusr_paramContext)
			i++
		}
	}

	return tst
}

func (s *Mkusr_paramsContext) Mkusr_param(i int) IMkusr_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkusr_paramContext); ok {
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

	return t.(IMkusr_paramContext)
}

func (s *Mkusr_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkusr_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkusr_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkusr_params(s)
	}
}

func (s *Mkusr_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkusr_params(s)
	}
}

func (s *Mkusr_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkusr_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkusr_params() (localctx IMkusr_paramsContext) {
	localctx = NewMkusr_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoDiskGrammarRULE_mkusr_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(230)
			p.Mkusr_param()
		}

		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IChgrp_paramsContext is an interface to support dynamic dispatch.
type IChgrp_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChgrp_param() []IChgrp_paramContext
	Chgrp_param(i int) IChgrp_paramContext

	// IsChgrp_paramsContext differentiates from other interfaces.
	IsChgrp_paramsContext()
}

type Chgrp_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChgrp_paramsContext() *Chgrp_paramsContext {
	var p = new(Chgrp_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chgrp_params
	return p
}

func InitEmptyChgrp_paramsContext(p *Chgrp_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chgrp_params
}

func (*Chgrp_paramsContext) IsChgrp_paramsContext() {}

func NewChgrp_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chgrp_paramsContext {
	var p = new(Chgrp_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_chgrp_params

	return p
}

func (s *Chgrp_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Chgrp_paramsContext) AllChgrp_param() []IChgrp_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChgrp_paramContext); ok {
			len++
		}
	}

	tst := make([]IChgrp_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChgrp_paramContext); ok {
			tst[i] = t.(IChgrp_paramContext)
			i++
		}
	}

	return tst
}

func (s *Chgrp_paramsContext) Chgrp_param(i int) IChgrp_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChgrp_paramContext); ok {
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

	return t.(IChgrp_paramContext)
}

func (s *Chgrp_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chgrp_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chgrp_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterChgrp_params(s)
	}
}

func (s *Chgrp_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitChgrp_params(s)
	}
}

func (s *Chgrp_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitChgrp_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Chgrp_params() (localctx IChgrp_paramsContext) {
	localctx = NewChgrp_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoDiskGrammarRULE_chgrp_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(235)
			p.Chgrp_param()
		}

		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMkfile_paramsContext is an interface to support dynamic dispatch.
type IMkfile_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMkfile_param() []IMkfile_paramContext
	Mkfile_param(i int) IMkfile_paramContext

	// IsMkfile_paramsContext differentiates from other interfaces.
	IsMkfile_paramsContext()
}

type Mkfile_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfile_paramsContext() *Mkfile_paramsContext {
	var p = new(Mkfile_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfile_params
	return p
}

func InitEmptyMkfile_paramsContext(p *Mkfile_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfile_params
}

func (*Mkfile_paramsContext) IsMkfile_paramsContext() {}

func NewMkfile_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkfile_paramsContext {
	var p = new(Mkfile_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkfile_params

	return p
}

func (s *Mkfile_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkfile_paramsContext) AllMkfile_param() []IMkfile_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMkfile_paramContext); ok {
			len++
		}
	}

	tst := make([]IMkfile_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMkfile_paramContext); ok {
			tst[i] = t.(IMkfile_paramContext)
			i++
		}
	}

	return tst
}

func (s *Mkfile_paramsContext) Mkfile_param(i int) IMkfile_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkfile_paramContext); ok {
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

	return t.(IMkfile_paramContext)
}

func (s *Mkfile_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkfile_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkfile_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkfile_params(s)
	}
}

func (s *Mkfile_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkfile_params(s)
	}
}

func (s *Mkfile_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkfile_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkfile_params() (localctx IMkfile_paramsContext) {
	localctx = NewMkfile_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GoDiskGrammarRULE_mkfile_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(240)
			p.Mkfile_param()
		}

		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMkdir_paramsContext is an interface to support dynamic dispatch.
type IMkdir_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMkdir_param() []IMkdir_paramContext
	Mkdir_param(i int) IMkdir_paramContext

	// IsMkdir_paramsContext differentiates from other interfaces.
	IsMkdir_paramsContext()
}

type Mkdir_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdir_paramsContext() *Mkdir_paramsContext {
	var p = new(Mkdir_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdir_params
	return p
}

func InitEmptyMkdir_paramsContext(p *Mkdir_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdir_params
}

func (*Mkdir_paramsContext) IsMkdir_paramsContext() {}

func NewMkdir_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkdir_paramsContext {
	var p = new(Mkdir_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkdir_params

	return p
}

func (s *Mkdir_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkdir_paramsContext) AllMkdir_param() []IMkdir_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMkdir_paramContext); ok {
			len++
		}
	}

	tst := make([]IMkdir_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMkdir_paramContext); ok {
			tst[i] = t.(IMkdir_paramContext)
			i++
		}
	}

	return tst
}

func (s *Mkdir_paramsContext) Mkdir_param(i int) IMkdir_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMkdir_paramContext); ok {
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

	return t.(IMkdir_paramContext)
}

func (s *Mkdir_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkdir_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkdir_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkdir_params(s)
	}
}

func (s *Mkdir_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkdir_params(s)
	}
}

func (s *Mkdir_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkdir_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkdir_params() (localctx IMkdir_paramsContext) {
	localctx = NewMkdir_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GoDiskGrammarRULE_mkdir_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(245)
			p.Mkdir_param()
		}

		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IRep_paramsContext is an interface to support dynamic dispatch.
type IRep_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRep_param() []IRep_paramContext
	Rep_param(i int) IRep_paramContext

	// IsRep_paramsContext differentiates from other interfaces.
	IsRep_paramsContext()
}

type Rep_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRep_paramsContext() *Rep_paramsContext {
	var p = new(Rep_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rep_params
	return p
}

func InitEmptyRep_paramsContext(p *Rep_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rep_params
}

func (*Rep_paramsContext) IsRep_paramsContext() {}

func NewRep_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rep_paramsContext {
	var p = new(Rep_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rep_params

	return p
}

func (s *Rep_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Rep_paramsContext) AllRep_param() []IRep_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRep_paramContext); ok {
			len++
		}
	}

	tst := make([]IRep_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRep_paramContext); ok {
			tst[i] = t.(IRep_paramContext)
			i++
		}
	}

	return tst
}

func (s *Rep_paramsContext) Rep_param(i int) IRep_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRep_paramContext); ok {
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

	return t.(IRep_paramContext)
}

func (s *Rep_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rep_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rep_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRep_params(s)
	}
}

func (s *Rep_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRep_params(s)
	}
}

func (s *Rep_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRep_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rep_params() (localctx IRep_paramsContext) {
	localctx = NewRep_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GoDiskGrammarRULE_rep_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(250)
			p.Rep_param()
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// IMkdisk_paramContext is an interface to support dynamic dispatch.
type IMkdisk_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Size() ISizeContext
	Fit() IFitContext
	Unit() IUnitContext
	Path() IPathContext

	// IsMkdisk_paramContext differentiates from other interfaces.
	IsMkdisk_paramContext()
}

type Mkdisk_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdisk_paramContext() *Mkdisk_paramContext {
	var p = new(Mkdisk_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdisk_param
	return p
}

func InitEmptyMkdisk_paramContext(p *Mkdisk_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdisk_param
}

func (*Mkdisk_paramContext) IsMkdisk_paramContext() {}

func NewMkdisk_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkdisk_paramContext {
	var p = new(Mkdisk_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkdisk_param

	return p
}

func (s *Mkdisk_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkdisk_paramContext) Size() ISizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *Mkdisk_paramContext) Fit() IFitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFitContext)
}

func (s *Mkdisk_paramContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *Mkdisk_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Mkdisk_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkdisk_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkdisk_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkdisk_param(s)
	}
}

func (s *Mkdisk_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkdisk_param(s)
	}
}

func (s *Mkdisk_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkdisk_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkdisk_param() (localctx IMkdisk_paramContext) {
	localctx = NewMkdisk_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GoDiskGrammarRULE_mkdisk_param)
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(255)
			p.Size()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(256)
			p.Fit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(257)
			p.Unit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(258)
			p.Path()
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

// IRmdisk_paramContext is an interface to support dynamic dispatch.
type IRmdisk_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext

	// IsRmdisk_paramContext differentiates from other interfaces.
	IsRmdisk_paramContext()
}

type Rmdisk_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmdisk_paramContext() *Rmdisk_paramContext {
	var p = new(Rmdisk_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rmdisk_param
	return p
}

func InitEmptyRmdisk_paramContext(p *Rmdisk_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rmdisk_param
}

func (*Rmdisk_paramContext) IsRmdisk_paramContext() {}

func NewRmdisk_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rmdisk_paramContext {
	var p = new(Rmdisk_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rmdisk_param

	return p
}

func (s *Rmdisk_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Rmdisk_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Rmdisk_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rmdisk_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rmdisk_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRmdisk_param(s)
	}
}

func (s *Rmdisk_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRmdisk_param(s)
	}
}

func (s *Rmdisk_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRmdisk_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rmdisk_param() (localctx IRmdisk_paramContext) {
	localctx = NewRmdisk_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GoDiskGrammarRULE_rmdisk_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Path()
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

// IFdisk_paramContext is an interface to support dynamic dispatch.
type IFdisk_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Size() ISizeContext
	Fit() IFitContext
	Unit() IUnitContext
	Path() IPathContext
	Type_() ITypeContext
	Name() INameContext

	// IsFdisk_paramContext differentiates from other interfaces.
	IsFdisk_paramContext()
}

type Fdisk_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFdisk_paramContext() *Fdisk_paramContext {
	var p = new(Fdisk_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fdisk_param
	return p
}

func InitEmptyFdisk_paramContext(p *Fdisk_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fdisk_param
}

func (*Fdisk_paramContext) IsFdisk_paramContext() {}

func NewFdisk_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fdisk_paramContext {
	var p = new(Fdisk_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_fdisk_param

	return p
}

func (s *Fdisk_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Fdisk_paramContext) Size() ISizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *Fdisk_paramContext) Fit() IFitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFitContext)
}

func (s *Fdisk_paramContext) Unit() IUnitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnitContext)
}

func (s *Fdisk_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Fdisk_paramContext) Type_() ITypeContext {
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

func (s *Fdisk_paramContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Fdisk_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fdisk_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fdisk_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFdisk_param(s)
	}
}

func (s *Fdisk_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFdisk_param(s)
	}
}

func (s *Fdisk_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFdisk_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Fdisk_param() (localctx IFdisk_paramContext) {
	localctx = NewFdisk_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GoDiskGrammarRULE_fdisk_param)
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Size()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Fit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.Unit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(266)
			p.Path()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(267)
			p.Type_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(268)
			p.Name()
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

// IMount_paramContext is an interface to support dynamic dispatch.
type IMount_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Name() INameContext

	// IsMount_paramContext differentiates from other interfaces.
	IsMount_paramContext()
}

type Mount_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMount_paramContext() *Mount_paramContext {
	var p = new(Mount_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mount_param
	return p
}

func InitEmptyMount_paramContext(p *Mount_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mount_param
}

func (*Mount_paramContext) IsMount_paramContext() {}

func NewMount_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mount_paramContext {
	var p = new(Mount_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mount_param

	return p
}

func (s *Mount_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mount_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Mount_paramContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Mount_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mount_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mount_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMount_param(s)
	}
}

func (s *Mount_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMount_param(s)
	}
}

func (s *Mount_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMount_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mount_param() (localctx IMount_paramContext) {
	localctx = NewMount_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GoDiskGrammarRULE_mount_param)
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.Name()
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

// IMkfs_paramContext is an interface to support dynamic dispatch.
type IMkfs_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_text() IId_textContext
	Type_() ITypeContext

	// IsMkfs_paramContext differentiates from other interfaces.
	IsMkfs_paramContext()
}

type Mkfs_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfs_paramContext() *Mkfs_paramContext {
	var p = new(Mkfs_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfs_param
	return p
}

func InitEmptyMkfs_paramContext(p *Mkfs_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfs_param
}

func (*Mkfs_paramContext) IsMkfs_paramContext() {}

func NewMkfs_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkfs_paramContext {
	var p = new(Mkfs_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkfs_param

	return p
}

func (s *Mkfs_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkfs_paramContext) Id_text() IId_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_textContext)
}

func (s *Mkfs_paramContext) Type_() ITypeContext {
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

func (s *Mkfs_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkfs_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkfs_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkfs_param(s)
	}
}

func (s *Mkfs_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkfs_param(s)
	}
}

func (s *Mkfs_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkfs_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkfs_param() (localctx IMkfs_paramContext) {
	localctx = NewMkfs_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GoDiskGrammarRULE_mkfs_param)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Id_text()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(276)
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

// ICat_paramContext is an interface to support dynamic dispatch.
type ICat_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Filen() IFilenContext

	// IsCat_paramContext differentiates from other interfaces.
	IsCat_paramContext()
}

type Cat_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCat_paramContext() *Cat_paramContext {
	var p = new(Cat_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_cat_param
	return p
}

func InitEmptyCat_paramContext(p *Cat_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_cat_param
}

func (*Cat_paramContext) IsCat_paramContext() {}

func NewCat_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cat_paramContext {
	var p = new(Cat_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_cat_param

	return p
}

func (s *Cat_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Cat_paramContext) Filen() IFilenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilenContext)
}

func (s *Cat_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cat_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cat_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCat_param(s)
	}
}

func (s *Cat_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCat_param(s)
	}
}

func (s *Cat_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCat_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Cat_param() (localctx ICat_paramContext) {
	localctx = NewCat_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GoDiskGrammarRULE_cat_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Filen()
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

// ILogin_paramContext is an interface to support dynamic dispatch.
type ILogin_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	User() IUserContext
	Pass() IPassContext
	Id_text() IId_textContext

	// IsLogin_paramContext differentiates from other interfaces.
	IsLogin_paramContext()
}

type Login_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogin_paramContext() *Login_paramContext {
	var p = new(Login_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_login_param
	return p
}

func InitEmptyLogin_paramContext(p *Login_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_login_param
}

func (*Login_paramContext) IsLogin_paramContext() {}

func NewLogin_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Login_paramContext {
	var p = new(Login_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_login_param

	return p
}

func (s *Login_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Login_paramContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *Login_paramContext) Pass() IPassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPassContext)
}

func (s *Login_paramContext) Id_text() IId_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_textContext)
}

func (s *Login_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Login_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Login_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterLogin_param(s)
	}
}

func (s *Login_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitLogin_param(s)
	}
}

func (s *Login_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitLogin_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Login_param() (localctx ILogin_paramContext) {
	localctx = NewLogin_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoDiskGrammarRULE_login_param)
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.User()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Pass()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(283)
			p.Id_text()
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

// IMkgrp_paramContext is an interface to support dynamic dispatch.
type IMkgrp_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsMkgrp_paramContext differentiates from other interfaces.
	IsMkgrp_paramContext()
}

type Mkgrp_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkgrp_paramContext() *Mkgrp_paramContext {
	var p = new(Mkgrp_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkgrp_param
	return p
}

func InitEmptyMkgrp_paramContext(p *Mkgrp_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkgrp_param
}

func (*Mkgrp_paramContext) IsMkgrp_paramContext() {}

func NewMkgrp_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkgrp_paramContext {
	var p = new(Mkgrp_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkgrp_param

	return p
}

func (s *Mkgrp_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkgrp_paramContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Mkgrp_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkgrp_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkgrp_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkgrp_param(s)
	}
}

func (s *Mkgrp_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkgrp_param(s)
	}
}

func (s *Mkgrp_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkgrp_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkgrp_param() (localctx IMkgrp_paramContext) {
	localctx = NewMkgrp_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoDiskGrammarRULE_mkgrp_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(286)
		p.Name()
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

// IRmgrp_paramContext is an interface to support dynamic dispatch.
type IRmgrp_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext

	// IsRmgrp_paramContext differentiates from other interfaces.
	IsRmgrp_paramContext()
}

type Rmgrp_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmgrp_paramContext() *Rmgrp_paramContext {
	var p = new(Rmgrp_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rmgrp_param
	return p
}

func InitEmptyRmgrp_paramContext(p *Rmgrp_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rmgrp_param
}

func (*Rmgrp_paramContext) IsRmgrp_paramContext() {}

func NewRmgrp_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rmgrp_paramContext {
	var p = new(Rmgrp_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rmgrp_param

	return p
}

func (s *Rmgrp_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Rmgrp_paramContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Rmgrp_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rmgrp_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rmgrp_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRmgrp_param(s)
	}
}

func (s *Rmgrp_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRmgrp_param(s)
	}
}

func (s *Rmgrp_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRmgrp_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rmgrp_param() (localctx IRmgrp_paramContext) {
	localctx = NewRmgrp_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoDiskGrammarRULE_rmgrp_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(288)
		p.Name()
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

// IMkusr_paramContext is an interface to support dynamic dispatch.
type IMkusr_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	User() IUserContext
	Pass() IPassContext
	Grp() IGrpContext

	// IsMkusr_paramContext differentiates from other interfaces.
	IsMkusr_paramContext()
}

type Mkusr_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkusr_paramContext() *Mkusr_paramContext {
	var p = new(Mkusr_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkusr_param
	return p
}

func InitEmptyMkusr_paramContext(p *Mkusr_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkusr_param
}

func (*Mkusr_paramContext) IsMkusr_paramContext() {}

func NewMkusr_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkusr_paramContext {
	var p = new(Mkusr_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkusr_param

	return p
}

func (s *Mkusr_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkusr_paramContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *Mkusr_paramContext) Pass() IPassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPassContext)
}

func (s *Mkusr_paramContext) Grp() IGrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrpContext)
}

func (s *Mkusr_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkusr_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkusr_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkusr_param(s)
	}
}

func (s *Mkusr_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkusr_param(s)
	}
}

func (s *Mkusr_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkusr_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkusr_param() (localctx IMkusr_paramContext) {
	localctx = NewMkusr_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoDiskGrammarRULE_mkusr_param)
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.User()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Pass()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(292)
			p.Grp()
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

// IRmusr_paramContext is an interface to support dynamic dispatch.
type IRmusr_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	User() IUserContext

	// IsRmusr_paramContext differentiates from other interfaces.
	IsRmusr_paramContext()
}

type Rmusr_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRmusr_paramContext() *Rmusr_paramContext {
	var p = new(Rmusr_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rmusr_param
	return p
}

func InitEmptyRmusr_paramContext(p *Rmusr_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rmusr_param
}

func (*Rmusr_paramContext) IsRmusr_paramContext() {}

func NewRmusr_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rmusr_paramContext {
	var p = new(Rmusr_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rmusr_param

	return p
}

func (s *Rmusr_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Rmusr_paramContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *Rmusr_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rmusr_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rmusr_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRmusr_param(s)
	}
}

func (s *Rmusr_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRmusr_param(s)
	}
}

func (s *Rmusr_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRmusr_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rmusr_param() (localctx IRmusr_paramContext) {
	localctx = NewRmusr_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoDiskGrammarRULE_rmusr_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.User()
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

// IChgrp_paramContext is an interface to support dynamic dispatch.
type IChgrp_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	User() IUserContext
	Grp() IGrpContext

	// IsChgrp_paramContext differentiates from other interfaces.
	IsChgrp_paramContext()
}

type Chgrp_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChgrp_paramContext() *Chgrp_paramContext {
	var p = new(Chgrp_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chgrp_param
	return p
}

func InitEmptyChgrp_paramContext(p *Chgrp_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chgrp_param
}

func (*Chgrp_paramContext) IsChgrp_paramContext() {}

func NewChgrp_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chgrp_paramContext {
	var p = new(Chgrp_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_chgrp_param

	return p
}

func (s *Chgrp_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Chgrp_paramContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *Chgrp_paramContext) Grp() IGrpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGrpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGrpContext)
}

func (s *Chgrp_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chgrp_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chgrp_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterChgrp_param(s)
	}
}

func (s *Chgrp_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitChgrp_param(s)
	}
}

func (s *Chgrp_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitChgrp_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Chgrp_param() (localctx IChgrp_paramContext) {
	localctx = NewChgrp_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoDiskGrammarRULE_chgrp_param)
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.User()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.Grp()
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

// IMkfile_paramContext is an interface to support dynamic dispatch.
type IMkfile_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	R() IRContext
	Size() ISizeContext
	Cont() IContContext

	// IsMkfile_paramContext differentiates from other interfaces.
	IsMkfile_paramContext()
}

type Mkfile_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkfile_paramContext() *Mkfile_paramContext {
	var p = new(Mkfile_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfile_param
	return p
}

func InitEmptyMkfile_paramContext(p *Mkfile_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkfile_param
}

func (*Mkfile_paramContext) IsMkfile_paramContext() {}

func NewMkfile_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkfile_paramContext {
	var p = new(Mkfile_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkfile_param

	return p
}

func (s *Mkfile_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkfile_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Mkfile_paramContext) R() IRContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRContext)
}

func (s *Mkfile_paramContext) Size() ISizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *Mkfile_paramContext) Cont() IContContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContContext)
}

func (s *Mkfile_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkfile_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkfile_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkfile_param(s)
	}
}

func (s *Mkfile_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkfile_param(s)
	}
}

func (s *Mkfile_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkfile_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkfile_param() (localctx IMkfile_paramContext) {
	localctx = NewMkfile_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoDiskGrammarRULE_mkfile_param)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.R()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(303)
			p.Size()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(304)
			p.Cont()
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

// IMkdir_paramContext is an interface to support dynamic dispatch.
type IMkdir_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	P() IPContext
	Path() IPathContext

	// IsMkdir_paramContext differentiates from other interfaces.
	IsMkdir_paramContext()
}

type Mkdir_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMkdir_paramContext() *Mkdir_paramContext {
	var p = new(Mkdir_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdir_param
	return p
}

func InitEmptyMkdir_paramContext(p *Mkdir_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_mkdir_param
}

func (*Mkdir_paramContext) IsMkdir_paramContext() {}

func NewMkdir_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mkdir_paramContext {
	var p = new(Mkdir_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_mkdir_param

	return p
}

func (s *Mkdir_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Mkdir_paramContext) P() IPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPContext)
}

func (s *Mkdir_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Mkdir_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mkdir_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mkdir_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMkdir_param(s)
	}
}

func (s *Mkdir_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMkdir_param(s)
	}
}

func (s *Mkdir_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMkdir_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Mkdir_param() (localctx IMkdir_paramContext) {
	localctx = NewMkdir_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GoDiskGrammarRULE_mkdir_param)
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.P()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Path()
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

// IRep_paramContext is an interface to support dynamic dispatch.
type IRep_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Path() IPathContext
	Id_text() IId_textContext
	Path_file_ls() IPath_file_lsContext

	// IsRep_paramContext differentiates from other interfaces.
	IsRep_paramContext()
}

type Rep_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRep_paramContext() *Rep_paramContext {
	var p = new(Rep_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rep_param
	return p
}

func InitEmptyRep_paramContext(p *Rep_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rep_param
}

func (*Rep_paramContext) IsRep_paramContext() {}

func NewRep_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rep_paramContext {
	var p = new(Rep_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rep_param

	return p
}

func (s *Rep_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Rep_paramContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Rep_paramContext) Path() IPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPathContext)
}

func (s *Rep_paramContext) Id_text() IId_textContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IId_textContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IId_textContext)
}

func (s *Rep_paramContext) Path_file_ls() IPath_file_lsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPath_file_lsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPath_file_lsContext)
}

func (s *Rep_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rep_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rep_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRep_param(s)
	}
}

func (s *Rep_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRep_param(s)
	}
}

func (s *Rep_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRep_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rep_param() (localctx IRep_paramContext) {
	localctx = NewRep_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GoDiskGrammarRULE_rep_param)
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
			p.Path()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(313)
			p.Id_text()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(314)
			p.Path_file_ls()
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
