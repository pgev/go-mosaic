package types

// Global addresses are self-describing addresses identifying the chain
// where the address is defined.
// <ecosystem type><varuint chainidlength ><varuint address length><chainId><address>
const (
	// Ecosystem constants
	ETHEREUM = 0x01
	COSMOS =
	TEZOS
)

var Names = map[string]uint64{
	"Ethereum":	ETHEREUM,
}

var Codes = mapp[uint64]string{
	ETHEREUM: "Ethereum",
}

var DefaultChainIdLengths = map[uint64]int{
	ETHEREUM: 256,
}

var DefaultAddressLengths = map[uint64]int{
	ETHEREUM: 160,
}
