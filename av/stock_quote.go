package av

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	stockQuote = "GLOBAL_QUOTE"

	querySymbol = "symbol"
)

type StockQuoteResult struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	Change           string `json:"09. change"`
	ChangePercent    string `json:"10. change percent"`
}

func (c *Client) StockQuote(symbol string) (StockQuoteResult, error) {
	endpoint := c.buildRequestString(requestParams{
		queryFunction: stockQuote,
		querySymbol:   symbol,
	})
	response, err := c.conn.Request(endpoint)
	if err != nil || response.StatusCode != http.StatusOK {
		return StockQuoteResult{}, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return StockQuoteResult{}, err
	}

	var res struct {
		StockQuote StockQuoteResult `json:"Global Quote"`
	}
	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return StockQuoteResult{}, err
	}

	return res.StockQuote, nil
}
