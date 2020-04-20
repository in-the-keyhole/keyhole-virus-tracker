package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	//	"bytes"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the LabResult.  Structure tags are used by encoding/json library
type WeatherResult struct {
	Lon      string `json:"lon"`
	Lat      string `json:"lat"`
	City     string `json:"city"`
	State    string `json:"state"`
	Temp     string `json:"temp"`
	Humidity string `json:"humidity"`
	Wind     string `json:"wind"`
	Pressure string `json:"pressure"`
	DateTime string `json:datetime"`
}

/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "queryAll" {
		return s.queryAll(APIstub)
	} else if function == "queryStateResults" {
		return s.queryStateResults(APIstub, args[0])
	} else if function == "create" {
		return s.create(APIstub, args)
	}

	fmt.Println("args ", args)

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	entries := []LabResult{
		LabResult{Gender: "M", DateTime: "2018:07:03:10:00", DOB: "08/01/1966", TestType: "nasal swab", Result: "A", City: "Leawood", State: "KS"},
		LabResult{Gender: "F", DateTime: "2018:07:03:10:00", DOB: "08/01/1972", TestType: "nasal swab", Result: "B", City: "Leawood", State: "KS"},
		LabResult{Gender: "M", DateTime: "2018:07:03:10:00", DOB: "08/01/1966", TestType: "nasal swab", Result: "A", City: "Georgetown", State: "SC"},
		LabResult{Gender: "M", DateTime: "2018:07:03:10:00", DOB: "08/01/1980", TestType: "mouth swab", Result: "A", City: "St. Louis", State: "MO"},
		LabResult{Gender: "F", DateTime: "2018:07:03:10:00", DOB: "09/01/1980", TestType: "mouth swab", Result: "B", City: "Everglades City", State: "FL"},
	}

	i := 0
	for i < len(entries) {
		fmt.Println("i is ", i)
		tsAsBytes, _ := json.Marshal(entries[i])
		key := time.Now().UnixNano()
		APIstub.PutState("LAB"+strconv.FormatInt(key, 10), tsAsBytes)
		fmt.Println("Added", entries[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) queryAllEntries(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := ""
	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllLabs:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryStateResults(APIstub shim.ChaincodeStubInterface, state string) sc.Response {

	startKey := ""
	endKey := ""

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- Query KS Labs:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) createLab(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	var labResult = LabResult{Gender: args[0], DOB: args[1], City: args[2], State: args[3], TestType: args[4], Result: args[5], DateTime: args[6]}

	key := time.Now().UnixNano()
	labAsBytes, _ := json.Marshal(labResult)
	APIstub.PutState("LAB"+strconv.FormatInt(key, 10), labAsBytes)

	fmt.Println("***Added Labs")

	return shim.Success([]byte("Lab Created"))
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
