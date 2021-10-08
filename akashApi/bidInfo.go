package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BidInfoDetails struct {
	Bid struct {
		BidID struct {
			Owner    string `json:"owner"`
			Dseq     string `json:"dseq"`
			Gseq     int    `json:"gseq"`
			Oseq     int    `json:"oseq"`
			Provider string `json:"provider"`
		} `json:"bid_id"`
		State string `json:"state"`
		Price struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"price"`
		CreatedAt string `json:"created_at"`
	} `json:"bid"`
	EscrowAccount struct {
		ID struct {
			Scope string `json:"scope"`
			Xid   string `json:"xid"`
		} `json:"id"`
		Owner   string `json:"owner"`
		State   string `json:"state"`
		Balance struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balance"`
		Transferred struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"transferred"`
		SettledAt string `json:"settled_at"`
	} `json:"escrow_account"`
}

func main() {
	// Relace owner, provider, DSEQ, GSEQ, OSEQ values in function call template below
	bidInfo("akash1ss3ty253h6yun0a0fly8s0prcx34x4q2qewpkk", "akash1vky0uh4wayh9npd74uqesglpaxwymynnspf6a4", "2970401", "1", "1")
}

func bidInfo(owner, provider, dseq, gseq, oseq string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/market/v1beta1/bids/info?id.owner=" + owner + "&id.provider=" + provider + "&id.dseq=" + dseq + "&id.gseq=" + gseq + "&id.oseq=" + oseq

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

	BidInfoDetails1 := BidInfoDetails{}

	json.Unmarshal(body, &BidInfoDetails1)

	fmt.Println(BidInfoDetails1)

}
