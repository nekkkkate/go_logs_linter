package lowercasefirst

import "log/slog"

func TestLowerCaseFirst() {
	logger := slog.Default()

	logger.Info("Starting server")    // want "first letter should be lowercase: Starting server"
	logger.Debug("debug message")     // OK - already lowercase
	logger.Warn("warning")            // OK - already lowercase
	logger.Error("Failed to connect") // want "first letter should be lowercase: Failed to connect"

	msg := "Error occurred"
	logger.Info(msg) // Should not report - not a string literal

	logger.Info("") // OK - empty string
}
