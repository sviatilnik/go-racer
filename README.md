# Go Racer

Параллельный «рейсер» URL-адресов на Go. Утилита принимает список сайтов и возвращает самый быстрый по времени ответа, используя горутины и `context` для отмены операций.

## Функционал

- Запуск HTTP-запросов к списку URL параллельно (по одному запросу на URL)
- Определение самого быстрого ответа (первый успешный)
- Отмена остальных запросов при получении первого успешного результата или при таймауте
- Проверка наличия интернета перед стартом
- Валидация URL-адресов
- Флаг `-timeout` для ограничения общего времени гонки (по умолчанию 5s)

## Сборка

```bash
go build -o racer ./cmd
```

Или с указанием модуля:

```bash
go build -o racer github.com/sviatilnik/go-racer/cmd
```

## Запуск

**Базовый запуск (таймаут по умолчанию 5s):**

```bash
./racer https://google.com https://ya.ru https://bing.com
```

**С указанием таймаута:**

```bash
./racer -timeout=3s https://google.com https://ya.ru https://bing.com
```

**Запуск через go run:**

```bash
go run ./cmd -timeout=2s https://google.com https://ya.ru
```

**Проверка с детектором гонок:**

```bash
go run -race ./cmd -timeout=2s https://google.com https://ya.ru
```

## Пример вывода

**Успешный результат:**

```text
🏁 Стартуем гонку для 3 URL с таймаутом 5.000000s
🚀 Самый быстрый: https://ya.ru (142 ms)

Остальные результаты: 
	✓ https://google.com (156 ms)
	✓ https://bing.com (203 ms)
```

**Все URL недоступны:**

```text
🏁 Стартуем гонку для 3 URL с таймаутом 2.000000s

❌ Ни один URL не ответил успешно

Ошибки: 
	 ✗ https://несуществующий-сайт.рф : dial tcp: lookup failed
	 ✗ https://google.com : контекст отменен
	 ✗ https://ya.ru : контекст отменен
```

## Структура проекта

```text
go_racer/
├── cmd/
│   └── racer.go          # Точка входа
├── internal/
│   ├── config/          # Парсинг флагов (-timeout)
│   │   ├── config.go    # Интерфейс Config
│   │   └── flag.go      # FlagConfig
│   ├── race/            # Логика гонки
│   │   ├── race.go      # Race, worker, оркестрация
│   │   └── result.go    # Форматирование вывода
│   └── utils/
│       └── urlvalidator.go  # Валидация URL, проверка сети
├── go.mod
└── README.md
```

## Требования

- Go 1.25+ (используется `sync.WaitGroup.Go()`)
