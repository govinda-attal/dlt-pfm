package app

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	"github.com/govinda-attal/dlt-pfm/pkg/core/status"
	"github.com/govinda-attal/dlt-pfm/pkg/pfm"
)

type AssetOnChain struct {
	stub shim.ChaincodeStubInterface
}

func NewOnChainService(stub shim.ChaincodeStubInterface) *AssetOnChain {
	return &AssetOnChain{stub}
}

func (ac *AssetOnChain) Save(asset *pfm.Asset) error {
	b, err := json.Marshal(asset)
	if err != nil {
		return status.ErrInternal.WithError(err)
	}
	err = ac.stub.PutState(asset.Name, b)
	if err != nil {
		return status.ErrInternal.WithError(err)
	}
	return nil
}

func (ac *AssetOnChain) Get(key string) (*pfm.Asset, error) {
	b, err := ac.stub.GetState(key)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}
	if b == nil {
		return nil, status.ErrNotFound
	}

	var asset pfm.Asset
	err = json.Unmarshal(b, &asset)
	if err != nil {
		return nil, status.ErrInternal.WithError(err)
	}
	return &asset, nil
}
