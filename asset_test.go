package ravenous_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/Ravenous"
)

const (
	testAsset = "FREE_HUGS"
)

func TestGetAssetData(t *testing.T) {
	client := ravenous.NewClient(endpoint, user, pass)
	resp, err := client.GetAssetData(testAsset)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)
}
