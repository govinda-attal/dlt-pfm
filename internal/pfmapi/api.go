package pfmapi

import (
	"github.com/govinda-attal/dlt-pfm/pkg/core/status"
	"github.com/govinda-attal/dlt-pfm/pkg/pfm"
)

type AssetSrv struct {
}

func NewAssetSrv( /* inject dependencies (typically instance of channel client or something similar)*/ ) *AssetSrv {
	return &AssetSrv{}
}

func (as *AssetSrv) Save(asset *pfm.Asset) (*pfm.Asset, error) {
	return nil, status.ErrNotImplemented
}

func (as *AssetSrv) Get(key string) (*pfm.Asset, error) {
	return nil, status.ErrNotImplemented
}
