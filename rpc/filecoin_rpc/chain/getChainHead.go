package chain

import (
	"encoding/json"

	"github.com/KimDaeikk/filmountain-oracle/rpc"
	"github.com/KimDaeikk/filmountain-oracle/types"
	"github.com/pkg/errors"
)

func GetChainHead() (*types.ChainHead, error) {
	jsonRpcResponse, err := rpc.ExecuteRpcCall("ChainHead", nil)
	if err != nil {
		errors.Wrap(err, "GetChainHead")
	}

	var chainHead types.ChainHead
	if err := json.Unmarshal(jsonRpcResponse.Result, &chainHead); err != nil {
		return nil, err
	}

	return &chainHead, nil
}
