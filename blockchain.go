package ravenous

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//'{"jsonrpc": "1.0", "id":"curltest", "method": "getblockchaininfo", "params": [] }'

func GetBlockchainInfo(endpoint string) error {
	rpcRequest := RPCRequest{
		Method:  "getblockchaininfo",
		Params:  nil,
		ID:      time.Now().UnixNano(),
		JSONRPC: "1.0",
	}
	payloadBuffer := &bytes.Buffer{}
	jsonEncoder := json.NewEncoder(payloadBuffer)
	if err := jsonEncoder.Encode(rpcRequest); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", endpoint, payloadBuffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	// add basic auth
	req.SetBasicAuth("user", "user")
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("unauthorized access")
	}
	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("bad request")
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", string(bodyBytes))
	return nil
}
