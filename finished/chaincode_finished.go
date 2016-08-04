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
	"strconv"
	"strings"

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
   
var counter int = 0;
var LOCs map[int]*loc;


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
	} else if function == "addLoc" {
		fmt.Println("**** First argument in addLoc:****" + args[0])
		return t.addLoc(stub, args)
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
	} else if function == "getLoc" {
	//	i,err := strconv.Atoi(args[0])
	//	fmt.Println(err); 
		return t.getLoc(stub, args);
		 
	} else if function == "getNumberOfLocs" {
	
		return t.getNumberOfLocs(stub, args);
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
func (t *SimpleChaincode) addLoc(stub *shim.ChaincodeStub, args []string) ([]byte,error){
  var err error;
   
  
    valAsbytes, err := stub.GetState("counter");
    
     counter, err = strconv.Atoi(string(valAsbytes));
  
     counter = counter+1;
    
     //fmt.Println(LOCs [counter].requester_name);
     
    counter_s := strconv.Itoa(counter) ;
    counter_b := []byte(counter_s);
     
     //err = stub.PutState(counter_s, []byte(LOCs[counter].requester_name)) //write the variable into the chaincode state
     //	if err != nil {
    // 		return nil, err
    //	}

    //fmt.Println(counter_b);

    
     err = stub.PutState("counter",counter_b);

     if err != nil {
		return nil, err
	}
	
   //-----------------------------------------------------------------------	
    s_requester := []string{counter_s, "requester"};
    s1 := strings.Join(s_requester, "_");
    
    /* Start
	
     err = stub.PutState(s1,[]byte(args[0]));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------	
    s_beneficiary := []string{counter_s, "beneficiary"};
    s1 = strings.Join(s_beneficiary, "_");
    
	
     err = stub.PutState(s1,[]byte(args[1]));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------
    s_amount := []string{counter_s, "amount"};
    s1 = strings.Join(s_amount, "_");
    
	
     err = stub.PutState(s1,[]byte(args[2]));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------
    s_expiry_date := []string{counter_s, "expiry_date"};
    s1 = strings.Join(s_expiry_date, "_");
    
	
     err = stub.PutState(s1,[]byte(args[3]));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------
   s_status := []string{counter_s, "status"};
    s1 = strings.Join(s_status, "_");
    
    err = stub.PutState(s1,[]byte("requested"));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------
  s_advising_bank := []string{counter_s, "advising_bank"};
    s1 = strings.Join(s_advising_bank, "_");
    
    err = stub.PutState(s1,[]byte("none"));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------
 s_document_hash := []string{counter_s, "document_hash"};
    s1 = strings.Join(s_document_hash, "_");
    
    err = stub.PutState(s1,[]byte(args[4]));

     if err != nil {
		return nil, err
	}
   //-----------------------------------------------------------------------
 s_loc_filename := []string{counter_s, "loc_filename"};
    s1 = strings.Join(s_loc_filename, "_");
    
    err = stub.PutState(s1,[]byte(args[5]));

     if err != nil {
		return nil, err
	}
	
     return nil, nil;
   end  */ 
  //-------------------------------------------------------------------------
   s_contract_hash := []string{counter_s, "contract_hash"};
    s1 = strings.Join(s_contract_hash, "_");
    
   // err = stub.PutState(s1,[]byte(args[6]));
    err = stub.PutState(s1,[]byte("test"));

     if err != nil {
		return nil, err
	}
	
     return nil, nil;
  //-------------------------------------------------------------------------
  s_bol_hash := []string{counter_s, "bol_hash"};
    s1 = strings.Join(s_bol_hash, "_");
    
   // err = stub.PutState(s1,[]byte(args[7]));
      err = stub.PutState(s1,[]byte("test2"));

     if err != nil {
		return nil, err
	}
	
     return nil, nil;
  //-------------------------------------------------------------------------
  
}

// Return specific LOC in the system
    func (t *SimpleChaincode) getLoc(stub *shim.ChaincodeStub , args []string) ([]byte,error) {
     
    	 
    	s := []string{args[0], "requester"};
        s1 := strings.Join(s, "_");
    	 
         
        requester_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
    	//------------------------------------------------------------
    	s = []string{args[0], "beneficiary"};
        s1 = strings.Join(s, "_");
    	 
         
        beneficiary_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
    	s = []string{args[0], "amount"};
        s1 = strings.Join(s, "_");
    	 
         
        amount_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
	s = []string{args[0], "expiry_date"};
        s1 = strings.Join(s, "_");
    	 
         
        expiry_date_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
	s = []string{args[0], "status"};
        s1 = strings.Join(s, "_");
    	 
         
        status_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
	s = []string{args[0], "advising_bank"};
        s1 = strings.Join(s, "_");
    	 
         
        advising_bank_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
	s = []string{args[0], "document_hash"};
        s1 = strings.Join(s, "_");
    	 
         
        document_hash_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
	s = []string{args[0], "loc_filename"};
        s1 = strings.Join(s, "_");
    	 
         
        loc_filename_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
    	s = []string{args[0], "contract_hash"};
        s1 = strings.Join(s, "_");
    	 
         
        contract_hash_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
	s = []string{args[0], "bol_hash"};
        s1 = strings.Join(s, "_");
    	 
         
        bol_hash_string, err := stub.GetState(s1);
    	
    	if err != nil {
		return nil, err
	}
	//--------------------------------------------------------------
    	
    	s = []string{string(requester_string),string(beneficiary_string),string(amount_string),string(expiry_date_string),string(status_string),string(advising_bank_string),string(document_hash_string),string(loc_filename_string),string(contract_hash_string),string(bol_hash_string)};
        
        s=[]string{string(contract_hash_string),string(bol_hash_string)};
        final_string := strings.Join(s, "|");
    	
    	
    	
        //s := strconv.Itoa(counter) ;
        //ret_s := []byte(s);
        return []byte(final_string), nil;
        
    }


 //Get number of LOCs in the system
    func (t *SimpleChaincode) getNumberOfLocs (stub *shim.ChaincodeStub, args []string) ([]byte, error){
    	valAsbytes, err := stub.GetState("counter");
    	
    	if err != nil {
		return nil, err
	}
    	
        //s := strconv.Itoa(counter) ;
        //ret_s := []byte(s);
        return valAsbytes, nil;
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
