package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMsgProcessCrossChainTx returns a new MsgProcessCrossChainTx
func NewMsgProcessCrossChainTx(submitter string, fromChainID uint64, proof string, header string, headerProof string, curHeader string) *MsgProcessCrossChainTx {
	return &MsgProcessCrossChainTx{submitter, fromChainID, proof, header, headerProof, curHeader}
}

// Route implements Msg
func (msg *MsgProcessCrossChainTx) Route() string {
	return RouterKey
}

// Type implements Msg
func (msg *MsgProcessCrossChainTx) Type() string {
	return "ProcessCrossChainTx"
}

// GetSigners implements Msg
func (msg *MsgProcessCrossChainTx) GetSigners() []sdk.AccAddress {
	submitter, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{submitter}
}

// GetSignBytes implements Msg
func (msg *MsgProcessCrossChainTx) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (msg *MsgProcessCrossChainTx) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Submitter)
	if err != nil {
		return ErrMsgProcessCrossChainTx(fmt.Sprintf("Invalid submitter (%s), error: %v", msg.Submitter, err))
	}
	if len(msg.Proof) == 0 {
		return ErrMsgProcessCrossChainTx(fmt.Sprintf("Proof should not be empty"))
	}
	if len(msg.Header) == 0 {
		return ErrMsgProcessCrossChainTx(fmt.Sprintf("Header should not be empty"))
	}
	return nil
}
