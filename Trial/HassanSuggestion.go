package main

import (
	"encoding/json"
	"errors"
	"fmt"
//	"strconv"
//	"time"
//	"strings"

  //  "os"
//    "strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"

)
//Models 


type Transaction struct {
	Username       string     `json:"Username"`
	ItemName       string     `json:"ItemName"`
	QRCode         string     `json:"QRCode"`
	TransID        string     `json:"TransID"`
	Date            string    `json:"Date"`
	Time           string     `json:"Time"`
	Price          string     `json:"Price"`

}

type UserAccount struct 
{
    Username       string                `json:"Username"`
    Transactions   []string              `json:"Transactions"`
    
    
}


//Functions
// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))     //To Start Chaincode
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("SwiftVending", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "Buy" {
		return t.Buy(stub , args)
	}
    
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	 } 
	fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}


// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func (t *SimpleChaincode) Buy(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) { 
  	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}
        // Insert Transaction
	trans:=Transaction{Username:args[0], ItemName:args[1], QRCode: args[2] ,TransID: args[3], Date:args[4], Time:args[5], Price:args[6]}
	
        transactionBytes, err := json.Marshal(&trans)
	// Missing Check here 
	err = stub.PutState(args[3] ,transactionBytes) 
	
	
	// Update account 
	
	var account UserAccount
	
	existingBytes, err3 := stub.GetState(args[0])
	// if user doens't exist : don't do the following 
	if err3!=nil{
		return nil, err
	}
	
	err3 = json.Unmarshal(existingBytes, &account)
	
	account.Transactions=append(account.Transactions, args[3]) 
	account.Username=args[0] 
	accountBytes, err3 := json.Marshal(&account)
	
	err2 := stub.PutState(args[0] ,accountBytes) 
	
	
	if err2==nil{
		return nil, err
	}
	
	if err3==nil{
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return nil, nil
}
