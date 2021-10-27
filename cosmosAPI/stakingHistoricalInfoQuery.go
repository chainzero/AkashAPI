package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type BlockDetails struct {
	Hist struct {
		Header struct {
			Version struct {
				Block string `json:"block"`
				App   string `json:"app"`
			} `json:"version"`
			ChainID     string    `json:"chain_id"`
			Height      string    `json:"height"`
			Time        time.Time `json:"time"`
			LastBlockID struct {
				Hash          string `json:"hash"`
				PartSetHeader struct {
					Total int    `json:"total"`
					Hash  string `json:"hash"`
				} `json:"part_set_header"`
			} `json:"last_block_id"`
			LastCommitHash     string `json:"last_commit_hash"`
			DataHash           string `json:"data_hash"`
			ValidatorsHash     string `json:"validators_hash"`
			NextValidatorsHash string `json:"next_validators_hash"`
			ConsensusHash      string `json:"consensus_hash"`
			AppHash            string `json:"app_hash"`
			LastResultsHash    string `json:"last_results_hash"`
			EvidenceHash       string `json:"evidence_hash"`
			ProposerAddress    string `json:"proposer_address"`
		} `json:"header"`
		Valset []struct {
			OperatorAddress string `json:"operator_address"`
			ConsensusPubkey struct {
				Type string `json:"@type"`
				Key  string `json:"key"`
			} `json:"consensus_pubkey"`
			Jailed          bool   `json:"jailed"`
			Status          string `json:"status"`
			Tokens          string `json:"tokens"`
			DelegatorShares string `json:"delegator_shares"`
			Description     struct {
				Moniker         string `json:"moniker"`
				Identity        string `json:"identity"`
				Website         string `json:"website"`
				SecurityContact string `json:"security_contact"`
				Details         string `json:"details"`
			} `json:"description"`
			UnbondingHeight string    `json:"unbonding_height"`
			UnbondingTime   time.Time `json:"unbonding_time"`
			Commission      struct {
				CommissionRates struct {
					Rate          string `json:"rate"`
					MaxRate       string `json:"max_rate"`
					MaxChangeRate string `json:"max_change_rate"`
				} `json:"commission_rates"`
				UpdateTime time.Time `json:"update_time"`
			} `json:"commission"`
			MinSelfDelegation string `json:"min_self_delegation"`
		} `json:"valset"`
	} `json:"hist"`
}

func main() {
	// Replace block height value in function call template below
	stakingHistoricalInfoQuery("3236235")
}

func stakingHistoricalInfoQuery(blockHeight string) {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/staking/v1beta1/historical_info/" + blockHeight

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

	BlockDetails1 := BlockDetails{}

	json.Unmarshal(body, &BlockDetails1)

	fmt.Println(BlockDetails1)

}
