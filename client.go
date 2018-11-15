package ravenous

import "net/http"

const (
	// DefaultUser is our default user for authentication
	DefaultUser = "user"
	// DefaultPass is our default password for authentication
	DefaultPass = "pass"
)

// Client represents a connection to the RavenCoin rpc connection
type Client struct {
	URL  string
	User string
	Pass string
	HC   *http.Client
}

// NewClient is used to instantiate our RavenCoin rpc client
func NewClient(endpoint, user, pass string) *Client {
	return &Client{
		URL:  endpoint,
		User: user,
		Pass: pass,
		HC:   &http.Client{},
	}
}
