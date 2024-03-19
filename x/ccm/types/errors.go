package types

// DONTCOVER

import (
	fmt "fmt"
	"reflect"

	errorsmod "cosmossdk.io/errors"
)

// x/ccm module sentinel errors
var (
	ErrMarshalSpecificType     = errorsmod.Register(ModuleName, 1, "ErrMarshalSpecificType")
	ErrUnmarshalSpecificType   = errorsmod.Register(ModuleName, 2, "ErrUnmarshalSpecificType")
	ErrProcessCrossChainTxType = errorsmod.Register(ModuleName, 3, "ErrProcessCrossChainTxType")
	ErrVerifyToCosmosTxType    = errorsmod.Register(ModuleName, 4, "ErrVerifyToCosmosTxType")

	ErrMsgProcessCrossChainTxType = errorsmod.Register(ModuleName, 5, "ErrMsgProcessCrossChainTxType")
	ErrMsgCreateCrossChainTxType  = errorsmod.Register(ModuleName, 6, "ErrMsgCreateCrossChainTxType")
	ErrCheckModuleContractType    = errorsmod.Register(ModuleName, 7, "ErrCheckModuleContractType")
	// this line is used by starport scaffolding # ibc/errors
)

func ErrMarshalSpecificTypeFail(o interface{}, err error) error {
	return errorsmod.Wrap(ErrMarshalSpecificType, fmt.Sprintf("Marshal type: %s, Error: %s", reflect.TypeOf(o).String(), err.Error()))
}

func ErrUnmarshalSpecificTypeFail(o interface{}, err error) error {
	return errorsmod.Wrap(ErrUnmarshalSpecificType, fmt.Sprintf("Umarshal type: %s, Error: %s", reflect.TypeOf(o).String(), err.Error()))
}

func ErrProcessCrossChainTx(reason string) error {
	return errorsmod.Wrapf(ErrProcessCrossChainTxType, "Reason: %s", reason)
}

func ErrVerifyToCosmosTx(reason string) error {
	return errorsmod.Wrapf(ErrVerifyToCosmosTxType, "Reason: %s", reason)
}

func ErrMsgProcessCrossChainTx(reason string) error {
	return errorsmod.Wrapf(ErrMsgProcessCrossChainTxType, "Reason: %s", reason)
}

func ErrMsgCreateCrossChainTx(reason string) error {
	return errorsmod.Wrapf(ErrMsgCreateCrossChainTxType, "Reason: %s", reason)
}

func ErrCheckModuleContract(reason string) error {
	return errorsmod.Wrapf(ErrCheckModuleContractType, "Reason: %s", reason)
}
