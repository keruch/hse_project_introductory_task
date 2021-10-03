package av

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	symSearch = "SYMBOL_SEARCH"
)

type SearchSymResult struct {
	Symbol      string `json:"1. symbol"`
	Name        string `json:"2. name"`
	Type        string `json:"3. type"`
	Region      string `json:"4. region"`
	MarketOpen  string `json:"5. marketOpen"`
	MarketClose string `json:"6. marketClose"`
	Timezone    string `json:"7. timezone"`
	Currency    string `json:"8. currency"`
	MatchScore  string `json:"9. matchScore"`
}

func (c *Client) SearchSym(keyword string) ([]SearchSymResult, error) {
	endpoint := c.buildRequestString(requestParams{
		queryFunction: symSearch,
		queryKeyword:  keyword,
	})
	response, err := c.conn.Request(endpoint)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	type searchSymData struct {
		BestMatches []SearchSymResult `json:"bestMatches"`
	}
	var res searchSymData

	err = json.Unmarshal(bodyBytes, &res)
	if err != nil {
		return nil, err
	}

	return res.BestMatches, nil
}
