package types

import (
	"encoding/hex"
	"math/big"

	fmt "fmt"

	polycommon "github.com/polynetwork/poly/common"
)

type TxArgs struct {
	FromAssetHash []byte
	ToAssetHash   []byte
	ToAddress     []byte
	Amount        *big.Int
	FeeAmount     *big.Int
	FeeAddress    []byte
	FromAddress   []byte
	Nonce         *big.Int
}

func (txargs *TxArgs) Serialize(sink *polycommon.ZeroCopySink, intLen int) error {
	sink.WriteVarBytes(txargs.FromAssetHash)
	sink.WriteVarBytes(txargs.ToAssetHash)
	sink.WriteVarBytes(txargs.ToAddress)
	paddedAmountBs, err := PadFixedBytes(txargs.Amount, intLen)
	if err != nil {
		return fmt.Errorf("serialization error:%v", err)
	}
	sink.WriteBytes(paddedAmountBs)
	paddedFeeAmountBs, err := PadFixedBytes(txargs.FeeAmount, intLen)
	if err != nil {
		return fmt.Errorf("serialization error:%v", err)
	}
	sink.WriteBytes(paddedFeeAmountBs)
	sink.WriteVarBytes(txargs.FeeAddress)
	sink.WriteVarBytes(txargs.FromAddress)
	paddedNonceBs, err := PadFixedBytes(txargs.Nonce, intLen)
	if err != nil {
		return fmt.Errorf("serialization error:%v", err)
	}
	sink.WriteBytes(paddedNonceBs)
	return nil
}

func (txargs *TxArgs) Deserialize(source *polycommon.ZeroCopySource, intLen int) error {
	fromAssetHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("deserialize fromAssetHash error")
	}
	toAssetHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("deserialize toAssetHash error")
	}
	toAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("deserialize ToAddress error")
	}
	paddedAmountBs, eof := source.NextBytes(uint64(intLen))
	if eof {
		return fmt.Errorf("deserialize Amount error")
	}
	amount, err := UnpadFixedBytes(paddedAmountBs, intLen)
	if err != nil {
		return fmt.Errorf("deserialization error:%v", err)
	}
	paddedFeeAmountBs, eof := source.NextBytes(uint64(intLen))
	if eof {
		return fmt.Errorf("deserialize FeeAmount error")
	}
	feeAmount, err := UnpadFixedBytes(paddedFeeAmountBs, intLen)
	if err != nil {
		return fmt.Errorf("deserialization error:%v", err)
	}
	feeAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("deserialize FeeAddress error")
	}
	fromAddress, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("deserialize FromAddress error")
	}
	paddedNonceBs, eof := source.NextBytes(uint64(intLen))
	if eof {
		return fmt.Errorf("deserialize Nonce error")
	}
	nonce, err := UnpadFixedBytes(paddedNonceBs, intLen)
	if err != nil {
		return fmt.Errorf("deserialization error:%v", err)
	}

	txargs.FromAssetHash = fromAssetHash
	txargs.ToAssetHash = toAssetHash
	txargs.ToAddress = toAddress
	txargs.Amount = amount
	txargs.FeeAmount = feeAmount
	txargs.FeeAddress = feeAddress
	txargs.FromAddress = fromAddress
	txargs.Nonce = nonce
	return nil
}

type RegisterAssetTxArgs struct {
	AssetHash       []byte
	NativeAssetHash []byte
}

func (this *RegisterAssetTxArgs) Serialize(sink *polycommon.ZeroCopySink) error {
	sink.WriteVarBytes(this.AssetHash)
	sink.WriteVarBytes(this.NativeAssetHash)
	return nil
}

func (this *RegisterAssetTxArgs) Deserialize(source *polycommon.ZeroCopySource) error {
	assetHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("TxArgs deserialize assetHash error")
	}

	nativeAssetHash, eof := source.NextVarBytes()
	if eof {
		return fmt.Errorf("TxArgs deserialize nativeAssetHash error")
	}

	this.AssetHash = assetHash
	this.NativeAssetHash = nativeAssetHash
	return nil
}

func PadFixedBytes(bigint *big.Int, intBsLen int) ([]byte, error) {
	ret := make([]byte, intBsLen)
	if bigint.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("PadFixedBytes doesnot support negative big.Int, but got:%s", bigint.String())
	}
	bigBs := bigint.Bytes()
	if len(bigBs) > intBsLen || (len(bigBs) == intBsLen && bigBs[0]>>7 == 1) {
		return nil, fmt.Errorf("PadFixedBytes only support maximum 2**255-1 big.Int, but got:%s", bigint.String())
	}
	copy(ret[:len(bigBs)], make([]byte, len(bigBs)))
	copy(ret[intBsLen-len(bigBs):], bigBs)
	return ToArrayReverse(ret), nil
}

func UnpadFixedBytes(paddedBs []byte, intBsLen int) (*big.Int, error) {
	if len(paddedBs) != intBsLen {
		return nil, fmt.Errorf("UnpadFixedBytes only support 32 bytes value, but got:%s", hex.EncodeToString(paddedBs))
	}
	nonZeroPos := intBsLen - 1
	for i := nonZeroPos; i >= 0; i-- {
		p := paddedBs[i]
		if p != 0x0 {
			nonZeroPos = i
			break
		}
	}
	if nonZeroPos == intBsLen-1 && paddedBs[intBsLen-1]>>7 == 1 {
		return nil, fmt.Errorf("UnpadFixedBytes only support 32 bytes nonnegative value, but got:%s", hex.EncodeToString(paddedBs))
	}

	return big.NewInt(0).SetBytes(ToArrayReverse(paddedBs[:nonZeroPos+1])), nil
}

func ToArrayReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}
