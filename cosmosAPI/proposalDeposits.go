package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DepositDetails struct {
	Deposits []struct {
		ProposalID string        `json:"proposal_id"`
		Depositor  string        `json:"depositor"`
		Amount     []interface{} `json:"amount"`
	} `json:"deposits"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func main() {
	// Replace ID value in function call template below
	proposalDeposits("13")
}

func proposalDeposits(id string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/gov/v1beta1/proposals/" + id + "/deposits"

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

	// ProposalDetail1 := ProposalDetail{}
	DepositDetails1 := DepositDetails{}

	json.Unmarshal(body, &DepositDetails1)

	fmt.Println(DepositDetails1)

}
