package main

import (
	"fmt"
	"hse_project_introductory_task/av"
)

const apiKey = "ZTQTVX829784GT8I"

func main() {
	client := av.NewClient(apiKey)

	SearchSymbolExample(&client)
	ExchangeRateExample(&client)
	StockQuoteExample(&client)
}

func SearchSymbolExample(client *av.Client) {
	res, err := client.SearchSymbol("microsoft")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func ExchangeRateExample(client *av.Client) {
	res, err := client.ExchangeRate("BTC", "CNY")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
}

func StockQuoteExample(client *av.Client) {
	res, err := client.StockQuote("AAPL")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
}
