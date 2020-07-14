package types

const (
	// Ethereum address has expected length of 20 bytes
	EthereumAddressLength = 20
)

type Address [EthereumAddressLength]byte
