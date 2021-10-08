package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeploymentGroupDetails struct {
	Deployment struct {
		DeploymentID struct {
			Owner string `json:"owner"`
			Dseq  string `json:"dseq"`
		} `json:"deployment_id"`
		State     string `json:"state"`
		Version   string `json:"version"`
		CreatedAt string `json:"created_at"`
	} `json:"deployment"`
	Groups []struct {
		GroupID struct {
			Owner string `json:"owner"`
			Dseq  string `json:"dseq"`
			Gseq  int    `json:"gseq"`
		} `json:"group_id"`
		State     string `json:"state"`
		GroupSpec struct {
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
		} `json:"group_spec"`
		CreatedAt string `json:"created_at"`
	} `json:"groups"`
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
	// Relace ownerID, DSEQ, and GSEQ values in function call template below
	deploymentGroupInfo("akash1ss3ty253h6yun0a0fly8s0prcx34x4q2qewpkk", "2963732", "1")
}

func deploymentGroupInfo(ownerID, dseq, gseq string) {
	method := "GET"

	url := "http://135.181.60.250:1317/akash/deployment/v1beta1/deployments/info?id.owner=" + ownerID + "&id.dseq=" + dseq + "&id.gseq=" + gseq

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

	DeploymentGroupDetails1 := DeploymentGroupDetails{}

	json.Unmarshal(body, &DeploymentGroupDetails1)

	fmt.Println(DeploymentGroupDetails1)

}
