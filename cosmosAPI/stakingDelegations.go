package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StakingDetail struct {
	DelegationResponses []struct {
		Delegation struct {
			DelegatorAddress string `json:"delegator_address"`
			ValidatorAddress string `json:"validator_address"`
			Shares           string `json:"shares"`
		} `json:"delegation"`
		Balance struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balance"`
	} `json:"delegation_responses"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func main() {
	// Replace ID value in function call template below
	stakingDelegations("akash1w3k6qpr4uz44py4z68chfrl7ltpxwtkngnc6xk")
}

func stakingDelegations(id string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/staking/v1beta1/delegations/" + id

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

	StakingDetail1 := StakingDetail{}

	json.Unmarshal(body, &StakingDetail1)

	fmt.Println(StakingDetail1)

}
