// Code generated from parser/GoDiskGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package godisk // GoDiskGrammar
import "github.com/antlr4-go/antlr/v4"

// GoDiskGrammarListener is a complete listener for a parse tree produced by GoDiskGrammar.
type GoDiskGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterMKDISK is called when entering the MKDISK production.
	EnterMKDISK(c *MKDISKContext)

	// EnterRMDISK is called when entering the RMDISK production.
	EnterRMDISK(c *RMDISKContext)

	// EnterFDISK is called when entering the FDISK production.
	EnterFDISK(c *FDISKContext)

	// EnterMOUNT is called when entering the MOUNT production.
	EnterMOUNT(c *MOUNTContext)

	// EnterMOUNTED is called when entering the MOUNTED production.
	EnterMOUNTED(c *MOUNTEDContext)

	// EnterMKFS is called when entering the MKFS production.
	EnterMKFS(c *MKFSContext)

	// EnterCAT is called when entering the CAT production.
	EnterCAT(c *CATContext)

	// EnterLOGIN is called when entering the LOGIN production.
	EnterLOGIN(c *LOGINContext)

	// EnterLOGOUT is called when entering the LOGOUT production.
	EnterLOGOUT(c *LOGOUTContext)

	// EnterMKGRP is called when entering the MKGRP production.
	EnterMKGRP(c *MKGRPContext)

	// EnterRMGRP is called when entering the RMGRP production.
	EnterRMGRP(c *RMGRPContext)

	// EnterMKUSR is called when entering the MKUSR production.
	EnterMKUSR(c *MKUSRContext)

	// EnterRMUSR is called when entering the RMUSR production.
	EnterRMUSR(c *RMUSRContext)

	// EnterCHGRP is called when entering the CHGRP production.
	EnterCHGRP(c *CHGRPContext)

	// EnterMKFILE is called when entering the MKFILE production.
	EnterMKFILE(c *MKFILEContext)

	// EnterMKDIR is called when entering the MKDIR production.
	EnterMKDIR(c *MKDIRContext)

	// EnterREP is called when entering the REP production.
	EnterREP(c *REPContext)

	// EnterSize is called when entering the size production.
	EnterSize(c *SizeContext)

	// EnterFit is called when entering the fit production.
	EnterFit(c *FitContext)

	// EnterUnit is called when entering the unit production.
	EnterUnit(c *UnitContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterId_text is called when entering the id_text production.
	EnterId_text(c *Id_textContext)

	// EnterPath is called when entering the path production.
	EnterPath(c *PathContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterFilen is called when entering the filen production.
	EnterFilen(c *FilenContext)

	// EnterUser is called when entering the user production.
	EnterUser(c *UserContext)

	// EnterPass is called when entering the pass production.
	EnterPass(c *PassContext)

	// EnterGrp is called when entering the grp production.
	EnterGrp(c *GrpContext)

	// EnterCont is called when entering the cont production.
	EnterCont(c *ContContext)

	// EnterPath_file_ls is called when entering the path_file_ls production.
	EnterPath_file_ls(c *Path_file_lsContext)

	// EnterP is called when entering the p production.
	EnterP(c *PContext)

	// EnterR is called when entering the r production.
	EnterR(c *RContext)

	// EnterMkdisk_params is called when entering the mkdisk_params production.
	EnterMkdisk_params(c *Mkdisk_paramsContext)

	// EnterFdisk_params is called when entering the fdisk_params production.
	EnterFdisk_params(c *Fdisk_paramsContext)

	// EnterMount_params is called when entering the mount_params production.
	EnterMount_params(c *Mount_paramsContext)

	// EnterMkfs_params is called when entering the mkfs_params production.
	EnterMkfs_params(c *Mkfs_paramsContext)

	// EnterLogin_params is called when entering the login_params production.
	EnterLogin_params(c *Login_paramsContext)

	// EnterCat_params is called when entering the cat_params production.
	EnterCat_params(c *Cat_paramsContext)

	// EnterMkusr_params is called when entering the mkusr_params production.
	EnterMkusr_params(c *Mkusr_paramsContext)

	// EnterChgrp_params is called when entering the chgrp_params production.
	EnterChgrp_params(c *Chgrp_paramsContext)

	// EnterMkfile_params is called when entering the mkfile_params production.
	EnterMkfile_params(c *Mkfile_paramsContext)

	// EnterMkdir_params is called when entering the mkdir_params production.
	EnterMkdir_params(c *Mkdir_paramsContext)

	// EnterRep_params is called when entering the rep_params production.
	EnterRep_params(c *Rep_paramsContext)

	// EnterMkdisk_param is called when entering the mkdisk_param production.
	EnterMkdisk_param(c *Mkdisk_paramContext)

	// EnterRmdisk_param is called when entering the rmdisk_param production.
	EnterRmdisk_param(c *Rmdisk_paramContext)

	// EnterFdisk_param is called when entering the fdisk_param production.
	EnterFdisk_param(c *Fdisk_paramContext)

	// EnterMount_param is called when entering the mount_param production.
	EnterMount_param(c *Mount_paramContext)

	// EnterMkfs_param is called when entering the mkfs_param production.
	EnterMkfs_param(c *Mkfs_paramContext)

	// EnterCat_param is called when entering the cat_param production.
	EnterCat_param(c *Cat_paramContext)

	// EnterLogin_param is called when entering the login_param production.
	EnterLogin_param(c *Login_paramContext)

	// EnterMkgrp_param is called when entering the mkgrp_param production.
	EnterMkgrp_param(c *Mkgrp_paramContext)

	// EnterRmgrp_param is called when entering the rmgrp_param production.
	EnterRmgrp_param(c *Rmgrp_paramContext)

	// EnterMkusr_param is called when entering the mkusr_param production.
	EnterMkusr_param(c *Mkusr_paramContext)

	// EnterRmusr_param is called when entering the rmusr_param production.
	EnterRmusr_param(c *Rmusr_paramContext)

	// EnterChgrp_param is called when entering the chgrp_param production.
	EnterChgrp_param(c *Chgrp_paramContext)

	// EnterMkfile_param is called when entering the mkfile_param production.
	EnterMkfile_param(c *Mkfile_paramContext)

	// EnterMkdir_param is called when entering the mkdir_param production.
	EnterMkdir_param(c *Mkdir_paramContext)

	// EnterRep_param is called when entering the rep_param production.
	EnterRep_param(c *Rep_paramContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitMKDISK is called when exiting the MKDISK production.
	ExitMKDISK(c *MKDISKContext)

	// ExitRMDISK is called when exiting the RMDISK production.
	ExitRMDISK(c *RMDISKContext)

	// ExitFDISK is called when exiting the FDISK production.
	ExitFDISK(c *FDISKContext)

	// ExitMOUNT is called when exiting the MOUNT production.
	ExitMOUNT(c *MOUNTContext)

	// ExitMOUNTED is called when exiting the MOUNTED production.
	ExitMOUNTED(c *MOUNTEDContext)

	// ExitMKFS is called when exiting the MKFS production.
	ExitMKFS(c *MKFSContext)

	// ExitCAT is called when exiting the CAT production.
	ExitCAT(c *CATContext)

	// ExitLOGIN is called when exiting the LOGIN production.
	ExitLOGIN(c *LOGINContext)

	// ExitLOGOUT is called when exiting the LOGOUT production.
	ExitLOGOUT(c *LOGOUTContext)

	// ExitMKGRP is called when exiting the MKGRP production.
	ExitMKGRP(c *MKGRPContext)

	// ExitRMGRP is called when exiting the RMGRP production.
	ExitRMGRP(c *RMGRPContext)

	// ExitMKUSR is called when exiting the MKUSR production.
	ExitMKUSR(c *MKUSRContext)

	// ExitRMUSR is called when exiting the RMUSR production.
	ExitRMUSR(c *RMUSRContext)

	// ExitCHGRP is called when exiting the CHGRP production.
	ExitCHGRP(c *CHGRPContext)

	// ExitMKFILE is called when exiting the MKFILE production.
	ExitMKFILE(c *MKFILEContext)

	// ExitMKDIR is called when exiting the MKDIR production.
	ExitMKDIR(c *MKDIRContext)

	// ExitREP is called when exiting the REP production.
	ExitREP(c *REPContext)

	// ExitSize is called when exiting the size production.
	ExitSize(c *SizeContext)

	// ExitFit is called when exiting the fit production.
	ExitFit(c *FitContext)

	// ExitUnit is called when exiting the unit production.
	ExitUnit(c *UnitContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitId_text is called when exiting the id_text production.
	ExitId_text(c *Id_textContext)

	// ExitPath is called when exiting the path production.
	ExitPath(c *PathContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitFilen is called when exiting the filen production.
	ExitFilen(c *FilenContext)

	// ExitUser is called when exiting the user production.
	ExitUser(c *UserContext)

	// ExitPass is called when exiting the pass production.
	ExitPass(c *PassContext)

	// ExitGrp is called when exiting the grp production.
	ExitGrp(c *GrpContext)

	// ExitCont is called when exiting the cont production.
	ExitCont(c *ContContext)

	// ExitPath_file_ls is called when exiting the path_file_ls production.
	ExitPath_file_ls(c *Path_file_lsContext)

	// ExitP is called when exiting the p production.
	ExitP(c *PContext)

	// ExitR is called when exiting the r production.
	ExitR(c *RContext)

	// ExitMkdisk_params is called when exiting the mkdisk_params production.
	ExitMkdisk_params(c *Mkdisk_paramsContext)

	// ExitFdisk_params is called when exiting the fdisk_params production.
	ExitFdisk_params(c *Fdisk_paramsContext)

	// ExitMount_params is called when exiting the mount_params production.
	ExitMount_params(c *Mount_paramsContext)

	// ExitMkfs_params is called when exiting the mkfs_params production.
	ExitMkfs_params(c *Mkfs_paramsContext)

	// ExitLogin_params is called when exiting the login_params production.
	ExitLogin_params(c *Login_paramsContext)

	// ExitCat_params is called when exiting the cat_params production.
	ExitCat_params(c *Cat_paramsContext)

	// ExitMkusr_params is called when exiting the mkusr_params production.
	ExitMkusr_params(c *Mkusr_paramsContext)

	// ExitChgrp_params is called when exiting the chgrp_params production.
	ExitChgrp_params(c *Chgrp_paramsContext)

	// ExitMkfile_params is called when exiting the mkfile_params production.
	ExitMkfile_params(c *Mkfile_paramsContext)

	// ExitMkdir_params is called when exiting the mkdir_params production.
	ExitMkdir_params(c *Mkdir_paramsContext)

	// ExitRep_params is called when exiting the rep_params production.
	ExitRep_params(c *Rep_paramsContext)

	// ExitMkdisk_param is called when exiting the mkdisk_param production.
	ExitMkdisk_param(c *Mkdisk_paramContext)

	// ExitRmdisk_param is called when exiting the rmdisk_param production.
	ExitRmdisk_param(c *Rmdisk_paramContext)

	// ExitFdisk_param is called when exiting the fdisk_param production.
	ExitFdisk_param(c *Fdisk_paramContext)

	// ExitMount_param is called when exiting the mount_param production.
	ExitMount_param(c *Mount_paramContext)

	// ExitMkfs_param is called when exiting the mkfs_param production.
	ExitMkfs_param(c *Mkfs_paramContext)

	// ExitCat_param is called when exiting the cat_param production.
	ExitCat_param(c *Cat_paramContext)

	// ExitLogin_param is called when exiting the login_param production.
	ExitLogin_param(c *Login_paramContext)

	// ExitMkgrp_param is called when exiting the mkgrp_param production.
	ExitMkgrp_param(c *Mkgrp_paramContext)

	// ExitRmgrp_param is called when exiting the rmgrp_param production.
	ExitRmgrp_param(c *Rmgrp_paramContext)

	// ExitMkusr_param is called when exiting the mkusr_param production.
	ExitMkusr_param(c *Mkusr_paramContext)

	// ExitRmusr_param is called when exiting the rmusr_param production.
	ExitRmusr_param(c *Rmusr_paramContext)

	// ExitChgrp_param is called when exiting the chgrp_param production.
	ExitChgrp_param(c *Chgrp_paramContext)

	// ExitMkfile_param is called when exiting the mkfile_param production.
	ExitMkfile_param(c *Mkfile_paramContext)

	// ExitMkdir_param is called when exiting the mkdir_param production.
	ExitMkdir_param(c *Mkdir_paramContext)

	// ExitRep_param is called when exiting the rep_param production.
	ExitRep_param(c *Rep_paramContext)
}
