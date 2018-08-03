/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"

	"github.com/govinda-attal/dlt-pfm/cmd/pfm-cc/app"
	"github.com/govinda-attal/dlt-pfm/pkg/core/status"
	"github.com/govinda-attal/dlt-pfm/pkg/pfm"
)

// SimpleAsset implements a simple chaincode to manage an asset.
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result []byte
	var err error
	if fn == "set" {
		result, err = set(stub, args)
	} else { // assume 'get' even if fn is nil
		result, err = get(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success(result)
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, status.ErrBadRequest
	}

	var rq pfm.SaveRq
	err := json.Unmarshal([]byte(args[0]), &rq)
	if err != nil {
		return nil, status.ErrBadRequest.WithError(err)
	}

	srv := app.NewOnChainService(stub)
	err = srv.Save(rq.Asset)
	if err != nil {
		return nil, err
	}

	rs := pfm.SaveRs{Status: &status.Success}
	b, _ := json.Marshal(&rs)
	return b, nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, status.ErrBadRequest
	}
	srv := app.NewOnChainService(stub)
	asset, err := srv.Get(args[0])
	if err != nil {
		return nil, err
	}
	rs := pfm.GetRs{Status: &status.Success, Asset: asset}

	b, _ := json.Marshal(rs)
	return b, nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
