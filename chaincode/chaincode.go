package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Patient struct {
	PatientID 	int		'json:"patientId"'
	FirstName 	string 	'json:"firstName"'
	LastName 	string 	'json:"lastName"'
	Age 		int 	'json:"age"'
}

type Doctor struct {
	DoctorID	int		'json:"doctorId"'
	Name 		string	'json:"doctorName"'
	Hospital	string	'json:"hospital"'
}

type Chaincode struct {

}

func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	fcn, params := stub.GetFunctionAndParameters()
	var result []byte
	var err error
	if fcn == "CreatePatient" {
		result, err = cc.CreatePatient(stub)
	} else if fcn == "CreateDoctor" {
		result, err = cc.CreateDoctor(stub, params)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(result)
}

func (cc *Chaincode) CreatePatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
		return nil, fmt.Errorf("Failed to create Patient: The number of arguments is incorrect")
	}

	var newPatient = Patient{PatientID: args[1], FirstName: args[2], LastName: args[3], Age: args[4]}

	newPatientAsBytes, err := json.Marshal(newPatient)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Patient")
	}

	stub.PutState(args[0], newPatientAsBytes)
	return newPatientAsBytes, nil
}

func (cc *Chaincode) CreateDoctor(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 3 {
		return nil, fmt.Errorf("Failed to create Doctor: The number of arguments is incorrect")
	}

	var newDoctor = Doctor{DoctorID: args[1], Name: args[2], Hospital: args[3]}

	newDoctorAsBytes, err := json.Marshal(newDoctor)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Doctor")
	}

	stub.PutState(args[0], newDoctorAsBytes)
	return newDoctorAsBytes, nil
}


func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}