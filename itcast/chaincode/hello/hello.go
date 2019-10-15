package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type HelloChaincode struct {

}

func (*HelloChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	f, args := stub.GetFunctionAndParameters()
	if len(args) ==0 {
		return shim.Error("指定了错误的参数数量")
	}

	fmt.Println("保存数据",f)

	if err := stub.PutState(args[0], []byte(args[0])); err != nil {
		return shim.Error("保存数据时发生错误")
	}

	return shim.Success([]byte("初始化成功"))
}

func (*HelloChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	f, str:= stub.GetFunctionAndParameters()
	if f == "query" {
		return query(str,stub)
	}
	return shim.Error("无指定函数")
}

func main() {
	err := shim.Start(new(HelloChaincode))
	if  err != nil {
		fmt.Println("链码启动失败")
	}
}

func query(args []string,stub shim.ChaincodeStubInterface)peer.Response   {
	if len(args) != 1 {
		return shim.Error("指定了错误的参数数量")
	}

	result, err := stub.GetState(args[0])
	if  err != nil {
		return shim.Error("查询错误")
	}

	if result ==  nil{
		return shim.Error("查无结果")
	}
	return shim.Success(result)
}
