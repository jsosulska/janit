package token

import (
	"flag"
)

// Token checking has an order of precedence.
func findTokenInSet(bulkFlagSet *flag.FlagSet) (string, error) {
	// Check if token was passed to CLI

	// Check if token is in ENV

	// Check if token is in File

	// Pass
}
