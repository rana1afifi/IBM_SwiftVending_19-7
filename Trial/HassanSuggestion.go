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
	ItemName      string      `json:"ItemName"`
	QRCode         string    `json:"QRCode"`

}

type UserAccount struct 
{
    Username       string     `json:"Username"`
    Items          []string       `json:"Items"`
    
    
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
         } else if function == "GetItems"{
		return t.GetItems(stub , args) 
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
	}Fget
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


func (t *SimpleChaincode)GetHistory(  username string , stub shim.ChaincodeStubInterface) ([]byte, error) {

	var history string

	// Get list of all the keys
	itemsBytes, err := stub.GetState(username)
	if err != nil {
		fmt.Println("Error retrieving history")
		return nil, errors.New("Error retrieving history")
	}
	var items []string
	err = json.Unmarshal(itemsBytes, &items)
	if err != nil {
		fmt.Println("Error unmarshalling item keys")
		return nil, errors.New("Error unmarshalling item keys")
	}

	// Get all the cps
	for _, value := range items {
		cpBytes, err := stub.GetState(value)

		var tr string
		err = json.Unmarshal(cpBytes, &tr)
		if err != nil {
			fmt.Println("Error retrieving tr " + value)
			return nil, errors.New("Error retrieving tr " + value)
		}

		fmt.Println("Appending CP" + value)
		history += tr
	}

	return []byte(history), nil


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




func (t *SimpleChaincode) Buy(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) { 
    
    var err error
    var qrcode string
	fmt.Println("running write()")

    
     if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}
    
    // Generate Random Number 
        qrcode= args[2] // the QRCODE is already stored in cloudant 
    
    // Create Object 
    trans:=Transaction{Username:args[0], ItemName:args[1], QRCode: qrcode}
    
   //Convert it to JsonObject
    
    transactionBytes, err := json.Marshal(&trans)
	if err != nil {
		fmt.Println("error creating transaction" + trans.Username) // add transaction code later 
		return nil, errors.New("Error creating transaction "+trans.Username)
	        
    } else {   //Store this Transaction into the database 
        
        err = stub.PutState(qrcode, transactionBytes)
    }
    
  
	
	existingBytes, err := stub.GetState(trans.Username)  // or args[0]
    // if error ==nil : the user does have an account 
      
    if err == nil {
		   var account UserAccount
	    err2 := json.Unmarshal(existingBytes, &account)
        // unmarshal bytes in order to append 
            if err2==nil  {
			fmt.Println("Error unmarshalling account "  + err2.Error())
            return nil, errors.New("Error  updating account "+trans.Username)
           
            } else {   // update array of items 
         
            account.Items= append(account.Items, qrcode)
	    accountInBytes,err:=json.Marshal(account)
		     if err==nil  {
			fmt.Println("Error marshalling account "  + err2.Error())
                        return nil, errors.New("Error  updating account "+trans.Username)
                      }
		    
            err = stub.PutState(args[0], accountInBytes)
        }
// if account doesn't exist
    } else {
           var qrarray []string 
	   qrarray=append(qrarray,qrcode)
          acc:=UserAccount{Username:args[0], Items: qrarray}
          accBytes, err := json.Marshal(&acc)
	     if err==nil  {
			fmt.Println("Error marshalling account "  + err2.Error())
                        return nil, errors.New("Error  updating account "+trans.Username)
                      }
          err = stub.PutState(args[0], accBytes)
    }
    
	return nil, nil
}
func (t *SimpleChaincode) GetItems(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) { 
    
       var err error
    // Get UserAccout 
     var transactions []Transaction
    
    accountBytes , err := stub.GetState(args[0]) 
    if err!=nil {
            fmt.Println("Error fetching account "  + err.Error())
            return nil, errors.New("Error  fetching account"+args[0])    
        
    }
    // unmarshal accountbytes and then loop 
    var qrCodes []string
    err=json.Unmarshal(accountBytes, &qrCodes)
    
    if err != nil {
		fmt.Println("Error unmarshalling account qr codes")
		return nil, errors.New("Error unmarshalling account qr codes")
	}
    
     //fetch the items of this account 
    for _, value := range qrCodes {
		transactionBytes, err := stub.GetState(value)

		var trans Transaction
		err = json.Unmarshal(transactionBytes, &trans)
		if err != nil {
			fmt.Println("Error retrieving transaction " + value)
			return nil, errors.New("Error retrieving transaction " + value)
		}

		fmt.Println("Appending CP" + value)
		transactions = append(transactions, trans)
    }
    
    resultTransBytes , err2:= json.Marshal(&transactions)
    
    if err2 != nil {
				fmt.Println("Error marshalling Transactions")
				return nil, err2
			}
			fmt.Println("All success, returning Transactions")
			return resultTransBytes, nil
	}

	
    
    
