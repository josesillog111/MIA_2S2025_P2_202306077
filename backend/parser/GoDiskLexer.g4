lexer grammar GoDiskLexer;

INT_LIT     : [0-9]+ ;
FLOAT_LIT   : [0-9]+ '.' [0-9]+ ;
STRING_LIT  : '"' (~["\\] | ESCAPE_SEQ)* '"' ;
CHAR_LIT    : '\'' . '\'' ;

// funciones
MKDISK       : ('Mkdisk' | 'mkdisk' | 'MKDISK') ;
RMDISK       : ('Rmdisk' | 'rmdisk' | 'RMDISK') ;
FDISK        : ('Fdisk' | 'fdisk' | 'FDISK') ;
MOUNT       : ('Mount' | 'mount' | 'MOUNT') ;
MOUNTED     : ('Mounted' | 'mounted' | 'MOUNTED') ;
MKFS        : ('Mkfs' | 'mkfs' | 'MKFS') ;
CAT          : ('Cat' | 'cat' | 'CAT') ;
LOGIN        : ('Login' | 'login' | 'LOGIN') ;
LOGOUT       : ('Logout' | 'logout' | 'LOGOUT') ;
MKGRP       : ('Mkgrp' | 'mkgrp' | 'MKGRP') ;
RMGRP       : ('Rmgrp' | 'rmgrp' | 'RMGRP') ;
MKUSR        : ('Mkusr' | 'mkusr' | 'MKUSR') ;
RMUSR        : ('Rmusr' | 'rmusr' | 'RMUSR') ;
CHGRP        : ('Chgrp' | 'chgrp' | 'CHGRP') ;
MKFILE       : ('Mkfile' | 'mkfile' | 'MKFILE') ;
MKDIR        : ('Mkdir' | 'mkdir' | 'MKDIR') ;
REP          : ('Rep' | 'rep' | 'REP') ;

// parametros
ID_TEXT      : ('id' | 'ID' | 'Id') ;
SIZE         : ('size' | 'SIZE' | 'Size');
FIT          : ('fit' | 'FIT' | 'Fit');
UNIT         : ('unit' | 'UNIT' | 'Unit');
PATH         : ('path' | 'PATH' | 'Path');
TYPE         : ('type' | 'TYPE' | 'Type');
NAME         : ('name' | 'NAME' | 'Name');
FILE         : ('file' | 'FILE' | 'File');
USER         : ('user' | 'USER' | 'User');
PASS         : ('pass' | 'PASS' | 'Pass');
GRP          : ('grp' | 'GRP' | 'Grp');
CONT         : ('cont' | 'CONT' | 'Cont');
PATH_FILE_LS : ('path_file_ls' | 'PATH_FILE_LS' | 'Path_File_Ls');
R            : ('r' | 'R');
P            : ('p' | 'P');

ASSIGN              : '=' ;
MINUS               : '-' ;
ID                  : [_a-zA-Z0-9][_a-zA-Z0-9]* ;
UNQUOTED_TEXT       : [a-zA-Z0-9_./\\:]+ ;
LINE_COMMENT        : '#' ~[\r\n]* -> skip ;
BLOCK_COMMENT       : '/*' .*? '*/' -> skip ;

fragment ESCAPE_SEQ : '\\' [btnrf0'"\\] ;
WS : [ \t\r\n]+ -> skip ;