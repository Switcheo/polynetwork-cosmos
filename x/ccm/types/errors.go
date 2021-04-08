package types

// DONTCOVER

import (
	fmt "fmt"
	"reflect"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/ccm module sentinel errors
var (
	ErrMarshalSpecificType     = sdkerrors.Register(ModuleName, 1, "ErrMarshalSpecificType")
	ErrUnmarshalSpecificType   = sdkerrors.Register(ModuleName, 2, "ErrUnmarshalSpecificType")
	ErrProcessCrossChainTxType = sdkerrors.Register(ModuleName, 3, "ErrProcessCrossChainTxType")
	ErrVerifyToCosmosTxType    = sdkerrors.Register(ModuleName, 4, "ErrVerifyToCosmosTxType")

	ErrMsgProcessCrossChainTxType = sdkerrors.Register(ModuleName, 5, "ErrMsgProcessCrossChainTxType")
	ErrMsgCreateCrossChainTxType  = sdkerrors.Register(ModuleName, 6, "ErrMsgCreateCrossChainTxType")
	ErrCheckModuleContractType    = sdkerrors.Register(ModuleName, 7, "ErrCheckModuleContractType")
	// this line is used by starport scaffolding # ibc/errors
)

func ErrMarshalSpecificTypeFail(o interface{}, err error) error {
	return sdkerrors.Wrap(ErrMarshalSpecificType, fmt.Sprintf("Marshal type: %s, Error: %s", reflect.TypeOf(o).String(), err.Error()))
}

func ErrUnmarshalSpecificTypeFail(o interface{}, err error) error {
	return sdkerrors.Wrap(ErrUnmarshalSpecificType, fmt.Sprintf("Umarshal type: %s, Error: %s", reflect.TypeOf(o).String(), err.Error()))
}

func ErrProcessCrossChainTx(reason string) error {
	return sdkerrors.Wrapf(ErrProcessCrossChainTxType, "Reason: %s", reason)
}

func ErrVerifyToCosmosTx(reason string) error {
	return sdkerrors.Wrapf(ErrVerifyToCosmosTxType, "Reason: %s", reason)
}

func ErrMsgProcessCrossChainTx(reason string) error {
	return sdkerrors.Wrapf(ErrMsgProcessCrossChainTxType, "Reason: %s", reason)
}

func ErrMsgCreateCrossChainTx(reason string) error {
	return sdkerrors.Wrapf(ErrMsgCreateCrossChainTxType, "Reason: %s", reason)
}

func ErrCheckModuleContract(reason string) error {
	return sdkerrors.Wrapf(ErrCheckModuleContractType, "Reason: %s", reason)
}
