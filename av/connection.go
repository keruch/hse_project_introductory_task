package av

import (
	"net/http"
	"net/url"
	"time"
)

const (
	hostName    = "www.alphavantage.co"
	httpsScheme = "https"

	queryName       = "query"

	queryAPI        = "apikey"
	queryAdjusted   = "adjusted"
	queryOutputSize = "outputsize"
	queryDataType   = "datatype"

	queryTrue       = "true"
	queryCompact    = "compact"
	queryJSON       = "json"
	queryFunction   = "function"
)

type avConnection struct {
	client *http.Client
	host   string
}

type Client struct {
	conn   Connection
	apiKey string
}

type Connection interface {
	Request(*url.URL) (*http.Response, error)
}

func NewConnection() Connection {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return &avConnection{
		client: client,
		host:   hostName,
	}
}

func NewClient(apiKey string) Client {
	return Client{
		conn:   NewConnection(),
		apiKey: apiKey,
	}
}

func (conn *avConnection) Request(endpoint *url.URL) (*http.Response, error) {
	endpoint.Scheme = httpsScheme
	endpoint.Host = conn.host
	uri := endpoint.String()
	return conn.client.Get(uri)
}

type requestParams map[string]string

func (c *Client) buildRequestString(params requestParams) *url.URL {
	endpoint := &url.URL{}
	endpoint.Path = queryName

	query := endpoint.Query()
	query.Set(queryAPI, c.apiKey)
	query.Set(queryAdjusted, queryTrue)
	query.Set(queryOutputSize, queryCompact)
	query.Set(queryDataType, queryJSON)

	for key, val := range params {
		query.Set(key, val)
	}

	endpoint.RawQuery = query.Encode()
	return endpoint
}
