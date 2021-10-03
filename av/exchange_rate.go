package av

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ExchangeRateResult struct {
	FromCurrencyCode string `json:"1. From_Currency Code"`
	FromCurrencyName string `json:"2. From_Currency Name"`
	ToCurrencyCode   string `json:"3. To_Currency Code"`
	ToCurrencyName   string `json:"4. To_Currency Name"`
	ExchangeRate     string `json:"5. Exchange Rate"`
	LastRefreshed    string `json:"6. Last Refreshed"`
	TimeZone         string `json:"7. Time Zone"`
	BidPrice         string `json:"8. Bid Price"`
	AskPrice         string `json:"9. Ask Price"`
}

const (
	exRate = "CURRENCY_EXCHANGE_RATE"

	queryFromCurrency = "from_currency"
	queryToCurrency   = "to_currency"
)

func (c *Client) ExchangeRate(fromCurrency, toCurrency string) (ExchangeRateResult, error) {
	endpoint := c.buildRequestString(requestParams{
		queryFunction:     exRate,
		queryFromCurrency: fromCurrency,
		queryToCurrency:   toCurrency,
	})
	response, err := c.conn.Request(endpoint)
	if err != nil || response.StatusCode != http.StatusOK {
		return ExchangeRateResult{}, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ExchangeRateResult{}, err
	}

	var res struct {
		ExchangeRate ExchangeRateResult `json:"Realtime Currency Exchange Rate"`
	}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return ExchangeRateResult{}, err
	}

	return res.ExchangeRate, nil
}
