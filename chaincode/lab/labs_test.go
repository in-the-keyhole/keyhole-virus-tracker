package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}

func TestExample02_Init(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("ex02", scc)

	// Init A=123 B=234
	checkInit(t, stub, [][]byte{[]byte("")})

	//checkState(t, stub, "A", "123")
	// checkState(t, stub, "B", "234")
}

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}

func TestExample02_Invoke(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("ex02", scc)

	// Init
	//checkInit(t, stub, [][]byte{[]byte("")})

	// Invoke Init Ledger
	checkInvoke(t, stub, [][]byte{[]byte("initLedger"), []byte("")})

	// Invoke Init Ledger
	checkInvoke(t, stub, [][]byte{[]byte("createLab"), []byte("M"), []byte("48"), []byte("Leawood"), []byte("Douglas"), []byte("KS"), []byte("Mouth Swab"), []byte("100"), []byte("2018:07:03:10:00"), []byte("active")})

	// Invoke Query All Labs
	checkInvoke(t, stub, [][]byte{[]byte("queryAllLabs"), []byte("")})

	// Invoke Query All Labs
	checkInvoke(t, stub, [][]byte{[]byte("queryStateResults"), []byte("KS")})

}
