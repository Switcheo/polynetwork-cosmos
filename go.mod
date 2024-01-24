module github.com/Switcheo/polynetwork-cosmos

go 1.15

require (
	cosmossdk.io/api v0.3.1
	cosmossdk.io/simapp v0.0.0-20230608160436-666c345ad23d
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cometbft/cometbft v0.37.2
	github.com/cometbft/cometbft-db v0.8.0
	github.com/cosmos/cosmos-sdk v0.47.5
	github.com/cosmos/gogoproto v1.4.10
	github.com/cosmos/ibc-go/v7 v7.3.1
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	github.com/polynetwork/poly v0.0.0-20200710095239-0596a3d7afe5
	github.com/rakyll/statik v0.1.7
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	google.golang.org/genproto/googleapis/api v0.0.0-20230629202037-9506855d4529
	google.golang.org/grpc v1.56.2
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/Switcheo/cosmos-sdk v0.47.5-0.20240119065259-675e01adc46f

replace github.com/btcsuite/btcd => github.com/btcsuite/btcd v0.22.2
