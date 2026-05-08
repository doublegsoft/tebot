// Generated from /Users/christian/export/local/works/doublegsoft.net/tebot/03.Development/tebot/gram/Tebot.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TebotParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, TEBOT_WHITESPACE=16, 
		TEBOT_COMMENT=17, TEBOT_QUOTED_STRING=18;
	public static final int
		RULE_tebot_selector = 0, RULE_tebot_value = 1, RULE_tebot_action = 2, 
		RULE_tebot_assign = 3, RULE_tebot_operation = 4, RULE_tebot_assert = 5, 
		RULE_tebot_operations = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"tebot_selector", "tebot_value", "tebot_action", "tebot_assign", "tebot_operation", 
			"tebot_assert", "tebot_operations"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'click'", "'input'", "'select'", "'capture'", "'assert'", "'sleep'", 
			"'goto'", "'move'", "'scroll'", "'save'", "'paste'", "'='", "'('", "')'", 
			"','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "TEBOT_WHITESPACE", "TEBOT_COMMENT", "TEBOT_QUOTED_STRING"
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
	public String getGrammarFileName() { return "Tebot.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public TebotParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class Tebot_selectorContext extends ParserRuleContext {
		public TerminalNode TEBOT_QUOTED_STRING() { return getToken(TebotParser.TEBOT_QUOTED_STRING, 0); }
		public Tebot_selectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_selector; }
	}

	public final Tebot_selectorContext tebot_selector() throws RecognitionException {
		Tebot_selectorContext _localctx = new Tebot_selectorContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_tebot_selector);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
			match(TEBOT_QUOTED_STRING);
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

	public static class Tebot_valueContext extends ParserRuleContext {
		public TerminalNode TEBOT_QUOTED_STRING() { return getToken(TebotParser.TEBOT_QUOTED_STRING, 0); }
		public Tebot_valueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_value; }
	}

	public final Tebot_valueContext tebot_value() throws RecognitionException {
		Tebot_valueContext _localctx = new Tebot_valueContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_tebot_value);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(16);
			match(TEBOT_QUOTED_STRING);
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

	public static class Tebot_actionContext extends ParserRuleContext {
		public Tebot_actionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_action; }
	}

	public final Tebot_actionContext tebot_action() throws RecognitionException {
		Tebot_actionContext _localctx = new Tebot_actionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_tebot_action);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(18);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10))) != 0)) ) {
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

	public static class Tebot_assignContext extends ParserRuleContext {
		public Tebot_selectorContext selector;
		public Tebot_valueContext value;
		public Tebot_valueContext tebot_value() {
			return getRuleContext(Tebot_valueContext.class,0);
		}
		public Tebot_selectorContext tebot_selector() {
			return getRuleContext(Tebot_selectorContext.class,0);
		}
		public Tebot_assignContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_assign; }
	}

	public final Tebot_assignContext tebot_assign() throws RecognitionException {
		Tebot_assignContext _localctx = new Tebot_assignContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_tebot_assign);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(23);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(20);
				((Tebot_assignContext)_localctx).selector = tebot_selector();
				setState(21);
				match(T__11);
				}
				break;
			}
			setState(25);
			((Tebot_assignContext)_localctx).value = tebot_value();
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

	public static class Tebot_operationContext extends ParserRuleContext {
		public Tebot_actionContext tebot_action() {
			return getRuleContext(Tebot_actionContext.class,0);
		}
		public Tebot_selectorContext tebot_selector() {
			return getRuleContext(Tebot_selectorContext.class,0);
		}
		public Tebot_assignContext tebot_assign() {
			return getRuleContext(Tebot_assignContext.class,0);
		}
		public Tebot_assertContext tebot_assert() {
			return getRuleContext(Tebot_assertContext.class,0);
		}
		public Tebot_operationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_operation; }
	}

	public final Tebot_operationContext tebot_operation() throws RecognitionException {
		Tebot_operationContext _localctx = new Tebot_operationContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_tebot_operation);
		try {
			setState(36);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				tebot_action();
				setState(28);
				match(T__12);
				setState(31);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
				case 1:
					{
					setState(29);
					tebot_selector();
					}
					break;
				case 2:
					{
					setState(30);
					tebot_assign();
					}
					break;
				}
				setState(33);
				match(T__13);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(35);
				tebot_assert();
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

	public static class Tebot_assertContext extends ParserRuleContext {
		public Tebot_valueContext dsn;
		public Tebot_valueContext sql;
		public Tebot_valueContext expected;
		public List<Tebot_valueContext> tebot_value() {
			return getRuleContexts(Tebot_valueContext.class);
		}
		public Tebot_valueContext tebot_value(int i) {
			return getRuleContext(Tebot_valueContext.class,i);
		}
		public Tebot_assertContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_assert; }
	}

	public final Tebot_assertContext tebot_assert() throws RecognitionException {
		Tebot_assertContext _localctx = new Tebot_assertContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_tebot_assert);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(38);
			match(T__4);
			setState(39);
			match(T__12);
			setState(40);
			((Tebot_assertContext)_localctx).dsn = tebot_value();
			setState(41);
			match(T__14);
			setState(42);
			((Tebot_assertContext)_localctx).sql = tebot_value();
			setState(43);
			match(T__14);
			setState(44);
			((Tebot_assertContext)_localctx).expected = tebot_value();
			setState(45);
			match(T__13);
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

	public static class Tebot_operationsContext extends ParserRuleContext {
		public List<Tebot_operationContext> tebot_operation() {
			return getRuleContexts(Tebot_operationContext.class);
		}
		public Tebot_operationContext tebot_operation(int i) {
			return getRuleContext(Tebot_operationContext.class,i);
		}
		public Tebot_operationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tebot_operations; }
	}

	public final Tebot_operationsContext tebot_operations() throws RecognitionException {
		Tebot_operationsContext _localctx = new Tebot_operationsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_tebot_operations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__0) | (1L << T__1) | (1L << T__2) | (1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8) | (1L << T__9) | (1L << T__10))) != 0)) {
				{
				{
				setState(47);
				tebot_operation();
				}
				}
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\248\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\3\3\3\3\4\3\4\3\5"+
		"\3\5\3\5\5\5\32\n\5\3\5\3\5\3\6\3\6\3\6\3\6\5\6\"\n\6\3\6\3\6\3\6\5\6"+
		"\'\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\b\7\b\63\n\b\f\b\16\b\66"+
		"\13\b\3\b\2\2\t\2\4\6\b\n\f\16\2\3\3\2\3\r\2\64\2\20\3\2\2\2\4\22\3\2"+
		"\2\2\6\24\3\2\2\2\b\31\3\2\2\2\n&\3\2\2\2\f(\3\2\2\2\16\64\3\2\2\2\20"+
		"\21\7\24\2\2\21\3\3\2\2\2\22\23\7\24\2\2\23\5\3\2\2\2\24\25\t\2\2\2\25"+
		"\7\3\2\2\2\26\27\5\2\2\2\27\30\7\16\2\2\30\32\3\2\2\2\31\26\3\2\2\2\31"+
		"\32\3\2\2\2\32\33\3\2\2\2\33\34\5\4\3\2\34\t\3\2\2\2\35\36\5\6\4\2\36"+
		"!\7\17\2\2\37\"\5\2\2\2 \"\5\b\5\2!\37\3\2\2\2! \3\2\2\2\"#\3\2\2\2#$"+
		"\7\20\2\2$\'\3\2\2\2%\'\5\f\7\2&\35\3\2\2\2&%\3\2\2\2\'\13\3\2\2\2()\7"+
		"\7\2\2)*\7\17\2\2*+\5\4\3\2+,\7\21\2\2,-\5\4\3\2-.\7\21\2\2./\5\4\3\2"+
		"/\60\7\20\2\2\60\r\3\2\2\2\61\63\5\n\6\2\62\61\3\2\2\2\63\66\3\2\2\2\64"+
		"\62\3\2\2\2\64\65\3\2\2\2\65\17\3\2\2\2\66\64\3\2\2\2\6\31!&\64";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}