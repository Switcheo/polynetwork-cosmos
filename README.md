# Polynetwork Cosmos

**polynetworkcosmos** is a developed for any Cosmos SDK based chain to integrate with [Poly Network](https://poly.network) conveniently. It is created using [Starport](https://github.com/tendermint/starport) and was ported from [cosmos-poly-module](https://github.com/polynetwork/cosmos-poly-module).

## Integration

This repo contains five modules.

```markdown
1. ccm: cross chain manager module
2. headersync: poly chain header synchronization module
3. lockproxy: proxy module for locking and unlocking asset between cross chain transaction
4. btcx: interoperable asset module for chains like BTC
5. ft: interoperable fungible token module without using lockproxy module *(Not Yet Ported)*
```

Please refer to [cosmos cross chain workflow documentation](https://github.com/polynetwork/docs/blob/master/cosmos/cosmos_cross_chain_workflow.md) for more info.

## Get started

```bash
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on multiple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
