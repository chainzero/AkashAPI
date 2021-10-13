package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinSupplyDetail struct {
	Amount struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"amount"`
}

func main() {
	// Replace Account denom value in function call template below
	coinSupply("uakt")
}

func coinSupply(denom string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/bank/v1beta1/supply/" + denom

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

	CoinSupplyDetail1 := CoinSupplyDetail{}

	json.Unmarshal(body, &CoinSupplyDetail1)

	fmt.Println(CoinSupplyDetail1)

}
