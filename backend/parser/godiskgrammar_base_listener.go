// Code generated from parser/GoDiskGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package godisk // GoDiskGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseGoDiskGrammarListener is a complete listener for a parse tree produced by GoDiskGrammar.
type BaseGoDiskGrammarListener struct{}

var _ GoDiskGrammarListener = &BaseGoDiskGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoDiskGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoDiskGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoDiskGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoDiskGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseGoDiskGrammarListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseGoDiskGrammarListener) ExitProg(ctx *ProgContext) {}

// EnterMKDISK is called when production MKDISK is entered.
func (s *BaseGoDiskGrammarListener) EnterMKDISK(ctx *MKDISKContext) {}

// ExitMKDISK is called when production MKDISK is exited.
func (s *BaseGoDiskGrammarListener) ExitMKDISK(ctx *MKDISKContext) {}

// EnterRMDISK is called when production RMDISK is entered.
func (s *BaseGoDiskGrammarListener) EnterRMDISK(ctx *RMDISKContext) {}

// ExitRMDISK is called when production RMDISK is exited.
func (s *BaseGoDiskGrammarListener) ExitRMDISK(ctx *RMDISKContext) {}

// EnterFDISK is called when production FDISK is entered.
func (s *BaseGoDiskGrammarListener) EnterFDISK(ctx *FDISKContext) {}

// ExitFDISK is called when production FDISK is exited.
func (s *BaseGoDiskGrammarListener) ExitFDISK(ctx *FDISKContext) {}

// EnterMOUNT is called when production MOUNT is entered.
func (s *BaseGoDiskGrammarListener) EnterMOUNT(ctx *MOUNTContext) {}

// ExitMOUNT is called when production MOUNT is exited.
func (s *BaseGoDiskGrammarListener) ExitMOUNT(ctx *MOUNTContext) {}

// EnterMOUNTED is called when production MOUNTED is entered.
func (s *BaseGoDiskGrammarListener) EnterMOUNTED(ctx *MOUNTEDContext) {}

// ExitMOUNTED is called when production MOUNTED is exited.
func (s *BaseGoDiskGrammarListener) ExitMOUNTED(ctx *MOUNTEDContext) {}

// EnterMKFS is called when production MKFS is entered.
func (s *BaseGoDiskGrammarListener) EnterMKFS(ctx *MKFSContext) {}

// ExitMKFS is called when production MKFS is exited.
func (s *BaseGoDiskGrammarListener) ExitMKFS(ctx *MKFSContext) {}

// EnterCAT is called when production CAT is entered.
func (s *BaseGoDiskGrammarListener) EnterCAT(ctx *CATContext) {}

// ExitCAT is called when production CAT is exited.
func (s *BaseGoDiskGrammarListener) ExitCAT(ctx *CATContext) {}

// EnterLOGIN is called when production LOGIN is entered.
func (s *BaseGoDiskGrammarListener) EnterLOGIN(ctx *LOGINContext) {}

// ExitLOGIN is called when production LOGIN is exited.
func (s *BaseGoDiskGrammarListener) ExitLOGIN(ctx *LOGINContext) {}

// EnterLOGOUT is called when production LOGOUT is entered.
func (s *BaseGoDiskGrammarListener) EnterLOGOUT(ctx *LOGOUTContext) {}

// ExitLOGOUT is called when production LOGOUT is exited.
func (s *BaseGoDiskGrammarListener) ExitLOGOUT(ctx *LOGOUTContext) {}

// EnterMKGRP is called when production MKGRP is entered.
func (s *BaseGoDiskGrammarListener) EnterMKGRP(ctx *MKGRPContext) {}

// ExitMKGRP is called when production MKGRP is exited.
func (s *BaseGoDiskGrammarListener) ExitMKGRP(ctx *MKGRPContext) {}

// EnterRMGRP is called when production RMGRP is entered.
func (s *BaseGoDiskGrammarListener) EnterRMGRP(ctx *RMGRPContext) {}

// ExitRMGRP is called when production RMGRP is exited.
func (s *BaseGoDiskGrammarListener) ExitRMGRP(ctx *RMGRPContext) {}

// EnterMKUSR is called when production MKUSR is entered.
func (s *BaseGoDiskGrammarListener) EnterMKUSR(ctx *MKUSRContext) {}

// ExitMKUSR is called when production MKUSR is exited.
func (s *BaseGoDiskGrammarListener) ExitMKUSR(ctx *MKUSRContext) {}

// EnterRMUSR is called when production RMUSR is entered.
func (s *BaseGoDiskGrammarListener) EnterRMUSR(ctx *RMUSRContext) {}

// ExitRMUSR is called when production RMUSR is exited.
func (s *BaseGoDiskGrammarListener) ExitRMUSR(ctx *RMUSRContext) {}

// EnterCHGRP is called when production CHGRP is entered.
func (s *BaseGoDiskGrammarListener) EnterCHGRP(ctx *CHGRPContext) {}

// ExitCHGRP is called when production CHGRP is exited.
func (s *BaseGoDiskGrammarListener) ExitCHGRP(ctx *CHGRPContext) {}

// EnterMKFILE is called when production MKFILE is entered.
func (s *BaseGoDiskGrammarListener) EnterMKFILE(ctx *MKFILEContext) {}

// ExitMKFILE is called when production MKFILE is exited.
func (s *BaseGoDiskGrammarListener) ExitMKFILE(ctx *MKFILEContext) {}

// EnterMKDIR is called when production MKDIR is entered.
func (s *BaseGoDiskGrammarListener) EnterMKDIR(ctx *MKDIRContext) {}

// ExitMKDIR is called when production MKDIR is exited.
func (s *BaseGoDiskGrammarListener) ExitMKDIR(ctx *MKDIRContext) {}

// EnterREP is called when production REP is entered.
func (s *BaseGoDiskGrammarListener) EnterREP(ctx *REPContext) {}

// ExitREP is called when production REP is exited.
func (s *BaseGoDiskGrammarListener) ExitREP(ctx *REPContext) {}

// EnterSize is called when production size is entered.
func (s *BaseGoDiskGrammarListener) EnterSize(ctx *SizeContext) {}

// ExitSize is called when production size is exited.
func (s *BaseGoDiskGrammarListener) ExitSize(ctx *SizeContext) {}

// EnterFit is called when production fit is entered.
func (s *BaseGoDiskGrammarListener) EnterFit(ctx *FitContext) {}

// ExitFit is called when production fit is exited.
func (s *BaseGoDiskGrammarListener) ExitFit(ctx *FitContext) {}

// EnterUnit is called when production unit is entered.
func (s *BaseGoDiskGrammarListener) EnterUnit(ctx *UnitContext) {}

// ExitUnit is called when production unit is exited.
func (s *BaseGoDiskGrammarListener) ExitUnit(ctx *UnitContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoDiskGrammarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoDiskGrammarListener) ExitType(ctx *TypeContext) {}

// EnterId_text is called when production id_text is entered.
func (s *BaseGoDiskGrammarListener) EnterId_text(ctx *Id_textContext) {}

// ExitId_text is called when production id_text is exited.
func (s *BaseGoDiskGrammarListener) ExitId_text(ctx *Id_textContext) {}

// EnterPath is called when production path is entered.
func (s *BaseGoDiskGrammarListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseGoDiskGrammarListener) ExitPath(ctx *PathContext) {}

// EnterName is called when production name is entered.
func (s *BaseGoDiskGrammarListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseGoDiskGrammarListener) ExitName(ctx *NameContext) {}

// EnterFilen is called when production filen is entered.
func (s *BaseGoDiskGrammarListener) EnterFilen(ctx *FilenContext) {}

// ExitFilen is called when production filen is exited.
func (s *BaseGoDiskGrammarListener) ExitFilen(ctx *FilenContext) {}

// EnterUser is called when production user is entered.
func (s *BaseGoDiskGrammarListener) EnterUser(ctx *UserContext) {}

// ExitUser is called when production user is exited.
func (s *BaseGoDiskGrammarListener) ExitUser(ctx *UserContext) {}

// EnterPass is called when production pass is entered.
func (s *BaseGoDiskGrammarListener) EnterPass(ctx *PassContext) {}

// ExitPass is called when production pass is exited.
func (s *BaseGoDiskGrammarListener) ExitPass(ctx *PassContext) {}

// EnterGrp is called when production grp is entered.
func (s *BaseGoDiskGrammarListener) EnterGrp(ctx *GrpContext) {}

// ExitGrp is called when production grp is exited.
func (s *BaseGoDiskGrammarListener) ExitGrp(ctx *GrpContext) {}

// EnterCont is called when production cont is entered.
func (s *BaseGoDiskGrammarListener) EnterCont(ctx *ContContext) {}

// ExitCont is called when production cont is exited.
func (s *BaseGoDiskGrammarListener) ExitCont(ctx *ContContext) {}

// EnterPath_file_ls is called when production path_file_ls is entered.
func (s *BaseGoDiskGrammarListener) EnterPath_file_ls(ctx *Path_file_lsContext) {}

// ExitPath_file_ls is called when production path_file_ls is exited.
func (s *BaseGoDiskGrammarListener) ExitPath_file_ls(ctx *Path_file_lsContext) {}

// EnterP is called when production p is entered.
func (s *BaseGoDiskGrammarListener) EnterP(ctx *PContext) {}

// ExitP is called when production p is exited.
func (s *BaseGoDiskGrammarListener) ExitP(ctx *PContext) {}

// EnterR is called when production r is entered.
func (s *BaseGoDiskGrammarListener) EnterR(ctx *RContext) {}

// ExitR is called when production r is exited.
func (s *BaseGoDiskGrammarListener) ExitR(ctx *RContext) {}

// EnterMkdisk_params is called when production mkdisk_params is entered.
func (s *BaseGoDiskGrammarListener) EnterMkdisk_params(ctx *Mkdisk_paramsContext) {}

// ExitMkdisk_params is called when production mkdisk_params is exited.
func (s *BaseGoDiskGrammarListener) ExitMkdisk_params(ctx *Mkdisk_paramsContext) {}

// EnterFdisk_params is called when production fdisk_params is entered.
func (s *BaseGoDiskGrammarListener) EnterFdisk_params(ctx *Fdisk_paramsContext) {}

// ExitFdisk_params is called when production fdisk_params is exited.
func (s *BaseGoDiskGrammarListener) ExitFdisk_params(ctx *Fdisk_paramsContext) {}

// EnterMount_params is called when production mount_params is entered.
func (s *BaseGoDiskGrammarListener) EnterMount_params(ctx *Mount_paramsContext) {}

// ExitMount_params is called when production mount_params is exited.
func (s *BaseGoDiskGrammarListener) ExitMount_params(ctx *Mount_paramsContext) {}

// EnterMkfs_params is called when production mkfs_params is entered.
func (s *BaseGoDiskGrammarListener) EnterMkfs_params(ctx *Mkfs_paramsContext) {}

// ExitMkfs_params is called when production mkfs_params is exited.
func (s *BaseGoDiskGrammarListener) ExitMkfs_params(ctx *Mkfs_paramsContext) {}

// EnterLogin_params is called when production login_params is entered.
func (s *BaseGoDiskGrammarListener) EnterLogin_params(ctx *Login_paramsContext) {}

// ExitLogin_params is called when production login_params is exited.
func (s *BaseGoDiskGrammarListener) ExitLogin_params(ctx *Login_paramsContext) {}

// EnterCat_params is called when production cat_params is entered.
func (s *BaseGoDiskGrammarListener) EnterCat_params(ctx *Cat_paramsContext) {}

// ExitCat_params is called when production cat_params is exited.
func (s *BaseGoDiskGrammarListener) ExitCat_params(ctx *Cat_paramsContext) {}

// EnterMkusr_params is called when production mkusr_params is entered.
func (s *BaseGoDiskGrammarListener) EnterMkusr_params(ctx *Mkusr_paramsContext) {}

// ExitMkusr_params is called when production mkusr_params is exited.
func (s *BaseGoDiskGrammarListener) ExitMkusr_params(ctx *Mkusr_paramsContext) {}

// EnterChgrp_params is called when production chgrp_params is entered.
func (s *BaseGoDiskGrammarListener) EnterChgrp_params(ctx *Chgrp_paramsContext) {}

// ExitChgrp_params is called when production chgrp_params is exited.
func (s *BaseGoDiskGrammarListener) ExitChgrp_params(ctx *Chgrp_paramsContext) {}

// EnterMkfile_params is called when production mkfile_params is entered.
func (s *BaseGoDiskGrammarListener) EnterMkfile_params(ctx *Mkfile_paramsContext) {}

// ExitMkfile_params is called when production mkfile_params is exited.
func (s *BaseGoDiskGrammarListener) ExitMkfile_params(ctx *Mkfile_paramsContext) {}

// EnterMkdir_params is called when production mkdir_params is entered.
func (s *BaseGoDiskGrammarListener) EnterMkdir_params(ctx *Mkdir_paramsContext) {}

// ExitMkdir_params is called when production mkdir_params is exited.
func (s *BaseGoDiskGrammarListener) ExitMkdir_params(ctx *Mkdir_paramsContext) {}

// EnterRep_params is called when production rep_params is entered.
func (s *BaseGoDiskGrammarListener) EnterRep_params(ctx *Rep_paramsContext) {}

// ExitRep_params is called when production rep_params is exited.
func (s *BaseGoDiskGrammarListener) ExitRep_params(ctx *Rep_paramsContext) {}

// EnterMkdisk_param is called when production mkdisk_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMkdisk_param(ctx *Mkdisk_paramContext) {}

// ExitMkdisk_param is called when production mkdisk_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMkdisk_param(ctx *Mkdisk_paramContext) {}

// EnterRmdisk_param is called when production rmdisk_param is entered.
func (s *BaseGoDiskGrammarListener) EnterRmdisk_param(ctx *Rmdisk_paramContext) {}

// ExitRmdisk_param is called when production rmdisk_param is exited.
func (s *BaseGoDiskGrammarListener) ExitRmdisk_param(ctx *Rmdisk_paramContext) {}

// EnterFdisk_param is called when production fdisk_param is entered.
func (s *BaseGoDiskGrammarListener) EnterFdisk_param(ctx *Fdisk_paramContext) {}

// ExitFdisk_param is called when production fdisk_param is exited.
func (s *BaseGoDiskGrammarListener) ExitFdisk_param(ctx *Fdisk_paramContext) {}

// EnterMount_param is called when production mount_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMount_param(ctx *Mount_paramContext) {}

// ExitMount_param is called when production mount_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMount_param(ctx *Mount_paramContext) {}

// EnterMkfs_param is called when production mkfs_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMkfs_param(ctx *Mkfs_paramContext) {}

// ExitMkfs_param is called when production mkfs_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMkfs_param(ctx *Mkfs_paramContext) {}

// EnterCat_param is called when production cat_param is entered.
func (s *BaseGoDiskGrammarListener) EnterCat_param(ctx *Cat_paramContext) {}

// ExitCat_param is called when production cat_param is exited.
func (s *BaseGoDiskGrammarListener) ExitCat_param(ctx *Cat_paramContext) {}

// EnterLogin_param is called when production login_param is entered.
func (s *BaseGoDiskGrammarListener) EnterLogin_param(ctx *Login_paramContext) {}

// ExitLogin_param is called when production login_param is exited.
func (s *BaseGoDiskGrammarListener) ExitLogin_param(ctx *Login_paramContext) {}

// EnterMkgrp_param is called when production mkgrp_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMkgrp_param(ctx *Mkgrp_paramContext) {}

// ExitMkgrp_param is called when production mkgrp_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMkgrp_param(ctx *Mkgrp_paramContext) {}

// EnterRmgrp_param is called when production rmgrp_param is entered.
func (s *BaseGoDiskGrammarListener) EnterRmgrp_param(ctx *Rmgrp_paramContext) {}

// ExitRmgrp_param is called when production rmgrp_param is exited.
func (s *BaseGoDiskGrammarListener) ExitRmgrp_param(ctx *Rmgrp_paramContext) {}

// EnterMkusr_param is called when production mkusr_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMkusr_param(ctx *Mkusr_paramContext) {}

// ExitMkusr_param is called when production mkusr_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMkusr_param(ctx *Mkusr_paramContext) {}

// EnterRmusr_param is called when production rmusr_param is entered.
func (s *BaseGoDiskGrammarListener) EnterRmusr_param(ctx *Rmusr_paramContext) {}

// ExitRmusr_param is called when production rmusr_param is exited.
func (s *BaseGoDiskGrammarListener) ExitRmusr_param(ctx *Rmusr_paramContext) {}

// EnterChgrp_param is called when production chgrp_param is entered.
func (s *BaseGoDiskGrammarListener) EnterChgrp_param(ctx *Chgrp_paramContext) {}

// ExitChgrp_param is called when production chgrp_param is exited.
func (s *BaseGoDiskGrammarListener) ExitChgrp_param(ctx *Chgrp_paramContext) {}

// EnterMkfile_param is called when production mkfile_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMkfile_param(ctx *Mkfile_paramContext) {}

// ExitMkfile_param is called when production mkfile_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMkfile_param(ctx *Mkfile_paramContext) {}

// EnterMkdir_param is called when production mkdir_param is entered.
func (s *BaseGoDiskGrammarListener) EnterMkdir_param(ctx *Mkdir_paramContext) {}

// ExitMkdir_param is called when production mkdir_param is exited.
func (s *BaseGoDiskGrammarListener) ExitMkdir_param(ctx *Mkdir_paramContext) {}

// EnterRep_param is called when production rep_param is entered.
func (s *BaseGoDiskGrammarListener) EnterRep_param(ctx *Rep_paramContext) {}

// ExitRep_param is called when production rep_param is exited.
func (s *BaseGoDiskGrammarListener) ExitRep_param(ctx *Rep_paramContext) {}
