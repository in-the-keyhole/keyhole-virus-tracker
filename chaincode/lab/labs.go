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
type LabResult struct {
	Status   string `json:status`
	Source   string `json:"source"`
	Gender   string `json:"gender"`
	DOB      string `json:"dob"`
	City     string `json:"city"`
	County   string `json:"county"`
	State    string `json:"state"`
	TestType string `json:"testtype"`
	Result   string `json:result"`
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
	} else if function == "initCovidLedger" {
		return s.initCovidLedger(APIstub)
	} else if function == "queryAllLabs" {
		return s.queryAllEntries(APIstub)
	} else if function == "queryStateResults" {
		return s.queryStateResults(APIstub, args[0])
	} else if function == "createLab" {
		return s.createLab(APIstub, args)
	}

	fmt.Println("args ", args)

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	entries := []LabResult{
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2019-08-01 10:00", DOB: "01/04/1980", TestType: "nasal swab", Result: "A", City: "Leawood", County: "Johnson", State: "KS", Status: "active"},
		LabResult{Source: "questlabs.com", Gender: "F", DateTime: "2019-08-30 10:00", DOB: "05/01/1975", TestType: "nasal swab", Result: "B", City: "Leawood", County: "Johnson", State: "KS", Status: "deceased"},
		LabResult{Source: "abclabs.com", Gender: "M", DateTime: "2019-09-03 10:00", DOB: "07/02/1945", TestType: "nasal swab", Result: "A", City: "Georgetown", County: "Georgetown", State: "SC", Status: "recovered"},
		LabResult{Source: "abclabs.com", Gender: "M", DateTime: "2019-09:04 10:00", DOB: "08/17/1963", TestType: "mouth swab", Result: "NEGATIVE", City: "St. Louis", County: "St. Louis", State: "MO", Status: "none"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2019-09-015 10:00", DOB: "04/15/1967", TestType: "mouth swab", Result: "NEGATIVE", City: "Everglades City", County: "Collier", State: "FL", Status: "none"},
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

func (s *SmartContract) initCovidLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	entries := []LabResult{
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-03 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Albany", County: "Albany", State: "NY", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-03 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Las Vegas", County: "Clark", State: "NV", Status: "None"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-03 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-05 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-05 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "recovered"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-13 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-13 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-23 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "recovered"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-23 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-01-23 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-03 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-03 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-03 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "recovered"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-13 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-13 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Boise", County: "Boise", State: "ID", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-15 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Batton Rouge", County: "Batton", State: "LA", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-15 15:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Isle", County: "Isle", State: "ME", Status: "deceased"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Dallas", County: "Dallas", State: "TX", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Austin", County: "Austin", State: "TX", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "St. Louis", County: "St. Louis", State: "MO", Status: "negative"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Lincoln", County: "Lincoln", State: "NE", Status: "recovered"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Banks", County: "Banks", State: "NH", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Walter", County: "Walter", State: "MT", Status: "active"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Aspen", County: "Aspen", State: "CO", Status: "none"},
		LabResult{Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", DOB: "01/04/1980", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Walter", County: "Walter", State: "NM", Status: "deceased"},
		LabResult{Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", DOB: "07/04/1976", TestType: "sputum sample", Result: "NEGATIVE", City: "Norse", County: "Clark", State: "SD", Status: "none"},
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

	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	status := "negative"

	if args[6] != "negative" {

		status = "active"

	}

	var labResult = LabResult{Gender: args[0], DOB: args[1], City: args[2], County: args[3], State: args[4], TestType: args[5], Result: args[6], DateTime: args[7], Source: args[8], Status: status}

	key := time.Now().UnixNano()
	labAsBytes, _ := json.Marshal(labResult)
	APIstub.PutState("LAB"+strconv.FormatInt(key, 10), labAsBytes)

	fmt.Println("*** Added Labs")
	fmt.Println(APIstub.GetCreator())

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
