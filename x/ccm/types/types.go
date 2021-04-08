package types

import "strconv"

// ChainID is this chain's ID, to be overriden in makefile
var ChainID string

// GetChainID gets the chain ID for this chain in poly network
func GetChainID() uint64 {
	parsedChainID, err := strconv.ParseUint(ChainID, 10, 64)
	if err != nil {
		panic(err)
	}
	if parsedChainID == 0 {
		panic("Invalid chain ID")
	}
	return parsedChainID
}
