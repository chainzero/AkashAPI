package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Provider struct {
	Provider struct {
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
	} `json:"provider"`
}

func main() {
	// Replace Provider address value in function call template below
	providerDetail("akash1vky0uh4wayh9npd74uqesglpaxwymynnspf6a4")
}

func providerDetail(provider string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/provider/v1beta1/providers/" + provider

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

	Provider1 := Provider{}

	json.Unmarshal(body, &Provider1)

	fmt.Println(Provider1)

}
