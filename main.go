package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	kosmosv1 "github.com/kosmos-industrie40/setting-up-kosmos-edge/api/v1"
	"github.com/pelletier/go-toml"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	scheme = runtime.NewScheme()
)

type contract struct {
	ContractId       int         `json:"contract_id,omitempty"`
	Activated        int         `json:"activated,omitempty"`
	ModificationTime string      `json:"modification_time,omitempty"`
	BcContent        interface{} `json:"bc_content,omitempty"`
	EdgeContent      interface{} `json:"edge_content,omitempty"`
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var config, err = toml.LoadFile("mock_contract_engine_parameters.toml")
	if err != nil {
		panic(fmt.Sprintf("Error %s", err))
	}
	activated_contracts, err := getActivatedContractsFromKosmosGlobal(config)
	if err != nil {
		panic(fmt.Sprintf("Error %s", err))
	}

	c, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		panic(fmt.Sprintf("Error %s", err))
	}

	for id, contract := range activated_contracts {
		// fmt.Printf("%+v", contract)
		err := createObject(c, id, contract)

		if err != nil {
			panic(fmt.Sprintf("Error %s", err))
		}
	}

}

func init() {
	if err := kosmosv1.AddToScheme(scheme); err != nil {
		panic(fmt.Sprintf("Error %s", err))
	}
}
