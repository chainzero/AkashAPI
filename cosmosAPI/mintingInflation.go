package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type InflationDetail struct {
	Inflation string `json:"inflation"`
}

func main() {
	mintingInflation()
}

func mintingInflation() {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/mint/v1beta1/inflation"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	InflationDetail1 := InflationDetail{}

	json.Unmarshal(body, &InflationDetail1)

	fmt.Println(InflationDetail1)

}
