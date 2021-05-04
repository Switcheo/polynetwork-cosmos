package types

// Btcx module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeCreateAsset         = "create_asset"
	AttributeKeySourceAssetDenom = "source_asset_denom"
	AttributeKeyFromAssetHash    = "from_asset_hash"
	AttributeKeyRedeemScript     = "redeem_script"

	EventTypeBindAsset      = "bind_asset"
	AttributeKeyCreator     = "creator"
	AttributeKeyToChainID   = "to_chain_id"
	AttributeKeyToAssetHash = "to_asset_hash"

	EventTypeLockAsset      = "lock_asset"
	AttributeKeyFromAddress = "from_address"
	AttributeKeyToAddress   = "to_address"
	AttributeKeyAmount      = "amount"

	EventTypeUnlock = "unlock"
)
