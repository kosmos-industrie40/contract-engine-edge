package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	kosmosv1 "github.com/kosmos-industrie40/setting-up-kosmos-edge/api/v1"
	"github.com/pelletier/go-toml"
)

func getActivatedContractsFromKosmosGlobal(config *toml.Tree) (map[int]kosmosv1.KosmosKubeEdgeSpec, error) {
	var KOSMoS_GLOBAL_HOST = config.Get("KOSMoS_Global.host").(string)

	// all_contracts := []string{}
	req, err := http.NewRequest(http.MethodGet, KOSMoS_GLOBAL_HOST+"/contractEngine/allContractInstances", nil)
	if err != nil {
		return nil, err
	}

	token, err := getKeycloakAccessToken(config)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got response: %d", resp.StatusCode)
	}
	activated_contracts, err := parseBodyForContracts(resp.Body)
	if err != nil {
		return nil, err
	}

	return activated_contracts, nil
}

// parseBody returns tall contracts, where activated==1
func parseBodyForContracts(body io.Reader) (map[int]kosmosv1.KosmosKubeEdgeSpec, error) {

	// contracts := []contract{}
	activated_contracts := map[int]kosmosv1.KosmosKubeEdgeSpec{}

	// body_instances := []interface{}{}
	contracts := []contract{}
	if err := json.NewDecoder(body).Decode(&contracts); err != nil {
		return nil, err
	}

	for i := range contracts {
		var helper map[string]interface{}
		helperType := reflect.TypeOf(helper)
		typeOfEdgeContent := reflect.TypeOf(contracts[i].EdgeContent)

		if contracts[i].Activated == 1 && helperType == typeOfEdgeContent {
			jsonData, err := json.Marshal(contracts[i].EdgeContent)
			if err != nil {
				return nil, err
			}
			var contract kosmosv1.KosmosKubeEdgeSpec
			err = json.Unmarshal(jsonData, &contract)
			if err != nil {
				return nil, err
			}
			activated_contracts[contracts[i].ContractId] = contract
		}
	}

	return activated_contracts, nil
}

func getKeycloakAccessToken(config *toml.Tree) (string, error) {
	var KEYCLOAK_URL = config.Get("keycloak.keycloak_url").(string)
	var KEYCLOAK_USERNAME = config.Get("keycloak.keycloak_username").(string)
	var KEYCLOAK_USER_PW = config.Get("keycloak.keycloak_pw").(string)
	var KEYCLOAK_CLIENT = config.Get("keycloak.keycloak_client").(string)
	var KEYCLOAK_SECRET = config.Get("keycloak.keycloak_secret").(string)

	data := url.Values{}
	data.Set("username", KEYCLOAK_USERNAME)
	data.Set("password", KEYCLOAK_USER_PW)
	data.Set("grant_type", "password")
	data.Set("client_id", KEYCLOAK_CLIENT)
	data.Set("client_secret", KEYCLOAK_SECRET)

	resp, err := http.Post(KEYCLOAK_URL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	keycloakResponse := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&keycloakResponse); err != nil {
		return "", err
	}

	token, ok := keycloakResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse token")
	}

	return token, err
}
