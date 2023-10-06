package crypto

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

// ListSupportedCryptos returns a list of supported cryptocurrencies from the API.
func ListSupportedCryptos() ([]string, error) {
	client := resty.New()
	url := "https://api.coingecko.com/api/v3/coins/list"

	var cryptos []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
	}

	_, err := client.R().SetResult(&cryptos).Get(url)
	if err != nil {
		return nil, err
	}

	var cryptoNames []string
	for _, crypto := range cryptos {
		cryptoNames = append(cryptoNames, crypto.Symbol)
	}

	return cryptoNames, nil
}

// FetchCryptoDataList fetches data for a list of cryptocurrencies.
func FetchCryptoDataList(cryptoList []string) (map[string]interface{}, error) {
	client := resty.New()
	url := "https://api.coingecko.com/api/v3/simple/price"

	params := map[string]string{
		"ids":           "",
		"vs_currencies": "usd",
	}

	// Compose the ids parameter with a comma-separated list of cryptocurrencies
	for _, crypto := range cryptoList {
		params["ids"] += crypto + ","
	}

	resp, err := client.R().SetQueryParams(params).Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == 200 {
		// Parse the JSON response into a map
		var data map[string]interface{}
		err := json.Unmarshal(resp.Body(), &data)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	return nil, nil
}
