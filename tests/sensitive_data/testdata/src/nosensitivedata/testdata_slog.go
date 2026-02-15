package nosensitivedata

import "log/slog"

func TestNoSensitiveData() {
	logger := slog.Default()

	// OK - no sensitive data, lowercase, english, no special chars
	logger.Info("server started")
	logger.Info("user logged in successfully")
	logger.Info("request processed")
	logger.Debug("cache hit")
	logger.Info("validation passed")
	
	// Want errors - contains word "password"
	logger.Info("user password is valid") // want "log-message should not contain sensitive data: user password is valid"
	logger.Error("invalid password provided") // want "log-message should not contain sensitive data: invalid password provided"
	logger.Debug("checking passwd") // want "log-message should not contain sensitive data: checking passwd"
	logger.Info("pwd validation failed") // want "log-message should not contain sensitive data: pwd validation failed"
	
	// Want errors - contains word "token"
	logger.Info("user token is valid") // want "log-message should not contain sensitive data: user token is valid"
	logger.Error("invalid token provided") // want "log-message should not contain sensitive data: invalid token provided"
	logger.Debug("refreshing token data") // want "log-message should not contain sensitive data: refreshing token data"
	
	// Want errors - contains "api_key" or "apikey"
	logger.Info("checking apikey") // want "log-message should not contain sensitive data: checking apikey"
	logger.Error("invalid apikey") // want "log-message should not contain sensitive data: invalid apikey"
	
	// Want errors - contains "secret"
	logger.Info("loading secret configuration") // want "log-message should not contain sensitive data: loading secret configuration"
	logger.Debug("secret validation passed") // want "log-message should not contain sensitive data: secret validation passed"
	
	// Want errors - contains "credit_card" or "creditcard"
	logger.Info("processing creditcard payment") // want "log-message should not contain sensitive data: processing creditcard payment"
	logger.Error("invalid creditcard number") // want "log-message should not contain sensitive data: invalid creditcard number"
	
	// Want errors - contains "bearer"
	logger.Info("bearer token received") // want "log-message should not contain sensitive data: bearer token received"
	
	// Want errors - contains "authorization"  
	logger.Debug("authorization header parsed") // want "log-message should not contain sensitive data: authorization header parsed"
	
	// Want errors - case insensitivity
	logger.Info("checking password") // want "log-message should not contain sensitive data: checking password"
	logger.Info("loading secret") // want "log-message should not contain sensitive data: loading secret"

	msg := "password is 123"
	logger.Info(msg) // Should not report - not a string literal
}
