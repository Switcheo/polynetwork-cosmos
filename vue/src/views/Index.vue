<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="Switcheo.polynetworkcosmos.btcx" type="DenomInfo" :fields="[ ['creator', 1, 'string'] , ['denom', 2, 'string'] , ['assetHash', 3, 'string'] , ['totalSupply', 4, 'string'] , ['redeemScript', 5, 'string'] , ['redeemScriptHash', 6, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
