package contract

import (
	"log"
	"net/http"
)

func FetchContractData() (*http.Response, *http.Response) {

	// GET request to access the ABI, Bytecode, and addresses of Token.sol
	tokenResp, tokenErr := http.Get("https://kimanikelly-contractapi.herokuapp.com/tokenContract")

	// If tokenErr does not equal nil(zero value) throw an err
	if tokenErr != nil {
		log.Fatalf("Failed to perform a GET request to the /tokenContract endpoint %v", tokenErr)
	}

	// GET request to access the ABI, Bytecode, and addresses of TTBank.sol
	ttBankResp, ttBankErr := http.Get("https://kimanikelly-contractapi.herokuapp.com/tokenContract/ttBank")

	// If ttBankErr does not equal nil(zero value) throw an err
	if ttBankErr != nil {
		log.Fatalf("Failed to perform a GET request to the /ttBank endpoint %v", ttBankErr)
	}

	// Returns the response made to the /tokenContract and /ttBank endpoints
	return tokenResp, ttBankResp
}
