package ravenous_test

import (
	"testing"

	ravenous "github.com/RTradeLtd/Ravenous"
)

const (
	endpoint = "http://192.168.1.233:8766"
)

func TestGetBlockchainInfo(t *testing.T) {
	if err := ravenous.GetBlockchainInfo(endpoint); err != nil {
		t.Fatal(err)
	}
}
