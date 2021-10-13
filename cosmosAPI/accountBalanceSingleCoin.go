package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountBalanceDetail struct {
	Balance struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balance"`
}

func main() {
	// Replace Account address and denom values in function call template below
	accountBalance("akash1ss3ty253h6yun0a0fly8s0prcx34x4q2qewpkk", "uakt")
}

func accountBalance(account, denom string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/bank/v1beta1/balances/" + account + "/" + denom

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

	AccountBalanceDetail1 := AccountBalanceDetail{}

	json.Unmarshal(body, &AccountBalanceDetail1)

	fmt.Println(AccountBalanceDetail1)

}
