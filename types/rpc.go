package types

import "encoding/json"

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
