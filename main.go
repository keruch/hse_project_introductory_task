package main

import (
	"fmt"
	"hse_project_introductory_task/av"
)

const apiKey = "ZTQTVX829784GT8I"

func main() {
	client := av.NewClient(apiKey)
	res, err := client.SearchSym("microsoft")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, val := range res {
		fmt.Printf("%+v\n", val)
	}
}
