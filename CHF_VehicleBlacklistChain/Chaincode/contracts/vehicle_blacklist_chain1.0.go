package contracts

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VehicleBlacklistContract manages vehicle blacklisting logic
type VehicleBlacklistContract struct {
	contractapi.Contract
}

type Vehicle struct {
	VehicleType       string `json:"vehicleType"`
	VehicleId         string `json:"vehicleId"`
	Model             string `json:"model"`
	Color             string `json:"color"`
	OwnerName         string `json:"ownername"`
	DateOfManufacture string `json:"dateOfManufacture"`
	Status            string `json:"status"` // e.g. "blacklisted" or "active"
	Reason            string `json:"reason"`
	ReportedBy        string `json:"reportedBy"`
	ReportedAt        string `json:"reportedAt"`
	RegisteredBy      string `json:"registeredBy"`
}

// ========== Vehicle: Check Existence ==========
func (v *VehicleBlacklistContract) VehicleExists(ctx contractapi.TransactionContextInterface, vehicleId string) (bool, error) {
	data, err := ctx.GetStub().GetState(vehicleId)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %v", err)
	}
	return data != nil, nil
}

// ========== Manufacturer: Create Vehicle ==========
func (c *VehicleBlacklistContract) CreateVehicle(ctx contractapi.TransactionContextInterface, vehicleId string, model string, color string,ownername string, dateOfManufacture string, registeredBy string) (string, error) {
	clientOrgID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return "", err
	}

	if clientOrgID != "owner-auto-com" {
		return "", fmt.Errorf("user under following MSPID: %v can't perform this action", clientOrgID)
	}

	exists, err := c.VehicleExists(ctx, vehicleId)
	if err != nil {
		return "", err
	} else if exists {
		return "", fmt.Errorf("the VEHICLE, %s already exists", vehicleId)
	}

	vehicle := Vehicle{
		VehicleType:       "vehicle",
		VehicleId:         vehicleId,
		Model:             model,
		Color:             color,
		OwnerName:         ownername,
		DateOfManufacture: dateOfManufacture,
		RegisteredBy:      registeredBy,
		Status:            "active",
	}

	bytes, err := json.Marshal(vehicle)
	if err != nil {
		return "", err
	}

	err = ctx.GetStub().PutState(vehicleId, bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Vehicle %s created successfully", vehicleId), nil
}

// ========== Dealer: Report Vehicle as Stolen ==========
func (v *VehicleBlacklistContract) ReportStolen(ctx contractapi.TransactionContextInterface, vehicleId string, reason string) error {
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return err
	}

	if clientMSPID != "dealer-auto-com" {
		return fmt.Errorf("only Dealer can report stolen vehicles, your MSPID: %s", clientMSPID)
	}

	vehicleBytes, err := ctx.GetStub().GetState(vehicleId)
	if err != nil || vehicleBytes == nil {
		return fmt.Errorf("vehicle %s not found", vehicleId)
	}

	var vehicle Vehicle
	err = json.Unmarshal(vehicleBytes, &vehicle)
	if err != nil {
		return err
	}

	vehicle.Status = "blacklisted"
	vehicle.Reason = reason
	vehicle.ReportedBy = clientMSPID
	vehicle.ReportedAt = time.Now().UTC().Format(time.RFC3339)

	updatedBytes, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	err = ctx.GetStub().SetEvent("VehicleBlacklistedEvent", updatedBytes)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(vehicleId, updatedBytes)
}

// ========== MVD: Remove From Blacklist ==========
func (v *VehicleBlacklistContract) RemoveFromBlacklist(ctx contractapi.TransactionContextInterface, vehicleId string) error {
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return err
	}

	if clientMSPID != "mvd-auto-com" {
		return fmt.Errorf("only MVD can remove vehicles from blacklist, your MSPID: %s", clientMSPID)
	}

	vehicleBytes, err := ctx.GetStub().GetState(vehicleId)
	if err != nil || vehicleBytes == nil {
		return fmt.Errorf("vehicle %s not found", vehicleId)
	}

	var vehicle Vehicle
	err = json.Unmarshal(vehicleBytes, &vehicle)
	if err != nil {
		return err
	}

	vehicle.Status = "active"
	vehicle.Reason = ""
	vehicle.ReportedBy = ""
	vehicle.ReportedAt = ""

	updatedBytes, err := json.Marshal(vehicle)
	if err != nil {
		return err
	}

	err = ctx.GetStub().SetEvent("VehicleRecoveredEvent", updatedBytes)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(vehicleId, updatedBytes)
}

// ========== Utility: Get Vehicle History ==========
func (v *VehicleBlacklistContract) GetVehicleHistory(ctx contractapi.TransactionContextInterface, vehicleId string) ([]*Vehicle, error) {
	historyIterator, err := ctx.GetStub().GetHistoryForKey(vehicleId)
	if err != nil {
		return nil, fmt.Errorf("failed to get history for vehicle %s: %v", vehicleId, err)
	}
	defer historyIterator.Close()

	var history []*Vehicle

	for historyIterator.HasNext() {
		response, err := historyIterator.Next()
		if err != nil {
			return nil, err
		}

		var vehicle Vehicle
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &vehicle)
			if err != nil {
				return nil, err
			}
			history = append(history, &vehicle)
		}
	}

	return history, nil
}
