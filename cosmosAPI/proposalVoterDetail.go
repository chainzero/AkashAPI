package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VoterDetail struct {
	Vote struct {
		ProposalID string `json:"proposal_id"`
		Voter      string `json:"voter"`
		Option     string `json:"option"`
	} `json:"vote"`
}

func main() {
	// Replace proposal ID and voter address values in function call template below
	proposalVoterDetail("13", "akash1qqf44hzms4uuqnhl39664gujus325uc3vexvmt")
}

func proposalVoterDetail(id, voterAddress string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/gov/v1beta1/proposals/" + id + "/votes/" + voterAddress

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

	VoterDetail1 := VoterDetail{}

	json.Unmarshal(body, &VoterDetail1)

	fmt.Println(VoterDetail1)

}
