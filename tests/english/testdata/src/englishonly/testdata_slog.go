package englishonly

import "log/slog"

func TestEnglishOnly() {
	logger := slog.Default()

	// OK - English only, lowercase, no special chars
	logger.Info("starting server")
	logger.Debug("processing request")
	logger.Warn("memory usage high")
	
	// Want errors - contains non-English characters
	logger.Info("привет мир")      // want "log-message should be in English only: привет мир"
	logger.Info("hello 世界")        // want "log-message should be in English only: hello 世界"
	logger.Info("café")            // want "log-message should be in English only: café"
	logger.Error("ошибка подключения") // want "log-message should be in English only: ошибка подключения"
	
	// OK - numbers are allowed
	logger.Info("123")
	logger.Info("server started on port 8080")

	msg := "привет"
	logger.Info(msg) // Should not report - not a string literal
}
