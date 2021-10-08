package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BidListDetails struct {
	Bids []struct {
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
	} `json:"bids"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func main() {
	// Relace ownerID and DSEQ values in function call template below
	bidList("2969932")
}

func bidList(dseq string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/market/v1beta1/bids/list?filters.dseq=" + dseq

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

	BidListDetails1 := BidListDetails{}

	json.Unmarshal(body, &BidListDetails1)

	fmt.Println(BidListDetails1)

}
