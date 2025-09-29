// Generated from /home/jose06/Documentos/MIA_2S2025_P2_202306077/backend/parser/GoDiskGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GoDiskGrammar extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		INT_LIT=1, FLOAT_LIT=2, STRING_LIT=3, CHAR_LIT=4, MKDISK=5, RMDISK=6, 
		FDISK=7, MOUNT=8, MOUNTED=9, UNMOUNT=10, MKFS=11, CAT=12, LOGIN=13, LOGOUT=14, 
		MKGRP=15, RMGRP=16, MKUSR=17, RMUSR=18, CHGRP=19, MKFILE=20, MKDIR=21, 
		REMOVE=22, EDIT=23, RENAME=24, COPY=25, MOVE=26, FIND=27, CHOWN=28, CHMOD=29, 
		RECOVERY=30, LOSS=31, JOURNALING=32, REP=33, ID_TEXT=34, SIZE=35, FIT=36, 
		UNIT=37, PATH=38, TYPE=39, NAME=40, FILE=41, USER=42, PASS=43, GRP=44, 
		CONT=45, PATH_FILE_LS=46, DELETE=47, ADD=48, FS=49, CONTENIDO=50, DESTINO=51, 
		USUARIO=52, UGO=53, R=54, P=55, ASSIGN=56, MINUS=57, ID=58, UNQUOTED_TEXT=59, 
		LINE_COMMENT=60, BLOCK_COMMENT=61, WS=62;
	public static final int
		RULE_prog = 0, RULE_stmt = 1, RULE_size = 2, RULE_fit = 3, RULE_unit = 4, 
		RULE_type = 5, RULE_id_text = 6, RULE_path = 7, RULE_name = 8, RULE_filen = 9, 
		RULE_user = 10, RULE_pass = 11, RULE_grp = 12, RULE_cont = 13, RULE_path_file_ls = 14, 
		RULE_delete = 15, RULE_add = 16, RULE_fs = 17, RULE_contenido = 18, RULE_destino = 19, 
		RULE_usuario = 20, RULE_ugo = 21, RULE_p = 22, RULE_r = 23, RULE_mkdisk_params = 24, 
		RULE_fdisk_params = 25, RULE_mount_params = 26, RULE_mkfs_params = 27, 
		RULE_login_params = 28, RULE_cat_params = 29, RULE_mkusr_params = 30, 
		RULE_chgrp_params = 31, RULE_mkfile_params = 32, RULE_mkdir_params = 33, 
		RULE_edit_params = 34, RULE_rename_params = 35, RULE_copy_params = 36, 
		RULE_move_params = 37, RULE_find_params = 38, RULE_chown_params = 39, 
		RULE_chmod_params = 40, RULE_rep_params = 41, RULE_mkdisk_param = 42, 
		RULE_rmdisk_param = 43, RULE_fdisk_param = 44, RULE_mount_param = 45, 
		RULE_unmount_param = 46, RULE_mkfs_param = 47, RULE_cat_param = 48, RULE_login_param = 49, 
		RULE_mkgrp_param = 50, RULE_rmgrp_param = 51, RULE_mkusr_param = 52, RULE_rmusr_param = 53, 
		RULE_chgrp_param = 54, RULE_mkfile_param = 55, RULE_mkdir_param = 56, 
		RULE_remove_param = 57, RULE_edit_param = 58, RULE_rename_param = 59, 
		RULE_copy_param = 60, RULE_move_param = 61, RULE_find_param = 62, RULE_chown_param = 63, 
		RULE_chmod_param = 64, RULE_recovery_param = 65, RULE_loss_param = 66, 
		RULE_journaling_param = 67, RULE_rep_param = 68;
	private static String[] makeRuleNames() {
		return new String[] {
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
			"recovery_param", "loss_param", "journaling_param", "rep_param"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "'='", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT_LIT", "FLOAT_LIT", "STRING_LIT", "CHAR_LIT", "MKDISK", "RMDISK", 
			"FDISK", "MOUNT", "MOUNTED", "UNMOUNT", "MKFS", "CAT", "LOGIN", "LOGOUT", 
			"MKGRP", "RMGRP", "MKUSR", "RMUSR", "CHGRP", "MKFILE", "MKDIR", "REMOVE", 
			"EDIT", "RENAME", "COPY", "MOVE", "FIND", "CHOWN", "CHMOD", "RECOVERY", 
			"LOSS", "JOURNALING", "REP", "ID_TEXT", "SIZE", "FIT", "UNIT", "PATH", 
			"TYPE", "NAME", "FILE", "USER", "PASS", "GRP", "CONT", "PATH_FILE_LS", 
			"DELETE", "ADD", "FS", "CONTENIDO", "DESTINO", "USUARIO", "UGO", "R", 
			"P", "ASSIGN", "MINUS", "ID", "UNQUOTED_TEXT", "LINE_COMMENT", "BLOCK_COMMENT", 
			"WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "GoDiskGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoDiskGrammar(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(GoDiskGrammar.EOF, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 17179869152L) != 0)) {
				{
				{
				setState(138);
				stmt();
				}
				}
				setState(143);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(144);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	 
		public StmtContext() { }
		public void copyFrom(StmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LOSSContext extends StmtContext {
		public TerminalNode LOSS() { return getToken(GoDiskGrammar.LOSS, 0); }
		public Loss_paramContext loss_param() {
			return getRuleContext(Loss_paramContext.class,0);
		}
		public LOSSContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MKGRPContext extends StmtContext {
		public TerminalNode MKGRP() { return getToken(GoDiskGrammar.MKGRP, 0); }
		public Mkgrp_paramContext mkgrp_param() {
			return getRuleContext(Mkgrp_paramContext.class,0);
		}
		public MKGRPContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RMGRPContext extends StmtContext {
		public TerminalNode RMGRP() { return getToken(GoDiskGrammar.RMGRP, 0); }
		public Rmgrp_paramContext rmgrp_param() {
			return getRuleContext(Rmgrp_paramContext.class,0);
		}
		public RMGRPContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MKFSContext extends StmtContext {
		public TerminalNode MKFS() { return getToken(GoDiskGrammar.MKFS, 0); }
		public Mkfs_paramsContext mkfs_params() {
			return getRuleContext(Mkfs_paramsContext.class,0);
		}
		public MKFSContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CHMODContext extends StmtContext {
		public TerminalNode CHMOD() { return getToken(GoDiskGrammar.CHMOD, 0); }
		public Chmod_paramsContext chmod_params() {
			return getRuleContext(Chmod_paramsContext.class,0);
		}
		public CHMODContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CHGRPContext extends StmtContext {
		public TerminalNode CHGRP() { return getToken(GoDiskGrammar.CHGRP, 0); }
		public Chgrp_paramsContext chgrp_params() {
			return getRuleContext(Chgrp_paramsContext.class,0);
		}
		public CHGRPContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RECOVERYContext extends StmtContext {
		public TerminalNode RECOVERY() { return getToken(GoDiskGrammar.RECOVERY, 0); }
		public Recovery_paramContext recovery_param() {
			return getRuleContext(Recovery_paramContext.class,0);
		}
		public RECOVERYContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RMDISKContext extends StmtContext {
		public TerminalNode RMDISK() { return getToken(GoDiskGrammar.RMDISK, 0); }
		public Rmdisk_paramContext rmdisk_param() {
			return getRuleContext(Rmdisk_paramContext.class,0);
		}
		public RMDISKContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FDISKContext extends StmtContext {
		public TerminalNode FDISK() { return getToken(GoDiskGrammar.FDISK, 0); }
		public Fdisk_paramsContext fdisk_params() {
			return getRuleContext(Fdisk_paramsContext.class,0);
		}
		public FDISKContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class REMOVEContext extends StmtContext {
		public TerminalNode REMOVE() { return getToken(GoDiskGrammar.REMOVE, 0); }
		public Remove_paramContext remove_param() {
			return getRuleContext(Remove_paramContext.class,0);
		}
		public REMOVEContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FINDContext extends StmtContext {
		public TerminalNode FIND() { return getToken(GoDiskGrammar.FIND, 0); }
		public Find_paramsContext find_params() {
			return getRuleContext(Find_paramsContext.class,0);
		}
		public FINDContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LOGINContext extends StmtContext {
		public TerminalNode LOGIN() { return getToken(GoDiskGrammar.LOGIN, 0); }
		public Login_paramsContext login_params() {
			return getRuleContext(Login_paramsContext.class,0);
		}
		public LOGINContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MOUNTContext extends StmtContext {
		public TerminalNode MOUNT() { return getToken(GoDiskGrammar.MOUNT, 0); }
		public Mount_paramsContext mount_params() {
			return getRuleContext(Mount_paramsContext.class,0);
		}
		public MOUNTContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MKDISKContext extends StmtContext {
		public TerminalNode MKDISK() { return getToken(GoDiskGrammar.MKDISK, 0); }
		public Mkdisk_paramsContext mkdisk_params() {
			return getRuleContext(Mkdisk_paramsContext.class,0);
		}
		public MKDISKContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RENAMEContext extends StmtContext {
		public TerminalNode RENAME() { return getToken(GoDiskGrammar.RENAME, 0); }
		public Rename_paramsContext rename_params() {
			return getRuleContext(Rename_paramsContext.class,0);
		}
		public RENAMEContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LOGOUTContext extends StmtContext {
		public TerminalNode LOGOUT() { return getToken(GoDiskGrammar.LOGOUT, 0); }
		public LOGOUTContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class JOURNALINGContext extends StmtContext {
		public TerminalNode JOURNALING() { return getToken(GoDiskGrammar.JOURNALING, 0); }
		public Journaling_paramContext journaling_param() {
			return getRuleContext(Journaling_paramContext.class,0);
		}
		public JOURNALINGContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EDITContext extends StmtContext {
		public TerminalNode EDIT() { return getToken(GoDiskGrammar.EDIT, 0); }
		public Edit_paramsContext edit_params() {
			return getRuleContext(Edit_paramsContext.class,0);
		}
		public EDITContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class COPYContext extends StmtContext {
		public TerminalNode COPY() { return getToken(GoDiskGrammar.COPY, 0); }
		public Copy_paramsContext copy_params() {
			return getRuleContext(Copy_paramsContext.class,0);
		}
		public COPYContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UNMOUNTContext extends StmtContext {
		public TerminalNode UNMOUNT() { return getToken(GoDiskGrammar.UNMOUNT, 0); }
		public Unmount_paramContext unmount_param() {
			return getRuleContext(Unmount_paramContext.class,0);
		}
		public UNMOUNTContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MKDIRContext extends StmtContext {
		public TerminalNode MKDIR() { return getToken(GoDiskGrammar.MKDIR, 0); }
		public Mkdir_paramsContext mkdir_params() {
			return getRuleContext(Mkdir_paramsContext.class,0);
		}
		public MKDIRContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CHOWNContext extends StmtContext {
		public TerminalNode CHOWN() { return getToken(GoDiskGrammar.CHOWN, 0); }
		public Chown_paramsContext chown_params() {
			return getRuleContext(Chown_paramsContext.class,0);
		}
		public CHOWNContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MOVEContext extends StmtContext {
		public TerminalNode MOVE() { return getToken(GoDiskGrammar.MOVE, 0); }
		public Move_paramsContext move_params() {
			return getRuleContext(Move_paramsContext.class,0);
		}
		public MOVEContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MKUSRContext extends StmtContext {
		public TerminalNode MKUSR() { return getToken(GoDiskGrammar.MKUSR, 0); }
		public Mkusr_paramsContext mkusr_params() {
			return getRuleContext(Mkusr_paramsContext.class,0);
		}
		public MKUSRContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CATContext extends StmtContext {
		public TerminalNode CAT() { return getToken(GoDiskGrammar.CAT, 0); }
		public Cat_paramsContext cat_params() {
			return getRuleContext(Cat_paramsContext.class,0);
		}
		public CATContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MOUNTEDContext extends StmtContext {
		public TerminalNode MOUNTED() { return getToken(GoDiskGrammar.MOUNTED, 0); }
		public MOUNTEDContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MKFILEContext extends StmtContext {
		public TerminalNode MKFILE() { return getToken(GoDiskGrammar.MKFILE, 0); }
		public Mkfile_paramsContext mkfile_params() {
			return getRuleContext(Mkfile_paramsContext.class,0);
		}
		public MKFILEContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class REPContext extends StmtContext {
		public TerminalNode REP() { return getToken(GoDiskGrammar.REP, 0); }
		public Rep_paramsContext rep_params() {
			return getRuleContext(Rep_paramsContext.class,0);
		}
		public REPContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RMUSRContext extends StmtContext {
		public TerminalNode RMUSR() { return getToken(GoDiskGrammar.RMUSR, 0); }
		public Rmusr_paramContext rmusr_param() {
			return getRuleContext(Rmusr_paramContext.class,0);
		}
		public RMUSRContext(StmtContext ctx) { copyFrom(ctx); }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(202);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
				_localctx = new MKDISKContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(146);
				match(MKDISK);
				setState(147);
				mkdisk_params();
				}
				break;
			case RMDISK:
				_localctx = new RMDISKContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(148);
				match(RMDISK);
				setState(149);
				rmdisk_param();
				}
				break;
			case FDISK:
				_localctx = new FDISKContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(150);
				match(FDISK);
				setState(151);
				fdisk_params();
				}
				break;
			case MOUNT:
				_localctx = new MOUNTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(152);
				match(MOUNT);
				setState(153);
				mount_params();
				}
				break;
			case MOUNTED:
				_localctx = new MOUNTEDContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(154);
				match(MOUNTED);
				}
				break;
			case UNMOUNT:
				_localctx = new UNMOUNTContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(155);
				match(UNMOUNT);
				setState(156);
				unmount_param();
				}
				break;
			case MKFS:
				_localctx = new MKFSContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(157);
				match(MKFS);
				setState(158);
				mkfs_params();
				}
				break;
			case CAT:
				_localctx = new CATContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(159);
				match(CAT);
				setState(160);
				cat_params();
				}
				break;
			case LOGIN:
				_localctx = new LOGINContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(161);
				match(LOGIN);
				setState(162);
				login_params();
				}
				break;
			case LOGOUT:
				_localctx = new LOGOUTContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(163);
				match(LOGOUT);
				}
				break;
			case MKGRP:
				_localctx = new MKGRPContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(164);
				match(MKGRP);
				setState(165);
				mkgrp_param();
				}
				break;
			case RMGRP:
				_localctx = new RMGRPContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(166);
				match(RMGRP);
				setState(167);
				rmgrp_param();
				}
				break;
			case MKUSR:
				_localctx = new MKUSRContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(168);
				match(MKUSR);
				setState(169);
				mkusr_params();
				}
				break;
			case RMUSR:
				_localctx = new RMUSRContext(_localctx);
				enterOuterAlt(_localctx, 14);
				{
				setState(170);
				match(RMUSR);
				setState(171);
				rmusr_param();
				}
				break;
			case CHGRP:
				_localctx = new CHGRPContext(_localctx);
				enterOuterAlt(_localctx, 15);
				{
				setState(172);
				match(CHGRP);
				setState(173);
				chgrp_params();
				}
				break;
			case MKFILE:
				_localctx = new MKFILEContext(_localctx);
				enterOuterAlt(_localctx, 16);
				{
				setState(174);
				match(MKFILE);
				setState(175);
				mkfile_params();
				}
				break;
			case MKDIR:
				_localctx = new MKDIRContext(_localctx);
				enterOuterAlt(_localctx, 17);
				{
				setState(176);
				match(MKDIR);
				setState(177);
				mkdir_params();
				}
				break;
			case REMOVE:
				_localctx = new REMOVEContext(_localctx);
				enterOuterAlt(_localctx, 18);
				{
				setState(178);
				match(REMOVE);
				setState(179);
				remove_param();
				}
				break;
			case EDIT:
				_localctx = new EDITContext(_localctx);
				enterOuterAlt(_localctx, 19);
				{
				setState(180);
				match(EDIT);
				setState(181);
				edit_params();
				}
				break;
			case RENAME:
				_localctx = new RENAMEContext(_localctx);
				enterOuterAlt(_localctx, 20);
				{
				setState(182);
				match(RENAME);
				setState(183);
				rename_params();
				}
				break;
			case COPY:
				_localctx = new COPYContext(_localctx);
				enterOuterAlt(_localctx, 21);
				{
				setState(184);
				match(COPY);
				setState(185);
				copy_params();
				}
				break;
			case MOVE:
				_localctx = new MOVEContext(_localctx);
				enterOuterAlt(_localctx, 22);
				{
				setState(186);
				match(MOVE);
				setState(187);
				move_params();
				}
				break;
			case FIND:
				_localctx = new FINDContext(_localctx);
				enterOuterAlt(_localctx, 23);
				{
				setState(188);
				match(FIND);
				setState(189);
				find_params();
				}
				break;
			case CHOWN:
				_localctx = new CHOWNContext(_localctx);
				enterOuterAlt(_localctx, 24);
				{
				setState(190);
				match(CHOWN);
				setState(191);
				chown_params();
				}
				break;
			case CHMOD:
				_localctx = new CHMODContext(_localctx);
				enterOuterAlt(_localctx, 25);
				{
				setState(192);
				match(CHMOD);
				setState(193);
				chmod_params();
				}
				break;
			case RECOVERY:
				_localctx = new RECOVERYContext(_localctx);
				enterOuterAlt(_localctx, 26);
				{
				setState(194);
				match(RECOVERY);
				setState(195);
				recovery_param();
				}
				break;
			case LOSS:
				_localctx = new LOSSContext(_localctx);
				enterOuterAlt(_localctx, 27);
				{
				setState(196);
				match(LOSS);
				setState(197);
				loss_param();
				}
				break;
			case JOURNALING:
				_localctx = new JOURNALINGContext(_localctx);
				enterOuterAlt(_localctx, 28);
				{
				setState(198);
				match(JOURNALING);
				setState(199);
				journaling_param();
				}
				break;
			case REP:
				_localctx = new REPContext(_localctx);
				enterOuterAlt(_localctx, 29);
				{
				setState(200);
				match(REP);
				setState(201);
				rep_params();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SizeContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode SIZE() { return getToken(GoDiskGrammar.SIZE, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode INT_LIT() { return getToken(GoDiskGrammar.INT_LIT, 0); }
		public SizeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_size; }
	}

	public final SizeContext size() throws RecognitionException {
		SizeContext _localctx = new SizeContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_size);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(MINUS);
			setState(205);
			match(SIZE);
			setState(206);
			match(ASSIGN);
			setState(207);
			match(INT_LIT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FitContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode FIT() { return getToken(GoDiskGrammar.FIT, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public FitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fit; }
	}

	public final FitContext fit() throws RecognitionException {
		FitContext _localctx = new FitContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fit);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(209);
			match(MINUS);
			setState(210);
			match(FIT);
			setState(211);
			match(ASSIGN);
			setState(212);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UnitContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode UNIT() { return getToken(GoDiskGrammar.UNIT, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public UnitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unit; }
	}

	public final UnitContext unit() throws RecognitionException {
		UnitContext _localctx = new UnitContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_unit);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(214);
			match(MINUS);
			setState(215);
			match(UNIT);
			setState(216);
			match(ASSIGN);
			setState(217);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode TYPE() { return getToken(GoDiskGrammar.TYPE, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			match(MINUS);
			setState(220);
			match(TYPE);
			setState(221);
			match(ASSIGN);
			setState(222);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Id_textContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode ID_TEXT() { return getToken(GoDiskGrammar.ID_TEXT, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public Id_textContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_id_text; }
	}

	public final Id_textContext id_text() throws RecognitionException {
		Id_textContext _localctx = new Id_textContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_id_text);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			match(MINUS);
			setState(225);
			match(ID_TEXT);
			setState(226);
			match(ASSIGN);
			setState(227);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PathContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode PATH() { return getToken(GoDiskGrammar.PATH, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public PathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_path; }
	}

	public final PathContext path() throws RecognitionException {
		PathContext _localctx = new PathContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_path);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(229);
			match(MINUS);
			setState(230);
			match(PATH);
			setState(231);
			match(ASSIGN);
			setState(232);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class NameContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode NAME() { return getToken(GoDiskGrammar.NAME, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_name);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(234);
			match(MINUS);
			setState(235);
			match(NAME);
			setState(236);
			match(ASSIGN);
			setState(237);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FilenContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode FILE() { return getToken(GoDiskGrammar.FILE, 0); }
		public TerminalNode INT_LIT() { return getToken(GoDiskGrammar.INT_LIT, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public FilenContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_filen; }
	}

	public final FilenContext filen() throws RecognitionException {
		FilenContext _localctx = new FilenContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_filen);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(MINUS);
			setState(240);
			match(FILE);
			setState(241);
			match(INT_LIT);
			setState(242);
			match(ASSIGN);
			setState(243);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UserContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode USER() { return getToken(GoDiskGrammar.USER, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public UserContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_user; }
	}

	public final UserContext user() throws RecognitionException {
		UserContext _localctx = new UserContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_user);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(245);
			match(MINUS);
			setState(246);
			match(USER);
			setState(247);
			match(ASSIGN);
			setState(248);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PassContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode PASS() { return getToken(GoDiskGrammar.PASS, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public PassContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_pass; }
	}

	public final PassContext pass() throws RecognitionException {
		PassContext _localctx = new PassContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_pass);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(250);
			match(MINUS);
			setState(251);
			match(PASS);
			setState(252);
			match(ASSIGN);
			setState(253);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class GrpContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode GRP() { return getToken(GoDiskGrammar.GRP, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public GrpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_grp; }
	}

	public final GrpContext grp() throws RecognitionException {
		GrpContext _localctx = new GrpContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_grp);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(255);
			match(MINUS);
			setState(256);
			match(GRP);
			setState(257);
			match(ASSIGN);
			setState(258);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode CONT() { return getToken(GoDiskGrammar.CONT, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public ContContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cont; }
	}

	public final ContContext cont() throws RecognitionException {
		ContContext _localctx = new ContContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_cont);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			match(MINUS);
			setState(261);
			match(CONT);
			setState(262);
			match(ASSIGN);
			setState(263);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Path_file_lsContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode PATH_FILE_LS() { return getToken(GoDiskGrammar.PATH_FILE_LS, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public Path_file_lsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_path_file_ls; }
	}

	public final Path_file_lsContext path_file_ls() throws RecognitionException {
		Path_file_lsContext _localctx = new Path_file_lsContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_path_file_ls);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			match(MINUS);
			setState(266);
			match(PATH_FILE_LS);
			setState(267);
			match(ASSIGN);
			setState(268);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeleteContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode DELETE() { return getToken(GoDiskGrammar.DELETE, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public DeleteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delete; }
	}

	public final DeleteContext delete() throws RecognitionException {
		DeleteContext _localctx = new DeleteContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(270);
			match(MINUS);
			setState(271);
			match(DELETE);
			setState(272);
			match(ASSIGN);
			setState(273);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AddContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode ADD() { return getToken(GoDiskGrammar.ADD, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode INT_LIT() { return getToken(GoDiskGrammar.INT_LIT, 0); }
		public AddContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_add; }
	}

	public final AddContext add() throws RecognitionException {
		AddContext _localctx = new AddContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_add);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(275);
			match(MINUS);
			setState(276);
			match(ADD);
			setState(277);
			match(ASSIGN);
			setState(278);
			match(INT_LIT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FsContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode FS() { return getToken(GoDiskGrammar.FS, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public FsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fs; }
	}

	public final FsContext fs() throws RecognitionException {
		FsContext _localctx = new FsContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_fs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(280);
			match(MINUS);
			setState(281);
			match(FS);
			setState(282);
			match(ASSIGN);
			setState(283);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ContenidoContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode CONTENIDO() { return getToken(GoDiskGrammar.CONTENIDO, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public ContenidoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contenido; }
	}

	public final ContenidoContext contenido() throws RecognitionException {
		ContenidoContext _localctx = new ContenidoContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_contenido);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			match(MINUS);
			setState(286);
			match(CONTENIDO);
			setState(287);
			match(ASSIGN);
			setState(288);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DestinoContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode DESTINO() { return getToken(GoDiskGrammar.DESTINO, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public DestinoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_destino; }
	}

	public final DestinoContext destino() throws RecognitionException {
		DestinoContext _localctx = new DestinoContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_destino);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			match(MINUS);
			setState(291);
			match(DESTINO);
			setState(292);
			match(ASSIGN);
			setState(293);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UsuarioContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode USUARIO() { return getToken(GoDiskGrammar.USUARIO, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode STRING_LIT() { return getToken(GoDiskGrammar.STRING_LIT, 0); }
		public TerminalNode UNQUOTED_TEXT() { return getToken(GoDiskGrammar.UNQUOTED_TEXT, 0); }
		public UsuarioContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_usuario; }
	}

	public final UsuarioContext usuario() throws RecognitionException {
		UsuarioContext _localctx = new UsuarioContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_usuario);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			match(MINUS);
			setState(296);
			match(USUARIO);
			setState(297);
			match(ASSIGN);
			setState(298);
			_la = _input.LA(1);
			if ( !(_la==STRING_LIT || _la==UNQUOTED_TEXT) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class UgoContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode UGO() { return getToken(GoDiskGrammar.UGO, 0); }
		public TerminalNode ASSIGN() { return getToken(GoDiskGrammar.ASSIGN, 0); }
		public TerminalNode ID() { return getToken(GoDiskGrammar.ID, 0); }
		public UgoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ugo; }
	}

	public final UgoContext ugo() throws RecognitionException {
		UgoContext _localctx = new UgoContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_ugo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(300);
			match(MINUS);
			setState(301);
			match(UGO);
			setState(302);
			match(ASSIGN);
			setState(303);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode P() { return getToken(GoDiskGrammar.P, 0); }
		public PContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_p; }
	}

	public final PContext p() throws RecognitionException {
		PContext _localctx = new PContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_p);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(305);
			match(MINUS);
			setState(306);
			match(P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RContext extends ParserRuleContext {
		public TerminalNode MINUS() { return getToken(GoDiskGrammar.MINUS, 0); }
		public TerminalNode R() { return getToken(GoDiskGrammar.R, 0); }
		public RContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_r; }
	}

	public final RContext r() throws RecognitionException {
		RContext _localctx = new RContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_r);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(308);
			match(MINUS);
			setState(309);
			match(R);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkdisk_paramsContext extends ParserRuleContext {
		public List<Mkdisk_paramContext> mkdisk_param() {
			return getRuleContexts(Mkdisk_paramContext.class);
		}
		public Mkdisk_paramContext mkdisk_param(int i) {
			return getRuleContext(Mkdisk_paramContext.class,i);
		}
		public Mkdisk_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdisk_params; }
	}

	public final Mkdisk_paramsContext mkdisk_params() throws RecognitionException {
		Mkdisk_paramsContext _localctx = new Mkdisk_paramsContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_mkdisk_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(312); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(311);
				mkdisk_param();
				}
				}
				setState(314); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Fdisk_paramsContext extends ParserRuleContext {
		public List<Fdisk_paramContext> fdisk_param() {
			return getRuleContexts(Fdisk_paramContext.class);
		}
		public Fdisk_paramContext fdisk_param(int i) {
			return getRuleContext(Fdisk_paramContext.class,i);
		}
		public Fdisk_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fdisk_params; }
	}

	public final Fdisk_paramsContext fdisk_params() throws RecognitionException {
		Fdisk_paramsContext _localctx = new Fdisk_paramsContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_fdisk_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(316);
				fdisk_param();
				}
				}
				setState(319); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mount_paramsContext extends ParserRuleContext {
		public List<Mount_paramContext> mount_param() {
			return getRuleContexts(Mount_paramContext.class);
		}
		public Mount_paramContext mount_param(int i) {
			return getRuleContext(Mount_paramContext.class,i);
		}
		public Mount_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mount_params; }
	}

	public final Mount_paramsContext mount_params() throws RecognitionException {
		Mount_paramsContext _localctx = new Mount_paramsContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_mount_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(322); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(321);
				mount_param();
				}
				}
				setState(324); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkfs_paramsContext extends ParserRuleContext {
		public List<Mkfs_paramContext> mkfs_param() {
			return getRuleContexts(Mkfs_paramContext.class);
		}
		public Mkfs_paramContext mkfs_param(int i) {
			return getRuleContext(Mkfs_paramContext.class,i);
		}
		public Mkfs_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfs_params; }
	}

	public final Mkfs_paramsContext mkfs_params() throws RecognitionException {
		Mkfs_paramsContext _localctx = new Mkfs_paramsContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_mkfs_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(327); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(326);
				mkfs_param();
				}
				}
				setState(329); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Login_paramsContext extends ParserRuleContext {
		public List<Login_paramContext> login_param() {
			return getRuleContexts(Login_paramContext.class);
		}
		public Login_paramContext login_param(int i) {
			return getRuleContext(Login_paramContext.class,i);
		}
		public Login_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_login_params; }
	}

	public final Login_paramsContext login_params() throws RecognitionException {
		Login_paramsContext _localctx = new Login_paramsContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_login_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(332); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(331);
				login_param();
				}
				}
				setState(334); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Cat_paramsContext extends ParserRuleContext {
		public List<Cat_paramContext> cat_param() {
			return getRuleContexts(Cat_paramContext.class);
		}
		public Cat_paramContext cat_param(int i) {
			return getRuleContext(Cat_paramContext.class,i);
		}
		public Cat_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cat_params; }
	}

	public final Cat_paramsContext cat_params() throws RecognitionException {
		Cat_paramsContext _localctx = new Cat_paramsContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_cat_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(336);
				cat_param();
				}
				}
				setState(339); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkusr_paramsContext extends ParserRuleContext {
		public List<Mkusr_paramContext> mkusr_param() {
			return getRuleContexts(Mkusr_paramContext.class);
		}
		public Mkusr_paramContext mkusr_param(int i) {
			return getRuleContext(Mkusr_paramContext.class,i);
		}
		public Mkusr_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkusr_params; }
	}

	public final Mkusr_paramsContext mkusr_params() throws RecognitionException {
		Mkusr_paramsContext _localctx = new Mkusr_paramsContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_mkusr_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(342); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(341);
				mkusr_param();
				}
				}
				setState(344); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Chgrp_paramsContext extends ParserRuleContext {
		public List<Chgrp_paramContext> chgrp_param() {
			return getRuleContexts(Chgrp_paramContext.class);
		}
		public Chgrp_paramContext chgrp_param(int i) {
			return getRuleContext(Chgrp_paramContext.class,i);
		}
		public Chgrp_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chgrp_params; }
	}

	public final Chgrp_paramsContext chgrp_params() throws RecognitionException {
		Chgrp_paramsContext _localctx = new Chgrp_paramsContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_chgrp_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(347); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(346);
				chgrp_param();
				}
				}
				setState(349); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkfile_paramsContext extends ParserRuleContext {
		public List<Mkfile_paramContext> mkfile_param() {
			return getRuleContexts(Mkfile_paramContext.class);
		}
		public Mkfile_paramContext mkfile_param(int i) {
			return getRuleContext(Mkfile_paramContext.class,i);
		}
		public Mkfile_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfile_params; }
	}

	public final Mkfile_paramsContext mkfile_params() throws RecognitionException {
		Mkfile_paramsContext _localctx = new Mkfile_paramsContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_mkfile_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(352); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(351);
				mkfile_param();
				}
				}
				setState(354); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkdir_paramsContext extends ParserRuleContext {
		public List<Mkdir_paramContext> mkdir_param() {
			return getRuleContexts(Mkdir_paramContext.class);
		}
		public Mkdir_paramContext mkdir_param(int i) {
			return getRuleContext(Mkdir_paramContext.class,i);
		}
		public Mkdir_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdir_params; }
	}

	public final Mkdir_paramsContext mkdir_params() throws RecognitionException {
		Mkdir_paramsContext _localctx = new Mkdir_paramsContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_mkdir_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(357); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(356);
				mkdir_param();
				}
				}
				setState(359); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Edit_paramsContext extends ParserRuleContext {
		public List<Edit_paramContext> edit_param() {
			return getRuleContexts(Edit_paramContext.class);
		}
		public Edit_paramContext edit_param(int i) {
			return getRuleContext(Edit_paramContext.class,i);
		}
		public Edit_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_edit_params; }
	}

	public final Edit_paramsContext edit_params() throws RecognitionException {
		Edit_paramsContext _localctx = new Edit_paramsContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_edit_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(362); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(361);
				edit_param();
				}
				}
				setState(364); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rename_paramsContext extends ParserRuleContext {
		public List<Rename_paramContext> rename_param() {
			return getRuleContexts(Rename_paramContext.class);
		}
		public Rename_paramContext rename_param(int i) {
			return getRuleContext(Rename_paramContext.class,i);
		}
		public Rename_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rename_params; }
	}

	public final Rename_paramsContext rename_params() throws RecognitionException {
		Rename_paramsContext _localctx = new Rename_paramsContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_rename_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(367); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(366);
				rename_param();
				}
				}
				setState(369); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Copy_paramsContext extends ParserRuleContext {
		public List<Copy_paramContext> copy_param() {
			return getRuleContexts(Copy_paramContext.class);
		}
		public Copy_paramContext copy_param(int i) {
			return getRuleContext(Copy_paramContext.class,i);
		}
		public Copy_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_copy_params; }
	}

	public final Copy_paramsContext copy_params() throws RecognitionException {
		Copy_paramsContext _localctx = new Copy_paramsContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_copy_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(372); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(371);
				copy_param();
				}
				}
				setState(374); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Move_paramsContext extends ParserRuleContext {
		public List<Move_paramContext> move_param() {
			return getRuleContexts(Move_paramContext.class);
		}
		public Move_paramContext move_param(int i) {
			return getRuleContext(Move_paramContext.class,i);
		}
		public Move_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_move_params; }
	}

	public final Move_paramsContext move_params() throws RecognitionException {
		Move_paramsContext _localctx = new Move_paramsContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_move_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(377); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(376);
				move_param();
				}
				}
				setState(379); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Find_paramsContext extends ParserRuleContext {
		public List<Find_paramContext> find_param() {
			return getRuleContexts(Find_paramContext.class);
		}
		public Find_paramContext find_param(int i) {
			return getRuleContext(Find_paramContext.class,i);
		}
		public Find_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_find_params; }
	}

	public final Find_paramsContext find_params() throws RecognitionException {
		Find_paramsContext _localctx = new Find_paramsContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_find_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(382); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(381);
				find_param();
				}
				}
				setState(384); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Chown_paramsContext extends ParserRuleContext {
		public List<Chown_paramContext> chown_param() {
			return getRuleContexts(Chown_paramContext.class);
		}
		public Chown_paramContext chown_param(int i) {
			return getRuleContext(Chown_paramContext.class,i);
		}
		public Chown_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chown_params; }
	}

	public final Chown_paramsContext chown_params() throws RecognitionException {
		Chown_paramsContext _localctx = new Chown_paramsContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_chown_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(387); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(386);
				chown_param();
				}
				}
				setState(389); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Chmod_paramsContext extends ParserRuleContext {
		public List<Chmod_paramContext> chmod_param() {
			return getRuleContexts(Chmod_paramContext.class);
		}
		public Chmod_paramContext chmod_param(int i) {
			return getRuleContext(Chmod_paramContext.class,i);
		}
		public Chmod_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chmod_params; }
	}

	public final Chmod_paramsContext chmod_params() throws RecognitionException {
		Chmod_paramsContext _localctx = new Chmod_paramsContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_chmod_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(392); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(391);
				chmod_param();
				}
				}
				setState(394); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rep_paramsContext extends ParserRuleContext {
		public List<Rep_paramContext> rep_param() {
			return getRuleContexts(Rep_paramContext.class);
		}
		public Rep_paramContext rep_param(int i) {
			return getRuleContext(Rep_paramContext.class,i);
		}
		public Rep_paramsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rep_params; }
	}

	public final Rep_paramsContext rep_params() throws RecognitionException {
		Rep_paramsContext _localctx = new Rep_paramsContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_rep_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(397); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(396);
				rep_param();
				}
				}
				setState(399); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==MINUS );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkdisk_paramContext extends ParserRuleContext {
		public SizeContext size() {
			return getRuleContext(SizeContext.class,0);
		}
		public FitContext fit() {
			return getRuleContext(FitContext.class,0);
		}
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public Mkdisk_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdisk_param; }
	}

	public final Mkdisk_paramContext mkdisk_param() throws RecognitionException {
		Mkdisk_paramContext _localctx = new Mkdisk_paramContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_mkdisk_param);
		try {
			setState(405);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(401);
				size();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(402);
				fit();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(403);
				unit();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(404);
				path();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rmdisk_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public Rmdisk_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rmdisk_param; }
	}

	public final Rmdisk_paramContext rmdisk_param() throws RecognitionException {
		Rmdisk_paramContext _localctx = new Rmdisk_paramContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_rmdisk_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(407);
			path();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Fdisk_paramContext extends ParserRuleContext {
		public SizeContext size() {
			return getRuleContext(SizeContext.class,0);
		}
		public FitContext fit() {
			return getRuleContext(FitContext.class,0);
		}
		public UnitContext unit() {
			return getRuleContext(UnitContext.class,0);
		}
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public DeleteContext delete() {
			return getRuleContext(DeleteContext.class,0);
		}
		public AddContext add() {
			return getRuleContext(AddContext.class,0);
		}
		public Fdisk_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fdisk_param; }
	}

	public final Fdisk_paramContext fdisk_param() throws RecognitionException {
		Fdisk_paramContext _localctx = new Fdisk_paramContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_fdisk_param);
		try {
			setState(417);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(409);
				size();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(410);
				fit();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(411);
				unit();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(412);
				path();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(413);
				type();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(414);
				name();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(415);
				delete();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(416);
				add();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mount_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public Mount_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mount_param; }
	}

	public final Mount_paramContext mount_param() throws RecognitionException {
		Mount_paramContext _localctx = new Mount_paramContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_mount_param);
		try {
			setState(421);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(419);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(420);
				name();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Unmount_paramContext extends ParserRuleContext {
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public Unmount_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unmount_param; }
	}

	public final Unmount_paramContext unmount_param() throws RecognitionException {
		Unmount_paramContext _localctx = new Unmount_paramContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_unmount_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(423);
			id_text();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkfs_paramContext extends ParserRuleContext {
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public FsContext fs() {
			return getRuleContext(FsContext.class,0);
		}
		public Mkfs_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfs_param; }
	}

	public final Mkfs_paramContext mkfs_param() throws RecognitionException {
		Mkfs_paramContext _localctx = new Mkfs_paramContext(_ctx, getState());
		enterRule(_localctx, 94, RULE_mkfs_param);
		try {
			setState(428);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(425);
				id_text();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(426);
				type();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(427);
				fs();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Cat_paramContext extends ParserRuleContext {
		public FilenContext filen() {
			return getRuleContext(FilenContext.class,0);
		}
		public Cat_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cat_param; }
	}

	public final Cat_paramContext cat_param() throws RecognitionException {
		Cat_paramContext _localctx = new Cat_paramContext(_ctx, getState());
		enterRule(_localctx, 96, RULE_cat_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(430);
			filen();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Login_paramContext extends ParserRuleContext {
		public UserContext user() {
			return getRuleContext(UserContext.class,0);
		}
		public PassContext pass() {
			return getRuleContext(PassContext.class,0);
		}
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public Login_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_login_param; }
	}

	public final Login_paramContext login_param() throws RecognitionException {
		Login_paramContext _localctx = new Login_paramContext(_ctx, getState());
		enterRule(_localctx, 98, RULE_login_param);
		try {
			setState(435);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(432);
				user();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(433);
				pass();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(434);
				id_text();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkgrp_paramContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public Mkgrp_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkgrp_param; }
	}

	public final Mkgrp_paramContext mkgrp_param() throws RecognitionException {
		Mkgrp_paramContext _localctx = new Mkgrp_paramContext(_ctx, getState());
		enterRule(_localctx, 100, RULE_mkgrp_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(437);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rmgrp_paramContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public Rmgrp_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rmgrp_param; }
	}

	public final Rmgrp_paramContext rmgrp_param() throws RecognitionException {
		Rmgrp_paramContext _localctx = new Rmgrp_paramContext(_ctx, getState());
		enterRule(_localctx, 102, RULE_rmgrp_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(439);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkusr_paramContext extends ParserRuleContext {
		public UserContext user() {
			return getRuleContext(UserContext.class,0);
		}
		public PassContext pass() {
			return getRuleContext(PassContext.class,0);
		}
		public GrpContext grp() {
			return getRuleContext(GrpContext.class,0);
		}
		public Mkusr_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkusr_param; }
	}

	public final Mkusr_paramContext mkusr_param() throws RecognitionException {
		Mkusr_paramContext _localctx = new Mkusr_paramContext(_ctx, getState());
		enterRule(_localctx, 104, RULE_mkusr_param);
		try {
			setState(444);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(441);
				user();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(442);
				pass();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(443);
				grp();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rmusr_paramContext extends ParserRuleContext {
		public UserContext user() {
			return getRuleContext(UserContext.class,0);
		}
		public Rmusr_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rmusr_param; }
	}

	public final Rmusr_paramContext rmusr_param() throws RecognitionException {
		Rmusr_paramContext _localctx = new Rmusr_paramContext(_ctx, getState());
		enterRule(_localctx, 106, RULE_rmusr_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(446);
			user();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Chgrp_paramContext extends ParserRuleContext {
		public UserContext user() {
			return getRuleContext(UserContext.class,0);
		}
		public GrpContext grp() {
			return getRuleContext(GrpContext.class,0);
		}
		public Chgrp_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chgrp_param; }
	}

	public final Chgrp_paramContext chgrp_param() throws RecognitionException {
		Chgrp_paramContext _localctx = new Chgrp_paramContext(_ctx, getState());
		enterRule(_localctx, 108, RULE_chgrp_param);
		try {
			setState(450);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(448);
				user();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(449);
				grp();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkfile_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public RContext r() {
			return getRuleContext(RContext.class,0);
		}
		public SizeContext size() {
			return getRuleContext(SizeContext.class,0);
		}
		public ContContext cont() {
			return getRuleContext(ContContext.class,0);
		}
		public Mkfile_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfile_param; }
	}

	public final Mkfile_paramContext mkfile_param() throws RecognitionException {
		Mkfile_paramContext _localctx = new Mkfile_paramContext(_ctx, getState());
		enterRule(_localctx, 110, RULE_mkfile_param);
		try {
			setState(456);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(452);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(453);
				r();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(454);
				size();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(455);
				cont();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Mkdir_paramContext extends ParserRuleContext {
		public PContext p() {
			return getRuleContext(PContext.class,0);
		}
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public Mkdir_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkdir_param; }
	}

	public final Mkdir_paramContext mkdir_param() throws RecognitionException {
		Mkdir_paramContext _localctx = new Mkdir_paramContext(_ctx, getState());
		enterRule(_localctx, 112, RULE_mkdir_param);
		try {
			setState(460);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(458);
				p();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(459);
				path();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Remove_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public Remove_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_remove_param; }
	}

	public final Remove_paramContext remove_param() throws RecognitionException {
		Remove_paramContext _localctx = new Remove_paramContext(_ctx, getState());
		enterRule(_localctx, 114, RULE_remove_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(462);
			path();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Edit_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public ContenidoContext contenido() {
			return getRuleContext(ContenidoContext.class,0);
		}
		public Edit_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_edit_param; }
	}

	public final Edit_paramContext edit_param() throws RecognitionException {
		Edit_paramContext _localctx = new Edit_paramContext(_ctx, getState());
		enterRule(_localctx, 116, RULE_edit_param);
		try {
			setState(466);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(464);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(465);
				contenido();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rename_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public Rename_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rename_param; }
	}

	public final Rename_paramContext rename_param() throws RecognitionException {
		Rename_paramContext _localctx = new Rename_paramContext(_ctx, getState());
		enterRule(_localctx, 118, RULE_rename_param);
		try {
			setState(470);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(468);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(469);
				name();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Copy_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public DestinoContext destino() {
			return getRuleContext(DestinoContext.class,0);
		}
		public Copy_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_copy_param; }
	}

	public final Copy_paramContext copy_param() throws RecognitionException {
		Copy_paramContext _localctx = new Copy_paramContext(_ctx, getState());
		enterRule(_localctx, 120, RULE_copy_param);
		try {
			setState(474);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(472);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(473);
				destino();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Move_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public DestinoContext destino() {
			return getRuleContext(DestinoContext.class,0);
		}
		public Move_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_move_param; }
	}

	public final Move_paramContext move_param() throws RecognitionException {
		Move_paramContext _localctx = new Move_paramContext(_ctx, getState());
		enterRule(_localctx, 122, RULE_move_param);
		try {
			setState(478);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(476);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(477);
				destino();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Find_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public Find_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_find_param; }
	}

	public final Find_paramContext find_param() throws RecognitionException {
		Find_paramContext _localctx = new Find_paramContext(_ctx, getState());
		enterRule(_localctx, 124, RULE_find_param);
		try {
			setState(482);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(480);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(481);
				name();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Chown_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public UsuarioContext usuario() {
			return getRuleContext(UsuarioContext.class,0);
		}
		public RContext r() {
			return getRuleContext(RContext.class,0);
		}
		public Chown_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chown_param; }
	}

	public final Chown_paramContext chown_param() throws RecognitionException {
		Chown_paramContext _localctx = new Chown_paramContext(_ctx, getState());
		enterRule(_localctx, 126, RULE_chown_param);
		try {
			setState(487);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(484);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(485);
				usuario();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(486);
				r();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Chmod_paramContext extends ParserRuleContext {
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public UgoContext ugo() {
			return getRuleContext(UgoContext.class,0);
		}
		public RContext r() {
			return getRuleContext(RContext.class,0);
		}
		public Chmod_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_chmod_param; }
	}

	public final Chmod_paramContext chmod_param() throws RecognitionException {
		Chmod_paramContext _localctx = new Chmod_paramContext(_ctx, getState());
		enterRule(_localctx, 128, RULE_chmod_param);
		try {
			setState(492);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,35,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(489);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(490);
				ugo();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(491);
				r();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Recovery_paramContext extends ParserRuleContext {
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public Recovery_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_recovery_param; }
	}

	public final Recovery_paramContext recovery_param() throws RecognitionException {
		Recovery_paramContext _localctx = new Recovery_paramContext(_ctx, getState());
		enterRule(_localctx, 130, RULE_recovery_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(494);
			id_text();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Loss_paramContext extends ParserRuleContext {
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public Loss_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loss_param; }
	}

	public final Loss_paramContext loss_param() throws RecognitionException {
		Loss_paramContext _localctx = new Loss_paramContext(_ctx, getState());
		enterRule(_localctx, 132, RULE_loss_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(496);
			id_text();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Journaling_paramContext extends ParserRuleContext {
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public Journaling_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_journaling_param; }
	}

	public final Journaling_paramContext journaling_param() throws RecognitionException {
		Journaling_paramContext _localctx = new Journaling_paramContext(_ctx, getState());
		enterRule(_localctx, 134, RULE_journaling_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(498);
			id_text();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Rep_paramContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public PathContext path() {
			return getRuleContext(PathContext.class,0);
		}
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public Path_file_lsContext path_file_ls() {
			return getRuleContext(Path_file_lsContext.class,0);
		}
		public Rep_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rep_param; }
	}

	public final Rep_paramContext rep_param() throws RecognitionException {
		Rep_paramContext _localctx = new Rep_paramContext(_ctx, getState());
		enterRule(_localctx, 136, RULE_rep_param);
		try {
			setState(504);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,36,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(500);
				name();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(501);
				path();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(502);
				id_text();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(503);
				path_file_ls();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001>\u01fb\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0002,\u0007,\u0002"+
		"-\u0007-\u0002.\u0007.\u0002/\u0007/\u00020\u00070\u00021\u00071\u0002"+
		"2\u00072\u00023\u00073\u00024\u00074\u00025\u00075\u00026\u00076\u0002"+
		"7\u00077\u00028\u00078\u00029\u00079\u0002:\u0007:\u0002;\u0007;\u0002"+
		"<\u0007<\u0002=\u0007=\u0002>\u0007>\u0002?\u0007?\u0002@\u0007@\u0002"+
		"A\u0007A\u0002B\u0007B\u0002C\u0007C\u0002D\u0007D\u0001\u0000\u0005\u0000"+
		"\u008c\b\u0000\n\u0000\f\u0000\u008f\t\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u0001\u00cb\b\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0018\u0004\u0018\u0139\b\u0018\u000b\u0018\f\u0018\u013a\u0001\u0019"+
		"\u0004\u0019\u013e\b\u0019\u000b\u0019\f\u0019\u013f\u0001\u001a\u0004"+
		"\u001a\u0143\b\u001a\u000b\u001a\f\u001a\u0144\u0001\u001b\u0004\u001b"+
		"\u0148\b\u001b\u000b\u001b\f\u001b\u0149\u0001\u001c\u0004\u001c\u014d"+
		"\b\u001c\u000b\u001c\f\u001c\u014e\u0001\u001d\u0004\u001d\u0152\b\u001d"+
		"\u000b\u001d\f\u001d\u0153\u0001\u001e\u0004\u001e\u0157\b\u001e\u000b"+
		"\u001e\f\u001e\u0158\u0001\u001f\u0004\u001f\u015c\b\u001f\u000b\u001f"+
		"\f\u001f\u015d\u0001 \u0004 \u0161\b \u000b \f \u0162\u0001!\u0004!\u0166"+
		"\b!\u000b!\f!\u0167\u0001\"\u0004\"\u016b\b\"\u000b\"\f\"\u016c\u0001"+
		"#\u0004#\u0170\b#\u000b#\f#\u0171\u0001$\u0004$\u0175\b$\u000b$\f$\u0176"+
		"\u0001%\u0004%\u017a\b%\u000b%\f%\u017b\u0001&\u0004&\u017f\b&\u000b&"+
		"\f&\u0180\u0001\'\u0004\'\u0184\b\'\u000b\'\f\'\u0185\u0001(\u0004(\u0189"+
		"\b(\u000b(\f(\u018a\u0001)\u0004)\u018e\b)\u000b)\f)\u018f\u0001*\u0001"+
		"*\u0001*\u0001*\u0003*\u0196\b*\u0001+\u0001+\u0001,\u0001,\u0001,\u0001"+
		",\u0001,\u0001,\u0001,\u0001,\u0003,\u01a2\b,\u0001-\u0001-\u0003-\u01a6"+
		"\b-\u0001.\u0001.\u0001/\u0001/\u0001/\u0003/\u01ad\b/\u00010\u00010\u0001"+
		"1\u00011\u00011\u00031\u01b4\b1\u00012\u00012\u00013\u00013\u00014\u0001"+
		"4\u00014\u00034\u01bd\b4\u00015\u00015\u00016\u00016\u00036\u01c3\b6\u0001"+
		"7\u00017\u00017\u00017\u00037\u01c9\b7\u00018\u00018\u00038\u01cd\b8\u0001"+
		"9\u00019\u0001:\u0001:\u0003:\u01d3\b:\u0001;\u0001;\u0003;\u01d7\b;\u0001"+
		"<\u0001<\u0003<\u01db\b<\u0001=\u0001=\u0003=\u01df\b=\u0001>\u0001>\u0003"+
		">\u01e3\b>\u0001?\u0001?\u0001?\u0003?\u01e8\b?\u0001@\u0001@\u0001@\u0003"+
		"@\u01ed\b@\u0001A\u0001A\u0001B\u0001B\u0001C\u0001C\u0001D\u0001D\u0001"+
		"D\u0001D\u0003D\u01f9\bD\u0001D\u0000\u0000E\u0000\u0002\u0004\u0006\b"+
		"\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02"+
		"468:<>@BDFHJLNPRTVXZ\\^`bdfhjlnprtvxz|~\u0080\u0082\u0084\u0086\u0088"+
		"\u0000\u0001\u0002\u0000\u0003\u0003;;\u0206\u0000\u008d\u0001\u0000\u0000"+
		"\u0000\u0002\u00ca\u0001\u0000\u0000\u0000\u0004\u00cc\u0001\u0000\u0000"+
		"\u0000\u0006\u00d1\u0001\u0000\u0000\u0000\b\u00d6\u0001\u0000\u0000\u0000"+
		"\n\u00db\u0001\u0000\u0000\u0000\f\u00e0\u0001\u0000\u0000\u0000\u000e"+
		"\u00e5\u0001\u0000\u0000\u0000\u0010\u00ea\u0001\u0000\u0000\u0000\u0012"+
		"\u00ef\u0001\u0000\u0000\u0000\u0014\u00f5\u0001\u0000\u0000\u0000\u0016"+
		"\u00fa\u0001\u0000\u0000\u0000\u0018\u00ff\u0001\u0000\u0000\u0000\u001a"+
		"\u0104\u0001\u0000\u0000\u0000\u001c\u0109\u0001\u0000\u0000\u0000\u001e"+
		"\u010e\u0001\u0000\u0000\u0000 \u0113\u0001\u0000\u0000\u0000\"\u0118"+
		"\u0001\u0000\u0000\u0000$\u011d\u0001\u0000\u0000\u0000&\u0122\u0001\u0000"+
		"\u0000\u0000(\u0127\u0001\u0000\u0000\u0000*\u012c\u0001\u0000\u0000\u0000"+
		",\u0131\u0001\u0000\u0000\u0000.\u0134\u0001\u0000\u0000\u00000\u0138"+
		"\u0001\u0000\u0000\u00002\u013d\u0001\u0000\u0000\u00004\u0142\u0001\u0000"+
		"\u0000\u00006\u0147\u0001\u0000\u0000\u00008\u014c\u0001\u0000\u0000\u0000"+
		":\u0151\u0001\u0000\u0000\u0000<\u0156\u0001\u0000\u0000\u0000>\u015b"+
		"\u0001\u0000\u0000\u0000@\u0160\u0001\u0000\u0000\u0000B\u0165\u0001\u0000"+
		"\u0000\u0000D\u016a\u0001\u0000\u0000\u0000F\u016f\u0001\u0000\u0000\u0000"+
		"H\u0174\u0001\u0000\u0000\u0000J\u0179\u0001\u0000\u0000\u0000L\u017e"+
		"\u0001\u0000\u0000\u0000N\u0183\u0001\u0000\u0000\u0000P\u0188\u0001\u0000"+
		"\u0000\u0000R\u018d\u0001\u0000\u0000\u0000T\u0195\u0001\u0000\u0000\u0000"+
		"V\u0197\u0001\u0000\u0000\u0000X\u01a1\u0001\u0000\u0000\u0000Z\u01a5"+
		"\u0001\u0000\u0000\u0000\\\u01a7\u0001\u0000\u0000\u0000^\u01ac\u0001"+
		"\u0000\u0000\u0000`\u01ae\u0001\u0000\u0000\u0000b\u01b3\u0001\u0000\u0000"+
		"\u0000d\u01b5\u0001\u0000\u0000\u0000f\u01b7\u0001\u0000\u0000\u0000h"+
		"\u01bc\u0001\u0000\u0000\u0000j\u01be\u0001\u0000\u0000\u0000l\u01c2\u0001"+
		"\u0000\u0000\u0000n\u01c8\u0001\u0000\u0000\u0000p\u01cc\u0001\u0000\u0000"+
		"\u0000r\u01ce\u0001\u0000\u0000\u0000t\u01d2\u0001\u0000\u0000\u0000v"+
		"\u01d6\u0001\u0000\u0000\u0000x\u01da\u0001\u0000\u0000\u0000z\u01de\u0001"+
		"\u0000\u0000\u0000|\u01e2\u0001\u0000\u0000\u0000~\u01e7\u0001\u0000\u0000"+
		"\u0000\u0080\u01ec\u0001\u0000\u0000\u0000\u0082\u01ee\u0001\u0000\u0000"+
		"\u0000\u0084\u01f0\u0001\u0000\u0000\u0000\u0086\u01f2\u0001\u0000\u0000"+
		"\u0000\u0088\u01f8\u0001\u0000\u0000\u0000\u008a\u008c\u0003\u0002\u0001"+
		"\u0000\u008b\u008a\u0001\u0000\u0000\u0000\u008c\u008f\u0001\u0000\u0000"+
		"\u0000\u008d\u008b\u0001\u0000\u0000\u0000\u008d\u008e\u0001\u0000\u0000"+
		"\u0000\u008e\u0090\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000"+
		"\u0000\u0090\u0091\u0005\u0000\u0000\u0001\u0091\u0001\u0001\u0000\u0000"+
		"\u0000\u0092\u0093\u0005\u0005\u0000\u0000\u0093\u00cb\u00030\u0018\u0000"+
		"\u0094\u0095\u0005\u0006\u0000\u0000\u0095\u00cb\u0003V+\u0000\u0096\u0097"+
		"\u0005\u0007\u0000\u0000\u0097\u00cb\u00032\u0019\u0000\u0098\u0099\u0005"+
		"\b\u0000\u0000\u0099\u00cb\u00034\u001a\u0000\u009a\u00cb\u0005\t\u0000"+
		"\u0000\u009b\u009c\u0005\n\u0000\u0000\u009c\u00cb\u0003\\.\u0000\u009d"+
		"\u009e\u0005\u000b\u0000\u0000\u009e\u00cb\u00036\u001b\u0000\u009f\u00a0"+
		"\u0005\f\u0000\u0000\u00a0\u00cb\u0003:\u001d\u0000\u00a1\u00a2\u0005"+
		"\r\u0000\u0000\u00a2\u00cb\u00038\u001c\u0000\u00a3\u00cb\u0005\u000e"+
		"\u0000\u0000\u00a4\u00a5\u0005\u000f\u0000\u0000\u00a5\u00cb\u0003d2\u0000"+
		"\u00a6\u00a7\u0005\u0010\u0000\u0000\u00a7\u00cb\u0003f3\u0000\u00a8\u00a9"+
		"\u0005\u0011\u0000\u0000\u00a9\u00cb\u0003<\u001e\u0000\u00aa\u00ab\u0005"+
		"\u0012\u0000\u0000\u00ab\u00cb\u0003j5\u0000\u00ac\u00ad\u0005\u0013\u0000"+
		"\u0000\u00ad\u00cb\u0003>\u001f\u0000\u00ae\u00af\u0005\u0014\u0000\u0000"+
		"\u00af\u00cb\u0003@ \u0000\u00b0\u00b1\u0005\u0015\u0000\u0000\u00b1\u00cb"+
		"\u0003B!\u0000\u00b2\u00b3\u0005\u0016\u0000\u0000\u00b3\u00cb\u0003r"+
		"9\u0000\u00b4\u00b5\u0005\u0017\u0000\u0000\u00b5\u00cb\u0003D\"\u0000"+
		"\u00b6\u00b7\u0005\u0018\u0000\u0000\u00b7\u00cb\u0003F#\u0000\u00b8\u00b9"+
		"\u0005\u0019\u0000\u0000\u00b9\u00cb\u0003H$\u0000\u00ba\u00bb\u0005\u001a"+
		"\u0000\u0000\u00bb\u00cb\u0003J%\u0000\u00bc\u00bd\u0005\u001b\u0000\u0000"+
		"\u00bd\u00cb\u0003L&\u0000\u00be\u00bf\u0005\u001c\u0000\u0000\u00bf\u00cb"+
		"\u0003N\'\u0000\u00c0\u00c1\u0005\u001d\u0000\u0000\u00c1\u00cb\u0003"+
		"P(\u0000\u00c2\u00c3\u0005\u001e\u0000\u0000\u00c3\u00cb\u0003\u0082A"+
		"\u0000\u00c4\u00c5\u0005\u001f\u0000\u0000\u00c5\u00cb\u0003\u0084B\u0000"+
		"\u00c6\u00c7\u0005 \u0000\u0000\u00c7\u00cb\u0003\u0086C\u0000\u00c8\u00c9"+
		"\u0005!\u0000\u0000\u00c9\u00cb\u0003R)\u0000\u00ca\u0092\u0001\u0000"+
		"\u0000\u0000\u00ca\u0094\u0001\u0000\u0000\u0000\u00ca\u0096\u0001\u0000"+
		"\u0000\u0000\u00ca\u0098\u0001\u0000\u0000\u0000\u00ca\u009a\u0001\u0000"+
		"\u0000\u0000\u00ca\u009b\u0001\u0000\u0000\u0000\u00ca\u009d\u0001\u0000"+
		"\u0000\u0000\u00ca\u009f\u0001\u0000\u0000\u0000\u00ca\u00a1\u0001\u0000"+
		"\u0000\u0000\u00ca\u00a3\u0001\u0000\u0000\u0000\u00ca\u00a4\u0001\u0000"+
		"\u0000\u0000\u00ca\u00a6\u0001\u0000\u0000\u0000\u00ca\u00a8\u0001\u0000"+
		"\u0000\u0000\u00ca\u00aa\u0001\u0000\u0000\u0000\u00ca\u00ac\u0001\u0000"+
		"\u0000\u0000\u00ca\u00ae\u0001\u0000\u0000\u0000\u00ca\u00b0\u0001\u0000"+
		"\u0000\u0000\u00ca\u00b2\u0001\u0000\u0000\u0000\u00ca\u00b4\u0001\u0000"+
		"\u0000\u0000\u00ca\u00b6\u0001\u0000\u0000\u0000\u00ca\u00b8\u0001\u0000"+
		"\u0000\u0000\u00ca\u00ba\u0001\u0000\u0000\u0000\u00ca\u00bc\u0001\u0000"+
		"\u0000\u0000\u00ca\u00be\u0001\u0000\u0000\u0000\u00ca\u00c0\u0001\u0000"+
		"\u0000\u0000\u00ca\u00c2\u0001\u0000\u0000\u0000\u00ca\u00c4\u0001\u0000"+
		"\u0000\u0000\u00ca\u00c6\u0001\u0000\u0000\u0000\u00ca\u00c8\u0001\u0000"+
		"\u0000\u0000\u00cb\u0003\u0001\u0000\u0000\u0000\u00cc\u00cd\u00059\u0000"+
		"\u0000\u00cd\u00ce\u0005#\u0000\u0000\u00ce\u00cf\u00058\u0000\u0000\u00cf"+
		"\u00d0\u0005\u0001\u0000\u0000\u00d0\u0005\u0001\u0000\u0000\u0000\u00d1"+
		"\u00d2\u00059\u0000\u0000\u00d2\u00d3\u0005$\u0000\u0000\u00d3\u00d4\u0005"+
		"8\u0000\u0000\u00d4\u00d5\u0005:\u0000\u0000\u00d5\u0007\u0001\u0000\u0000"+
		"\u0000\u00d6\u00d7\u00059\u0000\u0000\u00d7\u00d8\u0005%\u0000\u0000\u00d8"+
		"\u00d9\u00058\u0000\u0000\u00d9\u00da\u0005:\u0000\u0000\u00da\t\u0001"+
		"\u0000\u0000\u0000\u00db\u00dc\u00059\u0000\u0000\u00dc\u00dd\u0005\'"+
		"\u0000\u0000\u00dd\u00de\u00058\u0000\u0000\u00de\u00df\u0005:\u0000\u0000"+
		"\u00df\u000b\u0001\u0000\u0000\u0000\u00e0\u00e1\u00059\u0000\u0000\u00e1"+
		"\u00e2\u0005\"\u0000\u0000\u00e2\u00e3\u00058\u0000\u0000\u00e3\u00e4"+
		"\u0005:\u0000\u0000\u00e4\r\u0001\u0000\u0000\u0000\u00e5\u00e6\u0005"+
		"9\u0000\u0000\u00e6\u00e7\u0005&\u0000\u0000\u00e7\u00e8\u00058\u0000"+
		"\u0000\u00e8\u00e9\u0007\u0000\u0000\u0000\u00e9\u000f\u0001\u0000\u0000"+
		"\u0000\u00ea\u00eb\u00059\u0000\u0000\u00eb\u00ec\u0005(\u0000\u0000\u00ec"+
		"\u00ed\u00058\u0000\u0000\u00ed\u00ee\u0007\u0000\u0000\u0000\u00ee\u0011"+
		"\u0001\u0000\u0000\u0000\u00ef\u00f0\u00059\u0000\u0000\u00f0\u00f1\u0005"+
		")\u0000\u0000\u00f1\u00f2\u0005\u0001\u0000\u0000\u00f2\u00f3\u00058\u0000"+
		"\u0000\u00f3\u00f4\u0007\u0000\u0000\u0000\u00f4\u0013\u0001\u0000\u0000"+
		"\u0000\u00f5\u00f6\u00059\u0000\u0000\u00f6\u00f7\u0005*\u0000\u0000\u00f7"+
		"\u00f8\u00058\u0000\u0000\u00f8\u00f9\u0007\u0000\u0000\u0000\u00f9\u0015"+
		"\u0001\u0000\u0000\u0000\u00fa\u00fb\u00059\u0000\u0000\u00fb\u00fc\u0005"+
		"+\u0000\u0000\u00fc\u00fd\u00058\u0000\u0000\u00fd\u00fe\u0007\u0000\u0000"+
		"\u0000\u00fe\u0017\u0001\u0000\u0000\u0000\u00ff\u0100\u00059\u0000\u0000"+
		"\u0100\u0101\u0005,\u0000\u0000\u0101\u0102\u00058\u0000\u0000\u0102\u0103"+
		"\u0007\u0000\u0000\u0000\u0103\u0019\u0001\u0000\u0000\u0000\u0104\u0105"+
		"\u00059\u0000\u0000\u0105\u0106\u0005-\u0000\u0000\u0106\u0107\u00058"+
		"\u0000\u0000\u0107\u0108\u0007\u0000\u0000\u0000\u0108\u001b\u0001\u0000"+
		"\u0000\u0000\u0109\u010a\u00059\u0000\u0000\u010a\u010b\u0005.\u0000\u0000"+
		"\u010b\u010c\u00058\u0000\u0000\u010c\u010d\u0007\u0000\u0000\u0000\u010d"+
		"\u001d\u0001\u0000\u0000\u0000\u010e\u010f\u00059\u0000\u0000\u010f\u0110"+
		"\u0005/\u0000\u0000\u0110\u0111\u00058\u0000\u0000\u0111\u0112\u0005:"+
		"\u0000\u0000\u0112\u001f\u0001\u0000\u0000\u0000\u0113\u0114\u00059\u0000"+
		"\u0000\u0114\u0115\u00050\u0000\u0000\u0115\u0116\u00058\u0000\u0000\u0116"+
		"\u0117\u0005\u0001\u0000\u0000\u0117!\u0001\u0000\u0000\u0000\u0118\u0119"+
		"\u00059\u0000\u0000\u0119\u011a\u00051\u0000\u0000\u011a\u011b\u00058"+
		"\u0000\u0000\u011b\u011c\u0005:\u0000\u0000\u011c#\u0001\u0000\u0000\u0000"+
		"\u011d\u011e\u00059\u0000\u0000\u011e\u011f\u00052\u0000\u0000\u011f\u0120"+
		"\u00058\u0000\u0000\u0120\u0121\u0007\u0000\u0000\u0000\u0121%\u0001\u0000"+
		"\u0000\u0000\u0122\u0123\u00059\u0000\u0000\u0123\u0124\u00053\u0000\u0000"+
		"\u0124\u0125\u00058\u0000\u0000\u0125\u0126\u0007\u0000\u0000\u0000\u0126"+
		"\'\u0001\u0000\u0000\u0000\u0127\u0128\u00059\u0000\u0000\u0128\u0129"+
		"\u00054\u0000\u0000\u0129\u012a\u00058\u0000\u0000\u012a\u012b\u0007\u0000"+
		"\u0000\u0000\u012b)\u0001\u0000\u0000\u0000\u012c\u012d\u00059\u0000\u0000"+
		"\u012d\u012e\u00055\u0000\u0000\u012e\u012f\u00058\u0000\u0000\u012f\u0130"+
		"\u0005:\u0000\u0000\u0130+\u0001\u0000\u0000\u0000\u0131\u0132\u00059"+
		"\u0000\u0000\u0132\u0133\u00057\u0000\u0000\u0133-\u0001\u0000\u0000\u0000"+
		"\u0134\u0135\u00059\u0000\u0000\u0135\u0136\u00056\u0000\u0000\u0136/"+
		"\u0001\u0000\u0000\u0000\u0137\u0139\u0003T*\u0000\u0138\u0137\u0001\u0000"+
		"\u0000\u0000\u0139\u013a\u0001\u0000\u0000\u0000\u013a\u0138\u0001\u0000"+
		"\u0000\u0000\u013a\u013b\u0001\u0000\u0000\u0000\u013b1\u0001\u0000\u0000"+
		"\u0000\u013c\u013e\u0003X,\u0000\u013d\u013c\u0001\u0000\u0000\u0000\u013e"+
		"\u013f\u0001\u0000\u0000\u0000\u013f\u013d\u0001\u0000\u0000\u0000\u013f"+
		"\u0140\u0001\u0000\u0000\u0000\u01403\u0001\u0000\u0000\u0000\u0141\u0143"+
		"\u0003Z-\u0000\u0142\u0141\u0001\u0000\u0000\u0000\u0143\u0144\u0001\u0000"+
		"\u0000\u0000\u0144\u0142\u0001\u0000\u0000\u0000\u0144\u0145\u0001\u0000"+
		"\u0000\u0000\u01455\u0001\u0000\u0000\u0000\u0146\u0148\u0003^/\u0000"+
		"\u0147\u0146\u0001\u0000\u0000\u0000\u0148\u0149\u0001\u0000\u0000\u0000"+
		"\u0149\u0147\u0001\u0000\u0000\u0000\u0149\u014a\u0001\u0000\u0000\u0000"+
		"\u014a7\u0001\u0000\u0000\u0000\u014b\u014d\u0003b1\u0000\u014c\u014b"+
		"\u0001\u0000\u0000\u0000\u014d\u014e\u0001\u0000\u0000\u0000\u014e\u014c"+
		"\u0001\u0000\u0000\u0000\u014e\u014f\u0001\u0000\u0000\u0000\u014f9\u0001"+
		"\u0000\u0000\u0000\u0150\u0152\u0003`0\u0000\u0151\u0150\u0001\u0000\u0000"+
		"\u0000\u0152\u0153\u0001\u0000\u0000\u0000\u0153\u0151\u0001\u0000\u0000"+
		"\u0000\u0153\u0154\u0001\u0000\u0000\u0000\u0154;\u0001\u0000\u0000\u0000"+
		"\u0155\u0157\u0003h4\u0000\u0156\u0155\u0001\u0000\u0000\u0000\u0157\u0158"+
		"\u0001\u0000\u0000\u0000\u0158\u0156\u0001\u0000\u0000\u0000\u0158\u0159"+
		"\u0001\u0000\u0000\u0000\u0159=\u0001\u0000\u0000\u0000\u015a\u015c\u0003"+
		"l6\u0000\u015b\u015a\u0001\u0000\u0000\u0000\u015c\u015d\u0001\u0000\u0000"+
		"\u0000\u015d\u015b\u0001\u0000\u0000\u0000\u015d\u015e\u0001\u0000\u0000"+
		"\u0000\u015e?\u0001\u0000\u0000\u0000\u015f\u0161\u0003n7\u0000\u0160"+
		"\u015f\u0001\u0000\u0000\u0000\u0161\u0162\u0001\u0000\u0000\u0000\u0162"+
		"\u0160\u0001\u0000\u0000\u0000\u0162\u0163\u0001\u0000\u0000\u0000\u0163"+
		"A\u0001\u0000\u0000\u0000\u0164\u0166\u0003p8\u0000\u0165\u0164\u0001"+
		"\u0000\u0000\u0000\u0166\u0167\u0001\u0000\u0000\u0000\u0167\u0165\u0001"+
		"\u0000\u0000\u0000\u0167\u0168\u0001\u0000\u0000\u0000\u0168C\u0001\u0000"+
		"\u0000\u0000\u0169\u016b\u0003t:\u0000\u016a\u0169\u0001\u0000\u0000\u0000"+
		"\u016b\u016c\u0001\u0000\u0000\u0000\u016c\u016a\u0001\u0000\u0000\u0000"+
		"\u016c\u016d\u0001\u0000\u0000\u0000\u016dE\u0001\u0000\u0000\u0000\u016e"+
		"\u0170\u0003v;\u0000\u016f\u016e\u0001\u0000\u0000\u0000\u0170\u0171\u0001"+
		"\u0000\u0000\u0000\u0171\u016f\u0001\u0000\u0000\u0000\u0171\u0172\u0001"+
		"\u0000\u0000\u0000\u0172G\u0001\u0000\u0000\u0000\u0173\u0175\u0003x<"+
		"\u0000\u0174\u0173\u0001\u0000\u0000\u0000\u0175\u0176\u0001\u0000\u0000"+
		"\u0000\u0176\u0174\u0001\u0000\u0000\u0000\u0176\u0177\u0001\u0000\u0000"+
		"\u0000\u0177I\u0001\u0000\u0000\u0000\u0178\u017a\u0003z=\u0000\u0179"+
		"\u0178\u0001\u0000\u0000\u0000\u017a\u017b\u0001\u0000\u0000\u0000\u017b"+
		"\u0179\u0001\u0000\u0000\u0000\u017b\u017c\u0001\u0000\u0000\u0000\u017c"+
		"K\u0001\u0000\u0000\u0000\u017d\u017f\u0003|>\u0000\u017e\u017d\u0001"+
		"\u0000\u0000\u0000\u017f\u0180\u0001\u0000\u0000\u0000\u0180\u017e\u0001"+
		"\u0000\u0000\u0000\u0180\u0181\u0001\u0000\u0000\u0000\u0181M\u0001\u0000"+
		"\u0000\u0000\u0182\u0184\u0003~?\u0000\u0183\u0182\u0001\u0000\u0000\u0000"+
		"\u0184\u0185\u0001\u0000\u0000\u0000\u0185\u0183\u0001\u0000\u0000\u0000"+
		"\u0185\u0186\u0001\u0000\u0000\u0000\u0186O\u0001\u0000\u0000\u0000\u0187"+
		"\u0189\u0003\u0080@\u0000\u0188\u0187\u0001\u0000\u0000\u0000\u0189\u018a"+
		"\u0001\u0000\u0000\u0000\u018a\u0188\u0001\u0000\u0000\u0000\u018a\u018b"+
		"\u0001\u0000\u0000\u0000\u018bQ\u0001\u0000\u0000\u0000\u018c\u018e\u0003"+
		"\u0088D\u0000\u018d\u018c\u0001\u0000\u0000\u0000\u018e\u018f\u0001\u0000"+
		"\u0000\u0000\u018f\u018d\u0001\u0000\u0000\u0000\u018f\u0190\u0001\u0000"+
		"\u0000\u0000\u0190S\u0001\u0000\u0000\u0000\u0191\u0196\u0003\u0004\u0002"+
		"\u0000\u0192\u0196\u0003\u0006\u0003\u0000\u0193\u0196\u0003\b\u0004\u0000"+
		"\u0194\u0196\u0003\u000e\u0007\u0000\u0195\u0191\u0001\u0000\u0000\u0000"+
		"\u0195\u0192\u0001\u0000\u0000\u0000\u0195\u0193\u0001\u0000\u0000\u0000"+
		"\u0195\u0194\u0001\u0000\u0000\u0000\u0196U\u0001\u0000\u0000\u0000\u0197"+
		"\u0198\u0003\u000e\u0007\u0000\u0198W\u0001\u0000\u0000\u0000\u0199\u01a2"+
		"\u0003\u0004\u0002\u0000\u019a\u01a2\u0003\u0006\u0003\u0000\u019b\u01a2"+
		"\u0003\b\u0004\u0000\u019c\u01a2\u0003\u000e\u0007\u0000\u019d\u01a2\u0003"+
		"\n\u0005\u0000\u019e\u01a2\u0003\u0010\b\u0000\u019f\u01a2\u0003\u001e"+
		"\u000f\u0000\u01a0\u01a2\u0003 \u0010\u0000\u01a1\u0199\u0001\u0000\u0000"+
		"\u0000\u01a1\u019a\u0001\u0000\u0000\u0000\u01a1\u019b\u0001\u0000\u0000"+
		"\u0000\u01a1\u019c\u0001\u0000\u0000\u0000\u01a1\u019d\u0001\u0000\u0000"+
		"\u0000\u01a1\u019e\u0001\u0000\u0000\u0000\u01a1\u019f\u0001\u0000\u0000"+
		"\u0000\u01a1\u01a0\u0001\u0000\u0000\u0000\u01a2Y\u0001\u0000\u0000\u0000"+
		"\u01a3\u01a6\u0003\u000e\u0007\u0000\u01a4\u01a6\u0003\u0010\b\u0000\u01a5"+
		"\u01a3\u0001\u0000\u0000\u0000\u01a5\u01a4\u0001\u0000\u0000\u0000\u01a6"+
		"[\u0001\u0000\u0000\u0000\u01a7\u01a8\u0003\f\u0006\u0000\u01a8]\u0001"+
		"\u0000\u0000\u0000\u01a9\u01ad\u0003\f\u0006\u0000\u01aa\u01ad\u0003\n"+
		"\u0005\u0000\u01ab\u01ad\u0003\"\u0011\u0000\u01ac\u01a9\u0001\u0000\u0000"+
		"\u0000\u01ac\u01aa\u0001\u0000\u0000\u0000\u01ac\u01ab\u0001\u0000\u0000"+
		"\u0000\u01ad_\u0001\u0000\u0000\u0000\u01ae\u01af\u0003\u0012\t\u0000"+
		"\u01afa\u0001\u0000\u0000\u0000\u01b0\u01b4\u0003\u0014\n\u0000\u01b1"+
		"\u01b4\u0003\u0016\u000b\u0000\u01b2\u01b4\u0003\f\u0006\u0000\u01b3\u01b0"+
		"\u0001\u0000\u0000\u0000\u01b3\u01b1\u0001\u0000\u0000\u0000\u01b3\u01b2"+
		"\u0001\u0000\u0000\u0000\u01b4c\u0001\u0000\u0000\u0000\u01b5\u01b6\u0003"+
		"\u0010\b\u0000\u01b6e\u0001\u0000\u0000\u0000\u01b7\u01b8\u0003\u0010"+
		"\b\u0000\u01b8g\u0001\u0000\u0000\u0000\u01b9\u01bd\u0003\u0014\n\u0000"+
		"\u01ba\u01bd\u0003\u0016\u000b\u0000\u01bb\u01bd\u0003\u0018\f\u0000\u01bc"+
		"\u01b9\u0001\u0000\u0000\u0000\u01bc\u01ba\u0001\u0000\u0000\u0000\u01bc"+
		"\u01bb\u0001\u0000\u0000\u0000\u01bdi\u0001\u0000\u0000\u0000\u01be\u01bf"+
		"\u0003\u0014\n\u0000\u01bfk\u0001\u0000\u0000\u0000\u01c0\u01c3\u0003"+
		"\u0014\n\u0000\u01c1\u01c3\u0003\u0018\f\u0000\u01c2\u01c0\u0001\u0000"+
		"\u0000\u0000\u01c2\u01c1\u0001\u0000\u0000\u0000\u01c3m\u0001\u0000\u0000"+
		"\u0000\u01c4\u01c9\u0003\u000e\u0007\u0000\u01c5\u01c9\u0003.\u0017\u0000"+
		"\u01c6\u01c9\u0003\u0004\u0002\u0000\u01c7\u01c9\u0003\u001a\r\u0000\u01c8"+
		"\u01c4\u0001\u0000\u0000\u0000\u01c8\u01c5\u0001\u0000\u0000\u0000\u01c8"+
		"\u01c6\u0001\u0000\u0000\u0000\u01c8\u01c7\u0001\u0000\u0000\u0000\u01c9"+
		"o\u0001\u0000\u0000\u0000\u01ca\u01cd\u0003,\u0016\u0000\u01cb\u01cd\u0003"+
		"\u000e\u0007\u0000\u01cc\u01ca\u0001\u0000\u0000\u0000\u01cc\u01cb\u0001"+
		"\u0000\u0000\u0000\u01cdq\u0001\u0000\u0000\u0000\u01ce\u01cf\u0003\u000e"+
		"\u0007\u0000\u01cfs\u0001\u0000\u0000\u0000\u01d0\u01d3\u0003\u000e\u0007"+
		"\u0000\u01d1\u01d3\u0003$\u0012\u0000\u01d2\u01d0\u0001\u0000\u0000\u0000"+
		"\u01d2\u01d1\u0001\u0000\u0000\u0000\u01d3u\u0001\u0000\u0000\u0000\u01d4"+
		"\u01d7\u0003\u000e\u0007\u0000\u01d5\u01d7\u0003\u0010\b\u0000\u01d6\u01d4"+
		"\u0001\u0000\u0000\u0000\u01d6\u01d5\u0001\u0000\u0000\u0000\u01d7w\u0001"+
		"\u0000\u0000\u0000\u01d8\u01db\u0003\u000e\u0007\u0000\u01d9\u01db\u0003"+
		"&\u0013\u0000\u01da\u01d8\u0001\u0000\u0000\u0000\u01da\u01d9\u0001\u0000"+
		"\u0000\u0000\u01dby\u0001\u0000\u0000\u0000\u01dc\u01df\u0003\u000e\u0007"+
		"\u0000\u01dd\u01df\u0003&\u0013\u0000\u01de\u01dc\u0001\u0000\u0000\u0000"+
		"\u01de\u01dd\u0001\u0000\u0000\u0000\u01df{\u0001\u0000\u0000\u0000\u01e0"+
		"\u01e3\u0003\u000e\u0007\u0000\u01e1\u01e3\u0003\u0010\b\u0000\u01e2\u01e0"+
		"\u0001\u0000\u0000\u0000\u01e2\u01e1\u0001\u0000\u0000\u0000\u01e3}\u0001"+
		"\u0000\u0000\u0000\u01e4\u01e8\u0003\u000e\u0007\u0000\u01e5\u01e8\u0003"+
		"(\u0014\u0000\u01e6\u01e8\u0003.\u0017\u0000\u01e7\u01e4\u0001\u0000\u0000"+
		"\u0000\u01e7\u01e5\u0001\u0000\u0000\u0000\u01e7\u01e6\u0001\u0000\u0000"+
		"\u0000\u01e8\u007f\u0001\u0000\u0000\u0000\u01e9\u01ed\u0003\u000e\u0007"+
		"\u0000\u01ea\u01ed\u0003*\u0015\u0000\u01eb\u01ed\u0003.\u0017\u0000\u01ec"+
		"\u01e9\u0001\u0000\u0000\u0000\u01ec\u01ea\u0001\u0000\u0000\u0000\u01ec"+
		"\u01eb\u0001\u0000\u0000\u0000\u01ed\u0081\u0001\u0000\u0000\u0000\u01ee"+
		"\u01ef\u0003\f\u0006\u0000\u01ef\u0083\u0001\u0000\u0000\u0000\u01f0\u01f1"+
		"\u0003\f\u0006\u0000\u01f1\u0085\u0001\u0000\u0000\u0000\u01f2\u01f3\u0003"+
		"\f\u0006\u0000\u01f3\u0087\u0001\u0000\u0000\u0000\u01f4\u01f9\u0003\u0010"+
		"\b\u0000\u01f5\u01f9\u0003\u000e\u0007\u0000\u01f6\u01f9\u0003\f\u0006"+
		"\u0000\u01f7\u01f9\u0003\u001c\u000e\u0000\u01f8\u01f4\u0001\u0000\u0000"+
		"\u0000\u01f8\u01f5\u0001\u0000\u0000\u0000\u01f8\u01f6\u0001\u0000\u0000"+
		"\u0000\u01f8\u01f7\u0001\u0000\u0000\u0000\u01f9\u0089\u0001\u0000\u0000"+
		"\u0000%\u008d\u00ca\u013a\u013f\u0144\u0149\u014e\u0153\u0158\u015d\u0162"+
		"\u0167\u016c\u0171\u0176\u017b\u0180\u0185\u018a\u018f\u0195\u01a1\u01a5"+
		"\u01ac\u01b3\u01bc\u01c2\u01c8\u01cc\u01d2\u01d6\u01da\u01de\u01e2\u01e7"+
		"\u01ec\u01f8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}