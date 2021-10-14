package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ValidatorRewardDetail struct {
	Rewards struct {
		Rewards []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"rewards"`
	} `json:"rewards"`
}

func main() {
	// Replace Validator Account value in function call template below
	validatorRewards("akashvaloper14kn0kk33szpwus9nh8n87fjel8djx0y0uzn073")
}

func validatorRewards(validator string) {
	method := "GET"

	url := "http://135.181.60.250:1317//cosmos/distribution/v1beta1/validators/" + validator + "/outstanding_rewards"

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

	ValidatorRewardDetail1 := ValidatorRewardDetail{}

	json.Unmarshal(body, &ValidatorRewardDetail1)

	fmt.Println(ValidatorRewardDetail1)

}
