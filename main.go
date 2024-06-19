package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/KimDaeikk/filmountain-oracle/types"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	lotusRpcUrl := os.Getenv("LOTUS_RPC_URL")
	authToken := "Bearer " + os.Getenv("AUTH_TOKEN")
	types.LotusRpcUrl = &lotusRpcUrl
	types.AuthToken = &authToken

	if types.LotusRpcUrl == nil || types.AuthToken == nil {
		log.Fatalf("Environment variables not set properly")
	}
}

// JSON-RPC 요청 구조체
type JSONRPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

// JSON-RPC 응답 구조체
type JSONRPCResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
	ID      int             `json:"id"`
}

// RPC 에러 구조체
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Tipset 구조체
type Tipset struct {
	Cids []map[string]string `json:"Cids"`
	// 기타 필드 추가 가능
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

// t0118000
func main() {
	// 최신 Tipset 가져오기
	tipset, err := getLatestTipset()
	if err != nil {
		log.Fatalf("Error getting latest Tipset: %s", err)
	}

	performStateMinerSectorsCall(tipset)
}

func getLatestTipset() (*Tipset, error) {
	// JSON-RPC 요청 생성
	reqBody, err := json.Marshal(JSONRPCRequest{
		Jsonrpc: "2.0",
		Method:  "Filecoin.ChainHead",
		Params:  []interface{}{},
		ID:      1,
	})
	if err != nil {
		return nil, err
	}

	// HTTP POST 요청 생성
	req, err := http.NewRequest("POST", *types.LotusRpcUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", *types.AuthToken)

	// HTTP 클라이언트 생성
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 응답 디코딩
	var rpcResp JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, err
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC Error: %d - %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	// Tipset 결과 디코딩
	var tipset Tipset
	if err := json.Unmarshal(rpcResp.Result, &tipset); err != nil {
		return nil, err
	}

	return &tipset, nil
}

func performStateMinerSectorsCall(tipset *Tipset) {
	// JSON-RPC 요청 생성
	reqBody, err := json.Marshal(JSONRPCRequest{
		Jsonrpc: "2.0",
		Method:  "Filecoin.StateMinerSectors",
		Params: []interface{}{
			"t0120833", // Miner ID
			nil,        // Filter bitfield, nil to include all sectors
			tipset.Cids,
		},
		ID: 1,
	})
	if err != nil {
		panic(err)
	}

	// HTTP POST 요청 생성
	req, err := http.NewRequest("POST", *types.LotusRpcUrl, bytes.NewBuffer(reqBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", *types.AuthToken)

	// HTTP 클라이언트 생성
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 응답 디코딩
	var rpcResp JSONRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		panic(err)
	}

	if rpcResp.Error != nil {
		fmt.Printf("RPC Error: %d - %s\n", rpcResp.Error.Code, rpcResp.Error.Message)
		return
	}

	var sectors []SectorInfo
	if err := json.Unmarshal(rpcResp.Result, &sectors); err != nil {
		panic(err)
	}

	// 결과 출력
	for _, sector := range sectors {
		fmt.Printf("Sector Number: %d\n", sector.SectorNumber)
		fmt.Printf("Seal Proof: %d\n", sector.SealProof)
		fmt.Printf("Sealed CID: %s\n", sector.SealedCID["/"])
		fmt.Printf("Deal IDs: %v\n", sector.DealIDs)
		fmt.Printf("Activation: %d\n", sector.Activation)
		fmt.Printf("Expiration: %d\n", sector.Expiration)
		fmt.Printf("Deal Weight: %s\n", sector.DealWeight)
		fmt.Printf("Verified Deal Weight: %s\n", sector.VerifiedDealWeight)
		fmt.Printf("Initial Pledge: %s FIL\n", attoFILToFIL(sector.InitialPledge))
		fmt.Printf("Expected Day Reward: %s FIL\n", attoFILToFIL(sector.ExpectedDayReward))
		fmt.Printf("Expected Storage Pledge: %s FIL\n", attoFILToFIL(sector.ExpectedStoragePledge))
		fmt.Printf("Replaced Sector Age: %d\n", sector.ReplacedSectorAge)
		fmt.Printf("Replaced Day Reward: %s FIL\n", attoFILToFIL(sector.ReplacedDayReward))
		fmt.Printf("Sector Key CID: %v\n", sector.SectorKeyCID)
		fmt.Printf("Simple QA Power: %v\n", sector.SimpleQAPower)
		fmt.Println("-----------------------------------------------------")
	}
}

func attoFILToFIL(attoFIL string) string {
	value := new(big.Int)
	value.SetString(attoFIL, 10) // attoFIL 값 파싱
	fil := new(big.Float).Quo(new(big.Float).SetInt(value), big.NewFloat(1e18))
	return fil.Text('f', 18) // 18 소수점 자릿수로 출력
}
