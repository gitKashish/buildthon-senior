package main

import (
	"log"

	"github.com/p2eengineering/kalp-sdk-public/kalpsdk"
)

func main() {

	contract := kalpsdk.Contract{IsPayableContract: false}

	contract.Logger = kalpsdk.NewLogger()
	chaincode, err := kalpsdk.NewChaincode(&SmartContract{contract})
	contract.Logger.Info("DeVote")

	if err != nil {
		log.Panicf("Error creating KalpContractChaincode: %v", err)
	}

	// Start the chaincode
	if err := chaincode.Start(); err != nil {
		log.Panicf("Error starting chaincode: %v", err)
	}
}
