package litecoin

import (
	"blockbook/bchain/coins/btc"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
)

const (
	MainnetMagic wire.BitcoinNet = 0xdbb6c0fb
	TestnetMagic wire.BitcoinNet = 0xf1c8d2fd
	RegtestMagic wire.BitcoinNet = 0xdab5bffa
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
	RegtestParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = 48
	MainNetParams.ScriptHashAddrID = 5
	MainNetParams.Bech32HRPSegwit = "ltc"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = 111
	TestNetParams.ScriptHashAddrID = 196
	TestNetParams.Bech32HRPSegwit = "tltc"

	err := chaincfg.Register(&MainNetParams)
	if err == nil {
		err = chaincfg.Register(&TestNetParams)
	}
	if err != nil {
		panic(err)
	}
}

// LitecoinParser handle
type LitecoinParser struct {
	*btc.BitcoinParser
}

// NewLitecoinParser returns new LitecoinParser instance
func NewLitecoinParser(params *chaincfg.Params, c *btc.Configuration) *LitecoinParser {
	return &LitecoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Litecoin network,
// the regression test Litecoin network, the test Litecoin network and
// the simulation test Litecoin network, in this order
func GetChainParams(chain string) *chaincfg.Params {
	switch chain {
	case "test":
		return &TestNetParams
	case "regtest":
		return &RegtestParams
	default:
		return &MainNetParams
	}
}
