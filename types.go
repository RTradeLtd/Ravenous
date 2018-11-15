package ravenous

import (
	"bytes"
	"encoding/json"
	"time"
)

// RPCRequest is a request to the rpc daemon
type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int64       `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
}

// FormatRPCRequest is used to format our RPC request, and generate encoded payload data
func FormatRPCRequest(method string, params interface{}) (*bytes.Buffer, error) {
	rpcRequest := RPCRequest{
		Method:  method,
		Params:  params,
		ID:      time.Now().UnixNano(),
		JSONRPC: "1.0",
	}
	payloadBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(payloadBuffer)
	if err := jsonEncoder.Encode(rpcRequest); err != nil {
		return nil, err
	}
	return payloadBuffer, nil
}
