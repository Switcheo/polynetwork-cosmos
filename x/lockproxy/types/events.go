/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package types

// Minting module event types
const (
	AttributeValueCategory = ModuleName

	EventTypeCreateLockProxy = "create_lock_proxy"
	AttributeKeyCreator      = "creator"
	AttributeKeyProxyHash    = "lock_proxy_hash"

	EventTypeCreateAndDelegateCoinToProxy = "create_and_delegate_coin_to_proxy"

	EventTypeBindProxy           = "bind_proxy_hash"
	AttributeKeyLockProxy        = "lock_proxy_hash"
	AttributeKeyToChainID        = "to_chain_id"
	AttributeKeyToChainProxyHash = "to_chain_proxy_hash"

	EventTypeBindAsset           = "bind_asset_hash"
	AttributeKeySourceAssetDenom = "source_asset_denom"
	AttributeKeySourceAssetHash  = "source_asset_hash"
	AttributeKeyToChainAssetHash = "to_chain_asset_hash"
	EventTypeLock                = "lock"
	AttributeKeyFromAddress      = "from_address"
	AttributeKeyToAddress        = "to_address"
	AttributeKeyAmount           = "amount"
	AttributeKeyFeeAddress       = "fee_address"
	AttributeKeyFeeAmount        = "fee_amount"
	AttributeKeyNonce            = "nonce"

	AttributeKeyFromChainID = "from_chain_id"

	EventTypeUnlock              = "unlock"
	AttributeKeyFromContractHash = "from_contract_hash"
	AttributeKeyToAssetDenom     = "to_asset_denom"

	EventTypeRegisterAsset      = "register_asset"
	AttributeKeyToContractHash  = "to_contract_hash"
	AttributeKeyAssetHash       = "asset_hash"
	AttributeKeyNativeAssetHash = "native_asset_hash"
)
