package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type UnbondingDetails struct {
	Unbond struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Entries          []struct {
			CreationHeight string    `json:"creation_height"`
			CompletionTime time.Time `json:"completion_time"`
			InitialBalance string    `json:"initial_balance"`
			Balance        string    `json:"balance"`
		} `json:"entries"`
	} `json:"unbond"`
}

func main() {
	// Replace Validator and Delegator Account values in function call template below
	validatorUnbonding("akashvaloper1qvsus5qg8yhre7k2c78xkkw4nvqqgev7zw5wzs", "akash1qytlz4pam956hese8x77fwk5kamd9sxk99j0fz")
}

func validatorUnbonding(validatorId, delegatorId string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/staking/v1beta1/validators/" + validatorId + "/delegations/" + delegatorId + "/unbonding_delegation"

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

	UnbondingDetails1 := UnbondingDetails{}

	json.Unmarshal(body, &UnbondingDetails1)

	fmt.Println(UnbondingDetails1)

}
