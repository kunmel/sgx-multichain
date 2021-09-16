module sgx-multichain/sgx-multichain-linker

go 1.15

require (
	ethereumLink v0.0.0
	fabricLink v0.0.0

)

replace (
	ethereumLink => ../sgx-multichain-eth
	fabricLink => ../sgx-multichain-fab
	github.com/ethereum/go-ethereum v0.0.0 => ../../github.com/ethereum/go-ethereum
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.3
)