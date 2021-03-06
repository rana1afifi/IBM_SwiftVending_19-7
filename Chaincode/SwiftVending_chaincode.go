package main

import (
	"encoding/json"
	"errors"
	"fmt"
//	"strconv"
//	"time"
//	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

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
	username            string   `json:"username"`
	assetname          string   `json:"assetname"`
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
	} else if function == "Update" {
		return t.Update(stub , args)
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
	
         } else if function == "GetHistory"{
		return t.GetHistory(args[0] , stub) 
		/*fmt.Println("Getting all History")
		allTrans, err := GetHistory(args[0], stub)
		if err != nil {
			fmt.Println("Error from getHistory")
			return nil, err
		                 
		} else {
			allTransBytes, err1 := json.Marshal(&allTrans)
			if err1 != nil {
				fmt.Println("Error marshalling allTrans")
				return nil, err1}
			fmt.Println("All success, returning allTrans")
			return allTransBytes, nil} */
	}
        fmt.Println("query did not find func: " + function)
	return nil, errors.New("Received unknown function query: " + function)
}

// CreateTransaction to add asset to certain userID
func (t *SimpleChaincode) CreateTransaction(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//var userId, assetId string
	var err error
	fmt.Println("running write()")
/*
       if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
	userId = args[0] 
	assetId= args[1]*/
	//var assetIds []string
        //var transaction_arr []string
	 
	//var trans Transaction
	//TRIAL COMMENT STARTS HERE 
	/* trans:=Transaction{assetname: assetId, username:userId, Date:"Monday" , Time:11}
        transactionBytes, err := json.Marshal(&trans)
	if err != nil {
		fmt.Println("error creating transaction" + trans.username) // add transaction code later 
		return nil, errors.New("Error creating transaction "+trans.username)
	}
	
	fmt.Println("Attempting to get state of any existing transaction for " + trans.username)
	existingBytes, err := stub.GetState(trans.username)
       if err == nil {
		var company Transaction
		err = json.Unmarshal(existingBytes, &company)
		if err != nil {
			fmt.Println("Error unmarshalling account "  + err.Error())
*/
        //err = stub.PutState(userId, transactionBytes) //write the variable into the chaincode state
	err = stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		fmt.Println("failed to create create transaction")
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


func (t *SimpleChaincode)GetHistory(  username string , stub shim.ChaincodeStubInterface) pb.response {

	var history string

	// Get list of all the keys
	resultsIterator, err := stub.GetHistoryForKey(username)
	if err != nil {
		fmt.Println("Error retrieving history")
		return nil, errors.New("Error retrieving history")
	}
	//return resultsIterator , nil
	defer resultsIterator.Close()
	
	
	for resultsIterator.HasNext() {
		txID, historicValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		                           //copy transaction id over
		json.Unmarshal(historicValue, &history)     //un stringify it aka JSON.parse()
		if historicValue == nil {                  //marble has been deleted
			var emptyMarble string
			//tx.Value = emptyMarble                 //copy nil marble
		} else {
			json.Unmarshal(historicValue, &history) //un stringify it aka JSON.parse()
			//tx.Value = marble                      //copy marble over
		}
		//history = append(history, tx)              //add this tx to the list
	}
	fmt.Printf("- getHistoryForMarble returning:\n%s", history)

	//change to array of bytes
	historyAsBytes, _ := json.Marshal(history)     //convert to array of bytes
	return historyAsBytes , nil
}
	
	


func (t *SimpleChaincode) Update (stub shim.ChaincodeStubInterface, args[]string) ([]byte, error){

itemsBytes, err := stub.GetState(args[0])
	if err != nil {
		fmt.Println("Error retrieving history")
		return nil, errors.New("Error retrieving history")
	              }

    var items string=""
	err = json.Unmarshal(itemsBytes, &items)
	if err != nil {
		fmt.Println("Error unmarshalling item keys")
		return nil, errors.New("Error unmarshalling item keys")
	}
    items+=args[1]
    err = stub.PutState(args[0], []byte(items))
    
    
return nil , nil 
}
