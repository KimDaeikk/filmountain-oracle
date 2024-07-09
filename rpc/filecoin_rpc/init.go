package filecoin_rpc

import (
	"github.com/KimDaeikk/filmountain-oracle/rpc/filecoin_rpc/chain"
	// "github.com/KimDaeikk/filmountain-oracle/rpc/filecoin_rpc/client"
	// "github.com/KimDaeikk/filmountain-oracle/rpc/filecoin_rpc/state"
	"github.com/KimDaeikk/filmountain-oracle/types"
)

// -=-=-=-=-==-=-=- CHAIN -=-=-=-=-==-=-=-
func GetChainHead() (*types.ChainHead, error) {
	return chain.GetChainHead()
}

// -=-=-=-=-==-=-=- CLIENT -=-=-=-=-==-=-=-
func GetNodeStatus() (*types.NodeStatus, error) {
	// return client.GetNodeStatus()
}

// -=-=-=-=-==-=-=- STATE -=-=-=-=-==-=-=-
