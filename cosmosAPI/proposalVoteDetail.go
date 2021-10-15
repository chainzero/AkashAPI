package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VoteDetail struct {
	Votes []struct {
		ProposalID string `json:"proposal_id"`
		Voter      string `json:"voter"`
		Option     string `json:"option"`
	} `json:"votes"`
	Pagination struct {
		NextKey string `json:"next_key"`
		Total   string `json:"total"`
	} `json:"pagination"`
}

func main() {
	// Replace ID value in function call template below
	proposalVoteDetail("13")
}

func proposalVoteDetail(id string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/gov/v1beta1/proposals/" + id + "/votes"

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

	VoteDetail1 := VoteDetail{}

	json.Unmarshal(body, &VoteDetail1)

	fmt.Println(VoteDetail1)

}
