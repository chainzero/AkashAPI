package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DelegatorDetail struct {
	DelegationResponse struct {
		Delegation struct {
			DelegatorAddress string `json:"delegator_address"`
			ValidatorAddress string `json:"validator_address"`
			Shares           string `json:"shares"`
		} `json:"delegation"`
		Balance struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balance"`
	} `json:"delegation_response"`
}

func main() {
	// Replace Validator Account value in function call template below
	validatorsDelegations("akashvaloper1qvsus5qg8yhre7k2c78xkkw4nvqqgev7zw5wzs", "akash1qqpzacvv5mtct473nlv8xr34kmxku9ncgqrm9s")
}

func validatorsDelegations(validatorId, delegatorId string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/staking/v1beta1/validators/" + validatorId + "/delegations/" + delegatorId

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

	DelegatorDetail1 := DelegatorDetail{}

	json.Unmarshal(body, &DelegatorDetail1)

	fmt.Println(DelegatorDetail1)

}
