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
		return t.addLoc(stub, args[0],args[1],args[2],args[3],args[4],args[5],args[6],args[7])
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
		i,err := strconv.Atoi(args[0])
		fmt.Println(err); 
		return t.getLoc(stub, i)
		 
	} else if function == "getNumberOfLocs" {
		var m_key string;
		m_key := "counter";
		return t.getNumberOfLocs(stub, [m_key]);
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
  var err error;
  
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
     
     counter_s := strconv.Itoa(counter) ;
     counter_b := []byte(counter_s);
     
     err = stub.PutState(counter_s, []byte(LOCs[counter].requester_name)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}


	
     err = stub.PutState("counter",counter_b);

     if err != nil {
		return nil, err
	}
	
	
     return nil, nil;

}

// Return specific LOC in the system
    func (t *SimpleChaincode) getLoc(stub *shim.ChaincodeStub , location int) ([]byte,error) {
	b := make([]byte, 300)
             tracker:= 0;

        
      	for  i:=0 ; i<len(LOCs [location].requester_name) ; i++{
 		
		//fmt.Println(LOCs[location].requester_name[i]);
		b[tracker] = LOCs[location].requester_name[i] ;
	        tracker = tracker + 1;
         
           }

		  tracker = tracker + 1;
 	for  j:=0 ; j<len(LOCs [location].beneficiary_name) ; j++{
 		
		//fmt.Println(LOCs[location].beneficiary_name[j]);
		b[tracker] = LOCs[location].beneficiary_name[j] ;
	        tracker = tracker + 1;
         
           }
      
		  tracker = tracker + 1;

	 for  k:=0 ; k<len(LOCs [location].amount) ; k++{
 		
		//fmt.Println(LOCs[location].amount[k]);
		b[tracker] = LOCs[location].amount[k] ;
	        tracker = tracker + 1;
         
           }
		
		  tracker = tracker + 1;
 	for  l:=0 ; l<len(LOCs [location].expiry_date) ; l++{
 		
		//fmt.Println(LOCs[location].expiry_date[l]);
		b[tracker] = LOCs[location].expiry_date[l] ;
	        tracker = tracker + 1;
         
           }

		  tracker = tracker + 1;
 	for m:= 0; m <len(LOCs [location].status) ; m++{
            b[tracker] = LOCs[location].status[m] ;
            tracker = tracker + 1;
        }

	  tracker = tracker + 1;
 	for n:= 0; n <len(LOCs [location].advising_bank) ; n++{
            b[tracker] = LOCs[location].advising_bank[n] ;
            tracker = tracker + 1;
        }

	  tracker = tracker + 1;

	 for p:= 0; p <len(LOCs [location].document_hash); p++{
	     	//fmt.Println(LOCs[location].document_hash[p]);
            b[tracker] = LOCs[location].document_hash[p] ;
            tracker = tracker + 1;
        }
	
 	 tracker = tracker + 1;
		
	 for q:= 0; q <len(LOCs [location].loc_filename); q++{
	     	//fmt.Println(LOCs[location].loc_filename[q]);
            b[tracker] = LOCs[location].loc_filename[q] ;
            tracker = tracker + 1;
        }
		
	  tracker = tracker + 1;
		
	for r:= 0; r <len(LOCs [location].contract_hash); r++{
	     	//fmt.Println(LOCs[location].contract_hash[r]);
            b[tracker] = LOCs[location].contract_hash[r] ;
            tracker = tracker + 1;
        }

  	tracker = tracker + 1;
	
	for s:= 0; s <len(LOCs [location].bol_hash); s++{
	     	//fmt.Println(LOCs[location].bol_hash[s]);
            b[tracker] = LOCs[location].bol_hash[s] ;
            tracker = tracker + 1;
        }

		
        
                  return b, nil;
        
    }


 //Get number of LOCs in the system
    func (t *SimpleChaincode) getNumberOfLocs (stub *shim.ChaincodeStub, args []string) ([]byte, error){
    	valAsbytes, err := stub.GetState(args[0]);
    	
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
