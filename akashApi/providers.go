package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	providers()
}

type ProviderDetail struct {
	Providers []struct {
		Owner      string `json:"owner"`
		HostURI    string `json:"host_uri"`
		Attributes []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"attributes"`
		Info struct {
			Email   string `json:"email"`
			Website string `json:"website"`
		} `json:"info"`
	} `json:"providers"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func providers() {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/provider/v1beta1/providers"

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

	ProviderDetail1 := ProviderDetail{}

	json.Unmarshal(body, &ProviderDetail1)

	fmt.Println(ProviderDetail1)

}
