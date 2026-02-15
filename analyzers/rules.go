package analyzers

import (
	"go/token"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func lowerCaseFirst(pass *analysis.Pass, msg string, pos token.Pos) {
	msgRune := []rune(msg)

	if len(msgRune) > 0 && unicode.IsLetter(msgRune[0]) && unicode.IsUpper(msgRune[0]) {
		pass.Reportf(pos, "first letter should be lowercase: %s", msg)
	}
}

func englishOnly(pass *analysis.Pass, msg string, pos token.Pos) {
	msgRune := []rune(msg)
	for _, r := range msgRune {
		if unicode.IsLetter(r) && (!((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z'))) {
			pass.Reportf(pos, "log-message should be in English only: %s", msg)
			break
		}
	}
}

func noSpecialChars(pass *analysis.Pass, msg string, pos token.Pos) {
	msgRune := []rune(msg)
	for _, r := range msgRune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsSpace(r) {
			pass.Reportf(pos, "log-message should not contain special characters: %s", msg)
			break
		}
	}
}

func noSensitiveData(pass *analysis.Pass, msg string, pos token.Pos) {
	keywords := []string{
		"password", "passwd", "pwd", "secret", "token", "auth", "api_key",
		"apikey", "credit_card", "creditcard", "private_key", "privatekey",
		"access_key", "accesskey", "bearer", "authorization",
	}
	msgLower := strings.ToLower(msg)

	for _, keyword := range keywords {
		if strings.Contains(msgLower, keyword) {
			pass.Reportf(pos, "log-message should not contain sensitive data: %s", msg)
			break
		}
	}
}
