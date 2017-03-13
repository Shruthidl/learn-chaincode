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


   
var counter int = 0;
var stringvalues []string;



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("abc", []byte(args[0]))
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
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "addOutClearFile" {
		fmt.Println("**** First argument in addOutClearFile:****" + args[0])
		return t.addOutClearFile(stub, args)
	} 
	
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}
// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} 
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	
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

// Adding OutClear files 
func (t *SimpleChaincode) addOutClearFile(stub shim.ChaincodeStubInterface, args []string) ([]byte,error){
  var err error;
  var counter1 int;
  
	//prepareData
	err = stub.PutState("364924",[]byte("|City Bank - 130"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("364914",[]byte("|I Bank - 120"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("364927",[]byte("|My Bank - 140"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("4321432100",[]byte("|DCB Bank - 25"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("1234123400",[]byte("|Src Bank - 29"))
	if err != nil {
		return nil, err
	}
	
	
	// add out clear files
    valAsbytes,err :=stub.GetState(strconv.Itoa(counter))
    s:=string(valAsbytes);
	
     if len(s) != 0 {
	     lastByByte := s[len(s)-1:]
             counter1, err =  strconv.Atoi(lastByByte)
 		if err != nil {
     			return  nil,err
  	         }
	
   	  } else {
             counter1 = 0
    	   }
   
     counter = counter1+1;
    
     counter_s := strconv.Itoa(counter)
     stringvalues = append(args,counter_s)//string array (value)
     s_requester := counter_s //counter value(key)

     stringByte := strings.Join(stringvalues , "|") // x00 = null
     
      err = stub.PutState(s_requester, []byte(stringByte));

      if err != nil {
		return nil, err
	}
	
	//precapture
	 value,err :=stub.GetState(strconv.Itoa(counter))
		if err != nil {
		return nil, err
	}
	
		
		s1 := strings.Split(value, "|");
	
	       var mCount int = len(s1);
               parts  := make([]string, mCount );
                parts = s1;
	         parts[5] = "Validated";
		stringBytes := strings.Join(parts, "|") 

		err = stub.PutState(args[0], []byte(stringBytes));
	
	     
	
               return nil, nil
}
	
 func (t *SimpleChaincode) getStatus(stub shim.ChaincodeStubInterface, arg string) ([]byte, error){
        if(strings.HasPrefix(arg, "1240")){
         
		  return []byte("Validated|20-01-2017 07:20AM"), nil;
        }
     
	     return []byte("Invalid|20-01-2017 07:20AM"), nil;
    }

 func (t *SimpleChaincode) getCard(stub shim.ChaincodeStubInterface, arg string) ([]byte, error){
        if(strings.HasPrefix(arg, "364924")){
           
		 return []byte("364924"), nil;
        } else if(strings.HasPrefix(arg, "364914")){
         
		 return []byte("364914"), nil;
        } else if(strings.HasPrefix(arg, "364927")){
           
		 return []byte("364927"), nil;
        }
       
	   return []byte("364999"), nil;
    }

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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


