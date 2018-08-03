package pfm

import (
	"github.com/govinda-attal/dlt-pfm/pkg/core/status"
)

type Asset struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //unique name of the asset.
}

type SaveRq struct {
	Asset *Asset `json:"asset"`
}

type SaveRs struct {
	//Status ...
	Status *status.ServiceStatus `json:"status"`
}

type GetRs struct {
	// Status ...
	Status *status.ServiceStatus `json:"status"`
	// Asset ...
	Asset *Asset `json:"asset,omitempty"`
}
