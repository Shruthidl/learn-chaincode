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
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var key string ;
var value string ;

type loc struct  {
        requester_name string;
        beneficiary_name  string;
        amount  string;
        expiry_date string;
       	status string;
        advising_bank string;
        document_hash string;
        loc_filename string;
        contract_hash string;
        bol_hash string;
    }
   
var counter uint = 0;
var LOCs map[uint]*loc;


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
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
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	
	fmt.Println("saving state for key: " + key);
	
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Adding LOCs 
func (t *SimpleChaincode) addLoc(stub *shim.ChaincodeStub,requester_name, beneficiary_name, amount, expiry_date, document_hash,loc_filename, contract_hash,  bol_hash string) ([]byte,error){
  
     counter = counter+1;
     LOCs[counter] =  &loc{};
     LOCs[counter].requester_name = requester_name;
     LOCs[counter].beneficiary_name= beneficiary_name;
     LOCs[counter].amount= amount;
     LOCs[counter].expiry_date= expiry_date;
     LOCs[counter].status= "requested";
     LOCs[counter].advising_bank = "none";
     LOCs[counter].document_hash= document_hash;
     LOCs[counter].loc_filename= loc_filename;
     LOCs[counter].contract_hash= contract_hash;
     LOCs[counter].bol_hash = bol_hash ;
     //LOCs [counter]= loc{requester_name,beneficiary_name,amount,expiry_date,"requested","none",document_hash,loc_filename,contract_hash, bol_hash};
     fmt.Println(LOCs [counter].requester_name);
     return nil, nil;

}


// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	
	fmt.Println("retrieving state for key: " + key);
	
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

        //valAsbytes = []byte(valAsbytes);
	return valAsbytes, nil
}
