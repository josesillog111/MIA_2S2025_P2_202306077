// Generated from /home/jose06/Documentos/MIA_2S2025_P1_202306077/backend/parser/GoDiskGrammar.g4 by ANTLR 4.13.1
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
		FDISK=7, MOUNT=8, MOUNTED=9, MKFS=10, CAT=11, LOGIN=12, LOGOUT=13, MKGRP=14, 
		RMGRP=15, MKUSR=16, RMUSR=17, CHGRP=18, MKFILE=19, MKDIR=20, REP=21, ID_TEXT=22, 
		SIZE=23, FIT=24, UNIT=25, PATH=26, TYPE=27, NAME=28, FILE=29, USER=30, 
		PASS=31, GRP=32, CONT=33, PATH_FILE_LS=34, R=35, P=36, ASSIGN=37, MINUS=38, 
		ID=39, UNQUOTED_TEXT=40, LINE_COMMENT=41, BLOCK_COMMENT=42, WS=43;
	public static final int
		RULE_prog = 0, RULE_stmt = 1, RULE_size = 2, RULE_fit = 3, RULE_unit = 4, 
		RULE_type = 5, RULE_id_text = 6, RULE_path = 7, RULE_name = 8, RULE_filen = 9, 
		RULE_user = 10, RULE_pass = 11, RULE_grp = 12, RULE_cont = 13, RULE_path_file_ls = 14, 
		RULE_p = 15, RULE_r = 16, RULE_mkdisk_params = 17, RULE_fdisk_params = 18, 
		RULE_mount_params = 19, RULE_mkfs_params = 20, RULE_login_params = 21, 
		RULE_cat_params = 22, RULE_mkusr_params = 23, RULE_chgrp_params = 24, 
		RULE_mkfile_params = 25, RULE_mkdir_params = 26, RULE_rep_params = 27, 
		RULE_mkdisk_param = 28, RULE_rmdisk_param = 29, RULE_fdisk_param = 30, 
		RULE_mount_param = 31, RULE_mkfs_param = 32, RULE_cat_param = 33, RULE_login_param = 34, 
		RULE_mkgrp_param = 35, RULE_rmgrp_param = 36, RULE_mkusr_param = 37, RULE_rmusr_param = 38, 
		RULE_chgrp_param = 39, RULE_mkfile_param = 40, RULE_mkdir_param = 41, 
		RULE_rep_param = 42;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "stmt", "size", "fit", "unit", "type", "id_text", "path", "name", 
			"filen", "user", "pass", "grp", "cont", "path_file_ls", "p", "r", "mkdisk_params", 
			"fdisk_params", "mount_params", "mkfs_params", "login_params", "cat_params", 
			"mkusr_params", "chgrp_params", "mkfile_params", "mkdir_params", "rep_params", 
			"mkdisk_param", "rmdisk_param", "fdisk_param", "mount_param", "mkfs_param", 
			"cat_param", "login_param", "mkgrp_param", "rmgrp_param", "mkusr_param", 
			"rmusr_param", "chgrp_param", "mkfile_param", "mkdir_param", "rep_param"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "'='", "'-'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT_LIT", "FLOAT_LIT", "STRING_LIT", "CHAR_LIT", "MKDISK", "RMDISK", 
			"FDISK", "MOUNT", "MOUNTED", "MKFS", "CAT", "LOGIN", "LOGOUT", "MKGRP", 
			"RMGRP", "MKUSR", "RMUSR", "CHGRP", "MKFILE", "MKDIR", "REP", "ID_TEXT", 
			"SIZE", "FIT", "UNIT", "PATH", "TYPE", "NAME", "FILE", "USER", "PASS", 
			"GRP", "CONT", "PATH_FILE_LS", "R", "P", "ASSIGN", "MINUS", "ID", "UNQUOTED_TEXT", 
			"LINE_COMMENT", "BLOCK_COMMENT", "WS"
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
			setState(89);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4194272L) != 0)) {
				{
				{
				setState(86);
				stmt();
				}
				}
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(92);
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
	public static class LOGOUTContext extends StmtContext {
		public TerminalNode LOGOUT() { return getToken(GoDiskGrammar.LOGOUT, 0); }
		public LOGOUTContext(StmtContext ctx) { copyFrom(ctx); }
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
	public static class MKDIRContext extends StmtContext {
		public TerminalNode MKDIR() { return getToken(GoDiskGrammar.MKDIR, 0); }
		public Mkdir_paramsContext mkdir_params() {
			return getRuleContext(Mkdir_paramsContext.class,0);
		}
		public MKDIRContext(StmtContext ctx) { copyFrom(ctx); }
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
	public static class RMDISKContext extends StmtContext {
		public TerminalNode RMDISK() { return getToken(GoDiskGrammar.RMDISK, 0); }
		public Rmdisk_paramContext rmdisk_param() {
			return getRuleContext(Rmdisk_paramContext.class,0);
		}
		public RMDISKContext(StmtContext ctx) { copyFrom(ctx); }
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
	public static class FDISKContext extends StmtContext {
		public TerminalNode FDISK() { return getToken(GoDiskGrammar.FDISK, 0); }
		public Fdisk_paramsContext fdisk_params() {
			return getRuleContext(Fdisk_paramsContext.class,0);
		}
		public FDISKContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MOUNTEDContext extends StmtContext {
		public TerminalNode MOUNTED() { return getToken(GoDiskGrammar.MOUNTED, 0); }
		public MOUNTEDContext(StmtContext ctx) { copyFrom(ctx); }
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
	public static class MKDISKContext extends StmtContext {
		public TerminalNode MKDISK() { return getToken(GoDiskGrammar.MKDISK, 0); }
		public Mkdisk_paramsContext mkdisk_params() {
			return getRuleContext(Mkdisk_paramsContext.class,0);
		}
		public MKDISKContext(StmtContext ctx) { copyFrom(ctx); }
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
			setState(126);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MKDISK:
				_localctx = new MKDISKContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(94);
				match(MKDISK);
				setState(95);
				mkdisk_params();
				}
				break;
			case RMDISK:
				_localctx = new RMDISKContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(96);
				match(RMDISK);
				setState(97);
				rmdisk_param();
				}
				break;
			case FDISK:
				_localctx = new FDISKContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(98);
				match(FDISK);
				setState(99);
				fdisk_params();
				}
				break;
			case MOUNT:
				_localctx = new MOUNTContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(100);
				match(MOUNT);
				setState(101);
				mount_params();
				}
				break;
			case MOUNTED:
				_localctx = new MOUNTEDContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(102);
				match(MOUNTED);
				}
				break;
			case MKFS:
				_localctx = new MKFSContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(103);
				match(MKFS);
				setState(104);
				mkfs_params();
				}
				break;
			case CAT:
				_localctx = new CATContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(105);
				match(CAT);
				setState(106);
				cat_params();
				}
				break;
			case LOGIN:
				_localctx = new LOGINContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(107);
				match(LOGIN);
				setState(108);
				login_params();
				}
				break;
			case LOGOUT:
				_localctx = new LOGOUTContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(109);
				match(LOGOUT);
				}
				break;
			case MKGRP:
				_localctx = new MKGRPContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(110);
				match(MKGRP);
				setState(111);
				mkgrp_param();
				}
				break;
			case RMGRP:
				_localctx = new RMGRPContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(112);
				match(RMGRP);
				setState(113);
				rmgrp_param();
				}
				break;
			case MKUSR:
				_localctx = new MKUSRContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(114);
				match(MKUSR);
				setState(115);
				mkusr_params();
				}
				break;
			case RMUSR:
				_localctx = new RMUSRContext(_localctx);
				enterOuterAlt(_localctx, 13);
				{
				setState(116);
				match(RMUSR);
				setState(117);
				rmusr_param();
				}
				break;
			case CHGRP:
				_localctx = new CHGRPContext(_localctx);
				enterOuterAlt(_localctx, 14);
				{
				setState(118);
				match(CHGRP);
				setState(119);
				chgrp_params();
				}
				break;
			case MKFILE:
				_localctx = new MKFILEContext(_localctx);
				enterOuterAlt(_localctx, 15);
				{
				setState(120);
				match(MKFILE);
				setState(121);
				mkfile_params();
				}
				break;
			case MKDIR:
				_localctx = new MKDIRContext(_localctx);
				enterOuterAlt(_localctx, 16);
				{
				setState(122);
				match(MKDIR);
				setState(123);
				mkdir_params();
				}
				break;
			case REP:
				_localctx = new REPContext(_localctx);
				enterOuterAlt(_localctx, 17);
				{
				setState(124);
				match(REP);
				setState(125);
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
			setState(128);
			match(MINUS);
			setState(129);
			match(SIZE);
			setState(130);
			match(ASSIGN);
			setState(131);
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
			setState(133);
			match(MINUS);
			setState(134);
			match(FIT);
			setState(135);
			match(ASSIGN);
			setState(136);
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
			setState(138);
			match(MINUS);
			setState(139);
			match(UNIT);
			setState(140);
			match(ASSIGN);
			setState(141);
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
			setState(143);
			match(MINUS);
			setState(144);
			match(TYPE);
			setState(145);
			match(ASSIGN);
			setState(146);
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
			setState(148);
			match(MINUS);
			setState(149);
			match(ID_TEXT);
			setState(150);
			match(ASSIGN);
			setState(151);
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
			setState(153);
			match(MINUS);
			setState(154);
			match(PATH);
			setState(155);
			match(ASSIGN);
			setState(156);
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
			setState(158);
			match(MINUS);
			setState(159);
			match(NAME);
			setState(160);
			match(ASSIGN);
			setState(161);
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
			setState(163);
			match(MINUS);
			setState(164);
			match(FILE);
			setState(165);
			match(INT_LIT);
			setState(166);
			match(ASSIGN);
			setState(167);
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
			setState(169);
			match(MINUS);
			setState(170);
			match(USER);
			setState(171);
			match(ASSIGN);
			setState(172);
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
			setState(174);
			match(MINUS);
			setState(175);
			match(PASS);
			setState(176);
			match(ASSIGN);
			setState(177);
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
			setState(179);
			match(MINUS);
			setState(180);
			match(GRP);
			setState(181);
			match(ASSIGN);
			setState(182);
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
			setState(184);
			match(MINUS);
			setState(185);
			match(CONT);
			setState(186);
			match(ASSIGN);
			setState(187);
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
			setState(189);
			match(MINUS);
			setState(190);
			match(PATH_FILE_LS);
			setState(191);
			match(ASSIGN);
			setState(192);
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
		enterRule(_localctx, 30, RULE_p);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			match(MINUS);
			setState(195);
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
		enterRule(_localctx, 32, RULE_r);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(MINUS);
			setState(198);
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
		enterRule(_localctx, 34, RULE_mkdisk_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(201); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(200);
				mkdisk_param();
				}
				}
				setState(203); 
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
		enterRule(_localctx, 36, RULE_fdisk_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(205);
				fdisk_param();
				}
				}
				setState(208); 
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
		enterRule(_localctx, 38, RULE_mount_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(210);
				mount_param();
				}
				}
				setState(213); 
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
		enterRule(_localctx, 40, RULE_mkfs_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(215);
				mkfs_param();
				}
				}
				setState(218); 
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
		enterRule(_localctx, 42, RULE_login_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(221); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(220);
				login_param();
				}
				}
				setState(223); 
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
		enterRule(_localctx, 44, RULE_cat_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(226); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(225);
				cat_param();
				}
				}
				setState(228); 
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
		enterRule(_localctx, 46, RULE_mkusr_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(230);
				mkusr_param();
				}
				}
				setState(233); 
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
		enterRule(_localctx, 48, RULE_chgrp_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(235);
				chgrp_param();
				}
				}
				setState(238); 
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
		enterRule(_localctx, 50, RULE_mkfile_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(240);
				mkfile_param();
				}
				}
				setState(243); 
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
		enterRule(_localctx, 52, RULE_mkdir_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(246); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(245);
				mkdir_param();
				}
				}
				setState(248); 
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
		enterRule(_localctx, 54, RULE_rep_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(251); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(250);
				rep_param();
				}
				}
				setState(253); 
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
		enterRule(_localctx, 56, RULE_mkdisk_param);
		try {
			setState(259);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(255);
				size();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(256);
				fit();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(257);
				unit();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(258);
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
		enterRule(_localctx, 58, RULE_rmdisk_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
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
		public Fdisk_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fdisk_param; }
	}

	public final Fdisk_paramContext fdisk_param() throws RecognitionException {
		Fdisk_paramContext _localctx = new Fdisk_paramContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_fdisk_param);
		try {
			setState(269);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(263);
				size();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(264);
				fit();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(265);
				unit();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(266);
				path();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(267);
				type();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(268);
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
		enterRule(_localctx, 62, RULE_mount_param);
		try {
			setState(273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(271);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(272);
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
	public static class Mkfs_paramContext extends ParserRuleContext {
		public Id_textContext id_text() {
			return getRuleContext(Id_textContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public Mkfs_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mkfs_param; }
	}

	public final Mkfs_paramContext mkfs_param() throws RecognitionException {
		Mkfs_paramContext _localctx = new Mkfs_paramContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_mkfs_param);
		try {
			setState(277);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(275);
				id_text();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(276);
				type();
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
		enterRule(_localctx, 66, RULE_cat_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
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
		enterRule(_localctx, 68, RULE_login_param);
		try {
			setState(284);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(281);
				user();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(282);
				pass();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(283);
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
		enterRule(_localctx, 70, RULE_mkgrp_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
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
		enterRule(_localctx, 72, RULE_rmgrp_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
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
		enterRule(_localctx, 74, RULE_mkusr_param);
		try {
			setState(293);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(290);
				user();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(291);
				pass();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(292);
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
		enterRule(_localctx, 76, RULE_rmusr_param);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
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
		enterRule(_localctx, 78, RULE_chgrp_param);
		try {
			setState(299);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(297);
				user();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(298);
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
		enterRule(_localctx, 80, RULE_mkfile_param);
		try {
			setState(305);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(301);
				path();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(302);
				r();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(303);
				size();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(304);
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
		enterRule(_localctx, 82, RULE_mkdir_param);
		try {
			setState(309);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(307);
				p();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(308);
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
		enterRule(_localctx, 84, RULE_rep_param);
		try {
			setState(315);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(311);
				name();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(312);
				path();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(313);
				id_text();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(314);
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
		"\u0004\u0001+\u013e\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0001\u0000\u0005\u0000X\b\u0000"+
		"\n\u0000\f\u0000[\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0003\u0001\u007f\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0011\u0004\u0011\u00ca\b\u0011\u000b\u0011\f\u0011\u00cb\u0001"+
		"\u0012\u0004\u0012\u00cf\b\u0012\u000b\u0012\f\u0012\u00d0\u0001\u0013"+
		"\u0004\u0013\u00d4\b\u0013\u000b\u0013\f\u0013\u00d5\u0001\u0014\u0004"+
		"\u0014\u00d9\b\u0014\u000b\u0014\f\u0014\u00da\u0001\u0015\u0004\u0015"+
		"\u00de\b\u0015\u000b\u0015\f\u0015\u00df\u0001\u0016\u0004\u0016\u00e3"+
		"\b\u0016\u000b\u0016\f\u0016\u00e4\u0001\u0017\u0004\u0017\u00e8\b\u0017"+
		"\u000b\u0017\f\u0017\u00e9\u0001\u0018\u0004\u0018\u00ed\b\u0018\u000b"+
		"\u0018\f\u0018\u00ee\u0001\u0019\u0004\u0019\u00f2\b\u0019\u000b\u0019"+
		"\f\u0019\u00f3\u0001\u001a\u0004\u001a\u00f7\b\u001a\u000b\u001a\f\u001a"+
		"\u00f8\u0001\u001b\u0004\u001b\u00fc\b\u001b\u000b\u001b\f\u001b\u00fd"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u0104\b\u001c"+
		"\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0003\u001e\u010e\b\u001e\u0001\u001f\u0001\u001f"+
		"\u0003\u001f\u0112\b\u001f\u0001 \u0001 \u0003 \u0116\b \u0001!\u0001"+
		"!\u0001\"\u0001\"\u0001\"\u0003\"\u011d\b\"\u0001#\u0001#\u0001$\u0001"+
		"$\u0001%\u0001%\u0001%\u0003%\u0126\b%\u0001&\u0001&\u0001\'\u0001\'\u0003"+
		"\'\u012c\b\'\u0001(\u0001(\u0001(\u0001(\u0003(\u0132\b(\u0001)\u0001"+
		")\u0003)\u0136\b)\u0001*\u0001*\u0001*\u0001*\u0003*\u013c\b*\u0001*\u0000"+
		"\u0000+\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018"+
		"\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPRT\u0000\u0001\u0002\u0000"+
		"\u0003\u0003((\u0144\u0000Y\u0001\u0000\u0000\u0000\u0002~\u0001\u0000"+
		"\u0000\u0000\u0004\u0080\u0001\u0000\u0000\u0000\u0006\u0085\u0001\u0000"+
		"\u0000\u0000\b\u008a\u0001\u0000\u0000\u0000\n\u008f\u0001\u0000\u0000"+
		"\u0000\f\u0094\u0001\u0000\u0000\u0000\u000e\u0099\u0001\u0000\u0000\u0000"+
		"\u0010\u009e\u0001\u0000\u0000\u0000\u0012\u00a3\u0001\u0000\u0000\u0000"+
		"\u0014\u00a9\u0001\u0000\u0000\u0000\u0016\u00ae\u0001\u0000\u0000\u0000"+
		"\u0018\u00b3\u0001\u0000\u0000\u0000\u001a\u00b8\u0001\u0000\u0000\u0000"+
		"\u001c\u00bd\u0001\u0000\u0000\u0000\u001e\u00c2\u0001\u0000\u0000\u0000"+
		" \u00c5\u0001\u0000\u0000\u0000\"\u00c9\u0001\u0000\u0000\u0000$\u00ce"+
		"\u0001\u0000\u0000\u0000&\u00d3\u0001\u0000\u0000\u0000(\u00d8\u0001\u0000"+
		"\u0000\u0000*\u00dd\u0001\u0000\u0000\u0000,\u00e2\u0001\u0000\u0000\u0000"+
		".\u00e7\u0001\u0000\u0000\u00000\u00ec\u0001\u0000\u0000\u00002\u00f1"+
		"\u0001\u0000\u0000\u00004\u00f6\u0001\u0000\u0000\u00006\u00fb\u0001\u0000"+
		"\u0000\u00008\u0103\u0001\u0000\u0000\u0000:\u0105\u0001\u0000\u0000\u0000"+
		"<\u010d\u0001\u0000\u0000\u0000>\u0111\u0001\u0000\u0000\u0000@\u0115"+
		"\u0001\u0000\u0000\u0000B\u0117\u0001\u0000\u0000\u0000D\u011c\u0001\u0000"+
		"\u0000\u0000F\u011e\u0001\u0000\u0000\u0000H\u0120\u0001\u0000\u0000\u0000"+
		"J\u0125\u0001\u0000\u0000\u0000L\u0127\u0001\u0000\u0000\u0000N\u012b"+
		"\u0001\u0000\u0000\u0000P\u0131\u0001\u0000\u0000\u0000R\u0135\u0001\u0000"+
		"\u0000\u0000T\u013b\u0001\u0000\u0000\u0000VX\u0003\u0002\u0001\u0000"+
		"WV\u0001\u0000\u0000\u0000X[\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000"+
		"\u0000YZ\u0001\u0000\u0000\u0000Z\\\u0001\u0000\u0000\u0000[Y\u0001\u0000"+
		"\u0000\u0000\\]\u0005\u0000\u0000\u0001]\u0001\u0001\u0000\u0000\u0000"+
		"^_\u0005\u0005\u0000\u0000_\u007f\u0003\"\u0011\u0000`a\u0005\u0006\u0000"+
		"\u0000a\u007f\u0003:\u001d\u0000bc\u0005\u0007\u0000\u0000c\u007f\u0003"+
		"$\u0012\u0000de\u0005\b\u0000\u0000e\u007f\u0003&\u0013\u0000f\u007f\u0005"+
		"\t\u0000\u0000gh\u0005\n\u0000\u0000h\u007f\u0003(\u0014\u0000ij\u0005"+
		"\u000b\u0000\u0000j\u007f\u0003,\u0016\u0000kl\u0005\f\u0000\u0000l\u007f"+
		"\u0003*\u0015\u0000m\u007f\u0005\r\u0000\u0000no\u0005\u000e\u0000\u0000"+
		"o\u007f\u0003F#\u0000pq\u0005\u000f\u0000\u0000q\u007f\u0003H$\u0000r"+
		"s\u0005\u0010\u0000\u0000s\u007f\u0003.\u0017\u0000tu\u0005\u0011\u0000"+
		"\u0000u\u007f\u0003L&\u0000vw\u0005\u0012\u0000\u0000w\u007f\u00030\u0018"+
		"\u0000xy\u0005\u0013\u0000\u0000y\u007f\u00032\u0019\u0000z{\u0005\u0014"+
		"\u0000\u0000{\u007f\u00034\u001a\u0000|}\u0005\u0015\u0000\u0000}\u007f"+
		"\u00036\u001b\u0000~^\u0001\u0000\u0000\u0000~`\u0001\u0000\u0000\u0000"+
		"~b\u0001\u0000\u0000\u0000~d\u0001\u0000\u0000\u0000~f\u0001\u0000\u0000"+
		"\u0000~g\u0001\u0000\u0000\u0000~i\u0001\u0000\u0000\u0000~k\u0001\u0000"+
		"\u0000\u0000~m\u0001\u0000\u0000\u0000~n\u0001\u0000\u0000\u0000~p\u0001"+
		"\u0000\u0000\u0000~r\u0001\u0000\u0000\u0000~t\u0001\u0000\u0000\u0000"+
		"~v\u0001\u0000\u0000\u0000~x\u0001\u0000\u0000\u0000~z\u0001\u0000\u0000"+
		"\u0000~|\u0001\u0000\u0000\u0000\u007f\u0003\u0001\u0000\u0000\u0000\u0080"+
		"\u0081\u0005&\u0000\u0000\u0081\u0082\u0005\u0017\u0000\u0000\u0082\u0083"+
		"\u0005%\u0000\u0000\u0083\u0084\u0005\u0001\u0000\u0000\u0084\u0005\u0001"+
		"\u0000\u0000\u0000\u0085\u0086\u0005&\u0000\u0000\u0086\u0087\u0005\u0018"+
		"\u0000\u0000\u0087\u0088\u0005%\u0000\u0000\u0088\u0089\u0005\'\u0000"+
		"\u0000\u0089\u0007\u0001\u0000\u0000\u0000\u008a\u008b\u0005&\u0000\u0000"+
		"\u008b\u008c\u0005\u0019\u0000\u0000\u008c\u008d\u0005%\u0000\u0000\u008d"+
		"\u008e\u0005\'\u0000\u0000\u008e\t\u0001\u0000\u0000\u0000\u008f\u0090"+
		"\u0005&\u0000\u0000\u0090\u0091\u0005\u001b\u0000\u0000\u0091\u0092\u0005"+
		"%\u0000\u0000\u0092\u0093\u0005\'\u0000\u0000\u0093\u000b\u0001\u0000"+
		"\u0000\u0000\u0094\u0095\u0005&\u0000\u0000\u0095\u0096\u0005\u0016\u0000"+
		"\u0000\u0096\u0097\u0005%\u0000\u0000\u0097\u0098\u0005\'\u0000\u0000"+
		"\u0098\r\u0001\u0000\u0000\u0000\u0099\u009a\u0005&\u0000\u0000\u009a"+
		"\u009b\u0005\u001a\u0000\u0000\u009b\u009c\u0005%\u0000\u0000\u009c\u009d"+
		"\u0007\u0000\u0000\u0000\u009d\u000f\u0001\u0000\u0000\u0000\u009e\u009f"+
		"\u0005&\u0000\u0000\u009f\u00a0\u0005\u001c\u0000\u0000\u00a0\u00a1\u0005"+
		"%\u0000\u0000\u00a1\u00a2\u0007\u0000\u0000\u0000\u00a2\u0011\u0001\u0000"+
		"\u0000\u0000\u00a3\u00a4\u0005&\u0000\u0000\u00a4\u00a5\u0005\u001d\u0000"+
		"\u0000\u00a5\u00a6\u0005\u0001\u0000\u0000\u00a6\u00a7\u0005%\u0000\u0000"+
		"\u00a7\u00a8\u0007\u0000\u0000\u0000\u00a8\u0013\u0001\u0000\u0000\u0000"+
		"\u00a9\u00aa\u0005&\u0000\u0000\u00aa\u00ab\u0005\u001e\u0000\u0000\u00ab"+
		"\u00ac\u0005%\u0000\u0000\u00ac\u00ad\u0007\u0000\u0000\u0000\u00ad\u0015"+
		"\u0001\u0000\u0000\u0000\u00ae\u00af\u0005&\u0000\u0000\u00af\u00b0\u0005"+
		"\u001f\u0000\u0000\u00b0\u00b1\u0005%\u0000\u0000\u00b1\u00b2\u0007\u0000"+
		"\u0000\u0000\u00b2\u0017\u0001\u0000\u0000\u0000\u00b3\u00b4\u0005&\u0000"+
		"\u0000\u00b4\u00b5\u0005 \u0000\u0000\u00b5\u00b6\u0005%\u0000\u0000\u00b6"+
		"\u00b7\u0007\u0000\u0000\u0000\u00b7\u0019\u0001\u0000\u0000\u0000\u00b8"+
		"\u00b9\u0005&\u0000\u0000\u00b9\u00ba\u0005!\u0000\u0000\u00ba\u00bb\u0005"+
		"%\u0000\u0000\u00bb\u00bc\u0007\u0000\u0000\u0000\u00bc\u001b\u0001\u0000"+
		"\u0000\u0000\u00bd\u00be\u0005&\u0000\u0000\u00be\u00bf\u0005\"\u0000"+
		"\u0000\u00bf\u00c0\u0005%\u0000\u0000\u00c0\u00c1\u0007\u0000\u0000\u0000"+
		"\u00c1\u001d\u0001\u0000\u0000\u0000\u00c2\u00c3\u0005&\u0000\u0000\u00c3"+
		"\u00c4\u0005$\u0000\u0000\u00c4\u001f\u0001\u0000\u0000\u0000\u00c5\u00c6"+
		"\u0005&\u0000\u0000\u00c6\u00c7\u0005#\u0000\u0000\u00c7!\u0001\u0000"+
		"\u0000\u0000\u00c8\u00ca\u00038\u001c\u0000\u00c9\u00c8\u0001\u0000\u0000"+
		"\u0000\u00ca\u00cb\u0001\u0000\u0000\u0000\u00cb\u00c9\u0001\u0000\u0000"+
		"\u0000\u00cb\u00cc\u0001\u0000\u0000\u0000\u00cc#\u0001\u0000\u0000\u0000"+
		"\u00cd\u00cf\u0003<\u001e\u0000\u00ce\u00cd\u0001\u0000\u0000\u0000\u00cf"+
		"\u00d0\u0001\u0000\u0000\u0000\u00d0\u00ce\u0001\u0000\u0000\u0000\u00d0"+
		"\u00d1\u0001\u0000\u0000\u0000\u00d1%\u0001\u0000\u0000\u0000\u00d2\u00d4"+
		"\u0003>\u001f\u0000\u00d3\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d5\u0001"+
		"\u0000\u0000\u0000\u00d5\u00d3\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d6\'\u0001\u0000\u0000\u0000\u00d7\u00d9\u0003@"+
		" \u0000\u00d8\u00d7\u0001\u0000\u0000\u0000\u00d9\u00da\u0001\u0000\u0000"+
		"\u0000\u00da\u00d8\u0001\u0000\u0000\u0000\u00da\u00db\u0001\u0000\u0000"+
		"\u0000\u00db)\u0001\u0000\u0000\u0000\u00dc\u00de\u0003D\"\u0000\u00dd"+
		"\u00dc\u0001\u0000\u0000\u0000\u00de\u00df\u0001\u0000\u0000\u0000\u00df"+
		"\u00dd\u0001\u0000\u0000\u0000\u00df\u00e0\u0001\u0000\u0000\u0000\u00e0"+
		"+\u0001\u0000\u0000\u0000\u00e1\u00e3\u0003B!\u0000\u00e2\u00e1\u0001"+
		"\u0000\u0000\u0000\u00e3\u00e4\u0001\u0000\u0000\u0000\u00e4\u00e2\u0001"+
		"\u0000\u0000\u0000\u00e4\u00e5\u0001\u0000\u0000\u0000\u00e5-\u0001\u0000"+
		"\u0000\u0000\u00e6\u00e8\u0003J%\u0000\u00e7\u00e6\u0001\u0000\u0000\u0000"+
		"\u00e8\u00e9\u0001\u0000\u0000\u0000\u00e9\u00e7\u0001\u0000\u0000\u0000"+
		"\u00e9\u00ea\u0001\u0000\u0000\u0000\u00ea/\u0001\u0000\u0000\u0000\u00eb"+
		"\u00ed\u0003N\'\u0000\u00ec\u00eb\u0001\u0000\u0000\u0000\u00ed\u00ee"+
		"\u0001\u0000\u0000\u0000\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ee\u00ef"+
		"\u0001\u0000\u0000\u0000\u00ef1\u0001\u0000\u0000\u0000\u00f0\u00f2\u0003"+
		"P(\u0000\u00f1\u00f0\u0001\u0000\u0000\u0000\u00f2\u00f3\u0001\u0000\u0000"+
		"\u0000\u00f3\u00f1\u0001\u0000\u0000\u0000\u00f3\u00f4\u0001\u0000\u0000"+
		"\u0000\u00f43\u0001\u0000\u0000\u0000\u00f5\u00f7\u0003R)\u0000\u00f6"+
		"\u00f5\u0001\u0000\u0000\u0000\u00f7\u00f8\u0001\u0000\u0000\u0000\u00f8"+
		"\u00f6\u0001\u0000\u0000\u0000\u00f8\u00f9\u0001\u0000\u0000\u0000\u00f9"+
		"5\u0001\u0000\u0000\u0000\u00fa\u00fc\u0003T*\u0000\u00fb\u00fa\u0001"+
		"\u0000\u0000\u0000\u00fc\u00fd\u0001\u0000\u0000\u0000\u00fd\u00fb\u0001"+
		"\u0000\u0000\u0000\u00fd\u00fe\u0001\u0000\u0000\u0000\u00fe7\u0001\u0000"+
		"\u0000\u0000\u00ff\u0104\u0003\u0004\u0002\u0000\u0100\u0104\u0003\u0006"+
		"\u0003\u0000\u0101\u0104\u0003\b\u0004\u0000\u0102\u0104\u0003\u000e\u0007"+
		"\u0000\u0103\u00ff\u0001\u0000\u0000\u0000\u0103\u0100\u0001\u0000\u0000"+
		"\u0000\u0103\u0101\u0001\u0000\u0000\u0000\u0103\u0102\u0001\u0000\u0000"+
		"\u0000\u01049\u0001\u0000\u0000\u0000\u0105\u0106\u0003\u000e\u0007\u0000"+
		"\u0106;\u0001\u0000\u0000\u0000\u0107\u010e\u0003\u0004\u0002\u0000\u0108"+
		"\u010e\u0003\u0006\u0003\u0000\u0109\u010e\u0003\b\u0004\u0000\u010a\u010e"+
		"\u0003\u000e\u0007\u0000\u010b\u010e\u0003\n\u0005\u0000\u010c\u010e\u0003"+
		"\u0010\b\u0000\u010d\u0107\u0001\u0000\u0000\u0000\u010d\u0108\u0001\u0000"+
		"\u0000\u0000\u010d\u0109\u0001\u0000\u0000\u0000\u010d\u010a\u0001\u0000"+
		"\u0000\u0000\u010d\u010b\u0001\u0000\u0000\u0000\u010d\u010c\u0001\u0000"+
		"\u0000\u0000\u010e=\u0001\u0000\u0000\u0000\u010f\u0112\u0003\u000e\u0007"+
		"\u0000\u0110\u0112\u0003\u0010\b\u0000\u0111\u010f\u0001\u0000\u0000\u0000"+
		"\u0111\u0110\u0001\u0000\u0000\u0000\u0112?\u0001\u0000\u0000\u0000\u0113"+
		"\u0116\u0003\f\u0006\u0000\u0114\u0116\u0003\n\u0005\u0000\u0115\u0113"+
		"\u0001\u0000\u0000\u0000\u0115\u0114\u0001\u0000\u0000\u0000\u0116A\u0001"+
		"\u0000\u0000\u0000\u0117\u0118\u0003\u0012\t\u0000\u0118C\u0001\u0000"+
		"\u0000\u0000\u0119\u011d\u0003\u0014\n\u0000\u011a\u011d\u0003\u0016\u000b"+
		"\u0000\u011b\u011d\u0003\f\u0006\u0000\u011c\u0119\u0001\u0000\u0000\u0000"+
		"\u011c\u011a\u0001\u0000\u0000\u0000\u011c\u011b\u0001\u0000\u0000\u0000"+
		"\u011dE\u0001\u0000\u0000\u0000\u011e\u011f\u0003\u0010\b\u0000\u011f"+
		"G\u0001\u0000\u0000\u0000\u0120\u0121\u0003\u0010\b\u0000\u0121I\u0001"+
		"\u0000\u0000\u0000\u0122\u0126\u0003\u0014\n\u0000\u0123\u0126\u0003\u0016"+
		"\u000b\u0000\u0124\u0126\u0003\u0018\f\u0000\u0125\u0122\u0001\u0000\u0000"+
		"\u0000\u0125\u0123\u0001\u0000\u0000\u0000\u0125\u0124\u0001\u0000\u0000"+
		"\u0000\u0126K\u0001\u0000\u0000\u0000\u0127\u0128\u0003\u0014\n\u0000"+
		"\u0128M\u0001\u0000\u0000\u0000\u0129\u012c\u0003\u0014\n\u0000\u012a"+
		"\u012c\u0003\u0018\f\u0000\u012b\u0129\u0001\u0000\u0000\u0000\u012b\u012a"+
		"\u0001\u0000\u0000\u0000\u012cO\u0001\u0000\u0000\u0000\u012d\u0132\u0003"+
		"\u000e\u0007\u0000\u012e\u0132\u0003 \u0010\u0000\u012f\u0132\u0003\u0004"+
		"\u0002\u0000\u0130\u0132\u0003\u001a\r\u0000\u0131\u012d\u0001\u0000\u0000"+
		"\u0000\u0131\u012e\u0001\u0000\u0000\u0000\u0131\u012f\u0001\u0000\u0000"+
		"\u0000\u0131\u0130\u0001\u0000\u0000\u0000\u0132Q\u0001\u0000\u0000\u0000"+
		"\u0133\u0136\u0003\u001e\u000f\u0000\u0134\u0136\u0003\u000e\u0007\u0000"+
		"\u0135\u0133\u0001\u0000\u0000\u0000\u0135\u0134\u0001\u0000\u0000\u0000"+
		"\u0136S\u0001\u0000\u0000\u0000\u0137\u013c\u0003\u0010\b\u0000\u0138"+
		"\u013c\u0003\u000e\u0007\u0000\u0139\u013c\u0003\f\u0006\u0000\u013a\u013c"+
		"\u0003\u001c\u000e\u0000\u013b\u0137\u0001\u0000\u0000\u0000\u013b\u0138"+
		"\u0001\u0000\u0000\u0000\u013b\u0139\u0001\u0000\u0000\u0000\u013b\u013a"+
		"\u0001\u0000\u0000\u0000\u013cU\u0001\u0000\u0000\u0000\u0017Y~\u00cb"+
		"\u00d0\u00d5\u00da\u00df\u00e4\u00e9\u00ee\u00f3\u00f8\u00fd\u0103\u010d"+
		"\u0111\u0115\u011c\u0125\u012b\u0131\u0135\u013b";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}