package contract

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"math/big"

// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/kimanikelly/cli-ttBank/client"
// 	"github.com/kimanikelly/cli-ttBank/signer"
// 	"github.com/kimanikelly/cli-ttBank/utils"
// )

// type AddressList struct {
// 	Rinkeby string
// }

// type ContractResponseData struct {
// 	Addresses AddressList
// 	Abi       abi.ABI
// 	Bytecode  string
// }

// var ContractData ContractResponseData

// func DeployContracts() {

// 	fromAddress, privateKey := signer.Wallet()

// 	// Returns the nonce
// 	nonce := utils.Nonce(fromAddress)

// 	// Returns the gasPrice
// 	gasPrice := utils.GasPrice()

// 	// Returns the chainId
// 	chainId := utils.ChainID()

// 	// Binds the connected signer to the transaction options
// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

// 	// If err does not equal nil (zero value) throw an error
// 	if err != nil {
// 		log.Fatalf("Failed to build the NewKeyedTransactorWithChainID %v", err)
// 	}

// 	// Sets the nonce to send with a transaction
// 	auth.Nonce = big.NewInt(int64(nonce))

// 	// Sets the amount of test ETH to send with a transaction
// 	auth.Value = big.NewInt(0)

// 	// Sets the gasLimit for a transaction
// 	auth.GasLimit = uint64(300000)

// 	// Sets the gas price for a transaction
// 	auth.GasPrice = gasPrice

// 	tokenResp, _ := FetchContractData()

// 	tokenBody, err := io.ReadAll(tokenResp.Body)

// 	if err != nil {
// 		log.Fatalf("Failed to read the response body for the ./tokenContract endpoint %v", err)
// 	}

// 	json.Unmarshal([]byte(string(tokenBody)), &ContractData)

// 	address, _, contract, _ := bind.DeployContract(auth, ContractData.Abi, []byte(ContractData.Bytecode), client.ClientConnection())
// 	fmt.Println(address)
// 	fmt.Println(fromAddress)

// token.Initialize(auth, "TEST", "TT")

// name, _ := token.Name(&bind.CallOpts{})

// fmt.Println(name)

// fmt.Println(tx)
// fmt.Println(bound)

// }
