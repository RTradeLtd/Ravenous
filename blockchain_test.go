package ravenous_test

import (
	"testing"

	ravenous "github.com/RTradeLtd/Ravenous"
)

const (
	endpoint = "http://192.168.1.233:8766"
	user     = "user"
	pass     = "user"
)

func TestGetBlockchainInfo(t *testing.T) {
	client := ravenous.NewClient(endpoint, user, pass)
	if err := client.GetBlockchainInfo(); err != nil {
		t.Fatal(err)
	}
}
