package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ProvisionDetail struct {
	AnnualProvisions string `json:"annual_provisions"`
}

func main() {
	mintingAnuualProvisions()
}

func mintingAnuualProvisions() {
	method := "GET"

	url := "http://135.181.60.250:1317/cosmos/mint/v1beta1/annual_provisions"

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

	ProvisionDetail1 := ProvisionDetail{}

	json.Unmarshal(body, &ProvisionDetail1)

	fmt.Println(ProvisionDetail1)

}
