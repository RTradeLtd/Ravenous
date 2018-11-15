package ravenous

import (
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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected status 200 got %v", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", string(bodyBytes))
	return nil
}
