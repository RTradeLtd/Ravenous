package ravenous

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//'{"jsonrpc": "1.0", "id":"curltest", "method": "getblockchaininfo", "params": [] }'

// GetBlockchainInfo is used to fetch basic blockchain information
func (c *Client) GetBlockchainInfo() error {
	payloadBuffer, err := FormatRPCRequest("getblockchaininfo", nil)
	req, err := http.NewRequest("POST", c.URL, payloadBuffer)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	// add basic auth
	req.SetBasicAuth(c.User, c.Pass)
	resp, err := c.HC.Do(req)
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
