package types

import "strconv"

// ChainId is this chain's ID, to be overriden in makefile
var ChainId string

// GetChainId gets the chain ID for this chain in poly network
func GetChainId() uint64 {
	parsedChainId, err := strconv.ParseUint(ChainId, 10, 64)
	if err != nil {
		panic(err)
	}
	if parsedChainId == 0 {
		panic("Invalid chain ID")
	}
	return parsedChainId
}
