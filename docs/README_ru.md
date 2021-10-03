# GO API для Alpha Vantage
Alpha Vantage - инструмент, который предоставляет бесплатные API для работы с биржевыми данными в JSON и CSV форматах.

Данный API является Golang оберткой Alpha Vantage. С его помощью можно производить поиск на бирже по ключевым словам,
сравнивать стоимости котировок и просматривать информацию по заданным тикерам. 

Для работы с данным API пользователю
следует получить ключ на сайте Alpha Vantage. Это можно сделать по [ссылке](https://www.alphavantage.co/support/#api-key).

## Структура API

Дерево проекта:

```
root
  ├ .github
  |  ├ workflows
  |  |  └ golangci-lint.yaml - конфигурация ci линтера 
  ├ av                       - пакет с API
  |  ├ connection.go    
  |  ├ exchange_rate.go 
  |  ├ search_symbol.go 
  |  └ stock_quote.go   
  ├ docs                     - информация про репозиторий
  ├ .golangci.yaml           - конфигурация локального линтера
  └ main.go                  - примеры использования
```

В `connection.go` находится основной функционал API, который общается с сайтом и получает от него информацию. Особого
внимания достойна структура `Client`, по средствам которой производится все общение с API из внешнего кода. Она содержит
интерфейс `Connection` и API ключ пользователя, который указывается во время создания структуры при помощи `NewClient`.
Интерфейс `Connection` вводит функцию `Request`, которая реализует общение с сайтом с помощью GET запросов.
Структура `avConnection` реализует данный интерфейс. 
Запросы для функции `Request` создаются в `buildRequestString`.
Также в `connection.go` хранятся некоторые общие константы для составления запросов.

Файлы `exchange_rate.go`, `search_symbol.go` и `stock_quote.go` реализуют функции API.

## Функции API
Все функции в API формируют запросы с помощью `buildRequestString`, которая принимает хеш таблицу параметров и составляет
запрос правильного вида. Возвращаемым значением каждой функции является структура с полями, соответствующими ответу. 

### ExchangeRate
Файл `exchange_rate.go` содержит функцию `ExchangeRate`, которая сравнивает цены котировок.
Ее сигнатура: 
```
func (c *Client) ExchangeRate(fromCurrency, toCurrency string) (ExchangeRateResult, error)
```
`fromCurrency` - из чего переводим, `toCurrency` - во что переводим. 

Примеры запроса и ответа:

[https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo](https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo)

```javascript
{
  "Realtime Currency Exchange Rate": {
    "1. From_Currency Code": "BTC",
    "2. From_Currency Name": "Bitcoin",
    "3. To_Currency Code": "CNY",
    "4. To_Currency Name": "Chinese Yuan",
    "5. Exchange Rate": "27513.09000000",
    "6. Last Refreshed": "2019-02-20 20:07:15",
    "7. Time Zone": "UTC"
  }
}
```

### SearchSymbol
Файл `search_symbol.go` содержит функцию `SearchSymbol`, которая производит поиск тикеров по ключевому слову.
Ее сигнатура:
```
func (c *Client) SearchSymbol(keyword string) ([]SearchSymResult, error)
```
`keyword` - ключевое слово.

Примеры запроса и ответа:

[https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=Micro&apikey=demo](https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=Micro&apikey=demo)

```javascript
{
    "bestMatches": [
        {
            "1. symbol": "AMD",
            "2. name": "Advanced Micro Devices Inc.",
            "3. type": "Equity",
            "4. region": "United States",
            "5. marketOpen": "09:30",
            "6. marketClose": "16:00",
            "7. timezone": "UTC-05",
            "8. currency": "USD",
            "9. matchScore": "0.5000"
        },
        {
            "1. symbol": "MU",
            "2. name": "Micron Technology Inc.",
            "3. type": "Equity",
            "4. region": "United States",
            "5. marketOpen": "09:30",
            "6. marketClose": "16:00",
            "7. timezone": "UTC-05",
            "8. currency": "USD",
            "9. matchScore": "0.4545"
        },
        {
            "1. symbol": "MSFT",
            "2. name": "Microsoft Corporation",
            "3. type": "Equity",
            "4. region": "United States",
            "5. marketOpen": "09:30",
            "6. marketClose": "16:00",
            "7. timezone": "UTC-05",
            "8. currency": "USD",
            "9. matchScore": "0.4444"
        },
        { ... },
        { ... },
        { ... }
    ]
}
```

### StockQuote
Файл `stock_quote.go` содержит функцию `StockQuote`, которая выдает информацию (японская свеча, объемы и тд) по 
заданному тикеру.
Ее сигнатура:
```
func (c *Client) StockQuote(symbol string) (StockQuoteResult, error)
```
`symbol` - тикер.

Примеры запроса и ответа:

[https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo](https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=MSFT&apikey=demo)

```javascript
{
    "Global Quote": {
        "01. symbol": "MSFT",
            "02. open": "107.8600",
            "03. high": "107.9400",
            "04. low": "106.2950",
            "05. price": "106.7900",
            "06. volume": "13085190",
            "07. latest trading day": "2019-02-20",
            "08. previous close": "107.7100",
            "09. change": "-0.9200",
            "10. change percent": "-0.8541%"
    }
}
```
