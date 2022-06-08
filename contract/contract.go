package contract

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/kimanikelly/cli-ttBank/client"
)

type AddressList struct {
	Rinkeby string
}

type ContractResponseData struct {
	Addresses AddressList
	Abi       abi.ABI
	Bytecode  string
}

var ContractData ContractResponseData

var TTBankAddress string

func ContractInstance() *TTBank {

	ttBankResp := FetchContractData().Body

	ttBankBody, bodyErr := io.ReadAll(ttBankResp)

	if bodyErr != nil {
		log.Fatalf("Failed to read the response body %v", bodyErr)

	}

	json.Unmarshal([]byte(ttBankBody), &ContractData)

	// Checks if the 'NETWORK' env variable is set to localhost
	if os.Getenv("NETWORK") == "localhost" {

		// Assigns the value of the localhost TTBank address
		TTBankAddress = os.Getenv("LOCAL_TTBANK_ADDRESS")
	}

	// Checks if the 'NETWORK' env variable is set to rinkeby
	if os.Getenv("NETWORK") == "rinkeby" {

		// Assigns the value of the rinkeby TTBank address
		TTBankAddress = ContractData.Addresses.Rinkeby
	}

	// Creates the TTBank.sol contract instance
	ttBank, ttBankErr := NewTTBank(common.HexToAddress(TTBankAddress), client.ClientConnection())

	// If ttBankErr does not return nil(zero value) throw an error
	if ttBankErr != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", ttBankErr)

	}

	// Returns the TTBank.sol contract instance
	return ttBank
}
