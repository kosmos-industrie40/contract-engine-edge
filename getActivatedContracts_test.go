package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/jarcoal/httpmock"
	kosmosv1 "github.com/kosmos-industrie40/setting-up-kosmos-edge/api/v1"
	"github.com/pelletier/go-toml"
)

// func TestGetActivatedContractsFromKosmosGlobal(t *testing.T) {

// }

var testJson = `
[
   {
     "contract_id": 1,
     "activated": 0,
     "modification_time": "Waiting for activation by every Participant"
   },
   {
     "contract_id": 2,
     "activated": 1,
     "modification_time": "2022-01-01T00:00:01",
     "bc_content": "{'example': 'exampleexample'}",
     "edge_content": {
       "$schema": "formal.json",
       "body": {
         "contract": {
           "valid": {
             "start": "22022-01-01T00:00:01.000Z",
             "end": "2032-01-01T00:00:01.000Z"
           },
           "creationTime": "2022-01-01T00:00:01.0000Z",
           "partnes": [],
           "permissions": {
             "read": [],
             "write": [
               "/test"
             ]
           },
           "id": "2",
           "version": "v1"
         },
         "requiredTechnicalContainers": [
           {
             "system": "edge",
             "containers": [
               {
                 "url": "example.com/example-url-rtc",
                 "tag": "example-tag-rtc"
               }
             ]
           }
         ],
         "machine": "12345",
         "kosmosLocalSystems": [
           "edge"
         ],
         "sensors": [
           {
             "name": "sensor1",
             "storageDuration": [
               {
                 "duration": "24h",
                 "systemName": "example"
               }
             ],
             "meta": {}
           }
         ],
         "checkSignatures": false,
         "machineConnection": [
           {
             "connector": {
               "url": "example.com/example-url-mc",
               "tag": "example-tag-mc"
             },
             "connectionData": {
               "uri": "",
               "credentials": {},
               "machineID": "12345"
             },
             "subscribeData": {}
           }
         ],
         "blockchainConnection": {
           "uri": "blockchain",
           "containerList": [
             {
               "url": "example.com/example-url-bc",
               "tag": "example-tag-bc"
             }
           ],
           "sensors": []
         }
       },
       "signature": {
         "signature": "",
         "meta": {
           "date": "2022-01-01T00:00:00Z",
           "algorithm": "",
           "serialNumber": ""
         }
       }
     }
   }
]`

func TestKosmosKubeEdgeSpec(t *testing.T) {
	t.Run("Test fields KosmosKubeEdgeSpec", func(t *testing.T) {
		activated_contracts, err := parseBodyForContracts(strings.NewReader(testJson))

		if err != nil {
			t.Fatalf("%v", err)
		}
		contractID := 2
		var contract kosmosv1.KosmosKubeEdgeSpec = activated_contracts[contractID]

		if contract.Body.Machine != "12345" {
			t.Fatalf("Expected Maschine 12345 but got %v", contract.Body.Machine)
		}
		if len(contract.Body.Sensors) != 1 {
			t.Fatalf("Expected one sensor but got %v", len(contract.Body.Sensors))
		}
		var checkSignature = contract.Body.CheckSignature
		if checkSignature != false {
			t.Fatalf("Expected false but got %v", checkSignature)
		}
	})
}

func TestParseBodyForContracts(t *testing.T) {
	t.Run("test contract examples", func(t *testing.T) {
		activated_contracts, err := parseBodyForContracts(strings.NewReader(testJson))

		if err != nil {
			t.Fatalf("%v", err)
		}

		// Check one activated contract
		if len(activated_contracts) != 1 {
			t.Fatalf("Expected one activated contract, got %d", len(activated_contracts))
		}
		// Check if ContractID is in Map
		contractID := 2
		if _, found := activated_contracts[contractID]; !found {
			t.Fatalf("Excpected map with Key %v , but was not found", contractID)
		}

	})
}

func TestGetKeycloackAccessToken(t *testing.T) {
	var config, err = toml.LoadFile("mock_contract_engine_parameters.toml")
	if err != nil {
		t.Fatalf("Failed loading TOML: %v", err)
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mocking HTTP POST Request
	httpmock.RegisterResponder("POST",
		config.Get("keycloak.keycloak_url").(string),
		httpmock.NewStringResponder(200, `{"access_token":"TOKEN_TEST"}`))

	token, err := getKeycloakAccessToken(config)

	if err != nil {
		t.Fatalf("Getting error by request: %v", err)
	}

	type_token := reflect.TypeOf(token).Kind()
	if type_token != reflect.String {
		t.Fatalf("Should be a string but got %v", type_token)
	}

}
