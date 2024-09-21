package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/p2eengineering/kalp-sdk-public/kalpsdk"
)

// Define objectType names for prefix
const pollPrefix = "poll:"

// Define key names for options
const nameKey = "name"
const totalPollKey = "totalPoll"

type SmartContract struct {
	kalpsdk.Contract
}

type Vote struct {
	OptionIndex   int `json:"optionIndex"`
	LocationIndex int `json:"locationIndex"`
}

type Poll struct {
	Question  string    `json:"question"`
	Options   []string  `json:"options"`
	Locations []string  `json:"locations"`
	Expiry    time.Time `json:"expiry"`
	Votes     []Vote    `json:"votes"`
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

func CreatePoll(sdk kalpsdk.TransactionContextInterface, question string, options []string, locations []string, durationHours int) (*Poll, error) {
	// Check if context has initialized first
	initialized, err := checkInitialized(sdk)
	if err != nil {
		return nil, fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return nil, fmt.Errorf("contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Validating params
	if question == "" {
		return nil, fmt.Errorf("cannot create an empty poll")
	}

	if len(options) < 2 {
		return nil, fmt.Errorf("poll must have at least 2 options")
	}

	if durationHours < 1 {
		return nil, fmt.Errorf("poll must last for at least 1 hour")
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("cannot create poll id: %v", err)
	}

	expiry := time.Now().Add(time.Hour * time.Duration(durationHours))
	locations = append(locations, "other")

	poll := &Poll{
		Question:  question,
		Options:   options,
		Locations: locations,
		Expiry:    expiry,
		Votes:     []Vote{},
	}

	_pollKey, err := sdk.CreateCompositeKey(pollPrefix, []string{id.String()})
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}

	pollBytes, err := json.Marshal(poll)
	if err != nil {
		return nil, fmt.Errorf("error marshalling bytes: %v", err)
	}

	err = sdk.PutStateWithoutKYC(_pollKey, pollBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to put state with pollKey %s: %v", _pollKey, err)
	}

	var totalPoll int

	totalPollBytes, err := sdk.GetState(totalPollKey)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	// If not polls have been created yet, initialize the totalPoll
	if totalPollBytes == nil {
		totalPoll = 0
	} else {
		totalPoll, _ = strconv.Atoi(string(totalPollBytes))
	}
	// Increment total poll count
	totalPoll++

	err = sdk.PutStateWithoutKYC(totalPollKey, []byte(strconv.Itoa(totalPoll)))
	if err != nil {
		return nil, err
	}

	return poll, nil
}

func SubmitVote(sdk kalpsdk.TransactionContextInterface, pollId string, optionIndex int, locationIndex int) error {
	// Check if context has initialized first
	initialized, err := checkInitialized(sdk)
	if err != nil {
		return fmt.Errorf("failed to check if contract is already initialized: %v", err)
	}
	if !initialized {
		return fmt.Errorf("contract options need to be set before calling any function, call Initialize() to initialize contract")
	}

	// Validating params
	if len(pollId) != 36 {
		return fmt.Errorf("invalid poll id")
	}

	if optionIndex < 0 {
		return fmt.Errorf("option index cannot be negative")
	}

	if locationIndex < 0 {
		return fmt.Errorf("location index cannot be negative")
	}

	_pollKey, err := sdk.CreateCompositeKey(pollPrefix, []string{pollId})
	if err != nil {
		return fmt.Errorf("error creating composite key: %v", err.Error())
	}

	pollBytes, err := sdk.GetState(_pollKey)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if pollBytes == nil {
		return fmt.Errorf("the poll %s does not exist", pollId)
	}

	poll := new(Poll)
	err = json.Unmarshal(pollBytes, poll)
	if err != nil {
		return fmt.Errorf("failed unmarshall pollBytes: %v", err)
	}

	if optionIndex >= len(poll.Options) {
		return fmt.Errorf("optionIndex out of bound")
	}

	if locationIndex >= len(poll.Locations) {
		return fmt.Errorf("locationIndex out of bound")
	}

	poll.Votes = append(poll.Votes, Vote{
		OptionIndex:   optionIndex,
		LocationIndex: locationIndex,
	})

	pollBytes, err = json.Marshal(poll)
	if err != nil {
		return fmt.Errorf("failed to marshal poll after vote submission")
	}

	err = sdk.PutStateWithoutKYC(_pollKey, pollBytes)
	if err != nil {
		return fmt.Errorf("failed to put state with pollKey %s: %v", _pollKey, err)
	}

	return nil
}

// Helper Functions

func checkInitialized(sdk kalpsdk.TransactionContextInterface) (bool, error) {
	name, err := sdk.GetState(nameKey)
	if err != nil {
		return false, err
	}

	if name == nil {
		return false, nil
	}

	return true, nil
}
