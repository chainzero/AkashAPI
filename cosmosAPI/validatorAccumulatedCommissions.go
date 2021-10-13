package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CommissionDetail struct {
	Commission struct {
		Commission []struct {
			Denom  string `json:"denom"`
			Amount string `json:"amount"`
		} `json:"commission"`
	} `json:"commission"`
}

func main() {
	// Replace Validator Account value in function call template below
	validatorCommissions("akashvaloper14kn0kk33szpwus9nh8n87fjel8djx0y0uzn073")
}

func validatorCommissions(validator string) {
	method := "GET"

	url := "http://135.181.60.250:1317//cosmos/distribution/v1beta1/validators/" + validator + "/commission"

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

	CommissionDetail1 := CommissionDetail{}

	json.Unmarshal(body, &CommissionDetail1)

	fmt.Println(CommissionDetail1)

}
