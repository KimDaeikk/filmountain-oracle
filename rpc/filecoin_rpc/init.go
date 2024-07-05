package filecoin_rpc

import (
	"github.com/KimDaeikk/filmountain-oracle/rpc/filecoin_rpc/chain"
	// "github.com/KimDaeikk/filmountain-oracle/rpc/filecoin_rpc/node"
	"github.com/KimDaeikk/filmountain-oracle/types"
)

// -=-=-=-=-==-=-=- CHAIN -=-=-=-=-==-=-=-
func GetChainHead() (*types.Tipset, error) {
	return chain.GetChainHead()
}

// -=-=-=-=-==-=-=- NODE -=-=-=-=-==-=-=-
func GetNodeStatus() {
	// return node.GetNodeStatus()
}

// -=-=-=-=-==-=-=- STATE -=-=-=-=-==-=-=-
