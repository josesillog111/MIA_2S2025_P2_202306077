// Code generated from parser/GoDiskGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package godisk // GoDiskGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoDiskGrammar.
type GoDiskGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoDiskGrammar#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MKDISK.
	VisitMKDISK(ctx *MKDISKContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#RMDISK.
	VisitRMDISK(ctx *RMDISKContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#FDISK.
	VisitFDISK(ctx *FDISKContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MOUNT.
	VisitMOUNT(ctx *MOUNTContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MOUNTED.
	VisitMOUNTED(ctx *MOUNTEDContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MKFS.
	VisitMKFS(ctx *MKFSContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#CAT.
	VisitCAT(ctx *CATContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#LOGIN.
	VisitLOGIN(ctx *LOGINContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#LOGOUT.
	VisitLOGOUT(ctx *LOGOUTContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MKGRP.
	VisitMKGRP(ctx *MKGRPContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#RMGRP.
	VisitRMGRP(ctx *RMGRPContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MKUSR.
	VisitMKUSR(ctx *MKUSRContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#RMUSR.
	VisitRMUSR(ctx *RMUSRContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#CHGRP.
	VisitCHGRP(ctx *CHGRPContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MKFILE.
	VisitMKFILE(ctx *MKFILEContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#MKDIR.
	VisitMKDIR(ctx *MKDIRContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#REP.
	VisitREP(ctx *REPContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#size.
	VisitSize(ctx *SizeContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#fit.
	VisitFit(ctx *FitContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#unit.
	VisitUnit(ctx *UnitContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#id_text.
	VisitId_text(ctx *Id_textContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#filen.
	VisitFilen(ctx *FilenContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#user.
	VisitUser(ctx *UserContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#pass.
	VisitPass(ctx *PassContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#grp.
	VisitGrp(ctx *GrpContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#cont.
	VisitCont(ctx *ContContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#path_file_ls.
	VisitPath_file_ls(ctx *Path_file_lsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#p.
	VisitP(ctx *PContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#r.
	VisitR(ctx *RContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkdisk_params.
	VisitMkdisk_params(ctx *Mkdisk_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#fdisk_params.
	VisitFdisk_params(ctx *Fdisk_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mount_params.
	VisitMount_params(ctx *Mount_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkfs_params.
	VisitMkfs_params(ctx *Mkfs_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#login_params.
	VisitLogin_params(ctx *Login_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#cat_params.
	VisitCat_params(ctx *Cat_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkusr_params.
	VisitMkusr_params(ctx *Mkusr_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#chgrp_params.
	VisitChgrp_params(ctx *Chgrp_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkfile_params.
	VisitMkfile_params(ctx *Mkfile_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkdir_params.
	VisitMkdir_params(ctx *Mkdir_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#rep_params.
	VisitRep_params(ctx *Rep_paramsContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkdisk_param.
	VisitMkdisk_param(ctx *Mkdisk_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#rmdisk_param.
	VisitRmdisk_param(ctx *Rmdisk_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#fdisk_param.
	VisitFdisk_param(ctx *Fdisk_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mount_param.
	VisitMount_param(ctx *Mount_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkfs_param.
	VisitMkfs_param(ctx *Mkfs_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#cat_param.
	VisitCat_param(ctx *Cat_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#login_param.
	VisitLogin_param(ctx *Login_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkgrp_param.
	VisitMkgrp_param(ctx *Mkgrp_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#rmgrp_param.
	VisitRmgrp_param(ctx *Rmgrp_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkusr_param.
	VisitMkusr_param(ctx *Mkusr_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#rmusr_param.
	VisitRmusr_param(ctx *Rmusr_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#chgrp_param.
	VisitChgrp_param(ctx *Chgrp_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkfile_param.
	VisitMkfile_param(ctx *Mkfile_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#mkdir_param.
	VisitMkdir_param(ctx *Mkdir_paramContext) interface{}

	// Visit a parse tree produced by GoDiskGrammar#rep_param.
	VisitRep_param(ctx *Rep_paramContext) interface{}
}
