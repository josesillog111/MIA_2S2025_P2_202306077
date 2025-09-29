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

	// EnterUNMOUNT is called when entering the UNMOUNT production.
	EnterUNMOUNT(c *UNMOUNTContext)

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

	// EnterREMOVE is called when entering the REMOVE production.
	EnterREMOVE(c *REMOVEContext)

	// EnterEDIT is called when entering the EDIT production.
	EnterEDIT(c *EDITContext)

	// EnterRENAME is called when entering the RENAME production.
	EnterRENAME(c *RENAMEContext)

	// EnterCOPY is called when entering the COPY production.
	EnterCOPY(c *COPYContext)

	// EnterMOVE is called when entering the MOVE production.
	EnterMOVE(c *MOVEContext)

	// EnterFIND is called when entering the FIND production.
	EnterFIND(c *FINDContext)

	// EnterCHOWN is called when entering the CHOWN production.
	EnterCHOWN(c *CHOWNContext)

	// EnterCHMOD is called when entering the CHMOD production.
	EnterCHMOD(c *CHMODContext)

	// EnterRECOVERY is called when entering the RECOVERY production.
	EnterRECOVERY(c *RECOVERYContext)

	// EnterLOSS is called when entering the LOSS production.
	EnterLOSS(c *LOSSContext)

	// EnterJOURNALING is called when entering the JOURNALING production.
	EnterJOURNALING(c *JOURNALINGContext)

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

	// EnterName_find is called when entering the name_find production.
	EnterName_find(c *Name_findContext)

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

	// EnterDelete is called when entering the delete production.
	EnterDelete(c *DeleteContext)

	// EnterAdd is called when entering the add production.
	EnterAdd(c *AddContext)

	// EnterFs is called when entering the fs production.
	EnterFs(c *FsContext)

	// EnterContenido is called when entering the contenido production.
	EnterContenido(c *ContenidoContext)

	// EnterDestino is called when entering the destino production.
	EnterDestino(c *DestinoContext)

	// EnterUsuario is called when entering the usuario production.
	EnterUsuario(c *UsuarioContext)

	// EnterUgo is called when entering the ugo production.
	EnterUgo(c *UgoContext)

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

	// EnterEdit_params is called when entering the edit_params production.
	EnterEdit_params(c *Edit_paramsContext)

	// EnterRename_params is called when entering the rename_params production.
	EnterRename_params(c *Rename_paramsContext)

	// EnterCopy_params is called when entering the copy_params production.
	EnterCopy_params(c *Copy_paramsContext)

	// EnterMove_params is called when entering the move_params production.
	EnterMove_params(c *Move_paramsContext)

	// EnterFind_params is called when entering the find_params production.
	EnterFind_params(c *Find_paramsContext)

	// EnterChown_params is called when entering the chown_params production.
	EnterChown_params(c *Chown_paramsContext)

	// EnterChmod_params is called when entering the chmod_params production.
	EnterChmod_params(c *Chmod_paramsContext)

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

	// EnterUnmount_param is called when entering the unmount_param production.
	EnterUnmount_param(c *Unmount_paramContext)

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

	// EnterRemove_param is called when entering the remove_param production.
	EnterRemove_param(c *Remove_paramContext)

	// EnterEdit_param is called when entering the edit_param production.
	EnterEdit_param(c *Edit_paramContext)

	// EnterRename_param is called when entering the rename_param production.
	EnterRename_param(c *Rename_paramContext)

	// EnterCopy_param is called when entering the copy_param production.
	EnterCopy_param(c *Copy_paramContext)

	// EnterMove_param is called when entering the move_param production.
	EnterMove_param(c *Move_paramContext)

	// EnterFind_param is called when entering the find_param production.
	EnterFind_param(c *Find_paramContext)

	// EnterChown_param is called when entering the chown_param production.
	EnterChown_param(c *Chown_paramContext)

	// EnterChmod_param is called when entering the chmod_param production.
	EnterChmod_param(c *Chmod_paramContext)

	// EnterRecovery_param is called when entering the recovery_param production.
	EnterRecovery_param(c *Recovery_paramContext)

	// EnterLoss_param is called when entering the loss_param production.
	EnterLoss_param(c *Loss_paramContext)

	// EnterJournaling_param is called when entering the journaling_param production.
	EnterJournaling_param(c *Journaling_paramContext)

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

	// ExitUNMOUNT is called when exiting the UNMOUNT production.
	ExitUNMOUNT(c *UNMOUNTContext)

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

	// ExitREMOVE is called when exiting the REMOVE production.
	ExitREMOVE(c *REMOVEContext)

	// ExitEDIT is called when exiting the EDIT production.
	ExitEDIT(c *EDITContext)

	// ExitRENAME is called when exiting the RENAME production.
	ExitRENAME(c *RENAMEContext)

	// ExitCOPY is called when exiting the COPY production.
	ExitCOPY(c *COPYContext)

	// ExitMOVE is called when exiting the MOVE production.
	ExitMOVE(c *MOVEContext)

	// ExitFIND is called when exiting the FIND production.
	ExitFIND(c *FINDContext)

	// ExitCHOWN is called when exiting the CHOWN production.
	ExitCHOWN(c *CHOWNContext)

	// ExitCHMOD is called when exiting the CHMOD production.
	ExitCHMOD(c *CHMODContext)

	// ExitRECOVERY is called when exiting the RECOVERY production.
	ExitRECOVERY(c *RECOVERYContext)

	// ExitLOSS is called when exiting the LOSS production.
	ExitLOSS(c *LOSSContext)

	// ExitJOURNALING is called when exiting the JOURNALING production.
	ExitJOURNALING(c *JOURNALINGContext)

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

	// ExitName_find is called when exiting the name_find production.
	ExitName_find(c *Name_findContext)

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

	// ExitDelete is called when exiting the delete production.
	ExitDelete(c *DeleteContext)

	// ExitAdd is called when exiting the add production.
	ExitAdd(c *AddContext)

	// ExitFs is called when exiting the fs production.
	ExitFs(c *FsContext)

	// ExitContenido is called when exiting the contenido production.
	ExitContenido(c *ContenidoContext)

	// ExitDestino is called when exiting the destino production.
	ExitDestino(c *DestinoContext)

	// ExitUsuario is called when exiting the usuario production.
	ExitUsuario(c *UsuarioContext)

	// ExitUgo is called when exiting the ugo production.
	ExitUgo(c *UgoContext)

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

	// ExitEdit_params is called when exiting the edit_params production.
	ExitEdit_params(c *Edit_paramsContext)

	// ExitRename_params is called when exiting the rename_params production.
	ExitRename_params(c *Rename_paramsContext)

	// ExitCopy_params is called when exiting the copy_params production.
	ExitCopy_params(c *Copy_paramsContext)

	// ExitMove_params is called when exiting the move_params production.
	ExitMove_params(c *Move_paramsContext)

	// ExitFind_params is called when exiting the find_params production.
	ExitFind_params(c *Find_paramsContext)

	// ExitChown_params is called when exiting the chown_params production.
	ExitChown_params(c *Chown_paramsContext)

	// ExitChmod_params is called when exiting the chmod_params production.
	ExitChmod_params(c *Chmod_paramsContext)

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

	// ExitUnmount_param is called when exiting the unmount_param production.
	ExitUnmount_param(c *Unmount_paramContext)

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

	// ExitRemove_param is called when exiting the remove_param production.
	ExitRemove_param(c *Remove_paramContext)

	// ExitEdit_param is called when exiting the edit_param production.
	ExitEdit_param(c *Edit_paramContext)

	// ExitRename_param is called when exiting the rename_param production.
	ExitRename_param(c *Rename_paramContext)

	// ExitCopy_param is called when exiting the copy_param production.
	ExitCopy_param(c *Copy_paramContext)

	// ExitMove_param is called when exiting the move_param production.
	ExitMove_param(c *Move_paramContext)

	// ExitFind_param is called when exiting the find_param production.
	ExitFind_param(c *Find_paramContext)

	// ExitChown_param is called when exiting the chown_param production.
	ExitChown_param(c *Chown_paramContext)

	// ExitChmod_param is called when exiting the chmod_param production.
	ExitChmod_param(c *Chmod_paramContext)

	// ExitRecovery_param is called when exiting the recovery_param production.
	ExitRecovery_param(c *Recovery_paramContext)

	// ExitLoss_param is called when exiting the loss_param production.
	ExitLoss_param(c *Loss_paramContext)

	// ExitJournaling_param is called when exiting the journaling_param production.
	ExitJournaling_param(c *Journaling_paramContext)

	// ExitRep_param is called when exiting the rep_param production.
	ExitRep_param(c *Rep_paramContext)
}
