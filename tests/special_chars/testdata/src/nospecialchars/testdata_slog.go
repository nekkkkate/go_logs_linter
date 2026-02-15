package nospecialchars

import "log/slog"

func TestNoSpecialChars() {
	logger := slog.Default()

	// OK - no special chars, lowercase, english
	logger.Info("starting server")
	logger.Info("connection established")
	logger.Info("value is 42")
	
	// Want errors - contains special characters (punctuation)
	logger.Info("hello, world")    // want "log-message should not contain special characters: hello, world"
	logger.Info("error: timeout")   // want "log-message should not contain special characters: error: timeout"
	logger.Info("file not found!")  // want "log-message should not contain special characters: file not found!"
	logger.Info("user@example.com") // want "log-message should not contain special characters: user@example.com"
	logger.Info("request-123") // want "log-message should not contain special characters: request-123"
	
	// Want errors - contains emoji
	logger.Info("server started 🚀") // want "log-message should not contain special characters: server started 🚀"
	logger.Warn("warning message ⚠️") // want "log-message should not contain special characters: warning message ⚠️"

	msg := "hello!"
	logger.Info(msg) // Should not report - not a string literal
}
