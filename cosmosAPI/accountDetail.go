package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AccountDetail struct {
	Account struct {
		Type    string `json:"@type"`
		Address string `json:"address"`
		PubKey  struct {
			Type string `json:"@type"`
			Key  string `json:"key"`
		} `json:"pub_key"`
		AccountNumber string `json:"account_number"`
		Sequence      string `json:"sequence"`
	} `json:"account"`
}

func main() {
	// Replace Account address value in function call template below
	account("akash1ss3ty253h6yun0a0fly8s0prcx34x4q2qewpkk")
}

func account(account string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/auth/v1beta1/accounts/" + account

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

	AccountDetail1 := AccountDetail{}

	json.Unmarshal(body, &AccountDetail1)

	fmt.Println(AccountDetail1)

}
