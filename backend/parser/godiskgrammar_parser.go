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
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "'='", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "INT_LIT", "FLOAT_LIT", "STRING_LIT", "CHAR_LIT", "MKDISK", "RMDISK",
		"FDISK", "MOUNT", "MOUNTED", "UNMOUNT", "MKFS", "CAT", "LOGIN", "LOGOUT",
		"MKGRP", "RMGRP", "MKUSR", "RMUSR", "CHGRP", "MKFILE", "MKDIR", "REMOVE",
		"EDIT", "RENAME", "COPY", "MOVE", "FIND", "CHOWN", "CHMOD", "RECOVERY",
		"LOSS", "JOURNALING", "REP", "ID_TEXT", "SIZE", "FIT", "UNIT", "PATH",
		"TYPE", "NAME", "FILE", "USER", "PASS", "GRP", "CONT", "PATH_FILE_LS",
		"DELETE", "ADD", "FS", "CONTENIDO", "DESTINO", "USUARIO", "UGO", "R",
		"P", "ASSIGN", "MINUS", "ID", "UNQUOTED_TEXT", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"prog", "stmt", "size", "fit", "unit", "type", "id_text", "path", "name",
		"filen", "user", "pass", "grp", "cont", "path_file_ls", "delete", "add",
		"fs", "contenido", "destino", "usuario", "ugo", "p", "r", "mkdisk_params",
		"fdisk_params", "mount_params", "mkfs_params", "login_params", "cat_params",
		"mkusr_params", "chgrp_params", "mkfile_params", "mkdir_params", "edit_params",
		"rename_params", "copy_params", "move_params", "find_params", "chown_params",
		"chmod_params", "rep_params", "mkdisk_param", "rmdisk_param", "fdisk_param",
		"mount_param", "unmount_param", "mkfs_param", "cat_param", "login_param",
		"mkgrp_param", "rmgrp_param", "mkusr_param", "rmusr_param", "chgrp_param",
		"mkfile_param", "mkdir_param", "remove_param", "edit_param", "rename_param",
		"copy_param", "move_param", "find_param", "chown_param", "chmod_param",
		"recovery_param", "loss_param", "journaling_param", "rep_param",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 62, 507, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 1, 0, 5, 0, 140, 8, 0, 10, 0, 12, 0, 143, 9, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 203, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 4, 24, 313, 8, 24, 11, 24,
		12, 24, 314, 1, 25, 4, 25, 318, 8, 25, 11, 25, 12, 25, 319, 1, 26, 4, 26,
		323, 8, 26, 11, 26, 12, 26, 324, 1, 27, 4, 27, 328, 8, 27, 11, 27, 12,
		27, 329, 1, 28, 4, 28, 333, 8, 28, 11, 28, 12, 28, 334, 1, 29, 4, 29, 338,
		8, 29, 11, 29, 12, 29, 339, 1, 30, 4, 30, 343, 8, 30, 11, 30, 12, 30, 344,
		1, 31, 4, 31, 348, 8, 31, 11, 31, 12, 31, 349, 1, 32, 4, 32, 353, 8, 32,
		11, 32, 12, 32, 354, 1, 33, 4, 33, 358, 8, 33, 11, 33, 12, 33, 359, 1,
		34, 4, 34, 363, 8, 34, 11, 34, 12, 34, 364, 1, 35, 4, 35, 368, 8, 35, 11,
		35, 12, 35, 369, 1, 36, 4, 36, 373, 8, 36, 11, 36, 12, 36, 374, 1, 37,
		4, 37, 378, 8, 37, 11, 37, 12, 37, 379, 1, 38, 4, 38, 383, 8, 38, 11, 38,
		12, 38, 384, 1, 39, 4, 39, 388, 8, 39, 11, 39, 12, 39, 389, 1, 40, 4, 40,
		393, 8, 40, 11, 40, 12, 40, 394, 1, 41, 4, 41, 398, 8, 41, 11, 41, 12,
		41, 399, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 406, 8, 42, 1, 43, 1, 43, 1,
		44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 418, 8, 44,
		1, 45, 1, 45, 3, 45, 422, 8, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 3,
		47, 429, 8, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 3, 49, 436, 8, 49, 1,
		50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 3, 52, 445, 8, 52, 1, 53,
		1, 53, 1, 54, 1, 54, 3, 54, 451, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 3,
		55, 457, 8, 55, 1, 56, 1, 56, 3, 56, 461, 8, 56, 1, 57, 1, 57, 1, 58, 1,
		58, 3, 58, 467, 8, 58, 1, 59, 1, 59, 3, 59, 471, 8, 59, 1, 60, 1, 60, 3,
		60, 475, 8, 60, 1, 61, 1, 61, 3, 61, 479, 8, 61, 1, 62, 1, 62, 3, 62, 483,
		8, 62, 1, 63, 1, 63, 1, 63, 3, 63, 488, 8, 63, 1, 64, 1, 64, 1, 64, 3,
		64, 493, 8, 64, 1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 68, 1, 68,
		1, 68, 1, 68, 3, 68, 505, 8, 68, 1, 68, 0, 0, 69, 0, 2, 4, 6, 8, 10, 12,
		14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48,
		50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84,
		86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116,
		118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 0, 1, 2, 0, 3, 3, 59,
		59, 518, 0, 141, 1, 0, 0, 0, 2, 202, 1, 0, 0, 0, 4, 204, 1, 0, 0, 0, 6,
		209, 1, 0, 0, 0, 8, 214, 1, 0, 0, 0, 10, 219, 1, 0, 0, 0, 12, 224, 1, 0,
		0, 0, 14, 229, 1, 0, 0, 0, 16, 234, 1, 0, 0, 0, 18, 239, 1, 0, 0, 0, 20,
		245, 1, 0, 0, 0, 22, 250, 1, 0, 0, 0, 24, 255, 1, 0, 0, 0, 26, 260, 1,
		0, 0, 0, 28, 265, 1, 0, 0, 0, 30, 270, 1, 0, 0, 0, 32, 275, 1, 0, 0, 0,
		34, 280, 1, 0, 0, 0, 36, 285, 1, 0, 0, 0, 38, 290, 1, 0, 0, 0, 40, 295,
		1, 0, 0, 0, 42, 300, 1, 0, 0, 0, 44, 305, 1, 0, 0, 0, 46, 308, 1, 0, 0,
		0, 48, 312, 1, 0, 0, 0, 50, 317, 1, 0, 0, 0, 52, 322, 1, 0, 0, 0, 54, 327,
		1, 0, 0, 0, 56, 332, 1, 0, 0, 0, 58, 337, 1, 0, 0, 0, 60, 342, 1, 0, 0,
		0, 62, 347, 1, 0, 0, 0, 64, 352, 1, 0, 0, 0, 66, 357, 1, 0, 0, 0, 68, 362,
		1, 0, 0, 0, 70, 367, 1, 0, 0, 0, 72, 372, 1, 0, 0, 0, 74, 377, 1, 0, 0,
		0, 76, 382, 1, 0, 0, 0, 78, 387, 1, 0, 0, 0, 80, 392, 1, 0, 0, 0, 82, 397,
		1, 0, 0, 0, 84, 405, 1, 0, 0, 0, 86, 407, 1, 0, 0, 0, 88, 417, 1, 0, 0,
		0, 90, 421, 1, 0, 0, 0, 92, 423, 1, 0, 0, 0, 94, 428, 1, 0, 0, 0, 96, 430,
		1, 0, 0, 0, 98, 435, 1, 0, 0, 0, 100, 437, 1, 0, 0, 0, 102, 439, 1, 0,
		0, 0, 104, 444, 1, 0, 0, 0, 106, 446, 1, 0, 0, 0, 108, 450, 1, 0, 0, 0,
		110, 456, 1, 0, 0, 0, 112, 460, 1, 0, 0, 0, 114, 462, 1, 0, 0, 0, 116,
		466, 1, 0, 0, 0, 118, 470, 1, 0, 0, 0, 120, 474, 1, 0, 0, 0, 122, 478,
		1, 0, 0, 0, 124, 482, 1, 0, 0, 0, 126, 487, 1, 0, 0, 0, 128, 492, 1, 0,
		0, 0, 130, 494, 1, 0, 0, 0, 132, 496, 1, 0, 0, 0, 134, 498, 1, 0, 0, 0,
		136, 504, 1, 0, 0, 0, 138, 140, 3, 2, 1, 0, 139, 138, 1, 0, 0, 0, 140,
		143, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 144,
		1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 145, 5, 0, 0, 1, 145, 1, 1, 0, 0,
		0, 146, 147, 5, 5, 0, 0, 147, 203, 3, 48, 24, 0, 148, 149, 5, 6, 0, 0,
		149, 203, 3, 86, 43, 0, 150, 151, 5, 7, 0, 0, 151, 203, 3, 50, 25, 0, 152,
		153, 5, 8, 0, 0, 153, 203, 3, 52, 26, 0, 154, 203, 5, 9, 0, 0, 155, 156,
		5, 10, 0, 0, 156, 203, 3, 92, 46, 0, 157, 158, 5, 11, 0, 0, 158, 203, 3,
		54, 27, 0, 159, 160, 5, 12, 0, 0, 160, 203, 3, 58, 29, 0, 161, 162, 5,
		13, 0, 0, 162, 203, 3, 56, 28, 0, 163, 203, 5, 14, 0, 0, 164, 165, 5, 15,
		0, 0, 165, 203, 3, 100, 50, 0, 166, 167, 5, 16, 0, 0, 167, 203, 3, 102,
		51, 0, 168, 169, 5, 17, 0, 0, 169, 203, 3, 60, 30, 0, 170, 171, 5, 18,
		0, 0, 171, 203, 3, 106, 53, 0, 172, 173, 5, 19, 0, 0, 173, 203, 3, 62,
		31, 0, 174, 175, 5, 20, 0, 0, 175, 203, 3, 64, 32, 0, 176, 177, 5, 21,
		0, 0, 177, 203, 3, 66, 33, 0, 178, 179, 5, 22, 0, 0, 179, 203, 3, 114,
		57, 0, 180, 181, 5, 23, 0, 0, 181, 203, 3, 68, 34, 0, 182, 183, 5, 24,
		0, 0, 183, 203, 3, 70, 35, 0, 184, 185, 5, 25, 0, 0, 185, 203, 3, 72, 36,
		0, 186, 187, 5, 26, 0, 0, 187, 203, 3, 74, 37, 0, 188, 189, 5, 27, 0, 0,
		189, 203, 3, 76, 38, 0, 190, 191, 5, 28, 0, 0, 191, 203, 3, 78, 39, 0,
		192, 193, 5, 29, 0, 0, 193, 203, 3, 80, 40, 0, 194, 195, 5, 30, 0, 0, 195,
		203, 3, 130, 65, 0, 196, 197, 5, 31, 0, 0, 197, 203, 3, 132, 66, 0, 198,
		199, 5, 32, 0, 0, 199, 203, 3, 134, 67, 0, 200, 201, 5, 33, 0, 0, 201,
		203, 3, 82, 41, 0, 202, 146, 1, 0, 0, 0, 202, 148, 1, 0, 0, 0, 202, 150,
		1, 0, 0, 0, 202, 152, 1, 0, 0, 0, 202, 154, 1, 0, 0, 0, 202, 155, 1, 0,
		0, 0, 202, 157, 1, 0, 0, 0, 202, 159, 1, 0, 0, 0, 202, 161, 1, 0, 0, 0,
		202, 163, 1, 0, 0, 0, 202, 164, 1, 0, 0, 0, 202, 166, 1, 0, 0, 0, 202,
		168, 1, 0, 0, 0, 202, 170, 1, 0, 0, 0, 202, 172, 1, 0, 0, 0, 202, 174,
		1, 0, 0, 0, 202, 176, 1, 0, 0, 0, 202, 178, 1, 0, 0, 0, 202, 180, 1, 0,
		0, 0, 202, 182, 1, 0, 0, 0, 202, 184, 1, 0, 0, 0, 202, 186, 1, 0, 0, 0,
		202, 188, 1, 0, 0, 0, 202, 190, 1, 0, 0, 0, 202, 192, 1, 0, 0, 0, 202,
		194, 1, 0, 0, 0, 202, 196, 1, 0, 0, 0, 202, 198, 1, 0, 0, 0, 202, 200,
		1, 0, 0, 0, 203, 3, 1, 0, 0, 0, 204, 205, 5, 57, 0, 0, 205, 206, 5, 35,
		0, 0, 206, 207, 5, 56, 0, 0, 207, 208, 5, 1, 0, 0, 208, 5, 1, 0, 0, 0,
		209, 210, 5, 57, 0, 0, 210, 211, 5, 36, 0, 0, 211, 212, 5, 56, 0, 0, 212,
		213, 5, 58, 0, 0, 213, 7, 1, 0, 0, 0, 214, 215, 5, 57, 0, 0, 215, 216,
		5, 37, 0, 0, 216, 217, 5, 56, 0, 0, 217, 218, 5, 58, 0, 0, 218, 9, 1, 0,
		0, 0, 219, 220, 5, 57, 0, 0, 220, 221, 5, 39, 0, 0, 221, 222, 5, 56, 0,
		0, 222, 223, 5, 58, 0, 0, 223, 11, 1, 0, 0, 0, 224, 225, 5, 57, 0, 0, 225,
		226, 5, 34, 0, 0, 226, 227, 5, 56, 0, 0, 227, 228, 5, 58, 0, 0, 228, 13,
		1, 0, 0, 0, 229, 230, 5, 57, 0, 0, 230, 231, 5, 38, 0, 0, 231, 232, 5,
		56, 0, 0, 232, 233, 7, 0, 0, 0, 233, 15, 1, 0, 0, 0, 234, 235, 5, 57, 0,
		0, 235, 236, 5, 40, 0, 0, 236, 237, 5, 56, 0, 0, 237, 238, 7, 0, 0, 0,
		238, 17, 1, 0, 0, 0, 239, 240, 5, 57, 0, 0, 240, 241, 5, 41, 0, 0, 241,
		242, 5, 1, 0, 0, 242, 243, 5, 56, 0, 0, 243, 244, 7, 0, 0, 0, 244, 19,
		1, 0, 0, 0, 245, 246, 5, 57, 0, 0, 246, 247, 5, 42, 0, 0, 247, 248, 5,
		56, 0, 0, 248, 249, 7, 0, 0, 0, 249, 21, 1, 0, 0, 0, 250, 251, 5, 57, 0,
		0, 251, 252, 5, 43, 0, 0, 252, 253, 5, 56, 0, 0, 253, 254, 7, 0, 0, 0,
		254, 23, 1, 0, 0, 0, 255, 256, 5, 57, 0, 0, 256, 257, 5, 44, 0, 0, 257,
		258, 5, 56, 0, 0, 258, 259, 7, 0, 0, 0, 259, 25, 1, 0, 0, 0, 260, 261,
		5, 57, 0, 0, 261, 262, 5, 45, 0, 0, 262, 263, 5, 56, 0, 0, 263, 264, 7,
		0, 0, 0, 264, 27, 1, 0, 0, 0, 265, 266, 5, 57, 0, 0, 266, 267, 5, 46, 0,
		0, 267, 268, 5, 56, 0, 0, 268, 269, 7, 0, 0, 0, 269, 29, 1, 0, 0, 0, 270,
		271, 5, 57, 0, 0, 271, 272, 5, 47, 0, 0, 272, 273, 5, 56, 0, 0, 273, 274,
		5, 58, 0, 0, 274, 31, 1, 0, 0, 0, 275, 276, 5, 57, 0, 0, 276, 277, 5, 48,
		0, 0, 277, 278, 5, 56, 0, 0, 278, 279, 5, 1, 0, 0, 279, 33, 1, 0, 0, 0,
		280, 281, 5, 57, 0, 0, 281, 282, 5, 49, 0, 0, 282, 283, 5, 56, 0, 0, 283,
		284, 5, 58, 0, 0, 284, 35, 1, 0, 0, 0, 285, 286, 5, 57, 0, 0, 286, 287,
		5, 50, 0, 0, 287, 288, 5, 56, 0, 0, 288, 289, 7, 0, 0, 0, 289, 37, 1, 0,
		0, 0, 290, 291, 5, 57, 0, 0, 291, 292, 5, 51, 0, 0, 292, 293, 5, 56, 0,
		0, 293, 294, 7, 0, 0, 0, 294, 39, 1, 0, 0, 0, 295, 296, 5, 57, 0, 0, 296,
		297, 5, 52, 0, 0, 297, 298, 5, 56, 0, 0, 298, 299, 7, 0, 0, 0, 299, 41,
		1, 0, 0, 0, 300, 301, 5, 57, 0, 0, 301, 302, 5, 53, 0, 0, 302, 303, 5,
		56, 0, 0, 303, 304, 5, 58, 0, 0, 304, 43, 1, 0, 0, 0, 305, 306, 5, 57,
		0, 0, 306, 307, 5, 55, 0, 0, 307, 45, 1, 0, 0, 0, 308, 309, 5, 57, 0, 0,
		309, 310, 5, 54, 0, 0, 310, 47, 1, 0, 0, 0, 311, 313, 3, 84, 42, 0, 312,
		311, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314, 315,
		1, 0, 0, 0, 315, 49, 1, 0, 0, 0, 316, 318, 3, 88, 44, 0, 317, 316, 1, 0,
		0, 0, 318, 319, 1, 0, 0, 0, 319, 317, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0,
		320, 51, 1, 0, 0, 0, 321, 323, 3, 90, 45, 0, 322, 321, 1, 0, 0, 0, 323,
		324, 1, 0, 0, 0, 324, 322, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 53, 1,
		0, 0, 0, 326, 328, 3, 94, 47, 0, 327, 326, 1, 0, 0, 0, 328, 329, 1, 0,
		0, 0, 329, 327, 1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 55, 1, 0, 0, 0,
		331, 333, 3, 98, 49, 0, 332, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334,
		332, 1, 0, 0, 0, 334, 335, 1, 0, 0, 0, 335, 57, 1, 0, 0, 0, 336, 338, 3,
		96, 48, 0, 337, 336, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 337, 1, 0,
		0, 0, 339, 340, 1, 0, 0, 0, 340, 59, 1, 0, 0, 0, 341, 343, 3, 104, 52,
		0, 342, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 61, 1, 0, 0, 0, 346, 348, 3, 108, 54, 0, 347, 346,
		1, 0, 0, 0, 348, 349, 1, 0, 0, 0, 349, 347, 1, 0, 0, 0, 349, 350, 1, 0,
		0, 0, 350, 63, 1, 0, 0, 0, 351, 353, 3, 110, 55, 0, 352, 351, 1, 0, 0,
		0, 353, 354, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355,
		65, 1, 0, 0, 0, 356, 358, 3, 112, 56, 0, 357, 356, 1, 0, 0, 0, 358, 359,
		1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 67, 1, 0,
		0, 0, 361, 363, 3, 116, 58, 0, 362, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0,
		0, 364, 362, 1, 0, 0, 0, 364, 365, 1, 0, 0, 0, 365, 69, 1, 0, 0, 0, 366,
		368, 3, 118, 59, 0, 367, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 367,
		1, 0, 0, 0, 369, 370, 1, 0, 0, 0, 370, 71, 1, 0, 0, 0, 371, 373, 3, 120,
		60, 0, 372, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 372, 1, 0, 0, 0,
		374, 375, 1, 0, 0, 0, 375, 73, 1, 0, 0, 0, 376, 378, 3, 122, 61, 0, 377,
		376, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 377, 1, 0, 0, 0, 379, 380,
		1, 0, 0, 0, 380, 75, 1, 0, 0, 0, 381, 383, 3, 124, 62, 0, 382, 381, 1,
		0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 384, 385, 1, 0, 0,
		0, 385, 77, 1, 0, 0, 0, 386, 388, 3, 126, 63, 0, 387, 386, 1, 0, 0, 0,
		388, 389, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 389, 390, 1, 0, 0, 0, 390,
		79, 1, 0, 0, 0, 391, 393, 3, 128, 64, 0, 392, 391, 1, 0, 0, 0, 393, 394,
		1, 0, 0, 0, 394, 392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 81, 1, 0,
		0, 0, 396, 398, 3, 136, 68, 0, 397, 396, 1, 0, 0, 0, 398, 399, 1, 0, 0,
		0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 83, 1, 0, 0, 0, 401,
		406, 3, 4, 2, 0, 402, 406, 3, 6, 3, 0, 403, 406, 3, 8, 4, 0, 404, 406,
		3, 14, 7, 0, 405, 401, 1, 0, 0, 0, 405, 402, 1, 0, 0, 0, 405, 403, 1, 0,
		0, 0, 405, 404, 1, 0, 0, 0, 406, 85, 1, 0, 0, 0, 407, 408, 3, 14, 7, 0,
		408, 87, 1, 0, 0, 0, 409, 418, 3, 4, 2, 0, 410, 418, 3, 6, 3, 0, 411, 418,
		3, 8, 4, 0, 412, 418, 3, 14, 7, 0, 413, 418, 3, 10, 5, 0, 414, 418, 3,
		16, 8, 0, 415, 418, 3, 30, 15, 0, 416, 418, 3, 32, 16, 0, 417, 409, 1,
		0, 0, 0, 417, 410, 1, 0, 0, 0, 417, 411, 1, 0, 0, 0, 417, 412, 1, 0, 0,
		0, 417, 413, 1, 0, 0, 0, 417, 414, 1, 0, 0, 0, 417, 415, 1, 0, 0, 0, 417,
		416, 1, 0, 0, 0, 418, 89, 1, 0, 0, 0, 419, 422, 3, 14, 7, 0, 420, 422,
		3, 16, 8, 0, 421, 419, 1, 0, 0, 0, 421, 420, 1, 0, 0, 0, 422, 91, 1, 0,
		0, 0, 423, 424, 3, 12, 6, 0, 424, 93, 1, 0, 0, 0, 425, 429, 3, 12, 6, 0,
		426, 429, 3, 10, 5, 0, 427, 429, 3, 34, 17, 0, 428, 425, 1, 0, 0, 0, 428,
		426, 1, 0, 0, 0, 428, 427, 1, 0, 0, 0, 429, 95, 1, 0, 0, 0, 430, 431, 3,
		18, 9, 0, 431, 97, 1, 0, 0, 0, 432, 436, 3, 20, 10, 0, 433, 436, 3, 22,
		11, 0, 434, 436, 3, 12, 6, 0, 435, 432, 1, 0, 0, 0, 435, 433, 1, 0, 0,
		0, 435, 434, 1, 0, 0, 0, 436, 99, 1, 0, 0, 0, 437, 438, 3, 16, 8, 0, 438,
		101, 1, 0, 0, 0, 439, 440, 3, 16, 8, 0, 440, 103, 1, 0, 0, 0, 441, 445,
		3, 20, 10, 0, 442, 445, 3, 22, 11, 0, 443, 445, 3, 24, 12, 0, 444, 441,
		1, 0, 0, 0, 444, 442, 1, 0, 0, 0, 444, 443, 1, 0, 0, 0, 445, 105, 1, 0,
		0, 0, 446, 447, 3, 20, 10, 0, 447, 107, 1, 0, 0, 0, 448, 451, 3, 20, 10,
		0, 449, 451, 3, 24, 12, 0, 450, 448, 1, 0, 0, 0, 450, 449, 1, 0, 0, 0,
		451, 109, 1, 0, 0, 0, 452, 457, 3, 14, 7, 0, 453, 457, 3, 46, 23, 0, 454,
		457, 3, 4, 2, 0, 455, 457, 3, 26, 13, 0, 456, 452, 1, 0, 0, 0, 456, 453,
		1, 0, 0, 0, 456, 454, 1, 0, 0, 0, 456, 455, 1, 0, 0, 0, 457, 111, 1, 0,
		0, 0, 458, 461, 3, 44, 22, 0, 459, 461, 3, 14, 7, 0, 460, 458, 1, 0, 0,
		0, 460, 459, 1, 0, 0, 0, 461, 113, 1, 0, 0, 0, 462, 463, 3, 14, 7, 0, 463,
		115, 1, 0, 0, 0, 464, 467, 3, 14, 7, 0, 465, 467, 3, 36, 18, 0, 466, 464,
		1, 0, 0, 0, 466, 465, 1, 0, 0, 0, 467, 117, 1, 0, 0, 0, 468, 471, 3, 14,
		7, 0, 469, 471, 3, 16, 8, 0, 470, 468, 1, 0, 0, 0, 470, 469, 1, 0, 0, 0,
		471, 119, 1, 0, 0, 0, 472, 475, 3, 14, 7, 0, 473, 475, 3, 38, 19, 0, 474,
		472, 1, 0, 0, 0, 474, 473, 1, 0, 0, 0, 475, 121, 1, 0, 0, 0, 476, 479,
		3, 14, 7, 0, 477, 479, 3, 38, 19, 0, 478, 476, 1, 0, 0, 0, 478, 477, 1,
		0, 0, 0, 479, 123, 1, 0, 0, 0, 480, 483, 3, 14, 7, 0, 481, 483, 3, 16,
		8, 0, 482, 480, 1, 0, 0, 0, 482, 481, 1, 0, 0, 0, 483, 125, 1, 0, 0, 0,
		484, 488, 3, 14, 7, 0, 485, 488, 3, 40, 20, 0, 486, 488, 3, 46, 23, 0,
		487, 484, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 486, 1, 0, 0, 0, 488,
		127, 1, 0, 0, 0, 489, 493, 3, 14, 7, 0, 490, 493, 3, 42, 21, 0, 491, 493,
		3, 46, 23, 0, 492, 489, 1, 0, 0, 0, 492, 490, 1, 0, 0, 0, 492, 491, 1,
		0, 0, 0, 493, 129, 1, 0, 0, 0, 494, 495, 3, 12, 6, 0, 495, 131, 1, 0, 0,
		0, 496, 497, 3, 12, 6, 0, 497, 133, 1, 0, 0, 0, 498, 499, 3, 12, 6, 0,
		499, 135, 1, 0, 0, 0, 500, 505, 3, 16, 8, 0, 501, 505, 3, 14, 7, 0, 502,
		505, 3, 12, 6, 0, 503, 505, 3, 28, 14, 0, 504, 500, 1, 0, 0, 0, 504, 501,
		1, 0, 0, 0, 504, 502, 1, 0, 0, 0, 504, 503, 1, 0, 0, 0, 505, 137, 1, 0,
		0, 0, 37, 141, 202, 314, 319, 324, 329, 334, 339, 344, 349, 354, 359, 364,
		369, 374, 379, 384, 389, 394, 399, 405, 417, 421, 428, 435, 444, 450, 456,
		460, 466, 470, 474, 478, 482, 487, 492, 504,
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
	GoDiskGrammarUNMOUNT       = 10
	GoDiskGrammarMKFS          = 11
	GoDiskGrammarCAT           = 12
	GoDiskGrammarLOGIN         = 13
	GoDiskGrammarLOGOUT        = 14
	GoDiskGrammarMKGRP         = 15
	GoDiskGrammarRMGRP         = 16
	GoDiskGrammarMKUSR         = 17
	GoDiskGrammarRMUSR         = 18
	GoDiskGrammarCHGRP         = 19
	GoDiskGrammarMKFILE        = 20
	GoDiskGrammarMKDIR         = 21
	GoDiskGrammarREMOVE        = 22
	GoDiskGrammarEDIT          = 23
	GoDiskGrammarRENAME        = 24
	GoDiskGrammarCOPY          = 25
	GoDiskGrammarMOVE          = 26
	GoDiskGrammarFIND          = 27
	GoDiskGrammarCHOWN         = 28
	GoDiskGrammarCHMOD         = 29
	GoDiskGrammarRECOVERY      = 30
	GoDiskGrammarLOSS          = 31
	GoDiskGrammarJOURNALING    = 32
	GoDiskGrammarREP           = 33
	GoDiskGrammarID_TEXT       = 34
	GoDiskGrammarSIZE          = 35
	GoDiskGrammarFIT           = 36
	GoDiskGrammarUNIT          = 37
	GoDiskGrammarPATH          = 38
	GoDiskGrammarTYPE          = 39
	GoDiskGrammarNAME          = 40
	GoDiskGrammarFILE          = 41
	GoDiskGrammarUSER          = 42
	GoDiskGrammarPASS          = 43
	GoDiskGrammarGRP           = 44
	GoDiskGrammarCONT          = 45
	GoDiskGrammarPATH_FILE_LS  = 46
	GoDiskGrammarDELETE        = 47
	GoDiskGrammarADD           = 48
	GoDiskGrammarFS            = 49
	GoDiskGrammarCONTENIDO     = 50
	GoDiskGrammarDESTINO       = 51
	GoDiskGrammarUSUARIO       = 52
	GoDiskGrammarUGO           = 53
	GoDiskGrammarR             = 54
	GoDiskGrammarP             = 55
	GoDiskGrammarASSIGN        = 56
	GoDiskGrammarMINUS         = 57
	GoDiskGrammarID            = 58
	GoDiskGrammarUNQUOTED_TEXT = 59
	GoDiskGrammarLINE_COMMENT  = 60
	GoDiskGrammarBLOCK_COMMENT = 61
	GoDiskGrammarWS            = 62
)

// GoDiskGrammar rules.
const (
	GoDiskGrammarRULE_prog             = 0
	GoDiskGrammarRULE_stmt             = 1
	GoDiskGrammarRULE_size             = 2
	GoDiskGrammarRULE_fit              = 3
	GoDiskGrammarRULE_unit             = 4
	GoDiskGrammarRULE_type             = 5
	GoDiskGrammarRULE_id_text          = 6
	GoDiskGrammarRULE_path             = 7
	GoDiskGrammarRULE_name             = 8
	GoDiskGrammarRULE_filen            = 9
	GoDiskGrammarRULE_user             = 10
	GoDiskGrammarRULE_pass             = 11
	GoDiskGrammarRULE_grp              = 12
	GoDiskGrammarRULE_cont             = 13
	GoDiskGrammarRULE_path_file_ls     = 14
	GoDiskGrammarRULE_delete           = 15
	GoDiskGrammarRULE_add              = 16
	GoDiskGrammarRULE_fs               = 17
	GoDiskGrammarRULE_contenido        = 18
	GoDiskGrammarRULE_destino          = 19
	GoDiskGrammarRULE_usuario          = 20
	GoDiskGrammarRULE_ugo              = 21
	GoDiskGrammarRULE_p                = 22
	GoDiskGrammarRULE_r                = 23
	GoDiskGrammarRULE_mkdisk_params    = 24
	GoDiskGrammarRULE_fdisk_params     = 25
	GoDiskGrammarRULE_mount_params     = 26
	GoDiskGrammarRULE_mkfs_params      = 27
	GoDiskGrammarRULE_login_params     = 28
	GoDiskGrammarRULE_cat_params       = 29
	GoDiskGrammarRULE_mkusr_params     = 30
	GoDiskGrammarRULE_chgrp_params     = 31
	GoDiskGrammarRULE_mkfile_params    = 32
	GoDiskGrammarRULE_mkdir_params     = 33
	GoDiskGrammarRULE_edit_params      = 34
	GoDiskGrammarRULE_rename_params    = 35
	GoDiskGrammarRULE_copy_params      = 36
	GoDiskGrammarRULE_move_params      = 37
	GoDiskGrammarRULE_find_params      = 38
	GoDiskGrammarRULE_chown_params     = 39
	GoDiskGrammarRULE_chmod_params     = 40
	GoDiskGrammarRULE_rep_params       = 41
	GoDiskGrammarRULE_mkdisk_param     = 42
	GoDiskGrammarRULE_rmdisk_param     = 43
	GoDiskGrammarRULE_fdisk_param      = 44
	GoDiskGrammarRULE_mount_param      = 45
	GoDiskGrammarRULE_unmount_param    = 46
	GoDiskGrammarRULE_mkfs_param       = 47
	GoDiskGrammarRULE_cat_param        = 48
	GoDiskGrammarRULE_login_param      = 49
	GoDiskGrammarRULE_mkgrp_param      = 50
	GoDiskGrammarRULE_rmgrp_param      = 51
	GoDiskGrammarRULE_mkusr_param      = 52
	GoDiskGrammarRULE_rmusr_param      = 53
	GoDiskGrammarRULE_chgrp_param      = 54
	GoDiskGrammarRULE_mkfile_param     = 55
	GoDiskGrammarRULE_mkdir_param      = 56
	GoDiskGrammarRULE_remove_param     = 57
	GoDiskGrammarRULE_edit_param       = 58
	GoDiskGrammarRULE_rename_param     = 59
	GoDiskGrammarRULE_copy_param       = 60
	GoDiskGrammarRULE_move_param       = 61
	GoDiskGrammarRULE_find_param       = 62
	GoDiskGrammarRULE_chown_param      = 63
	GoDiskGrammarRULE_chmod_param      = 64
	GoDiskGrammarRULE_recovery_param   = 65
	GoDiskGrammarRULE_loss_param       = 66
	GoDiskGrammarRULE_journaling_param = 67
	GoDiskGrammarRULE_rep_param        = 68
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
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17179869152) != 0 {
		{
			p.SetState(138)
			p.Stmt()
		}

		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(144)
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

type LOSSContext struct {
	StmtContext
}

func NewLOSSContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LOSSContext {
	var p = new(LOSSContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *LOSSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LOSSContext) LOSS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarLOSS, 0)
}

func (s *LOSSContext) Loss_param() ILoss_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoss_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoss_paramContext)
}

func (s *LOSSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterLOSS(s)
	}
}

func (s *LOSSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitLOSS(s)
	}
}

func (s *LOSSContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitLOSS(s)

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

type CHMODContext struct {
	StmtContext
}

func NewCHMODContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CHMODContext {
	var p = new(CHMODContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *CHMODContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CHMODContext) CHMOD() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCHMOD, 0)
}

func (s *CHMODContext) Chmod_params() IChmod_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChmod_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChmod_paramsContext)
}

func (s *CHMODContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCHMOD(s)
	}
}

func (s *CHMODContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCHMOD(s)
	}
}

func (s *CHMODContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCHMOD(s)

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

type RECOVERYContext struct {
	StmtContext
}

func NewRECOVERYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RECOVERYContext {
	var p = new(RECOVERYContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *RECOVERYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RECOVERYContext) RECOVERY() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarRECOVERY, 0)
}

func (s *RECOVERYContext) Recovery_param() IRecovery_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRecovery_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRecovery_paramContext)
}

func (s *RECOVERYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRECOVERY(s)
	}
}

func (s *RECOVERYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRECOVERY(s)
	}
}

func (s *RECOVERYContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRECOVERY(s)

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

type REMOVEContext struct {
	StmtContext
}

func NewREMOVEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *REMOVEContext {
	var p = new(REMOVEContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *REMOVEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *REMOVEContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarREMOVE, 0)
}

func (s *REMOVEContext) Remove_param() IRemove_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRemove_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRemove_paramContext)
}

func (s *REMOVEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterREMOVE(s)
	}
}

func (s *REMOVEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitREMOVE(s)
	}
}

func (s *REMOVEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitREMOVE(s)

	default:
		return t.VisitChildren(s)
	}
}

type FINDContext struct {
	StmtContext
}

func NewFINDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FINDContext {
	var p = new(FINDContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *FINDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FINDContext) FIND() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarFIND, 0)
}

func (s *FINDContext) Find_params() IFind_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFind_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFind_paramsContext)
}

func (s *FINDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFIND(s)
	}
}

func (s *FINDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFIND(s)
	}
}

func (s *FINDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFIND(s)

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

type RENAMEContext struct {
	StmtContext
}

func NewRENAMEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RENAMEContext {
	var p = new(RENAMEContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *RENAMEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RENAMEContext) RENAME() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarRENAME, 0)
}

func (s *RENAMEContext) Rename_params() IRename_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRename_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRename_paramsContext)
}

func (s *RENAMEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRENAME(s)
	}
}

func (s *RENAMEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRENAME(s)
	}
}

func (s *RENAMEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRENAME(s)

	default:
		return t.VisitChildren(s)
	}
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

type JOURNALINGContext struct {
	StmtContext
}

func NewJOURNALINGContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JOURNALINGContext {
	var p = new(JOURNALINGContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *JOURNALINGContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JOURNALINGContext) JOURNALING() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarJOURNALING, 0)
}

func (s *JOURNALINGContext) Journaling_param() IJournaling_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJournaling_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJournaling_paramContext)
}

func (s *JOURNALINGContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterJOURNALING(s)
	}
}

func (s *JOURNALINGContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitJOURNALING(s)
	}
}

func (s *JOURNALINGContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitJOURNALING(s)

	default:
		return t.VisitChildren(s)
	}
}

type EDITContext struct {
	StmtContext
}

func NewEDITContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EDITContext {
	var p = new(EDITContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *EDITContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EDITContext) EDIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarEDIT, 0)
}

func (s *EDITContext) Edit_params() IEdit_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdit_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEdit_paramsContext)
}

func (s *EDITContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterEDIT(s)
	}
}

func (s *EDITContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitEDIT(s)
	}
}

func (s *EDITContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitEDIT(s)

	default:
		return t.VisitChildren(s)
	}
}

type COPYContext struct {
	StmtContext
}

func NewCOPYContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *COPYContext {
	var p = new(COPYContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *COPYContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *COPYContext) COPY() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCOPY, 0)
}

func (s *COPYContext) Copy_params() ICopy_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICopy_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICopy_paramsContext)
}

func (s *COPYContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCOPY(s)
	}
}

func (s *COPYContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCOPY(s)
	}
}

func (s *COPYContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCOPY(s)

	default:
		return t.VisitChildren(s)
	}
}

type UNMOUNTContext struct {
	StmtContext
}

func NewUNMOUNTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UNMOUNTContext {
	var p = new(UNMOUNTContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *UNMOUNTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UNMOUNTContext) UNMOUNT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNMOUNT, 0)
}

func (s *UNMOUNTContext) Unmount_param() IUnmount_paramContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnmount_paramContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnmount_paramContext)
}

func (s *UNMOUNTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterUNMOUNT(s)
	}
}

func (s *UNMOUNTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitUNMOUNT(s)
	}
}

func (s *UNMOUNTContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitUNMOUNT(s)

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

type CHOWNContext struct {
	StmtContext
}

func NewCHOWNContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CHOWNContext {
	var p = new(CHOWNContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *CHOWNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CHOWNContext) CHOWN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCHOWN, 0)
}

func (s *CHOWNContext) Chown_params() IChown_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChown_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IChown_paramsContext)
}

func (s *CHOWNContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCHOWN(s)
	}
}

func (s *CHOWNContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCHOWN(s)
	}
}

func (s *CHOWNContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCHOWN(s)

	default:
		return t.VisitChildren(s)
	}
}

type MOVEContext struct {
	StmtContext
}

func NewMOVEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MOVEContext {
	var p = new(MOVEContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *MOVEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MOVEContext) MOVE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMOVE, 0)
}

func (s *MOVEContext) Move_params() IMove_paramsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMove_paramsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMove_paramsContext)
}

func (s *MOVEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMOVE(s)
	}
}

func (s *MOVEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMOVE(s)
	}
}

func (s *MOVEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMOVE(s)

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
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoDiskGrammarMKDISK:
		localctx = NewMKDISKContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)
			p.Match(GoDiskGrammarMKDISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(147)
			p.Mkdisk_params()
		}

	case GoDiskGrammarRMDISK:
		localctx = NewRMDISKContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(GoDiskGrammarRMDISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(149)
			p.Rmdisk_param()
		}

	case GoDiskGrammarFDISK:
		localctx = NewFDISKContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.Match(GoDiskGrammarFDISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(151)
			p.Fdisk_params()
		}

	case GoDiskGrammarMOUNT:
		localctx = NewMOUNTContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(152)
			p.Match(GoDiskGrammarMOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(153)
			p.Mount_params()
		}

	case GoDiskGrammarMOUNTED:
		localctx = NewMOUNTEDContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(154)
			p.Match(GoDiskGrammarMOUNTED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoDiskGrammarUNMOUNT:
		localctx = NewUNMOUNTContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(155)
			p.Match(GoDiskGrammarUNMOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(156)
			p.Unmount_param()
		}

	case GoDiskGrammarMKFS:
		localctx = NewMKFSContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(157)
			p.Match(GoDiskGrammarMKFS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(158)
			p.Mkfs_params()
		}

	case GoDiskGrammarCAT:
		localctx = NewCATContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(159)
			p.Match(GoDiskGrammarCAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(160)
			p.Cat_params()
		}

	case GoDiskGrammarLOGIN:
		localctx = NewLOGINContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(161)
			p.Match(GoDiskGrammarLOGIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Login_params()
		}

	case GoDiskGrammarLOGOUT:
		localctx = NewLOGOUTContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(163)
			p.Match(GoDiskGrammarLOGOUT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoDiskGrammarMKGRP:
		localctx = NewMKGRPContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(164)
			p.Match(GoDiskGrammarMKGRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(165)
			p.Mkgrp_param()
		}

	case GoDiskGrammarRMGRP:
		localctx = NewRMGRPContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(166)
			p.Match(GoDiskGrammarRMGRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(167)
			p.Rmgrp_param()
		}

	case GoDiskGrammarMKUSR:
		localctx = NewMKUSRContext(p, localctx)
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(168)
			p.Match(GoDiskGrammarMKUSR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(169)
			p.Mkusr_params()
		}

	case GoDiskGrammarRMUSR:
		localctx = NewRMUSRContext(p, localctx)
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(170)
			p.Match(GoDiskGrammarRMUSR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Rmusr_param()
		}

	case GoDiskGrammarCHGRP:
		localctx = NewCHGRPContext(p, localctx)
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(172)
			p.Match(GoDiskGrammarCHGRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(173)
			p.Chgrp_params()
		}

	case GoDiskGrammarMKFILE:
		localctx = NewMKFILEContext(p, localctx)
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(174)
			p.Match(GoDiskGrammarMKFILE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.Mkfile_params()
		}

	case GoDiskGrammarMKDIR:
		localctx = NewMKDIRContext(p, localctx)
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(176)
			p.Match(GoDiskGrammarMKDIR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Mkdir_params()
		}

	case GoDiskGrammarREMOVE:
		localctx = NewREMOVEContext(p, localctx)
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(178)
			p.Match(GoDiskGrammarREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(179)
			p.Remove_param()
		}

	case GoDiskGrammarEDIT:
		localctx = NewEDITContext(p, localctx)
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(180)
			p.Match(GoDiskGrammarEDIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(181)
			p.Edit_params()
		}

	case GoDiskGrammarRENAME:
		localctx = NewRENAMEContext(p, localctx)
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(182)
			p.Match(GoDiskGrammarRENAME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(183)
			p.Rename_params()
		}

	case GoDiskGrammarCOPY:
		localctx = NewCOPYContext(p, localctx)
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(184)
			p.Match(GoDiskGrammarCOPY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(185)
			p.Copy_params()
		}

	case GoDiskGrammarMOVE:
		localctx = NewMOVEContext(p, localctx)
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(186)
			p.Match(GoDiskGrammarMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(187)
			p.Move_params()
		}

	case GoDiskGrammarFIND:
		localctx = NewFINDContext(p, localctx)
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(188)
			p.Match(GoDiskGrammarFIND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(189)
			p.Find_params()
		}

	case GoDiskGrammarCHOWN:
		localctx = NewCHOWNContext(p, localctx)
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(190)
			p.Match(GoDiskGrammarCHOWN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(191)
			p.Chown_params()
		}

	case GoDiskGrammarCHMOD:
		localctx = NewCHMODContext(p, localctx)
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(192)
			p.Match(GoDiskGrammarCHMOD)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(193)
			p.Chmod_params()
		}

	case GoDiskGrammarRECOVERY:
		localctx = NewRECOVERYContext(p, localctx)
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(194)
			p.Match(GoDiskGrammarRECOVERY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.Recovery_param()
		}

	case GoDiskGrammarLOSS:
		localctx = NewLOSSContext(p, localctx)
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(196)
			p.Match(GoDiskGrammarLOSS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Loss_param()
		}

	case GoDiskGrammarJOURNALING:
		localctx = NewJOURNALINGContext(p, localctx)
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(198)
			p.Match(GoDiskGrammarJOURNALING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(199)
			p.Journaling_param()
		}

	case GoDiskGrammarREP:
		localctx = NewREPContext(p, localctx)
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(200)
			p.Match(GoDiskGrammarREP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(201)
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
		p.SetState(204)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(GoDiskGrammarSIZE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
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
		p.SetState(209)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.Match(GoDiskGrammarFIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
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
		p.SetState(214)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(GoDiskGrammarUNIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(217)
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
		p.SetState(219)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(220)
		p.Match(GoDiskGrammarTYPE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(221)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(222)
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
		p.SetState(224)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.Match(GoDiskGrammarID_TEXT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(226)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(227)
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
		p.SetState(229)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.Match(GoDiskGrammarPATH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(231)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
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
		p.SetState(234)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.Match(GoDiskGrammarNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(237)
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
		p.SetState(239)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(240)
		p.Match(GoDiskGrammarFILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(241)
		p.Match(GoDiskGrammarINT_LIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(242)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
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
		p.SetState(245)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
		p.Match(GoDiskGrammarUSER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
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
		p.SetState(250)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(251)
		p.Match(GoDiskGrammarPASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
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
		p.SetState(255)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(256)
		p.Match(GoDiskGrammarGRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(257)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
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
		p.SetState(260)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(261)
		p.Match(GoDiskGrammarCONT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(262)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(263)
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
		p.SetState(265)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(266)
		p.Match(GoDiskGrammarPATH_FILE_LS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(268)
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

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	DELETE() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_delete
	return p
}

func InitEmptyDeleteContext(p *DeleteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_delete
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarDELETE, 0)
}

func (s *DeleteContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *DeleteContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (s *DeleteContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitDelete(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Delete_() (localctx IDeleteContext) {
	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoDiskGrammarRULE_delete)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(271)
		p.Match(GoDiskGrammarDELETE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(272)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
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

// IAddContext is an interface to support dynamic dispatch.
type IAddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	ADD() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	INT_LIT() antlr.TerminalNode

	// IsAddContext differentiates from other interfaces.
	IsAddContext()
}

type AddContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddContext() *AddContext {
	var p = new(AddContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_add
	return p
}

func InitEmptyAddContext(p *AddContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_add
}

func (*AddContext) IsAddContext() {}

func NewAddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddContext {
	var p = new(AddContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_add

	return p
}

func (s *AddContext) GetParser() antlr.Parser { return s.parser }

func (s *AddContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *AddContext) ADD() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarADD, 0)
}

func (s *AddContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *AddContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarINT_LIT, 0)
}

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitAdd(s)
	}
}

func (s *AddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitAdd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Add() (localctx IAddContext) {
	localctx = NewAddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoDiskGrammarRULE_add)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Match(GoDiskGrammarADD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(277)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(278)
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

// IFsContext is an interface to support dynamic dispatch.
type IFsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	FS() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsFsContext differentiates from other interfaces.
	IsFsContext()
}

type FsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFsContext() *FsContext {
	var p = new(FsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fs
	return p
}

func InitEmptyFsContext(p *FsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_fs
}

func (*FsContext) IsFsContext() {}

func NewFsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FsContext {
	var p = new(FsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_fs

	return p
}

func (s *FsContext) GetParser() antlr.Parser { return s.parser }

func (s *FsContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *FsContext) FS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarFS, 0)
}

func (s *FsContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *FsContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *FsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFs(s)
	}
}

func (s *FsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFs(s)
	}
}

func (s *FsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Fs() (localctx IFsContext) {
	localctx = NewFsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoDiskGrammarRULE_fs)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(280)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(281)
		p.Match(GoDiskGrammarFS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(283)
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

// IContenidoContext is an interface to support dynamic dispatch.
type IContenidoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	CONTENIDO() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsContenidoContext differentiates from other interfaces.
	IsContenidoContext()
}

type ContenidoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContenidoContext() *ContenidoContext {
	var p = new(ContenidoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_contenido
	return p
}

func InitEmptyContenidoContext(p *ContenidoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_contenido
}

func (*ContenidoContext) IsContenidoContext() {}

func NewContenidoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContenidoContext {
	var p = new(ContenidoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_contenido

	return p
}

func (s *ContenidoContext) GetParser() antlr.Parser { return s.parser }

func (s *ContenidoContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *ContenidoContext) CONTENIDO() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarCONTENIDO, 0)
}

func (s *ContenidoContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *ContenidoContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *ContenidoContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *ContenidoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContenidoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContenidoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterContenido(s)
	}
}

func (s *ContenidoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitContenido(s)
	}
}

func (s *ContenidoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitContenido(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Contenido() (localctx IContenidoContext) {
	localctx = NewContenidoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoDiskGrammarRULE_contenido)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.Match(GoDiskGrammarCONTENIDO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(287)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
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

// IDestinoContext is an interface to support dynamic dispatch.
type IDestinoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	DESTINO() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsDestinoContext differentiates from other interfaces.
	IsDestinoContext()
}

type DestinoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinoContext() *DestinoContext {
	var p = new(DestinoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_destino
	return p
}

func InitEmptyDestinoContext(p *DestinoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_destino
}

func (*DestinoContext) IsDestinoContext() {}

func NewDestinoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinoContext {
	var p = new(DestinoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_destino

	return p
}

func (s *DestinoContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinoContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *DestinoContext) DESTINO() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarDESTINO, 0)
}

func (s *DestinoContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *DestinoContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *DestinoContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *DestinoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterDestino(s)
	}
}

func (s *DestinoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitDestino(s)
	}
}

func (s *DestinoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitDestino(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Destino() (localctx IDestinoContext) {
	localctx = NewDestinoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoDiskGrammarRULE_destino)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(291)
		p.Match(GoDiskGrammarDESTINO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(292)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(293)
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

// IUsuarioContext is an interface to support dynamic dispatch.
type IUsuarioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	USUARIO() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	STRING_LIT() antlr.TerminalNode
	UNQUOTED_TEXT() antlr.TerminalNode

	// IsUsuarioContext differentiates from other interfaces.
	IsUsuarioContext()
}

type UsuarioContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUsuarioContext() *UsuarioContext {
	var p = new(UsuarioContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_usuario
	return p
}

func InitEmptyUsuarioContext(p *UsuarioContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_usuario
}

func (*UsuarioContext) IsUsuarioContext() {}

func NewUsuarioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UsuarioContext {
	var p = new(UsuarioContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_usuario

	return p
}

func (s *UsuarioContext) GetParser() antlr.Parser { return s.parser }

func (s *UsuarioContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *UsuarioContext) USUARIO() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUSUARIO, 0)
}

func (s *UsuarioContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *UsuarioContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarSTRING_LIT, 0)
}

func (s *UsuarioContext) UNQUOTED_TEXT() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUNQUOTED_TEXT, 0)
}

func (s *UsuarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UsuarioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UsuarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterUsuario(s)
	}
}

func (s *UsuarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitUsuario(s)
	}
}

func (s *UsuarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitUsuario(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Usuario() (localctx IUsuarioContext) {
	localctx = NewUsuarioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoDiskGrammarRULE_usuario)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(295)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(296)
		p.Match(GoDiskGrammarUSUARIO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(297)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(298)
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

// IUgoContext is an interface to support dynamic dispatch.
type IUgoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MINUS() antlr.TerminalNode
	UGO() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsUgoContext differentiates from other interfaces.
	IsUgoContext()
}

type UgoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUgoContext() *UgoContext {
	var p = new(UgoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_ugo
	return p
}

func InitEmptyUgoContext(p *UgoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_ugo
}

func (*UgoContext) IsUgoContext() {}

func NewUgoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UgoContext {
	var p = new(UgoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_ugo

	return p
}

func (s *UgoContext) GetParser() antlr.Parser { return s.parser }

func (s *UgoContext) MINUS() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarMINUS, 0)
}

func (s *UgoContext) UGO() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarUGO, 0)
}

func (s *UgoContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarASSIGN, 0)
}

func (s *UgoContext) ID() antlr.TerminalNode {
	return s.GetToken(GoDiskGrammarID, 0)
}

func (s *UgoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UgoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UgoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterUgo(s)
	}
}

func (s *UgoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitUgo(s)
	}
}

func (s *UgoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitUgo(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Ugo() (localctx IUgoContext) {
	localctx = NewUgoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoDiskGrammarRULE_ugo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(300)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(301)
		p.Match(GoDiskGrammarUGO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(302)
		p.Match(GoDiskGrammarASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(303)
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
	p.EnterRule(localctx, 44, GoDiskGrammarRULE_p)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(306)
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
	p.EnterRule(localctx, 46, GoDiskGrammarRULE_r)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(GoDiskGrammarMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
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
	p.EnterRule(localctx, 48, GoDiskGrammarRULE_mkdisk_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(311)
			p.Mkdisk_param()
		}

		p.SetState(314)
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
	p.EnterRule(localctx, 50, GoDiskGrammarRULE_fdisk_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(316)
			p.Fdisk_param()
		}

		p.SetState(319)
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
	p.EnterRule(localctx, 52, GoDiskGrammarRULE_mount_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(321)
			p.Mount_param()
		}

		p.SetState(324)
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
	p.EnterRule(localctx, 54, GoDiskGrammarRULE_mkfs_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(326)
			p.Mkfs_param()
		}

		p.SetState(329)
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
	p.EnterRule(localctx, 56, GoDiskGrammarRULE_login_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(331)
			p.Login_param()
		}

		p.SetState(334)
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
	p.EnterRule(localctx, 58, GoDiskGrammarRULE_cat_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(336)
			p.Cat_param()
		}

		p.SetState(339)
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
	p.EnterRule(localctx, 60, GoDiskGrammarRULE_mkusr_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(341)
			p.Mkusr_param()
		}

		p.SetState(344)
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
	p.EnterRule(localctx, 62, GoDiskGrammarRULE_chgrp_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(346)
			p.Chgrp_param()
		}

		p.SetState(349)
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
	p.EnterRule(localctx, 64, GoDiskGrammarRULE_mkfile_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(351)
			p.Mkfile_param()
		}

		p.SetState(354)
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
	p.EnterRule(localctx, 66, GoDiskGrammarRULE_mkdir_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(356)
			p.Mkdir_param()
		}

		p.SetState(359)
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

// IEdit_paramsContext is an interface to support dynamic dispatch.
type IEdit_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEdit_param() []IEdit_paramContext
	Edit_param(i int) IEdit_paramContext

	// IsEdit_paramsContext differentiates from other interfaces.
	IsEdit_paramsContext()
}

type Edit_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdit_paramsContext() *Edit_paramsContext {
	var p = new(Edit_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_edit_params
	return p
}

func InitEmptyEdit_paramsContext(p *Edit_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_edit_params
}

func (*Edit_paramsContext) IsEdit_paramsContext() {}

func NewEdit_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edit_paramsContext {
	var p = new(Edit_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_edit_params

	return p
}

func (s *Edit_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Edit_paramsContext) AllEdit_param() []IEdit_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEdit_paramContext); ok {
			len++
		}
	}

	tst := make([]IEdit_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEdit_paramContext); ok {
			tst[i] = t.(IEdit_paramContext)
			i++
		}
	}

	return tst
}

func (s *Edit_paramsContext) Edit_param(i int) IEdit_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEdit_paramContext); ok {
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

	return t.(IEdit_paramContext)
}

func (s *Edit_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edit_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edit_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterEdit_params(s)
	}
}

func (s *Edit_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitEdit_params(s)
	}
}

func (s *Edit_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitEdit_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Edit_params() (localctx IEdit_paramsContext) {
	localctx = NewEdit_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GoDiskGrammarRULE_edit_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(361)
			p.Edit_param()
		}

		p.SetState(364)
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

// IRename_paramsContext is an interface to support dynamic dispatch.
type IRename_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRename_param() []IRename_paramContext
	Rename_param(i int) IRename_paramContext

	// IsRename_paramsContext differentiates from other interfaces.
	IsRename_paramsContext()
}

type Rename_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRename_paramsContext() *Rename_paramsContext {
	var p = new(Rename_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rename_params
	return p
}

func InitEmptyRename_paramsContext(p *Rename_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rename_params
}

func (*Rename_paramsContext) IsRename_paramsContext() {}

func NewRename_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rename_paramsContext {
	var p = new(Rename_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rename_params

	return p
}

func (s *Rename_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Rename_paramsContext) AllRename_param() []IRename_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRename_paramContext); ok {
			len++
		}
	}

	tst := make([]IRename_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRename_paramContext); ok {
			tst[i] = t.(IRename_paramContext)
			i++
		}
	}

	return tst
}

func (s *Rename_paramsContext) Rename_param(i int) IRename_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRename_paramContext); ok {
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

	return t.(IRename_paramContext)
}

func (s *Rename_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rename_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rename_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRename_params(s)
	}
}

func (s *Rename_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRename_params(s)
	}
}

func (s *Rename_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRename_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rename_params() (localctx IRename_paramsContext) {
	localctx = NewRename_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GoDiskGrammarRULE_rename_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(366)
			p.Rename_param()
		}

		p.SetState(369)
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

// ICopy_paramsContext is an interface to support dynamic dispatch.
type ICopy_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllCopy_param() []ICopy_paramContext
	Copy_param(i int) ICopy_paramContext

	// IsCopy_paramsContext differentiates from other interfaces.
	IsCopy_paramsContext()
}

type Copy_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopy_paramsContext() *Copy_paramsContext {
	var p = new(Copy_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_copy_params
	return p
}

func InitEmptyCopy_paramsContext(p *Copy_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_copy_params
}

func (*Copy_paramsContext) IsCopy_paramsContext() {}

func NewCopy_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Copy_paramsContext {
	var p = new(Copy_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_copy_params

	return p
}

func (s *Copy_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Copy_paramsContext) AllCopy_param() []ICopy_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICopy_paramContext); ok {
			len++
		}
	}

	tst := make([]ICopy_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICopy_paramContext); ok {
			tst[i] = t.(ICopy_paramContext)
			i++
		}
	}

	return tst
}

func (s *Copy_paramsContext) Copy_param(i int) ICopy_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICopy_paramContext); ok {
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

	return t.(ICopy_paramContext)
}

func (s *Copy_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Copy_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Copy_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCopy_params(s)
	}
}

func (s *Copy_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCopy_params(s)
	}
}

func (s *Copy_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCopy_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Copy_params() (localctx ICopy_paramsContext) {
	localctx = NewCopy_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GoDiskGrammarRULE_copy_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(371)
			p.Copy_param()
		}

		p.SetState(374)
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

// IMove_paramsContext is an interface to support dynamic dispatch.
type IMove_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMove_param() []IMove_paramContext
	Move_param(i int) IMove_paramContext

	// IsMove_paramsContext differentiates from other interfaces.
	IsMove_paramsContext()
}

type Move_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMove_paramsContext() *Move_paramsContext {
	var p = new(Move_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_move_params
	return p
}

func InitEmptyMove_paramsContext(p *Move_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_move_params
}

func (*Move_paramsContext) IsMove_paramsContext() {}

func NewMove_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Move_paramsContext {
	var p = new(Move_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_move_params

	return p
}

func (s *Move_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Move_paramsContext) AllMove_param() []IMove_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMove_paramContext); ok {
			len++
		}
	}

	tst := make([]IMove_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMove_paramContext); ok {
			tst[i] = t.(IMove_paramContext)
			i++
		}
	}

	return tst
}

func (s *Move_paramsContext) Move_param(i int) IMove_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMove_paramContext); ok {
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

	return t.(IMove_paramContext)
}

func (s *Move_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Move_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Move_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMove_params(s)
	}
}

func (s *Move_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMove_params(s)
	}
}

func (s *Move_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMove_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Move_params() (localctx IMove_paramsContext) {
	localctx = NewMove_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GoDiskGrammarRULE_move_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(376)
			p.Move_param()
		}

		p.SetState(379)
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

// IFind_paramsContext is an interface to support dynamic dispatch.
type IFind_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFind_param() []IFind_paramContext
	Find_param(i int) IFind_paramContext

	// IsFind_paramsContext differentiates from other interfaces.
	IsFind_paramsContext()
}

type Find_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFind_paramsContext() *Find_paramsContext {
	var p = new(Find_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_find_params
	return p
}

func InitEmptyFind_paramsContext(p *Find_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_find_params
}

func (*Find_paramsContext) IsFind_paramsContext() {}

func NewFind_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Find_paramsContext {
	var p = new(Find_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_find_params

	return p
}

func (s *Find_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Find_paramsContext) AllFind_param() []IFind_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFind_paramContext); ok {
			len++
		}
	}

	tst := make([]IFind_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFind_paramContext); ok {
			tst[i] = t.(IFind_paramContext)
			i++
		}
	}

	return tst
}

func (s *Find_paramsContext) Find_param(i int) IFind_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFind_paramContext); ok {
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

	return t.(IFind_paramContext)
}

func (s *Find_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Find_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Find_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFind_params(s)
	}
}

func (s *Find_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFind_params(s)
	}
}

func (s *Find_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFind_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Find_params() (localctx IFind_paramsContext) {
	localctx = NewFind_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GoDiskGrammarRULE_find_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(381)
			p.Find_param()
		}

		p.SetState(384)
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

// IChown_paramsContext is an interface to support dynamic dispatch.
type IChown_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChown_param() []IChown_paramContext
	Chown_param(i int) IChown_paramContext

	// IsChown_paramsContext differentiates from other interfaces.
	IsChown_paramsContext()
}

type Chown_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChown_paramsContext() *Chown_paramsContext {
	var p = new(Chown_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chown_params
	return p
}

func InitEmptyChown_paramsContext(p *Chown_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chown_params
}

func (*Chown_paramsContext) IsChown_paramsContext() {}

func NewChown_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chown_paramsContext {
	var p = new(Chown_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_chown_params

	return p
}

func (s *Chown_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Chown_paramsContext) AllChown_param() []IChown_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChown_paramContext); ok {
			len++
		}
	}

	tst := make([]IChown_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChown_paramContext); ok {
			tst[i] = t.(IChown_paramContext)
			i++
		}
	}

	return tst
}

func (s *Chown_paramsContext) Chown_param(i int) IChown_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChown_paramContext); ok {
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

	return t.(IChown_paramContext)
}

func (s *Chown_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chown_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chown_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterChown_params(s)
	}
}

func (s *Chown_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitChown_params(s)
	}
}

func (s *Chown_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitChown_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Chown_params() (localctx IChown_paramsContext) {
	localctx = NewChown_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GoDiskGrammarRULE_chown_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(386)
			p.Chown_param()
		}

		p.SetState(389)
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

// IChmod_paramsContext is an interface to support dynamic dispatch.
type IChmod_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllChmod_param() []IChmod_paramContext
	Chmod_param(i int) IChmod_paramContext

	// IsChmod_paramsContext differentiates from other interfaces.
	IsChmod_paramsContext()
}

type Chmod_paramsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChmod_paramsContext() *Chmod_paramsContext {
	var p = new(Chmod_paramsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chmod_params
	return p
}

func InitEmptyChmod_paramsContext(p *Chmod_paramsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chmod_params
}

func (*Chmod_paramsContext) IsChmod_paramsContext() {}

func NewChmod_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chmod_paramsContext {
	var p = new(Chmod_paramsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_chmod_params

	return p
}

func (s *Chmod_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Chmod_paramsContext) AllChmod_param() []IChmod_paramContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IChmod_paramContext); ok {
			len++
		}
	}

	tst := make([]IChmod_paramContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IChmod_paramContext); ok {
			tst[i] = t.(IChmod_paramContext)
			i++
		}
	}

	return tst
}

func (s *Chmod_paramsContext) Chmod_param(i int) IChmod_paramContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IChmod_paramContext); ok {
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

	return t.(IChmod_paramContext)
}

func (s *Chmod_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chmod_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chmod_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterChmod_params(s)
	}
}

func (s *Chmod_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitChmod_params(s)
	}
}

func (s *Chmod_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitChmod_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Chmod_params() (localctx IChmod_paramsContext) {
	localctx = NewChmod_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GoDiskGrammarRULE_chmod_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(391)
			p.Chmod_param()
		}

		p.SetState(394)
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
	p.EnterRule(localctx, 82, GoDiskGrammarRULE_rep_params)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GoDiskGrammarMINUS {
		{
			p.SetState(396)
			p.Rep_param()
		}

		p.SetState(399)
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
	p.EnterRule(localctx, 84, GoDiskGrammarRULE_mkdisk_param)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.Size()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.Fit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)
			p.Unit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(404)
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
	p.EnterRule(localctx, 86, GoDiskGrammarRULE_rmdisk_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
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
	Delete_() IDeleteContext
	Add() IAddContext

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

func (s *Fdisk_paramContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *Fdisk_paramContext) Add() IAddContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAddContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAddContext)
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
	p.EnterRule(localctx, 88, GoDiskGrammarRULE_fdisk_param)
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Size()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Fit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(411)
			p.Unit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(412)
			p.Path()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(413)
			p.Type_()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(414)
			p.Name()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(415)
			p.Delete_()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(416)
			p.Add()
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
	p.EnterRule(localctx, 90, GoDiskGrammarRULE_mount_param)
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(419)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
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

// IUnmount_paramContext is an interface to support dynamic dispatch.
type IUnmount_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_text() IId_textContext

	// IsUnmount_paramContext differentiates from other interfaces.
	IsUnmount_paramContext()
}

type Unmount_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnmount_paramContext() *Unmount_paramContext {
	var p = new(Unmount_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_unmount_param
	return p
}

func InitEmptyUnmount_paramContext(p *Unmount_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_unmount_param
}

func (*Unmount_paramContext) IsUnmount_paramContext() {}

func NewUnmount_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unmount_paramContext {
	var p = new(Unmount_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_unmount_param

	return p
}

func (s *Unmount_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Unmount_paramContext) Id_text() IId_textContext {
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

func (s *Unmount_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unmount_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unmount_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterUnmount_param(s)
	}
}

func (s *Unmount_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitUnmount_param(s)
	}
}

func (s *Unmount_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitUnmount_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Unmount_param() (localctx IUnmount_paramContext) {
	localctx = NewUnmount_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GoDiskGrammarRULE_unmount_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(423)
		p.Id_text()
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
	Fs() IFsContext

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

func (s *Mkfs_paramContext) Fs() IFsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFsContext)
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
	p.EnterRule(localctx, 94, GoDiskGrammarRULE_mkfs_param)
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
			p.Id_text()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(426)
			p.Type_()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(427)
			p.Fs()
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
	p.EnterRule(localctx, 96, GoDiskGrammarRULE_cat_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
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
	p.EnterRule(localctx, 98, GoDiskGrammarRULE_login_param)
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			p.User()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			p.Pass()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(434)
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
	p.EnterRule(localctx, 100, GoDiskGrammarRULE_mkgrp_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
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
	p.EnterRule(localctx, 102, GoDiskGrammarRULE_rmgrp_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
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
	p.EnterRule(localctx, 104, GoDiskGrammarRULE_mkusr_param)
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)
			p.User()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Pass()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(443)
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
	p.EnterRule(localctx, 106, GoDiskGrammarRULE_rmusr_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(446)
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
	p.EnterRule(localctx, 108, GoDiskGrammarRULE_chgrp_param)
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			p.User()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
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
	p.EnterRule(localctx, 110, GoDiskGrammarRULE_mkfile_param)
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(452)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(453)
			p.R()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(454)
			p.Size()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(455)
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
	p.EnterRule(localctx, 112, GoDiskGrammarRULE_mkdir_param)
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(458)
			p.P()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(459)
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

// IRemove_paramContext is an interface to support dynamic dispatch.
type IRemove_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext

	// IsRemove_paramContext differentiates from other interfaces.
	IsRemove_paramContext()
}

type Remove_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemove_paramContext() *Remove_paramContext {
	var p = new(Remove_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_remove_param
	return p
}

func InitEmptyRemove_paramContext(p *Remove_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_remove_param
}

func (*Remove_paramContext) IsRemove_paramContext() {}

func NewRemove_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Remove_paramContext {
	var p = new(Remove_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_remove_param

	return p
}

func (s *Remove_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Remove_paramContext) Path() IPathContext {
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

func (s *Remove_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Remove_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Remove_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRemove_param(s)
	}
}

func (s *Remove_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRemove_param(s)
	}
}

func (s *Remove_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRemove_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Remove_param() (localctx IRemove_paramContext) {
	localctx = NewRemove_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, GoDiskGrammarRULE_remove_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(462)
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

// IEdit_paramContext is an interface to support dynamic dispatch.
type IEdit_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Contenido() IContenidoContext

	// IsEdit_paramContext differentiates from other interfaces.
	IsEdit_paramContext()
}

type Edit_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEdit_paramContext() *Edit_paramContext {
	var p = new(Edit_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_edit_param
	return p
}

func InitEmptyEdit_paramContext(p *Edit_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_edit_param
}

func (*Edit_paramContext) IsEdit_paramContext() {}

func NewEdit_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Edit_paramContext {
	var p = new(Edit_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_edit_param

	return p
}

func (s *Edit_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Edit_paramContext) Path() IPathContext {
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

func (s *Edit_paramContext) Contenido() IContenidoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContenidoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContenidoContext)
}

func (s *Edit_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Edit_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Edit_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterEdit_param(s)
	}
}

func (s *Edit_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitEdit_param(s)
	}
}

func (s *Edit_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitEdit_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Edit_param() (localctx IEdit_paramContext) {
	localctx = NewEdit_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, GoDiskGrammarRULE_edit_param)
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(464)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(465)
			p.Contenido()
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

// IRename_paramContext is an interface to support dynamic dispatch.
type IRename_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Name() INameContext

	// IsRename_paramContext differentiates from other interfaces.
	IsRename_paramContext()
}

type Rename_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRename_paramContext() *Rename_paramContext {
	var p = new(Rename_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rename_param
	return p
}

func InitEmptyRename_paramContext(p *Rename_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_rename_param
}

func (*Rename_paramContext) IsRename_paramContext() {}

func NewRename_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Rename_paramContext {
	var p = new(Rename_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_rename_param

	return p
}

func (s *Rename_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Rename_paramContext) Path() IPathContext {
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

func (s *Rename_paramContext) Name() INameContext {
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

func (s *Rename_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Rename_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Rename_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRename_param(s)
	}
}

func (s *Rename_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRename_param(s)
	}
}

func (s *Rename_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRename_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Rename_param() (localctx IRename_paramContext) {
	localctx = NewRename_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, GoDiskGrammarRULE_rename_param)
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(468)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(469)
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

// ICopy_paramContext is an interface to support dynamic dispatch.
type ICopy_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Destino() IDestinoContext

	// IsCopy_paramContext differentiates from other interfaces.
	IsCopy_paramContext()
}

type Copy_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopy_paramContext() *Copy_paramContext {
	var p = new(Copy_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_copy_param
	return p
}

func InitEmptyCopy_paramContext(p *Copy_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_copy_param
}

func (*Copy_paramContext) IsCopy_paramContext() {}

func NewCopy_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Copy_paramContext {
	var p = new(Copy_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_copy_param

	return p
}

func (s *Copy_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Copy_paramContext) Path() IPathContext {
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

func (s *Copy_paramContext) Destino() IDestinoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestinoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestinoContext)
}

func (s *Copy_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Copy_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Copy_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterCopy_param(s)
	}
}

func (s *Copy_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitCopy_param(s)
	}
}

func (s *Copy_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitCopy_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Copy_param() (localctx ICopy_paramContext) {
	localctx = NewCopy_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, GoDiskGrammarRULE_copy_param)
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(472)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(473)
			p.Destino()
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

// IMove_paramContext is an interface to support dynamic dispatch.
type IMove_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Destino() IDestinoContext

	// IsMove_paramContext differentiates from other interfaces.
	IsMove_paramContext()
}

type Move_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMove_paramContext() *Move_paramContext {
	var p = new(Move_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_move_param
	return p
}

func InitEmptyMove_paramContext(p *Move_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_move_param
}

func (*Move_paramContext) IsMove_paramContext() {}

func NewMove_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Move_paramContext {
	var p = new(Move_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_move_param

	return p
}

func (s *Move_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Move_paramContext) Path() IPathContext {
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

func (s *Move_paramContext) Destino() IDestinoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDestinoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDestinoContext)
}

func (s *Move_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Move_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Move_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterMove_param(s)
	}
}

func (s *Move_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitMove_param(s)
	}
}

func (s *Move_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitMove_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Move_param() (localctx IMove_paramContext) {
	localctx = NewMove_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, GoDiskGrammarRULE_move_param)
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.Destino()
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

// IFind_paramContext is an interface to support dynamic dispatch.
type IFind_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Name() INameContext

	// IsFind_paramContext differentiates from other interfaces.
	IsFind_paramContext()
}

type Find_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFind_paramContext() *Find_paramContext {
	var p = new(Find_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_find_param
	return p
}

func InitEmptyFind_paramContext(p *Find_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_find_param
}

func (*Find_paramContext) IsFind_paramContext() {}

func NewFind_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Find_paramContext {
	var p = new(Find_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_find_param

	return p
}

func (s *Find_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Find_paramContext) Path() IPathContext {
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

func (s *Find_paramContext) Name() INameContext {
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

func (s *Find_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Find_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Find_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterFind_param(s)
	}
}

func (s *Find_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitFind_param(s)
	}
}

func (s *Find_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitFind_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Find_param() (localctx IFind_paramContext) {
	localctx = NewFind_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, GoDiskGrammarRULE_find_param)
	p.SetState(482)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(480)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(481)
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

// IChown_paramContext is an interface to support dynamic dispatch.
type IChown_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Usuario() IUsuarioContext
	R() IRContext

	// IsChown_paramContext differentiates from other interfaces.
	IsChown_paramContext()
}

type Chown_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChown_paramContext() *Chown_paramContext {
	var p = new(Chown_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chown_param
	return p
}

func InitEmptyChown_paramContext(p *Chown_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chown_param
}

func (*Chown_paramContext) IsChown_paramContext() {}

func NewChown_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chown_paramContext {
	var p = new(Chown_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_chown_param

	return p
}

func (s *Chown_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Chown_paramContext) Path() IPathContext {
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

func (s *Chown_paramContext) Usuario() IUsuarioContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUsuarioContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUsuarioContext)
}

func (s *Chown_paramContext) R() IRContext {
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

func (s *Chown_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chown_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chown_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterChown_param(s)
	}
}

func (s *Chown_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitChown_param(s)
	}
}

func (s *Chown_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitChown_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Chown_param() (localctx IChown_paramContext) {
	localctx = NewChown_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, GoDiskGrammarRULE_chown_param)
	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(484)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)
			p.Usuario()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(486)
			p.R()
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

// IChmod_paramContext is an interface to support dynamic dispatch.
type IChmod_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Path() IPathContext
	Ugo() IUgoContext
	R() IRContext

	// IsChmod_paramContext differentiates from other interfaces.
	IsChmod_paramContext()
}

type Chmod_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChmod_paramContext() *Chmod_paramContext {
	var p = new(Chmod_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chmod_param
	return p
}

func InitEmptyChmod_paramContext(p *Chmod_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_chmod_param
}

func (*Chmod_paramContext) IsChmod_paramContext() {}

func NewChmod_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Chmod_paramContext {
	var p = new(Chmod_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_chmod_param

	return p
}

func (s *Chmod_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Chmod_paramContext) Path() IPathContext {
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

func (s *Chmod_paramContext) Ugo() IUgoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUgoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUgoContext)
}

func (s *Chmod_paramContext) R() IRContext {
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

func (s *Chmod_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Chmod_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Chmod_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterChmod_param(s)
	}
}

func (s *Chmod_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitChmod_param(s)
	}
}

func (s *Chmod_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitChmod_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Chmod_param() (localctx IChmod_paramContext) {
	localctx = NewChmod_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, GoDiskGrammarRULE_chmod_param)
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Path()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(490)
			p.Ugo()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(491)
			p.R()
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

// IRecovery_paramContext is an interface to support dynamic dispatch.
type IRecovery_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_text() IId_textContext

	// IsRecovery_paramContext differentiates from other interfaces.
	IsRecovery_paramContext()
}

type Recovery_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecovery_paramContext() *Recovery_paramContext {
	var p = new(Recovery_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_recovery_param
	return p
}

func InitEmptyRecovery_paramContext(p *Recovery_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_recovery_param
}

func (*Recovery_paramContext) IsRecovery_paramContext() {}

func NewRecovery_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Recovery_paramContext {
	var p = new(Recovery_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_recovery_param

	return p
}

func (s *Recovery_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Recovery_paramContext) Id_text() IId_textContext {
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

func (s *Recovery_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Recovery_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Recovery_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterRecovery_param(s)
	}
}

func (s *Recovery_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitRecovery_param(s)
	}
}

func (s *Recovery_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitRecovery_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Recovery_param() (localctx IRecovery_paramContext) {
	localctx = NewRecovery_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, GoDiskGrammarRULE_recovery_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(494)
		p.Id_text()
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

// ILoss_paramContext is an interface to support dynamic dispatch.
type ILoss_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_text() IId_textContext

	// IsLoss_paramContext differentiates from other interfaces.
	IsLoss_paramContext()
}

type Loss_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoss_paramContext() *Loss_paramContext {
	var p = new(Loss_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_loss_param
	return p
}

func InitEmptyLoss_paramContext(p *Loss_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_loss_param
}

func (*Loss_paramContext) IsLoss_paramContext() {}

func NewLoss_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Loss_paramContext {
	var p = new(Loss_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_loss_param

	return p
}

func (s *Loss_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Loss_paramContext) Id_text() IId_textContext {
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

func (s *Loss_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Loss_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Loss_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterLoss_param(s)
	}
}

func (s *Loss_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitLoss_param(s)
	}
}

func (s *Loss_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitLoss_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Loss_param() (localctx ILoss_paramContext) {
	localctx = NewLoss_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, GoDiskGrammarRULE_loss_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Id_text()
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

// IJournaling_paramContext is an interface to support dynamic dispatch.
type IJournaling_paramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Id_text() IId_textContext

	// IsJournaling_paramContext differentiates from other interfaces.
	IsJournaling_paramContext()
}

type Journaling_paramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJournaling_paramContext() *Journaling_paramContext {
	var p = new(Journaling_paramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_journaling_param
	return p
}

func InitEmptyJournaling_paramContext(p *Journaling_paramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoDiskGrammarRULE_journaling_param
}

func (*Journaling_paramContext) IsJournaling_paramContext() {}

func NewJournaling_paramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Journaling_paramContext {
	var p = new(Journaling_paramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoDiskGrammarRULE_journaling_param

	return p
}

func (s *Journaling_paramContext) GetParser() antlr.Parser { return s.parser }

func (s *Journaling_paramContext) Id_text() IId_textContext {
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

func (s *Journaling_paramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Journaling_paramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Journaling_paramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.EnterJournaling_param(s)
	}
}

func (s *Journaling_paramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoDiskGrammarListener); ok {
		listenerT.ExitJournaling_param(s)
	}
}

func (s *Journaling_paramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoDiskGrammarVisitor:
		return t.VisitJournaling_param(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoDiskGrammar) Journaling_param() (localctx IJournaling_paramContext) {
	localctx = NewJournaling_paramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, GoDiskGrammarRULE_journaling_param)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.Id_text()
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
	p.EnterRule(localctx, 136, GoDiskGrammarRULE_rep_param)
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(500)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)
			p.Path()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(502)
			p.Id_text()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(503)
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
