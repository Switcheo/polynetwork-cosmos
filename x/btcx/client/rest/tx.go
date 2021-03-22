package rest

import (
	"net/http"
	"strconv"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createDenomRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Denom        string       `json:"denom"`
	RedeemScript string       `json:"redeemScript"`
}

func createDenomInfoHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createDenomRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		parsedDenom := req.Denom

		parsedRedeemScript := req.RedeemScript

		msg := types.NewMsgCreateDenom(
			req.BaseReq.From,
			parsedDenom,
			parsedRedeemScript,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
