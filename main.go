package main

import (
	"fmt"
	"log"

	"github.com/KimDaeikk/filmountain-oracle/rpc/filecoin_rpc"
	"github.com/KimDaeikk/filmountain-oracle/types"
)

func init() {
	// 전역변수 초기화
	err := types.Init()
	if err != nil {
		log.Fatalf("Failed to initialize global variables: %v", err)
	}
}

// SectorInfo 구조체
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

func main() {
	// 전역변수 리소스 정리
	defer types.CleanUp()

	// 최신 Tipset 가져오기
	tipset, err := filecoin_rpc.GetChainHead()
	if err != nil {
		log.Printf("Error getting latest Tipset: %s\n", err)
	}

	fmt.Printf("%+v\n", tipset.Cids)
	// performStateMinerSectorsCall(tipset)
}

// func performStateMinerSectorsCall(tipset *Tipset) {
// 	// JSON-RPC 요청 생성
// 	reqBody, err := json.Marshal(JSONRPCRequest{
// 		Jsonrpc: "2.0",
// 		Method:  "Filecoin.StateMinerSectors",
// 		Params: []interface{}{
// 			"f01083914", // Miner ID
// 			nil,         // Filter bitfield, nil to include all sectors
// 			tipset.Cids,
// 		},
// 		ID: 1,
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// HTTP POST 요청 생성
// 	req, err := http.NewRequest("POST", *types.LotusRpcUrl, bytes.NewBuffer(reqBody))
// 	if err != nil {
// 		panic(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", *types.AuthToken)

// 	// HTTP 클라이언트 생성
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	// 응답 디코딩
// 	var rpcResp JSONRPCResponse
// 	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
// 		panic(err)
// 	}

// 	if rpcResp.Error != nil {
// 		fmt.Printf("RPC Error: %d - %s\n", rpcResp.Error.Code, rpcResp.Error.Message)
// 		return
// 	}

// 	var sectors []SectorInfo
// 	if err := json.Unmarshal(rpcResp.Result, &sectors); err != nil {
// 		panic(err)
// 	}

// 	// 결과 출력
// 	for _, sector := range sectors {
// 		fmt.Printf("Sector Number: %d\n", sector.SectorNumber)
// 		fmt.Printf("Seal Proof: %d\n", sector.SealProof)
// 		fmt.Printf("Sealed CID: %s\n", sector.SealedCID["/"])
// 		fmt.Printf("Deal IDs: %v\n", sector.DealIDs)
// 		fmt.Printf("Activation: %d\n", sector.Activation)
// 		fmt.Printf("Expiration: %d\n", sector.Expiration)
// 		fmt.Printf("Deal Weight: %s\n", sector.DealWeight)
// 		fmt.Printf("Verified Deal Weight: %s\n", sector.VerifiedDealWeight)
// 		fmt.Printf("Initial Pledge: %s FIL\n", attoFILToFIL(sector.InitialPledge))
// 		fmt.Printf("Expected Day Reward: %s FIL\n", attoFILToFIL(sector.ExpectedDayReward))
// 		fmt.Printf("Expected Storage Pledge: %s FIL\n", attoFILToFIL(sector.ExpectedStoragePledge))
// 		fmt.Printf("Replaced Sector Age: %d\n", sector.ReplacedSectorAge)
// 		fmt.Printf("Replaced Day Reward: %s FIL\n", attoFILToFIL(sector.ReplacedDayReward))
// 		fmt.Printf("Sector Key CID: %v\n", sector.SectorKeyCID)
// 		fmt.Printf("Simple QA Power: %v\n", sector.SimpleQAPower)
// 		fmt.Println("-----------------------------------------------------")
// 	}
// }

// func attoFILToFIL(attoFIL string) string {
// 	value := new(big.Int)
// 	value.SetString(attoFIL, 10) // attoFIL 값 파싱
// 	fil := new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(1e18))
// 	return fil.Text('f', 18) // 18 소수점 자릿수로 출력
// }
