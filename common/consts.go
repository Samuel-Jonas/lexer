package common

const (
	LETTERS_LOWER  = "abcdefghiklmnopqrstuvwxyz"
	LETTERS_UPPER  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS         = "0123456789"
	LETTERS        = LETTERS_LOWER + LETTERS_UPPER
	LETTERS_DIGITS = DIGITS + LETTERS
	INT            = "INT"     // integer data type
	FLOAT          = "FLOAT"   // float data ype
	STRING         = "STRING"  // string data type
	IDENTIFIER     = "ID"      // identifier or variable type
	KEYWORD        = "KEYWORD" // reserved word type
	PLUS           = "PLUS"    // (+) addition operation symbol
	MINUS          = "MINUS"   // (-) subtraction operation symbol
	TIMES          = "TIMES"   // (*) multiplication operation symbol
	DIV            = "DIV"     // (/) division operation symbol
	POW            = "POW"     // (^) exponentiation opeation symbol
	EQ             = "EQ"      // (==) equality comparison symbol
	LPAR           = "LPAR"    // ( left parenthesis
	RPAR           = "RPAR"    // ) right parenthesis
	LSQUARE        = "LSQUARE" // ([) left square bracket
	RSQUARE        = "RSQUARE" // (]) right square bracket
	EQUAL          = "EQUAL"   // (==) equal comparison symbol
	NE             = "NE"      // (!=) not equal comparison symbol
	LT             = "LT"      // (<) less then comparison symbol
	GT             = "GT"      // (>) greater then comparison symbol
	LTE            = "LTE"     // (<=) less then or equal to comparison symbol
	GTE            = "GTE"     // (>=) greater then or equal to comparison symbol
	COMMA          = "COMMA"   // (,) comma symbol
	ARROW          = "ARROW"   // (->) arrow symbol
	NEWLINE        = "NEWLINE" // (/n) new line
	EOF            = "EOF"     // end of file
	VAR            = "var"     // keyword var
	AND            = "and"     // keyword and
	OR             = "or"      // keyword or
	NOT            = "not"     // keyword not
	WHILE          = "while"   // keyword while
	DO             = "do"      // keyword do
	END            = "end"     // keyword end
	FALSE          = "false"   // keyword false
	TRUE           = "true"    // keyword true
	NULL           = "null"    // keyword null
)
