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
UNMOUNT     : ('Unmount' | 'unmount' | 'UNMOUNT') ;
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
REMOVE       : ('Remove' | 'remove' | 'REMOVE') ;
EDIT         : ('Edit' | 'edit' | 'EDIT') ;
RENAME       : ('Rename' | 'rename' | 'RENAME') ;
COPY         : ('Copy' | 'copy' | 'COPY') ;
MOVE         : ('Move' | 'move' | 'MOVE') ;
FIND         : ('Find' | 'find' | 'FIND') ;
CHOWN        : ('Chown' | 'chown' | 'CHOWN') ;
CHMOD        : ('Chmod' | 'chmod' | 'CHMOD') ;
RECOVERY     : ('Recovery' | 'recovery' | 'RECOVERY') ;
LOSS         : ('Loss' | 'loss' | 'LOSS') ;
JOURNALING   : ('Journaling' | 'journaling' | 'JOURNALING') ;
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
DELETE       : ('delete' | 'DELETE' | 'Delete');
ADD          : ('add' | 'ADD' | 'Add');
FS           : ('fs' | 'FS' | 'Fs');
CONTENIDO    : ('contenido' | 'CONTENIDO' | 'Contenido') ;
DESTINO      : ('destino' | 'DESTINO' | 'Destino') ;
USUARIO      : ('usuario' | 'USUARIO' | 'Usuario') ;
UGO          : ('ugo' | 'UGO' | 'Ugo') ;
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