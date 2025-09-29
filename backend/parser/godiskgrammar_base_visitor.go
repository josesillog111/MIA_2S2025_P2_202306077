// Code generated from parser/GoDiskGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package godisk // GoDiskGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseGoDiskGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoDiskGrammarVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMKDISK(ctx *MKDISKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRMDISK(ctx *RMDISKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFDISK(ctx *FDISKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMOUNT(ctx *MOUNTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMOUNTED(ctx *MOUNTEDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitUNMOUNT(ctx *UNMOUNTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMKFS(ctx *MKFSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCAT(ctx *CATContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitLOGIN(ctx *LOGINContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitLOGOUT(ctx *LOGOUTContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMKGRP(ctx *MKGRPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRMGRP(ctx *RMGRPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMKUSR(ctx *MKUSRContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRMUSR(ctx *RMUSRContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCHGRP(ctx *CHGRPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMKFILE(ctx *MKFILEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMKDIR(ctx *MKDIRContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitREMOVE(ctx *REMOVEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitEDIT(ctx *EDITContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRENAME(ctx *RENAMEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCOPY(ctx *COPYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMOVE(ctx *MOVEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFIND(ctx *FINDContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCHOWN(ctx *CHOWNContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCHMOD(ctx *CHMODContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRECOVERY(ctx *RECOVERYContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitLOSS(ctx *LOSSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitJOURNALING(ctx *JOURNALINGContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitREP(ctx *REPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitSize(ctx *SizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFit(ctx *FitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitUnit(ctx *UnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitId_text(ctx *Id_textContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFilen(ctx *FilenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitUser(ctx *UserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitPass(ctx *PassContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitGrp(ctx *GrpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCont(ctx *ContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitPath_file_ls(ctx *Path_file_lsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitDelete(ctx *DeleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFs(ctx *FsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitContenido(ctx *ContenidoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitDestino(ctx *DestinoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitUsuario(ctx *UsuarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitUgo(ctx *UgoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitP(ctx *PContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitR(ctx *RContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkdisk_params(ctx *Mkdisk_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFdisk_params(ctx *Fdisk_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMount_params(ctx *Mount_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkfs_params(ctx *Mkfs_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitLogin_params(ctx *Login_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCat_params(ctx *Cat_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkusr_params(ctx *Mkusr_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitChgrp_params(ctx *Chgrp_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkfile_params(ctx *Mkfile_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkdir_params(ctx *Mkdir_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitEdit_params(ctx *Edit_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRename_params(ctx *Rename_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCopy_params(ctx *Copy_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMove_params(ctx *Move_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFind_params(ctx *Find_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitChown_params(ctx *Chown_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitChmod_params(ctx *Chmod_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRep_params(ctx *Rep_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkdisk_param(ctx *Mkdisk_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRmdisk_param(ctx *Rmdisk_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFdisk_param(ctx *Fdisk_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMount_param(ctx *Mount_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitUnmount_param(ctx *Unmount_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkfs_param(ctx *Mkfs_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCat_param(ctx *Cat_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitLogin_param(ctx *Login_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkgrp_param(ctx *Mkgrp_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRmgrp_param(ctx *Rmgrp_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkusr_param(ctx *Mkusr_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRmusr_param(ctx *Rmusr_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitChgrp_param(ctx *Chgrp_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkfile_param(ctx *Mkfile_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMkdir_param(ctx *Mkdir_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRemove_param(ctx *Remove_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitEdit_param(ctx *Edit_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRename_param(ctx *Rename_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitCopy_param(ctx *Copy_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitMove_param(ctx *Move_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitFind_param(ctx *Find_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitChown_param(ctx *Chown_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitChmod_param(ctx *Chmod_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRecovery_param(ctx *Recovery_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitLoss_param(ctx *Loss_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitJournaling_param(ctx *Journaling_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoDiskGrammarVisitor) VisitRep_param(ctx *Rep_paramContext) interface{} {
	return v.VisitChildren(ctx)
}
