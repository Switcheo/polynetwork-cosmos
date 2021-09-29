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

	EventTypeCreateCrossChainTx = "make_from_cosmos_proof"
	AttributeKeyStatus          = "status"
	AttributeKeyFromAddress     = "from_address"
	AttributeKeyFromContract    = "from_contract"
	AttributeKeyToChainId       = "to_chain_id"
	AttributeCrossChainId       = "cross_chain_id"
	AttributeKeyTxParamHash     = "make_tx_param_hash"
	AttributeKeyMakeTxParam     = "make_tx_param"

	EventTypeVerifyToCosmosProof                        = "verify_to_cosmos_proof"
	AttributeKeyMerkleValueTxHash                       = "merkle_value:txhash"
	AttributeKeyMerkleValueMakeTxParamTxHash            = "merkle_value:make_tx_param:txhash"
	AttributeKeyMerkleValueMakeTxParamToContractAddress = "merkle_value:make_tx_param:to_contract_address"
	AttributeKeyFromChainId                             = "from_chain_id"
)
