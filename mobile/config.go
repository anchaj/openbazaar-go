package mobile

import (
	"github.com/op/go-logging"
)

var stdoutLogFormat = logging.MustStringFormatter(
	`%{color:reset}%{color}%{time:15:04:05.000} [%{shortfunc}] [%{level}] %{message}`,
)

var logger logging.Backend

// NodeConfig type definition includes path for data directory, authentication token, Testnet flag, user-agent, and trusted peer for wallet
type NodeConfig struct {

	// Path for the node's data directory
	RepoPath string

	// An API authentication. Authentication is turned on if not "".
	AuthenticationToken string

	// Use testnet
	Testnet bool

	// Add a custom user-agent
	UserAgent string

	// Add a trusted peer for the wallet
	WalletTrustedPeer string

	// Processes to disable
	DisableWallet        bool
	DisableExchangerates bool
}
