package chain

import (
	"encoding/json"

	"github.com/KimDaeikk/filmountain-oracle/rpc"
	"github.com/KimDaeikk/filmountain-oracle/types"
	"github.com/pkg/errors"
)

func GetChainHead() (*types.Tipset, error) {
	jsonRpcResponse, err := rpc.ExecuteRpcCall("ChainHead")
	if err != nil {
		errors.Wrap(err, "GetChainHead")
	}

	var tipset types.Tipset
	if err := json.Unmarshal(jsonRpcResponse.Result, &tipset); err != nil {
		return nil, err
	}

	return &tipset, nil
}
