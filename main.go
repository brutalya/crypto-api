package main

import (
	"fmt"

	"github.com/brutalya/crypto-api/crypto"
)

func main() {
	// List supported cryptocurrencies
	supportedCryptos, err := crypto.ListSupportedCryptos()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Supported Cryptocurrencies:")
	for _, crypto := range supportedCryptos {
		fmt.Println(crypto)
	}

	// Fetch data for a list of cryptocurrencies
	cryptoList := []string{"bitcoin", "ethereum", "litecoin"}
	data, err := crypto.FetchCryptoDataList(cryptoList)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Crypto Data:")
	for crypto, values := range data {
		fmt.Println(crypto)
		fmt.Println("Price (USD):", values.(map[string]interface{})["usd"])
	}

	// Get the current price of Bitcoin
	bitcoinPrice, err := crypto.GetCryptoPrice("bitcoin")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Current Bitcoin Price (USD): %.2f\n", bitcoinPrice)
}
