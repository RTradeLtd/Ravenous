package ravenous

// RPCRequest is a request to the rpc daemon
type RPCRequest struct {
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int64       `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
}
