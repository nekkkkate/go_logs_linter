# go_logs_linter
Линтер для Go, совместимый с golangci-lint, который анализирует лог-записи в коде и проверяет их соответствие установленным правилами


# Инструкция по установке и запуску линтера:

## Использование кастомного линтера в командной строке

1) Собираем линтер:

    ```bash
    go build -o path/to/compiled/linter ./cmd/logslint 
    ```

2) Запускаем линтер для всех файлов проекта:

    ```bash
    go run path/to/compiled/linter ./...
    ```

3) Запускаем линтер для определенного пакета:

    ```bash
    go run path/to/compiled/linter ./path/to/package
    ```

## Запуск линтера с помощью golangci-lint

### Если не установлен golangci-lint, необходимо это исправить:

1) Установка golangci-lint

    ```bash
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
    ```

2) Очень важно проверить, что кастомный линтер собирается с той же версией, что и golangci-lint - в противном случае исправить это (так как плагин просто не заработает)

   Версия golangci-lint:

    ```bash
    golangci-lint --version
    ```

   Версия Go:

   ```bash
    go version
    ```
   
   Если версии не совпадают, нужно пересобрать линтер с той же версией Go, что и golangci-lint
   
### Основная инструкция

1) Билдим линтер как плагин (то есть нам нужно чтобы плагин был скомпилирован как shared object (.so) файлы)

    ```bash
    go build -buildmode=plugin -o loglinter.so ./cmd/loglinter-plugin
    ```

2) Запускаем с кастомным линтером

    ```bash
    golangci-lint run --config .golangci.yml path/to/file
    ```
   Пример запуска линтера:
   ![img.png](img.png)

# Полезные материалы

    - https://www.youtube.com/live/BkTJysDTlb0?si=2TbEHMcd1DfyzFYr (чтобы разобраться, как именно устроен пакет golang.org/x/tools/go/analysis, а также как подключить кастомный линтер к golangci-lint)
    - https://habr.com/ru/companies/ostrovok/articles/908768/#habracut (очень хорошая статься, чтобы разобраться, как подходить к разработке своего линтера)
    - https://youtu.be/Q9ymKAVgQ_8?si=uvCPx7LBYGehgYKB (для общего понимания, что из себя представляет golangci-lint)
    - https://habr.com/ru/articles/457970/ (более подробно о линтерах в Go, возможностях golangci-lint)
