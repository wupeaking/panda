package token

// 定义说有的token类型

const (
	// 关键字
	_ = iota
	BREAK
	FUNC
	STRUCT
	ELSE
	GOTO
	PACKAGE
	SWITCH
	IF
	TYPE
	CONTINUE
	FOR
	IMPORT
	RETURN
	VAR
	NILL
	MAP
	LIST
	TRUE
	FALSE

	// 标识符
	IDENTIFIER

	// 标点
	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]
	COMMA    // ,
	SEMI     // ;
	COLON    // ,
	DOT      // .

	// Logical
	LOGICALOR  // '||';
	LOGICALAND // '&&';

	BANG // "!"
	// Relation operators
	EQUALS          // '==';
	NOTEQUALS       // '!=';
	LESS            // '<';
	LESSOREQUALS    //'<=';
	GREATER         // '>';
	GREATEROREQUALS // '>=';

	// 位运算
	BITOR  // '|';
	LSHIFT // '<<';
	RSHIFT // '>>';
	BITAND // '&';

	// 赋值和运算
	ASSIGN     // "="
	PLUS       // "+"
	MINUS      // "-"
	ASTERISK   //  "*"
	SLASH      // "/"
	MOD        // "%"
	PLUSPLUS   // ++
	MINUSMINUS // --

	//
	NUMBER
	STRING

	//
	ILLEGAL
	EOF
)

type TokenType int

var TokenTypeLiteral = map[TokenType]string{
	BREAK:      "break",
	FUNC:       "function",
	STRUCT:     "struct",
	ELSE:       "else",
	GOTO:       "goto",
	PACKAGE:    "package",
	SWITCH:     "switch",
	IF:         "if",
	TYPE:       "type",
	CONTINUE:   "continue",
	FOR:        "for",
	IMPORT:     "import",
	RETURN:     "return",
	VAR:        "var",
	NILL:       "nil",
	MAP:        "map",
	LIST:       "list",
	TRUE:       "true",
	FALSE:      "false",
	IDENTIFIER: "IDENTIFIER",
	// 标点
	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",
	COMMA:    ",",
	SEMI:     ";",
	COLON:    ":",
	DOT:      ".",

	// Logical
	LOGICALOR:  "||",
	LOGICALAND: "&&",

	BANG: "!",

	// Relation operators
	EQUALS:          "==",
	NOTEQUALS:       "!=",
	LESS:            "<",
	LESSOREQUALS:    "<=",
	GREATER:         ">",
	GREATEROREQUALS: ">=",

	// 位运算
	BITOR:  "|",
	LSHIFT: "<<",
	RSHIFT: ">>",
	BITAND: "&",

	// 赋值和运算
	ASSIGN:     "=",
	PLUS:       "+",
	MINUS:      "-",
	ASTERISK:   "*",
	SLASH:      "/",
	MOD:        "%",
	PLUSPLUS:   "++",
	MINUSMINUS: "--",

	NUMBER:  "NUMBER",
	STRING:  "STRING",
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
}

var LiteralTokenType = map[string]TokenType{
	"break":    BREAK,
	"function": FUNC,
	"struct":   STRUCT,
	"else":     ELSE,
	"goto":     GOTO,
	"package":  PACKAGE,
	"switch":   SWITCH,
	"if":       IF,
	"type":     TYPE,
	"continue": CONTINUE,
	"for":      FOR,
	"import":   IMPORT,
	"return":   RETURN,
	"var":      VAR,
	"nil":      NILL,
	"map":      MAP,
	"list":     LIST,
	"true":     TRUE,
	"false":    FALSE,

	// 标点
	"(": LPAREN,
	")": RPAREN,
	"{": LBRACE,
	"}": RBRACE,
	"[": LBRACKET,
	"]": RBRACKET,
	",": COMMA,
	";": SEMI,
	":": COLON,
	".": DOT,

	// Logical
	"||": LOGICALOR,
	"&&": LOGICALAND,

	"!": BANG,

	// Relation operators
	"==": EQUALS,
	"!=": NOTEQUALS,
	"<":  LESS,
	"<=": LESSOREQUALS,
	">":  GREATER,
	">=": GREATEROREQUALS,

	// 位运算
	"|":  BITOR,
	"<<": LSHIFT,
	">>": RSHIFT,
	"&":  BITAND,

	// 赋值和运算
	"=":  ASSIGN,
	"+":  PLUS,
	"-":  MINUS,
	"*":  ASTERISK,
	"/":  SLASH,
	"%":  MOD,
	"++": PLUSPLUS,
	"--": MINUSMINUS,
}

type Token struct {
	Type     TokenType
	Literal  string
	Line     int
	Position int
}
