package types

type ChainHead struct {
	Cids    []map[string]string `json:"Cids"`
	Blocks  []interface{}       `json:"Blocks"`
	Heights int                 `json:"Height"`
}

type SectorInfo struct {
	SectorNumber          int               `json:"SectorNumber"`
	SealProof             int               `json:"SealProof"`
	SealedCID             map[string]string `json:"SealedCID"`
	DealIDs               []int             `json:"DealIDs"`
	Activation            int               `json:"Activation"`
	Expiration            int               `json:"Expiration"`
	DealWeight            string            `json:"DealWeight"`
	VerifiedDealWeight    string            `json:"VerifiedDealWeight"`
	InitialPledge         string            `json:"InitialPledge"`
	ExpectedDayReward     string            `json:"ExpectedDayReward"`
	ExpectedStoragePledge string            `json:"ExpectedStoragePledge"`
	ReplacedSectorAge     int               `json:"ReplacedSectorAge"`
	ReplacedDayReward     string            `json:"ReplacedDayReward"`
	SectorKeyCID          map[string]string `json:"SectorKeyCID"`
	SimpleQAPower         bool              `json:"SimpleQAPower"`
}

// NodeStatus
type SyncStatus struct {
	Epoch  uint `json:"Epoch"`
	Behind uint `json:"Behind"`
}

type PeerStatus struct {
	PeersToPublishMsgs   uint `json:"PeersToPublishMsgs"`
	PeersToPublishBlocks uint `json:"PeersToPublishBlocks"`
}

type ChainStatus struct {
	BlocksPerTipsetLast100      float64 `json:"BlocksPerTipsetLast100"`
	BlocksPerTipsetLastFinality float64 `json:"BlocksPerTipsetLastFinality"`
}

type NodeStatus struct {
	SyncStatus  SyncStatus  `json:"SyncStatus"`
	PeerStatus  PeerStatus  `json:"PeerStatus"`
	ChainStatus ChainStatus `json:"ChainStatus"`
}
