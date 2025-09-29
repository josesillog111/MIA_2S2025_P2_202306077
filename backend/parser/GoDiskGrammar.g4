parser grammar GoDiskGrammar;

options { tokenVocab=GoDiskLexer; }

prog: stmt* EOF
    ;

stmt: MKDISK mkdisk_params   # MKDISK
    | RMDISK rmdisk_param    # RMDISK   
    | FDISK fdisk_params     # FDISK
    | MOUNT mount_params     # MOUNT   
    | MOUNTED                # MOUNTED
    | MKFS mkfs_params       # MKFS
    | CAT cat_params         # CAT   
    | LOGIN login_params     # LOGIN
    | LOGOUT                 # LOGOUT
    | MKGRP mkgrp_param      # MKGRP 
    | RMGRP rmgrp_param      # RMGRP
    | MKUSR mkusr_params     # MKUSR
    | RMUSR rmusr_param      # RMUSR
    | CHGRP chgrp_params     # CHGRP
    | MKFILE mkfile_params   # MKFILE
    | MKDIR mkdir_params     # MKDIR
    | REP rep_params         # REP
    ;

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
    ;

mount_param
    : path 
    | name ;

mkfs_param
    : id_text 
    | type ;

cat_param
    : filen ;

login_param
    : user 
    | pass 
    | id_text ;

mkgrp_param
    : name ;

rmgrp_param
    : name ;

mkusr_param
    : user 
    | pass 
    | grp ;


rmusr_param
    : user ;

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

rep_param
    : name 
    | path 
    | id_text 
    | path_file_ls ;
