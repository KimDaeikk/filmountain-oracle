package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"github.com/KimDaeikk/filmountain-oracle/types"
)

func ExecuteRpcCall(_method string, _parameters []interface{}) (types.JSONRPCResponse, error) {
	requestBody, err := generateRequestBody(_method, _parameters)
	if err != nil {
		return types.JSONRPCResponse{}, errors.Wrap(err, "ExecuteRpcCall")
	}
	postRequest, err := generatePOSTRequest(requestBody)
	if err != nil {
		return types.JSONRPCResponse{}, errors.Wrap(err, "ExecuteRpcCall")
	}
	httpResponse, err := sendHTTPRequest(postRequest)
	if err != nil {
		return types.JSONRPCResponse{}, errors.Wrap(err, "ExecuteRpcCall")
	}
	jsonRPCResponse, err := decodeResponse(httpResponse)
	if err != nil {
		return types.JSONRPCResponse{}, errors.Wrap(err, "ExecuteRpcCall")
	}

	return jsonRPCResponse, nil
}

func generateRequestBody(_method string, _parameters []interface{}) ([]byte, error) {
	body, err := json.Marshal(types.JSONRPCRequest{
		Jsonrpc: "2.0",
		Method:  "Filecoin." + _method,
		Params:  _parameters,
		ID:      1,
	})
	if err != nil {
		return nil, errors.Wrap(err, "generateRequestBody: fail to generate request body")
	}
	return body, nil
}

func generatePOSTRequest(_requestBody []byte) (*http.Request, error) {
	postRequest, err := http.NewRequest("POST", *types.LotusRpcUrl, bytes.NewBuffer(_requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "generatePOSTRequest: fail to generate post request")
	}
	postRequest.Header.Set("Content-Type", "application/json")
	postRequest.Header.Set("Authorization", *types.AuthToken)
	return postRequest, nil
}

func sendHTTPRequest(_postRequest *http.Request) (*http.Response, error) {
	response, err := types.Client.Do(_postRequest)
	if err != nil {
		return nil, errors.Wrap(err, "sendHTTPRequest: fail to send http request")
	}
	return response, nil
}

func decodeResponse(_response *http.Response) (types.JSONRPCResponse, error) {
	defer _response.Body.Close()
	var rpcResp types.JSONRPCResponse
	if err := json.NewDecoder(_response.Body).Decode(&rpcResp); err != nil {
		return types.JSONRPCResponse{}, errors.Wrap(err, "decodeResponse: fail to decode json response")
	}

	if rpcResp.Error != nil {
		return types.JSONRPCResponse{}, errors.Errorf("decodeResponse: RPC Error: %d - %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return rpcResp, nil
}
