# Тесткейсы для линтера логов

Тестовые файлы содержат намеренные нарушения правил для проверки линтера. Каждый файл соответствует одному логгеру

## Правило 1: Сообщения должны начинаться со строчной буквы

| Файл | Функция | Нарушение |
|------|---------|-----------|
| slog_logger.go | slogHandleHTTP | "API Request" |
| slog_logger.go | slogHandleHTTP | "Request token" |
| slog_logger.go | slogHandleHTTP | "API Key used" |
| slog_logger.go | slogAuthenticate | "Token is empty" |
| slog_logger.go | slogAuthenticate | "User authentication" |
| slog_logger.go | slogAuthenticate | "Пароль пользователя" |
| slog_logger.go | slogAuthenticate | "Password validation" |
| slog_logger.go | slogAuthenticate | "Validating password" |
| slog_logger.go | slogAuthenticate | "User john_doe authenticated" |
| slog_logger.go | slogStartServer | "Starting application" |
| slog_logger.go | slogStartServer | "Port not specified" |
| slog_logger.go | slogStartServer | "Сервер успешно" |
| slog_logger.go | slogStartServer | "Application ready" |
| slog_logger.go | slogInitializeDB | "Connecting to database" |
| slog_logger.go | slogInitializeDB | "Using password" |
| slog_logger.go | slogProcessPayment | "Processing payment" |
| slog_logger.go | slogProcessPayment | "CVV validation" |
| slog_logger.go | slogProcessPayment | "Payment declined" |
| slog_logger.go | slogConfigLoaded | "Config loaded" |
| slog_logger.go | slogConfigLoaded | "Config contains" |
| zap_logger.go | zapHandleHTTP | "API Request" |
| zap_logger.go | zapHandleHTTP | "API Key used" |
| zap_logger.go | zapAuthenticate | "Token is empty" |
| zap_logger.go | zapAuthenticate | "User authentication" |
| zap_logger.go | zapAuthenticate | "Пароль пользователя" |
| zap_logger.go | zapAuthenticate | "Password validation" |
| zap_logger.go | zapAuthenticate | "Validating password" |
| zap_logger.go | zapAuthenticate | "User authenticated" |
| zap_logger.go | zapStartServer | "Starting application" |
| zap_logger.go | zapStartServer | "Port not specified" |
| zap_logger.go | zapStartServer | "Сервер успешно" |
| zap_logger.go | zapStartServer | "Application ready" |
| zap_logger.go | zapInitializeDB | "Connecting to database" |
| zap_logger.go | zapInitializeDB | "Using password" |
| zap_logger.go | zapProcessPayment | "Processing payment" |
| zap_logger.go | zapProcessPayment | "CVV validation" |
| zap_logger.go | zapProcessPayment | "Payment declined" |
| zap_logger.go | zapConfigLoaded | "Config loaded" |
| zap_logger.go | zapMiddleware | "Starting service" |
| zap_logger.go | zapMiddleware | "ошибка при загрузке" |
| zap_logger.go | zapMiddleware | "configuration loaded" |
| zap_logger.go | zapMiddleware | "api_key=" |

## Правило 2: Только английский язык

| Файл | Функция | Нарушение |
|------|---------|-----------|
| slog_logger.go | slogHandleHTTP | "ошибка валидации API ключа" |
| slog_logger.go | slogHandleHTTP | "Внимание" |
| slog_logger.go | slogAuthenticate | "Начинаем аутентификацию" |
| slog_logger.go | slogAuthenticate | "Пароль пользователя" |
| slog_logger.go | slogStartServer | "Сервер успешно запущен" |
| slog_logger.go | slogStartServer | "ошибка подключения к базе данных" |
| zap_logger.go | zapHandleHTTP | "ошибка валидации API ключа" |
| zap_logger.go | zapHandleHTTP | "Внимание" |
| zap_logger.go | zapAuthenticate | "Начинаем аутентификацию" |
| zap_logger.go | zapAuthenticate | "Пароль пользователя" |
| zap_logger.go | zapStartServer | "Сервер успешно запущен" |
| zap_logger.go | zapStartServer | "ошибка подключения к базе данных" |
| zap_logger.go | zapMiddleware | "ошибка при загрузке конфигурации" |

## Правило 3: Без спецсимволов и эмодзи

| Файл | Функция | Нарушение |
|------|---------|-----------|
| slog_logger.go | slogHandleHTTP | 📨, ! |
| slog_logger.go | slogHandleHTTP | !!!, ..., ⏰ |
| slog_logger.go | slogAuthenticate | ..., 🔐 |
| slog_logger.go | slogAuthenticate | !!!, 😱 |
| slog_logger.go | slogAuthenticate | !, 🎉 |
| slog_logger.go | slogAuthenticate | !!! |
| slog_logger.go | slogStartServer | !!!, ..., 🔧 |
| slog_logger.go | slogStartServer | !, 🎉 |
| slog_logger.go | slogStartServer | ✅, 🚀 |
| slog_logger.go | slogProcessPayment | !!!, ..., 💳 |
| zap_logger.go | zapHandleHTTP | 📨, ! |
| zap_logger.go | zapHandleHTTP | !!!, ..., ⏰ |
| zap_logger.go | zapAuthenticate | ..., 🔐 |
| zap_logger.go | zapAuthenticate | !!!, 😱 |
| zap_logger.go | zapAuthenticate | !, 🎉 |
| zap_logger.go | zapAuthenticate | !!! |
| zap_logger.go | zapStartServer | !!!, ..., 🔧 |
| zap_logger.go | zapStartServer | !, 🎉 |
| zap_logger.go | zapStartServer | !, ✅, 🚀 |
| zap_logger.go | zapProcessPayment | !!!, 💳 |
| zap_logger.go | zapMiddleware | !!!, 🎯 |

## Правило 4: Без чувствительных данных

| Файл | Функция | Нарушение |
|------|---------|-----------|
| slog_logger.go | slogHandleHTTP | apiKey в строке |
| slog_logger.go | slogHandleHTTP | userToken в строке |
| slog_logger.go | slogHandleHTTP | "API Key used: " + apiKey |
| slog_logger.go | slogHandleHTTP | token=secret-token-123 |
| slog_logger.go | slogAuthenticate | token в строке |
| slog_logger.go | slogAuthenticate | "Пароль пользователя: " + password |
| slog_logger.go | slogAuthenticate | "Validating password: " + pass |
| slog_logger.go | slogStartServer | api_key=, dbPassword, apiSecret |
| slog_logger.go | slogStartServer | db_password, api_secret |
| slog_logger.go | slogInitializeDB | username, password |
| slog_logger.go | slogProcessPayment | cardNumber, cvv |
| slog_logger.go | slogConfigLoaded | secret_key= |
| zap_logger.go | zapHandleHTTP | apiKey (zap.String), userToken |
| zap_logger.go | zapHandleHTTP | "API Key used: " + apiKey |
| zap_logger.go | zapHandleHTTP | token= + userToken |
| zap_logger.go | zapAuthenticate | token, password, pass |
| zap_logger.go | zapStartServer | api_key=, dbPassword, api_secret |
| zap_logger.go | zapInitializeDB | username, password |
| zap_logger.go | zapProcessPayment | cardNumber, cvv |
| zap_logger.go | zapConfigLoaded | secret_key= |
| zap_logger.go | zapMiddleware | api_key= |
