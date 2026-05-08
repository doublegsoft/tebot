// Generated from /Users/christian/export/local/works/doublegsoft.net/tebot/03.Development/tebot/gram/Tebot.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class TebotLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, TEBOT_WHITESPACE=16, 
		TEBOT_COMMENT=17, TEBOT_QUOTED_STRING=18;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "TEBOT_WHITESPACE", 
			"TEBOT_COMMENT", "TEBOT_QUOTED_STRING"
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


	public TebotLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Tebot.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\24\u008c\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3"+
		"\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\6\21u\n\21\r"+
		"\21\16\21v\3\21\3\21\3\22\3\22\7\22}\n\22\f\22\16\22\u0080\13\22\3\22"+
		"\3\22\3\23\3\23\7\23\u0086\n\23\f\23\16\23\u0089\13\23\3\23\3\23\2\2\24"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\3\2\5\5\2\13\f\17\17\"\"\4\2\f\f\17\17\5\2\f\f\17\17"+
		"$$\2\u008e\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2"+
		"\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27"+
		"\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2"+
		"\2\2#\3\2\2\2\2%\3\2\2\2\3\'\3\2\2\2\5-\3\2\2\2\7\63\3\2\2\2\t:\3\2\2"+
		"\2\13B\3\2\2\2\rI\3\2\2\2\17O\3\2\2\2\21T\3\2\2\2\23Y\3\2\2\2\25`\3\2"+
		"\2\2\27e\3\2\2\2\31k\3\2\2\2\33m\3\2\2\2\35o\3\2\2\2\37q\3\2\2\2!t\3\2"+
		"\2\2#z\3\2\2\2%\u0083\3\2\2\2\'(\7e\2\2()\7n\2\2)*\7k\2\2*+\7e\2\2+,\7"+
		"m\2\2,\4\3\2\2\2-.\7k\2\2./\7p\2\2/\60\7r\2\2\60\61\7w\2\2\61\62\7v\2"+
		"\2\62\6\3\2\2\2\63\64\7u\2\2\64\65\7g\2\2\65\66\7n\2\2\66\67\7g\2\2\67"+
		"8\7e\2\289\7v\2\29\b\3\2\2\2:;\7e\2\2;<\7c\2\2<=\7r\2\2=>\7v\2\2>?\7w"+
		"\2\2?@\7t\2\2@A\7g\2\2A\n\3\2\2\2BC\7c\2\2CD\7u\2\2DE\7u\2\2EF\7g\2\2"+
		"FG\7t\2\2GH\7v\2\2H\f\3\2\2\2IJ\7u\2\2JK\7n\2\2KL\7g\2\2LM\7g\2\2MN\7"+
		"r\2\2N\16\3\2\2\2OP\7i\2\2PQ\7q\2\2QR\7v\2\2RS\7q\2\2S\20\3\2\2\2TU\7"+
		"o\2\2UV\7q\2\2VW\7x\2\2WX\7g\2\2X\22\3\2\2\2YZ\7u\2\2Z[\7e\2\2[\\\7t\2"+
		"\2\\]\7q\2\2]^\7n\2\2^_\7n\2\2_\24\3\2\2\2`a\7u\2\2ab\7c\2\2bc\7x\2\2"+
		"cd\7g\2\2d\26\3\2\2\2ef\7r\2\2fg\7c\2\2gh\7u\2\2hi\7v\2\2ij\7g\2\2j\30"+
		"\3\2\2\2kl\7?\2\2l\32\3\2\2\2mn\7*\2\2n\34\3\2\2\2op\7+\2\2p\36\3\2\2"+
		"\2qr\7.\2\2r \3\2\2\2su\t\2\2\2ts\3\2\2\2uv\3\2\2\2vt\3\2\2\2vw\3\2\2"+
		"\2wx\3\2\2\2xy\b\21\2\2y\"\3\2\2\2z~\7%\2\2{}\n\3\2\2|{\3\2\2\2}\u0080"+
		"\3\2\2\2~|\3\2\2\2~\177\3\2\2\2\177\u0081\3\2\2\2\u0080~\3\2\2\2\u0081"+
		"\u0082\b\22\2\2\u0082$\3\2\2\2\u0083\u0087\7$\2\2\u0084\u0086\n\4\2\2"+
		"\u0085\u0084\3\2\2\2\u0086\u0089\3\2\2\2\u0087\u0085\3\2\2\2\u0087\u0088"+
		"\3\2\2\2\u0088\u008a\3\2\2\2\u0089\u0087\3\2\2\2\u008a\u008b\7$\2\2\u008b"+
		"&\3\2\2\2\6\2v~\u0087\3\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}