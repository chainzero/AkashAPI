package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ProposalDetail struct {
	Proposal struct {
		ProposalID string `json:"proposal_id"`
		Content    struct {
			Type        string `json:"@type"`
			Title       string `json:"title"`
			Description string `json:"description"`
			Changes     []struct {
				Subspace string `json:"subspace"`
				Key      string `json:"key"`
				Value    string `json:"value"`
			} `json:"changes"`
			Plan struct {
				Name                string      `json:"name"`
				Time                time.Time   `json:"time"`
				Height              string      `json:"height"`
				Info                string      `json:"info"`
				UpgradedClientState interface{} `json:"upgraded_client_state"`
			} `json:"plan"`
			Amount []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"amount"`
		} `json:"content,omitempty"`
		Status           string `json:"status"`
		FinalTallyResult struct {
			Yes        string `json:"yes"`
			Abstain    string `json:"abstain"`
			No         string `json:"no"`
			NoWithVeto string `json:"no_with_veto"`
		} `json:"final_tally_result"`
		SubmitTime     time.Time `json:"submit_time"`
		DepositEndTime time.Time `json:"deposit_end_time"`
		TotalDeposit   []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"total_deposit"`
		VotingStartTime time.Time `json:"voting_start_time"`
		VotingEndTime   time.Time `json:"voting_end_time"`
	} `json:"proposal"`
}

func main() {
	// Replace ID value in function call template below
	proposalID("13")
}

func proposalID(id string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/gov/v1beta1/proposals/" + id

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

	ProposalDetail1 := ProposalDetail{}

	json.Unmarshal(body, &ProposalDetail1)

	fmt.Println(ProposalDetail1)

}
