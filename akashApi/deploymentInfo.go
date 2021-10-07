package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	deploymentInfo("akash1ss3ty253h6yun0a0fly8s0prcx34x4q2qewpkk", "2963732")
}

func deploymentInfo(ownerID, dseq string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/deployment/v1beta1/deployments/info?id.owner=" + ownerID + "&id.dseq=" + dseq

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

	fmt.Println(string(body))

}
