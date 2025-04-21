package main

import (
	"log"

	"vehicleauto/contracts"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	VehicleContract := new(contracts.VehicleBlacklistContract)

	chaincode, err := contractapi.NewChaincode(VehicleContract)

	if err != nil {
		log.Panicf("Could not create chaincode : %v", err)
	}

	err = chaincode.Start()

	if err != nil {
		log.Panicf("Failed to start chaincode : %v", err)
	}
}