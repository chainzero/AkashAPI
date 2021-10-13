package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ParamDetails struct {
	Params struct {
		MaxMemoCharacters      string `json:"max_memo_characters"`
		TxSigLimit             string `json:"tx_sig_limit"`
		TxSizeCostPerByte      string `json:"tx_size_cost_per_byte"`
		SigVerifyCostEd25519   string `json:"sig_verify_cost_ed25519"`
		SigVerifyCostSecp256K1 string `json:"sig_verify_cost_secp256k1"`
	} `json:"params"`
}

func main() {
	parameters()
}

func parameters() {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/auth/v1beta1/params"

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

	fmt.Println(string(body))

	ParamDetails1 := ParamDetails{}

	json.Unmarshal(body, &ParamDetails1)

	fmt.Println(ParamDetails1)

}
