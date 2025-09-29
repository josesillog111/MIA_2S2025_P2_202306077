parser grammar GoDiskGrammar;

options { tokenVocab=GoDiskLexer; }

prog: stmt* EOF
    ;

stmt: MKDISK mkdisk_params          # MKDISK
    | RMDISK rmdisk_param           # RMDISK   
    | FDISK fdisk_params            # FDISK
    | MOUNT mount_params            # MOUNT   
    | MOUNTED                       # MOUNTED
    | UNMOUNT unmount_param         # UNMOUNT
    | MKFS mkfs_params              # MKFS
    | CAT cat_params                # CAT   
    | LOGIN login_params            # LOGIN
    | LOGOUT                        # LOGOUT
    | MKGRP mkgrp_param             # MKGRP 
    | RMGRP rmgrp_param             # RMGRP
    | MKUSR mkusr_params            # MKUSR
    | RMUSR rmusr_param             # RMUSR
    | CHGRP chgrp_params            # CHGRP
    | MKFILE mkfile_params          # MKFILE
    | MKDIR mkdir_params            # MKDIR
    | REMOVE remove_param           # REMOVE
    | EDIT edit_params              # EDIT
    | RENAME rename_params          # RENAME
    | COPY copy_params              # COPY
    | MOVE move_params              # MOVE
    | FIND find_params              # FIND
    | CHOWN chown_params            # CHOWN
    | CHMOD chmod_params            # CHMOD
    | RECOVERY recovery_param       # RECOVERY
    | LOSS loss_param               # LOSS
    | JOURNALING journaling_param   # JOURNALING
    | REP rep_params                # REP
    ;

/*
   
PARÁMETROS

*/

size            : MINUS SIZE ASSIGN INT_LIT ;
fit             : MINUS FIT ASSIGN ID ;
unit            : MINUS UNIT ASSIGN ID ;
type            : MINUS TYPE ASSIGN ID;
id_text         : MINUS ID_TEXT ASSIGN ID ;
path            : MINUS PATH ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
name            : MINUS NAME ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
filen           : MINUS FILE INT_LIT ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
user            : MINUS USER ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
pass            : MINUS PASS ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
grp             : MINUS GRP ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
cont            : MINUS CONT ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
path_file_ls    : MINUS PATH_FILE_LS ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
delete          : MINUS DELETE ASSIGN ID ;
add             : MINUS ADD ASSIGN INT_LIT ;
fs              : MINUS FS ASSIGN ID ;
contenido       : MINUS CONTENIDO ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
destino         : MINUS DESTINO ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
usuario         : MINUS USUARIO ASSIGN (STRING_LIT | UNQUOTED_TEXT) ;
ugo             : MINUS UGO ASSIGN ID ;
p               : MINUS P ;
r               : MINUS R ;


/* 

CONJUNTO DE PARÁMETROS

 */
mkdisk_params
  : mkdisk_param+  
  ;

fdisk_params
    : fdisk_param+  
    ;

mount_params
    : mount_param+   
    ;

mkfs_params
    : mkfs_param+  
    ;

login_params
    : login_param+  
    ;

cat_params
    : cat_param+  
    ;

mkusr_params
    : mkusr_param+  
    ;

chgrp_params
    : chgrp_param+  
    ;

mkfile_params
    : mkfile_param+  
    ;

mkdir_params
    : mkdir_param+  
    ;

edit_params
    : edit_param+  
    ;

rename_params
    : rename_param+
    ;

copy_params
    : copy_param+  
    ;

move_params
    : move_param+  
    ;

find_params
    : find_param+  
    ;

chown_params
    : chown_param+  
    ;

chmod_params
    : chmod_param+  
    ;

rep_params
    : rep_param+  
    ;


/*

PARÁMETROS INDIVIDUALES

 */
mkdisk_param
  : size
  | fit
  | unit
  | path
  ;

rmdisk_param: path ;

fdisk_param
    : size
    | fit
    | unit
    | path
    | type
    | name
    | delete
    | add
    ;

mount_param
    : path 
    | name 
    ;

unmount_param
    : id_text 
    ;

mkfs_param
    : id_text 
    | type 
    | fs
    ;

cat_param
    : filen ;

login_param
    : user 
    | pass 
    | id_text 
    ;

mkgrp_param
    : name 
    ;

rmgrp_param
    : name 
    ;

mkusr_param
    : user 
    | pass 
    | grp 
    ;


rmusr_param
    : user 
    ;

chgrp_param
    : user 
    | grp ;

mkfile_param
    : path 
    | r 
    | size 
    | cont ;

mkdir_param
    : p 
    | path ;


remove_param
    : path 
    ;

edit_param
    : path 
    | contenido 
    ;

rename_param
    : path 
    | name 
    ;

copy_param
    : path 
    | destino 
    ;

move_param
    : path 
    | destino 
    ;

find_param
    : path 
    | name 
    ;

chown_param
    : path 
    | usuario 
    | r 
    ;

chmod_param
    : path 
    | ugo
    | r 
    ;

recovery_param
    : id_text 
    ;

loss_param
    : id_text 
    ;

journaling_param
    : id_text 
    ;

rep_param
    : name 
    | path 
    | id_text 
    | path_file_ls 
    ;
