package types

import (
	fmt "fmt"

	polycommon "github.com/polynetwork/poly/common"
)

const (
	BtcChainID = uint64(1)
)

type ToBTCArgs struct {
	ToBtcAddress []byte
	Amount       uint64
	RedeemScript []byte
}

func (this *ToBTCArgs) Serialization(sink *polycommon.ZeroCopySink) error {
	sink.WriteVarBytes(this.ToBtcAddress)
	sink.WriteUint64(this.Amount)
	sink.WriteVarBytes(this.RedeemScript)
	return nil
}

func (this *ToBTCArgs) Deserialization(source *polycommon.ZeroCopySource) error {
	toBtcAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("ToBTCArgs deserialize toBtcAddress error")
	}
	amt, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("ToBTCArgs deserialize Amount error")
	}
	redeemScript, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("ToBTCArgs deserialize redeemScript error")
	}

	this.ToBtcAddress = toBtcAddress
	this.Amount = amt
	this.RedeemScript = redeemScript
	return nil
}

type BTCArgs struct {
	ToBtcAddress []byte
	Amount       uint64
}

func (this *BTCArgs) Serialization(sink *polycommon.ZeroCopySink) error {
	sink.WriteVarBytes(this.ToBtcAddress)
	sink.WriteUint64(this.Amount)
	return nil
}

func (this *BTCArgs) Deserialization(source *polycommon.ZeroCopySource) error {
	toBtcAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("ToBTCArgs deserialize toBtcAddress error")
	}
	amt, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("ToBTCArgs deserialize Amount error")
	}
	this.ToBtcAddress = toBtcAddress
	this.Amount = amt
	return nil
}
