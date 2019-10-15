package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type AssCC struct {
}

func (*AssCC) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetStringArgs()
	if len(args) == 0 {
		return shim.Error("22222222")
	}

	var a = args[0]
	var av = args[1]
	var b = args[2]
	var bv = args[3]

	if err := stub.PutState(a, []byte(av)); err != nil {
		return shim.Error("33333333")
	}

	if err := stub.PutState(b, []byte(bv)); err != nil {
		return shim.Error("33333333")
	}

	return shim.Success([]byte("初始化成功"))
}

func (*AssCC) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	f, args := stub.GetFunctionAndParameters()
	var (
		result string
		err    error
	)

	switch f {
	case "find":
		{
			return find(args, stub)
		}
	case "payment":
		{
			return payment(args, stub)
		}
	case "del":
		{
			result, err = set(args, stub)
		}
	case "set":
		{
			result, err = set(args, stub)
		}
	case "get":
		{
			result, err = get(args, stub)
		}

	}

	if err != nil {
		return shim.Error("44444444444")
	}

	return shim.Success([]byte(result))
}

func find(args []string, stub shim.ChaincodeStubInterface) peer.Response {

	bytes, _ := stub.GetState(args[0])
	return shim.Success(bytes)
}

func payment(args []string, stub shim.ChaincodeStubInterface) peer.Response {

	s := args[0]
	t := args[1]
	v := args[2]

	sv := stub.GetState(s)
	tv := stub.GetState(t)

}

func set(args []string, stub shim.ChaincodeStubInterface) (string, error) {

	if len(args) == 0 {
		return "", fmt.Errorf("555555555555")
	}

	if err := stub.PutState(args[0], []byte(args[1])); err != nil {
		return "", fmt.Errorf("6666666666")
	}

	return string(args[0]), nil
}

func get(args []string, stub shim.ChaincodeStubInterface) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("7777777")
	}

	result, err := stub.GetState(args[0])
	if err != nil {
		return "", fmt.Errorf("8888888")
	}

	if result == nil {
		return "", fmt.Errorf("9999999")
	}

	return string(result), nil

}

func main() {
	if err := shim.Start(new(AssCC)); err != nil {
		fmt.Println("11111111")
	}
}
