# Линтер для проверки лог-записей
**log-linter** — статический анализатор для Go, который в виде плагина для golangci-lint проверяет лог-записи на корректное форматирование и безопасность.

## Поддерживаемые логгеры

* `log`
* `log/slog`
* `go.uber.org/zap`

---

## Проверяемые правила

### 1. Сообщение должно начинаться со строчной буквы

❌ Неправильно

```go
log.Info("Starting server on port 8080")
slog.Error("Failed to connect to database")
```

✅ Правильно

```go
log.Info("starting server on port 8080")
slog.Error("failed to connect to database")
```

### 2. Сообщения должны быть на английском языке

❌ Неправильно

```go
log.Info("запуск сервера")
log.Error("ошибка подключения к базе данных")
```

✅ Правильно

```go
log.Info("starting server")
log.Error("failed to connect to database")
```

### 3. Сообщения не должны содержать спецсимволы и эмодзи

❌ Неправильно

```go
log.Info("server started! 🚀")
log.Error("connection failed!!!")
log.Warn("warning: something went wrong...")
```

✅ Правильно

```go
log.Info("server started")
log.Error("connection failed")
log.Warn("something went wrong")
```

### 4. Сообщения не должны содержать чувствительные данные

Линтер проверяет использование потенциально чувствительных переменных, например:

* `password`
* `token`
* `apiKey`

❌ Неправильно

```go
log.Info("user password: " + password)
log.Debug("api_key=" + apiKey)
log.Info("token: " + token)
```

✅ Правильно

```go
log.Info("user authenticated successfully")
log.Debug("api request completed")
log.Info("token validated")
```

---

## Установка и сборка

Требования:

* **Go 1.22+**
* **golangci-lint 1.52+**
* **CGO_ENABLED=1**

Клонирование репозитория:

```bash
git clone https://github.com/montonyyy/log-linter.git
cd log-linter
```

Сборка плагина:
```bash
CGO_ENABLED=1 go build -buildmode=plugin -o loglint.so ./plugin
```
После сборки в корне появится файл `loglint.so` — это сам плагин.

---

## Использование с golangci-lint

1. Скопируйте плагин в свой проект

```bash
cp /path/to/log-linter/loglint.so /path/to/your/project/
cd /path/to/your/project
```

2. В корне вашего проекта создайте или отредактируйте файл `.golangci.yml`:

```yaml
linters-settings:
  custom:
    loglint:
      path: ./loglint.so
      description: "log-linter"
linters:
  enable:
    - loglint
  disable-all: true
```

3. Запустите проверку

```bash
golangci-lint run ./...
```

## Примеры использования

Базовые ошибки:

```go
package main

import "log"

func main() {
    log.Println("Starting server")  // contains capital letter
    log.Println("запуск сервера")    // contains not an english letter
    log.Println("hello!")            // contains symbol letter
}
```

Чувствительные данные:

```go
package main

import "log/slog"

func main() {
    password := "qwerty1337"
    token := "ABCD:EFGH"
    user := "Ivan"

    slog.Info("password " + password)  // contains sensitive data
    slog.Info("token " + token)        // contains sensitive data
    slog.Info("user " + user)          // OK
}
```

Бинарные выражения также работают:

```go
package main

import "log/slog"

func main() {
    slog.Info("Hello " + "мир" + "!")  // ❌ contains capital letter, contains not an enlgish letter, contains symbol letter, 
}
```

---

## Тестирование

Запуск unit-тестов

```bash
go test ./...
```

Ручная проверка на тестовых файлах

```bash
golangci-lint run --disable-all --enable=loglint ./testdata/...
```
