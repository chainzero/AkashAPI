package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OrderListDetail struct {
	Orders []struct {
		OrderID struct {
			Owner string `json:"owner"`
			Dseq  string `json:"dseq"`
			Gseq  int    `json:"gseq"`
			Oseq  int    `json:"oseq"`
		} `json:"order_id"`
		State string `json:"state"`
		Spec  struct {
			Name         string `json:"name"`
			Requirements struct {
				SignedBy struct {
					AllOf []interface{} `json:"all_of"`
					AnyOf []string      `json:"any_of"`
				} `json:"signed_by"`
				Attributes []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"attributes"`
			} `json:"requirements"`
			Resources []struct {
				Resources struct {
					CPU struct {
						Units struct {
							Val string `json:"val"`
						} `json:"units"`
						Attributes []interface{} `json:"attributes"`
					} `json:"cpu"`
					Memory struct {
						Quantity struct {
							Val string `json:"val"`
						} `json:"quantity"`
						Attributes []interface{} `json:"attributes"`
					} `json:"memory"`
					Storage struct {
						Quantity struct {
							Val string `json:"val"`
						} `json:"quantity"`
						Attributes []interface{} `json:"attributes"`
					} `json:"storage"`
					Endpoints []struct {
						Kind string `json:"kind"`
					} `json:"endpoints"`
				} `json:"resources"`
				Count int `json:"count"`
				Price struct {
					Denom  string `json:"denom"`
					Amount string `json:"amount"`
				} `json:"price"`
			} `json:"resources"`
		} `json:"spec"`
		CreatedAt string `json:"created_at"`
	} `json:"orders"`
	Pagination struct {
		NextKey interface{} `json:"next_key"`
		Total   string      `json:"total"`
	} `json:"pagination"`
}

func main() {
	// Replace DSEQ value in function call template below
	ordersList("2980326")
}

func ordersList(dseq string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/market/v1beta1/orders/list?filters.dseq=" + dseq

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

	OrderListDetail1 := OrderListDetail{}

	json.Unmarshal(body, &OrderListDetail1)

	fmt.Println(OrderListDetail1)

}
