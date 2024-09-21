package main

import (
	"fmt"

	"github.com/p2eengineering/kalp-sdk-public/kalpsdk"
)

const nameKey = "name"
const pollKey = "poll"
const totalPollKey = "totalPoll"

type SmartContract struct {
	kalpsdk.Contract
}

type Vote struct {
	optionIndex   int
	locationIndex int
}

type Poll struct {
	id        string
	question  string
	options   []string
	locations []string
	votes     []Vote
}

func (s *SmartContract) Initialize(sdk kalpsdk.TransactionContextInterface, name string) (bool, error) {
	// Check minter authorization - this sample assumes `mailabs` is the central banker with privilege to intitialize contract
	clientMSPID, err := sdk.GetClientIdentity().GetMSPID()
	if err != nil {
		return false, fmt.Errorf("failed to get MSPID: %v", err)
	}
	if clientMSPID != "mailabs" {
		return false, fmt.Errorf("client is not authorized to initialize contract")
	}

	// Check contract are not already set, client is not authorized to change them once initialized
	bytes, err := sdk.GetState(nameKey)
	if err != nil {
		return false, fmt.Errorf("failed to get name %s: %v", nameKey, err)
	}

	if bytes != nil {
		return false, fmt.Errorf("contract options are set, client is not allowed to change them")
	}

	err = sdk.PutStateWithoutKYC(nameKey, []byte(name))
	if err != nil {
		return false, fmt.Errorf("failed to set name: %v", err)
	}

	return true, nil
}
