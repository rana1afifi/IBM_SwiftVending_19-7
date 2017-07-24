/*
Copyright IBM Corp 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"

)
//Models 
type Item struct {
	Flavor           string  `json:"flavor"`
	ExpiryDate       string  `json:"expirydate"`
	Price            float64 `json:"price"`
	Calories         int     `json:"calories"`
	Brand            string  `json:"brand"`
	Ingredients      string  `json:"ingredients"`
	Size             string  `json:"size"`	
    Code 	         int     `json:"code"`
	Category         string  `json:"category"`
	
}



type Account struct {
	Email          string  `json:"email"`
	Name           string  `json:"name"`
	CashBalance    float64 `json:"cashBalance"`
	Password       string  `json:"password"`
	AssetsIds   []string `json:"assetIds"`

}

type Transaction struct {
	Code            int   `json:"code"`
	Email          string   `json:"email"`
	Date  		   string   `json:"date"`
	Time   		   int      `json:"time"`

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

	err := stub.PutState("hello_world", []byte(args[0]))
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
	} else if function == "CreateTransaction" {
		return t.CreateTransaction(stub, args)
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

// CreateTransaction to add asset to certain userID
func (t *SimpleChaincode) CreateTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var userId, assetId string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	userId = args[0] 
	assetId= args[1]
	var assetIds []string
        var transaction []string
	 
	var trans = Transaction{Code: 5, Email:userId}
        accountBytes, err := json.Marshal(&trans)
	if err != nil {
		fmt.Println("error creating transaction" + Transaction.Code)
		return nil, errors.New("Error creating transaction " + Transaction.Code)
	}

        err = stub.PutState(userId, accountBytes) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
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
