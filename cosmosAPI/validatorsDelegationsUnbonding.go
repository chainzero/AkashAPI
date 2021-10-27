package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type UnbondingDetails struct {
	UnbondingResponses []struct {
		DelegatorAddress string `json:"delegator_address"`
		ValidatorAddress string `json:"validator_address"`
		Entries          []struct {
			CreationHeight string    `json:"creation_height"`
			CompletionTime time.Time `json:"completion_time"`
			InitialBalance string    `json:"initial_balance"`
			Balance        string    `json:"balance"`
		} `json:"entries"`
	} `json:"unbonding_responses"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}

func main() {
	// Replace Validator Account value in function call template below
	validatorsDelegations("akashvaloper1qvsus5qg8yhre7k2c78xkkw4nvqqgev7zw5wzs")
}

func validatorsDelegations(id string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/staking/v1beta1/validators/" + id + "/unbonding_delegations"

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
