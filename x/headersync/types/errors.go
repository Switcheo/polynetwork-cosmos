package types

// DONTCOVER

import (
	fmt "fmt"
	"reflect"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/headersync module sentinel errors
var (
	ErrDeserializeHeaderFailType    = sdkerrors.Register(ModuleName, 1, "ErrDeserializeHeaderFailType")
	ErrSerializeHeaderFailType      = sdkerrors.Register(ModuleName, 2, "ErrSerializeHeaderFailType")
	ErrHeaderEmptyType              = sdkerrors.Register(ModuleName, 3, "ErrHeaderEmptyType")
	ErrGetKeyHeaderHashType         = sdkerrors.Register(ModuleName, 4, "ErrGetKeyHeaderHashType")
	ErrGetConsensusPeersFailType    = sdkerrors.Register(ModuleName, 5, "ErrGetConsensusPeersFailType")
	ErrBookKeeperNumErrType         = sdkerrors.Register(ModuleName, 6, "ErrBookKeeperNumErrType")
	ErrInvalidPublicKeyType         = sdkerrors.Register(ModuleName, 7, "ErrInvalidPublicKeyType")
	ErrUnmarshalSpecificType        = sdkerrors.Register(ModuleName, 8, "ErrUnmarshalSpecificType")
	ErrVerifyMultiSigFailType       = sdkerrors.Register(ModuleName, 9, "ErrVerifyMultiSignatureFailType")
	ErrDeserializeConsensusPeerType = sdkerrors.Register(ModuleName, 10, "ErrDeserializeConsensusPeerType")
	ErrSyncGenesisHeaderType        = sdkerrors.Register(ModuleName, 11, "ErrSyncGenesisHeaderType")
	ErrSyncBlockHeaderType          = sdkerrors.Register(ModuleName, 12, "ErrSyncBlockHeaderType")
	// this line is used by starport scaffolding # ibc/errors
)

func ErrSyncBlockHeader(operation string, chainID uint64, height uint32, err error) error {
	return sdkerrors.Wrapf(ErrSyncBlockHeaderType, "operation: %s, chainID: %d, height: %d,  Error :%s", operation, chainID, height, err.Error())
}
func ErrDeserializeHeader(err error) error {
	return sdkerrors.Wrap(ErrDeserializeHeaderFailType, fmt.Sprintf("Header deserialization Error:%s", err.Error()))
}

func ErrSerializeHeader(err error) error {
	return sdkerrors.Wrap(ErrSerializeHeaderFailType, fmt.Sprintf("Header serialization Error:%s", err.Error()))
}
func ErrHeaderEmpty(headerhash []byte) error {
	return sdkerrors.Wrap(ErrHeaderEmptyType, fmt.Sprintf("Header of headerHash: %x is empty", headerhash))
}
func ErrDeserializeConsensusPeer(err error) error {
	return sdkerrors.Wrap(ErrDeserializeConsensusPeerType, fmt.Sprintf("ConsensusPeer deserialization Error:%s", err.Error()))
}

func ErrMarshalSpecificTypeFail(o interface{}, err error) error {
	return sdkerrors.Wrap(ErrUnmarshalSpecificType, fmt.Sprintf("Marshal type: %s, Error: %s", reflect.TypeOf(o).String(), err.Error()))
}
func ErrUnmarshalSpecificTypeFail(o interface{}, err error) error {
	return sdkerrors.Wrap(ErrUnmarshalSpecificType, fmt.Sprintf("Umarshal type: %s, Error: %s", reflect.TypeOf(o).String(), err.Error()))
}

func ErrGetKeyHeaderHash(reason string) error {
	return sdkerrors.Wrap(ErrGetKeyHeaderHashType, fmt.Sprintf("Reason: %s", reason))
}

func ErrGetConsensusPeers(chainID uint64) error {
	return sdkerrors.Wrap(ErrGetConsensusPeersFailType, fmt.Sprintf("For chainID: %d, Get consensus peers empty error", chainID))
}

func ErrBookKeeperNum(headerBookKeeperNum int, consensusNodeNum int) error {
	return sdkerrors.Wrap(ErrBookKeeperNumErrType, fmt.Sprintf("Header Bookkeepers number: %d must more than 2/3 consensus node number: %d", headerBookKeeperNum, consensusNodeNum))
}

func ErrInvalidPublicKey(pubkey string) error {
	return sdkerrors.Wrap(ErrInvalidPublicKeyType, fmt.Sprintf("Invalid pubkey: %s", pubkey))
}

func ErrVerifyMultiSigFail(err error, height uint32) error {
	return sdkerrors.Wrap(ErrVerifyMultiSigFailType, fmt.Sprintf("Verify multi signature Error: %s of height: %d", err.Error(), height))
}

func ErrSyncGenesisHeader(reason string) error {
	return sdkerrors.Wrapf(ErrSyncGenesisHeaderType, fmt.Sprintf("Reason: %s", reason))
}
