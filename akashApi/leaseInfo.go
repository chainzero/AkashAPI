package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LeaseDetails struct {
	Lease struct {
		LeaseID struct {
			Owner    string `json:"owner"`
			Dseq     string `json:"dseq"`
			Gseq     int    `json:"gseq"`
			Oseq     int    `json:"oseq"`
			Provider string `json:"provider"`
		} `json:"lease_id"`
		State string `json:"state"`
		Price struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"price"`
		CreatedAt string `json:"created_at"`
	} `json:"lease"`
	EscrowPayment struct {
		AccountID struct {
			Scope string `json:"scope"`
			Xid   string `json:"xid"`
		} `json:"account_id"`
		PaymentID string `json:"payment_id"`
		Owner     string `json:"owner"`
		State     string `json:"state"`
		Rate      struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"rate"`
		Balance struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"balance"`
		Withdrawn struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"withdrawn"`
	} `json:"escrow_payment"`
}

func main() {
	// Replace owner, provider, DSEQ, GSEQ, OSEQ values in function call template below
	leaseInfo("akash1ss3ty253h6yun0a0fly8s0prcx34x4q2qewpkk", "akash1vky0uh4wayh9npd74uqesglpaxwymynnspf6a4", "2980326", "1", "1")
}

func leaseInfo(owner, provider, dseq, gseq, oseq string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/market/v1beta1/leases/info?id.owner=" + owner + "&id.provider=" + provider + "&id.dseq=" + dseq + "&id.gseq=" + gseq + "&id.oseq=" + oseq

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

	LeaseDetails1 := LeaseDetails{}

	json.Unmarshal(body, &LeaseDetails1)

	fmt.Println(LeaseDetails1)

}
