package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TallyDetail struct {
	Tally struct {
		Yes        string `json:"yes"`
		Abstain    string `json:"abstain"`
		No         string `json:"no"`
		NoWithVeto string `json:"no_with_veto"`
	} `json:"tally"`
}

func main() {
	// Replace ID value in function call template below
	proposalVoteTally("13")
}

func proposalVoteTally(id string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/gov/v1beta1/proposals/" + id + "/tally"

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

	// DepositDetails1 := DepositDetails{}
	TallyDetail1 := TallyDetail{}

	json.Unmarshal(body, &TallyDetail1)

	fmt.Println(TallyDetail1)

}
