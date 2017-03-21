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
	"bytes"
	

	
        "github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var key string ;
var value string ;


   
var counter int = 0;
var txncounter int = 0;
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
	} else if function == "addInClearFile" {
		fmt.Println("**** First argument in addInClearFile:****" + args[0])
		return t.addInClearFile(stub, args)
	} else if function == "markTxnCleared" {
		fmt.Println("**** First argument in markTxnCleared:****" + args[0])
		return t.markTxnCleared(stub, args)
	} else if function == "markFilesCleared" {
		fmt.Println("**** First argument in markFilesCleared:****" + args[0])
		return t.markFilesCleared(stub, args)
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
	} else if function == "getFiles" {
	       return t.getFiles(stub, args);
	} else if function == "getAlltxns" {
	       return t.getAlltxns(stub, args);
	} else if function == "getCurrentFileId" {
	       return t.getCurrentFileId(stub, args);
	} else if function == "getCounts" {
	       return t.getCounts(stub, args);
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
func (t *SimpleChaincode) addOutClearFile(stub shim.ChaincodeStubInterface, args []string) ([]byte ,error){
  var err error;
  var counter1 int;
  var stringslice []string;
  
	//prepareData
	err = stub.PutState("364924",[]byte("City Bank - 130"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("364914",[]byte("I Bank - 120"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("364927",[]byte("My Bank - 140"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("4321432100",[]byte("DCB Bank - 25"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("1234123400",[]byte("Src Bank - 29"))
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
	acq_name,err :=stub.GetState(args[0])
	stringslice = append(stringslice,args[2],args[1],args[3],args[4],string(acq_name),"In Process");
	stringvalues = append(stringslice,counter_s,counter_s);//string array (value)
        s_requester := counter_s //counter value(key)
          	
	     var mCount int = len(stringvalues);
                parts  := make([]string, mCount );
                parts = stringvalues;
	
		 if(!strings.HasPrefix(args[5] , "H-")){
			
			 parts[5] = "Rejected";

			stringBytes1 := strings.Join(parts, "|") 

			err = stub.PutState(s_requester, []byte(stringBytes1));
			if err != nil {
				return nil, err
			}
			     return nil, nil;
		}

		if(!strings.HasPrefix(args[6] , "T-")){
			
			 parts[5] = "Rejected";

			stringBytes2 := strings.Join(parts, "|") 

			err = stub.PutState(s_requester, []byte(stringBytes2));
			if err != nil {
				return nil, err
			}
			 return nil, nil;
		}
	
	
	        parts[5] = "Validated";
		stringBytes := strings.Join(parts, "|") 

		err = stub.PutState(s_requester, []byte(stringBytes));
			if err != nil {
				return nil, err
			}
	
		
	
	//enrich data
	content := strings.Split(args[7], ",");
	var m int = len(content);
          cont  := make([]string, m );
          cont = content;
	
  for i := 0; i < len(cont); i++ {
	 txncounter = txncounter + 1;
	 fmt.Println(cont[i]);
	 content1:=strings.Split(cont[i],"|")  
         var mCount1 int = len(content1);
         parts1  := make([]string, mCount1 );
         parts1 = content1;
		
		
        var buffer bytes.Buffer
		
	 buffer.WriteString(strconv.Itoa(txncounter));
          fmt.Println(buffer);
          buffer.WriteString("|");
          buffer.WriteString(strconv.Itoa(counter));
          buffer.WriteString("|");
          buffer.WriteString(cont[i]);	
	  status := "Validated|20-01-2017 07:20AM";	
	 if(strings.HasPrefix(parts1[0], "1240")){
         status = "Validated|20-01-2017 07:20AM";
	 }else{
          status = "Invalid|20-01-2017 07:20AM";
	 }
   	 fmt.Println(status);
	card := "364924";	
	if(strings.HasPrefix(parts1[0], "364924")){
           card = "364924";
		
        } else if(strings.HasPrefix(parts1[0], "364914")){
          card = "364914";
		
        } else if(strings.HasPrefix(parts1[0], "364927")){
           
          card = "364927";
        }
	
               fmt.Println(card);
		 cardname,err :=stub.GetState(card)
		
		buffer.WriteString(string(cardname));
	        buffer.WriteString((parts1[0]));
		buffer.WriteString("|0.25|");
		buffer.WriteString(status);
		
		err = stub.PutState("t"+strconv.Itoa(txncounter), []byte(buffer.String()));	
	     
		if err != nil {
			return nil, err
			}
	}
	     
	  
	
              return nil, nil
}
	

// Adding InClear files 
func (t *SimpleChaincode) addInClearFile(stub shim.ChaincodeStubInterface, args []string) ([]byte ,error){
  var err error;
  var counter2 int;
  var stringslice1 []string;
  
	//prepareData
	err = stub.PutState("364924",[]byte("City Bank - 130"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("364914",[]byte("I Bank - 120"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("364927",[]byte("My Bank - 140"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("4321432100",[]byte("DCB Bank - 25"))
	if err != nil {
		return nil, err
	}
	err = stub.PutState("1234123400",[]byte("Src Bank - 29"))
	if err != nil {
		return nil, err
	}
	
	valAsbytes,err :=stub.GetState(strconv.Itoa(counter))
        s:=string(valAsbytes);
	
        if len(s) != 0 {
	     lastByByte := s[len(s)-1:]
             counter2, err =  strconv.Atoi(lastByByte)
 		if err != nil {
     			return  nil,err
  	         }
	
   	  } else {
             counter2 = 0
    	   }
   
	counter = counter2+1;
	
	counter_s := strconv.Itoa(counter)
	acq_name,err :=stub.GetState(args[0])
	stringslice1 = append(stringslice1,args[2],args[1],args[3],args[4],string(acq_name),"In Process");
	
	stringvalues = append(stringslice1,counter_s,counter_s);//string array (value)
        s_requester := counter_s //counter value(key)
       
		stringBytes := strings.Join(stringvalues, "|") 

		err = stub.PutState(s_requester, []byte(stringBytes));
			if err != nil {
				return nil, err
			}
	    
	  return nil, nil
}



// mark transaction cleared
func (t *SimpleChaincode) markTxnCleared(stub shim.ChaincodeStubInterface, args []string) ([]byte,error){
        var s = args;
        var mCount int = len(s);
        parts  := make([]string, mCount );
   	parts = s;
        for j := 0; j < len(parts); j++ {
           
		valueAsBytes , err := stub.GetState("t"+parts[j]);
	if err != nil {
	 return nil,err	
	}
 	 var str bytes.Buffer;
		//str.WriteString(string(valueAsBytes));
		str.WriteString("1|2|1240|364914020023481|123.00|Gloria Jeans-CH|4321432100|City Bank - 1301240|0.25|Validated|20-01-2017 07:20AM");
                str.WriteString("|Cleared");
                err = stub.PutState("t"+parts[j], []byte(str.String()));	
	     
		if err != nil {
			return nil, err
			}
        }
	return nil, nil
    }


//mark files cleared
func (t *SimpleChaincode) markFilesCleared(stub shim.ChaincodeStubInterface, args []string) ([]byte,error){
	var s = args;
        var mCount int = len(s);
        parts  := make([]string, mCount );
   	parts = s;
         
	for j := 0; j < len(parts); j++ {
		valueAsBytes , err := stub.GetState(parts[j]);
	if err != nil {
	 return nil,err	
	}
	s1:=string(valueAsBytes);
	var s2 = strings.Split(s1, "|");;
        var mCount int = len(s2);
        parts1  := make([]string, mCount );
   	parts1 = s2;
	parts1[5] = "Cleared";	 
        stringbyte := strings.Join(parts1, "|") 
		 err = stub.PutState(parts[j],[]byte(stringbyte));	
	     
		if err != nil {
			return nil, err
			}
        }
        	return nil, nil
    }


// Return all files
    func (t *SimpleChaincode) getFiles (stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {
     
    	var list []string;
	
	for i := 1; i <=counter; i++ {
	 valueAsBytes , err := stub.GetState(strconv.Itoa(i));
	if err != nil {
	 return nil,err	
	}
	  s:=string(valueAsBytes);
	  var s2 = strings.Split(s, "|");
       
          parts1  := make([]string, len(s2)+2);
	  parts1[0] = strconv.Itoa(i)
	  copy(parts1[1:], s2)
	  parts1[1] = strconv.Itoa(i)
	  copy(parts1[2:], s2)
   	  s2 = parts1;
	  s2= append(s2[:9], s2[9+1:]...)	
	  s2= append(s2[:8], s2[8+1:]...)	
	  s = strings.Join(s2, "|") 
	  list =append(list,s);
	}

	listByte := strings.Join(list, ",");
	
	return []byte(listByte), nil;
        
    }

// Return all transactions
    func (t *SimpleChaincode) getAlltxns (stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {
     
    	var list []string;
	
	for i := 1; i <=txncounter; i++ {
	 valueAsBytes , err := stub.GetState("t"+strconv.Itoa(i));
	if err != nil {
	 return nil,err	
	}
	  s:=string(valueAsBytes);
	  list =append(list,s);
	}

	txnsByte := strings.Join(list, ",");
	
	return []byte(txnsByte), nil;
        
    }


// Return getCurrentFileId
   func (t *SimpleChaincode) getCurrentFileId (stub shim.ChaincodeStubInterface, args []string)([]byte,error){
	   
	   return []byte(strconv.Itoa(counter)),nil;
}


// Return count
func (t *SimpleChaincode) getCounts(stub shim.ChaincodeStubInterface, args []string)([]byte,error){
       
	        var str bytes.Buffer;
		str.WriteString("Files:");
                str.WriteString(strconv.Itoa(counter));
	        str.WriteString(",");
	        str.WriteString("Txns:");
                str.WriteString(strconv.Itoa(txncounter));
	        return []byte(str.String()),nil;
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

