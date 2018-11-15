package ravenous

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetAssetData is used to get information about a particular asset
func (c *Client) GetAssetData(assetName string) (*GetAssetDataResponse, error) {
	payloadBuffer, err := FormatRPCRequest("getassetdata", map[string]string{
		"asset_name": assetName,
	})
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.URL, payloadBuffer)
	if err != nil {
		return nil, err
	}
	// set header
	req.Header.Add("Content-Type", "application/json")
	// set basic auth
	req.SetBasicAuth(c.User, c.Pass)
	resp, err := c.HC.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected status 200 got %v", resp.StatusCode)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var intf struct {
		Result GetAssetDataResponse
	}
	if err = json.Unmarshal(bodyBytes, &intf); err != nil {
		return nil, err
	}
	return &intf.Result, nil
}
