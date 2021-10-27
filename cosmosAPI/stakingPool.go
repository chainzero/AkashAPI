package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PoolDetails struct {
	Pool struct {
		NotBondedTokens string `json:"not_bonded_tokens"`
		BondedTokens    string `json:"bonded_tokens"`
	} `json:"pool"`
}

func main() {
	stakingPoolQuery()
}

func stakingPoolQuery() {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/staking/v1beta1/pool"

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

	PoolDetails1 := PoolDetails{}

	json.Unmarshal(body, &PoolDetails1)

	fmt.Println(PoolDetails1)

}
