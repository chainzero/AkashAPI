package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type ProposalListDetail struct {
	Proposals []struct {
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
	} `json:"proposals"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func main() {
	proposalList()
}

func proposalList() {
	method := "GET"

	url := "http://135.181.181.122:1518/cosmos/gov/v1beta1/proposals"

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

	ProposalListDetail1 := ProposalListDetail{}

	json.Unmarshal(body, &ProposalListDetail1)

	fmt.Println(ProposalListDetail1)

}
