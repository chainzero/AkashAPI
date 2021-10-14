package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DepositByAddressDetail struct {
	Deposit struct {
		ProposalID string `json:"proposal_id"`
		Depositor  string `json:"depositor"`
		Amount     []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"amount"`
	} `json:"deposit"`
}

func main() {
	// Replace Proposal ID and depositor address values in function call template below
	proposalDepositsByAddress("13", "akash1uu4e4lt460h6etczv4pfngzqe436vekym3lh0x")
}

func proposalDepositsByAddress(id, address string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/gov/v1beta1/proposals/" + id + "/deposits/" + address

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

	DepositByAddressDetail1 := DepositByAddressDetail{}

	json.Unmarshal(body, &DepositByAddressDetail1)

	fmt.Println(DepositByAddressDetail1)

}
