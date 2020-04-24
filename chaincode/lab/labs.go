package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	//	"bytes"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

func IDGen(lab string) string {

	//nano := time.Now().UnixNano()
	//snano := strconv.FormatInt(nano, 10)
	//id := sha256.Sum256([]byte(snano + lab))
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	// fmt.Println(uuid)

	//return string(id[:])
	return uuid

}

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the LabResult.  Structure tags are used by encoding/json library
type LabResult struct {
	ID       string `json:id`
	Status   string `json:status`
	Source   string `json:"source"`
	Gender   string `json:"gender"`
	Age      string `json:"age"`
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
	} else if function == "recovered" {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1 UUID")
		}
		return s.changeStatus(APIstub, args[0], "recovered")
	} else if function == "deceased" {
		if len(args) != 1 {
			return shim.Error("Incorrect number of arguments. Expecting 1 UUID")
		}
		return s.changeStatus(APIstub, args[0], "deceased")
	} else if function == "queryById" {

		return s.queryByID(APIstub, args)

	}

	fmt.Println("args ", args)

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	entries := []LabResult{
		LabResult{ID: IDGen("questlabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2019-08-01 10:00", Age: "33", TestType: "nasal swab", Result: "A", City: "Leawood", County: "Johnson", State: "KS", Status: "active"},
		LabResult{ID: IDGen("questlabs.com"), Source: "questlabs.com", Gender: "F", DateTime: "2019-08-30 10:00", Age: "40", TestType: "nasal swab", Result: "B", City: "Leawood", County: "Johnson", State: "KS", Status: "deceased"},
		LabResult{ID: IDGen("questlabs.com"), Source: "abclabs.com", Gender: "M", DateTime: "2019-09-03 10:00", Age: "55", TestType: "nasal swab", Result: "A", City: "Georgetown", County: "Georgetown", State: "SC", Status: "recovered"},
		LabResult{ID: IDGen("questlabs.com"), Source: "questlabs.com", Gender: "F", DateTime: "2019-08-30 10:00", Age: "40", TestType: "nasal swab", Result: "B", City: "Leawood", County: "Johnson", State: "KS", Status: "deceased"},
		LabResult{ID: IDGen("questlabs.com"), Source: "questlabs.com", Gender: "F", DateTime: "2019-08-30 10:00", Age: "40", TestType: "nasal swab", Result: "B", City: "Leawood", County: "Johnson", State: "KS", Status: "deceased"},
		LabResult{ID: IDGen("questlabs.com"), Source: "questlabs.com", Gender: "F", DateTime: "2019-08-30 10:00", Age: "40", TestType: "nasal swab", Result: "B", City: "Leawood", County: "Johnson", State: "KS", Status: "deceased"},
		LabResult{ID: IDGen("abclabs.com"), Source: "abclabs.com", Gender: "M", DateTime: "2019-09:04 10:00", Age: "57", TestType: "mouth swab", Result: "NEGATIVE", City: "St. Louis", County: "St. Louis", State: "MO", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2019-09-015 10:00", Age: "75", TestType: "mouth swab", Result: "NEGATIVE", City: "Everglades City", County: "Collier", State: "FL", Status: "none"},
	}

	i := 0
	for i < len(entries) {
		fmt.Println("i is ", i)
		tsAsBytes, _ := json.Marshal(entries[i])
		APIstub.PutState(IDGen(entries[i].ID), tsAsBytes)
		fmt.Println("Added", entries[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) initCovidLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	entries := []LabResult{
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-03 10:00", Age: "25", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Albany", County: "Albany", State: "NY", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-03 10:00", Age: "35", TestType: "sputum sample", Result: "NEGATIVE", City: "Las Vegas", County: "Clark", State: "NV", Status: "None"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-03 10:00", Age: "45", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-05 10:00", Age: "47", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-05 10:00", Age: "55", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "recovered"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-13 10:00", Age: "60", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-13 10:00", Age: "62", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", Age: "64", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-23 10:00", Age: "40", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "recovered"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", Age: "41", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-23 10:00", Age: "42", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", Age: "43", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-01-23 10:00", Age: "47", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-01-23 10:00", Age: "80", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-03 10:00", Age: "81", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Hollwood", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-03 10:00", Age: "82", TestType: "sputum sample", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "deceased"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-03 10:00", Age: "85", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Los Angeles", County: "Los Angeles", State: "CA", Status: "recovered"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-13 10:00", Age: "15", TestType: "sputum sample", Result: "NEGATIVE", City: "Los Angeles", County: "Clark", State: "NV", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-13 10:00", Age: "28", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Boise", County: "Boise", State: "ID", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-15 10:00", Age: "29", TestType: "sputum sample", Result: "NEGATIVE", City: "Batton Rouge", County: "Batton", State: "LA", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-15 15:00", Age: "30", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Isle", County: "Isle", State: "ME", Status: "deceased"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", Age: "31", TestType: "sputum sample", Result: "NEGATIVE", City: "Dallas", County: "Dallas", State: "TX", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", Age: "38", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Austin", County: "Austin", State: "TX", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", Age: "40", TestType: "sputum sample", Result: "NEGATIVE", City: "St. Louis", County: "St. Louis", State: "MO", Status: "negative"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", Age: "50", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Lincoln", County: "Lincoln", State: "NE", Status: "recovered"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", Age: "59", TestType: "sputum sample", Result: "NEGATIVE", City: "Banks", County: "Banks", State: "NH", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", Age: "60", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Walter", County: "Walter", State: "MT", Status: "active"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", Age: "44", TestType: "sputum sample", Result: "NEGATIVE", City: "Aspen", County: "Aspen", State: "CO", Status: "none"},
		LabResult{ID: IDGen("abslabs.com"), Source: "questlabs.com", Gender: "M", DateTime: "2020-02-20 10:00", Age: "45", TestType: "nasopharyngeal swab", Result: "POSITIVE", City: "Walter", County: "Walter", State: "NM", Status: "deceased"},
		LabResult{ID: IDGen("abslabs.com"), Source: "abclabs.com", Gender: "F", DateTime: "2020-02-20 10:00", Age: "55", TestType: "sputum sample", Result: "NEGATIVE", City: "Norse", County: "Clark", State: "SD", Status: "none"},
	}

	i := 0
	for i < len(entries) {
		fmt.Println("i is ", i)
		tsAsBytes, _ := json.Marshal(entries[i])
		APIstub.PutState(entries[i].ID, tsAsBytes)
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

func (s *SmartContract) queryByID(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 UUID")
	}

	key := args[0]

	result, err := APIstub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(result)

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

func (s *SmartContract) changeStatus(APIstub shim.ChaincodeStubInterface, id string, status string) sc.Response {

	byt, err := APIstub.GetState(id)

	if err != nil {
		return shim.Error("Lab ID " + id + " not found")
	}
	var labResult LabResult
	if err := json.Unmarshal(byt, &labResult); err != nil {
		panic(err)
	}

	// change status to recovered
	labResult.Status = status
	labAsBytes, _ := json.Marshal(labResult)
	APIstub.PutState(id, labAsBytes)

	return shim.Success([]byte("Lab Found Marked as " + labResult.Status))
}

func (s *SmartContract) createLab(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	status := "negative"

	if args[6] != "negative" {

		status = "active"

	}

	id := IDGen(args[8])
	var labResult = LabResult{Gender: args[0], Age: args[1], City: args[2], County: args[3], State: args[4], TestType: args[5], Result: args[6], DateTime: args[7], Source: args[8], Status: status}
	labAsBytes, _ := json.Marshal(labResult)
	APIstub.PutState(id, labAsBytes)

	fmt.Println("*** Added Labs")
	fmt.Println(APIstub.GetCreator())

	return shim.Success(labAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
