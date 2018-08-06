package pfm

import (
	"github.com/govinda-attal/dlt-pfm/pkg/core/status"
)

type Asset struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //unique name of the asset.
}

type AssetAPI interface {
	Save(asset *Asset) (*Asset, error)
	Get(id string) (*Asset, error)
}

// Following are message(s) that can be used to pass to-fro to chaincode.
// These are simple json bindings or can be protobuf if need be.

type SaveRq struct {
	Asset *Asset `json:"asset"`
}

type SaveRs struct {
	Status *status.ServiceStatus `json:"status"`
}

type GetRs struct {
	Status *status.ServiceStatus `json:"status"`
	Asset  *Asset                `json:"asset,omitempty"`
}
